.PHONY: up stop build package


up: stop build ## redeploy fenice
	docker-compose up -d --build

stop: ## stop fenice docker container
	docker-compose stop; docker-compose rm -f

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend

package:
	docker build -t accident/backend .

	