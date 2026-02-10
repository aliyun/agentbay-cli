# AgentBay CLI User Guide

Quick guide to get you started with AgentBay CLI.

## Prerequisites

- AgentBay CLI installed
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
agentbay image list                    # List user images (default)
agentbay image list --include-system   # List both user and system images
agentbay image list --system-only      # List only system images
agentbay image list --os-type Android --size 5
```

**Options:**
- `--os-type, -o`: Filter by OS (Linux, Windows, Android)
- `--include-system`: Include system images in addition to user images
- `--system-only`: Show only system images
- `--page, -p`: Page number (default: 1)
- `--size, -s`: Items per page (default: 10)

**Example output:**
```
=== USER IMAGES (17) ===
IMAGE ID              IMAGE NAME       TYPE                 STATUS        OS                 APPLY SCENE
--------              ----------       ----                 ------        --                 -----------
imgc-xxxxx...xxx      my-app           DockerBuilder        Available     Android 14         CodeSpace
imgc-xxxxx...xxx      web-server       DockerBuilder        Available     Linux Ubuntu       CodeSpace
imgc-xxxxx...xxx      test-img         DockerBuilder        Creating      Windows 2022       CodeSpace

=== SYSTEM IMAGES (3) ===
IMAGE ID                  IMAGE NAME                     TYPE                 STATUS        OS                 APPLY SCENE
--------                  ----------                     ----                 ------        --                 -----------
mobile-use-android-14     Mobile Use Android 14          DedicatedDesktop     Available     Android 14         MobileUse
computer-use-windows-2022 Computer Use Windows Server... DedicatedDesktop     Available     Windows 2022       ComputerUse
computer-use-ubuntu-2204  Computer Use Linux Ubuntu 2... DedicatedDesktop     Available     Linux Ubuntu 2204  ComputerUse
```

**Status meanings:**
- **Creating**: Building
- **Available**: Ready to activate
- **Activated**: Running
- **Create Failed**: Build failed

**Type meanings:**
- **DockerBuilder**: User-created images built from Dockerfile
- **DedicatedDesktop**: System images or dedicated desktop images


**Note**: System images are always available and don't require activation. Only user-created images need to be activated before use.

## 4. Download Dockerfile Template

```bash
agentbay image init
```

Downloads a Dockerfile template from the cloud and saves it as `Dockerfile` in the current directory.

**Output:**
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

**Note**: 
- If a `Dockerfile` already exists in the current directory, it will be overwritten. The command will warn you before overwriting.
- **Important**: The first N lines (N is returned by the system) of the Dockerfile template are system-defined and cannot be modified. Only modify content after line N+1, otherwise the image build may fail.

## 5. Create Image

```bash
agentbay image create my-app --dockerfile ./Dockerfile --imageId code-space-debian-12
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
[STEP 3/4] Uploading ADD/COPY files (N files)... Done.   # Only when Dockerfile contains COPY/ADD
[STEP 4/4] Creating Docker image task... Done.
[STEP 4/4] Building image (Task ID: task-xxxxx)...
[STATUS] Build status: RUNNING
[SUCCESS] Image created successfully!
[RESULT] Image ID: imgc-xxxxx...xxx
```

Build time varies based on image size. Use `-v` for detailed logs.

### ADD/COPY File Upload

When creating an image, the CLI parses `COPY` and `ADD` instructions in your Dockerfile and automatically uploads the referenced local files:

- **Path rules**: File paths are relative to the directory containing the Dockerfile
- **Supported**: Single files, multiple files, subdirectories, wildcards (e.g. `*.py`), `--chown` option
- **Not supported**: Absolute paths, path traversal (e.g. `../`), URL sources in `ADD` (e.g. `ADD https://...`)
- **Note**: Ensure all files referenced by COPY/ADD exist in the Dockerfile directory or its subdirectories

## 6. Activate Image

User-created images need to be activated before use. System images are always available and don't require activation.

```bash
agentbay image activate imgc-xxxxx...xxx
```

Starts the image instance.

**Options:**
- `--cpu, -c`: CPU cores (2, 4, or 8) - must be paired with memory; default: 2 when not specified
- `--memory, -m`: Memory in GB (4, 8, or 16) - must be paired with CPU; default: 4 when not specified

**Supported Resource Combinations:**
- `2c4g` - 2 CPU cores with 4 GB memory **(default when --cpu/--memory not specified)**
- `4c8g` - 4 CPU cores with 8 GB memory
- `8c16g` - 8 CPU cores with 16 GB memory

**Examples:**
```bash
# Activate with default resources (2c4g)
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

## 7. Deactivate Image

Deactivate custom images when done to save resources. Deactivating an activated user image releases related resources.

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
- Check base image ID is valid (use `agentbay image list --include-system` to find valid system image IDs)
- Check if you modified the first N lines of the Dockerfile (N is shown when downloading the template)
- Use `agentbay image init` to download a template Dockerfile
- Use `-v` option to view detailed error information

**Q: Which parts of the Dockerfile cannot be modified?**
- The first N lines (N is returned by the system) of the Dockerfile template downloaded by `agentbay image init` are system-defined and cannot be modified
- These lines typically contain base image definitions and system-required configurations
- Only modify content after line N+1, otherwise the image build may fail

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
Endpoint: xiaoying.cn-shanghai.aliyuncs.com
```

### Supported Environment Values

- Production: `production`, `prod`, or not set (default)
- Pre-release: `prerelease`, `pre`, `staging`

---

For technical support, provide Request ID from error messages.


