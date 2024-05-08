ORG    := yykhomenko
NAME   := hashes
REPO   := ${ORG}/${NAME}
TAG    := ${REPO}:$(shell date '+%y%m%d')-$(shell git log -1 --pretty=format:'%h')
LATEST := ${REPO}:latest

include .env
export

update:## Update dependencies
	go get -u ./...
	go mod tidy

lint: ## Run linters
	golangci-lint run --no-config --issues-exit-code=0 --timeout=10m \
    --disable-all --enable=deadcode  --enable=gocyclo --enable=revive --enable=varcheck \
    --enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
    --enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck

test:	## Run tests
	go test -race -timeout 30s ./...

bench: ## Run benchmarks
	go test ./... -bench=. -benchmem

build: ## Build application
	go build ./cmd/${NAME}

run: ## Build and start version
	go run ./cmd/${NAME}

start: ## Start version
	./${NAME}

clean: ## Clean project
	rm -f ${NAME}

install: ## Install version
	make build
	make test
	go install ./cmd/${NAME}

image: ## Build image
	docker build -t ${TAG} -t ${LATEST} .

publish: ## Publish image
	docker push ${TAG}
	docker push ${LATEST}

deploy: ## Deploy to k8s cluster
	kubectl apply -f deployments/hashes-namespace.yml
	kubectl apply -f deployments

undeploy: ## Undeploy from k8s cluster
	kubectl delete -f deployments

show: ## Show service
	minikube service hashes-service

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / \
  {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
