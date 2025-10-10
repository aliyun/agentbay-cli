# AgentBay CLI

A command-line interface for AgentBay services.

## Features

AgentBay CLI provides comprehensive image management capabilities:

- **Authentication**: Secure OAuth-based login with Google account integration
- **Image Creation**: Build custom images from Dockerfiles with base image support
- **Image Management**: Activate, deactivate, and monitor image instances
- **Image Listing**: Browse user and system images with pagination and filtering support
- **Configuration Management**: Secure token storage and automatic token refresh

## Quick Start

```bash
# 1. Log in to AgentBay
agentbay login

# 2. List available user images
agentbay image list

# 3. Create a custom image
agentbay image create myapp --dockerfile ./Dockerfile --imageId code_latest

# 4. Activate the image
agentbay image activate imgc-xxxxx...xxx

# 5. Deactivate when done
agentbay image deactivate imgc-xxxxx...xxx
```

For detailed usage instructions and examples, see the [User Guide](docs/USER_GUIDE.md).


## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details. 