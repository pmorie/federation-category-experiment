package inject

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	federationv1beta1 "github.com/pmorie/federation-category-experiment/pkg/apis/federation/v1beta1"
	rscheme "github.com/pmorie/federation-category-experiment/pkg/client/clientset/versioned/scheme"
	"github.com/pmorie/federation-category-experiment/pkg/controller/federatedreplicaset"
	"github.com/pmorie/federation-category-experiment/pkg/controller/federatedreplicasetoverride"
	"github.com/pmorie/federation-category-experiment/pkg/controller/federatedreplicasetplacement"
	"github.com/pmorie/federation-category-experiment/pkg/inject/args"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&federationv1beta1.FederatedReplicaSet{}, arguments.Informers.Federation().V1beta1().FederatedReplicaSets()); err != nil {
			return err
		}
		if err := arguments.ControllerManager.AddInformerProvider(&federationv1beta1.FederatedReplicaSetOverride{}, arguments.Informers.Federation().V1beta1().FederatedReplicaSetOverrides()); err != nil {
			return err
		}
		if err := arguments.ControllerManager.AddInformerProvider(&federationv1beta1.FederatedReplicaSetPlacement{}, arguments.Informers.Federation().V1beta1().FederatedReplicaSetPlacements()); err != nil {
			return err
		}

		// Add Kubernetes informers

		if c, err := federatedreplicaset.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		if c, err := federatedreplicasetoverride.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		if c, err := federatedreplicasetplacement.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &federationv1beta1.FederatedReplicaSetCRD)
	Injector.CRDs = append(Injector.CRDs, &federationv1beta1.FederatedReplicaSetOverrideCRD)
	Injector.CRDs = append(Injector.CRDs, &federationv1beta1.FederatedReplicaSetPlacementCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"federation.pmorie.toy"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "federation.pmorie.toy",
		Version: "v1beta1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
