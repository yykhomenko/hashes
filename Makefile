ORG    := yykhomenko
NAME   := hashes
REPO   := ${ORG}/${NAME}
TAG    := $(shell git log -1 --pretty=format:"%h")
IMG    := ${REPO}:${TAG}
LATEST := ${REPO}:latest

build: ## Build version
	go build ./cmd/${NAME}

lint: ## Run linters
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
    --disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
    --enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
    --enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck

test:	## Run tests
	go test -race -timeout 30s ./...

bench: ## Run benchmarks
	go test ./... -bench=. -benchmem

run: ## Build and start version
	go run ./cmd/${NAME}

start: ## Start version
	./cmd/${NAME}

clean: ## Clean project
	rm -f ${NAME}

install: ## Install version
	make build
	make test
	go install ./cmd/${NAME}

image: ## Build image
	docker build -t ${IMG} -t ${LATEST} .

publish: ## Publish image
	docker push ${IMG}
	docker push ${LATEST}

deploy: ## Deploy to k8s cluster
	kubectl apply -f deployments/hashes-namespace.yml
	kubectl apply -f deployments

undeploy: ## Undeploy from k8s cluster
	kubectl delete -f deployments

show: ## Show service
	minikube service hashes-service

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
  awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
