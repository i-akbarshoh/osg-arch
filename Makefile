run:
	go run cmd/main.go
build:
	go build -o bin/app cmd/app/main.go
test:
	go test -v ./...
lint:
	golangci-lint run
init_db:
	sh ./commands/init_postgres.sh