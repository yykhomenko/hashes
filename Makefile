.PHONY: build
build: ## Build a version
	go build -v ./cmd/hashes

.PHONY: test
test: ## Run all the tests
	go test -v -race -timeout 30s ./...

.PHONY: image
image: ## Build an image
	docker build -t cbiot/hashes .

.PHONY: publish
publish: ## Publish an image
	docker push cbiot/hashes

.PHONY: pull
pull: ## Pull an image
	docker pull cbiot/hashes:latest

.PHONY: run
run: ## Run a container
	docker run --rm --name=cbiot_hashes -p 8080:8080 -it cbiot/hashes:latest

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
