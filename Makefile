.PHONY: build
build: ## Build a version
	go build -v ./cmd/hashes

.PHONY: image
image: ## Build a image
	docker build -t cbiot/hashes .
.PHONY: publish-image

publish: ## Publish a image
	docker push cbiot/hashes

.PHONY: test
test: ## Run all the tests
	go test -v -race -timeout 30s ./...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
