[English](../en/docker.md) | **中文**

# Docker 操作 — `agentbay docker`

这组命令封装本地 `docker` CLI，用于与 AgentBay ACR 镜像仓库交互。

## 命令

### `docker login`

通过 `GetACRRepoCredential` 接口获取临时凭证并**自动登录本地 docker**（有效期约 1 小时），同时返回该账号专属的镜像上传地址。`docker build` 和 `docker push` 都依赖该登录状态。凭证信息（`RegistryUrl`、`Namespace`、`RepoName`、`ImageTag`）也会被缓存供后续 `agentbay docker tag` / `agentbay docker push` 使用。

```bash
agentbay docker login
```

**输出示例：**

```
Credential expires at: 2026-05-11 12:28:55
Image registry path:   ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674

WARNING! Your credentials are stored unencrypted in '/home/moushuai.ms/.docker/config.json'.
Configure a credential helper to remove this warning. See
https://docs.docker.com/go/credential-store/

Login Succeeded

Note: Credentials will expire after the time above. You can run 'agentbay docker login' again to refresh.
Note: When tagging images, use: ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:<your-tag>
```

**注意事项：**

- 调用 `GetACRRepoCredential` 获取临时 ACR 凭证。
- 凭证有效期约 1 小时，到期后重新执行 `agentbay docker login` 即可刷新。
- 凭证会缓存到本地，供后续 `tag` 和 `push` 命令使用。

**登录后即可执行 docker 构建和推送**（在执行 `agentbay image create-from-template` 创建自定义镜像前必须完成）：

```bash
# 本地构建 docker 镜像（tag 必须以登录返回的 Image registry path 为前缀）
docker build \
  -t ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:<your-tag> \
  -f Dockerfile .

# 将本地 docker 镜像推送到 ACR 仓库
docker push ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:<your-tag>
```

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
