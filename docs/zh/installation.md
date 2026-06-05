[English](../en/installation.md) | **中文**

# 安装

## Windows

### 快速安装（PowerShell）

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

安装脚本会：

1. 检测系统架构（amd64/arm64）
2. 从 GitHub Releases 下载最新版本
3. 创建安装目录（默认 `%LOCALAPPDATA%\agentbay`）
4. 安装二进制文件 `agentbay.exe`
5. 更新 PATH 环境变量（用户级别）
6. 自动验证安装

### 前置条件

- Windows 10 或更高版本（Windows Server 2016 或更高版本）
- PowerShell 5.1 或更高版本（推荐 PowerShell 7+）
- 网络连接

### 验证安装

```powershell
# 重启 PowerShell 或刷新 PATH，然后：
agentbay version
```

### 更新

```powershell
# 重新执行安装命令即可原地升级
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

### 卸载

> 如果安装时指定了 `-InstallPath` 或设置了 `$env:AGENTBAY_PATH`，请把下面的 `$env:LOCALAPPDATA\agentbay` 替换成实际安装目录。

```powershell
# 删除安装目录
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force

# 从用户 PATH 中移除
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")

# 请重启 PowerShell 让 PATH 变更生效。
```

---

## macOS / Linux（Homebrew）

### 前置条件

- macOS 或 Linux
- 已安装 [Homebrew](https://brew.sh)（Linux 用户参考 [Homebrew on Linux](https://docs.brew.sh/Homebrew-on-Linux)，与 macOS 共用同一个 `brew` 命令）
- 网络连接

### 安装

```bash
# 1. 添加 Homebrew tap
brew tap aliyun/agentbay

# 2. 安装 agentbay
brew install agentbay

# 3. 验证
agentbay version
```

> 在现代 macOS（sonoma / ventura / sequoia，含 Apple Silicon）和 Linux（x86_64 / aarch64）上会直接下载预编译 bottle，通常几秒钟完成。仅当没有匹配 bottle 的环境（如较旧的 macOS）才会从源码编译并自动拉取 Go 作为构建依赖，需要几分钟。

### 更新

日常升级推荐只刷新 `aliyun/agentbay` 这一个 tap，通常几秒钟即可完成：

```bash
git -C "$(brew --repository aliyun/agentbay)" pull --ff-only && brew upgrade agentbay
```

这条命令会跳过 Homebrew 的全量元数据同步（`formula.jws.json` / `cask.jws.json` 等几十 MB 的 JSON 下载，以及 brew 自身升级），只从 tap 拉取最新的 `agentbay` formula，然后直接 pour 新版 bottle。

如果 `brew` 本身开始报错（例如长时间没刷新 brew 之后，或者 Homebrew 出现破坏性更新），改用完整更新作为回退方案：

```bash
brew update && brew upgrade agentbay
```

这条命令会刷新 Homebrew 自身、所有已安装的 tap，以及 core formula 元数据，然后再升级 agentbay。更慢但更彻底。

### 卸载

```bash
brew uninstall agentbay
brew untap aliyun/agentbay   # 可选
```

### 故障排除（Homebrew）

```bash
brew update && brew cleanup
brew reinstall agentbay
```

---

## 预编译二进制

各平台的预编译二进制发布在 [GitHub Releases](https://github.com/aliyun/agentbay-cli/releases) 页面。下载与您操作系统/架构匹配的产物（Linux/macOS 是 `.tar.gz`，Windows 是 `.zip` 或 `.exe`），按下面对应平台的步骤安装。

### Linux

```bash
# 解压(以 amd64 为例，arm64 同理)
tar -xzf agentbay-*-linux-amd64.tar.gz

# 安装到 PATH 内的目录
chmod +x agentbay
sudo mv agentbay /usr/local/bin/

agentbay version
```

### macOS

```bash
# 解压(以 Apple Silicon 为例，Intel 用 darwin-amd64)
tar -xzf agentbay-*-darwin-arm64.tar.gz

chmod +x agentbay

# 浏览器下载会被 Gatekeeper 标记为隔离文件，首次运行会被拦截；
# 使用 curl/wget 下载则无需此步。
xattr -d com.apple.quarantine agentbay 2>/dev/null || true

sudo mv agentbay /usr/local/bin/

agentbay version
```

### Windows

```powershell
# 解压 zip（或直接下载 .exe，跳过此步）
Expand-Archive agentbay-*-windows-amd64.zip -DestinationPath .

# 放到 PATH 内的目录(与一键安装脚本默认位置一致)
$dst = "$env:LOCALAPPDATA\agentbay"
New-Item -ItemType Directory -Force -Path $dst | Out-Null
Move-Item .\agentbay.exe "$dst\agentbay.exe" -Force

# 如果 $dst 不在 PATH 中，把它加入用户 PATH 并重启 PowerShell；
# 或直接使用上文 PowerShell 一键安装脚本。
agentbay version
```
