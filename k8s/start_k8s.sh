#!/bin/bash

# Step #1. Ensure logged in
# az login

# Step #2. Create K8s
aks-engine deploy \
    --dns-prefix dfreilich-interview \
    --resource-group dfreilich-interview \
    --location uksouth \
    --api-model cluster.json \
    --force-overwrite

export KUBECONFIG="_output/dfreilich-interview/kubeconfig/kubeconfig.uksouth.json"
kubectl cluster-info
kubectl get pods -A

# Use HAProxy Ingress Controller
# Install based on their latest docs: https://www.haproxy.com/documentation/kubernetes/latest/usage/ingress/
helm repo add haproxytech https://haproxytech.github.io/helm-charts
helm repo update
helm upgrade --install kubernetes-ingress haproxytech/kubernetes-ingress \
    --create-namespace \
    --namespace haproxy-controller \
    --set controller.service.type=LoadBalancer

# Apply Services and Networking
kubectl apply -f service-a.yaml
kubectl apply -f service-b.yaml
kubectl apply -f networking.yaml
