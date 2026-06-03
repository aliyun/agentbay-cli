[中文](../zh/image.md) | **English**

# Image Management — `agentbay image`

Manage AgentBay images: create, list, activate, deactivate, delete, and query lifecycle status.

> The current CLI version supports creating and activating **CodeSpace** type images only.

## Commands

### `image list`

List available AgentBay images.

```bash
agentbay image list                      # User images (default)
agentbay image list --include-system     # User + system images
agentbay image list --system-only        # System images only
agentbay image list --os-type Linux      # Filter by OS type: Linux / Android / Windows
agentbay image list --page 2 --size 5    # Pagination
agentbay image list --output json        # JSON output (for AI/scripts)
```

**Flags:**

| Flag               | Short | Type   | Required | Description                                                                        |
| ------------------ | ----- | ------ | -------- | ---------------------------------------------------------------------------------- |
| `--os-type`        | `-o`  | string | No       | Filter by OS (Linux, Windows, Android)                                             |
| `--include-system` |       |        | No       | Include system images in addition to user images                                   |
| `--system-only`    |       |        | No       | Show only system images                                                            |
| `--page`           | `-p`  | int    | No       | Page number (default: 1)                                                           |
| `--size`           | `-s`  | int    | No       | Items per page (default: 10)                                                       |
| `--output`         |       | string | No       | Output format. Use `json` for machine-readable complete data (e.g. for AI/scripts) |

**Output example:**

Default table output:

```
=== USER IMAGES (17) ===
IMAGE ID              IMAGE NAME       TYPE                 STATUS        OS                 PHYSICAL IMAGE                 APPLY SCENE
--------              ----------       ----                 ------        --                 --------------                 -----------
imgc-xxxxx...xxx      my-app           DockerBuilder        Available     Android 14         registry.example.com/img:tag   CodeSpace

=== SYSTEM IMAGES (3) ===
IMAGE ID                  IMAGE NAME                     TYPE                 STATUS        OS                 PHYSICAL IMAGE                 APPLY SCENE
--------                  ----------                     ----                 ------        --                 --------------                 -----------
mobile-use-android-14     Mobile Use Android 14          DedicatedDesktop     Available     Android 14                                        MobileUse
```

Use `--output json` for complete JSON output (note: `-o` short flag is taken by `--os-type` on this command, use the full flag name):

```bash
agentbay image list --output json
```

```json
{
  "totalCount": 2,
  "images": [
    {
      "imageId": "imgc-xxxxxxxxxxxxxx",
      "imageName": "my-app",
      "type": "DockerBuilder",
      "status": "IMAGE_AVAILABLE",
      "statusDisplay": "Available",
      "osName": "Linux",
      "osVersion": "Debian 12",
      "osDisplay": "Linux Debian 12",
      "physicalImage": "registry.example.com/my-app:latest",
      "applyScene": "CodeSpace"
    }
  ]
}
```

**Notes:**

- System images are always available and don't require activation. Only user-created images need to be activated before use.
- **Status meanings**: Creating, Available, Activated, Create Failed
- **Type meanings**: DockerBuilder (user-created), DedicatedDesktop (system)

**Involved APIs:**

| Action          | Required Permission      |
| --------------- | ------------------------ |
| `ListMcpImages` | `agentbay:ListMcpImages` |

```json
{
  "Action": ["agentbay:ListMcpImages"]
}
```

---

### `image init`

Download a Dockerfile template from the cloud to the current directory.

```bash
agentbay image init --sourceImageId code-space-debian-12
agentbay image init -i code-space-debian-12
```

**Flags:**

| Flag              | Short | Type   | Required | Description                    |
| ----------------- | ----- | ------ | -------- | ------------------------------ |
| `--sourceImageId` | `-i`  | string | Yes      | System image ID to use as base |

Available `sourceImageId` values for production:

- `code-space-debian-12`
- `code-space-debian-12-enhanced`
- `aio-ubuntu-2404`

**Output:**

```
[INIT] Downloading Dockerfile template...
Requesting Dockerfile template... Done.
Downloading Dockerfile from OSS... Done.
Writing Dockerfile to /path/to/current/directory/Dockerfile... Done.
[SUCCESS] Dockerfile template downloaded successfully!
[IMPORTANT] The first 5 line(s) of the Dockerfile are system-defined and cannot be modified.
[IMPORTANT] Please only modify content after line 5.
```

**Notes:**

- If a `Dockerfile` already exists, it will be overwritten.
- The first N lines of the Dockerfile are system-defined and must not be modified. Only edit content after line N+1.

**Involved APIs:**

| Action                  | Required Permission              |
| ----------------------- | -------------------------------- |
| `GetDockerfileTemplate` | `agentbay:GetDockerfileTemplate` |

```json
{
  "Action": ["agentbay:GetDockerfileTemplate"]
}
```

