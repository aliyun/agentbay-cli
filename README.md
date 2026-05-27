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

### Update

**macOS / Linux (Homebrew) — fast path (recommended for routine updates):**

```bash
git -C "$(brew --repository aliyun/agentbay)" pull --ff-only && brew upgrade agentbay
```

Refreshes only the `aliyun/agentbay` tap and then upgrades agentbay. Skips Homebrew's full metadata sync (large `formula.jws.json` / `cask.jws.json` downloads and brew self-update), so it usually finishes in seconds.

**macOS / Linux (Homebrew) — fallback if `brew` itself reports errors:**

```bash
brew update && brew upgrade agentbay
```

Refreshes Homebrew itself, all taps, and the core formula metadata before upgrading agentbay. Slower but more thorough — use this if the fast path fails (e.g., after a long time without a brew refresh, or after a Homebrew breaking change).

**Windows (PowerShell):** re-run the install command to upgrade in place.

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

### Uninstall

```bash
# macOS / Linux (Homebrew)
brew uninstall agentbay
brew untap aliyun/agentbay   # optional
```

```powershell
# Windows (PowerShell)
# Note: if you installed with a custom -InstallPath or $env:AGENTBAY_PATH,
# replace "$env:LOCALAPPDATA\agentbay" below with your actual install directory.
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
# Restart PowerShell for the PATH change to take effect.
```

> **Note (Homebrew):** The first `brew install agentbay` builds from source and will automatically install Go as a build dependency, so it may take a few minutes. Subsequent upgrades reuse the cache.

See [Installation Guide](docs/en/installation.md) for details (including pre-built binaries and troubleshooting).

---

## Authentication

**AccessKey (recommended for scripts/CI):**

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

See [Authentication & Environment](docs/en/authentication.md) for STS, OAuth (not recommended), and environment variables.

---

## Documentation

| Document                                                       | Description                                                           |
| -------------------------------------------------------------- | --------------------------------------------------------------------- |
| [Installation Guide](docs/en/installation.md)                  | Detailed installation steps and troubleshooting                       |
| [Authentication & Environment](docs/en/authentication.md)      | AccessKey, STS, OAuth, and environment variables                      |
| [Image Creation & Sharing Workflow](docs/en/image-workflow.md) | End-to-end tutorial from Dockerfile template to cross-account sharing |
| [Image Management](docs/en/image.md)                           | Image lifecycle management command reference                          |
| [Docker Operations](docs/en/docker.md)                         | ACR login, image push, and sharing                                    |
| [API Key Management](docs/en/apikey.md)                        | Key creation, enable, disable, delete                                 |
| [RAM Permissions](docs/en/ram-permissions.md)                  | Required RAM permissions by command group                             |
| [FAQ](docs/en/faq.md)                                          | Frequently asked questions                                            |

For full command details, see the [Command Reference](docs/en/README.md).

---

## RAM Permissions (RAM Sub-accounts Only)

The main Alibaba Cloud account does **not** require any additional permission configuration. If you are using a **RAM sub-account** with AK/SK authentication, grant the required permissions via the [RAM console](https://ram.console.aliyun.com/users).

For the complete list of required permissions and Policy JSON examples for each command group, see [RAM Permissions](docs/en/ram-permissions.md).

---

## Command Overview

| Group   | Commands                                                                                                                           | Description      | Details                 |
| ------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------- | ----------------------- |
| Core    | `version`, `login`, `logout`                                                                                                       | Version & auth   | [→](docs/en/core.md)    |
| Image   | `list`, `init`, `create`, `create-from-template`, `activate`, `deactivate`, `delete`, `status`, `set-max-session`, `warmup-status` | Image lifecycle  | [→](docs/en/image.md)   |
| API Key | `create`, `enable`, `disable`, `delete`, `list`, `concurrency set`, `describe-key-content`                                         | Key management   | [→](docs/en/apikey.md)  |
| Network | `package list`                                                                                                                     | Network config   | [→](docs/en/network.md) |
| Skills  | `push`, `update`, `show`, `list`, `delete`                                                                                         | Skill management | [→](docs/en/skills.md)  |
| Docker  | `login`, `tag`, `push`, `share`, `unshare`, `list-shares`                                                                          | Docker registry  | [→](docs/en/docker.md)  |

---

## Quick Start

```bash
# 1. Authenticate (AccessKey recommended)
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. Create an API key (account real-name verification is required)
agentbay apikey create "my-api-key"

# 3. List your API keys and find the API Key (akm-xxxxxxxxxxxxxxxx) from the output
agentbay apikey list

# 4. Disable the API key when temporarily not needed
agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx

# 5. Re-enable it later
agentbay apikey enable --api-key akm-xxxxxxxxxxxxxxxx

# 6. Delete the API key permanently (must be DISABLED first; --yes skips prompts)
agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx --yes
```

> **Tip:** For automation scripts, you can use `--api-key-id ak-xxxxxxxxxxxxxxxx` (returned by `apikey create`) instead of `--api-key`. See [API Key docs](docs/en/apikey.md#terminology) for details.

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
This project is licensed under the Apache License 2.0 — see the [LICENSE](LICENSE) file for details.
