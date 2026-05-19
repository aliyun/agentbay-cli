# AgentBay CLI

[English](README.md) | **中文**

AgentBay 服务的命令行工具。

---

## 概述

AgentBay CLI 是基于 Cobra 框架的命令行工具，通过阿里云 OpenAPI 与 AgentBay 服务交互。功能包括：

- **镜像管理** —— 创建、列出、激活 / 停用 / 删除自定义镜像，查询生命周期状态，配置会话并发数
- **API Key 管理** —— 创建 / 启用 / 禁用 / 删除密钥，设置每个密钥的并发上限
- **网络管理** —— 按区域查询网络包及其 EIP 绑定信息
- **技能管理** —— 推送本地技能包，按 ID 查看技能详情
- **Docker 操作** —— 登录 ACR、为 AgentBay 打 tag 并推送镜像
- **认证** —— 推荐使用 AccessKey / STS 环境变量；OAuth 登录仅供本地开发使用
- **配置** —— 安全的 Token 存储、自动刷新、多环境支持

> 当前 CLI 版本仅支持创建和激活 **CodeSpace** 类型的镜像。

---

## 安装

预编译的二进制文件可在 `bin/` 与 `packages/` 目录下找到。macOS / Linux 也可以通过 Homebrew tap 安装（参考 `homebrew/agentbay.rb`）。

```bash
# 验证安装
agentbay version
```

---

## 认证方式

CLI 支持三种认证方式。**生产脚本与 CI/CD 推荐使用 AccessKey 或 STS。**

> 优先级：`AGENTBAY_ACCESS_KEY_ID` / `AGENTBAY_ACCESS_KEY_SECRET` 环境变量 > 本地存储的 OAuth Token。

### 1. AccessKey（推荐）

设置以下环境变量。这是自动化、脚本和 CI/CD 场景下的首选方式：

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

### 2. STS 临时凭证（推荐用于短期会话）

使用 STS（Security Token Service）临时凭证时，除了 AK/SK 外还需设置 Session Token：

```bash
export AGENTBAY_ACCESS_KEY_ID="STS.xxx"
export AGENTBAY_ACCESS_KEY_SECRET="your-sts-secret"
export AGENTBAY_ACCESS_KEY_SESSION_TOKEN="your-sts-session-token"
```

### 3. OAuth 登录（已废弃，不推荐使用）

> 警告：`agentbay login` **已不推荐使用，后续版本将被废弃移除**。请改用 AccessKey 或 STS 方式。

```bash
agentbay login    # 打开浏览器进行 OAuth 登录
agentbay logout   # 注销服务端会话并清理本地凭证
```

---

## 环境变量

如无特殊说明，所有 AgentBay CLI 环境变量均为可选。

### 环境变量设置方式

**方式 1：当前终端 export**（关闭终端后失效）

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

**方式 2：工作目录下的 `.env` 文件**（推荐，项目级配置）

CLI 启动时会自动加载当前工作目录下的 `.env` 文件（通过 `godotenv`）。创建 `.env` 文件：

```dotenv
AGENTBAY_ACCESS_KEY_ID=your-access-key-id
AGENTBAY_ACCESS_KEY_SECRET=your-access-key-secret
AGENTBAY_ENV=production
```

> 请勿将 `.env` 提交到版本控制——确保已加入 `.gitignore`。

**方式 3：Shell 配置文件**（全部终端会话持久生效）

追加到 `~/.bashrc`、`~/.zshrc` 或对应文件中：

```bash
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
```

然后重新加载：`source ~/.zshrc`（或打开新终端）。

**方式 4：命令前内联指定**（仅对当次命令生效）

```bash
AGENTBAY_ENV=prerelease agentbay image list
```

### 认证相关

| 变量                                | 说明                                                   |
| ----------------------------------- | ------------------------------------------------------ |
| `AGENTBAY_ACCESS_KEY_ID`            | AccessKey ID（使用 STS 时为以 `STS.` 开头的 Token ID） |
| `AGENTBAY_ACCESS_KEY_SECRET`        | AccessKey Secret（或 STS Secret）                      |
| `AGENTBAY_ACCESS_KEY_SESSION_TOKEN` | STS Session Token（仅在使用 STS 凭证时需要）           |

### 环境选择

| 变量           | 默认值       | 可选值                                                                                                                                     |
| -------------- | ------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `AGENTBAY_ENV` | `production` | `production` / `prod`、`prerelease` / `pre` / `staging`、`international` / `intl` / `prod-international`、`international-pre` / `intl-pre` |

```bash
# 切换到预发环境
export AGENTBAY_ENV=prerelease
# 切换到国际站生产环境
export AGENTBAY_ENV=international
```

各环境对应的 Endpoint：

| 环境                | Endpoint                                   |
| ------------------- | ------------------------------------------ |
| `production`        | `xiaoying.cn-shanghai.aliyuncs.com`        |
| `prerelease`        | `xiaoying-pre.cn-hangzhou.aliyuncs.com`    |
| `international`     | `xiaoying.ap-southeast-1.aliyuncs.com`     |
| `international-pre` | `xiaoying-pre.ap-southeast-1.aliyuncs.com` |

