[中文](../zh/installation.md) | **English**

# Installation

## Windows

### Quick Installation (PowerShell)

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

The installation script will:

1. Detect system architecture (amd64/arm64)
2. Download the latest version from GitHub Releases
3. Create installation directory (`%LOCALAPPDATA%\agentbay` by default)
4. Install the binary as `agentbay.exe`
5. Update PATH environment variable (user-level)
6. Verify installation automatically

### Prerequisites

- Windows 10 or later (Windows Server 2016 or later)
- PowerShell 5.1 or later (PowerShell 7+ recommended)
- Internet connection

### Verify Installation

```powershell
# Restart PowerShell or refresh PATH, then:
agentbay version
```

### Update

```powershell
# Re-run the install command to upgrade in place
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

### Uninstallation

> If you installed with a custom `-InstallPath` or `$env:AGENTBAY_PATH`, replace `$env:LOCALAPPDATA\agentbay` below with your actual install directory.

```powershell
# Remove installation directory
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force

# Remove from user PATH
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")

# Restart PowerShell so the PATH change takes effect.
```

---

## macOS / Linux (Homebrew)

```bash
# 1. Add the Homebrew tap
brew tap aliyun/agentbay

# 2. Install agentbay
brew install agentbay

# 3. Verify
agentbay version
```

> The first install builds from source and will automatically pull Go as a build dependency, so it may take a few minutes. Subsequent upgrades reuse the cache.

### Update

```bash
brew update && brew upgrade agentbay
```

### Uninstall

```bash
brew uninstall agentbay
brew untap aliyun/agentbay   # optional
```

### Troubleshooting (Homebrew)

```bash
brew update && brew cleanup
brew reinstall agentbay
```

---

## Pre-built Binaries

Pre-built binaries for every platform are published on [GitHub Releases](https://github.com/aliyun/agentbay-cli/releases). Download the binary that matches your OS/architecture, make it executable, and place it in your PATH.

```bash
chmod +x agentbay
sudo mv agentbay /usr/local/bin/
agentbay version
```
