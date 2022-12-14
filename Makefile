install:
	brew update && brew install azure-cli && brew install Azure/aks-engine/aks-engine && brew install helm

test:
	go test ./... -v

containerize:
	pack build dfreilich/azure-k8s-service --builder paketobuildpacks/builder:tiny --default-process service-a
	docker push dfreilich/azure-k8s-service
