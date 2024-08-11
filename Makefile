run:
	go run ./cmd

test:
	go test ./...

install:
	go get ./...

build:
	go fmt ./... && go vet ./... && go test ./... && go build -o bot ./cmd