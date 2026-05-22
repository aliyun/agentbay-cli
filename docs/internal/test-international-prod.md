# AgentBay CLI 国际站生产环境测试文档

本文档描述在国际站生产环境（International Production）下，从登录到构建、激活镜像的完整手动测试步骤，用于验证 CLI 在国际站 endpoint 与 OAuth 下的行为。

---

## 1. 前置条件

- 已安装 AgentBay CLI（`agentbay` 可用）
- 拥有**阿里云国际站**（alibabacloud.com）账号，并已完成 RAM/OAuth 应用配置（国际站 OAuth 应用 ID：4192690673476752832）
- 网络可访问 signin.alibabacloud.com、oauth.alibabacloud.com、xiaoying.ap-southeast-1.aliyuncs.com
- 当前目录可写（用于下载 Dockerfile 和构建）

---

## 2. 环境准备与校验

### 2.1 设置国际站生产环境

```bash
export AGENTBAY_ENV=international
```

可选：若希望仅当前 shell 生效，无需写入 profile。

### 2.2 确认环境与 Endpoint

```bash
agentbay version
```

**预期输出应包含：**

- `Environment: international`
- `Endpoint: xiaoying.ap-southeast-1.aliyuncs.com`

若为其他 Environment 或 Endpoint，请检查是否已执行 `export AGENTBAY_ENV=international` 且未设置 `AGENTBAY_CLI_ENDPOINT` 覆盖。

---

## 3. 登录（Login）

### 3.1 执行登录

```bash
agentbay login
```

### 3.2 预期行为

1. 终端输出类似：
   - `Starting AgentBay authentication...`
   - `Trying to start callback server on port 3001... Success!`
   - `Opening browser for authentication...`
2. 浏览器自动打开，**地址栏应为**：`https://signin.alibabacloud.com/oauth2/v1/auth?...`（即国际站登录页，而非 signin.aliyun.com）
3. 使用国际站账号完成授权后，浏览器显示 “Authentication Successful” 或类似成功页
4. 终端输出：
   - `Authentication successful!`
   - `You are now logged in to AgentBay!`

### 3.3 校验点

- [ ] 浏览器打开的是 **signin.alibabacloud.com**
- [ ] URL 中 `client_id=4192690673476752832`（国际站默认 Client ID）
- [ ] 登录成功后无 `invalid clientId` 等报错

---

## 4. 查看镜像列表（List Images）

### 4.1 查看用户镜像

```bash
agentbay image list
```

**预期：** 返回当前账号在国际站生产环境下的用户镜像列表（可能为空或已有镜像）。表头应包含 IMAGE ID、IMAGE NAME、TYPE、STATUS、OS、APPLY SCENE。

### 4.2 查看系统镜像（获取基础镜像 ID）

```bash
agentbay image list --system-only
```

**预期：** 列出系统镜像，其中应包含可用于 CodeSpace 的基础镜像，例如 `code-space-debian-12`。记录后续步骤将用到的 **Image ID**（如 `code-space-debian-12`）。

### 4.3 校验点

- [ ] 请求成功，无 `Not authenticated` 或 `invalid_client` 错误
- [ ] 能至少看到系统镜像（如 code-space-debian-12）

---

## 5. 下载 Dockerfile 模板（Init）

### 5.1 创建工作目录（可选）

```bash
mkdir -p ~/agentbay-intl-test && cd ~/agentbay-intl-test
```

### 5.2 下载模板

```bash
agentbay image init --sourceImageId code-space-debian-12
```

或简短形式：

```bash
agentbay image init -i code-space-debian-12
```

### 5.3 预期输出

- `[INIT] Downloading Dockerfile template...`
- `Requesting Dockerfile template... Done.`
- `Downloading Dockerfile from OSS... Done.`
- `Writing Dockerfile to .../Dockerfile...`
- `[SUCCESS] Dockerfile template downloaded successfully!`
- 提示前 N 行为系统定义，仅可修改 N 行之后内容

### 5.4 校验点

- [ ] 当前目录下生成 `Dockerfile` 文件
- [ ] 无认证或 endpoint 相关报错

---

## 6. 创建镜像（Create Image）

### 6.1 执行构建

```bash
agentbay image create my-intl-test-image --dockerfile ./Dockerfile --imageId code-space-debian-12
```

或使用短选项：

