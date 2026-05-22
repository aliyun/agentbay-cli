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

## RAM Permissions (RAM Sub-accounts Only)

> The main Alibaba Cloud account does **not** require any additional permission configuration.
> This section applies only to **RAM sub-accounts** using AK/SK authentication.

If you are using a RAM sub-account's AK/SK, grant the required permissions via the [RAM console](https://ram.console.aliyun.com/users).

### `apikey` Command Group

| OpenAPI Action | Required Permission | Used By |
|---|---|---|
| `CreateApiKey` | `agentbay:CreateApiKey` | `apikey create` |
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` | `apikey enable`, `apikey disable`, `apikey delete`, `apikey list`, `apikey concurrency set` |
| `DescribeApiKeys` | `agentbay:DescribeApiKeys` | `apikey delete`, `apikey list` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` | `apikey enable`, `apikey disable`, `apikey delete` |
| `DeleteApiKey` | `agentbay:DeleteApiKey` | `apikey delete` |
| `ModifyMcpApiKeyConfig` | `agentbay:ModifyMcpApiKeyConfig` | `apikey concurrency set` |
| `DescribeKeyContent` | `agentbay:DescribeKeyContent` | `apikey describe-key-content` |

**RAM Policy example (full access to `apikey` commands):**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:CreateApiKey",
        "agentbay:DescribeMcpApiKey",
        "agentbay:DescribeApiKeys",
        "agentbay:ModifyApiKeyStatus",
        "agentbay:DeleteApiKey",
        "agentbay:ModifyMcpApiKeyConfig",
        "agentbay:DescribeKeyContent"
      ],
      "Resource": "*"
    }
  ]
}
```

> If you only use specific commands, refer to the **Involved APIs** section in [API Key docs](docs/en/apikey.md) and grant only the required subset.

### `image` Command Group

| OpenAPI Action | Required Permission | Used By |
|---|---|---|
| `ListMcpImages` | `agentbay:ListMcpImages` | `image list`, `image deactivate` |
| `GetMcpImageInfo` | `agentbay:GetMcpImageInfo` | `image create`, `image activate`, `image deactivate`, `image delete`, `image status`, `image set-max-session` |
| `GetDockerFileStoreCredential` | `agentbay:GetDockerFileStoreCredential` | `image create` |
| `CreateDockerImageTask` | `agentbay:CreateDockerImageTask` | `image create` |
| `GetDockerImageTask` | `agentbay:GetDockerImageTask` | `image create` |
| `CreateImageFromTemplate` | `agentbay:CreateImageFromTemplate` | `image create-from-template` |
| `DescribeInstanceTypes` | `agentbay:DescribeInstanceTypes` | `image activate` |
| `DescribeMcpPolicyData` | `agentbay:DescribeMcpPolicyData` | `image activate` |
| `CreateMcpPolicyData` | `agentbay:CreateMcpPolicyData` | `image activate` |
| `ModifyMcpPolicyData` | `agentbay:ModifyMcpPolicyData` | `image activate` |
| `DescribeOfficeSites` | `agentbay:DescribeOfficeSites` | `image activate` |
| `SaveMcpPolicyData` | `agentbay:SaveMcpPolicyData` | `image activate` |
| `CreateResourceGroup` | `agentbay:CreateResourceGroup` | `image activate` |
| `DeleteResourceGroup` | `agentbay:DeleteResourceGroup` | `image deactivate` |
| `DeleteMcpImage` | `agentbay:DeleteMcpImage` | `image delete` |
| `GetDockerfileTemplate` | `agentbay:GetDockerfileTemplate` | `image init` |
| `BatchCreateHideResourceGroupsWithMaxSession` | `agentbay:BatchCreateHideResourceGroupsWithMaxSession` | `image set-max-session` |
| `DescribeWarmUpStatusOpen` | `agentbay:DescribeWarmUpStatusOpen` | `image warmup-status` |

**RAM Policy example (full access to `image` commands):**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:ListMcpImages",
        "agentbay:GetMcpImageInfo",
        "agentbay:GetDockerFileStoreCredential",
        "agentbay:CreateDockerImageTask",
        "agentbay:GetDockerImageTask",
        "agentbay:CreateImageFromTemplate",
        "agentbay:DescribeInstanceTypes",
        "agentbay:DescribeMcpPolicyData",
        "agentbay:CreateMcpPolicyData",
        "agentbay:ModifyMcpPolicyData",
        "agentbay:DescribeOfficeSites",
        "agentbay:SaveMcpPolicyData",
        "agentbay:CreateResourceGroup",
        "agentbay:DeleteResourceGroup",
        "agentbay:DeleteMcpImage",
        "agentbay:GetDockerfileTemplate",
        "agentbay:BatchCreateHideResourceGroupsWithMaxSession",
        "agentbay:DescribeWarmUpStatusOpen"
      ],
      "Resource": "*"
    }
  ]
}
```

> If you only use specific commands, refer to the **Involved APIs** section in [Image docs](docs/en/image.md) and grant only the required subset.

### `network` Command Group

| OpenAPI Action | Required Permission | Used By |
|---|---|---|
| `DescribeNetworkPackages` | `agentbay:DescribeNetworkPackages` | `network package list` |

**RAM Policy example:**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:DescribeNetworkPackages"
      ],
      "Resource": "*"
    }
  ]
}
```

### `skills` Command Group

| OpenAPI Action | Required Permission | Used By |
|---|---|---|
| `GetMarketSkillCredential` | `agentbay:GetMarketSkillCredential` | `skills push` |
| `CreateMarketSkill` | `agentbay:CreateMarketSkill` | `skills push` |
| `DescribeMarketSkillDetail` | `agentbay:DescribeMarketSkillDetail` | `skills show` |

**RAM Policy example:**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:GetMarketSkillCredential",
        "agentbay:CreateMarketSkill",
        "agentbay:DescribeMarketSkillDetail"
      ],
      "Resource": "*"
    }
  ]
}
```

### `docker` Command Group

| OpenAPI Action | Required Permission | Used By |
|---|---|---|
| `GetACRRepoCredential` | `agentbay:GetACRRepoCredential` | `docker login` |

**RAM Policy example:**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:GetACRRepoCredential"
      ],
      "Resource": "*"
    }
  ]
}
```

> `docker tag` and `docker push` are wrappers around the native `docker` CLI and do not call any AgentBay API directly.

---

## Command Overview

| Group | Commands | Description | Details |
|-------|----------|-------------|---------|
| Core | `version`, `login`, `logout` | Version & auth | [→](docs/en/core.md) |
| Image | `list`, `init`, `create`, `create-from-template`, `activate`, `deactivate`, `delete`, `status`, `set-max-session`, `warmup-status` | Image lifecycle | [→](docs/en/image.md) |
| API Key | `create`, `enable`, `disable`, `delete`, `list`, `concurrency set`, `describe-key-content` | Key management | [→](docs/en/apikey.md) |
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
