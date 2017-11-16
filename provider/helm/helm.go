package helm

import (
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/kubernetes/pkg/kubelet/kubeletconfig/util/log"
)

var (
	TillerAddress = "tiller-deploy:44134"
)

func InstallRelease(address string, release string, namespace string, releasename string) *services.InstallReleaseResponse {
	var client = NewHelmImplementer(address)
	response, _ := client.InstallRelease(release, namespace, helm.ReleaseName(releasename))
	return response
}

func DeleteRelease(address string, releasename string) *services.UninstallReleaseResponse {
	var client = NewHelmImplementer(address)
	response, _ := client.DeleteRelease(releasename)
	return response
}

func NewHelmImplementer(address string) *helm.Client {
	if address == "" {
		address = TillerAddress
	} else {
		log.Infof("provider.helm: tiller address '%s' supplied", address)
	}

	return helm.NewClient(helm.Host(address))
}
