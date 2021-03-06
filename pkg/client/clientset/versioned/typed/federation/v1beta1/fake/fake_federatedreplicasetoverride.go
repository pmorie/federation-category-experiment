// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/pmorie/federation-category-experiment/pkg/apis/federation/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedReplicaSetOverrides implements FederatedReplicaSetOverrideInterface
type FakeFederatedReplicaSetOverrides struct {
	Fake *FakeFederationV1beta1
	ns   string
}

var federatedreplicasetoverridesResource = schema.GroupVersionResource{Group: "federation.pmorie.toy", Version: "v1beta1", Resource: "federatedreplicasetoverrides"}

var federatedreplicasetoverridesKind = schema.GroupVersionKind{Group: "federation.pmorie.toy", Version: "v1beta1", Kind: "FederatedReplicaSetOverride"}

// Get takes name of the federatedReplicaSetOverride, and returns the corresponding federatedReplicaSetOverride object, and an error if there is any.
func (c *FakeFederatedReplicaSetOverrides) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedreplicasetoverridesResource, c.ns, name), &v1beta1.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicaSetOverride), err
}

// List takes label and field selectors, and returns the list of FederatedReplicaSetOverrides that match those selectors.
func (c *FakeFederatedReplicaSetOverrides) List(opts v1.ListOptions) (result *v1beta1.FederatedReplicaSetOverrideList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedreplicasetoverridesResource, federatedreplicasetoverridesKind, c.ns, opts), &v1beta1.FederatedReplicaSetOverrideList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedReplicaSetOverrideList{}
	for _, item := range obj.(*v1beta1.FederatedReplicaSetOverrideList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedReplicaSetOverrides.
func (c *FakeFederatedReplicaSetOverrides) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedreplicasetoverridesResource, c.ns, opts))

}

// Create takes the representation of a federatedReplicaSetOverride and creates it.  Returns the server's representation of the federatedReplicaSetOverride, and an error, if there is any.
func (c *FakeFederatedReplicaSetOverrides) Create(federatedReplicaSetOverride *v1beta1.FederatedReplicaSetOverride) (result *v1beta1.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedreplicasetoverridesResource, c.ns, federatedReplicaSetOverride), &v1beta1.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicaSetOverride), err
}

// Update takes the representation of a federatedReplicaSetOverride and updates it. Returns the server's representation of the federatedReplicaSetOverride, and an error, if there is any.
func (c *FakeFederatedReplicaSetOverrides) Update(federatedReplicaSetOverride *v1beta1.FederatedReplicaSetOverride) (result *v1beta1.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedreplicasetoverridesResource, c.ns, federatedReplicaSetOverride), &v1beta1.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicaSetOverride), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedReplicaSetOverrides) UpdateStatus(federatedReplicaSetOverride *v1beta1.FederatedReplicaSetOverride) (*v1beta1.FederatedReplicaSetOverride, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedreplicasetoverridesResource, "status", c.ns, federatedReplicaSetOverride), &v1beta1.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicaSetOverride), err
}

// Delete takes name of the federatedReplicaSetOverride and deletes it. Returns an error if one occurs.
func (c *FakeFederatedReplicaSetOverrides) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedreplicasetoverridesResource, c.ns, name), &v1beta1.FederatedReplicaSetOverride{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedReplicaSetOverrides) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedreplicasetoverridesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedReplicaSetOverrideList{})
	return err
}

// Patch applies the patch and returns the patched federatedReplicaSetOverride.
func (c *FakeFederatedReplicaSetOverrides) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedreplicasetoverridesResource, c.ns, name, data, subresources...), &v1beta1.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicaSetOverride), err
}
