package common

import (
	"github.com/derekparker/delve/pkg/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/helm/helmpath"
)

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
	Config          *config.Config
	RestConfig      *rest.Config
	Clientkub       kubernetes.Interface
	TillerAddress   string
	TillerNamespace string
	TillerTunnel    bool
	Home            helmpath.Home
	DebugHelm       bool
	KubeContext     string
	KubeConfig      string
}
