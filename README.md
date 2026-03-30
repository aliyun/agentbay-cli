# AgentBay CLI

A command-line interface for AgentBay services.

## Features

AgentBay CLI provides image management and skills management:

**Note**: The current version of the CLI tool supports creating and activating CodeSpace type images only.

- **Authentication**: OAuth login with Aliyun, or AccessKey via environment variables (`AGENTBAY_ACCESS_KEY_ID` / `AGENTBAY_ACCESS_KEY_SECRET`) for automation and CI
- **Dockerfile Template**: Download Dockerfile templates from the cloud
- **Image Creation**: Build custom images from Dockerfiles with base image support; automatically parses and uploads COPY/ADD referenced files
- **Image Management**: Activate, deactivate, and monitor image instances
- **Image Listing**: Browse user and system images with separated display, pagination and filtering support
- **Image Status**: Query resource lifecycle status for an image by ID (`agentbay image status`)
- **Skills**: Push local skills and show skill details by ID (`skills list` is a placeholder until the backend list API is available)
- **Configuration Management**: Secure token storage and automatic token refresh

## Quick Start

```bash
# 1. Authenticate (pick one)
agentbay login
# Or set AGENTBAY_ACCESS_KEY_ID and AGENTBAY_ACCESS_KEY_SECRET (optional: AGENTBAY_ACCESS_KEY_SESSION_TOKEN for STS)

# 2. List available images
agentbay image list                    # List user images (default)
agentbay image list --include-system   # List both user and system images
agentbay image list --system-only      # List only system images

# 3. Download Dockerfile template
agentbay image init --sourceImageId code-space-debian-12    # Download Dockerfile template to current directory
# Or use short form:
agentbay image init -i code-space-debian-12

# 4. Create a custom image (using system image as base)
agentbay image create myapp --dockerfile ./Dockerfile --imageId code-space-debian-12

# 5. Activate the image (uses 2c4g by default; specify --cpu/--memory for other sizes)
agentbay image activate imgc-xxxxx...xxx

# 6. Deactivate when done
agentbay image deactivate imgc-xxxxx...xxx

# Optional: check resource status (activate/deactivate lifecycle, not Docker build task)
agentbay image status imgc-xxxxx...xxx

# Skills (optional; directory or .zip with SKILL.md frontmatter; list is a placeholder)
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
agentbay skills show <skill-id>              # Show skill details
```

**Note**: 
- With both OAuth tokens and AccessKey env set, the CLI prefers AccessKey for API calls.
- System images are always available and don't require activation. Only user-created images need to be activated before use.
- When downloading Dockerfile templates, the first N lines (N is returned by the system) are system-defined and cannot be modified. Only modify content after line N+1.
- Available sourceImageID for production environment: 
`code-space-debian-12`
`code-space-debian-12-enhanced`. 

For detailed usage instructions and examples, see the [User Guide](docs/USER_GUIDE.md) .


## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details. 