.PHONY: build, lint, test

build:
	go build -v ./cmd/api/

lint:
	golint ./... && golangci-lint run

test:	
	go test -v ./...

.DEFAULT_GOAL := build