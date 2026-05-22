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

- This calls `GetACRRepoCredential` to obtain temporary ACR credentials.
- Credentials are valid for about 1 hour. Re-run `agentbay docker login` to refresh.
- Credentials are cached locally for subsequent `tag` and `push` commands.

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

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<source-image>` | string | Yes | Local image name with tag |
| `<target-tag>` | string | Yes | Target tag for the ACR image |

> Run `agentbay docker login` first.

---

### `docker push`

Push a tagged image to the AgentBay ACR registry. The image name must match `$RegistryUrl/$Namespace/$RepoName[:tag]`; mismatched names are rejected.

```bash
agentbay docker push <registry>/<namespace>/<repo>:v1.0
```

**Arguments:**

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<image>` | string | Yes | Full image name matching the ACR path |

> Run `agentbay docker login` first.
