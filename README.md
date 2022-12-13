[![Build](https://github.com/dfreilich/azure-k8s-cluster/actions/workflows/build.yml/badge.svg)](https://github.com/dfreilich/azure-k8s-cluster/actions/workflows/build.yml)

# azure-k8s-cluster
Repo used to generate Azure K8s Cluster

The prompt provided is to create a Azure K8S cluster with services.

## Cluster details
1.      Use Azure and AKS-Engine (not pre-defined AKS solutions)
1.      Set up K8S cluster , with RBAC enabled
1.      Cluster should have 2 services – A and B
1.      Cluster should have Ingress controller, redirecting traffic by URL: xxx/service-A or xxx/service-B
1.      Service-A should not be able to “talk” with Service-B (policy disabled)
1.      For Service A, write a script\application which retrieves the bitcoin value in dollar from an API on the web (you should find one) every minute and prints it, And also every 10 minutes it should print the average value of the last 10 minutes.

## General Guideline
1.      Please, consider this as process for setting up “production-ready” cluster by all meaning, the following cluster buildout should be automated and fully repeatable, Pods should utilize liveness and readiness
1.      Code should be supportable
