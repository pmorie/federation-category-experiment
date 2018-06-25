# federation-category-experiment

This repository is a place to explore use of categories to provide a good
experience for [federation-v2](https://github.com/kubernetes-sigs/federation-v2)
resources. This repository was created using
[kubebuilder](https://github.com/kubernetes-sigs/kubebuilder).

Categories are sets of arbitrary resources. Resources declare themselves to be
in categories; Kubernetes CRDs can declare themselves to have categories using
the `spec.names.categories` field:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: federatedreplicasets.federation.pmorie.toy
spec:
  group: federation.pmorie.toy
  names:
    categories:
    - federation
    kind: FederatedReplicaSet
    plural: federatedreplicasets
  scope: Namespaced
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          type: string
        kind:
          type: string
        metadata:
          type: object
        spec:
          type: object
        status:
          type: object
      type: object
  version: v1beta1
```

The `kubectl` tool can use categories for `get`:

```shell
$ k get federation --field-selector=metadata.name=my-rs --show-kind                                      
NAME                                                      AGE
federatedreplicasetoverride.federation.pmorie.toy/my-rs   6m
                                                    
NAME                                                       AGE                                           
federatedreplicasetplacement.federation.pmorie.toy/my-rs   6m
                                                    
NAME                                              AGE
federatedreplicaset.federation.pmorie.toy/my-rs   6m
```

## Running this example

Install the CRDS as follows:

```shell
$ kubectl create -f crd.yaml

$ kubectl get crd
NAME                                                  CREATED AT
federatedreplicasetoverrides.federation.pmorie.toy    2018-06-22T15:07:26Z
federatedreplicasetplacements.federation.pmorie.toy   2018-06-22T15:07:26Z
federatedreplicasets.federation.pmorie.toy            2018-06-22T15:07:26Z
```

You can then create example resources using:

```shell
$ kubectl create -f example-resources.yaml
```

Now you can try out using categories with `kubectl`. Try:

```shell
$ kubectl get federation --field-selector=metadata.name=my-rs --show-kind
NAME                                              AGE
federatedreplicaset.federation.pmorie.toy/my-rs   2d

NAME                                                      AGE
federatedreplicasetoverride.federation.pmorie.toy/my-rs   2d

NAME                                                       AGE
federatedreplicasetplacement.federation.pmorie.toy/my-rs   2d
```

and:

```shell
$ kubectl get federation --selector=app=my-federated-app --show-kind
NAME                                              AGE
federatedreplicaset.federation.pmorie.toy/my-rs   2d

NAME                                                      AGE
federatedreplicasetoverride.federation.pmorie.toy/my-rs   2d

NAME                                                       AGE
federatedreplicasetplacement.federation.pmorie.toy/my-rs   2d
```