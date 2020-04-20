docker-build:
	docker-compose build alug

docker-login:
	docker-compose run --rm alug /bin/bash

run:
	go run cmd/main.go

build:
	CGO_ENABLED=0 go build -o ./bin/alug ./cmd/alug/main.go

test:
	go test -v ./...

coverage:
	go test -race -cover -covermode=atomic ./...

fmt:
	go fmt ./...
