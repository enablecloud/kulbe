package operator

import (
	"fmt"
	"reflect"

	"github.com/derekparker/delve/pkg/config"
	crv1 "github.com/enablecloud/kulbe/apis/cr/application/v1"
	hlm "github.com/enablecloud/kulbe/provider/helm"
	nspace "github.com/enablecloud/kulbe/provider/namespace"

	"k8s.io/client-go/kubernetes"
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
	Init(c *config.Config, kube kubernetes.Interface) error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
	config        *config.Config
	clientkub     kubernetes.Interface
	tillerAddress string
}

// Init initializes handler configuration
// Do nothing for default handler
func (d *Default) Init(conf *config.Config, clientb kubernetes.Interface) error {
	d.config = conf
	d.clientkub = clientb
	return nil

}

func (d *Default) ObjectCreated(obj interface{}) {
	fmt.Println("Processing create to ObjectCreated ")
	//deploymentsClient := d.clientkub.AppsV1beta2().Deployments(v1.NamespaceDefault)
	fmt.Println(reflect.TypeOf(obj))
	objAppFolder, ok := obj.(*crv1.Application)
	if ok && objAppFolder != nil && reflect.TypeOf(obj).String() == "*v1.Application" {
		//create Namespace
		name := objAppFolder.Name
		namespace := objAppFolder.Namespace
		newNamespape := namespace + name
		nspace.CreateNameSpace(d.clientkub, newNamespape)
		for i, v := range objAppFolder.Spec.Components.Items {
			compName := v.Name
			helmName := v.Spec.HelmName
			helmVersion := v.Spec.Version
			hlm.InstallRelease(d.tillerAddress, helmName, helmVersion, newNamespape, compName)
			fmt.Println(helmName)
			fmt.Println(helmVersion)
			fmt.Println(i)

		}

	}

}

func (d *Default) ObjectDeleted(obj interface{}) {

	fmt.Println("Processing remove to ObjectDeleted ")
	//deploymentsClient := d.clientkub.AppsV1beta2().Deployments(v1.NamespaceDefault)
	objAppFolder, ok := obj.(*crv1.Application)
	if ok && objAppFolder != nil && reflect.TypeOf(obj).String() == "*v1.Application" {
		name := objAppFolder.Name
		namespace := objAppFolder.Namespace
		newNamespape := namespace + name
		for i, v := range objAppFolder.Spec.Components.Items {
			compName := v.Name
			helmName := v.Spec.HelmName
			helmVersion := v.Spec.Version
			hlm.DeleteRelease(d.tillerAddress, compName)
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
