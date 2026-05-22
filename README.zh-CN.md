# AgentBay CLI

[English](README.md) | **中文**

AgentBay 服务的命令行工具。

---

## 概述

AgentBay CLI 是基于 Cobra 框架的命令行工具，通过阿里云 OpenAPI 与 AgentBay 服务交互。提供镜像管理、API Key 管理、网络管理、技能管理、Docker 操作及灵活的认证方式。

> 当前 CLI 版本仅支持创建和激活 **CodeSpace** 类型的镜像。

---

## 安装

```bash
# macOS / Linux（Homebrew）
brew tap aliyun/agentbay && brew install agentbay

# Windows（PowerShell）
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"

# 验证
agentbay version
```

详见 [安装指南](docs/zh/installation.md)。

---

## 认证方式

**AccessKey（推荐用于脚本/CI）：**

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

STS、OAuth（不推荐使用）及环境变量详见 [认证与环境](docs/zh/authentication.md)。

---

## RAM 账号接口权限（仅 RAM 子账号需要配置）

> 阿里云**主账号**无需任何额外权限配置。
> 本节仅适用于使用 AK/SK 认证的 **RAM 子账号**。

如果使用 RAM 子账号的 AK/SK，请在 [RAM 控制台](https://ram.console.aliyun.com/users) 为该账号授予所需接口权限。

### `apikey` 命令分组

| OpenAPI Action | 所需权限 | 调用命令 |
|---|---|---|
| `CreateApiKey` | `agentbay:CreateApiKey` | `apikey create` |
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` | `apikey enable`、`apikey disable`、`apikey delete`、`apikey list`、`apikey concurrency set` |
| `DescribeApiKeys` | `agentbay:DescribeApiKeys` | `apikey delete`、`apikey list` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` | `apikey enable`、`apikey disable`、`apikey delete` |
| `DeleteApiKey` | `agentbay:DeleteApiKey` | `apikey delete` |
| `ModifyMcpApiKeyConfig` | `agentbay:ModifyMcpApiKeyConfig` | `apikey concurrency set` |
| `DescribeKeyContent` | `agentbay:DescribeKeyContent` | `apikey describe-key-content` |

**RAM Policy 示例（`apikey` 命令完整授权）：**

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

> 如果只使用特定命令，请参考 [API Key 文档](docs/zh/apikey.md) 中各命令的**涉及接口**章节，仅授予所需的最小权限。

---

## 命令概览

| 分组 | 命令 | 说明 | 详情 |
|------|------|------|------|
| 核心 | `version`, `login`, `logout` | 版本与认证 | [→](docs/zh/core.md) |
| 镜像 | `list`, `init`, `create`, `create-from-template`, `activate`, `deactivate`, `delete`, `status`, `set-max-session`, `warmup-status` | 镜像生命周期 | [→](docs/zh/image.md) |
| API Key | `create`, `enable`, `disable`, `delete`, `list`, `concurrency set`, `describe-key-content` | 密钥管理 | [→](docs/zh/apikey.md) |
| 网络 | `package list` | 网络配置 | [→](docs/zh/network.md) |
| 技能 | `push`, `show`, `list` | 技能管理 | [→](docs/zh/skills.md) |
| Docker | `login`, `tag`, `push` | Docker 仓库 | [→](docs/zh/docker.md) |

---

## 快速入门

```bash
# 1. 完成认证（推荐 AccessKey）
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. 创建 API Key（账户需先完成实名认证）
agentbay apikey create "my-api-key"

# 3. 查看已创建的 API Key，从输出中获取 API Key（akm-xxxxxxxxxxxxxxxx）
agentbay apikey list

# 4. 临时不需要时禁用 API Key
agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx

# 5. 需要时重新启用
agentbay apikey enable --api-key akm-xxxxxxxxxxxxxxxx

# 6. 永久删除 API Key（必须先 DISABLED；--yes 跳过确认）
agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx --yes
```

> **提示：** 自动化脚本可使用 `--api-key-id ak-xxxxxxxxxxxxxxxx`（由 `apikey create` 返回）代替 `--api-key`。详见 [API Key 文档](docs/zh/apikey.md#术语说明)。

完整命令说明请参考 [命令参考](docs/zh/README.md)。

---

## 更新日志

查看 [CHANGELOG.md](CHANGELOG.md) 了解版本更新记录。

---

## 注意事项

- 同时设置了 AccessKey 环境变量与 OAuth Token 时，CLI 优先使用 AccessKey 调用 API。
- 系统镜像始终可用，无需激活；只有用户镜像必须先激活才能使用。
- API Key 创建前账户需完成实名认证，且每个 Key 必须使用唯一名称。
- 在非交互式环境中执行破坏性命令（`apikey delete`、`image delete`）时，请使用 `--yes` / `-y` 跳过确认提示。

---

## 许可证

本项目基于 Apache License 2.0 协议开源，详见 [LICENSE](LICENSE) 文件。
