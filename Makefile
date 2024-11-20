SHELL := /bin/bash

VERSION ?= "v0.0.0-dev"
REVISION ?= ""
CREATED ?= ""

.PHONY: all
.SILENT: all
all: tidy lint fmt vet test build

.PHONY: lint
.SILENT: lint
lint:
	golangci-lint run

.PHONY: fmt
.SILENT: fmt
fmt:
	go fmt ./...

.PHONY: tidy
.SILENT: tidy
tidy:
	go mod tidy

.PHONY: vet
.SILENT: vet
vet:
	go vet ./...

.PHONY: test 
.SILENT: test
test:
	go test -timeout 1m ./... -cover

.PHONY: cover
.SILENT: cover
cover:
	mkdir -p .tmp
	go test -timeout 1m ./... -coverprofile=.tmp/coverage.out
	go tool cover -html=.tmp/coverage.out	

.PHONY: build
.SILENT: build
build:
	CGO_ENABLED=0 go build -ldflags "-w -s" -o bin/pimctl ./main.go
