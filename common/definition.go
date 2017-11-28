package common

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
	RestConfig      *rest.Config
	Clientkub       kubernetes.Interface
	TillerAddress   string
	TillerNamespace string
	TillerTunnel    bool
	Home            string
	DebugHelm       bool
	KubeContext     string
	KubeConfig      string
}

func BuildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
