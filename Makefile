.PHONY: all test build_deamon build_client build

OUTPUT_DEAMON=container-deamon
OUTPUT_CLIENT=container-client
GO?=go

all: build

install_lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1

lint:
	golangci-lint --version
	golangci-lint run --timeout 5m


clean:
	go clean

test:
	go test ./...

build: build_deamon build_client

build_deamon:
	go build -o $(OUTPUT_DEAMON) ./deamon.go

build_client:
	go build -o $(OUTPUT_CLIENT) ./client.go