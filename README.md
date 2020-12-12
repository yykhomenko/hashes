docker build -t hashes .

docker run -p 8080:8081 -it hashes

kubectl create deployment hashes --image cbiot/hashes:latest

kubectl expose deployment hashes --type=LoadBalancer --port=8080
