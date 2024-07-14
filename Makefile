build:
	go build -o ./bin/download-service
	chmod +x ./bin/download-service

run:
	./bin/download-service

docker-build:
	docker build -t artifactory-download:latest .

k8s-apply:
	kubectl apply -k ./k8s/_base
	kubectl rollout restart -n artifactory deployment/artifactory-download

deploy: docker-build k8s-apply