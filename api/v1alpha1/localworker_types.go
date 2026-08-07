package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LocalWorkerSpec struct {
	// +optional
	Image string `json:"image,omitempty"`

	// +optional
	ConfigSecretRef string `json:"configSecretRef,omitempty"`
}

type LocalWorkerStatus struct {

	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type LocalWorker struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	Spec LocalWorkerSpec `json:"spec"`

	// +optional
	Status LocalWorkerStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

type LocalWorkerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []LocalWorker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LocalWorker{}, &LocalWorkerList{})
}
