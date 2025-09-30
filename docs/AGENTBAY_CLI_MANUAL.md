# AgentBay CLI 使用手册

## 简介

AgentBay CLI 是 AgentBay 服务的命令行接口工具，提供完整的镜像管理功能。通过 CLI 工具，用户可以创建、构建、激活和管理自定义镜像，用于 AgentBay 环境。

### 主要功能

- **认证管理**：基于 OAuth 的安全登录机制，支持阿里云账号集成
- **镜像列表**：浏览用户镜像和系统镜像，支持分页和过滤
- **镜像创建**：从 Dockerfile 构建自定义镜像，支持基础镜像
- **镜像激活**：激活用户镜像实例，支持自定义资源配置
- **镜像停用**：停用已激活的镜像实例
- **模板下载**：从云端下载 Dockerfile 模板
- **配置管理**：安全的令牌存储和自动令牌刷新

### 支持的镜像类型

当前版本的 CLI 工具仅支持创建和激活 **CodeSpace** 类型的镜像。

---


## 快速开始

### 1. 登录

首先需要登录到 AgentBay：

```bash
agentbay login
```

CLI 将自动打开浏览器进行阿里云认证。完成登录后返回终端。

### 2. 查看可用镜像

列出可用的镜像：

```bash
# 仅列出用户镜像（默认）
agentbay image list

# 包含系统镜像与用户镜像
agentbay image list --include-system

# 仅显示系统镜像
agentbay image list --system-only
```

### 3. 下载 Dockerfile 模板

下载 Dockerfile 模板到当前目录：

```bash
agentbay image init
```

### 4. 创建自定义镜像

使用 Dockerfile 创建自定义镜像：

```bash
agentbay image create myapp --dockerfile ./Dockerfile --imageId code-space-debian-12
```

### 5. 激活镜像

用户自定义镜像使用之前需要激活，系统镜像无需激活。激活用户镜像，使其可用于部署：

```bash
agentbay image activate imgc-xxxxx...xxx
```

### 6. 停用镜像

使用完毕后停用自定义镜像以节约资源。停用已激活的用户镜像，释放相关资源：

```bash
agentbay image deactivate imgc-xxxxx...xxx
```

---

## 命令参考

### 全局选项

所有命令都支持以下全局选项：

- `--help, -h`: 显示命令帮助信息
- `--verbose, -v`: 启用详细输出模式，显示调试信息
- `--version`: 显示 CLI 版本信息

### 命令结构

```
agentbay [全局选项] <命令> [命令选项] [参数]
```

---

## 镜像管理

### 镜像激活与停用说明

- **用户自定义镜像**：使用之前需要激活，激活后可用于部署
- **系统镜像**：始终可用，无需激活
- **停用镜像**：使用完毕后停用自定义镜像以节约资源，停用已激活的用户镜像会释放相关资源

---

### image list - 列出镜像

列出可用的 AgentBay 镜像。

**语法：**

```bash
agentbay image list [选项]
```

**选项：**

- `--os-type, -o <类型>`: 按操作系统类型过滤（Linux、Android、Windows）
- `--include-system`: 同时显示用户镜像和系统镜像
- `--system-only`: 仅显示系统镜像
- `--page, -p <数字>`: 页码（默认：1）
- `--size, -s <数字>`: 每页显示数量（默认：10）

**示例：**

```bash
# 列出用户镜像
agentbay image list

# 列出 Linux 镜像
agentbay image list --os-type Linux

# 列出所有镜像（用户 + 系统）
agentbay image list --include-system

# 仅列出系统镜像
agentbay image list --system-only

# 分页查询
agentbay image list --page 2 --size 5
```

**输出说明：**

命令输出包含以下列：

- **IMAGE ID**: 镜像唯一标识符
- **IMAGE NAME**: 镜像名称
- **TYPE**: 镜像类型（DockerBuilder 或 DedicatedDesktop）
- **STATUS**: 镜像状态
- **OS**: 操作系统类型和版本
- **APPLY SCENE**: 应用场景

