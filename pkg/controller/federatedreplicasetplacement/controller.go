package federatedreplicasetplacement

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	federationv1beta1 "github.com/pmorie/federation-category-experiment/pkg/apis/federation/v1beta1"
	federationv1beta1client "github.com/pmorie/federation-category-experiment/pkg/client/clientset/versioned/typed/federation/v1beta1"
	federationv1beta1informer "github.com/pmorie/federation-category-experiment/pkg/client/informers/externalversions/federation/v1beta1"
	federationv1beta1lister "github.com/pmorie/federation-category-experiment/pkg/client/listers/federation/v1beta1"

	"github.com/pmorie/federation-category-experiment/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for FederatedReplicaSetPlacement resources goes here.

func (bc *FederatedReplicaSetPlacementController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on federatedreplicasetplacement.FederatedReplicaSetPlacementController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=federation,version=v1beta1,kind=FederatedReplicaSetPlacement,resource=federatedreplicasetplacements
type FederatedReplicaSetPlacementController struct {
	// INSERT ADDITIONAL FIELDS HERE
	federatedreplicasetplacementLister federationv1beta1lister.FederatedReplicaSetPlacementLister
	federatedreplicasetplacementclient federationv1beta1client.FederationV1beta1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	federatedreplicasetplacementrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &FederatedReplicaSetPlacementController{
		federatedreplicasetplacementLister: arguments.ControllerManager.GetInformerProvider(&federationv1beta1.FederatedReplicaSetPlacement{}).(federationv1beta1informer.FederatedReplicaSetPlacementInformer).Lister(),

		federatedreplicasetplacementclient:   arguments.Clientset.FederationV1beta1(),
		federatedreplicasetplacementrecorder: arguments.CreateRecorder("FederatedReplicaSetPlacementController"),
	}

	// Create a new controller that will call FederatedReplicaSetPlacementController.Reconcile on changes to FederatedReplicaSetPlacements
	gc := &controller.GenericController{
		Name:             "FederatedReplicaSetPlacementController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&federationv1beta1.FederatedReplicaSetPlacement{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a FederatedReplicaSetPlacement Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the FederatedReplicaSetPlacementController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
