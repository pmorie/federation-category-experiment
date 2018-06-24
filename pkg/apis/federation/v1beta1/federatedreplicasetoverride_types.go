package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the FederatedReplicaSetOverride resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FederatedReplicaSetOverrideSpec defines the desired state of FederatedReplicaSetOverride
type FederatedReplicaSetOverrideSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// FederatedReplicaSetOverrideStatus defines the observed state of FederatedReplicaSetOverride
type FederatedReplicaSetOverrideStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedReplicaSetOverride
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=federatedreplicasetoverrides
// +kubebuilder:categories=federation
type FederatedReplicaSetOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedReplicaSetOverrideSpec   `json:"spec,omitempty"`
	Status FederatedReplicaSetOverrideStatus `json:"status,omitempty"`
}