**状态说明：**

- **Creating**: Building
- **Available**: Ready to activate
- **Activated**: Running
- **Create Failed**: Build failed

**注意事项：**

- 系统镜像始终可用，无需激活
- 只有用户创建的镜像需要激活后才能使用
- 默认情况下仅显示用户镜像
- 镜像列表会按用户镜像和系统镜像分组显示
- **重要**：`image list` 命令只显示镜像的构建状态（Creating、Available、Create Failed），不显示激活/停用状态。激活和停用状态只在执行 `image activate` 或 `image deactivate` 命令时显示

---

### image init - 下载 Dockerfile 模板

从云端下载 Dockerfile 模板到当前目录。

**语法：**

```bash
agentbay image init
```

**示例：**

```bash
# 下载 Dockerfile 模板
agentbay image init
```

**输出示例：**

```
[INIT] Downloading Dockerfile template...
Requesting Dockerfile template... Done.
Downloading Dockerfile from OSS... Done.
Writing Dockerfile to /path/to/current/directory/Dockerfile...
[WARN] Dockerfile already exists at /path/to/current/directory/Dockerfile
[INFO] The existing file will be overwritten.
 Done.
[SUCCESS] Dockerfile template downloaded successfully!
[INFO] Dockerfile saved to: /path/to/current/directory/Dockerfile
[IMPORTANT] The first 5 line(s) of the Dockerfile are system-defined and cannot be modified.
[IMPORTANT] Please only modify content after line 5.
```

**注意事项：**

- 如果当前目录已存在 `Dockerfile`，命令会覆盖现有文件
- 覆盖前会显示警告信息
- 此步骤为可选，您也可以手动创建 Dockerfile 或使用现有文件
- **重要提示**：Dockerfile 模板的前 N 行（N 由系统返回）是系统定义的，不能修改
- 只能修改第 N+1 行之后的内容，否则可能导致镜像构建失败
- 系统会在下载成功后显示不可编辑的行数信息，请务必遵守此限制

---

### image create - 创建镜像

从 Dockerfile 创建新的 AgentBay 镜像。

**语法：**

```bash
agentbay image create <镜像名称> --dockerfile <路径> --imageId <基础镜像ID>
```

**参数：**

- `<镜像名称>`: 自定义镜像名称（必需）

**选项：**

- `--dockerfile, -f <路径>`: Dockerfile 文件路径（必需）
- `--imageId, -i <ID>`: 基础镜像 ID（必需）

**示例：**

```bash
# 使用完整选项名称
agentbay image create my-app --dockerfile ./Dockerfile --imageId code-space-debian-12

# 使用短选项名称
agentbay image create my-app -f ./Dockerfile -i code-space-debian-12

# 使用详细输出模式
agentbay image create my-app -f ./Dockerfile -i code-space-debian-12 -v
```

**输出示例：**

```
[BUILD] Creating image 'my-app'...
[STEP 1/4] Getting upload credentials... Done.
[STEP 2/4] Uploading Dockerfile... Done.
[STEP 3/4] Creating Docker image task... Done.
[STEP 4/4] Building image (Task ID: task-xxxxx)...
[STATUS] Build status: RUNNING
[SUCCESS] Image created successfully!
[RESULT] Image ID: imgc-xxxxx...xxx
```

**构建流程：**

1. 获取上传凭证
2. 上传 Dockerfile 到对象存储
3. 创建 Docker 镜像构建任务
4. 启动镜像构建

**注意事项：**

- 构建时间取决于镜像大小和复杂度
- 使用 `-v` 选项可查看详细的构建日志
- 构建过程中可以通过 `image list` 命令查看状态
- 基础镜像 ID 必须是有效的系统镜像 ID，可通过 `image list --system-only` 查看

---

### image activate - 激活镜像

激活用户镜像，使其可用于部署。

