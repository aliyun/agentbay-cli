[中文](../zh/docker.md) | **English**

# Docker Operations — `agentbay docker`

These commands wrap the local `docker` CLI to interact with the AgentBay ACR registry.

## Commands

### `docker login`

Obtain temporary credentials via the `GetACRRepoCredential` API and **automatically log in to your local docker** (valid for ~1 hour). The command also returns the dedicated image registry path for your account. Both `docker build` and `docker push` rely on this login state. The credential info (`RegistryUrl`, `Namespace`, `RepoName`, `ImageTag`) is also cached locally for `agentbay docker tag` / `agentbay docker push`.

```bash
agentbay docker login
```

**Example output:**

```
Credential expires at: 2026-05-11 12:28:55
Image registry path:   ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674

WARNING! Your credentials are stored unencrypted in '/home/moushuai.ms/.docker/config.json'.
Configure a credential helper to remove this warning. See
https://docs.docker.com/go/credential-store/

Login Succeeded

Note: Credentials will expire after the time above. You can run 'agentbay docker login' again to refresh.
Note: When tagging images, use: ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:<your-tag>
```

**Notes:**

- This calls `GetACRRepoCredential` to obtain temporary ACR credentials **for your own repository only**.
- Credentials are valid for about 1 hour. Both `docker build` and `docker push` depend on this login state; re-run `agentbay docker login` to refresh.
- Credentials are cached locally for subsequent `tag` and `push` commands.
- The command also returns your account's dedicated **image registry path** (the `Image registry path`), which must be used as the prefix when building and pushing images.

**Involved APIs:**

| Action                 | Required Permission             |
| ---------------------- | ------------------------------- |
| `GetACRRepoCredential` | `agentbay:GetACRRepoCredential` |

```json
{
  "Action": ["agentbay:GetACRRepoCredential"]
}
```

