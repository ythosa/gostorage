SHELL := /bin/bash

.PHONY: build
build:
	go build -v ./cmd/storageserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./internal/...

.PHONY: run
run:
	go run ./cmd/storageserver

.PHONY: swag
swag:
	swag init -g cmd/gostorage/main.go

.PHONY: pipeline
pipeline:
	make swag && make run

.PHONY: lint
lint:
	golangci-lint run --color always

.DEFAULT_GOAL := pipeline
