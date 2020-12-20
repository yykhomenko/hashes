ORG    := cbiot
NAME   := hashes
REPO   := ${ORG}/${NAME}
TAG    := $(shell git log -1 --pretty=format:"%h")
IMG    := ${REPO}:${TAG}
LATEST := ${REPO}:latest

build: ## Build a version
	go build -v ./cmd/...

test:	## Run all the tests
	go test -v -race -timeout 30s ./...

image: ## Build an image
	docker build -t ${IMG} -t ${LATEST} .

publish: ## Publish an image
	docker push ${IMG}
	docker push ${LATEST}

deploy: ## Deploy to k8s cluster
	kubectl apply -f deployments

undeploy: ## Undeploy from k8s cluster
	kubectl delete -f deployments

show: ## Show a service
	minikube service hashes-service

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
