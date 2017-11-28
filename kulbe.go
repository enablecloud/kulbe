package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	comm "github.com/enablecloud/kulbe/common"
	serv "github.com/enablecloud/kulbe/go-server-server/go"
	kubeappoperator "github.com/enablecloud/kulbe/operators/application"
	"k8s.io/kubernetes/pkg/api"
)

func main() {

	var eventHandler kubeappoperator.Handler

	kubeconfig := flag.String("kube-config", "", "Path to a kube config. Only required if out-of-cluster.")
	kubeContext := flag.String("kube-context", "", "name of the kubeconfig context to use.")
	namespace := flag.String("namespace", api.NamespaceAll, "Namespace managed by the controller (All by default).")

	flag.Parse() // Create the client config. Use kubeconfig if given, otherwise assume in-cluster.
	config, err := comm.BuildConfig(*kubeconfig)
	if err != nil {
		panic(err)
	}
	eventHandler = new(kubeappoperator.Default)
	var clientCluster = kubeappoperator.GetClient()
	if clientCluster == nil {
		clientCluster = kubeappoperator.GetClientOutOfCluster()
	}
	eventHandler.Init(config, clientCluster, *kubeContext, *kubeconfig)
	fmt.Println("Start with namespace : '" + *namespace + "' and config: '" + *kubeconfig + "'.")
	log.Printf("Server started")

	router := serv.NewRouter()

	go log.Fatal(http.ListenAndServe(":8080", router))
	kubeappoperator.Start(*namespace, config, eventHandler)

}
