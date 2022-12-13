install:
	brew update && brew install azure-cli && brew install Azure/aks-engine/aks-engine

test:
	go test ./... -v