**语法：**

```bash
agentbay image activate <镜像ID> [选项]
```

**参数：**

- `<镜像ID>`: 要激活的镜像 ID（必需）

**选项：**

- `--cpu, -c <核心数>`: CPU 核心数（2、4 或 8）
- `--memory, -m <GB>`: 内存大小，单位 GB（4、8 或 16）

**支持的资源配置组合：**

- `2c4g`: 2 个 CPU 核心，4 GB 内存
- `4c8g`: 4 个 CPU 核心，8 GB 内存
- `8c16g`: 8 个 CPU 核心，16 GB 内存

**示例：**

```bash
# 使用默认资源配置激活
agentbay image activate imgc-xxxxx...xxx

# 使用 2c4g 配置激活
agentbay image activate imgc-xxxxx...xxx --cpu 2 --memory 4

# 使用 4c8g 配置激活
agentbay image activate imgc-xxxxx...xxx --cpu 4 --memory 8

# 使用 8c16g 配置激活
agentbay image activate imgc-xxxxx...xxx --cpu 8 --memory 16

# 使用详细输出
agentbay image activate imgc-xxxxx...xxx --cpu 4 --memory 8 -v
```

**输出示例：**

```
[ACTIVATE] Activating image...
Checking current image status... Done.
Creating resource group... Done.
Waiting for activation to complete...
  Status: Activating (elapsed: 5s, attempt: 2/60)
  Status: Activating (elapsed: 13s, attempt: 3/60)
[SUCCESS] Image activated successfully!
```

**注意事项：**

- 只有用户镜像可以激活
- 系统镜像始终可用，无需激活
- CPU 和内存选项必须同时指定，且必须匹配支持的组合
- 如果未指定 CPU 和内存，将使用默认资源配置
- 激活过程通常需要 1-2 分钟
- 如果镜像已经激活，命令会提示无需操作
- 激活过程中会轮询镜像状态，最多尝试 60 次，总超时时间为 30 分钟

---

### image deactivate - 停用镜像

停用已激活的用户镜像，释放相关资源。

**语法：**

```bash
agentbay image deactivate <镜像ID>
```

**参数：**

- `<镜像ID>`: 要停用的镜像 ID（必需）

**示例：**

```bash
# 停用镜像
agentbay image deactivate imgc-xxxxx...xxx

# 使用详细输出
agentbay image deactivate imgc-xxxxx...xxx -v
```

**输出示例：**

```
[DEACTIVATE] Deactivating image...
Deleting resource group... Done.
Waiting for deactivation to complete...
  Status: Deactivating (elapsed: 5s, attempt: 2/40)
[SUCCESS] Image deactivated successfully!
```



---

## 认证管理

### login - 登录

通过 OAuth 流程登录到 AgentBay。

**语法：**

```bash
agentbay login
```

**功能说明：**

1. 启动本地回调服务器（端口 3001）
2. 打开浏览器进行阿里云认证
3. 接收授权码并交换访问令牌
4. 保存认证令牌到本地配置文件

**输出示例：**

```
Starting AgentBay authentication...
Starting local callback server on port 3001...
Opening browser for authentication...
Browser opened successfully!
Waiting for callback on http://localhost:3001/callback...
Authentication successful!
Received authorization code: xxxxx...
Exchanging authorization code for access token...
Saving authentication tokens...
Authentication tokens saved successfully!
You are now logged in to AgentBay!
```

**注意事项：**

- 如果已经登录且令牌未过期，命令会提示已登录
- 如果端口 3001 被占用，命令会显示错误信息并提供排查建议
- 如果浏览器无法自动打开，命令会显示认证 URL，可手动复制到浏览器
- 认证超时时间为 5 分钟
- 令牌会自动刷新，无需频繁登录

**故障排查：**

如果端口被占用，可以使用以下命令检查：

- **macOS/Linux**: `lsof -i :3001`
- **Windows**: `netstat -ano | findstr :3001`

