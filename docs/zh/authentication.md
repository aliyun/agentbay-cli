[English](../../en/authentication.md) | **中文**

# 认证与环境

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

### 3. OAuth 登录（不推荐使用）

> 提示：`agentbay login` **不推荐使用**。请优先使用 AccessKey 或 STS 方式。

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

## 环境切换

AgentBay CLI 支持通过 `AGENTBAY_ENV` 环境变量在生产和预发环境间切换。

### 验证当前环境

```bash
agentbay version
```

输出中的 `Environment` 和 `Endpoint` 字段会反映当前配置。

### 国际站（阿里云国际站）

设置 `AGENTBAY_ENV=international` 即可使用国际站生产环境。CLI 将自动使用：

- Endpoint：`xiaoying.ap-southeast-1.aliyuncs.com`
- OAuth：signin.alibabacloud.com 及默认国际站 OAuth Client ID

除非需要覆盖默认值，否则无需设置 `AGENTBAY_OAUTH_REGION` 或 `AGENTBAY_OAUTH_CLIENT_ID`。
