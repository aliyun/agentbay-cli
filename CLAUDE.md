# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

AgentBay CLI is a Go-based command-line tool for managing AgentBay services, particularly image management (CodeSpace type images). The project uses Cobra for CLI framework, OAuth-based authentication with Google, and Alibaba Cloud SDK for API interactions.

## Build and Development Commands

### Build Commands
- `make build` - Build the binary with debug symbols
- `make build-optimized` - Build optimized binary (smaller size, no debug symbols)
- `make dev-build` - Quick development build
- `make clean` - Clean build artifacts

### Test Commands
- `make test` or `make test-unit` - Run unit tests
- `make test-integration` - Run integration tests (tagged with `integration`)
- `make test-all` - Run all tests (unit + integration)
- `make coverage` - Run tests with coverage report (generates coverage.html)

### Development Workflow
- `make all` - Run unit tests and build (default workflow)
- `make dev-test` - Quick development test
- `make dev-run` - Build and run binary
- `make deps` - Download Go module dependencies

## Architecture

### Core Structure
- **main.go**: Entry point with Cobra root command setup
- **cmd/**: Command implementations (login, logout, image, version)
- **internal/agentbay/**: Main AgentBay API client with XML response handling
- **internal/auth/**: OAuth authentication and token management
- **internal/client/**: HTTP client wrapper with retry logic and API models
- **internal/config/**: Configuration management and environment handling
- **internal/models/**: Data models for Docker operations

### Key Components

#### CLI Commands
- `agentbay login` - OAuth-based authentication
- `agentbay image` - Image management (create, activate, deactivate, list)
- `agentbay logout` - Clear authentication
- `agentbay version` - Version information

#### Authentication Flow
The authentication system uses OAuth with Google integration and automatic token refresh. Tokens are securely stored and managed through the auth package.

#### API Client Architecture
- Uses Alibaba Cloud OpenAPI SDK for HTTP operations  
- Custom retry logic with exponential backoff
- XML response parsing with fallback mechanisms
- Context-aware request handling

#### Testing Strategy
- Unit tests in `test/unit/` directory
- Integration tests in `test/integration/` with build tags
- Tests cover command validation, auth flows, API interactions, and error handling

### Environment Configuration
The application supports multiple environments (development, staging, production) configured through the config package. Environment variables can be loaded from `.env` files.

## Important Notes

- Integration tests require the `integration` build tag
- The project uses structured logging with logrus
- Cross-compilation is supported for Linux, Windows, and macOS
- UPX compression targets available for production builds
- All API models are auto-generated in internal/client/