---

### logout - 登出

登出 AgentBay，清除本地认证数据。

**语法：**

```bash
agentbay logout
```

**功能说明：**

1. 尝试撤销服务器端的刷新令牌
2. 清除本地配置文件中的认证令牌

**输出示例：**

```
Logging out from AgentBay...
Revoking server tokens...
Refresh token revoked successfully
Clearing local authentication data...
Successfully logged out from AgentBay
```

**注意事项：**

- 即使服务器端撤销失败，本地数据仍会被清除
- 访问令牌是短期有效的，会自动过期
- 撤销刷新令牌会同时使相关的访问令牌失效

---

### version - 版本信息

显示 CLI 版本和构建信息。

**语法：**

```bash
agentbay version
```

**输出示例：**

```
AgentBay CLI version 1.0.0
Git commit: abc1234
Build date: 2025-01-15
Environment: production
Endpoint: xiaoying-share.cn-shanghai.aliyuncs.com
```

**信息说明：**

- **Version**: CLI 版本号
- **Git commit**: 构建时的 Git 提交哈希
- **Build date**: 构建日期
- **Environment**: 当前环境（production 或 prerelease）
- **Endpoint**: 当前使用的 API 端点

---

## 配置说明

### 配置文件结构

配置文件采用 JSON 格式，包含以下信息：

- 访问令牌（Access Token）
- 刷新令牌（Refresh Token）
- ID 令牌（ID Token）
- 令牌类型（Token Type）
- 令牌过期时间（Expires At）

### 令牌管理

CLI 工具会自动管理令牌：

- **自动刷新**：访问令牌过期前会自动使用刷新令牌获取新令牌
- **安全存储**：令牌存储在用户配置目录，仅当前用户可访问
- **令牌验证**：每次 API 调用前会检查令牌有效性

### 环境变量

CLI 工具支持通过环境变量进行配置：

- `AGENTBAY_ENV`: 设置运行环境（production 或 prerelease）

---

## 常见问题

### 认证相关

**Q: 登录时提示端口被占用怎么办？**

A: 端口 3001 可能被其他程序占用。可以：
1. 关闭占用端口的程序
2. 使用 `lsof -i :3001`（macOS/Linux）或 `netstat -ano | findstr :3001`（Windows）查找占用进程
3. 终止占用进程后重试

**Q: 浏览器无法自动打开怎么办？**

A: CLI 会显示认证 URL，您可以手动复制 URL 到浏览器中打开。

**Q: 登录超时怎么办？**

A: 认证流程有 5 分钟超时限制。如果超时，请重新运行 `agentbay login`。

**Q: 如何检查当前登录状态？**

A: 运行 `agentbay image list` 或其他需要认证的命令，如果未登录会提示先登录。

### 镜像相关

**Q: 如何查看可用的基础镜像？**

A: 使用 `agentbay image list --system-only` 查看所有可用的系统镜像。

**Q: 镜像构建失败怎么办？**

A: 请检查：
1. Dockerfile 语法是否正确
2. 基础镜像 ID 是否有效
3. 是否修改了 Dockerfile 中系统定义的前 N 行（N 在下载模板时会显示）
4. 使用 `-v` 选项查看详细错误信息
5. 使用 `agentbay image init` 下载模板参考

**Q: Dockerfile 的哪些部分不能修改？**

A: 使用 `agentbay image init` 下载的 Dockerfile 模板中，前 N 行（N 由系统返回）是系统定义的，不能修改。这些行通常包含基础镜像定义和系统必需的配置。命令成功后会显示不可编辑的行数，例如：
```
[IMPORTANT] The first 5 line(s) of the Dockerfile are system-defined and cannot be modified.
[IMPORTANT] Please only modify content after line 5.
```
只能修改第 N+1 行之后的内容。如果修改了前 N 行，可能导致镜像构建失败。

**Q: 如何查看镜像构建状态？**

