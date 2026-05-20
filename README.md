# AgentBay CLI

[中文版](README.zh-CN.md) | **English**

A command-line interface for AgentBay services.

---

## Overview

AgentBay CLI is a Cobra-based command-line tool that talks to AgentBay services through Alibaba Cloud OpenAPI. It provides:

- **Image Management** — create, list, activate / deactivate / delete custom images, query lifecycle status, check warm-up quota, and configure session concurrency
- **API Key Management** — create / list / enable / disable / delete keys, set per-key concurrency limits
- **Network Management** — query network packages and EIP bindings by region
- **Skills Management** — push local skills and inspect details by ID
- **Docker Operations** — log in to ACR, tag and push images for AgentBay
- **Authentication** — AccessKey / STS environment variables (recommended), or OAuth login for local development
- **Configuration** — secure token storage, automatic token refresh, multi-environment support

> The current CLI version supports creating and activating **CodeSpace** type images only.

---

## Installation

Pre-built binaries are available under `bin/` and `packages/`. On macOS / Linux you can also install via Homebrew tap (see `homebrew/agentbay.rb`).

```bash
# Verify installation
agentbay version
```

---

## Authentication

The CLI supports three authentication methods. **AccessKey or STS is the recommended method for production scripts and CI/CD.**

> Priority: `AGENTBAY_ACCESS_KEY_ID` / `AGENTBAY_ACCESS_KEY_SECRET` env vars > OAuth tokens stored locally.

### 1. AccessKey (Recommended)

Set the following environment variables. This is the preferred method for automation, scripts, and CI/CD:

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

### 2. STS Temporary Credentials (Recommended for short-lived sessions)

For STS (Security Token Service) temporary credentials, set the session token in addition to the AK/SK pair:

```bash
export AGENTBAY_ACCESS_KEY_ID="STS.xxx"
export AGENTBAY_ACCESS_KEY_SECRET="your-sts-secret"
export AGENTBAY_ACCESS_KEY_SESSION_TOKEN="your-sts-session-token"
```

### 3. OAuth Login (Deprecated — not recommended)

> WARNING: `agentbay login` is **deprecated and will be removed in a future release**. Please use AccessKey or STS instead.

```bash
agentbay login    # Opens a browser for OAuth login
agentbay logout   # Invalidate session and clear local credentials
```

---

## Environment Variables

All AgentBay CLI environment variables are optional unless noted otherwise.

### How to Set Environment Variables

**Method 1: Export in current session** (lost when terminal closes)

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

**Method 2: `.env` file in the working directory** (recommended for project-scoped configuration)

The CLI automatically loads `.env` from the current working directory on startup (via `godotenv`). Create a `.env` file:

```dotenv
AGENTBAY_ACCESS_KEY_ID=your-access-key-id
AGENTBAY_ACCESS_KEY_SECRET=your-access-key-secret
AGENTBAY_ENV=production
```

> Do **not** commit `.env` to version control — add it to `.gitignore`.

**Method 3: Shell profile** (persists across all terminal sessions)

Append to `~/.bashrc`, `~/.zshrc`, or equivalent:

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

Then reload: `source ~/.zshrc` (or open a new terminal).

**Method 4: Inline per command**

```bash
AGENTBAY_ENV=prerelease agentbay image list
```

### Authentication

| Variable                            | Description                                                  |
| ----------------------------------- | ------------------------------------------------------------ |
| `AGENTBAY_ACCESS_KEY_ID`            | AccessKey ID (or STS Token ID prefixed with `STS.`)          |
| `AGENTBAY_ACCESS_KEY_SECRET`        | AccessKey Secret (or STS secret)                             |
| `AGENTBAY_ACCESS_KEY_SESSION_TOKEN` | STS session token (only required when using STS credentials) |

### Environment Selection

| Variable       | Default      | Allowed values                                                                                                                             |
| -------------- | ------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `AGENTBAY_ENV` | `production` | `production` / `prod`, `prerelease` / `pre` / `staging`, `international` / `intl` / `prod-international`, `international-pre` / `intl-pre` |

```bash
# Switch to pre-release
export AGENTBAY_ENV=prerelease
# Switch to international production
export AGENTBAY_ENV=international
```

Endpoints per environment:

| Environment         | Endpoint                                   |
| ------------------- | ------------------------------------------ |
| `production`        | `xiaoying.cn-shanghai.aliyuncs.com`        |
| `prerelease`        | `xiaoying-pre.cn-hangzhou.aliyuncs.com`    |
| `international`     | `xiaoying.ap-southeast-1.aliyuncs.com`     |
| `international-pre` | `xiaoying-pre.ap-southeast-1.aliyuncs.com` |

### Advanced Configuration

