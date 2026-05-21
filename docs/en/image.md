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
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--os-type` | `-o` | string | No | Filter by OS (Linux, Windows, Android) |
| `--include-system` | | | No | Include system images in addition to user images |
| `--system-only` | | | No | Show only system images |
| `--page` | `-p` | int | No | Page number (default: 1) |
| `--size` | `-s` | int | No | Items per page (default: 10) |

**Output example:**

```
=== USER IMAGES (17) ===
IMAGE ID              IMAGE NAME       TYPE                 STATUS        OS                 APPLY SCENE
--------              ----------       ----                 ------        --                 -----------
imgc-xxxxx...xxx      my-app           DockerBuilder        Available     Android 14         CodeSpace

=== SYSTEM IMAGES (3) ===
IMAGE ID                  IMAGE NAME                     TYPE                 STATUS        OS                 APPLY SCENE
--------                  ----------                     ----                 ------        --                 -----------
mobile-use-android-14     Mobile Use Android 14          DedicatedDesktop     Available     Android 14         MobileUse
```

**Notes:**

- System images are always available and don't require activation. Only user-created images need to be activated before use.
- **Status meanings**: Creating, Available, Activated, Create Failed
- **Type meanings**: DockerBuilder (user-created), DedicatedDesktop (system)

---

### `image init`

Download a Dockerfile template from the cloud to the current directory.

```bash
agentbay image init --sourceImageId code-space-debian-12
agentbay image init -i code-space-debian-12
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--sourceImageId` | `-i` | string | Yes | System image ID to use as base |

Available `sourceImageId` values for production:

- `code-space-debian-12`
- `code-space-debian-12-enhanced`

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

Create a custom image from a system image template (calls the `CreateImageFromTemplate` API).

```bash
agentbay image create-from-template \
  --source-image registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 \
  --name my-custom-image \
  --imageId <system-image-id>

# Short form
agentbay image create-from-template -s registry.cn-hangzhou.aliyuncs.com/myrepo/myimage:v1.0 -n my-custom-image -i <system-image-id>
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--source-image` | `-s` | string | Yes | Source image registry path |
| `--name` | `-n` | string | Yes | Custom image name |
| `--imageId` | `-i` | string | Yes | Base system image ID |

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

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--cpu` | `-c` | int | No | CPU cores (2, 4, or 8); must pair with `--memory` |
| `--memory` | `-m` | int | No | Memory in GB (4, 8, or 16); must pair with `--cpu` |
| `--network-type` | | string | No | Network type: `DEFAULT` or `ADVANCED` |
| `--session-bandwidth` | | int | No | Session bandwidth (required for ADVANCED network) |
| `--dns-address` | | string | No | DNS address (repeatable; required for ADVANCED network) |
| `--lifecycle-mode` | | string | No | Lifecycle mode: `auto` or `manual` |
| `--lifecycle-max-runtime` | | int | No | Max runtime in seconds |
| `--lifecycle-hibernate` | | int | No | Hibernate timeout in seconds |
| `--lifecycle-idle-timeout` | | int | No | Idle timeout in seconds |
| `--region-id` | | string | No | Region ID for resource deployment |

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

---

### `image delete`

Delete a User image **permanently**. Only deactivated User images can be deleted.

```bash
agentbay image delete imgc-xxxxxxxxxxxxxx          # With confirmation
agentbay image delete imgc-xxxxxxxxxxxxxx --yes    # Skip confirmation (CI / scripts)
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--yes` | `-y` | | No | Skip confirmation prompt (required in non-interactive mode) |

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

---

### `image status`

Query the resource lifecycle status of an image (different from the Docker build task status during `image create`).

```bash
agentbay image status imgc-xxxxxxxxxxxxxx
```

**Common status values:**

| Status | Meaning |
|--------|---------|
| `IMAGE_CREATING` | Image is being created |
| `IMAGE_CREATE_FAILED` | Image creation failed |
| `IMAGE_AVAILABLE` | Available, not activated |
| `RESOURCE_DEPLOYING` | Activation in progress |
| `RESOURCE_PUBLISHED` | Activated and in use |
| `RESOURCE_DELETING` | Deactivation in progress |
| `RESOURCE_FAILED` | Activation or resource operation failed |
| `RESOURCE_CEASED` | Resource ceased |

---

### `image set-max-session`

Set the maximum concurrent session count for an activated User image. Requires the image to be in `RESOURCE_PUBLISHED` state and use **advanced network**.

```bash
agentbay image set-max-session --image-id imgc-xxxxxxxxxxxxxx --max-session-num 10
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--image-id` | string | Yes | Image ID |
| `--max-session-num` | int | Yes | Maximum concurrent sessions |

> The command polls until the resource group is ready (typically ~5 minutes).

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
