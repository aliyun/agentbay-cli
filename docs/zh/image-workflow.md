[English](../en/image-workflow.md) | **中文**

# 镜像创建与共享

本文档分为两章：

- [一、镜像创建](#一镜像创建)：任何账号都可以独立完成的流程 —— 基于 Dockerfile 模板构建并注册为自定义镜像。
- [二、镜像共享](#二镜像共享)：把自己的镜像仓库共享给其他阿里云账号使用（可选）。

---

## 一、镜像创建

**场景**：单账号基于 Dockerfile 模板构建镜像，推送至 ACR，并注册为可激活的自定义镜像。

> 本章示例使用 UID `****4214` 的账号。如果你只是想为自己创建镜像，到本章末尾即可结束。

### Step 1：下载 Dockerfile 模板

```bash
agentbay image init --sourceImageId aio-ubuntu-2404
```

### Step 2：登录 Docker

获取临时凭证并自动登录本地 docker（有效期约 1 小时），同时返回镜像上传地址。

```bash
agentbay docker login
```

输出示例：

```
Credential expires at: 2026-05-11 12:28:55
Image registry path:   ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214

Login Succeeded
```

> `docker login` 只返回用户自己仓库的临时凭证，自动登陆本地 docker，有效期 1 个小时，因为 `docker build` 和 `docker push` 依赖登陆状态。同时会返回镜像上传地址。

### Step 3：本地构建镜像

确保本地已安装 Docker 环境，然后执行：

- **macOS**：推荐安装 [OrbStack](https://orbstack.dev/)（轻量、快速，资源占用远低于 Docker Desktop）
- **Windows**：推荐安装 Docker Desktop + WSL2 后端
- **Linux**：直接使用系统包管理器安装 Docker Engine

```bash
docker build \
  -t ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1 \
  -f Dockerfile .
```

### Step 4：推送镜像到 ACR

```bash
docker push ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1
```

### Step 5：创建自定义镜像

推荐使用短路径 `/namespace/repo:tag`（与 `image list` 返回的 `physicalImage` 字段格式一致）：

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****4214:cli-test-0.0.1 \
  --name cli-template-create-1 \
  --imageId aio-ubuntu-2404
```

输出示例：

```
[IMAGE] Creating custom image from template...
  SourceImage:      ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1
  PhysicalImageId:  /customer_cli/****4214:cli-test-0.0.1
  Name:             cli-template-create-1
  ImageId:          aio-ubuntu-2404
Requesting CreateImageFromTemplate... Done. (HTTP 200)

[RESPONSE]
  RequestId:      6921CD80-B902-1294-9C4A-6078DBBA7B2F
  Code:           ok
  Message:
  Success:        true
  HttpStatusCode: 200

[DATA]
  ImageId: imgc-0a9mg1hbjw1b7r564

[SUCCESS] CreateImageFromTemplate call completed.
```

### 关键说明

1. **`--source-image` 格式**：推荐使用短路径 `/namespace/repo:tag`（与 `image list` 输出格式一致）；也支持完整 registry 路径。
2. **物理镜像 ID 获取**：`image create-from-template` 成功后返回的 `PhysicalImageId` 即为后续可使用的短路径；也可通过 `image list` 查看已有镜像的 `physicalImage` 字段。

---

## 二、镜像共享

**场景**：A 账号（UID: `****4214`）把自己的镜像仓库（整体）共享给 B 账号（UID: `****7069`），B 账号基于其中的镜像创建自己的自定义镜像。

> 前置条件：A 账号已完成 [一、镜像创建](#一镜像创建) 中的全部步骤，仓库中至少有一个镜像。

### A 账号（共享方）

#### Step 1：共享镜像仓库给 B 账号

> **注意**：被授权账号必须是**主账号**（RAM 子账号无法作为共享目标）。

将当前用户的 Docker 镜像仓库（整体）授权给指定用户只读拉取。被授权用户仅有 pull 权限，不可 push 或删除 A 的镜像。授权永久有效，直到主动调用 `agentbay docker unshare` 撤销。

```bash
agentbay docker share --target-uid ****7069
```

#### Step 2：确认共享状态

```bash
agentbay docker list-shares --direction Outgoing
```

输出示例：

```
[INFO] ListSharedDockerRepos Request ID: 89469103-92EF-12BD-BD9B-1F1B9A2F9D6D
PeerAliUid            Status
--------------------  ---------------
****7069              ACTIVE

Total: 1
```

#### Step 3（可选）：撤销共享

如果后续不再希望 B 账号继续拉取自己的镜像，可随时撤销授权。`target-uid` 既可作为位置参数传入，也可使用 `--target-uid` 标志：

```bash
agentbay docker unshare ****7069
# 或
agentbay docker unshare --target-uid ****7069
```

输出示例：

```
[STEP 1/1] Cancelling Docker repo sharing with UID ****7069...
[INFO] UnshareDockerRepo Request ID: 7F2A1B3C-4D5E-6F70-8192-A3B4C5D6E7F8

[SUCCESS] Docker repo sharing cancelled.
  Revoked : true
```

撤销后可再次执行 `agentbay docker list-shares --direction Outgoing` 确认该条目已不在列表中。

### B 账号（接收方）

#### Step 1：查看接收到的共享仓库

```bash
agentbay docker list-shares --direction Incoming
```

输出示例：

```
[INFO] ListSharedDockerRepos Request ID: A4C0FF35-AA8A-1BDA-B807-3FA595048431
PeerAliUid            Status
--------------------  ---------------
****4214              ACTIVE

Total: 1
```

#### Step 2：基于 A 账号的镜像创建自定义镜像

通过 A 账号 `image create-from-template` 成功之后返回的 `PhysicalImageId`，或通过 `image list` 查看到的物理镜像 ID，参考 [一、镜像创建 → Step 5](#step-5创建自定义镜像) 完成创建：

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****4214:cli-test-0.0.1 \
  --name cli-test \
  --imageId aio-ubuntu-2404
```

### 关键说明

1. **权限范围**：被共享方仅有 **pull** 权限，不可 push 或删除 A 账号的镜像。
2. **授权有效期**：共享授权**永久有效**，直到 A 账号主动调用 `agentbay docker unshare` 撤销。
3. **主账号限制**：被授权账号必须是阿里云**主账号**，RAM 子账号无法作为共享目标。
