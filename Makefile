.PHONY: build clean test

# Binary name
BINARY_NAME=dcgo

# Build the application
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Clean build files
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test ./...

# Build and run
run: build
	./$(BINARY_NAME)

# Install dependencies
deps:
	go mod download

# Default target
all: deps test build 