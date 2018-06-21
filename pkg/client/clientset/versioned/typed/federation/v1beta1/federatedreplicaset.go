// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/pmorie/federation-category-experiment/pkg/apis/federation/v1beta1"
	scheme "github.com/pmorie/federation-category-experiment/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedReplicaSetsGetter has a method to return a FederatedReplicaSetInterface.
// A group's client should implement this interface.
type FederatedReplicaSetsGetter interface {
	FederatedReplicaSets(namespace string) FederatedReplicaSetInterface
}

// FederatedReplicaSetInterface has methods to work with FederatedReplicaSet resources.
type FederatedReplicaSetInterface interface {
	Create(*v1beta1.FederatedReplicaSet) (*v1beta1.FederatedReplicaSet, error)
	Update(*v1beta1.FederatedReplicaSet) (*v1beta1.FederatedReplicaSet, error)
	UpdateStatus(*v1beta1.FederatedReplicaSet) (*v1beta1.FederatedReplicaSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedReplicaSet, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedReplicaSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedReplicaSet, err error)
	FederatedReplicaSetExpansion
}

// federatedReplicaSets implements FederatedReplicaSetInterface
type federatedReplicaSets struct {
	client rest.Interface
	ns     string
}

// newFederatedReplicaSets returns a FederatedReplicaSets
func newFederatedReplicaSets(c *FederationV1beta1Client, namespace string) *federatedReplicaSets {
	return &federatedReplicaSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedReplicaSet, and returns the corresponding federatedReplicaSet object, and an error if there is any.
func (c *federatedReplicaSets) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedReplicaSet, err error) {
	result = &v1beta1.FederatedReplicaSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedReplicaSets that match those selectors.
func (c *federatedReplicaSets) List(opts v1.ListOptions) (result *v1beta1.FederatedReplicaSetList, err error) {
	result = &v1beta1.FederatedReplicaSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedReplicaSets.
func (c *federatedReplicaSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedReplicaSet and creates it.  Returns the server's representation of the federatedReplicaSet, and an error, if there is any.
func (c *federatedReplicaSets) Create(federatedReplicaSet *v1beta1.FederatedReplicaSet) (result *v1beta1.FederatedReplicaSet, err error) {
	result = &v1beta1.FederatedReplicaSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		Body(federatedReplicaSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedReplicaSet and updates it. Returns the server's representation of the federatedReplicaSet, and an error, if there is any.
func (c *federatedReplicaSets) Update(federatedReplicaSet *v1beta1.FederatedReplicaSet) (result *v1beta1.FederatedReplicaSet, err error) {
	result = &v1beta1.FederatedReplicaSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		Name(federatedReplicaSet.Name).
		Body(federatedReplicaSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedReplicaSets) UpdateStatus(federatedReplicaSet *v1beta1.FederatedReplicaSet) (result *v1beta1.FederatedReplicaSet, err error) {
	result = &v1beta1.FederatedReplicaSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		Name(federatedReplicaSet.Name).
		SubResource("status").
		Body(federatedReplicaSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedReplicaSet and deletes it. Returns an error if one occurs.
func (c *federatedReplicaSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedReplicaSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedreplicasets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedReplicaSet.
func (c *federatedReplicaSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedReplicaSet, err error) {
	result = &v1beta1.FederatedReplicaSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedreplicasets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