---

### `image create` _(deprecated — use `create-from-template` instead)_

> WARNING: `image create` is **deprecated and will be removed in a future release**. To create a custom image, please use [`image create-from-template`](#image-create-from-template) instead.

Build a custom image from a Dockerfile. Files referenced by `COPY` / `ADD` are parsed and uploaded automatically.

```bash
agentbay image create myapp --dockerfile ./Dockerfile --imageId code-space-debian-12
agentbay image create myapp -f ./Dockerfile -i code-space-debian-12
```

---

### `image create-from-template`

Create a custom image from a system image template + a Docker image already pushed to AgentBay ACR (calls the `CreateImageFromTemplate` API). The source image can come from your own repository or from another account's Docker repository shared with you.

> **Prerequisites**: Before running this command, you need a Docker image already pushed to the AgentBay ACR registry. The full workflow is:
>
> 1. `agentbay image init -i <system-image-id>` — download the base Dockerfile template and edit the editable section as needed.
> 2. `agentbay docker login` — automatically log in to local docker and obtain the image registry path (valid for ~1 hour).
> 3. Build the Docker image locally — make sure Docker is installed before running `docker build`:
>    - **macOS**: [OrbStack](https://orbstack.dev/) is recommended (lightweight, fast, much lower resource usage than Docker Desktop)
>    - **Windows**: Docker Desktop + WSL2 backend is recommended
>    - **Linux**: install Docker Engine directly via your system package manager
>
>    Then run `docker build -t <registry-path>:<your-tag> -f Dockerfile .`
>
> 4. `docker push <registry-path>:<your-tag>` — push to ACR.
> 5. `agentbay image create-from-template ...` — create the custom image based on the pushed Docker image (this command).

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****9674:<your-tag> \
  --name imageName \
  --imageId code-space-debian-12

# Short form
agentbay image create-from-template -s /customer_cli/****9674:<your-tag> -n imageName -i code-space-debian-12
```

**Flags:**

| Flag             | Short | Type   | Required | Description                                                        |
| ---------------- | ----- | ------ | -------- | ------------------------------------------------------------------ |
| `--source-image` | `-s`  | string | Yes      | Source Docker image registry path (with tag) already pushed to ACR |
| `--name`         | `-n`  | string | Yes      | Custom image name                                                  |
| `--imageId`      | `-i`  | string | Yes      | Base system image ID (e.g. `code-space-debian-12`)                 |

`--source-image` supports two formats:

1. **Recommended**: short path `/customer_cli/<aliuid>:tag` (matches the `physicalImage` field returned by `image list`, can be copied directly)
2. **Also supported**: full registry path `<registry>/customer_cli/<aliuid>:tag`

The CLI first extracts the AliUID from `source-image`:

- If the AliUID matches the local ACR cache created by `agentbay docker login`, the image is treated as your own repository image.
- If it does not match, the CLI calls `ListSharedDockerRepos` to check whether the current account has received Docker repository sharing authorization from that AliUID. The command continues only when data is returned.
- Short paths for your own repository are expanded in terminal output with the current account registry URL; short paths for shared repositories remain short to avoid implying they belong to the current account's ACR path.
- The command prints OpenAPI Request IDs. Shared repository flows print the `ListSharedDockerRepos` Request ID first, then the `CreateImageFromTemplate` Request ID.

**Creation flow (server-side):**

1. Locate the Docker image specified by `--source-image`.
2. Apply the configuration parameters from the system image identified by `--imageId`.
3. Create a Docker custom image.

**Notes for the supplied Docker image:**

1. **Do NOT include any `CMD` instruction.**
2. **Do NOT modify the `FROM` instruction** (keep the `FROM` from the `agentbay image init` template).
3. The Dockerfile **must** end with `USER root`, or never specify a `USER` at all.
4. Currently **only Linux x86 architecture is supported**.
5. If you need an EntryPoint, rewrite it as a supervisor program entry:

   ```dockerfile
   RUN echo '[program:user-define]' > /etc/supervisor/conf.d/user-define.conf && \
       echo 'command=%s' >> /etc/supervisor/conf.d/user-define.conf && \
       echo 'priority=33' >> /etc/supervisor/conf.d/user-define.conf
   ```

   Replace `%s` with the actual command to run.

**Involved APIs:**

| Action                    | Required Permission                |
| ------------------------- | ---------------------------------- |
| `ListSharedDockerRepos`   | `agentbay:ListSharedDockerRepos`   |
| `CreateImageFromTemplate` | `agentbay:CreateImageFromTemplate` |

```json
{
  "Action": [
    "agentbay:ListSharedDockerRepos",
    "agentbay:CreateImageFromTemplate"
  ]
}
```

---

### `image activate`

Activate a User image so it can be used.

```bash
# Default resources (2c4g)
agentbay image activate imgc-xxxxxxxxxxxxxx

# Specific CPU and memory (must be specified together)
agentbay image activate imgc-xxxxxxxxxxxxxx --cpu 2 --memory 4

# Advanced network
agentbay image activate imgc-xxxxxxxxxxxxxx \
  --network-type ADVANCED \
  --session-bandwidth 100 \
  --dns-address 8.8.8.8 \
  --dns-address 8.8.4.4

# Sandbox lifecycle
agentbay image activate imgc-xxxxxxxxxxxxxx \
  --lifecycle-mode auto \
  --lifecycle-max-runtime 3600 \
  --lifecycle-hibernate 1800 \
  --lifecycle-idle-timeout 600

# Specify region
agentbay image activate imgc-xxxxxxxxxxxxxx --region-id cn-shanghai
```

**Flags:**

| Flag                       | Short | Type   | Required | Description                                             |
| -------------------------- | ----- | ------ | -------- | ------------------------------------------------------- |
| `--cpu`                    | `-c`  | int    | No       | CPU cores (2, 4, or 8); must pair with `--memory`       |
| `--memory`                 | `-m`  | int    | No       | Memory in GB (4, 8, or 16); must pair with `--cpu`      |
| `--network-type`           |       | string | No       | Network type: `DEFAULT` or `ADVANCED`                   |
| `--session-bandwidth`      |       | int    | No       | Session bandwidth (required for ADVANCED network)       |
| `--dns-address`            |       | string | No       | DNS address (repeatable; required for ADVANCED network) |
| `--lifecycle-mode`         |       | string | No       | Release mode: `auto` (auto-release) or `manual` (manual release)               |
| `--lifecycle-max-runtime`  |       | int    | No       | Max runtime per session (minutes); requires `--lifecycle-mode auto`            |
| `--lifecycle-hibernate`    |       | int    | No       | Max hibernate duration (hours); requires `--lifecycle-mode auto`               |
| `--lifecycle-idle-timeout` |       | int    | No       | Max idle duration (minutes); requires `--lifecycle-mode auto`                  |
| `--region-id`              |       | string | No       | Region ID for resource deployment                       |

**Supported resource combinations:** `2c4g` (default), `4c8g`, `8c16g`

**Notes:**

- `--cpu` and `--memory` must be specified together.
- `--network-type ADVANCED` requires `--session-bandwidth` and `--dns-address`.
- Activation typically takes 1-2 minutes. If already activated, you'll see "No action needed."

**Output:**

```
[ACTIVATE] Activating image...
Checking current image status... Done.
Creating resource group... Done.
Waiting for activation to complete...
  Status: Activating (elapsed: 5s, attempt: 2/60)
[SUCCESS] Image activated successfully!
```

**Involved APIs:**

| Action                  | Required Permission              |
| ----------------------- | -------------------------------- |
| `GetMcpImageInfo`       | `agentbay:GetMcpImageInfo`       |
| `DescribeInstanceTypes` | `agentbay:DescribeInstanceTypes` |
| `DescribeMcpPolicyData` | `agentbay:DescribeMcpPolicyData` |
| `CreateMcpPolicyData`   | `agentbay:CreateMcpPolicyData`   |
| `ModifyMcpPolicyData`   | `agentbay:ModifyMcpPolicyData`   |
| `DescribeOfficeSites`   | `agentbay:DescribeOfficeSites`   |
| `SaveMcpPolicyData`     | `agentbay:SaveMcpPolicyData`     |
| `CreateResourceGroup`   | `agentbay:CreateResourceGroup`   |

```json
{
  "Action": [
    "agentbay:GetMcpImageInfo",
    "agentbay:DescribeInstanceTypes",
    "agentbay:DescribeMcpPolicyData",
    "agentbay:CreateMcpPolicyData",
    "agentbay:ModifyMcpPolicyData",
    "agentbay:DescribeOfficeSites",
    "agentbay:SaveMcpPolicyData",
    "agentbay:CreateResourceGroup"
  ]
}
```

---

### `image deactivate`

Deactivate an activated User image.

```bash
agentbay image deactivate imgc-xxxxxxxxxxxxxx
```

**Output:**

```
[DEACTIVATE] Deactivating image...
Deleting resource group... Done.
Waiting for deactivation to complete...
  Status: Deactivating (elapsed: 5s, attempt: 2/40)
[SUCCESS] Image deactivated successfully!
```

Usually completes in seconds.

**Involved APIs:**

| Action                | Required Permission            |
| --------------------- | ------------------------------ |
| `GetMcpImageInfo`     | `agentbay:GetMcpImageInfo`     |
| `ListMcpImages`       | `agentbay:ListMcpImages`       |
| `DeleteResourceGroup` | `agentbay:DeleteResourceGroup` |

```json
{
  "Action": [
    "agentbay:GetMcpImageInfo",
    "agentbay:ListMcpImages",
    "agentbay:DeleteResourceGroup"
  ]
}
```

---

### `image delete`

Delete a User image **permanently**. Only deactivated User images can be deleted.

```bash
agentbay image delete imgc-xxxxxxxxxxxxxx          # With confirmation
agentbay image delete imgc-xxxxxxxxxxxxxx --yes    # Skip confirmation (CI / scripts)
```

**Flags:**

| Flag    | Short | Type | Required | Description                                                 |
| ------- | ----- | ---- | -------- | ----------------------------------------------------------- |
| `--yes` | `-y`  |      | No       | Skip confirmation prompt (required in non-interactive mode) |

**Restrictions:**

- Only User images can be deleted (System images cannot be deleted)
- Images in `IMAGE_CREATING`, `RESOURCE_DEPLOYING`, `RESOURCE_DELETING`, `RESOURCE_PUBLISHED`, `RESOURCE_FAILED`, or `RESOURCE_MAINTAINING` states cannot be deleted

**Output:**

```
[DELETE] Deleting image 'imgc-xxxxx'...
Checking current image status... Done.
[INFO] GetMcpImageInfo Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
[INFO] Image Type: User
[INFO] Current Status: Available (Deactivated)
Are you sure you want to permanently delete image 'imgc-xxxxx'? This action is irreversible. [y/N]: y
Deleting image... Done.
[INFO] DeleteMcpImage Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
[SUCCESS] Image 'imgc-xxxxx' has been permanently deleted.
```

**Involved APIs:**

| Action            | Required Permission        |
| ----------------- | -------------------------- |
| `GetMcpImageInfo` | `agentbay:GetMcpImageInfo` |
| `DeleteMcpImage`  | `agentbay:DeleteMcpImage`  |

```json
{
  "Action": ["agentbay:GetMcpImageInfo", "agentbay:DeleteMcpImage"]
}
```

---

### `image status`

Query the resource lifecycle status of an image (different from the Docker build task status during `image create`).

```bash
agentbay image status imgc-xxxxxxxxxxxxxx
```

**Common status values:**

| Status                | Meaning                                 |
| --------------------- | --------------------------------------- |
| `IMAGE_CREATING`      | Image is being created                  |
| `IMAGE_CREATE_FAILED` | Image creation failed                   |
| `IMAGE_AVAILABLE`     | Available, not activated                |
| `RESOURCE_DEPLOYING`  | Activation in progress                  |
| `RESOURCE_PUBLISHED`  | Activated and in use                    |
| `RESOURCE_DELETING`   | Deactivation in progress                |
| `RESOURCE_FAILED`     | Activation or resource operation failed |
| `RESOURCE_CEASED`     | Resource ceased                         |

**Involved APIs:**

| Action            | Required Permission        |
| ----------------- | -------------------------- |
| `GetMcpImageInfo` | `agentbay:GetMcpImageInfo` |

```json
{
  "Action": ["agentbay:GetMcpImageInfo"]
}
```

---

### `image set-max-session`

Set the maximum concurrent session count for an activated User image. Requires the image to be in `RESOURCE_PUBLISHED` state and use **advanced network**.

```bash
agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 10
```

**Flags:**

| Flag                | Type   | Required | Description                 |
| ------------------- | ------ | -------- | --------------------------- |
| `--image-id`        | string | Yes      | Image ID                    |
| `--max-session-num` | int    | Yes      | Maximum concurrent sessions |

> The command polls until the resource group is ready (typically ~5 minutes).

**Involved APIs:**

| Action                                        | Required Permission                                    |
| --------------------------------------------- | ------------------------------------------------------ |
| `GetMcpImageInfo`                             | `agentbay:GetMcpImageInfo`                             |
| `BatchCreateHideResourceGroupsWithMaxSession` | `agentbay:BatchCreateHideResourceGroupsWithMaxSession` |

```json
{
  "Action": [
    "agentbay:GetMcpImageInfo",
    "agentbay:BatchCreateHideResourceGroupsWithMaxSession"
  ]
}
```

---

### `image warmup-status`

Query the warm-up status for the current account, including session quota, image quota, and details of warm-up images.

```bash
agentbay image warmup-status
```

**Output includes:**

- **Session Quota** — max session limit, total used, and available sessions
- **Image Quota** — max image count and current image count
- **Warm-up Images** — table of image IDs, total max size, and group count

**Involved APIs:**

| Action                     | Required Permission                 |
| -------------------------- | ----------------------------------- |
| `DescribeWarmUpStatusOpen` | `agentbay:DescribeWarmUpStatusOpen` |

```json
{
  "Action": ["agentbay:DescribeWarmUpStatusOpen"]
}
```
