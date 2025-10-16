# Makefile untuk Go Bitcoin Wallet

.PHONY: build run clean test lint help install deps

# Variables
BINARY_NAME=go-wallet
MAIN_PATH=cmd/wallet/main.go
BUILD_DIR=bin

# Default target
all: deps build

# Help command
help:
	@echo "Go Bitcoin Wallet - Makefile Commands"
	@echo ""
	@echo "Usage:"
	@echo "  make deps       - Download dependencies"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run the application"
	@echo "  make test       - Run tests"
	@echo "  make lint       - Run linter"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make install    - Install the binary"
	@echo "  make all        - Download deps and build"

# Download dependencies
deps:
	@echo "📦 Downloading dependencies..."
	go mod download
	go mod verify

# Build the application
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for multiple platforms
build-all:
	@echo "🔨 Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "✅ Multi-platform build complete"

# Run the application
run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Run with arguments
run-args: build
	./$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -v -cover ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

# Run linter
lint:
	@echo "🔍 Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "✨ Formatting code..."
	go fmt ./...
	goimports -w .

# Install the binary to GOPATH
install: build
	@echo "📥 Installing $(BINARY_NAME)..."
	go install $(MAIN_PATH)
	@echo "✅ Installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	go clean
	@echo "✅ Clean complete"

# Update dependencies
update-deps:
	@echo "🔄 Updating dependencies..."
	go get -u ./...
	go mod tidy

# Generate documentation
docs:
	@echo "📚 Generating documentation..."
	godoc -http=:6060 &
	@echo "✅ Documentation server running at http://localhost:6060"

# Create example wallet
example:
	@echo "💡 Creating example wallet..."
	./$(BUILD_DIR)/$(BINARY_NAME) create ExampleWallet

# Security audit
security:
	@echo "🔒 Running security audit..."
	go list -json -m all | nancy sleuth
