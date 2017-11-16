package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ComponentResourcePlural = "kapps"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ComponentSpec   `json:"spec"`
	Status            ComponentStatus `json:"status,omitempty"`
}

type ComponentSpec struct {
	HelmName string `json:"helmname"`
	Version  string `json:"helmversion"`
}

type ComponentStatus struct {
	State   ComponentState `json:"state,omitempty"`
	Message string         `json:"message,omitempty"`
}

type ComponentState string

const (
	ComponentStateCreated   ComponentState = "Created"
	ComponentStateProcessed ComponentState = "Processed"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Component `json:"items"`
}
