/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LocalWorkerSpec defines the desired state of LocalWorker.
//
// A LocalWorker is one customer-side backup worker: a Deployment running the
// local-worker image, configured entirely by the `config.yaml` it is handed.
// Placeholder shape only, the fields below are the minimum needed to express
// the intent, not the final API.
type LocalWorkerSpec struct {
	// image is the local-worker container image to run.
	// +optional
	Image string `json:"image,omitempty"`

	// configSecretRef names a Secret in the same namespace holding the worker's
	// `config.yaml` under that key. It carries a WorkOS org-scoped API key, so it
	// is a Secret rather than a ConfigMap.
	// +optional
	ConfigSecretRef string `json:"configSecretRef,omitempty"`
}

// LocalWorkerStatus defines the observed state of LocalWorker.
type LocalWorkerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the LocalWorker resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// LocalWorker is the Schema for the localworkers API
type LocalWorker struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of LocalWorker
	// +required
	Spec LocalWorkerSpec `json:"spec"`

	// status defines the observed state of LocalWorker
	// +optional
	Status LocalWorkerStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// LocalWorkerList contains a list of LocalWorker
type LocalWorkerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []LocalWorker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LocalWorker{}, &LocalWorkerList{})
}
