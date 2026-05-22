[中文](../zh/network.md) | **English**

# Network Management — `agentbay network`

Query network packages and EIP bindings by region.

## Commands

### `network package list`

List network packages for a region.

```bash
agentbay network package list                              # Default region: cn-hangzhou
agentbay network package list --biz-region-id cn-shanghai  # Custom region
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--biz-region-id` | string | No | Region ID (default: `cn-hangzhou`) |

**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeNetworkPackages` | `agentbay:DescribeNetworkPackages` |

```json
{
  "Action": [
    "agentbay:DescribeNetworkPackages"
  ]
}
```
