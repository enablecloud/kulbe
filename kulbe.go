package main

import (
	"flag"
	"fmt"

	"github.com/derekparker/delve/pkg/config"
	kubeappoperator "github.com/enablecloud/kulbe/operators/application"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	env "k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/helm/helmpath"
	"k8s.io/kubernetes/pkg/api"
)

func main() {

	conf := &config.Config{}
	var eventHandler kubeappoperator.Handler

	helmhome := flag.String("helm-home", env.DefaultHelmHome, "location of your Helm config. Overrides $HELM_HOME.")
	kubeconfig := flag.String("kube-config", "", "Path to a kube config. Only required if out-of-cluster.")
	kubeContext := flag.String("kube-context", "", "name of the kubeconfig context to use.")
	debug := flag.Bool("debug-helm", false, "debug mode.")
	namespace := flag.String("namespace", api.NamespaceAll, "Namespace managed by the controller (All by default).")
	tillernamespace := flag.String("tiller-namespace", "kube-system", "Tiller Namespace for helm deployment.")
	tilleraddress := flag.String("tiller-host", "tiller-deploy:44134", "Tiller Address for helm deployment.")
	tillertunnel := flag.Bool("tiller-tunnel", false, "Tiller tunnel active ?.(Default: false)")
	flag.Parse() // Create the client config. Use kubeconfig if given, otherwise assume in-cluster.
	config, err := buildConfig(*kubeconfig)
	if err != nil {
		panic(err)
	}
	eventHandler = new(kubeappoperator.Default)
	//var homeLad = new(helmpath.Home)
	//homeLad = (helmpath.Home)(*helmhome)
	//homeLad = *helmhome

	eventHandler.Init(conf, config, kubeappoperator.GetClientOutOfCluster(), *tillernamespace, *tilleraddress, *tillertunnel, (helmpath.Home)(*helmhome), *debug, *kubeContext, *kubeconfig)
	fmt.Println("Start with namespace : '" + *namespace + "' and config: '" + *kubeconfig + "'.")
	kubeappoperator.Start(conf, *namespace, config, eventHandler)

}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
