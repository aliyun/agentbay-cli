# Agentbay CLI Windows Installation Guide

This guide provides instructions for installing Agentbay CLI on Windows using PowerShell.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Verification](#verification)
- [Usage](#usage)
- [Troubleshooting](#troubleshooting)
- [Uninstallation](#uninstallation)

## Prerequisites

Before installing Agentbay CLI, please ensure:
- Windows 10 or later (Windows Server 2016 or later)
- PowerShell 5.1 or later (PowerShell 7+ recommended)
- Internet connection
- Administrator privileges (recommended for PATH configuration)

## Installation

### Quick Installation

Install Agentbay CLI with a single PowerShell command:

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

### Installation Process

The installation script will:
1. **Detect system architecture** (amd64/arm64)
2. **Download the latest version** from GitHub Releases
3. **Create installation directory** (`%LOCALAPPDATA%\agentbay` by default)
4. **Install the binary** as `agentbay.exe`
5. **Update PATH environment variable** (user-level)
6. **Verify installation** automatically

## Verification

After installation, verify that Agentbay CLI is installed correctly:

### Step 1: Restart PowerShell
```powershell
# Close current PowerShell window and open a new one
Start-Process powershell -Verb RunAs; exit
# Or refresh the environment variables
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
```

### Step 2: Check Installation
```powershell
# Check if agentbay command is available
agentbay --version
```

**Expected Output:**
```
Agentbay CLI version current_version
Git commit: current_commit_id
Build date: current_build_date
```

### Step 3: Verify Command Help
```powershell
# Display help information
agentbay --help
```

**Expected Output:**
```
Command line interface for Agentbay services

Usage:
  agentbay [command]

Available Commands:
  image       Manage images
  login       Log in to Agentbay
  logout      Log out from Agentbay
  version     Show version information
  help        Help about any command

Flags:
  -h, --help      help for agentbay
  -v, --verbose   Enable verbose output

Use "agentbay [command] --help" for more information about a command.
```

### Step 4: Test Core Functionality
```powershell
# Test image command
agentbay image --help

# Test version command
agentbay version
```

## Usage

### Basic Commands

```powershell
# Show help
agentbay --help

# Show version
agentbay version

# Login to Agentbay
agentbay login

# List available images
agentbay image list
```

For more detailed usage instructions, examples, and advanced features, see the [User Guide](../docs/USER_GUIDE.md).


### Enable Verbose Output
```powershell
# Use -v flag for detailed output
agentbay -v image list
agentbay --verbose login
```

## Troubleshooting

### Common Issues

#### Issue 1: Command Not Found
```powershell
# Error: 'agentbay' is not recognized as an internal or external command
```

**Solutions:**
1. **Restart PowerShell** or refresh environment variables:
   ```powershell
   $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
   ```

2. **Check installation directory:**
   ```powershell
   Get-ChildItem "$env:LOCALAPPDATA\agentbay"
   ```

#### Issue 2: Installation Failed
```powershell
# Error: Failed to download Agentbay CLI
```

**Solutions:**
1. **Check internet connection**
2. **Try running the installation command again**
3. **Use manual download from GitHub Releases if needed**

#### Issue 3: Permission Denied
```powershell
# Error: Access denied or execution policy restriction
```

**Solutions:**
1. **Run as Administrator:**
   ```powershell
   # Right-click PowerShell and select "Run as Administrator"
   ```

2. **Check execution policy:**
   ```powershell
   Get-ExecutionPolicy
   Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
   ```


### Getting Help

If you encounter issues:
1. **Check the installation log** for error messages
2. **Verify system requirements** (Windows version, PowerShell version)
3. **Try manual installation** from GitHub Releases
4. **Contact support** with error details and system information

## Uninstallation

To remove Agentbay CLI:

### Step 1: Remove Binary
```powershell
# Remove installation directory
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force
```

### Step 2: Clean PATH
```powershell
# Remove from user PATH
$agentbayPath = "$env:LOCALAPPDATA\agentbay"
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $agentbayPath }) -join ';'
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
```

### Step 3: Verify Removal
```powershell
# This should return an error
agentbay --version
```

---

## Additional Information

### System Requirements
- **OS**: Windows 10/11, Windows Server 2016+
- **Architecture**: x64 (amd64) or ARM64
- **PowerShell**: 5.1+ (7+ recommended)
- **Disk Space**: ~50MB
- **Network**: Internet connection for download

### Installation Locations
- **Default**: `%LOCALAPPDATA%\agentbay\agentbay.exe`

### Links
- **GitHub Repository**: https://github.com/aliyun/agentbay-cli
- **Documentation**:  [USER_GUIDE.md](../docs/USER_GUIDE.md)
- **Issues**: https://github.com/aliyun/agentbay-cli/issues

---

**Note**: This installation method downloads the latest stable release. For development versions or specific releases, please visit the GitHub Releases page.