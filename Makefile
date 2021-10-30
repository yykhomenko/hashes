ORG    := yykhomenko
NAME   := hashes
REPO   := ${ORG}/${NAME}
TAG    := $(shell git log -1 --pretty=format:"%h")
IMG    := ${REPO}:${TAG}
LATEST := ${REPO}:latest

build: ## Build a version
	GOOS=linux GOARCH=amd64 go build -v ./cmd/...

lint: ## Run linters
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
    --disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
    --enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
    --enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck

test:	## Run all the tests
	go test -v -race -timeout 30s ./...

image: ## Build an image
	docker build -t ${IMG} -t ${LATEST} .

publish: ## Publish an image
	docker push ${IMG}
	docker push ${LATEST}

deploy: ## Deploy to k8s cluster
	kubectl apply -f deployments/hashes-namespace.yml
	kubectl apply -f deployments

undeploy: ## Undeploy from k8s cluster
	kubectl delete -f deployments

show: ## Show a service
	minikube service hashes-service

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
