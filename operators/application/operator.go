package operator

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/derekparker/delve/pkg/config"
	crv1 "github.com/enablecloud/kulbe/apis/cr/application/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

const maxRetries = 5

// Controller object
type Controller struct {
	clientset    kubernetes.Interface
	queue        workqueue.RateLimitingInterface
	informer     cache.SharedIndexInformer
	eventHandler Handler
}

// GetClient returns a k8s clientset to the request from inside of cluster
func GetClient() kubernetes.Interface {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("Can not get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Can not create kubernetes client: %v", err)
	}

	return clientset
}

func buildOutOfClusterConfig() (*rest.Config, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
	}
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

// GetClientOutOfCluster returns a k8s clientset to the request from outside of cluster
func GetClientOutOfCluster() kubernetes.Interface {
	config, err := buildOutOfClusterConfig()
	if err != nil {
		fmt.Println("Can not get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	return clientset
}

func Start(conf *config.Config, namespace string, cfg *rest.Config, eventHandler Handler) {
	kubeClient := GetClientOutOfCluster()

	stopCh := make(chan struct{})
	defer close(stopCh)

	// make a new config for our extension's API group, using the first config as a baseline
	appClient, _, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}

	appF := watchAppFolder(kubeClient, namespace, appClient, eventHandler)
	go appF.Run(stopCh)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)
	<-sigterm
}

func NewClient(cfg *rest.Config) (*rest.RESTClient, *runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	if err := crv1.AddToScheme(scheme); err != nil {
		return nil, nil, err
	}

	config := *cfg
	config.GroupVersion = &crv1.SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, nil, err
	}

	return client, scheme, nil
}

type ApplicationController struct {
	ApplicationClient *rest.RESTClient
	ApplicationScheme *runtime.Scheme
}

// Run starts the kubewatch controller
func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	fmt.Println("Starting kubewatch controller")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	fmt.Println("Kubewatch controller synced and ready")

	wait.Until(c.runWorker, time.Second, stopCh)
}

// HasSynced is required for the cache.Controller interface.
func (c *Controller) HasSynced() bool {
	return c.informer.HasSynced()
}

// LastSyncResourceVersion is required for the cache.Controller interface.
func (c *Controller) LastSyncResourceVersion() string {
	return c.informer.LastSyncResourceVersion()
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
		// continue looping
	}

}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.processItem(key.(EventQueued))
	if err == nil {
		// No error, reset the ratelimit counters
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		fmt.Println("Error processing %s (will retry): %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		fmt.Println("Error processing %s (giving up): %v", key, err)
		c.queue.Forget(key)
		utilruntime.HandleError(err)
	}

	keyD, quitD := c.queue.Get()
	if quitD {
		return false
	}
	defer c.queue.Done(keyD)

	return true
}

func (c *Controller) processItem(key EventQueued) error {
	obj, _, err := c.informer.GetIndexer().GetByKey(key.Key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}
	if key.Event == EventTypeCreate {
		c.eventHandler.ObjectCreated(obj)
		return nil
	}
	if key.Event == EventTypeUpdate {
		c.eventHandler.ObjectUpdated(key.Old, obj)
		return nil
	}
	if key.Event == EventTypeDelete {
		c.eventHandler.ObjectDeleted(key.Old)
		return nil
	}

	return nil
}

func watchAppFolder(clientkub kubernetes.Interface, namespace string, client *rest.RESTClient, eventHandler Handler) *Controller {

	//Define what we want to look for (Services)
	watchlist := cache.NewListWatchFromClient(client, "kapps", namespace, fields.Everything())

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	//	resyncPeriod := 30 * time.Minute

	//Setup an informer to call functions when the watchlist changes
	informer := cache.NewSharedIndexInformer(
		watchlist,
		&crv1.KApplication{},
		0, //Skip resync
		cache.Indexers{},
	)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Println("Create ", key)
				queue.Add(EventQueued{Event: EventTypeCreate, Key: key})
			}
		},
		UpdateFunc: func(old, new interface{}) {
			_, err := cache.DeletionHandlingMetaNamespaceKeyFunc(old)
			_, err = cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				//Workaround as queue is not take into similar object
				eventHandler.ObjectUpdated(old, new)
				//queue.Add(EventQueued{Event: EventTypeUpdate, Key: key, Old: old})
				//queue.Add(EventQueued{Event: EventTypeUpdate, Key: key, Old: old})

			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Println("Delete ", key)
				//Workaround as queue is not take into similar object
				//queue.Add(EventQueued{Event: EventTypeDelete, Key: key, Old: obj})
				eventHandler.ObjectDeleted(obj)
			}
		},
	})

	return &Controller{
		clientset:    clientkub,
		informer:     informer,
		queue:        queue,
		eventHandler: eventHandler,
	}
}
