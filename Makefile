run:
	go run cmd/app/main.go
docker:
	sh ./commands/docker.sh
compose:docker
	docker-compose up -d
rm:
	docker-compose down && docker rmi ${DOCKER_USERNAME}/${DOCKER_IMAGE_NAME}
build:
	go build -o bin/app cmd/app/main.go
test:
	go test -v ./...
lint:
	golangci-lint run
init_db:
	sh ./commands/init_db.sh