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

- 调用 `GetACRRepoCredential` 获取临时 ACR 凭证，**仅针对当前用户自己的仓库**。
- 凭证有效期约 1 小时，`docker build` 和 `docker push` 都依赖该登录状态，到期后重新执行 `agentbay docker login` 即可刷新。
- 凭证会缓存到本地，供后续 `tag` 和 `push` 命令使用。
- 命令同时返回该账号专属的**镜像上传地址**（即 `Image registry path`），构建和推送时必须以此地址为前缀。

**涉及接口：**

| Action                 | 所需权限                        |
| ---------------------- | ------------------------------- |
| `GetACRRepoCredential` | `agentbay:GetACRRepoCredential` |

```json
{
  "Action": ["agentbay:GetACRRepoCredential"]
}
```

> **Docker 环境准备**：`docker build` 和 `docker push` 需要本地 Docker Daemon 支持，CLI 本身不提供容器运行时。请根据操作系统安装 Docker 环境：
>
> - **macOS**：推荐安装 [OrbStack](https://orbstack.dev/)（轻量、快速，资源占用远低于 Docker Desktop）
> - **Windows**：推荐安装 Docker Desktop + WSL2 后端
> - **Linux**：直接使用系统包管理器安装 Docker Engine

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

| 参数             | 类型   | 必填 | 说明                 |
| ---------------- | ------ | ---- | -------------------- |
| `<source-image>` | string | 是   | 本地镜像名（含 tag） |
| `<target-tag>`   | string | 是   | 目标 tag             |

> 必须先执行 `agentbay docker login`。

> **注意**：`docker tag` 是本地 docker CLI 的封装命令，不调用任何 AgentBay OpenAPI 接口，无需配置额外的 RAM 权限。

---

### `docker push`

推送已 tag 的镜像到 AgentBay ACR。镜像名必须匹配 `$RegistryUrl/$Namespace/$RepoName[:tag]`，否则将被拒绝。

```bash
agentbay docker push <registry>/<namespace>/<repo>:v1.0
```

**参数：**

| 参数      | 类型   | 必填 | 说明              |
| --------- | ------ | ---- | ----------------- |
| `<image>` | string | 是   | 完整的 ACR 镜像名 |

> 必须先执行 `agentbay docker login`。

> **注意**：`docker push` 是本地 docker CLI 的封装命令，不调用任何 AgentBay OpenAPI 接口，无需配置额外的 RAM 权限。

---

### `docker share`

将当前用户的 Docker 镜像仓库（整体）授权给指定阿里云账号只读拉取。

```bash
agentbay docker share <TARGET_ALI_UID>
agentbay docker share --target-uid <TARGET_ALI_UID>
```

**参数：**

| 参数           | 类型  | 必填 | 说明                           |
| -------------- | ----- | ---- | ------------------------------ |
| `<target-uid>` | int64 | 是   | 目标阿里云账号 UID（位置参数） |
| `--target-uid` | int64 | 是   | 目标阿里云账号 UID（命名参数） |

**注意事项：**

- 被授权账号必须是**主账号**（RAM 子账号无法作为共享目标）。
- 被授权用户仅有 **pull** 权限，不可 push 或删除你的镜像。
- 授权**永久有效**，直到主动调用 `docker unshare` 撤销。

**输出示例：**

```
[INFO] Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Docker repo shared successfully.
  Target UID  : 1234567890
  Owner UID   : 9876543210
  Repo Name   : my-acr-repo
  Status      : active
```

**涉及接口：**

| Action            | 所需权限                   |
| ----------------- | -------------------------- |
| `ShareDockerRepo` | `agentbay:ShareDockerRepo` |

---

### `docker unshare`

取消将 Docker 镜像仓库共享给指定阿里云账号。

```bash
agentbay docker unshare <TARGET_ALI_UID>
agentbay docker unshare --target-uid <TARGET_ALI_UID>
```

**参数：**

| 参数           | 类型  | 必填 | 说明                                       |
| -------------- | ----- | ---- | ------------------------------------------ |
| `<target-uid>` | int64 | 是   | 要取消共享的目标阿里云账号 UID（位置参数） |
| `--target-uid` | int64 | 是   | 要取消共享的目标阿里云账号 UID（命名参数） |

**输出示例：**

```
[INFO] Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Docker repo unshared successfully.
```

**涉及接口：**

| Action              | 所需权限                     |
| ------------------- | ---------------------------- |
| `UnshareDockerRepo` | `agentbay:UnshareDockerRepo` |

---

### `docker list-shares`

列出 Docker 镜像仓库的共享信息。使用 `--direction` 指定方向：

- `Outgoing`：你共享给其他账号的仓库
- `Incoming`：其他账号共享给你的仓库

```bash
agentbay docker list-shares --direction Outgoing
agentbay docker list-shares --direction Incoming
```

**参数：**

| 参数              | 类型   | 必填 | 说明                                   |
| ----------------- | ------ | ---- | -------------------------------------- |
| `--direction`     | string | 是   | 共享方向：`Outgoing` 或 `Incoming`     |
| `--output` / `-o` | string | 否   | 输出格式，填 `json` 可获得机器可读输出 |

**输出示例（默认表格）：**

Outgoing（我共享给其他账号）：

```
[INFO] ListSharedDockerRepos Request ID: 89469103-92EF-12BD-BD9B-1F1B9A2F9D6D
PeerAliUid            Status
--------------------  ---------------
****7069              ACTIVE

Total: 1
```

Incoming（其他账号共享给我）：

```
[INFO] ListSharedDockerRepos Request ID: A4C0FF35-AA8A-1BDA-B807-3FA595048431
PeerAliUid            Status
--------------------  ---------------
****4214              ACTIVE

Total: 1
```

**输出示例（`--output json`）：**

```json
{
  "totalCount": 2,
  "items": [
    {
      "peerAliUid": 1234567890,
      "status": "active"
    },
    {
      "peerAliUid": 2345678901,
      "status": "pending"
    }
  ]
}
```

**涉及接口：**

| Action                  | 所需权限                         |
| ----------------------- | -------------------------------- |
| `ListSharedDockerRepos` | `agentbay:ListSharedDockerRepos` |
