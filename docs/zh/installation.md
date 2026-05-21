[English](../../en/installation.md) | **中文**

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

### 卸载

```powershell
# 删除安装目录
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force

# 从用户 PATH 中移除
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
```

---

## macOS / Linux（Homebrew）

```bash
# 1. 添加 Homebrew tap
brew tap aliyun/agentbay

# 2. 安装 agentbay
brew install agentbay

# 3. 验证
agentbay version
```

### 更新

```bash
brew upgrade agentbay
```

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

预编译的二进制文件也可在仓库的 `bin/` 和 `packages/` 目录下获取。下载适合您平台的二进制文件，添加执行权限并放入 PATH 中。

```bash
chmod +x agentbay
sudo mv agentbay /usr/local/bin/
agentbay version
```
