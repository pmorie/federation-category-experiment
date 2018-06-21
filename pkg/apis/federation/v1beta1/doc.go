// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/pmorie/federation-category-experiment/pkg/apis/federation
// +k8s:defaulter-gen=TypeMeta
// +groupName=federation.pmorie.toy
package v1beta1 // import "github.com/pmorie/federation-category-experiment/pkg/apis/federation/v1beta1"
