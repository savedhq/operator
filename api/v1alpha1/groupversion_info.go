// +kubebuilder:object:generate=true
// +groupName=worker.saved.sh
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	SchemeGroupVersion = schema.GroupVersion{Group: "worker.saved.sh", Version: "v1alpha1"}

	GroupVersion = SchemeGroupVersion

	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	AddToScheme = SchemeBuilder.AddToScheme
)
