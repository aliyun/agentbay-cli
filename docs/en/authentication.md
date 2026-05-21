[中文](../../zh/authentication.md) | **English**

# Authentication & Environment

## Authentication Methods

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

## Environment Switching

AgentBay CLI supports switching between production and pre-release environments using the `AGENTBAY_ENV` environment variable.

### Verify Current Environment

```bash
agentbay version
```

Output includes `Environment` and `Endpoint` fields reflecting the current configuration.

### International (Alibaba Cloud International)

For international production (ap-southeast-1, alibabacloud.com), set `AGENTBAY_ENV=international`. The CLI then uses:

- Endpoint: `xiaoying.ap-southeast-1.aliyuncs.com`
- OAuth: signin.alibabacloud.com and the default international OAuth client ID

You do not need to set `AGENTBAY_OAUTH_REGION` or `AGENTBAY_OAUTH_CLIENT_ID` unless you want to override them.
