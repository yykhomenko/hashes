build: ## Build a version
	go build -v ./cmd/...

test:	## Run all the tests
	go test -v -race -timeout 30s ./...

image: ## Build an image
	docker build -t cbiot/hashes .

publish: ## Publish an image
	docker push cbiot/hashes

deploy: ## Deploy a container
	kubectl apply -f hashes-deployment.yml
	kubectl apply -f hashes-service.yml

undeploy: ## Undeploy a container
	kubectl delete -f hashes-deployment.yml
	kubectl delete -f hashes-service.yml

show: ## Show a service
	minikube service hashes-service

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
