[English](../en/core.md) | **中文**

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

### `agentbay login`

> **仅支持主账号。** RAM 子账号和 RAM 角色登录时会被拒绝 —— 推荐使用 AccessKey 环境变量方式；如需继续使用 OAuth，请先在浏览器访问 [阿里云官网](https://www.aliyun.com/) 并退出当前阿里云登录态，再重新执行 `agentbay login` 并选择/登录阿里云主账号（详见 [认证与环境](authentication.md)）。

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
