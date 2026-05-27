# AgentBay CLI

[中文版](README.zh-CN.md) | **English**

A command-line interface for AgentBay services on Alibaba Cloud — image lifecycle, API keys, Docker, and skills management.

> The current CLI version supports creating and activating **CodeSpace** type images only.

---

## Features

- **Image lifecycle** — create from Dockerfile/template, activate, list, delete
- **Docker integration** — ACR login, push, cross-account share / unshare
- **API key management** — create, enable/disable, delete, concurrency control
- **Skills & Network** — push/update skills, list network packages
- **Multi-auth** — AccessKey (AK/SK), STS, OAuth
- **Cross-platform** — macOS, Linux, Windows

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

> First `brew install agentbay` builds from source and will install Go as a build dependency, so it may take a few minutes. Subsequent upgrades reuse the cache.

<details>
<summary><b>Update</b></summary>

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

</details>

<details>
<summary><b>Uninstall</b></summary>

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

</details>

See [Installation Guide](docs/en/installation.md) for pre-built binaries and troubleshooting.

---

## Quick Start — API Key in 60 seconds

```bash
# 1. Authenticate
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. Create an API key (account real-name verification is required)
agentbay apikey create "my-api-key"

# 3. Inspect / disable / re-enable / delete
agentbay apikey list
agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx
agentbay apikey enable  --api-key akm-xxxxxxxxxxxxxxxx
agentbay apikey delete  --api-key akm-xxxxxxxxxxxxxxxx --yes
```

> **Tip:** For automation scripts, you can use `--api-key-id ak-xxxxxxxxxxxxxxxx` (returned by `apikey create`) instead of `--api-key`. See [API Key docs](docs/en/apikey.md#terminology).
>
> Using a RAM sub-account? See [RAM Permissions](docs/en/ram-permissions.md) for the required policies.

---

## Tutorial — Image Workflow (end-to-end)

Build a custom image from a Dockerfile template, push it to ACR, then share it across Alibaba Cloud accounts.

**Scenario:** Account A builds a custom image and shares it with Account B; Account B creates its own image from the shared repository.

```bash
# ── Account A: build & publish ─────────────────────────────
agentbay image init --sourceImageId aio-ubuntu-2404            # 1. download Dockerfile template
agentbay docker login                                          # 2. ACR login (temp credentials, ~1h)
docker build -t <registry>/<namespace>/<uid>:<tag> -f Dockerfile .   # 3. build locally
docker push  <registry>/<namespace>/<uid>:<tag>                # 4. push to ACR
agentbay image create-from-template \                          # 5. create custom image
  --source-image /<namespace>/<uid>:<tag> \
  --name my-image --imageId aio-ubuntu-2404
agentbay docker share <ACCOUNT_B_UID>                          # 6. share repo to Account B
agentbay docker list-shares --direction Outgoing               # 7. verify the share

# ── Account B: receive & use ───────────────────────────────
agentbay docker list-shares --direction Incoming               # 8. view incoming shares
agentbay image create-from-template ...                        # 9. create own image from A's repo
```

→ **Full walkthrough** with concrete example values, expected output, and troubleshooting: **[Image Creation & Sharing Workflow](docs/en/image-workflow.md)**

> **Prerequisite:** Docker installed locally. On macOS we recommend [OrbStack](https://orbstack.dev/) — it is lightweight, fast, and uses far fewer resources than Docker Desktop.

---

## Commands

| Group   | Commands                                                                                                                           | Description      | Details                 |
| ------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------- | ----------------------- |
| Core    | `version`, `login`, `logout`                                                                                                       | Version & auth   | [→](docs/en/core.md)    |
| Image   | `list`, `init`, `create`, `create-from-template`, `activate`, `deactivate`, `delete`, `status`, `set-max-session`, `warmup-status` | Image lifecycle  | [→](docs/en/image.md)   |
| API Key | `create`, `enable`, `disable`, `delete`, `list`, `concurrency set`, `describe-key-content`                                         | Key management   | [→](docs/en/apikey.md)  |
| Network | `package list`                                                                                                                     | Network config   | [→](docs/en/network.md) |
| Skills  | `push`, `update`, `show`, `list`, `delete`                                                                                         | Skill management | [→](docs/en/skills.md)  |
| Docker  | `login`, `tag`, `push`, `share`, `unshare`, `list-shares`                                                                          | Docker registry  | [→](docs/en/docker.md)  |

Full command reference → [docs/en/README.md](docs/en/README.md)

---

## Documentation

| Topic                          | Doc                                            |
| ------------------------------ | ---------------------------------------------- |
| Installation & troubleshooting | [installation.md](docs/en/installation.md)     |
| Authentication & env vars      | [authentication.md](docs/en/authentication.md) |
| Image workflow (end-to-end)    | [image-workflow.md](docs/en/image-workflow.md) |
| Image management               | [image.md](docs/en/image.md)                   |
| Docker operations              | [docker.md](docs/en/docker.md)                 |
| API key management             | [apikey.md](docs/en/apikey.md)                 |
| RAM permissions (sub-accounts) | [ram-permissions.md](docs/en/ram-permissions.md) |
| FAQ                            | [faq.md](docs/en/faq.md)                       |

---

## Authentication

AccessKey is recommended for scripts and CI. The CLI also supports STS and OAuth (not recommended). See [Authentication & Environment](docs/en/authentication.md) for details.

The main Alibaba Cloud account does **not** require any additional permission configuration. If you are using a RAM sub-account with AK/SK authentication, grant the required permissions via the [RAM console](https://ram.console.aliyun.com/users) — see [RAM Permissions](docs/en/ram-permissions.md) for the complete policy list.

---

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

---

## License

This project is licensed under the Apache License 2.0 — see the [LICENSE](LICENSE) file for details.
