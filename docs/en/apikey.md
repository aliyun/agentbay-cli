[中文](../../zh/apikey.md) | **English**

# API Key Management — `agentbay apikey`

Manage API keys: create, list, enable, disable, delete, and set per-key concurrency limits.

> API key creation requires account real-name verification. Each API key must have a unique name.

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

---

### `apikey enable`

Re-enable a disabled API key.

```bash
agentbay apikey enable akm-xxxxxxxxxxxxxxxx
```

**Arguments:**

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<api-key-id>` | string | Yes | API key ID (e.g., `akm-xxxxxxxxxxxxxxxx`) |

---

### `apikey disable`

Disable an API key (it can no longer authenticate requests).

```bash
agentbay apikey disable akm-xxxxxxxxxxxxxxxx
```

**Arguments:**

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<api-key-id>` | string | Yes | API key ID |

---

### `apikey delete`

Delete an API key permanently. Only `DISABLED` keys can be deleted directly; if the key is `ENABLED` you will be prompted to disable it first.

```bash
agentbay apikey delete akm-xxxxxxxxxxxxxxxx          # Interactive (with confirmation)
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes    # Skip all prompts (CI / scripts)
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--yes` | `-y` | | No | Skip all confirmation prompts (for non-interactive use) |

**Notes:**

- If the key is `ENABLED`, the command will prompt you to disable it first before deletion.
- In non-interactive environments, use `--yes` / `-y` to skip all prompts.

---

### `apikey list`

List API keys with optional filtering and pagination.

```bash
agentbay apikey list                                        # List up to 10 API keys
agentbay apikey list --max-results 20                       # List up to 20 API keys
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx         # Query a specific API key
agentbay apikey list --next-token <token>                   # Fetch the next page
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--max-results` | int | No | Maximum number of results (default: 10) |
| `--api-key` | string | No | Query a specific API key by ID |
| `--next-token` | string | No | Pagination token for next page |

---

### `apikey concurrency set`

Set the maximum concurrent session limit for an API key.

```bash
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--api-key` | string | Yes | API key ID |
| `--concurrency` | int | Yes | Maximum concurrent sessions |
