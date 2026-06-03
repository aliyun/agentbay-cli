# AgentBay CLI

[English](README.md) | **中文**

> 🤖 **LLM 友好**:本项目提供 [llms.txt](llms.txt) 与 [llms-full.txt](llms-full.txt),便于 AI 助手理解。

阿里云 AgentBay 服务的命令行工具 —— 镜像生命周期、API Key、Docker 与技能管理。

> 当前 CLI 版本仅支持创建和激活 **CodeSpace** 类型的镜像。

---

## 功能特性

- **镜像生命周期** —— 基于 Dockerfile/模板创建、激活、查询、删除
- **Docker 集成** —— ACR 登录、镜像推送、跨账号共享/取消共享
- **API Key 管理** —— 创建、启用/禁用、删除、并发配置
- **技能与网络** —— 技能推送/更新、网络包查询
- **多种认证方式** —— AccessKey（AK/SK）、STS、OAuth
- **跨平台支持** —— macOS、Linux、Windows

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

> 首次执行 `brew install agentbay` 会从源码编译，并自动安装 Go 作为构建依赖，整个过程可能需要几分钟。后续升级会复用缓存。

<details>
<summary><b>更新</b></summary>

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

</details>

<details>
<summary><b>卸载</b></summary>

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

</details>

详见 [安装指南](docs/zh/installation.md)（含预编译二进制及故障排除）。

---

## 快速入门 —— 60 秒上手 API Key

```bash
# 1. 完成认证（推荐 AccessKey）
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. 创建 API Key（账户需先完成实名认证）
agentbay apikey create "my-api-key"

# 3. 查询 / 禁用 / 重新启用 / 删除
agentbay apikey list
agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx
agentbay apikey enable  --api-key akm-xxxxxxxxxxxxxxxx
agentbay apikey delete  --api-key akm-xxxxxxxxxxxxxxxx --yes
```

> **提示：** 自动化脚本可使用 `--api-key-id ak-xxxxxxxxxxxxxxxx`（由 `apikey create` 返回）代替 `--api-key`。详见 [API Key 文档](docs/zh/apikey.md#术语说明)。
>
> 使用 RAM 子账号？请先参照 [RAM 权限配置](docs/zh/ram-permissions.md) 授予所需权限。

---

## 进阶教程 —— 镜像创建与共享

基于 Dockerfile 模板构建自定义镜像、推送至 ACR，并按需跨阿里云账号共享。

```bash
# ── 镜像创建（任何账号都能独立完成） ─────────────────────────
agentbay image init --sourceImageId aio-ubuntu-2404            # 1. 下载 Dockerfile 模板
agentbay docker login                                          # 2. 登录 ACR（临时凭证，~1 小时）
docker build -t <registry>/<namespace>/<uid>:<tag> -f Dockerfile .   # 3. 本地构建
docker push  <registry>/<namespace>/<uid>:<tag>                # 4. 推送到 ACR
agentbay image create-from-template \                          # 5. 创建自定义镜像
  --source-image /<namespace>/<uid>:<tag> \
  --name my-image --imageId aio-ubuntu-2404

# ── 镜像共享（可选，A 账号 → B 账号） ────────────────────────
# A 账号（共享方）：
agentbay docker share <ACCOUNT_B_UID>                          # 1. 共享仓库给 B 账号
agentbay docker list-shares --direction Outgoing               # 2. 确认共享生效
# B 账号（接收方）：
agentbay docker list-shares --direction Incoming               # 3. 查看收到的共享
agentbay image create-from-template ...                        # 4. 基于 A 的镜像创建自己的镜像（复用上面 Step 5）
```

→ **完整流程**（含真实参数示例、命令输出、故障排查）：**[镜像创建与共享](docs/zh/image-workflow.md)**

> **前置条件：** 本机已安装 Docker。macOS 推荐使用 [OrbStack](https://orbstack.dev/) —— 轻量、快速，资源占用远低于 Docker Desktop。

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

完整命令说明请参考 [命令参考](docs/zh/README.md)

---

## 文档导航

| 主题                       | 文档                                              |
| -------------------------- | ------------------------------------------------- |
| 安装与故障排除             | [installation.md](docs/zh/installation.md)        |
| 认证与环境变量             | [authentication.md](docs/zh/authentication.md)    |
| 镜像创建与共享             | [image-workflow.md](docs/zh/image-workflow.md)    |
| 镜像管理                   | [image.md](docs/zh/image.md)                      |
| Docker 操作                | [docker.md](docs/zh/docker.md)                    |
| API Key 管理               | [apikey.md](docs/zh/apikey.md)                    |
| RAM 权限配置（子账号专用） | [ram-permissions.md](docs/zh/ram-permissions.md)  |
| 常见问题                   | [faq.md](docs/zh/faq.md)                          |

---

## 认证方式

推荐使用 AccessKey（脚本/CI/RAM 子账号必选）。CLI 同时支持 STS 与 OAuth 登录（`agentbay login`，**仅支持阿里云主账号** —— RAM 子账号会被拒绝）。详见 [认证与环境](docs/zh/authentication.md)。

阿里云**主账号**无需额外配置。若使用 **RAM 子账号**的 AK/SK 认证，可以到 [RAM 控制台](https://ram.console.aliyun.com/policies) 先新建或修改权限策略，再将策略授权给对应的 RAM 子账号 —— 完整 Policy 列表参见 [RAM 权限配置](docs/zh/ram-permissions.md)。

---

## 更新日志

查看 [CHANGELOG.md](CHANGELOG.md) 了解版本更新记录。

---

## 许可证

本项目基于 Apache License 2.0 协议开源，详见 [LICENSE](LICENSE) 文件。
