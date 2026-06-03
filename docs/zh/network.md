[English](../en/network.md) | **中文**

# 网络管理 — `agentbay network`

按区域查询网络包及其 EIP 绑定信息。

## 命令

### `network package list`

按区域列出网络包。

```bash
agentbay network package list                              # 默认区域 cn-hangzhou
agentbay network package list --biz-region-id cn-shanghai  # 指定区域
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `--biz-region-id` | string | 否 | 区域 ID（默认：`cn-hangzhou`） |

**输出示例：**

```
[LIST] Fetching network packages...
Requesting network packages... Done. (Action: DescribeNetworkPackages, Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx)

[OK] Found 1 network package(s)

NETWORK PACKAGE ID         OFFICE SITE ID                     EIP ADDRESSES
------------------         --------------                     -------------
np-xxxxxxxxxxxx            cn-hangzhou+dir-xxxxxxxxxxxx       1.2.3.4
```

**输出列说明：**

| 列名 | 说明 |
|---|---|
| `NETWORK PACKAGE ID` | 网络包唯一标识符 |
| `OFFICE SITE ID` | 关联的办公网络 ID |
| `EIP ADDRESSES` | 绑定的弹性公网 IP 地址（可能为多个） |

**注意事项：**

- 默认查询 `cn-hangzhou` 区域，可通过 `--biz-region-id` 指定其他区域。
- 指定区域下没有网络包时，会输出 `[EMPTY] No network packages found.`。
- 使用 `-v` / `--verbose` 可额外打印 Request ID 等调试信息。

**涉及接口：**

| Action | 所需权限 |
|---|---|
| `DescribeNetworkPackages` | `agentbay:DescribeNetworkPackages` |

```json
{
  "Action": [
    "agentbay:DescribeNetworkPackages"
  ]
}
```
