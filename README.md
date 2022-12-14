[![Build](https://github.com/dfreilich/azure-k8s-cluster/actions/workflows/build.yml/badge.svg)](https://github.com/dfreilich/azure-k8s-cluster/actions/workflows/build.yml)

# azure-k8s-cluster
Repo used to generate Azure K8s Cluster

The prompt provided is to create a Azure K8S cluster with services.

## Repo Overview:

There are three elements to the repository: 
- [k8s](https://github.com/dfreilich/azure-k8s-cluster/tree/main/k8s) &rarr; Contains all the files relevant to the creation of the Kubernetes environment, done using the [aks-engine](https://github.com/Azure/aks-engine/) CLI
  * `clean_up.sh` &rarr; Deletes the current AZ K8s cluster
  * `start_k8s.sh` &rarr; Creates the cluster, using `aks-engine`, installs the `haproxy-controller` for `Ingress` and applies the relevant `yaml` files
  * `cluster.json` &rarr; The `aks-engine` configuration
  * `service-a.yaml` &rarr; Contains the definition of the `Deployment`, `Service`, and `Ingress` for service-a
  * `service-b.yaml` &rarr; Contains the definition of the `Deployment`, `Service`, and `Ingress` for service-b
  * `networking.yaml` &rarr; Contains the definition of the `NetworkPolicy` used to deny all unfiltered access between services (to ensure `service-a` can't talk to `service-b`), as well as another policy to allow `service-b` to talk to `service-a` (since the prompt seemed to imply that it was only `service-a` which couldn't access `service-b`, but not vice versa)
- [service-a](https://github.com/dfreilich/azure-k8s-cluster/tree/main/service-a) &rarr; Defines a simple tested service in Go which retrieves the bitcoin value in dollar from an API every minute and logs it, as well as logging the average value of the last 10 minutes. In addition, it exposes three endpoints: 
  * `/` &rarr; Returns `Hello World`, as well as the current Bitcoin price in `Euro`. 
  * `/healthz` &rarr; A simple health-check
  * `/readyz` &rarr; A simple liveness check
- [service-b](https://github.com/dfreilich/azure-k8s-cluster/tree/main/service-b) &rarr; Defines a simple, tested service in Go which  exposes three endpoints: 
  * `/` &rarr; Returns `Hello World` 
  * `/healthz` &rarr; A simple health-check
  * `/readyz` &rarr; A simple liveness check


## Cluster details

- [x] Use Azure and AKS-Engine (not pre-defined AKS solutions)
- [x] Set up K8S cluster , with RBAC enabled
- [x] Cluster should have 2 services – A and B
- [x] Cluster should have Ingress controller, redirecting traffic by URL: xxx/service-A or xxx/service-B
- [x] Service-A should not be able to “talk” with Service-B (policy disabled)
- [x] For Service A, write a script\application which retrieves the bitcoin value in dollar from an API on the web (you should find one) every minute and prints it, And also every 10 minutes it should print the average value of the last 10 minutes.


## General Guideline
- [x] Please, consider this as process for setting up “production-ready” cluster by all meaning, the following cluster buildout should be automated and fully repeatable, Pods should utilize liveness and readiness
- [x] Code should be supportable

