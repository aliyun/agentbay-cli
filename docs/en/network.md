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

**Output example:**

```
[LIST] Fetching network packages...
Requesting network packages... Done. (Action: DescribeNetworkPackages, Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)

[OK] Found 1 network package(s)

NETWORK PACKAGE ID         OFFICE SITE ID                     EIP ADDRESSES
------------------         --------------                     -------------
np-xxxxxxxxxxxx            cn-hangzhou+dir-xxxxxxxxxxxx       1.2.3.4
```

**Output columns:**

| Column | Description |
|---|---|
| `NETWORK PACKAGE ID` | Unique identifier of the network package |
| `OFFICE SITE ID` | ID of the office network the package belongs to |
| `EIP ADDRESSES` | Elastic IP addresses bound to the package (may be multiple) |

**Notes:**

- The default region is `cn-hangzhou`; use `--biz-region-id` to query another region.
- When no network packages exist in the region, the command prints `[EMPTY] No network packages found.`
- Use `-v` / `--verbose` to additionally print the Request ID and other debug information.

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
