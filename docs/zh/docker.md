[English](../en/docker.md) | **中文**

# Docker 操作 — `agentbay docker`

这组命令封装本地 `docker` CLI，用于与 AgentBay ACR 镜像仓库交互。

## 命令

### `docker login`

通过 `GetACRRepoCredential` 接口获取临时凭证并登录 AgentBay ACR 仓库。凭证信息（`RegistryUrl`、`Namespace`、`RepoName`、`ImageTag`）会被缓存供 `tag` / `push` 使用。

```bash
agentbay docker login
```

**注意事项：**

- 调用 `GetACRRepoCredential` 获取临时 ACR 凭证。
- 凭证会缓存到本地，供后续 `tag` 和 `push` 命令使用。

---

### `docker tag`

为本地镜像打 tag 以推送到 AgentBay ACR。目标镜像名将自动构造为 `$RegistryUrl/$Namespace/$RepoName:<target-tag>`。

```bash
agentbay docker tag myapp:latest v1.0
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `<source-image>` | string | 是 | 本地镜像名（含 tag） |
| `<target-tag>` | string | 是 | 目标 tag |

> 必须先执行 `agentbay docker login`。

---

### `docker push`

推送已 tag 的镜像到 AgentBay ACR。镜像名必须匹配 `$RegistryUrl/$Namespace/$RepoName[:tag]`，否则将被拒绝。

```bash
agentbay docker push <registry>/<namespace>/<repo>:v1.0
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `<image>` | string | 是 | 完整的 ACR 镜像名 |

> 必须先执行 `agentbay docker login`。
