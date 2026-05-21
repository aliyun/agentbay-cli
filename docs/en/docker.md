[中文](../zh/docker.md) | **English**

# Docker Operations — `agentbay docker`

These commands wrap the local `docker` CLI to interact with the AgentBay ACR registry.

## Commands

### `docker login`

Log in to the AgentBay ACR registry using temporary credentials obtained from the `GetACRRepoCredential` API. The credential info (`RegistryUrl`, `Namespace`, `RepoName`, `ImageTag`) is cached for `tag` / `push`.

```bash
agentbay docker login
```

**Notes:**

- This calls `GetACRRepoCredential` to obtain temporary ACR credentials.
- Credentials are cached locally for subsequent `tag` and `push` commands.

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
