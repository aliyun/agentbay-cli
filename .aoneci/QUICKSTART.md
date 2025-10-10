# Quick Start Guide - AgentBay CLI CI/CD

## Overview

This document provides a quick start guide for using the aone.ci CI/CD pipeline for AgentBay CLI.

## Prerequisites

- aone.ci account with access to the repository
- Familiarity with Git and aone.ci platform
- (Optional) OSS bucket and credentials if deploying packages

## Getting Started

### 1. Pipeline Configuration

The pipeline is already configured in `.aoneci/cicd.yml`. It will automatically trigger when:
- Code is pushed to the `master` branch
- Code is pushed to the `main` branch

### 2. Local Testing Before Push

Always test locally before pushing to ensure the CI/CD pipeline will succeed:

```bash
# Run unit tests
make test-unit

# Build for all platforms
make build-all

# Verify binaries were created
ls -lh bin/
```

Expected output:
```
bin/
├── agentbay-darwin-amd64
├── agentbay-darwin-arm64
├── agentbay-linux-amd64
├── agentbay-linux-arm64
├── agentbay-windows-amd64.exe
└── agentbay-windows-arm64.exe
```

### 3. Push Code to Trigger CI/CD

```bash
# Stage your changes
git add .

# Commit your changes
git commit -m "your commit message"

# Push to master (or main)
git push origin master
```

The CI/CD pipeline will automatically start.

## Monitoring the Pipeline

1. Go to the aone.ci console: https://aone.ci
2. Navigate to your repository
3. Click on "Pipelines" or "CI/CD"
4. Find the latest pipeline run
5. Click on it to view detailed logs

## Pipeline Stages

### Stage 1: Unit Tests (约 2-3 分钟)

The pipeline first runs all unit tests to ensure code quality.

**What it does:**
- Checks out code
- Sets up Go 1.23.0 environment
- Downloads dependencies
- Runs unit tests with `make test-unit`
- Generates coverage report

**Success criteria:**
- All unit tests pass
- No compilation errors

### Stage 2: Build and Package (约 5-8 分钟)

After tests pass, the pipeline builds binaries for all platforms.

**What it does:**
- Sets up build variables (version, timestamp, git commit)
- Builds binaries for 6 platforms
- Creates distribution packages (tar.gz, zip)
- Generates SHA256 checksums
- Prints build summary

**Success criteria:**
- All 6 binaries built successfully
- All packages created with checksums
- No build errors

## Build Artifacts

After successful build, the following artifacts are available:

### Binary Archives
- `agentbay-{VERSION}-darwin-amd64.tar.gz`
- `agentbay-{VERSION}-darwin-arm64.tar.gz`
- `agentbay-{VERSION}-linux-amd64.tar.gz`
- `agentbay-{VERSION}-linux-arm64.tar.gz`
- `agentbay-{VERSION}-windows-amd64.zip`
- `agentbay-{VERSION}-windows-arm64.zip`

### Direct Executables (Windows)
- `agentbay-{VERSION}-windows-amd64.exe`
- `agentbay-{VERSION}-windows-arm64.exe`

### Checksums
- All packages include `.sha256` checksum files

## Troubleshooting

### Pipeline Fails on Unit Tests

**Symptom:** Red ❌ on the "Unit Tests" stage

**Solution:**
1. Check the test output in the pipeline logs
2. Run `make test-unit` locally to reproduce the issue
3. Fix the failing tests
4. Push the fix

### Pipeline Fails on Build

**Symptom:** Red ❌ on the "Build and Package" stage

**Solution:**
1. Check the build output in the pipeline logs
2. Run `make build-all` locally to reproduce the issue
3. Common issues:
   - Missing dependencies: Run `make deps`
   - Compilation errors: Fix the code
   - Build tool issues: Ensure Go 1.23.0 is installed
4. Push the fix

### Build is Slow

**Expected times:**
- Unit Tests: 2-3 minutes
- Build and Package: 5-8 minutes
- Total: 7-11 minutes

If significantly slower:
- Check aone.ci system status
- Check if resources are properly allocated (2-8Gi for tests, 4-16Gi for builds)
- Review pipeline logs for any hanging processes

### Artifacts Not Found

**Solution:**
1. Ensure the pipeline completed successfully
2. Check the "Build and Package" stage logs
3. Look for "✅ All packages created" message
4. If using OSS upload (optional), verify OSS credentials

## Environment Variables

The pipeline uses the following variables:

### Global Variables (configured in aone.ci)
- `BINARY_NAME`: `agentbay`
- `VERSION_PREFIX`: `dev` (or your preferred prefix)

### Auto-generated Variables
- `VERSION`: `{VERSION_PREFIX}-{TIMESTAMP}`
- `TIMESTAMP`: Build timestamp (YYYYMMDD-HHMM)
- `GIT_COMMIT`: Short git commit hash
- `BUILD_DATE`: ISO 8601 format date

## Best Practices

### 1. Test Before Push
Always run `make test-unit` and `make build-all` locally before pushing.

### 2. Meaningful Commit Messages
Use clear commit messages that describe what changed:
```bash
git commit -m "feat: add new command for image management"
git commit -m "fix: correct error handling in login flow"
git commit -m "docs: update README with installation instructions"
```

### 3. Monitor Pipeline
Check the pipeline status after pushing to catch issues early.

### 4. Keep Dependencies Updated
Regularly update Go dependencies:
```bash
make deps
git add go.mod go.sum
git commit -m "chore: update dependencies"
```

### 5. Review Logs
If something fails, carefully review the pipeline logs to understand the root cause.

## Advanced Configuration

### Changing Trigger Branches

To trigger on different branches, edit `.aoneci/cicd.yml`:

```yaml
triggers:
  push:
    branches:
      - master
      - main
      - develop  # Add more branches here
```

### Adjusting Resource Allocation

If builds are slow or failing due to resource constraints, adjust in `.aoneci/cicd.yml`:

```yaml
jobs:
  unit-tests:
    runs-on: 4-16Gi  # Increase from 2-8Gi if needed

  build-and-package:
    runs-on: 8-32Gi  # Increase from 4-16Gi if needed
```

### Adding OSS Upload

To automatically upload artifacts to OSS (like agbcloud-cli does):

1. Configure OSS credentials in aone.ci secrets:
   - `OSS_ACCESS_KEY_ID`
   - `OSS_ACCESS_KEY_SECRET`
   - `OSS_BUCKET`
   - `OSS_ENDPOINT`

2. Add upload step to `.aoneci/cicd.yml` (refer to agbcloud-cli project)

## Getting Help

### Documentation
- CI/CD Pipeline Details: See [README.md](README.md)
- aone.ci Documentation: https://aone.ci/docs
- AgentBay CLI Documentation: See [../README.md](../README.md)

### Support
- Check pipeline logs first
- Review this quickstart and README.md
- Contact DevOps team for aone.ci platform issues
- Open an issue in the repository for code-related problems

## Summary

The CI/CD pipeline is designed to be simple and reliable:

1. **Push code** → Automatic trigger
2. **Tests run** → Ensure quality
3. **Build multi-platform binaries** → Ready for distribution
4. **View results** → Monitor in aone.ci console

That's it! The pipeline handles the rest automatically.

