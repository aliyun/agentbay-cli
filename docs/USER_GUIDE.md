# AgentBay CLI User Guide

Quick guide to get you started with AgentBay CLI.

## Installation

### Windows

Install AgentBay CLI using PowerShell:

```powershell
powershell -Command "irm https://aliyun.github.io/agentbay-cli/windows | iex"
```

For detailed Windows installation instructions, troubleshooting, and system requirements, see [scripts/README.md](../scripts/README.md).

### macOS/Linux

Install AgentBay CLI using Homebrew:

```bash
# 1. Add Agentbay Cloud's Homebrew tap
brew tap aliyun/agentbay

# 2. Install agentbay command-line tool
brew install agentbay

# 3. Verify installation
agentbay version
```

For more details about Homebrew installation, visit the [ Homebrew-agentbay tap repository](https://github.com/aliyun/homebrew-agentbay).

### Uninstallation

#### Windows
```powershell
# Remove installation directory
Remove-Item -Path "$env:LOCALAPPDATA\agentbay" -Recurse -Force
```

#### macOS/Linux
```bash
# Uninstall agentbay
brew uninstall agentbay

# Remove tap (optional)
brew untap aliyun/agentbay
```

## Prerequisites

- AgentBay CLI installed (see Installation section above)
- Aliyun account
- Network connection

**Supported Image Types**: The current version of the CLI tool supports creating and activating **CodeSpace** type images only.

## 1. Login

```bash
agentbay login
```

The CLI will open your browser for Aliyun authentication. Complete the login and return to the terminal.

**Output:**
```
Starting AgentBay authentication...
Opening browser for authentication...
...
Authentication successful!
You are now logged in to AgentBay!
```

## 2. Logout

```bash
agentbay logout
```

Clears your authentication tokens.

## 3. List Images

```bash
agentbay image list
agentbay image list --os-type Android --size 5
```

**Options:**
- `--os-type, -o`: Filter by OS (Linux, Windows, Android)
- `--page, -p`: Page number (default: 1)
- `--size, -s`: Items per page (default: 10)

**Example output:**
```
[OK] Found 5 images (Total: 17)
IMAGE ID              IMAGE NAME       TYPE            STATUS        OS
--------              ----------       ----            ------        --
imgc-xxxxx...xxx      my-app           ImageBuilder    Available     Android 14
imgc-xxxxx...xxx      web-server       ImageBuilder    Activated     Linux Ubuntu
imgc-xxxxx...xxx      test-img         ImageBuilder    Creating      Windows 2022
```

**Status meanings:**
- **Creating**: Building
- **Available**: Ready to activate
- **Activated**: Running
- **Create Failed**: Build failed

## 4. Create Image

```bash
agentbay image create my-app --dockerfile ./Dockerfile --imageId code_latest
```

**Required:**
- `<image-name>`: Your custom image name
- `--dockerfile, -f`: Path to Dockerfile
- `--imageId, -i`: Base image ID

**Output:**
```
[BUILD] Creating image 'my-app'...
[STEP 1/4] Getting upload credentials... Done.
[STEP 2/4] Uploading Dockerfile... Done.
[STEP 3/4] Creating Docker image task... Done.
[STEP 4/4] Building image (Task ID: task-xxxxx)...
[STATUS] Build status: RUNNING
[SUCCESS] ✅ Image created successfully!
[RESULT] Image ID: imgc-xxxxx...xxx
```

Build time varies based on image size. Use `-v` for detailed logs.

## 5. Activate Image

```bash
agentbay image activate imgc-xxxxx...xxx
```

Starts the image instance.

**Options:**
- `--cpu, -c`: CPU cores (2, 4, or 8) - must be paired with memory
- `--memory, -m`: Memory in GB (4, 8, or 16) - must be paired with CPU

**Supported Resource Combinations:**
- `2c4g` - 2 CPU cores with 4 GB memory
- `4c8g` - 4 CPU cores with 8 GB memory
- `8c16g` - 8 CPU cores with 16 GB memory

**Examples:**
```bash
# Activate with default resources
agentbay image activate imgc-xxxxx...xxx

# Activate with specific resources
agentbay image activate imgc-xxxxx...xxx --cpu 2 --memory 4
agentbay image activate imgc-xxxxx...xxx --cpu 4 --memory 8
agentbay image activate imgc-xxxxx...xxx --cpu 8 --memory 16
```

**Output:**
```
[ACTIVATE] Activating image...
Checking current image status... Done.
Creating resource group... Done.
Waiting for activation to complete...
  Status: Activating (elapsed: 5s, attempt: 2/60)
  Status: Activating (elapsed: 13s, attempt: 3/60)
[SUCCESS] Image activated successfully!
```

Activation typically takes 1-2 minutes. If already activated, you'll see "No action needed."

## 6. Deactivate Image

```bash
agentbay image deactivate imgc-xxxxx...xxx
```

Stops the running image instance.

**Output:**
```
[DEACTIVATE] Deactivating image...
Deleting resource group... Done.
Waiting for deactivation to complete...
  Status: Deactivating (elapsed: 5s, attempt: 2/40)
[SUCCESS] Image deactivated successfully!
```

Usually completes in seconds.

## FAQ

**Q: How to view help?**
```bash
agentbay --help
agentbay image --help
```

**Q: Check CLI version?**
```bash
agentbay version
```

**Q: Enable detailed logs?**
```bash
agentbay -v image list
```

**Q: Login issues?**
- Check network connection
- Ensure browser can access signin.aliyun.com
- Check firewall settings

**Q: Image build fails?**
- Verify Dockerfile syntax
- Check base image ID is valid
- Use `agentbay image list` to find valid IDs

**Q: Where is config stored?**
`~/.config/agentbay/config.json` (macOS/Linux) or `%APPDATA%\agentbay\config.json` (Windows)

**Q: Supported OS types?**
Linux, Windows, Android

---

## Environment Switching (Internal Use Only)

**Note: This section is for internal developers and testing purposes only.**

AgentBay CLI supports switching between production and pre-release environments using the `AGENTBAY_ENV` environment variable.

### Switch to Pre-release Environment

```bash
# Temporary (single command)
AGENTBAY_ENV=prerelease agentbay login

# Session-wide (current terminal)
export AGENTBAY_ENV=prerelease
agentbay login
agentbay image list

# Permanent (add to ~/.zshrc or ~/.bashrc)
echo 'export AGENTBAY_ENV=prerelease' >> ~/.zshrc
source ~/.zshrc
```

### Switch back to Production

```bash
# Remove environment variable
unset AGENTBAY_ENV

# Or explicitly set to production
export AGENTBAY_ENV=production
```

### Verify Current Environment

```bash
agentbay version
```

**Output:**
```
AgentBay CLI version x.x.x
Git commit: xxxxxxx
Build date: 2025-xx-xx
Environment: production
Endpoint: xiaoying-share.cn-shanghai.aliyuncs.com
```

### Supported Environment Values

- Production: `production`, `prod`, or not set (default)
- Pre-release: `prerelease`, `pre`, `staging`

---

For technical support, provide Request ID from error messages.


