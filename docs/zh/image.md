[English](../en/image.md) | **中文**

# 镜像管理 — `agentbay image`

管理 AgentBay 镜像：创建、列出、激活、停用、删除、查询生命周期状态。

> 当前 CLI 版本仅支持创建和激活 **CodeSpace** 类型的镜像。

## 命令

### `image list`

列出可用的 AgentBay 镜像。

```bash
agentbay image list                      # 用户镜像（默认）
agentbay image list --include-system     # 用户镜像 + 系统镜像
agentbay image list --system-only        # 仅系统镜像
agentbay image list --os-type Linux      # 按 OS 过滤：Linux / Android / Windows
agentbay image list --page 2 --size 5    # 分页
agentbay image list --output json        # JSON 输出（AI/脚本使用）
```

**参数：**

| 参数               | 短参数 | 类型   | 必填 | 说明                                                             |
| ------------------ | ------ | ------ | ---- | ---------------------------------------------------------------- |
| `--os-type`        | `-o`   | string | 否   | 按 OS 过滤（Linux、Windows、Android）                            |
| `--include-system` |        |        | 否   | 在用户镜像基础上包含系统镜像                                     |
| `--system-only`    |        |        | 否   | 仅显示系统镜像                                                   |
| `--page`           | `-p`   | int    | 否   | 页码（默认：1）                                                  |
| `--size`           | `-s`   | int    | 否   | 每页条数（默认：10）                                             |
| `--output`         |        | string | 否   | 输出格式。使用 `json` 获取机器可读的完整数据（适合 AI/脚本使用） |

**输出示例：**

默认表格输出：

```
=== USER IMAGES (17) ===
IMAGE ID              IMAGE NAME       TYPE                 STATUS        OS                 PHYSICAL IMAGE                 APPLY SCENE
--------              ----------       ----                 ------        --                 --------------                 -----------
imgc-xxxxx...xxx      my-app           DockerBuilder        Available     Android 14         registry.example.com/img:tag   CodeSpace

=== SYSTEM IMAGES (3) ===
IMAGE ID                  IMAGE NAME                     TYPE                 STATUS        OS                 PHYSICAL IMAGE                 APPLY SCENE
--------                  ----------                     ----                 ------        --                 --------------                 -----------
mobile-use-android-14     Mobile Use Android 14          DedicatedDesktop     Available     Android 14                                        MobileUse
```

使用 `--output json` 输出完整 JSON（注：`image list` 的 `-o` 短参数已被 `--os-type` 占用，请使用全称）：

```bash
agentbay image list --output json
```

```json
{
  "totalCount": 2,
  "images": [
    {
      "imageId": "imgc-xxxxxxxxxxxxxx",
      "imageName": "my-app",
      "type": "DockerBuilder",
      "status": "IMAGE_AVAILABLE",
      "statusDisplay": "Available",
      "osName": "Linux",
      "osVersion": "Debian 12",
      "osDisplay": "Linux Debian 12",
      "physicalImage": "registry.example.com/my-app:latest",
      "applyScene": "CodeSpace"
    }
  ]
}
```

**注意事项：**

- 系统镜像始终可用，无需激活；只有用户镜像必须先激活才能使用。
- **状态含义**：Creating（构建中）、Available（可激活）、Activated（已激活）、Create Failed（构建失败）
- **类型含义**：DockerBuilder（用户创建）、DedicatedDesktop（系统镜像）

**涉及接口：**

| Action          | 所需权限                 |
| --------------- | ------------------------ |
| `ListMcpImages` | `agentbay:ListMcpImages` |

```json
{
  "Action": ["agentbay:ListMcpImages"]
}
```

---

### `image init`

从云端下载 Dockerfile 模板到当前目录。

```bash
agentbay image init --sourceImageId code-space-debian-12
agentbay image init -i code-space-debian-12
```

**参数：**

