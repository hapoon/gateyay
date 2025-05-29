BIN := gateyay
VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS := "-s -w -X main.revision=$(CURRENT_REVISION)"
GOBIN ?= $(shell go env GOPATH)/bin
export GO111MODULE=on

.PHONY: run
run:
	@go run main.go

.PHONY: test
test:
	@go test -v -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: build
build:
	@go build -o ${BIN} main.go

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: format
format:
	@go fmt ./...

.PHONY: show-version
show-version: $(GOBIN)/gobump
	@gobump show -r .

$(GOBIN)/gobump:
	go install github.com/x-motemen/gobump/cmd/gobump@latest