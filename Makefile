# Copyright 2025 AgentBay CLI Contributors
# SPDX-License-Identifier: Apache-2.0

# Build variables
BINARY_NAME=agentbay
VERSION?=dev
GIT_COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X github.com/agentbay/agentbay-cli/cmd.Version=$(VERSION) -X github.com/agentbay/agentbay-cli/cmd.GitCommit=$(GIT_COMMIT) -X github.com/agentbay/agentbay-cli/cmd.BuildDate=$(BUILD_DATE)"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build targets
.PHONY: all build clean test test-unit test-integration test-all coverage test-coverage deps help

all: test-unit build

build: deps
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf bin/
	rm -f coverage.out coverage.html

test: test-unit

test-unit: deps
	$(GOTEST) -v ./test/unit/...

test-integration: deps
	$(GOTEST) -v -tags=integration ./test/integration/...

test-all: deps
	$(GOTEST) -v ./test/unit/... ./test/integration/...

coverage: deps
	$(GOTEST) -v -coverprofile=coverage.out ./test/unit/... ./internal/... ./cmd/...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

test-coverage: coverage

deps:
	$(GOMOD) download

# Tidy go modules (for local development only, not used in CI)
.PHONY: tidy
tidy:
	$(GOMOD) tidy

# Development targets
.PHONY: dev-build dev-test dev-run

dev-build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) .

dev-test:
	$(GOTEST) -v ./...

dev-run: dev-build
	./$(BINARY_NAME)

# Installation targets
.PHONY: install uninstall

install: build
	cp $(BINARY_NAME) /usr/local/bin/

uninstall:
	rm -f /usr/local/bin/$(BINARY_NAME)

# Cross-compilation targets
.PHONY: build-linux build-windows build-darwin build-all

build-linux: deps
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 .

build-windows: deps
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-arm64.exe .

build-darwin: deps
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 .

build-all: deps
	@mkdir -p bin
	$(MAKE) build-linux
	$(MAKE) build-windows
	$(MAKE) build-darwin

# Help target
help:
	@echo "Available targets:"
	@echo "  all          - Run unit tests and build binary"
	@echo "  build        - Build binary"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run unit tests (default)"
	@echo "  test-unit    - Run unit tests"
	@echo "  test-integration - Run integration tests"
	@echo "  test-all     - Run all tests"
	@echo "  coverage     - Run tests with coverage"
	@echo "  test-coverage - Run tests with coverage (alias)"
	@echo "  deps         - Download dependencies"
	@echo "  tidy         - Tidy go modules (local development only)"
	@echo "  dev-build    - Quick development build"
	@echo "  dev-test     - Quick development test"
	@echo "  dev-run      - Build and run binary"
	@echo "  install      - Install binary to /usr/local/bin"
	@echo "  uninstall    - Remove binary from /usr/local/bin"
	@echo "  build-linux  - Cross-compile for Linux"
	@echo "  build-windows- Cross-compile for Windows"
	@echo "  build-darwin - Cross-compile for macOS"
	@echo "  build-all    - Cross-compile for all platforms"
	@echo "  help         - Show this help message" 