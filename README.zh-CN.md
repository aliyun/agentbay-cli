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

### 更新

**macOS / Linux（Homebrew）—— 快速通道（推荐用于日常升级）：**

```bash
git -C "$(brew --repository aliyun/agentbay)" pull --ff-only && brew upgrade agentbay
```

只刷新 `aliyun/agentbay` 这一个 tap，然后升级 agentbay。跳过 Homebrew 的全量元数据同步（`formula.jws.json` / `cask.jws.json` 等几十 MB 的 JSON 下载，以及 brew 自身的升级），通常几秒钟就能完成。

**macOS / Linux（Homebrew）—— `brew` 本身报错时的回退方案：**

```bash
brew update && brew upgrade agentbay
```

会刷新 Homebrew 自身、所有已安装的 tap，以及 core formula 元数据，然后再升级 agentbay。更慢但更彻底 —— 在快速通道失败时使用（例如长时间没刷新 brew 之后，或者 Homebrew 出现破坏性更新时）。

**Windows（PowerShell）：** 重新执行安装命令即可原地升级。

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

### 卸载

```bash
# macOS / Linux（Homebrew）
brew uninstall agentbay
brew untap aliyun/agentbay   # 可选
```

```powershell
# Windows（PowerShell）
# 注意：如果安装时指定了 -InstallPath 或设置了 $env:AGENTBAY_PATH，
# 请把下面的 "$env:LOCALAPPDATA\agentbay" 替换成实际安装目录。
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
# 请重启 PowerShell 让 PATH 变更生效。
```

> **Homebrew 提示：** 首次执行 `brew install agentbay` 会从源码编译，并自动安装 Go 作为构建依赖，整个过程可能需要几分钟。后续升级会复用缓存。

详见 [安装指南](docs/zh/installation.md)（含预编译二进制及故障排除）。

---

## 认证方式

**AccessKey（推荐用于脚本/CI）：**

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

STS、OAuth（不推荐使用）及环境变量详见 [认证与环境](docs/zh/authentication.md)。

---

## 文档导航

| 文档                                                | 说明                                       |
| --------------------------------------------------- | ------------------------------------------ |
| [安装指南](docs/zh/installation.md)                 | 详细安装步骤与故障排除                     |
| [认证与环境](docs/zh/authentication.md)             | AccessKey、STS、OAuth 及环境变量           |
| [镜像创建与共享完整流程](docs/zh/image-workflow.md) | 从 Dockerfile 模板到跨账号共享的端到端教程 |
| [镜像管理](docs/zh/image.md)                        | 镜像生命周期管理命令参考                   |
| [Docker 操作](docs/zh/docker.md)                    | ACR 登录、镜像推送与共享                   |
| [API Key 管理](docs/zh/apikey.md)                   | 密钥创建、启用、禁用、删除                 |
| [RAM 权限配置](docs/zh/ram-permissions.md)          | 各命令组所需 RAM 权限                      |
| [常见问题](docs/zh/faq.md)                          | FAQ                                        |

完整命令说明请参考 [命令参考](docs/zh/README.md)。

---

## RAM 账号接口权限（仅 RAM 子账号需要配置）

阿里云**主账号**无需额外配置。若使用 **RAM 子账号**的 AK/SK 认证，请在 [RAM 控制台](https://ram.console.aliyun.com/users) 为该账号授予所需接口权限。

各命令组所需的完整权限列表及 Policy JSON 示例，请参考 [RAM 权限配置](docs/zh/ram-permissions.md)。

---

## 命令概览

| 分组    | 命令                                                                                                                               | 说明         | 详情                    |
| ------- | ---------------------------------------------------------------------------------------------------------------------------------- | ------------ | ----------------------- |
| 核心    | `version`, `login`, `logout`                                                                                                       | 版本与认证   | [→](docs/zh/core.md)    |
| 镜像    | `list`, `init`, `create`, `create-from-template`, `activate`, `deactivate`, `delete`, `status`, `set-max-session`, `warmup-status` | 镜像生命周期 | [→](docs/zh/image.md)   |
| API Key | `create`, `enable`, `disable`, `delete`, `list`, `concurrency set`, `describe-key-content`                                         | 密钥管理     | [→](docs/zh/apikey.md)  |
| 网络    | `package list`                                                                                                                     | 网络配置     | [→](docs/zh/network.md) |
| 技能    | `push`, `update`, `show`, `list`, `delete`                                                                                         | 技能管理     | [→](docs/zh/skills.md)  |
| Docker  | `login`, `tag`, `push`, `share`, `unshare`, `list-shares`                                                                          | Docker 仓库  | [→](docs/zh/docker.md)  |

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
