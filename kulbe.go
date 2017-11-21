package main

import (
	"flag"
	"fmt"

	"github.com/derekparker/delve/pkg/config"
	kubeappoperator "github.com/enablecloud/kulbe/operators/application"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/api"
)

func main() {

	conf := &config.Config{}
	var eventHandler kubeappoperator.Handler

	kubeconfig := flag.String("kubeconfig", "", "Path to a kube config. Only required if out-of-cluster.")
	namespace := flag.String("namespace", api.NamespaceAll, "Namespace managed by the controller (All by default).")
	flag.Parse() // Create the client config. Use kubeconfig if given, otherwise assume in-cluster.
	config, err := buildConfig(*kubeconfig)
	if err != nil {
		panic(err)
	}
	eventHandler = new(kubeappoperator.Default)
	eventHandler.Init(conf, kubeappoperator.GetClientOutOfCluster())
	fmt.Println("Start with namespace : '" + *namespace + "' and config: '" + *kubeconfig + "'.")
	kubeappoperator.Start(conf, *namespace, config, eventHandler)

}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
