package operator

import (
	"fmt"
	"reflect"

	"github.com/derekparker/delve/pkg/config"
	crv1 "github.com/enablecloud/kulbe/apis/cr/application/v1"
	"github.com/enablecloud/kulbe/common"
	hlm "github.com/enablecloud/kulbe/provider/helm"
	nspace "github.com/enablecloud/kulbe/provider/namespace"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/helm/helmpath"
)

type EventType string

const (
	EventTypeCreate EventType = "Create"
	EventTypeUpdate EventType = "Update"
	EventTypeDelete EventType = "Delete"
)

type EventQueued struct {
	Event EventType `json:"state,omitempty"`
	Key   string    `json:"message,omitempty"`
}

type Handler interface {
	Init(c *config.Config, rest *rest.Config, kube kubernetes.Interface, tillernamespace string, tillerAddress string, tillertunnel bool, home helmpath.Home, debugHelm bool, kubeContext string, kubeConfig string) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
	config          *config.Config
	restConfig      *rest.Config
	clientkub       kubernetes.Interface
	tillerAddress   string
	tillerNamespace string
	tillerTunnel    bool
	home            helmpath.Home
	debugHelm       bool
	kubeContext     string
	kubeConfig      string
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(conf *config.Config, rest *rest.Config, clientb kubernetes.Interface, tillernamespace string, tillerAddress string, tillertunnel bool, home helmpath.Home, debugHelm bool, kubeContext string, kubeConfig string) error {
	d.config = conf
	d.restConfig = rest
	d.clientkub = clientb
	d.tillerAddress = tillerAddress
	d.tillerNamespace = tillernamespace
	d.tillerTunnel = tillertunnel
	d.home = home
	d.debugHelm = debugHelm
	d.kubeContext = kubeContext
	d.kubeConfig = d.kubeConfig
	return nil

}

func (d *Default) ObjectCreated(obj interface{}) {
	fmt.Println("Processing create to ObjectCreated ")
	//deploymentsClient := d.clientkub.AppsV1beta2().Deployments(v1.NamespaceDefault)
	fmt.Println(reflect.TypeOf(obj))
	objAppFolder, ok := obj.(*crv1.KApplication)
	if ok && objAppFolder != nil && reflect.TypeOf(obj).String() == "*v1.KApplication" {
		//create Namespace
		name := objAppFolder.Name
		namespace := objAppFolder.Namespace
		newNamespape := namespace + name
		nspace.CreateNameSpace(d.clientkub, newNamespape)
		for i, v := range objAppFolder.Spec.Components.Items {
			fmt.Println(v)
			fmt.Println(v.Spec)
			compName := v.Name
			helmName := v.Spec.HelmName
			helmVersion := v.Spec.Version
			newConf := common.Default{TillerAddress: d.tillerAddress, TillerNamespace: d.tillerNamespace, TillerTunnel: d.tillerTunnel, Config: d.config, RestConfig: d.restConfig, Clientkub: d.clientkub, Home: d.home, DebugHelm: d.debugHelm, KubeContext: d.kubeContext, KubeConfig: d.kubeConfig}
			hlm.InstallRelease(helmName, helmVersion, newNamespape, compName, &newConf)
			fmt.Println(helmName)
			fmt.Println(helmVersion)
			fmt.Println(i)

		}

	}

}

func (d *Default) ObjectDeleted(obj interface{}) {

	fmt.Println("Processing remove to ObjectDeleted ")
	//deploymentsClient := d.clientkub.AppsV1beta2().Deployments(v1.NamespaceDefault)
	objAppFolder, ok := obj.(*crv1.KApplication)
	if ok && objAppFolder != nil && reflect.TypeOf(obj).String() == "*v1.KApplication" {
		name := objAppFolder.Name
		namespace := objAppFolder.Namespace
		newNamespape := namespace + name
		for i, v := range objAppFolder.Spec.Components.Items {
			compName := v.Name
			helmName := v.Spec.HelmName
			helmVersion := v.Spec.Version
			newConf := common.Default{TillerAddress: d.tillerAddress, TillerNamespace: d.tillerNamespace, TillerTunnel: d.tillerTunnel, Config: d.config, RestConfig: d.restConfig, Clientkub: d.clientkub, Home: d.home, DebugHelm: d.debugHelm, KubeContext: d.kubeContext, KubeConfig: d.kubeConfig}

			hlm.DeleteRelease(compName, &newConf)
			fmt.Println(helmName)
			fmt.Println(helmVersion)
			fmt.Println(i)

		}
		nspace.DeleteNameSpace(d.clientkub, newNamespape)
	}
}

func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {
	fmt.Println("Processing update to ObjectCreated ")
}
