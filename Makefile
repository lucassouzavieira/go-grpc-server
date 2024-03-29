mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
base_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

# Golang project related settings
GO?=go
APP?=$(base_dir)
APP_NAME=my-app
APP_DIR=app
CLI_NAME=$(APP_NAME)-cli
CLI_DIR=cli
PROJECT_PACKAGE=github.com/lucassouzavieira/$(APP)
WORK_DIR=/go/src/$(PROJECT_PACKAGE)
BUILD_DIRECTORY=build
LINTER_EXECUTABLE := golangci-lint
LINTER_PATH := $(GOPATH)/bin/$(LINTER_EXECUTABLE)
GOFMT := gofmt
BUILD_ENV :=
BUILD_ENV += CGO_ENABLED=0

# Docker 
DOCKER_REPOSITORY=localhost
DOCKER_IMAGE=$(APP)
DOCKER_TAG=1.0
DOCKER_NAMESPACE=grpc-learning
DOCKER_IMAGE_NAME=$(DOCKER_NAMESPACE)/$(DOCKER_IMAGE)
DOCKER_BUILD_ENV :=
DOCKER_BUILD_ENV += --log-level=debug --platform=linux/amd64 --rm

# Make commands
.PHONY: build
build:
	$(BUILD_ENV) go build -o build/$(APP_NAME) -a ./cmd/$(APP_DIR)

build-cli:
	$(BUILD_ENV) go build -o build/$(CLI_NAME) -a ./cmd/$(CLI_DIR)

docker-build:
	docker build $(DOCKER_BUILD_ENV) -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

proto:
	protoc --go_out=:. --go-grpc_out=:. internal/protobuf/schema/*.proto

.PHONY: lint
lint:
	$(LINTER_PATH) run ./...

.PHONY: fmt
fmt:
	$(GOFMT) -s -w .

.PHONY: test
test:
	git restore -- test/
	$(GO) test -v -cover -timeout 3m ./...
	git restore -- test/

# Utilitary commands
install-tools:
	go get github.com/google/wire/cmd/wire
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2