| Variable                   | Description                                                                 |
| -------------------------- | --------------------------------------------------------------------------- |
| `AGENTBAY_CLI_ENDPOINT`    | Override the default API endpoint for the current environment               |
| `AGENTBAY_CLI_TIMEOUT_MS`  | API request timeout in milliseconds                                         |
| `AGENTBAY_CLI_CONFIG_DIR`  | Override the default config directory (default: `~/.agentbay`)              |
| `AGENTBAY_OAUTH_CLIENT_ID` | Override the default OAuth client ID (only relevant for `agentbay login`)   |
| `AGENTBAY_OAUTH_REGION`    | Override the OAuth region (`cn` or `intl`)                                  |
| `AGENTBAY_API_URL`         | _(Legacy)_ Same as `AGENTBAY_CLI_ENDPOINT`, kept for backward compatibility |

---

## Command Reference

Run `agentbay --help` to discover commands. Commands are organised into two groups: **Core Commands** and **Management Commands**.

### Core Commands

#### `agentbay version`

Show version, git commit, build date, current environment and endpoint.

```bash
agentbay version
```

#### `agentbay logout`

Log out from AgentBay (invalidate the OAuth session on the server and clear local credentials).

```bash
agentbay logout
```

#### `agentbay login` _(deprecated — will be removed in a future release)_

