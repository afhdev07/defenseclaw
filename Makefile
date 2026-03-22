BINARY    := defenseclaw
VERSION   := 0.1.0
GOFLAGS   := -ldflags "-X main.version=$(VERSION)"

.PHONY: build build-all build-linux-arm64 build-linux-amd64 build-darwin-arm64 build-darwin-amd64 test lint clean

build:
	go build $(GOFLAGS) -o $(BINARY) ./cmd/defenseclaw

build-all: build-linux-arm64 build-linux-amd64 build-darwin-arm64 build-darwin-amd64

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -o $(BINARY)-linux-arm64 ./cmd/defenseclaw

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $(BINARY)-linux-amd64 ./cmd/defenseclaw

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) -o $(BINARY)-darwin-arm64 ./cmd/defenseclaw

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -o $(BINARY)-darwin-amd64 ./cmd/defenseclaw

test:
	go test -race ./...

lint:
	golangci-lint run

clean:
	rm -f $(BINARY) $(BINARY)-*
