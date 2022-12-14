#!/bin/bash

# Step #1. Ensure logged in
# az login

# Step #2. Create K8s
aks-engine deploy \
    --dns-prefix dfreilich-interview \
    --resource-group dfreilich-interview \
    --location uksouth \
    --api-model cluster.json

export KUBECONFIG="_output/dfreilich-interview/kubeconfig/kubeconfig.uksouth.json"
kubectl cluster-info
kubectl get pods -A

# Use HAProxy Ingress Controller
# Install based on their latest docs: https://www.haproxy.com/documentation/kubernetes/latest/usage/ingress/
helm repo add haproxytech https://haproxytech.github.io/helm-charts
helm repo update
helm install kubernetes-ingress haproxytech/kubernetes-ingress \
    --create-namespace \
    --namespace haproxy-controller \
    --set controller.service.type=LoadBalancer