See [Authentication](#authentication).

---

### Image Management — `agentbay image`

#### `image list`

List available AgentBay images.

```bash
agentbay image list                      # User images (default)
agentbay image list --include-system     # User + system images
agentbay image list --system-only        # System images only
agentbay image list --os-type Linux      # Filter by OS type: Linux / Android / Windows
agentbay image list --page 2 --size 5    # Pagination
```

#### `image init`

Download a Dockerfile template from the cloud to the current directory.

```bash
agentbay image init --sourceImageId code-space-debian-12
agentbay image init -i code-space-debian-12
```

> The first N lines of the downloaded Dockerfile are system-defined and must not be modified. Only edit content after line N+1.

Available `sourceImageId` values for production:

- `code-space-debian-12`
- `code-space-debian-12-enhanced`

#### `image create` _(deprecated — use `create-from-template` instead)_

> WARNING: `image create` is **deprecated and will be removed in a future release**. To create a custom image, please use [`agentbay image create-from-template`](#image-create-from-template) instead.

Build a custom image from a Dockerfile. Files referenced by `COPY` / `ADD` are parsed and uploaded automatically.

```bash
agentbay image create myapp --dockerfile ./Dockerfile --imageId code-space-debian-12
agentbay image create myapp -f ./Dockerfile -i code-space-debian-12
```

#### `image create-from-template`

Create a custom image from a system image template (calls the `CreateImageFromTemplate` API).

```bash
agentbay image create-from-template \
  --source-image registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 \
  --name my-custom-image \
  --imageId <system-image-id>

# Short form
agentbay image create-from-template -s registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 -n my-custom-image -i <system-image-id>
```

#### `image activate`

Activate a User image so it can be used.

```bash
# Default resources
agentbay image activate imgc-xxxxxxxxxxxxxx

# Specific CPU and memory (must be specified together)
agentbay image activate imgc-xxxxxxxxxxxxxx --cpu 2 --memory 4

# Advanced network
agentbay image activate imgc-xxxxxxxxxxxxxx \
  --network-type ADVANCED \
  --session-bandwidth 100 \
  --dns-address 8.8.8.8 \
  --dns-address 8.8.4.4

# Sandbox lifecycle
agentbay image activate imgc-xxxxxxxxxxxxxx \
  --lifecycle-mode auto \
  --lifecycle-max-runtime 3600 \
  --lifecycle-hibernate 1800 \
  --lifecycle-idle-timeout 600

# Specify region
agentbay image activate imgc-xxxxxxxxxxxxxx --region-id cn-shanghai
```

Notes:

- `--cpu` and `--memory` must be specified together.
- `--network-type ADVANCED` requires `--session-bandwidth` and `--dns-address`.
- `--lifecycle-mode` accepts `auto` or `manual`.

#### `image deactivate`

Deactivate an activated User image.

```bash
agentbay image deactivate imgc-xxxxxxxxxxxxxx
```

#### `image delete`

Delete a User image **permanently**. Only deactivated User images can be deleted.

```bash
agentbay image delete imgc-xxxxxxxxxxxxxx          # With confirmation
agentbay image delete imgc-xxxxxxxxxxxxxx --yes    # Skip confirmation (CI / scripts)
```

#### `image status`

Query the resource lifecycle status of an image (different from the Docker build task status during `image create`).

```bash
agentbay image status imgc-xxxxxxxxxxxxxx
```

Common statuses: `IMAGE_CREATING`, `IMAGE_CREATE_FAILED`, `IMAGE_AVAILABLE`, `RESOURCE_DEPLOYING`, `RESOURCE_PUBLISHED`, `RESOURCE_DELETING`, `RESOURCE_FAILED`, `RESOURCE_CEASED`.

#### `image set-max-session`

Set the maximum concurrent session count for an activated User image. Requires the image to be in `RESOURCE_PUBLISHED` state and use **advanced network**.

```bash
agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 10
```

> The command polls until the resource group is ready (typically ~5 minutes).

#### `image warmup-status`

Query the warm-up status for the current account, including session quota, image quota, and details of warm-up images.

```bash
agentbay image warmup-status
```

Output includes:

- **Session Quota** — max session limit, total used, and available sessions
- **Image Quota** — max image count and current image count
- **Warm-up Images** — table of image IDs, total max size, and group count

---

### API Key Management — `agentbay apikey`

> API key creation requires account real-name verification. Each API key must have a unique name.

#### `apikey create`

Create a new API key.

```bash
agentbay apikey create --name "my-api-key"
```

#### `apikey enable`

Re-enable a disabled API key.

```bash
agentbay apikey enable akm-xxxxxxxxxxxxxxxx
```

#### `apikey disable`

Disable an API key (it can no longer authenticate requests).

```bash
agentbay apikey disable akm-xxxxxxxxxxxxxxxx
```

#### `apikey delete`

Delete an API key permanently. Only `DISABLED` keys can be deleted directly; if the key is `ENABLED` you will be prompted to disable it first.

```bash
agentbay apikey delete akm-xxxxxxxxxxxxxxxx          # Interactive (with confirmation)
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes    # Skip all prompts (CI / scripts)
```

#### `apikey list`

List API keys with optional filtering and pagination.

```bash
agentbay apikey list                                        # List up to 10 API keys
agentbay apikey list --max-results 20                       # List up to 20 API keys
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx         # Query a specific API key
agentbay apikey list --next-token <token>                   # Fetch the next page
```

#### `apikey concurrency set`

Set the maximum concurrent session limit for an API key.

```bash
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10
```

---

### Network Management — `agentbay network`

#### `network package list`

List network packages for a region.

```bash
agentbay network package list                              # Default region: cn-hangzhou
agentbay network package list --biz-region-id cn-shanghai  # Custom region
```

---

### Skills Management — `agentbay skills`

#### `skills push`

Push a local skill (directory or `.zip`) to the cloud. A directory must contain `SKILL.md` with `name` / `description` frontmatter; a directory is packed into a zip and uploaded.

```bash
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
```

#### `skills show`

Show skill details by ID.

```bash
agentbay skills show <skill-id>
```

#### `skills list` _(placeholder)_

Lists cloud skills. Backend list API is not yet available; this command currently acts as a placeholder.

---

### Docker Operations — `agentbay docker`

These commands wrap the local `docker` CLI to interact with the AgentBay ACR registry.

#### `docker login`

Log in to the AgentBay ACR registry using temporary credentials obtained from the `GetACRRepoCredential` API. The credential info (`RegistryUrl`, `Namespace`, `RepoName`, `ImageTag`) is cached for `tag` / `push`.

```bash
agentbay docker login
```

#### `docker tag`

Tag a local image for the AgentBay ACR registry. The target image name is constructed as `$RegistryUrl/$Namespace/$RepoName:<target-tag>`.

```bash
agentbay docker tag myapp:latest v1.0
```

> Run `agentbay docker login` first.

#### `docker push`

Push a tagged image to the AgentBay ACR registry. The image name must match `$RegistryUrl/$Namespace/$RepoName[:tag]`; mismatched names are rejected.

```bash
agentbay docker push <registry>/<namespace>/<repo>:v1.0
```

> Run `agentbay docker login` first.

---

## Quick Start

A minimal end-to-end flow:

```bash
# 1. Authenticate (AccessKey recommended)
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. Create an API key (account real-name verification is required)
agentbay apikey create --name "my-api-key"

# 3. Disable the API key when temporarily not needed
agentbay apikey disable akm-xxxxxxxxxxxxxxxx

# 4. Re-enable it later
agentbay apikey enable akm-xxxxxxxxxxxxxxxx

# 5. Delete the API key permanently (must be DISABLED first; --yes skips prompts)
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes
```

For full command details, see the [Command Reference](#command-reference) above and the [User Guide](docs/USER_GUIDE.md).

---

## Notes

- When both AccessKey env vars and OAuth tokens are present, the CLI prefers AccessKey for API calls.
- System images are always available and don't need activation; only User images must be activated.
- API keys require real-name verification before creation, and each key must have a unique name.
- Use `--yes` / `-y` on destructive commands (`apikey delete`, `image delete`) to skip prompts in non-interactive environments.

---

## License

This project is licensed under the Apache License 2.0 — see the [LICENSE](LICENSE) file for details.
