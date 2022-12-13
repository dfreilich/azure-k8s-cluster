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
