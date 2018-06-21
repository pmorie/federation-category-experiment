package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the FederatedReplicaSet resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FederatedReplicaSetSpec defines the desired state of FederatedReplicaSet
type FederatedReplicaSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// FederatedReplicaSetStatus defines the observed state of FederatedReplicaSet
type FederatedReplicaSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedReplicaSet
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=federatedreplicasets
// +kubebuilder:categories=federation
type FederatedReplicaSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedReplicaSetSpec   `json:"spec,omitempty"`
	Status FederatedReplicaSetStatus `json:"status,omitempty"`
}
