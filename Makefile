docker-build:
	docker-compose build alug

docker-login:
	docker-compose run --rm alug /bin/bash

run:
	go run cmd/main.go

test:
	go test -v ./...

fmt:
	go fmt ./...
