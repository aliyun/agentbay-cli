# AgentBay CLI

A command-line interface for AgentBay services.

## Features

AgentBay CLI provides image management, API key management, network management, and skills management:

**Note**: The current version of the CLI tool supports creating and activating CodeSpace type images only.

- **Authentication**: OAuth login with Aliyun, or AccessKey via environment variables (`AGENTBAY_ACCESS_KEY_ID` / `AGENTBAY_ACCESS_KEY_SECRET`) for automation and CI
- **Dockerfile Template**: Download Dockerfile templates from the cloud
- **Image Creation**: Build custom images from Dockerfiles with base image support; automatically parses and uploads COPY/ADD referenced files
- **Image Management**: Activate, deactivate, and monitor image instances with configurable resource specifications (CPU/memory) and network types
- **Image Listing**: Browse user and system images with separated display, pagination and filtering support
- **Image Status**: Query resource lifecycle status for an image by ID (`agentbay image status`)
- **API Key Management**: Create API keys and configure session concurrency limits for authentication and access control
- **Network Management**: Query network packages by region, view package details including EIP addresses and office site bindings
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

# 5. Activate the image (uses default resources; specify --cpu/--memory for other sizes)
agentbay image activate imgc-xxxxx...xxx

# Activate with specific CPU and memory
agentbay image activate imgc-xxxxx...xxx --cpu 2 --memory 4

# Activate with advanced network configuration
agentbay image activate imgc-xxxxx...xxx --network-type ADVANCED --session-bandwidth 100 --dns-address 8.8.8.8 --dns-address 8.8.4.4

# Activate with sandbox lifecycle parameters
agentbay image activate imgc-xxxxx...xxx --lifecycle-mode auto --lifecycle-max-runtime 3600 --lifecycle-hibernate 1800 --lifecycle-idle-timeout 600

# 6. Deactivate when done
agentbay image deactivate imgc-xxxxx...xxx

# Optional: check resource status (activate/deactivate lifecycle, not Docker build task)
agentbay image status imgc-xxxxx...xxx

# API Key Management (optional)
agentbay apikey create --name "my-api-key"                        # Create a new API key
agentbay apikey concurrency set --api-key-id ak-xxx --concurrency 10  # Set concurrency limit

# Network Management (optional)
agentbay network package list                              # List network packages (default region: cn-hangzhou)
agentbay network package list --biz-region-id cn-shanghai  # List for a specific region

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
- Image activation uses default resource configuration if `--cpu` and `--memory` are not specified. CPU and memory must be specified together.
- Advanced network type (`--network-type ADVANCED`) requires `--session-bandwidth` and `--dns-address` parameters.
- Sandbox lifecycle parameters (`--lifecycle-mode`, `--lifecycle-max-runtime`, `--lifecycle-hibernate`, `--lifecycle-idle-timeout`) are optional and override existing policy values. `--lifecycle-mode` accepts `auto` or `manual`.
- API keys require account real-name verification before creation. Each API key must have a unique name.
- Network package list uses `cn-hangzhou` as the default region. Use `--biz-region-id` to query other regions.

For detailed usage instructions and examples, see the [User Guide](docs/USER_GUIDE.md) .

## Development Workflow

This project is synchronized across three repositories:

| Remote                          | URL                                                    | Role                                                               |
| ------------------------------- | ------------------------------------------------------ | ------------------------------------------------------------------ |
| `aliyun`                        | https://github.com/aliyun/agentbay-cli                 | Public-facing repository, single source of truth                   |
| `origin`                        | https://code.alibaba-inc.com/InnoArchClub/agentbay-cli | Internal mirror, triggers aone.ci build for pre-merge verification |
| `upstream` _(legacy, optional)_ | https://github.com/LXY02/agentbay-cli                  | Personal fork, historical PR proxy (no longer required)            |

### Branches

- **`aliyun/master`** is the baseline for all new development.
- **`origin/master`** is frozen for historical audit purposes. **No new feature should be based on it.**
- Feature branches use the `feat/**`, `fix/**`, or `release/**` naming convention.

### Standard Flow (new features / bug fixes)

```bash
# 0. One-time setup — add the aliyun remote if missing
git remote add aliyun git@github.com:aliyun/agentbay-cli.git
git fetch aliyun

# 1. Create a feature branch from aliyun/master
git fetch aliyun
git checkout -b feat/<topic> aliyun/master

# 2. Develop and commit
git commit -am "feat: <description>"

# 3. Push to internal GitLab → triggers .aoneci/cicd.yml for build verification
git push origin feat/<topic>

#    (Optional) adjust .aoneci/cicd.yml on this branch if the feature needs
#    special build behavior, then push again to rerun the pipeline.

# 4. After internal verification passes, push to aliyun and open a PR
git push aliyun feat/<topic>
#    Open PR on https://github.com/aliyun/agentbay-cli : feat/<topic> → master
```

### Why this workflow

- `.aoneci/cicd.yml` now triggers on `feat/**`, `fix/**`, and `release/**` branches,
  so every feature branch gets an internal multi-platform build + OSS upload automatically.
- Because feature branches are cut from `aliyun/master`, no cherry-pick is required
  when promoting a change to the public repository.
- `.github/workflows/` on the public repo only runs on tag pushes (`v*`) or manual
  dispatch, so pushing feature branches to `aliyun` is safe and does not trigger releases.

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.
