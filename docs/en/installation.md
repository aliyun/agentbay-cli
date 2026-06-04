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

### Prerequisites

- macOS or Linux
- [Homebrew](https://brew.sh) installed (Linux users: see [Homebrew on Linux](https://docs.brew.sh/Homebrew-on-Linux) — same `brew` command as macOS)
- Internet connection

### Install

```bash
# 1. Add the Homebrew tap
brew tap aliyun/agentbay

# 2. Install agentbay
brew install agentbay

# 3. Verify
agentbay version
```

> On modern macOS (sonoma / ventura / sequoia, including Apple Silicon) and Linux (x86_64 / aarch64), `brew install` pours a pre-built bottle — usually done in seconds. Only environments without a matching bottle (e.g., older macOS) fall back to building from source and will pull Go as a build dependency, taking a few minutes.

### Update

For routine updates, refresh only the `aliyun/agentbay` tap — usually completes in seconds:

```bash
git -C "$(brew --repository aliyun/agentbay)" pull --ff-only && brew upgrade agentbay
```

This skips Homebrew's full metadata sync (large `formula.jws.json` / `cask.jws.json` downloads and brew self-update) and only pulls the latest `agentbay` formula from the tap before pouring the new bottle.

If `brew` itself starts reporting errors (e.g., after a long time without a refresh, or after a Homebrew breaking change), fall back to the full update:

```bash
brew update && brew upgrade agentbay
```

This refreshes Homebrew itself, all installed taps, and the core formula metadata before upgrading. Slower but more thorough.

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

Pre-built binaries for every platform are published on [GitHub Releases](https://github.com/aliyun/agentbay-cli/releases). Download the asset that matches your OS/architecture (`.tar.gz` for Linux/macOS, `.zip` or `.exe` for Windows), then follow the steps for your platform.

### Linux

```bash
# Extract (amd64 shown; arm64 is analogous)
tar -xzf agentbay-*-linux-amd64.tar.gz

# Install into a PATH directory
chmod +x agentbay
sudo mv agentbay /usr/local/bin/

agentbay version
```

### macOS

```bash
# Extract (Apple Silicon shown; Intel uses darwin-amd64)
tar -xzf agentbay-*-darwin-arm64.tar.gz

chmod +x agentbay

# Browser downloads are quarantined by Gatekeeper and blocked on first run;
# this step is unnecessary if you downloaded with curl/wget.
xattr -d com.apple.quarantine agentbay 2>/dev/null || true

sudo mv agentbay /usr/local/bin/

agentbay version
```

### Windows

```powershell
# Extract the zip (or skip this step if you downloaded the standalone .exe)
Expand-Archive agentbay-*-windows-amd64.zip -DestinationPath .

# Move into a PATH directory (matches the one-click installer's default)
$dst = "$env:LOCALAPPDATA\agentbay"
New-Item -ItemType Directory -Force -Path $dst | Out-Null
Move-Item .\agentbay.exe "$dst\agentbay.exe" -Force

# If $dst is not on PATH, add it to your user PATH and restart PowerShell;
# alternatively use the PowerShell one-click installer above.
agentbay version
```
