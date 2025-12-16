# AgentBay CLI CI/CD Configuration

This directory contains the continuous integration and deployment configuration for AgentBay CLI on aone.ci platform.

## Overview

The CI/CD pipeline automatically triggers when code is merged to the `master` or `main` branch. It performs the following steps:

1. **Unit Tests** - Runs all unit tests and generates coverage reports
2. **Build and Package** - Builds binaries for multiple platforms and creates distribution packages

## Pipeline Structure

### Jobs

#### 1. Unit Tests (`unit-tests`)
- **Purpose**: Validate code quality and functionality
- **Resources**: 2 CPU cores, 8GB RAM
- **Timeout**: 20 minutes
- **Steps**:
  - Checkout code
  - Setup Go 1.23.0 environment
  - Install dependencies
  - Run unit tests
  - Generate coverage report

#### 2. Build and Package (`build-and-package`)
- **Purpose**: Build multi-platform binaries and create distribution packages
- **Resources**: 4 CPU cores, 16GB RAM
- **Timeout**: 45 minutes
- **Dependencies**: Requires `unit-tests` to pass
- **Steps**:
  - Setup build variables (version, timestamp, git commit)
  - Build binaries for all platforms
  - Create distribution packages
  - Print build summary

### Supported Platforms

The pipeline builds binaries for the following platforms:

- **macOS**
  - Intel (x64)
  - Apple Silicon (ARM64)
- **Linux**
  - x64
  - ARM64
- **Windows**
  - x64
  - ARM64

### Package Formats

- **Linux/macOS**: `.tar.gz` archives
- **Windows**: `.zip` archives and standalone `.exe` files
- All packages include `.sha256` checksum files

## Configuration Variables

The pipeline uses the following variables that must be configured in aone.ci:

### Global Variables (vars)
- `BINARY_NAME`: The name of the binary (must be set to: `agentbay`)
- `VERSION_PREFIX`: Version prefix for builds (default: `dev`)
- `OSS_BUCKET`: OSS bucket name for storing build artifacts
- `OSS_ENDPOINT`: OSS endpoint URL (e.g., `oss-cn-hangzhou.aliyuncs.com`)

### Secrets (secrets)
- `OSS_ACCESS_KEY_ID`: OSS access key ID for authentication
- `OSS_ACCESS_KEY_SECRET`: OSS access key secret for authentication

### Runtime Variables
- `VERSION`: Auto-generated as `{VERSION_PREFIX}-{TIMESTAMP}`
- `TIMESTAMP`: Build timestamp in format `YYYYMMDD-HHMM` (Asia/Shanghai timezone)
- `GIT_COMMIT`: Short git commit hash
- `BUILD_DATE`: Build date in ISO 8601 format

## Trigger Conditions

The pipeline automatically triggers on:
- Push to `master` branch
- Push to `main` branch

## Build Artifacts

After successful build, the following artifacts are created in the `packages/` directory:

### Binary Archives
```
agentbay-{VERSION}-darwin-amd64.tar.gz
agentbay-{VERSION}-darwin-arm64.tar.gz
agentbay-{VERSION}-linux-amd64.tar.gz
agentbay-{VERSION}-linux-arm64.tar.gz
agentbay-{VERSION}-windows-amd64.zip
agentbay-{VERSION}-windows-arm64.zip
```

### Standalone Executables (Windows)
```
agentbay-{VERSION}-windows-amd64.exe
agentbay-{VERSION}-windows-arm64.exe
```

### Checksum Files
Each package includes a corresponding `.sha256` file for integrity verification.

## OSS Upload

After successful build, all packages are automatically uploaded to Alibaba Cloud OSS.

### Upload Structure
```
oss://{OSS_BUCKET}/agentbay/releases/{VERSION}/
├── agentbay-{VERSION}-darwin-amd64.tar.gz
├── agentbay-{VERSION}-darwin-amd64.tar.gz.sha256
├── agentbay-{VERSION}-darwin-arm64.tar.gz
├── agentbay-{VERSION}-darwin-arm64.tar.gz.sha256
├── agentbay-{VERSION}-linux-amd64.tar.gz
├── agentbay-{VERSION}-linux-amd64.tar.gz.sha256
├── agentbay-{VERSION}-linux-arm64.tar.gz
├── agentbay-{VERSION}-linux-arm64.tar.gz.sha256
├── agentbay-{VERSION}-windows-amd64.zip
├── agentbay-{VERSION}-windows-amd64.zip.sha256
├── agentbay-{VERSION}-windows-amd64.exe
├── agentbay-{VERSION}-windows-amd64.exe.sha256
├── agentbay-{VERSION}-windows-arm64.zip
├── agentbay-{VERSION}-windows-arm64.zip.sha256
├── agentbay-{VERSION}-windows-arm64.exe
└── agentbay-{VERSION}-windows-arm64.exe.sha256
```

### Public Access
All uploaded files are set with `public-read` ACL, allowing public downloads via:
```
https://{OSS_BUCKET}.{OSS_ENDPOINT}/agentbay/releases/{VERSION}/{filename}
```

## Usage

### Modifying the Pipeline

To modify the pipeline configuration, edit the `cicd.yml` file in this directory.

### Adding New Build Steps

To add new steps to the pipeline, add them under the appropriate job's `steps` section:

```yaml
- id: your-step-id
  name: Your Step Name
  run: |
    echo "Your commands here"
```

### Changing Trigger Conditions

To change when the pipeline triggers, modify the `triggers` section:

```yaml
triggers:
  push:
    branches:
      - master
      - main
      - develop  # Add more branches
```

## Testing Locally

Before pushing changes, you can test the build process locally:

```bash
# Run unit tests
make test-unit

# Build for all platforms
make build-all

# Check generated binaries
ls -lh bin/
```

## Troubleshooting

### Build Failures

If the build fails:

1. Check the logs in the aone.ci console
2. Verify that all unit tests pass locally: `make test-unit`
3. Ensure dependencies are up to date: `make deps`
4. Test multi-platform build locally: `make build-all`

### Missing Packages

If packages are not created:

1. Verify the `bin/` directory contains all expected binaries
2. Check that the `create-packages` step completed successfully
3. Ensure `zip` and `tar` utilities are available in the CI environment

### Version Issues

If version numbers are incorrect:

1. Check the `VERSION_PREFIX` variable in `cicd.yml`
2. Verify that git commit information is available
3. Review the `setup-build-vars` step output

## References

- [aone.ci Documentation](https://aone.ci/docs)
- [AgentBay CLI Documentation](../README.md)
- [Makefile Reference](../Makefile)