A: 使用 `agentbay image list` 查看镜像状态，状态为 "Creating" 表示正在构建。

**Q: 激活镜像需要多长时间？**

A: 通常需要 1-2 分钟。激活过程中会显示进度信息。

**Q: 可以同时激活多个镜像吗？**

A: 可以，每个镜像独立管理，互不影响。

**Q: 停用镜像后数据会丢失吗？**

A: 停用镜像会释放计算资源，但镜像本身不会删除，可以重新激活。

### 命令使用

**Q: 如何查看命令帮助？**

A: 使用 `--help` 或 `-h` 选项：
```bash
agentbay --help
agentbay image --help
agentbay image create --help
```

**Q: 如何启用详细日志？**

A: 使用 `-v` 或 `--verbose` 选项：
```bash
agentbay -v image create my-app -f ./Dockerfile -i code-space-debian-12
```

**Q: 配置文件在哪里？**

A: 
- macOS/Linux: `~/.config/agentbay/config.json`
- Windows: `%APPDATA%\agentbay\config.json`

**Q: 如何重置配置？**

A: 删除配置文件后重新登录：
```bash
# macOS/Linux
rm ~/.config/agentbay/config.json

# Windows
del %APPDATA%\agentbay\config.json
```

### 错误处理

**Q: 遇到 "Request ID" 错误怎么办？**

A: 错误信息中会包含 Request ID，请记录此 ID 并联系技术支持。

**Q: 网络连接问题怎么办？**

A: 请检查：
1. 网络连接是否正常
2. 防火墙设置是否阻止了连接
3. 是否能够访问 AgentBay 服务端点

---

## 环境切换

### 概述

AgentBay CLI 支持在生产环境和预发布环境之间切换。此功能主要用于内部开发和测试。

### 环境说明

- **生产环境（production）**: 默认环境，用于正式使用
- **预发布环境（prerelease）**: 用于测试和验证

### 切换方法

#### 临时切换（单次命令）

```bash
AGENTBAY_ENV=prerelease agentbay login
```

#### 会话级切换（当前终端）

```bash
# macOS/Linux
export AGENTBAY_ENV=prerelease
agentbay login
agentbay image list

# Windows (PowerShell)
$env:AGENTBAY_ENV="prerelease"
agentbay login
agentbay image list
```

#### 永久切换（添加到配置文件）

```bash
# macOS/Linux - 添加到 ~/.zshrc 或 ~/.bashrc
echo 'export AGENTBAY_ENV=prerelease' >> ~/.zshrc
source ~/.zshrc

# Windows - 添加到系统环境变量
```

### 切换回生产环境

```bash
# 取消环境变量
unset AGENTBAY_ENV

# 或显式设置为生产环境
export AGENTBAY_ENV=production
```

### 验证当前环境

使用 `agentbay version` 命令查看当前环境：

```bash
agentbay version
```

输出中的 "Environment" 字段显示当前使用的环境。

### 支持的环境值

- **生产环境**: `production`、`prod` 或不设置（默认）
- **预发布环境**: `prerelease`、`pre`、`staging`

### 注意事项

- 不同环境的认证令牌是独立的，需要分别登录
- 不同环境的镜像和资源是隔离的
- 切换环境后需要重新登录
- 此功能主要用于内部测试，普通用户应使用默认的生产环境

---

## 技术支持

如遇到问题，请提供以下信息：

1. CLI 版本（`agentbay version`）
2. 错误信息（包括 Request ID）
3. 操作步骤
4. 系统信息（操作系统、版本）

---

## 附录



### 资源配置参考

支持的 CPU 和内存组合：

| CPU 核心数 | 内存 (GB) | 配置名称 |
|-----------|----------|---------|
| 2 | 4 | 2c4g |
| 4 | 8 | 4c8g |
| 8 | 16 | 8c16g |

---

**文档版本**: 1.0  
**最后更新**: 2025-01-15  
**版权所有**: AgentBay CLI Contributors

