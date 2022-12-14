install:
	brew update && brew install azure-cli && brew install Azure/aks-engine/aks-engine

test:
	go test ./... -v

containerize:
	pack build dfreilich/azure-k8s/service-b --builder paketobuildpacks/builder:tiny --env BP_GO_TARGETS="./service-b"
