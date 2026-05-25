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

```bash
# 1. 添加 Homebrew tap
brew tap aliyun/agentbay

# 2. 安装 agentbay
brew install agentbay

# 3. 验证
agentbay version
```

> 首次安装会从源码编译，并自动安装 Go 作为构建依赖，整个过程可能需要几分钟。后续升级会复用缓存。

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

各平台的预编译二进制发布在 [GitHub Releases](https://github.com/aliyun/agentbay-cli/releases) 页面。下载与您操作系统/架构匹配的二进制，添加执行权限并放入 PATH 中。

```bash
chmod +x agentbay
sudo mv agentbay /usr/local/bin/
agentbay version
```
