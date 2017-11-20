package main

import (
	"flag"
	"fmt"

	"github.com/derekparker/delve/pkg/config"
	kubeappoperator "github.com/enablecloud/kulbe/operators/application"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("Start")
	conf := &config.Config{}
	var eventHandler kubeappoperator.Handler

	kubeconfig := flag.String("kubeconfig", "", "Path to a kube config. Only required if out-of-cluster.")
	flag.Parse() // Create the client config. Use kubeconfig if given, otherwise assume in-cluster.
	config, err := buildConfig(*kubeconfig)
	if err != nil {
		panic(err)
	}
	eventHandler = new(kubeappoperator.Default)
	eventHandler.Init(conf, kubeappoperator.GetClientOutOfCluster())

	kubeappoperator.Start(conf, config, eventHandler)

}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
