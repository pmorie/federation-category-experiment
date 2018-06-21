package v1beta1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: "federation.pmorie.toy", Version: "v1beta1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&FederatedReplicaSet{},
		&FederatedReplicaSetList{},
		&FederatedReplicaSetOverride{},
		&FederatedReplicaSetOverrideList{},
		&FederatedReplicaSetPlacement{},
		&FederatedReplicaSetPlacementList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSet `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSetOverride `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedReplicaSetPlacement `json:"items"`
}

// CRD Generation
func getFloat(f float64) *float64 {
	return &f
}

func getInt(i int64) *int64 {
	return &i
}

var (
	// Define CRDs for resources
	FederatedReplicaSetCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "federatedreplicasets.federation.pmorie.toy",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "federation.pmorie.toy",
			Version: "v1beta1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "FederatedReplicaSet",
				Plural: "federatedreplicasets",
				Categories: []string{
					"federation",
				},
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
						"status": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
					},
				},
			},
		},
	}
	// Define CRDs for resources
	FederatedReplicaSetOverrideCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "federatedreplicasetoverrides.federation.pmorie.toy",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "federation.pmorie.toy",
			Version: "v1beta1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "FederatedReplicaSetOverride",
				Plural: "federatedreplicasetoverrides",
				Categories: []string{
					"federation",
				},
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
						"status": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
					},
				},
			},
		},
	}
	// Define CRDs for resources
	FederatedReplicaSetPlacementCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "federatedreplicasetplacements.federation.pmorie.toy",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "federation.pmorie.toy",
			Version: "v1beta1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "FederatedReplicaSetPlacement",
				Plural: "federatedreplicasetplacements",
				Categories: []string{
					"federation",
				},
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
						"status": v1beta1.JSONSchemaProps{
							Type:       "object",
							Properties: map[string]v1beta1.JSONSchemaProps{},
						},
					},
				},
			},
		},
	}
)
