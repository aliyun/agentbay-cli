[中文](../zh/image-workflow.md) | **English**

# Image Creation & Sharing

This document is split into two chapters:

- [Part 1: Image Creation](#part-1-image-creation) — the flow any account can complete on its own: build a custom image from a Dockerfile template and register it.
- [Part 2: Image Sharing](#part-2-image-sharing) — share your image repository with another Alibaba Cloud account (optional).

---

## Part 1: Image Creation

**Scenario**: A single account builds an image from a Dockerfile template, pushes it to ACR, and registers it as an activatable custom image.

> The examples in this chapter use account UID `****4214`. If you only need to create images for yourself, you can stop at the end of this chapter.

### Step 1: Download the Dockerfile Template

```bash
agentbay image init --sourceImageId aio-ubuntu-2404
```

### Step 2: Log in to Docker

Obtain temporary credentials and automatically log in to local docker (valid for ~1 hour). The command also returns the dedicated image registry path.

```bash
agentbay docker login
```

Example output:

```
Credential expires at: 2026-05-11 12:28:55
Image registry path:   ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214

Login Succeeded
```

> `docker login` returns temporary credentials for your own repository, automatically logs in to local docker, valid for 1 hour (both `docker build` and `docker push` depend on this login state). It also returns the image upload address.

### Step 3: Build the Image Locally

Make sure Docker is installed locally before running:

- **macOS**: [OrbStack](https://orbstack.dev/) is recommended (lightweight, fast, much lower resource usage than Docker Desktop)
- **Windows**: Docker Desktop + WSL2 backend is recommended
- **Linux**: install Docker Engine directly via your system package manager

```bash
docker build \
  -t ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1 \
  -f Dockerfile .
```

### Step 4: Push the Image to ACR

```bash
docker push ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1
```

### Step 5: Create the Custom Image

The short path `/namespace/repo:tag` is recommended (matches the `physicalImage` field returned by `image list`):

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****4214:cli-test-0.0.1 \
  --name cli-template-create-1 \
  --imageId aio-ubuntu-2404
```

Example output:

```
[IMAGE] Creating custom image from template...
  SourceImage:      ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1
  PhysicalImageId:  /customer_cli/****4214:cli-test-0.0.1
  Name:             cli-template-create-1
  ImageId:          aio-ubuntu-2404
Requesting CreateImageFromTemplate... Done. (HTTP 200)

[RESPONSE]
  RequestId:      6921CD80-B902-1294-9C4A-6078DBBA7B2F
  Code:           ok
  Message:
  Success:        true
  HttpStatusCode: 200

[DATA]
  ImageId: imgc-0a9mg1hbjw1b7r564

[SUCCESS] CreateImageFromTemplate call completed.
```

### Key Notes

1. **`--source-image` Format**: The short path `/namespace/repo:tag` is recommended (consistent with `image list` output); full registry paths are also supported.
2. **Physical Image ID**: The `PhysicalImageId` returned by `image create-from-template` can be used directly as the short path; you can also find it in the `physicalImage` field from `image list`.

---

## Part 2: Image Sharing

**Scenario**: Account A (UID: `****4214`) shares its entire Docker image repository with Account B (UID: `****7069`); Account B then creates its own custom image from one of A's images.

> Prerequisite: Account A has completed every step in [Part 1: Image Creation](#part-1-image-creation) and the repository contains at least one image.

### Account A (the sharer)

#### Step 1: Share the Repository with Account B

> **Note**: The target account must be a **primary account** (RAM sub-accounts cannot be sharing targets).

Share your entire Docker image repository with the specified user for read-only pull access. The recipient has pull permission only, cannot push or delete your images. The authorization is permanent until explicitly revoked via `agentbay docker unshare`.

```bash
agentbay docker share --target-uid ****7069
```

#### Step 2: Verify the Share

```bash
agentbay docker list-shares --direction Outgoing
```

Example output:

```
[INFO] ListSharedDockerRepos Request ID: 89469103-92EF-12BD-BD9B-1F1B9A2F9D6D
PeerAliUid            Status
--------------------  ---------------
****7069              ACTIVE

Total: 1
```

#### Step 3 (Optional): Revoke the Share

If you no longer want Account B to pull your images, revoke the authorization at any time. `target-uid` can be passed as a positional argument or via the `--target-uid` flag:

```bash
agentbay docker unshare ****7069
# or
agentbay docker unshare --target-uid ****7069
```

Example output:

```
[STEP 1/1] Cancelling Docker repo sharing with UID ****7069...
[INFO] UnshareDockerRepo Request ID: 7F2A1B3C-4D5E-6F70-8192-A3B4C5D6E7F8

[SUCCESS] Docker repo sharing cancelled.
  Revoked : true
```

After revocation, run `agentbay docker list-shares --direction Outgoing` again to confirm the entry is gone.

### Account B (the recipient)

#### Step 1: View Incoming Shares

```bash
agentbay docker list-shares --direction Incoming
```

Example output:

```
[INFO] ListSharedDockerRepos Request ID: A4C0FF35-AA8A-1BDA-B807-3FA595048431
PeerAliUid            Status
--------------------  ---------------
****4214              ACTIVE

Total: 1
```

#### Step 2: Create a Custom Image from A's Image

Use the `PhysicalImageId` returned by Account A's `image create-from-template` success, or find it via `image list`, then follow [Part 1 → Step 5](#step-5-create-the-custom-image) to create the image:

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****4214:cli-test-0.0.1 \
  --name cli-test \
  --imageId aio-ubuntu-2404
```

### Key Notes

1. **Permission Scope**: The recipient has **pull** permission only; they cannot push or delete images in Account A's repository.
2. **Authorization Duration**: The share is **permanently valid** until Account A explicitly calls `agentbay docker unshare` to revoke it.
3. **Primary Account Required**: The target account must be an Alibaba Cloud **primary account**; RAM sub-accounts cannot be sharing targets.
