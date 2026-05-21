[中文](../../zh/core.md) | **English**

# Core Commands — `agentbay`

Core commands for version info and authentication.

## Commands

### `agentbay version`

Show version, git commit, build date, current environment and endpoint.

```bash
agentbay version
```

**Output:**

```
AgentBay CLI version x.x.x
Git commit: xxxxxxx
Build date: 2025-xx-xx
Environment: production
Endpoint: xiaoying.cn-shanghai.aliyuncs.com
```

---

### `agentbay login` _(not recommended)_

> NOTE: `agentbay login` is **not recommended**. Please prefer AccessKey or STS environment variables instead.

Opens a browser for Aliyun OAuth authentication. After completing the login in the browser, return to the terminal.

```bash
agentbay login
```

**Output:**

```
Starting AgentBay authentication...
Opening browser for authentication...
...
Authentication successful!
You are now logged in to AgentBay!
```

**Notes:**

- Requires a browser and network access to `signin.aliyun.com` (or `signin.alibabacloud.com` for international).
- The OAuth callback server runs on `localhost:3001` by default.
- When both AccessKey env vars and OAuth tokens are present, the CLI prefers AccessKey for API calls.

---

### `agentbay logout`

Log out from AgentBay — invalidate the OAuth session on the server and clear local credentials.

```bash
agentbay logout
```

**Notes:**

- Clears **OAuth** tokens stored in the CLI config file.
- Does **not** unset environment variables — if `AGENTBAY_ACCESS_KEY_ID` and `AGENTBAY_ACCESS_KEY_SECRET` are still set, commands may remain authenticated via AccessKey.
