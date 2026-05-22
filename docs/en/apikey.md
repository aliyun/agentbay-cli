[中文](../zh/apikey.md) | **English**

# API Key Management — `agentbay apikey`

Manage API keys: create, list, enable, disable, delete, and set per-key concurrency limits.

> API key creation requires account real-name verification. Each API key must have a unique name.

## Terminology

- **API Key** (`akm-xxx`): The user-visible key value displayed in the management console and returned by `apikey list`. Used with the `--api-key` flag.
- **API Key ID** (`ak-xxx`): The internal key identifier returned by `apikey create` and shown in `apikey list` output. Used with the `--api-key-id` flag.

> The `--api-key` flag is recommended for interactive use. The `--api-key-id` flag is useful for automation scripts that start from `apikey create` output.

## Commands

### `apikey create`

Create a new API key.

```bash
# Positional argument (recommended)
agentbay apikey create "my-api-key"

# --name flag (backward compatible)
agentbay apikey create --name "my-api-key"
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--name` | string | Yes | API key name (must be unique) |

**Output:** The command displays `ApiKeyId` (ak-xxx format) and `Name` of the newly created key.

---

### `apikey enable`

Re-enable a disabled API key.

```bash
# Enable using the user-visible API Key (recommended)
agentbay apikey enable --api-key akm-xxxxxxxxxxxxxxxx

# Enable using the internal API Key ID
agentbay apikey enable --api-key-id ak-xxxxxxxxxxxxxxxx
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--api-key` | string | Yes* | User-visible API Key (akm-xxx format, recommended) |
| `--api-key-id` | string | Yes* | Internal API Key ID (ak-xxx). Prefer `--api-key` for normal usage |

\* Either `--api-key` or `--api-key-id` must be specified, but not both.

---

### `apikey disable`

Disable an API key (it can no longer authenticate requests).

```bash
# Disable using the user-visible API Key (recommended)
agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx

# Disable using the internal API Key ID
agentbay apikey disable --api-key-id ak-xxxxxxxxxxxxxxxx
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--api-key` | string | Yes* | User-visible API Key (akm-xxx format, recommended) |
| `--api-key-id` | string | Yes* | Internal API Key ID (ak-xxx). Prefer `--api-key` for normal usage |

\* Either `--api-key` or `--api-key-id` must be specified, but not both.

---

### `apikey delete`

Delete an API key permanently. Only `DISABLED` keys can be deleted directly; if the key is `ENABLED` you will be prompted to disable it first.

```bash
# Delete using the user-visible API Key (interactive, with confirmation)
agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx

# Delete using the internal API Key ID
agentbay apikey delete --api-key-id ak-xxxxxxxxxxxxxxxx

# Skip all prompts (CI / scripts)
agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx --yes
agentbay apikey delete --api-key-id ak-xxxxxxxxxxxxxxxx -y
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--api-key` | | string | Yes* | User-visible API Key (akm-xxx format, recommended) |
| `--api-key-id` | | string | Yes* | Internal API Key ID (ak-xxx). Prefer `--api-key` for normal usage |
| `--yes` | `-y` | | No | Skip all confirmation prompts (for non-interactive use) |

\* Either `--api-key` or `--api-key-id` must be specified, but not both.

**Notes:**

- If the key is `ENABLED`, the command will prompt you to disable it first before deletion.
- In non-interactive environments, use `--yes` / `-y` to skip all prompts.

---

### `apikey list`

List API keys with optional filtering and pagination.

```bash
# List up to 10 API keys (default)
agentbay apikey list

# List up to 20 API keys
agentbay apikey list --max-results 20

# Query a specific API key by its user-visible value (recommended)
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx

# Query a specific API key by its internal ID
agentbay apikey list --api-key-id ak-xxxxxxxxxxxxxxxx

# Fetch the next page
agentbay apikey list --next-token <token>
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--max-results` | int | No | Maximum number of results (default: 10) |
| `--api-key` | string | No | Filter by user-visible API Key (akm-xxx format) |
| `--api-key-id` | string | No | Filter by internal API Key ID (ak-xxx). `--api-key` and `--api-key-id` are mutually exclusive |
| `--next-token` | string | No | Pagination token for next page |

---

### `apikey concurrency set`

Set the maximum concurrent session limit for an API key.

```bash
# Set concurrency using the user-visible API Key (recommended)
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10

# Set concurrency using the internal API Key ID
agentbay apikey concurrency set --api-key-id ak-xxx --concurrency 10
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--api-key` | string | Yes* | User-visible API Key (akm-xxx format, recommended) |
| `--api-key-id` | string | Yes* | Internal API Key ID (ak-xxx). Prefer `--api-key` for normal usage |
| `--concurrency` | int | Yes | Maximum concurrent sessions (must be >= 1) |

\* Either `--api-key` or `--api-key-id` must be specified, but not both.

---

### `apikey describe-key-content`

Retrieve the plaintext API key (akm-xxx format) for a given API key ID (ak-xxx).

```bash
# Retrieve the plaintext API key for a given API key ID
agentbay apikey describe-key-content --api-key-id ak-xxxxxxxxxxxxxxxx
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--api-key-id` | string | Yes | Internal API key ID (ak-xxx format) |

**Output:** The command displays the plaintext `ApiKey` (akm-xxx format) and the `ApiKeyId` used to query it.