| 参数              | 短参数 | 类型   | 必填 | 说明        |
| ----------------- | ------ | ------ | ---- | ----------- |
| `--sourceImageId` | `-i`   | string | 是   | 系统镜像 ID |

生产环境可用的 `sourceImageId`：

- `code-space-debian-12`
- `code-space-debian-12-enhanced`
- `aio-ubuntu-2404`

**输出：**

```
[INIT] Downloading Dockerfile template...
Requesting Dockerfile template... Done.
Downloading Dockerfile from OSS... Done.
Writing Dockerfile to /path/to/current/directory/Dockerfile... Done.
[SUCCESS] Dockerfile template downloaded successfully!
[IMPORTANT] The first 5 line(s) of the Dockerfile are system-defined and cannot be modified.
[IMPORTANT] Please only modify content after line 5.
```

**注意事项：**

- 如果当前目录已有 `Dockerfile`，将被覆盖。
- Dockerfile 的前 N 行由系统定义，不可修改，仅可编辑第 N+1 行之后的内容。

---

### `image create`（已废弃，请改用 `create-from-template`）

> 警告：`image create` **已不推荐使用，后续版本将被移除**。如需创建自定义镜像，请改用 [`image create-from-template`](#image-create-from-template)。

基于 Dockerfile 构建自定义镜像。`COPY` / `ADD` 引用的文件会被自动解析并上传。

```bash
agentbay image create myapp --dockerfile ./Dockerfile --imageId code-space-debian-12
agentbay image create myapp -f ./Dockerfile -i code-space-debian-12
```

---

### `image create-from-template`

基于系统镜像模板 + 已推送到 AgentBay ACR 的 Docker 物理镜像创建自定义镜像（调用 `CreateImageFromTemplate` 接口）。源镜像可以来自当前账号自己的仓库，也可以来自其他账号共享给当前账号的 Docker 仓库。

> **前置依赖**：执行该命令前，需要先准备好可用的 Docker 物理镜像，并已推送到 AgentBay ACR 仓库。完整流程：
>
> 1. `agentbay image init -i <system-image-id>` —— 下载基础 Dockerfile 模板，根据需要在模板可编辑区域修改。
> 2. `agentbay docker login` —— 自动登录本地 docker 并获取镜像上传地址（有效期约 1 小时）。
> 3. 本地构建 Docker 镜像 —— 执行 `docker build` 前需要确保本地已安装 Docker 环境：
>    - **macOS**：推荐安装 [OrbStack](https://orbstack.dev/)（轻量、快速，资源占用远低于 Docker Desktop）
>    - **Windows**：推荐安装 Docker Desktop + WSL2 后端
>    - **Linux**：直接使用系统包管理器安装 Docker Engine
>
>    然后执行 `docker build -t <registry-path>:<your-tag> -f Dockerfile .`
>
> 4. `docker push <registry-path>:<your-tag>` —— 推送到 ACR。
> 5. `agentbay image create-from-template ...` —— 基于上述镜像创建自定义镜像（即本命令）。

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****9674:<your-tag> \
  --name imageName \
  --imageId code-space-debian-12

# 短参数形式
agentbay image create-from-template -s /customer_cli/****9674:<your-tag> -n imageName -i code-space-debian-12
```

**参数：**

| 参数             | 短参数 | 类型   | 必填 | 说明                                            |
| ---------------- | ------ | ------ | ---- | ----------------------------------------------- |
| `--source-image` | `-s`   | string | 是   | 已推送到 ACR 的源 Docker 镜像仓库路径（含 tag） |
| `--name`         | `-n`   | string | 是   | 自定义镜像名称                                  |
| `--imageId`      | `-i`   | string | 是   | 基础系统镜像 ID（如 `code-space-debian-12`）    |

`--source-image` 支持两种格式：

1. **推荐**：短路径 `/customer_cli/<aliuid>:tag`（与 `image list` 返回的 `physicalImage` 字段格式一致，可直接复制使用）
2. **也支持**：完整 registry 路径 `<registry>/customer_cli/<aliuid>:tag`

CLI 会先解析 `source-image` 中的 AliUID：

- 如果该 AliUID 与当前本地 `agentbay docker login` 缓存的 ACR 仓库匹配，则按自有仓库处理。
- 如果不匹配，则自动调用 `ListSharedDockerRepos` 查询当前账号是否收到该 AliUID 的 Docker 仓库共享授权；返回数据非空才会继续创建。
- 自有短路径会在终端输出中补全为当前账号 registry URL；共享短路径会保持短路径展示，避免误导为当前账号 ACR 地址。
- 命令会输出涉及的 OpenAPI Request ID；共享仓库场景会先输出 `ListSharedDockerRepos` 的 Request ID，再输出 `CreateImageFromTemplate` 的 Request ID。

**创建流程（服务端）：**

1. 依据传入的 `--source-image` 锁定 Docker 物理镜像。
2. 参考传入的 `--imageId`（系统镜像 ID）对应的配置参数。
3. 创建一个 Docker 自定义镜像。

**对传入的 Docker 镜像注意事项：**

1. **不要包含 `CMD` 指令**。
2. **不要修改 `FROM` 指令**（保持 `agentbay image init` 模板中的 `FROM`）。
3. Dockerfile 末尾**必须**包含 `USER root`，或者全程未指定 `USER`。
4. 目前**仅支持 Linux x86 架构**。
5. 如需指定 EntryPoint，请按以下规则改写为 supervisor 配置：

   ```dockerfile
   RUN echo '[program:user-define]' > /etc/supervisor/conf.d/user-define.conf && \
       echo 'command=%s' >> /etc/supervisor/conf.d/user-define.conf && \
       echo 'priority=33' >> /etc/supervisor/conf.d/user-define.conf
   ```

   将 `%s` 替换为实际要执行的命令。

**涉及接口：**

| Action                    | 所需权限                           |
| ------------------------- | ---------------------------------- |
| `ListSharedDockerRepos`   | `agentbay:ListSharedDockerRepos`   |
| `CreateImageFromTemplate` | `agentbay:CreateImageFromTemplate` |

```json
{
  "Action": [
    "agentbay:ListSharedDockerRepos",
    "agentbay:CreateImageFromTemplate"
  ]
}
```

---

### `image activate`

激活用户镜像使其可用。

```bash
# 使用默认资源规格（2c4g）
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

**参数：**

| 参数                       | 短参数 | 类型   | 必填 | 说明                                          |
| -------------------------- | ------ | ------ | ---- | --------------------------------------------- |
| `--cpu`                    | `-c`   | int    | 否   | CPU 核数（2、4、8），须与 `--memory` 同时指定 |
| `--memory`                 | `-m`   | int    | 否   | 内存 GB（4、8、16），须与 `--cpu` 同时指定    |
| `--network-type`           |        | string | 否   | 网络类型：`DEFAULT` 或 `ADVANCED`             |
| `--session-bandwidth`      |        | int    | 否   | 单 session 最高公网带宽，单位 Mbps，建议设置范围 2-200；仅 ADVANCED 网络可用，不传表示不限制单 session 公网访问带宽上限 |
| `--dns-address`            |        | string | 否   | DNS 地址；仅 ADVANCED 网络可用，可重复指定；不传则 CLI 自动使用当前 office network 的默认 DNS |
| `--lifecycle-mode`         |        | string | 否   | 释放模式：`auto`（自动释放）或 `manual`（手动释放） |
| `--lifecycle-max-runtime`  |        | int    | 否   | 单次运行最长时长（分钟）；需 `--lifecycle-mode` 为 `auto` |
| `--lifecycle-hibernate`    |        | int    | 否   | 休眠最大时长（小时）；需 `--lifecycle-mode` 为 `auto`   |
| `--lifecycle-idle-timeout` |        | int    | 否   | 无活动最大时长（分钟）；需 `--lifecycle-mode` 为 `auto` |
| `--region-id`              |        | string | 否   | 资源部署的区域 ID                             |

**支持的资源规格：** `2c4g`（默认）、`4c8g`、`8c16g`

**注意事项：**

- `--cpu` 和 `--memory` 必须同时指定。
- `--session-bandwidth` 与 `--dns-address` 仅在 `--network-type ADVANCED` 时可用，且均为可选；DEFAULT 网络下传这两个参数会被拒绝。不传 `--dns-address` 时，CLI 会自动从当前 office network 拉取默认 DNS 填充。
- 激活通常需要 1-2 分钟。如果已激活，会提示 "No action needed."

**输出：**

```
[ACTIVATE] Activating image...
Checking current image status... Done.
Creating resource group... Done.
Waiting for activation to complete...
  Status: Activating (elapsed: 5s, attempt: 2/60)
[SUCCESS] Image activated successfully!
```

**涉及接口：**

| Action                  | 所需权限                         |
| ----------------------- | -------------------------------- |
| `GetMcpImageInfo`       | `agentbay:GetMcpImageInfo`       |
| `DescribeInstanceTypes` | `agentbay:DescribeInstanceTypes` |
| `DescribeMcpPolicyData` | `agentbay:DescribeMcpPolicyData` |
| `CreateMcpPolicyData`   | `agentbay:CreateMcpPolicyData`   |
| `ModifyMcpPolicyData`   | `agentbay:ModifyMcpPolicyData`   |
| `DescribeOfficeSites`   | `agentbay:DescribeOfficeSites`   |
| `SaveMcpPolicyData`     | `agentbay:SaveMcpPolicyData`     |
| `CreateResourceGroup`   | `agentbay:CreateResourceGroup`   |

```json
{
  "Action": [
    "agentbay:GetMcpImageInfo",
    "agentbay:DescribeInstanceTypes",
    "agentbay:DescribeMcpPolicyData",
    "agentbay:CreateMcpPolicyData",
    "agentbay:ModifyMcpPolicyData",
    "agentbay:DescribeOfficeSites",
    "agentbay:SaveMcpPolicyData",
    "agentbay:CreateResourceGroup"
  ]
}
```

---

### `image deactivate`

停用已激活的用户镜像。

```bash
agentbay image deactivate imgc-xxxxxxxxxxxxxx
```

**输出：**

```
[DEACTIVATE] Deactivating image...
Deleting resource group... Done.
Waiting for deactivation to complete...
  Status: Deactivating (elapsed: 5s, attempt: 2/40)
[SUCCESS] Image deactivated successfully!
```

通常几秒内完成。

**涉及接口：**

| Action                | 所需权限                       |
| --------------------- | ------------------------------ |
| `GetMcpImageInfo`     | `agentbay:GetMcpImageInfo`     |
| `ListMcpImages`       | `agentbay:ListMcpImages`       |
| `DeleteResourceGroup` | `agentbay:DeleteResourceGroup` |

```json
{
  "Action": [
    "agentbay:GetMcpImageInfo",
    "agentbay:ListMcpImages",
    "agentbay:DeleteResourceGroup"
  ]
}
```

---

### `image delete`

**永久删除**用户镜像。仅已停用的用户镜像可删除。

```bash
agentbay image delete imgc-xxxxxxxxxxxxxx          # 带确认提示
agentbay image delete imgc-xxxxxxxxxxxxxx --yes    # 跳过确认（脚本 / CI）
```

**参数：**

| 参数    | 短参数 | 类型 | 必填 | 说明                           |
| ------- | ------ | ---- | ---- | ------------------------------ |
| `--yes` | `-y`   |      | 否   | 跳过确认提示（非交互模式必填） |

**限制：**

- 仅用户镜像可删除（系统镜像不可删除）
- 以下状态的镜像不可删除：`IMAGE_CREATING`、`RESOURCE_DEPLOYING`、`RESOURCE_DELETING`、`RESOURCE_PUBLISHED`、`RESOURCE_FAILED`、`RESOURCE_MAINTAINING`

**输出：**

```
[DELETE] Deleting image 'imgc-xxxxx'...
Checking current image status... Done.
[INFO] GetMcpImageInfo Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
[INFO] Image Type: User
[INFO] Current Status: Available (Deactivated)
Are you sure you want to permanently delete image 'imgc-xxxxx'? This action is irreversible. [y/N]: y
Deleting image... Done.
[INFO] DeleteMcpImage Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
[SUCCESS] Image 'imgc-xxxxx' has been permanently deleted.
```

**涉及接口：**

| Action            | 所需权限                   |
| ----------------- | -------------------------- |
| `GetMcpImageInfo` | `agentbay:GetMcpImageInfo` |
| `DeleteMcpImage`  | `agentbay:DeleteMcpImage`  |

```json
{
  "Action": ["agentbay:GetMcpImageInfo", "agentbay:DeleteMcpImage"]
}
```

---

### `image status`

查询镜像的资源生命周期状态（与 `image create` 时的 Docker 构建任务状态不同）。

```bash
agentbay image status imgc-xxxxxxxxxxxxxx
```

**常见状态值：**

| 状态                  | 含义               |
| --------------------- | ------------------ |
| `IMAGE_CREATING`      | 镜像创建中         |
| `IMAGE_CREATE_FAILED` | 镜像创建失败       |
| `IMAGE_AVAILABLE`     | 可用，未激活       |
| `RESOURCE_DEPLOYING`  | 激活中             |
| `RESOURCE_PUBLISHED`  | 已激活，使用中     |
| `RESOURCE_DELETING`   | 停用中             |
| `RESOURCE_FAILED`     | 激活或资源操作失败 |
| `RESOURCE_CEASED`     | 资源已释放         |

**涉及接口：**

| Action            | 所需权限                   |
| ----------------- | -------------------------- |
| `GetMcpImageInfo` | `agentbay:GetMcpImageInfo` |

```json
{
  "Action": ["agentbay:GetMcpImageInfo"]
}
```

---

### `image set-max-session`

设置已激活用户镜像的最大并发会话数。要求镜像处于 `RESOURCE_PUBLISHED` 状态且使用**高级网络**。

```bash
agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 10
```

**参数：**

| 参数                | 类型   | 必填 | 说明           |
| ------------------- | ------ | ---- | -------------- |
| `--image-id`        | string | 是   | 镜像 ID        |
| `--max-session-num` | int    | 是   | 最大并发会话数 |

> 该命令会轮询直到资源组就绪（通常约 5 分钟）。

**涉及接口：**

| Action                                        | 所需权限                                               |
| --------------------------------------------- | ------------------------------------------------------ |
| `GetMcpImageInfo`                             | `agentbay:GetMcpImageInfo`                             |
| `BatchCreateHideResourceGroupsWithMaxSession` | `agentbay:BatchCreateHideResourceGroupsWithMaxSession` |

```json
{
  "Action": [
    "agentbay:GetMcpImageInfo",
    "agentbay:BatchCreateHideResourceGroupsWithMaxSession"
  ]
}
```

---

### `image warmup-status`

查询当前账户的预热状态，包括会话配额、镜像配额以及预热镜像详情。

```bash
agentbay image warmup-status
```

**输出包括：**

- **会话配额** —— 最大会话数限制、已使用的会话数、可用的会话数
- **镜像配额** —— 最大镜像数、当前镜像数
- **预热镜像** —— 镜像 ID、总最大容量、资源组数量

**涉及接口：**

| Action                     | 所需权限                            |
| -------------------------- | ----------------------------------- |
| `DescribeWarmUpStatusOpen` | `agentbay:DescribeWarmUpStatusOpen` |

```json
{
  "Action": ["agentbay:DescribeWarmUpStatusOpen"]
}
```
