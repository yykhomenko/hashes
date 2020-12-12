build:						## Build a version
	go build -v ./cmd/hashes

test: 						## Run all the tests
	go test -v -race -timeout 30s ./...

image: 						## Build an image
	docker build -t cbiot/hashes .

publish: 					## Publish an image
	docker push cbiot/hashes

pull: 						## Pull an image
	docker pull cbiot/hashes:latest

run: 							## Run a container
	docker run --rm --name=cbiot_hashes -p 8080:8080 -it cbiot/hashes:latest

deploy: 					## Deploy a container
	kubectl create deployment hashes --image cbiot/hashes:latest
	kubectl expose deployment hashes --type=LoadBalancer --port=8080

undeploy: 				## Undeploy a container
	kubectl delete service hashes
	kubectl delete deployment hashes

query: 						## Query to container
	minikube service hashes

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