### 高级配置

| 变量                       | 说明                                                        |
| -------------------------- | ----------------------------------------------------------- |
| `AGENTBAY_CLI_ENDPOINT`    | 覆盖当前环境的默认 API Endpoint                             |
| `AGENTBAY_CLI_TIMEOUT_MS`  | API 请求超时时间（毫秒）                                    |
| `AGENTBAY_CLI_CONFIG_DIR`  | 覆盖默认配置目录（默认 `~/.agentbay`）                      |
| `AGENTBAY_OAUTH_CLIENT_ID` | 覆盖默认的 OAuth Client ID（仅对 `agentbay login` 生效）    |
| `AGENTBAY_OAUTH_REGION`    | 覆盖 OAuth 区域（`cn` 或 `intl`）                           |
| `AGENTBAY_API_URL`         | _(Legacy)_ 等同于 `AGENTBAY_CLI_ENDPOINT`，仅为向后兼容保留 |

---

## 命令参考

可通过 `agentbay --help` 查看所有命令。命令分为两个组：**Core Commands**（核心命令）与 **Management Commands**（管理命令）。

### 核心命令

#### `agentbay version`

显示版本号、Git Commit、构建日期、当前环境与 Endpoint。

```bash
agentbay version
```

#### `agentbay logout`

退出 AgentBay（注销服务端 OAuth 会话并清除本地凭证）。

```bash
agentbay logout
```

#### `agentbay login`（已废弃，后续版本将移除）

