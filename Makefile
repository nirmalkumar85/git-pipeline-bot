# Go project configuration
GO := go
APP_NAME := pullrequest-cli
BUILD_DIR := ./bin
SRC_DIR := ./src
TEST_DIR := ./tests

# Default Go binary output
BINARY_NAME := $(APP_NAME)

# Directories for source, build, and test files
SRC := $(shell find . -type f -name "*.go" -not -path "./vendor/*")
VENDOR := ./vendor
GO111MODULE := on

# Define the go version
GO_VERSION := 1.20

# Build the application
build: fmt lint test
	@echo "Building the Go binary..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Format the code using gofmt
fmt:
	@echo "Running gofmt..."
	$(GO) fmt $(SRC_DIR)

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Lint the code using golangci-lint
lint:
	@echo "Running linting..."
	golangci-lint run --config .golangci.yml

# Run the application (assuming you've built it)
run:
	@echo "Running the application..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(BUILD_DIR)

# Install dependencies (if you are using Go Modules)
install:
	@echo "Installing dependencies..."
	$(GO) mod tidy

# Install golangci-lint (if not already installed)
install-linter:
	@echo "Installing golangci-lint..."
	curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v1.52.0/golangci-lint-1.52.0-linux-amd64.tar.gz | tar -xvzf - -C /tmp/
	mv /tmp/golangci-lint-1.52.0-linux-amd64/golangci-lint /usr/local/bin

# Run unit tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# Help command to list available make targets
help:
	@echo "Makefile commands:"
	@echo "  build           Build the Go binary."
	@echo "  fmt             Format the Go code."
	@echo "  lint            Lint the code using golangci-lint."
	@echo "  test            Run tests."
	@echo "  run             Run the application."
	@echo "  clean           Clean up build artifacts."
	@echo "  install         Install dependencies."
	@echo "  install-linter  Install golangci-lint."
	@echo "  test-coverage   Run tests with coverage report."
