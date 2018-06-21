package v1beta1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/pmorie/federation-category-experiment/pkg/apis/federation/v1beta1"
	. "github.com/pmorie/federation-category-experiment/pkg/client/clientset/versioned/typed/federation/v1beta1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the FederatedReplicaSetPlacement resource tests

var _ = Describe("FederatedReplicaSetPlacement", func() {
	var instance FederatedReplicaSetPlacement
	var expected FederatedReplicaSetPlacement
	var client FederatedReplicaSetPlacementInterface

	BeforeEach(func() {
		instance = FederatedReplicaSetPlacement{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	// INSERT YOUR CODE HERE - add more "Describe" tests

	// Automatically created storage tests
	Describe("when sending a storage request", func() {
		Context("for a valid config", func() {
			It("should provide CRUD access to the object", func() {
				client = cs.FederationV1beta1().FederatedReplicaSetPlacements("default")

				By("returning success from the create request")
				actual, err := client.Create(&instance)
				Expect(err).ShouldNot(HaveOccurred())

				By("defaulting the expected fields")
				Expect(actual.Spec).To(Equal(expected.Spec))

				By("returning the item for list requests")
				result, err := client.List(metav1.ListOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Items).To(HaveLen(1))
				Expect(result.Items[0].Spec).To(Equal(expected.Spec))

				By("returning the item for get requests")
				actual, err = client.Get(instance.Name, metav1.GetOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual.Spec).To(Equal(expected.Spec))

				By("deleting the item for delete requests")
				err = client.Delete(instance.Name, &metav1.DeleteOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				result, err = client.List(metav1.ListOptions{})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Items).To(HaveLen(0))
			})
		})
	})
})
