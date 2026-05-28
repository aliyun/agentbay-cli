[中文](../zh/image-workflow.md) | **English**

# Image Creation & Sharing Workflow

Below is an end-to-end example: Account A builds a custom image from a Dockerfile template and shares it with Account B.

## Scenario

- **Account A** (UID: `****4214`): creates a custom image and shares it with Account B
- **Account B** (UID: `****7069`): receives the shared repository and creates its own custom image from it

---

## Account A Workflow

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

### Step 6: Share the Repository with Account B

> **Note**: The target account must be a **primary account** (RAM sub-accounts cannot be sharing targets).

Share your entire Docker image repository with the specified user for read-only pull access. The recipient has pull permission only, cannot push or delete your images. The authorization is permanent until explicitly revoked via `docker unshare`.

```bash
agentbay docker share --target-uid ****7069
```

### Step 7: Verify the Share

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

---

## Account B Workflow

### View Incoming Shares

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

### Create a Custom Image from A's Image

Use the `PhysicalImageId` returned by Account A's `image create-from-template` success, or find it via `image list`, and create using the short path format:

```bash
agentbay image create-from-template \
  --source-image /customer_cli/****4214:cli-test-0.0.1 \
  --name cli-test \
  --imageId aio-ubuntu-2404
```

---

## Key Notes

1. **Permission Scope**: The recipient has **pull** permission only; they cannot push or delete images in Account A's repository.
2. **Authorization Duration**: The share is **permanently valid** until Account A explicitly calls `docker unshare` to revoke it.
3. **Physical Image ID**: The `PhysicalImageId` returned by `image create-from-template` can be used directly as the short path; you can also find it in the `physicalImage` field from `image list`.
4. **`--source-image` Format**: The short path `/namespace/repo:tag` is recommended (consistent with `image list` output); full registry paths are also supported.
