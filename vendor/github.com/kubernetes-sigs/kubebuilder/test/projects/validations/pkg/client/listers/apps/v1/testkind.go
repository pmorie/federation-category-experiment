// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/apis/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TestKindLister helps list TestKinds.
type TestKindLister interface {
	// List lists all TestKinds in the indexer.
	List(selector labels.Selector) (ret []*v1.TestKind, err error)
	// TestKinds returns an object that can list and get TestKinds.
	TestKinds(namespace string) TestKindNamespaceLister
	TestKindListerExpansion
}

// testKindLister implements the TestKindLister interface.
type testKindLister struct {
	indexer cache.Indexer
}

// NewTestKindLister returns a new TestKindLister.
func NewTestKindLister(indexer cache.Indexer) TestKindLister {
	return &testKindLister{indexer: indexer}
}

// List lists all TestKinds in the indexer.
func (s *testKindLister) List(selector labels.Selector) (ret []*v1.TestKind, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TestKind))
	})
	return ret, err
}

// TestKinds returns an object that can list and get TestKinds.
func (s *testKindLister) TestKinds(namespace string) TestKindNamespaceLister {
	return testKindNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TestKindNamespaceLister helps list and get TestKinds.
type TestKindNamespaceLister interface {
	// List lists all TestKinds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TestKind, err error)
	// Get retrieves the TestKind from the indexer for a given namespace and name.
	Get(name string) (*v1.TestKind, error)
	TestKindNamespaceListerExpansion
}

// testKindNamespaceLister implements the TestKindNamespaceLister
// interface.
type testKindNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TestKinds in the indexer for a given namespace.
func (s testKindNamespaceLister) List(selector labels.Selector) (ret []*v1.TestKind, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TestKind))
	})
	return ret, err
}

// Get retrieves the TestKind from the indexer for a given namespace and name.
func (s testKindNamespaceLister) Get(name string) (*v1.TestKind, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("testkind"), name)
	}
	return obj.(*v1.TestKind), nil
}