详见 [认证方式](#认证方式)。

---

### 镜像管理 —— `agentbay image`

#### `image list`

列出可用的 AgentBay 镜像。

```bash
agentbay image list                      # 用户镜像（默认）
agentbay image list --include-system     # 用户镜像 + 系统镜像
agentbay image list --system-only        # 仅系统镜像
agentbay image list --os-type Linux      # 按 OS 过滤：Linux / Android / Windows
agentbay image list --page 2 --size 5    # 分页
```

#### `image init`

从云端下载 Dockerfile 模板到当前目录。

```bash
agentbay image init --sourceImageId code-space-debian-12
agentbay image init -i code-space-debian-12
```

> 下载的 Dockerfile 前 N 行由系统定义，不可修改，仅可编辑第 N+1 行之后的内容。

生产环境可用的 `sourceImageId`：

- `code-space-debian-12`
- `code-space-debian-12-enhanced`

#### `image create`（已废弃，请改用 `create-from-template`）

> 警告：`image create` **已不推荐使用，后续版本将被移除**。如需创建自定义镜像，请改用 [`agentbay image create-from-template`](#image-create-from-template)。

基于 Dockerfile 构建自定义镜像。`COPY` / `ADD` 引用的文件会被自动解析并上传。

```bash
agentbay image create myapp --dockerfile ./Dockerfile --imageId code-space-debian-12
agentbay image create myapp -f ./Dockerfile -i code-space-debian-12
```

#### `image create-from-template`

基于系统镜像模板创建自定义镜像（调用 `CreateImageFromTemplate` 接口）。

```bash
agentbay image create-from-template \
  --source-image registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 \
  --name my-custom-image \
  --imageId <system-image-id>

# 短参数形式
agentbay image create-from-template -s registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 -n my-custom-image -i <system-image-id>
```

#### `image activate`

激活用户镜像使其可用。

```bash
# 使用默认资源规格
agentbay image activate imgc-xxxxxxxxxxxxxx

# 指定 CPU 和内存（必须同时指定）
agentbay image activate imgc-xxxxxxxxxxxxxx --cpu 2 --memory 4

# 高级网络
agentbay image activate imgc-xxxxxxxxxxxxxx \
  --network-type ADVANCED \
  --session-bandwidth 100 \
  --dns-address 8.8.8.8 \
  --dns-address 8.8.4.4

# 沙箱生命周期
agentbay image activate imgc-xxxxxxxxxxxxxx \
  --lifecycle-mode auto \
  --lifecycle-max-runtime 3600 \
  --lifecycle-hibernate 1800 \
  --lifecycle-idle-timeout 600

# 指定区域
agentbay image activate imgc-xxxxxxxxxxxxxx --region-id cn-shanghai
```

注意事项：

- `--cpu` 和 `--memory` 必须同时指定。
- `--network-type ADVANCED` 时必须同时指定 `--session-bandwidth` 与 `--dns-address`。
- `--lifecycle-mode` 可选值为 `auto` 或 `manual`。

#### `image deactivate`

停用已激活的用户镜像。

```bash
agentbay image deactivate imgc-xxxxxxxxxxxxxx
```

#### `image delete`

**永久删除**用户镜像。仅已停用的用户镜像可删除。

```bash
agentbay image delete imgc-xxxxxxxxxxxxxx          # 带确认提示
agentbay image delete imgc-xxxxxxxxxxxxxx --yes    # 跳过确认（脚本 / CI）
```

#### `image status`

查询镜像的资源生命周期状态（与 `image create` 时的 Docker 构建任务状态不同）。

```bash
agentbay image status imgc-xxxxxxxxxxxxxx
```

常见状态：`IMAGE_CREATING`、`IMAGE_CREATE_FAILED`、`IMAGE_AVAILABLE`、`RESOURCE_DEPLOYING`、`RESOURCE_PUBLISHED`、`RESOURCE_DELETING`、`RESOURCE_FAILED`、`RESOURCE_CEASED`。

#### `image set-max-session`

设置已激活用户镜像的最大并发会话数。要求镜像处于 `RESOURCE_PUBLISHED` 状态且使用**高级网络**。

```bash
agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 10
```

> 该命令会轮询直到资源组就绪（通常约 5 分钟）。

---

### API Key 管理 —— `agentbay apikey`

> 创建 API Key 前账户必须完成实名认证，且每个 API Key 名称必须唯一。

#### `apikey create`

创建新的 API Key。

```bash
agentbay apikey create --name "my-api-key"
```

#### `apikey enable`

重新启用已禁用的 API Key。

```bash
agentbay apikey enable akm-xxxxxxxxxxxxxxxx
```

#### `apikey disable`

禁用 API Key（禁用后无法用于认证）。

```bash
agentbay apikey disable akm-xxxxxxxxxxxxxxxx
```

#### `apikey delete`

永久删除 API Key。仅 `DISABLED` 状态的 Key 可直接删除；若处于 `ENABLED` 状态会先提示禁用。

```bash
agentbay apikey delete akm-xxxxxxxxxxxxxxxx          # 交互式（带确认）
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes    # 跳过所有提示（脚本 / CI）
```

#### `apikey concurrency set`

设置 API Key 的最大并发会话数。

```bash
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10
```

---

### 网络管理 —— `agentbay network`

#### `network package list`

按区域列出网络包。

```bash
agentbay network package list                              # 默认区域 cn-hangzhou
agentbay network package list --biz-region-id cn-shanghai  # 指定区域
```

---

### 技能管理 —— `agentbay skills`

#### `skills push`

推送本地技能（目录或 `.zip`）到云端。目录形式必须包含带 `name` / `description` frontmatter 的 `SKILL.md`，目录会被打包为 zip 后上传。

```bash
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
```

#### `skills show`

按 ID 查看技能详情。

```bash
agentbay skills show <skill-id>
```

#### `skills list`（占位）

列出云端技能。后端 list 接口尚未提供，该命令目前为占位实现。

---

### Docker 操作 —— `agentbay docker`

这组命令封装本地 `docker` CLI，用于与 AgentBay ACR 镜像仓库交互。

#### `docker login`

通过 `GetACRRepoCredential` 接口获取临时凭证并登录 AgentBay ACR 仓库。凭证信息（`RegistryUrl`、`Namespace`、`RepoName`、`ImageTag`）会被缓存供 `tag` / `push` 使用。

```bash
agentbay docker login
```

#### `docker tag`

为本地镜像打 tag 以推送到 AgentBay ACR。目标镜像名将自动构造为 `$RegistryUrl/$Namespace/$RepoName:<target-tag>`。

```bash
agentbay docker tag myapp:latest v1.0
```

> 必须先执行 `agentbay docker login`。

#### `docker push`

推送已 tag 的镜像到 AgentBay ACR。镜像名必须匹配 `$RegistryUrl/$Namespace/$RepoName[:tag]`，否则将被拒绝。

```bash
agentbay docker push <registry>/<namespace>/<repo>:v1.0
```

> 必须先执行 `agentbay docker login`。

---

## 快速入门

最小可运行流程：

```bash
# 1. 完成认证（推荐 AccessKey）
export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"

# 2. 创建 API Key（账户需先完成实名认证）
agentbay apikey create --name "my-api-key"

# 3. 临时不需要时禁用 API Key
agentbay apikey disable akm-xxxxxxxxxxxxxxxx

# 4. 需要时重新启用
agentbay apikey enable akm-xxxxxxxxxxxxxxxx

# 5. 永久删除 API Key（必须先 DISABLED；--yes 跳过确认）
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes
```

完整命令说明请参考上方 [命令参考](#命令参考) 章节，以及 [User Guide](docs/USER_GUIDE.md)。

---

## 注意事项

- 同时设置了 AccessKey 环境变量与 OAuth Token 时，CLI 优先使用 AccessKey 调用 API。
- 系统镜像始终可用，无需激活；只有用户镜像必须先激活才能使用。
- API Key 创建前账户需完成实名认证，且每个 Key 必须使用唯一名称。
- 在非交互式环境中执行破坏性命令（`apikey delete`、`image delete`）时，请使用 `--yes` / `-y` 跳过确认提示。

---

## 许可证

本项目基于 Apache License 2.0 协议开源，详见 [LICENSE](LICENSE) 文件。
