[中文](../../zh/installation.md) | **English**

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

### Uninstallation

```powershell
# Remove installation directory
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force

# Remove from user PATH
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
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

### Update

```bash
brew upgrade agentbay
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

Pre-built binaries are also available under `bin/` and `packages/` in the repository. Download the appropriate binary for your platform, make it executable, and place it in your PATH.

```bash
chmod +x agentbay
sudo mv agentbay /usr/local/bin/
agentbay version
```
