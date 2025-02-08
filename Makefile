.PHONY: all test build_deamon build_client build protoc protoc_install

OUTPUT_DEAMON=container-deamon
OUTPUT_CLIENT=container-client
GO?=go

all: build

install_lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1

lint: install_lint
	golangci-lint --version
	golangci-lint run --timeout 5m


clean:
	go clean

protoc_install:
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


protoc: protoc_install
	@protoc --go_out=. --go-grpc_out=. ./google/rpc/service.proto

test:
	$(GO) test ./...

build: build_deamon build_client

build_deamon: protoc
	$(GO) build -o $(OUTPUT_DEAMON) ./deamon.go

build_client: protoc
	$(GO) build -o $(OUTPUT_CLIENT) ./client.go