> **Docker Environment Setup**: `docker build` and `docker push` require a local Docker Daemon; the CLI itself does not provide a container runtime. Install Docker according to your OS:
>
> - **macOS**: [OrbStack](https://orbstack.dev/) is recommended (lightweight, fast, much lower resource usage than Docker Desktop)
> - **Windows**: Docker Desktop + WSL2 backend is recommended
> - **Linux**: install Docker Engine directly via your system package manager

**After login you can build and push your docker image** (this must be done before running `agentbay image create-from-template`):

```bash
# Build the docker image locally (the tag MUST be prefixed with the Image registry path returned above)
docker build \
  -t ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:<your-tag> \
  -f Dockerfile .

# Push the local docker image to the ACR registry
docker push ai-container-pre-9543-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/1160165251879674:<your-tag>
```

---

### `docker tag`

Tag a local image for the AgentBay ACR registry. The target image name is constructed as `$RegistryUrl/$Namespace/$RepoName:<target-tag>`.

```bash
agentbay docker tag myapp:latest v1.0
```

**Arguments:**

| Argument         | Type   | Required | Description                  |
| ---------------- | ------ | -------- | ---------------------------- |
| `<source-image>` | string | Yes      | Local image name with tag    |
| `<target-tag>`   | string | Yes      | Target tag for the ACR image |

> Run `agentbay docker login` first.

> `docker tag` is a wrapper around the native `docker tag` CLI and does not call any AgentBay API directly. No additional RAM permissions are required.

---

### `docker push`

Push a tagged image to the AgentBay ACR registry. The image name must match `$RegistryUrl/$Namespace/$RepoName[:tag]`; mismatched names are rejected.

```bash
agentbay docker push <registry>/<namespace>/<repo>:v1.0
```

**Arguments:**

| Argument  | Type   | Required | Description                           |
| --------- | ------ | -------- | ------------------------------------- |
| `<image>` | string | Yes      | Full image name matching the ACR path |

> Run `agentbay docker login` first.

> `docker push` is a wrapper around the native `docker push` CLI and does not call any AgentBay API directly. No additional RAM permissions are required.

---

### `docker share`

Grant read-only pull access to your entire Docker image repository to another Alibaba Cloud account.

```bash
agentbay docker share <TARGET_ALI_UID>
agentbay docker share --target-uid <TARGET_ALI_UID>
```

**Arguments:**

| Argument       | Type  | Required | Description                                   |
| -------------- | ----- | -------- | --------------------------------------------- |
| `<target-uid>` | int64 | Yes      | Target Alibaba Cloud account UID (positional) |
| `--target-uid` | int64 | Yes      | Target Alibaba Cloud account UID (named flag) |

**Notes:**

- The target account must be a **primary account** (RAM sub-accounts cannot be sharing targets).
- The recipient has **pull** permission only; they cannot push or delete your images.
- The share is **permanently valid** until explicitly revoked via `docker unshare`.

**Example output:**

```
[INFO] Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Docker repo shared successfully.
  Target UID  : 1234567890
  Owner UID   : 9876543210
  Repo Name   : my-acr-repo
  Status      : active
```

**Involved APIs:**

| Action            | Required Permission        |
| ----------------- | -------------------------- |
| `ShareDockerRepo` | `agentbay:ShareDockerRepo` |

---

### `docker unshare`

Cancel sharing the Docker image repository with a specific Alibaba Cloud account.

```bash
agentbay docker unshare <TARGET_ALI_UID>
agentbay docker unshare --target-uid <TARGET_ALI_UID>
```

**Arguments:**

| Argument       | Type  | Required | Description                                                          |
| -------------- | ----- | -------- | -------------------------------------------------------------------- |
| `<target-uid>` | int64 | Yes      | Target Alibaba Cloud account UID to cancel sharing with (positional) |
| `--target-uid` | int64 | Yes      | Target Alibaba Cloud account UID to cancel sharing with (named flag) |

**Example output:**

```
[INFO] Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
Docker repo unshared successfully.
```

**Involved APIs:**

| Action              | Required Permission          |
| ------------------- | ---------------------------- |
| `UnshareDockerRepo` | `agentbay:UnshareDockerRepo` |

---

### `docker list-shares`

List Docker image repository sharing information. Use `--direction` to specify:

- `Outgoing`: repos you have shared with other accounts
- `Incoming`: repos that other accounts have shared with you

```bash
agentbay docker list-shares --direction Outgoing
agentbay docker list-shares --direction Incoming
agentbay docker list-shares --direction Outgoing --page 2 --size 5
```

**Flags:**

| Flag              | Type   | Required | Default | Description                                           |
| ----------------- | ------ | -------- | ------- | ----------------------------------------------------- |
| `--direction`     | string | Yes      | —       | Sharing direction: `Outgoing` or `Incoming`           |
| `--page`          | int    | No       | 1       | Page number                                           |
| `--size`          | int    | No       | 10      | Page size                                             |
| `--output` / `-o` | string | No       | —       | Output format. Use `json` for machine-readable output |

**Example output (default table):**

Outgoing (repos you shared with others):

```
[INFO] ListSharedDockerRepos Request ID: 89469103-92EF-12BD-BD9B-1F1B9A2F9D6D
PeerAliUid            Status
--------------------  ---------------
****7069              ACTIVE

Total: 1
```

Incoming (repos shared with you):

```
[INFO] ListSharedDockerRepos Request ID: A4C0FF35-AA8A-1BDA-B807-3FA595048431
PeerAliUid            Status
--------------------  ---------------
****4214              ACTIVE

Total: 1
```

**Example output (`--output json`):**

```json
{
  "totalCount": 2,
  "pageNumber": 1,
  "pageSize": 10,
  "items": [
    {
      "peerAliUid": 1234567890,
      "status": "active"
    },
    {
      "peerAliUid": 2345678901,
      "status": "pending"
    }
  ]
}
```

**Involved APIs:**

| Action                  | Required Permission              |
| ----------------------- | -------------------------------- |
| `ListSharedDockerRepos` | `agentbay:ListSharedDockerRepos` |
