# AgentBay CLI

[中文版](README.zh-CN.md) | **English**

A command-line interface for AgentBay services.

---

## Overview

AgentBay CLI is a Cobra-based command-line tool that talks to AgentBay services through Alibaba Cloud OpenAPI. It provides image management, API key management, network management, skills management, Docker operations, and flexible authentication.

> The current CLI version supports creating and activating **CodeSpace** type images only.

---

## Installation

```bash
# macOS / Linux (Homebrew)
brew tap aliyun/agentbay && brew install agentbay

# Windows (PowerShell)
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"

# Verify
agentbay version
```

See [Installation Guide](docs/en/installation.md) for details.

---

## Authentication

**AccessKey (recommended for scripts/CI):**

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

See [Authentication & Environment](docs/en/authentication.md) for STS, OAuth (not recommended), and environment variables.

---

## Command Overview

| Group | Commands | Description | Details |
|-------|----------|-------------|---------|
| Core | `version`, `login`, `logout` | Version & auth | [→](docs/en/core.md) |
| Image | `list`, `init`, `create`, `create-from-template`, `activate`, `deactivate`, `delete`, `status`, `set-max-session`, `warmup-status` | Image lifecycle | [→](docs/en/image.md) |
| API Key | `create`, `enable`, `disable`, `delete`, `list`, `concurrency set` | Key management | [→](docs/en/apikey.md) |
| Network | `package list` | Network config | [→](docs/en/network.md) |
| Skills | `push`, `show`, `list` | Skill management | [→](docs/en/skills.md) |
| Docker | `login`, `tag`, `push` | Docker registry | [→](docs/en/docker.md) |

---

## Quick Start

```bash
# 1. Authenticate (AccessKey recommended)
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. Create an API key (account real-name verification is required)
agentbay apikey create "my-api-key"

# 3. List your API keys and find the API Key ID (akm-xxxxxxxxxxxxxxxx) from the output
agentbay apikey list

# 4. Disable the API key when temporarily not needed
agentbay apikey disable akm-xxxxxxxxxxxxxxxx

# 5. Re-enable it later
agentbay apikey enable akm-xxxxxxxxxxxxxxxx

# 6. Delete the API key permanently (must be DISABLED first; --yes skips prompts)
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes
```

For full command details, see the [Command Reference](docs/en/README.md).

---

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

---

## Notes

- When both AccessKey env vars and OAuth tokens are present, the CLI prefers AccessKey for API calls.
- System images are always available and don't need activation; only User images must be activated.
- API keys require real-name verification before creation, and each key must have a unique name.
- Use `--yes` / `-y` on destructive commands (`apikey delete`, `image delete`) to skip prompts in non-interactive environments.

---

## License

This project is licensed under the Apache License 2.0 — see the [LICENSE](LICENSE) file for details.
