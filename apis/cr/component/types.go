/**
 */
package application

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestType is a top-level type. A client is created for it.
type ApplicationType struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Status TestTypeStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TestTypeList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type ApplicationTypeList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []ApplicationType
}

type ApplicationTypeStatus struct {
	Status string
}
