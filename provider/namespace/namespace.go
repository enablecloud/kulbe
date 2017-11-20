package namespace

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateNameSpace(clientkub kubernetes.Interface, namespace string) (*v1.Namespace, error) {
	nsSpec := &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}
	return clientkub.Core().Namespaces().Create(nsSpec)

}

func DeleteNameSpace(clientkub kubernetes.Interface, namespace string) error {
	deletePolicy := metav1.DeletePropagationBackground
	return clientkub.Core().Namespaces().Delete(namespace, &metav1.DeleteOptions{PropagationPolicy: &deletePolicy})

}
