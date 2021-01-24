# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif
 
all: manager
 
# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out
 
# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go
 
# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Run go fmt against code
fmt:
	go fmt ./...
 
# Run go vet against code
vet:
	go vet ./...
	