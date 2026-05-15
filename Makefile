.PHONY: build run dev test clean lint vet fmt

BIN_NAME=bin/server
MAIN_PATH=./cmd/server

build:
	go build -o $(BIN_NAME) $(MAIN_PATH)

run: build
	./$(BIN_NAME)

dev:
	go run $(MAIN_PATH)

test:
	go test -v ./...

clean:
	rm -f $(BIN_NAME)

lint:
	gofmt -l .

vet:
	go vet ./...

fmt:
	go fmt ./...
