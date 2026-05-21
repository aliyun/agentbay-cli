[English](../../en/core.md) | **中文**

# 核心命令 — `agentbay`

版本信息与认证相关命令。

## 命令

### `agentbay version`

显示版本号、Git Commit、构建日期、当前环境与 Endpoint。

```bash
agentbay version
```

**输出：**

```
AgentBay CLI version x.x.x
Git commit: xxxxxxx
Build date: 2025-xx-xx
Environment: production
Endpoint: xiaoying.cn-shanghai.aliyuncs.com
```

---

### `agentbay login`（已废弃，后续版本将移除）

> 警告：`agentbay login` **已不推荐使用，后续版本将被废弃移除**。请改用 AccessKey 或 STS 环境变量方式。

打开浏览器进行阿里云 OAuth 认证。在浏览器中完成登录后返回终端。

```bash
agentbay login
```

**输出：**

```
Starting AgentBay authentication...
Opening browser for authentication...
...
Authentication successful!
You are now logged in to AgentBay!
```

**注意事项：**

- 需要浏览器且能访问 `signin.aliyun.com`（国际站为 `signin.alibabacloud.com`）。
- OAuth 回调服务器默认运行在 `localhost:3001`。
- 同时设置了 AccessKey 环境变量与 OAuth Token 时，CLI 优先使用 AccessKey 调用 API。

---

### `agentbay logout`

退出 AgentBay——注销服务端 OAuth 会话并清除本地凭证。

```bash
agentbay logout
```

**注意事项：**

- 清除 CLI 配置文件中存储的 **OAuth** Token。
- **不会**取消设置环境变量——如果 `AGENTBAY_ACCESS_KEY_ID` 和 `AGENTBAY_ACCESS_KEY_SECRET` 仍存在，命令可能仍通过 AccessKey 保持认证状态。
