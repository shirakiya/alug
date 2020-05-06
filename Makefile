RUN_CONTEXT ?= docker-compose run --rm alug

docker-build:
	docker-compose build alug

docker-login:
	docker-compose run --rm alug /bin/bash

run:
	$(RUN_CONTEXT) go run cmd/alug/main.go

build:
	CGO_ENABLED=0 go build -o ./bin/alug ./cmd/alug/main.go

test:
	$(RUN_CONTEXT) go test -v ./...

coverage:
	$(RUN_CONTEXT) go test -race -cover -covermode=atomic ./...

fmt:
	$(RUN_CONTEXT) go fmt ./...
