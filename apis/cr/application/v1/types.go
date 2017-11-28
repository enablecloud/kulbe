package v1

import (
	cp "github.com/enablecloud/kulbe/apis/cr/component/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ApplicationResourcePlural = "kapps"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KApplication `json:"items"`
}

type ApplicationSpec struct {
	Components cp.ComponentList `json:"components"`
}

type ApplicationStatus struct {
	State   ApplicationState `json:"state,omitempty"`
	Message string           `json:"message,omitempty"`
}

type ApplicationState string

const (
	ApplicationStateCreated   ApplicationState = "Created"
	ApplicationStateProcessed ApplicationState = "Processed"
)