```bash
agentbay image create my-intl-test-image -f ./Dockerfile -i code-space-debian-12
```

### 6.2 预期输出（按步骤）

- `[BUILD] Creating image 'my-intl-test-image'...`
- `[STEP 1/4] Getting upload credentials... Done.`
- `[STEP 2/4] Uploading Dockerfile... Done.`
- `[STEP 3/4] Uploading ADD/COPY files...`（若 Dockerfile 中有 COPY/ADD）
- `[STEP 4/4] Creating Docker image task... Done.`
- `[STEP 4/4] Building image (Task ID: task-xxxxx)...`
- 构建过程中可能轮询状态：`[STATUS] Build status: RUNNING`
- 最终：`[SUCCESS] Image created successfully!`
- `[RESULT] Image ID: imgc-xxxxxxxxxxxxxxxx`

### 6.3 校验点

- [ ] 成功拿到 **Image ID**（形如 `imgc-xxxxxxxxxxxxxxxx`），并记录该 ID 用于激活
- [ ] 未出现认证失败、endpoint 不可达或 OSS 上传失败

---

## 7. 激活镜像（Activate Image）

将上一步得到的 `imgc-xxxxxxxxxxxxxxxx` 替换为实际 Image ID。

### 7.1 使用默认规格激活（2c4g）

```bash
agentbay image activate imgc-xxxxxxxxxxxxxxxx
```

### 7.2 或指定规格

```bash
agentbay image activate imgc-xxxxxxxxxxxxxxxx --cpu 2 --memory 4
# 或 4c8g
agentbay image activate imgc-xxxxxxxxxxxxxxxx --cpu 4 --memory 8
```

### 7.3 预期输出

- `[ACTIVATE] Activating image...`
- `Checking current image status... Done.`
- `Creating resource group... Done.`
- `Waiting for activation to complete...`
- 轮询状态：`Status: Activating (elapsed: Xs, attempt: Y/60)`
- `[SUCCESS] Image activated successfully!`

若镜像已处于激活状态，可能提示 “No action needed.”

### 7.4 校验点

- [ ] 激活流程结束无报错
- [ ] 可通过 `agentbay image list` 再次查看该镜像，STATUS 为 **Activated**

---

## 8. 可选：查看列表与停用

### 8.1 再次查看镜像列表

```bash
agentbay image list
```

确认刚创建并激活的镜像在列表中且状态为 **Activated**。

### 8.2 停用镜像（释放资源）

```bash
agentbay image deactivate imgc-xxxxxxxxxxxxxxxx
```

预期：`[SUCCESS] Image deactivated successfully!`

### 8.3 登出

```bash
agentbay logout
```

预期：`Successfully logged out from AgentBay`

---

## 9. 测试通过标准

| 步骤           | 通过标准                                                                 |
|----------------|--------------------------------------------------------------------------|
| 环境与版本     | `agentbay version` 显示 Environment: international，Endpoint: xiaoying.ap-southeast-1.aliyuncs.com |
| 登录           | 浏览器为 signin.alibabacloud.com，登录成功且无 invalid clientId         |
| 镜像列表       | `image list` 与 `image list --system-only` 正常返回                      |
| 下载模板       | `image init -i code-space-debian-12` 成功生成 Dockerfile                 |
| 创建镜像       | `image create` 完成并返回 Image ID                                       |
| 激活镜像       | `image activate` 成功，列表中该镜像状态为 Activated                      |

---

## 10. 常见问题

- **浏览器打开的是 signin.aliyun.com 而不是 signin.alibabacloud.com**  
  确认已执行 `export AGENTBAY_ENV=international`，且未设置 `AGENTBAY_OAUTH_REGION=domestic` 等覆盖。

- **报错 invalid clientId**  
  国际站需使用国际站 OAuth 应用；当前默认 Client ID 为 4192690673476752832。若使用自定义应用，请设置 `AGENTBAY_OAUTH_CLIENT_ID`。

- **请求超时或连接失败**  
  检查本机网络是否能访问 xiaoying.ap-southeast-1.aliyuncs.com（国际站生产 endpoint）。

- **想用国际站预发**  
  使用 `export AGENTBAY_ENV=international-pre`，endpoint 将切换为 xiaoying-pre.ap-southeast-1.aliyuncs.com（预发配置以实际环境为准）。
