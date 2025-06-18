# Sprint Capacity Calculator Makefile

# Binary name
BINARY_NAME=sizely
BINARY_PATH=./bin/$(BINARY_NAME)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt

# Build info
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT = $(shell git rev-parse --short HEAD)
DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Linker flags
LDFLAGS = -ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# Build tags
BUILD_TAGS =

# Coverage
COVERAGE_FILE = coverage.out

.PHONY: all build clean test coverage lint fmt vet install uninstall help deps check release

# Default target
all: check build

## build: Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	$(GOBUILD) $(LDFLAGS) -tags '$(BUILD_TAGS)' -o $(BINARY_PATH) ./cmd/$(BINARY_NAME)

## install: Install the binary to $GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	$(GOBUILD) $(LDFLAGS) -tags '$(BUILD_TAGS)' -o $(GOPATH)/bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

## uninstall: Remove the binary from $GOPATH/bin
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@rm -f $(GOPATH)/bin/$(BINARY_NAME)

## clean: Remove build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf bin/
	@rm -f $(COVERAGE_FILE)

## test: Run all tests
test:
	@echo "Running tests..."
	$(GOTEST) -v -race ./...

## coverage: Generate test coverage report
coverage:
	@echo "Generating coverage report..."
	$(GOTEST) -race -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report generated: coverage.html"

## lint: Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install it with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.54.2"; \
	fi

## fmt: Format Go code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

## vet: Run go vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

## check: Run all checks (fmt, vet, lint, test)
check: fmt vet test
	@echo "All checks passed!"

## build-all: Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p bin
	# Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/$(BINARY_NAME)
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 ./cmd/$(BINARY_NAME)
	# macOS
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/$(BINARY_NAME)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 ./cmd/$(BINARY_NAME)
	# Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/$(BINARY_NAME)

## release: Create a release (requires goreleaser)
release:
	@echo "Creating release..."
	@if command -v goreleaser >/dev/null 2>&1; then \
		goreleaser release --rm-dist; \
	else \
		echo "goreleaser not found. Install it from: https://goreleaser.com/install/"; \
	fi

## snapshot: Create a snapshot release
snapshot:
	@echo "Creating snapshot release..."
	@if command -v goreleaser >/dev/null 2>&1; then \
		goreleaser release --snapshot --rm-dist; \
	else \
		echo "goreleaser not found. Install it from: https://goreleaser.com/install/"; \
	fi

## run-example: Run example with sample data
run-example: build
	@echo "Running example calculation..."
	$(BINARY_PATH) -calc -input examples/basic/tasks.json
	@echo "\nRunning reverse calculation example..."
	$(BINARY_PATH) -reverse -points 33

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):$(VERSION) .

## help: Show this help message
help:
	@echo "Available targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

# Check if required tools are installed
.PHONY: check-tools
check-tools:
	@echo "Checking required tools..."
	@command -v go >/dev/null 2>&1 || { echo "Go is required but not installed. Aborting." >&2; exit 1; }
	@command -v git >/dev/null 2>&1 || { echo "Git is required but not installed. Aborting." >&2; exit 1; }
	@echo "All required tools are installed."
