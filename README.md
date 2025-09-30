# AgentBay CLI

Command line interface for AgentBay services.

## Features

- **Authentication**: OAuth-based login/logout
- **Image Management**: 
  - Create custom images from Dockerfiles
  - List available images with filtering and pagination
  - Activate User images for deployment
  - Deactivate activated User images *(requires API enhancement)*
- **Configuration Management**: Secure token storage and management

## Installation

```bash
# Build from source
make build

# Run tests
make test
```

## Usage

### Authentication
```bash
# Login to AgentBay
agentbay login

# Logout from AgentBay
agentbay logout
```

### Image Management
```bash
# List all user images
agentbay image list

# List images with filters
agentbay image list --os-type Linux --size 5

# Create a custom image
agentbay image create my-image --dockerfile ./Dockerfile --imageId base_image_id

# Activate an image
agentbay image activate imgc-xxxxxxxxxxxxxx

# Deactivate an image (API enhancement required)
agentbay image deactivate imgc-xxxxxxxxxxxxxx
```

## API Status

### Image Deactivate Command

The `image deactivate` command has been implemented but requires backend API enhancement:

**Current Status**: ✅ Command implemented, ❌ API not ready

**Issue**: The `DeleteResourceGroup` API currently requires a `resourceGroupId` parameter, but the command only has access to `imageId`.

**Required API Enhancement**: 
- Option 1: Modify `DeleteResourceGroup` API to accept `imageId` parameter and internally resolve to `resourceGroupId`
- Option 2: Add a new API endpoint specifically for image deactivation that takes `imageId`

**Error Message**: 
```
MissingParameter: There is a missing parameter resourceGroupId
```

**Implementation**: The command structure, validation, status polling, and user experience are fully implemented and ready for testing once the API is enhanced.

## Development

```bash
# Run unit tests
make test

# Build the CLI
make build

# Run with verbose output
./agentbay --verbose <command>
``` 