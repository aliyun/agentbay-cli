[English](../en/apikey.md) | **中文**

# API Key 管理 — `agentbay apikey`

管理 API Key：创建、列出、启用、禁用、删除以及设置每个密钥的并发上限。

> 创建 API Key 前账户必须完成实名认证，且每个 API Key 名称必须唯一。

## 术语说明

- **API Key**（`akm-xxx`）：用户可见的密钥值，在管理控制台中展示，由 `apikey list` 返回。对应 `--api-key` 参数。
- **API Key ID**（`ak-xxx`）：内部密钥标识符，由 `apikey create` 返回，在 `apikey list` 输出中显示。对应 `--api-key-id` 参数。

> 交互式操作推荐使用 `--api-key` 参数。自动化脚本从 `apikey create` 输出开始时，使用 `--api-key-id` 更方便。

## 命令

### `apikey create`

创建新的 API Key。

```bash
# 使用位置参数（推荐，更简洁）
agentbay apikey create "my-api-key"

# 使用 --name 参数（向后兼容）
agentbay apikey create --name "my-api-key"
```

**参数：**

| 参数     | 类型   | 必填 | 说明                     |
| -------- | ------ | ---- | ------------------------ |
| `--name` | string | 是   | API Key 名称（必须唯一） |

**输出：** 命令显示新创建密钥的 `ApiKeyId`（ak-xxx 格式）和 `Name`。

**涉及接口：**

| Action         | 所需权限                |
| -------------- | ----------------------- |
| `CreateApiKey` | `agentbay:CreateApiKey` |

```json
{
  "Action": ["agentbay:CreateApiKey"]
}
```

---

### `apikey enable`

重新启用已禁用的 API Key。

```bash
# 使用用户可见的 API Key（推荐）
agentbay apikey enable --api-key akm-xxxxxxxxxxxxxxxx

# 使用内部 API Key ID
agentbay apikey enable --api-key-id ak-xxxxxxxxxxxxxxxx
```

**参数：**

| 参数           | 类型   | 必填 | 说明                                                |
| -------------- | ------ | ---- | --------------------------------------------------- |
| `--api-key`    | string | 是\* | 用户可见的 API Key（akm-xxx 格式，推荐）            |
| `--api-key-id` | string | 是\* | 内部 API Key ID（ak-xxx）。日常使用推荐 `--api-key` |

\* 必须指定 `--api-key` 或 `--api-key-id` 其中之一，不可同时指定。

**涉及接口：**

| Action               | 所需权限                      |
| -------------------- | ----------------------------- |
| `DescribeMcpApiKey`  | `agentbay:DescribeMcpApiKey`  |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` |

```json
{
  "Action": ["agentbay:DescribeMcpApiKey", "agentbay:ModifyApiKeyStatus"]
}
```

---

### `apikey disable`

禁用 API Key（禁用后无法用于认证）。

```bash
# 使用用户可见的 API Key（推荐）
agentbay apikey disable --api-key akm-xxxxxxxxxxxxxxxx

# 使用内部 API Key ID
agentbay apikey disable --api-key-id ak-xxxxxxxxxxxxxxxx
```

**参数：**

| 参数           | 类型   | 必填 | 说明                                                |
| -------------- | ------ | ---- | --------------------------------------------------- |
| `--api-key`    | string | 是\* | 用户可见的 API Key（akm-xxx 格式，推荐）            |
| `--api-key-id` | string | 是\* | 内部 API Key ID（ak-xxx）。日常使用推荐 `--api-key` |

\* 必须指定 `--api-key` 或 `--api-key-id` 其中之一，不可同时指定。

**涉及接口：**

| Action               | 所需权限                      |
| -------------------- | ----------------------------- |
| `DescribeMcpApiKey`  | `agentbay:DescribeMcpApiKey`  |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` |

```json
{
  "Action": ["agentbay:DescribeMcpApiKey", "agentbay:ModifyApiKeyStatus"]
}
```

---

### `apikey delete`

永久删除 API Key。仅 `DISABLED` 状态的 Key 可直接删除；若处于 `ENABLED` 状态会先提示禁用。

```bash
# 使用用户可见的 API Key（交互式，带确认）
agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx

# 使用内部 API Key ID
agentbay apikey delete --api-key-id ak-xxxxxxxxxxxxxxxx

# 跳过所有提示（脚本 / CI）
agentbay apikey delete --api-key akm-xxxxxxxxxxxxxxxx --yes
agentbay apikey delete --api-key-id ak-xxxxxxxxxxxxxxxx -y
```

**参数：**

| 参数           | 短参数 | 类型   | 必填 | 说明                                                |
| -------------- | ------ | ------ | ---- | --------------------------------------------------- |
| `--api-key`    |        | string | 是\* | 用户可见的 API Key（akm-xxx 格式，推荐）            |
| `--api-key-id` |        | string | 是\* | 内部 API Key ID（ak-xxx）。日常使用推荐 `--api-key` |
| `--yes`        | `-y`   |        | 否   | 跳过所有确认提示（非交互模式使用）                  |

\* 必须指定 `--api-key` 或 `--api-key-id` 其中之一，不可同时指定。

**注意事项：**

- 如果 Key 处于 `ENABLED` 状态，命令会先提示禁用再删除。
- 非交互环境中请使用 `--yes` / `-y` 跳过所有提示。

**涉及接口：**

| Action               | 所需权限                      |
| -------------------- | ----------------------------- |
| `DescribeMcpApiKey`  | `agentbay:DescribeMcpApiKey`  |
| `DescribeApiKeys`    | `agentbay:DescribeApiKeys`    |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` |
| `DeleteApiKey`       | `agentbay:DeleteApiKey`       |

```json
{
  "Action": [
    "agentbay:DescribeMcpApiKey",
    "agentbay:DescribeApiKeys",
    "agentbay:ModifyApiKeyStatus",
    "agentbay:DeleteApiKey"
  ]
}
```

---

### `apikey list`

列出 API Key，支持筛选和分页。

```bash
# 最多列出 10 个 API Key
agentbay apikey list

# 最多列出 20 个 API Key
agentbay apikey list --max-results 20

# 按用户可见的 API Key 查询（推荐）
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx

# 按内部 API Key ID 查询
agentbay apikey list --api-key-id ak-xxxxxxxxxxxxxxxx

# 获取下一页
agentbay apikey list --next-token <token>

# JSON 输出（AI/脚本使用）
agentbay apikey list -o json
```

**参数：**

| 参数            | 短参数 | 类型   | 必填 | 说明                                                             |
| --------------- | ------ | ------ | ---- | ---------------------------------------------------------------- |
| `--max-results` |        | int    | 否   | 最大返回数量（默认：10）                                         |
| `--api-key`     |        | string | 否   | 按用户可见的 API Key（akm-xxx 格式）筛选                         |
| `--api-key-id`  |        | string | 否   | 按内部 API Key ID（ak-xxx）筛选。与 `--api-key` 互斥             |
| `--next-token`  |        | string | 否   | 分页 Token                                                       |
| `--output`      | `-o`   | string | 否   | 输出格式。使用 `json` 获取机器可读的完整数据（适合 AI/脚本使用） |

**输出示例：**

使用 `--output json`（或 `-o json`）输出完整 JSON：

```bash
agentbay apikey list -o json
```

```json
{
  "totalCount": 2,
  "apiKeys": [
    {
      "keyId": "ak-xxxxxxxxxxxxxxxx",
      "name": "my-key",
      "apiKey": "akm-xxxxxxxxxxxxxxxx",
      "status": "ENABLED",
      "concurrency": 5,
      "gmtCreate": "2026-01-01T00:00:00.000+00:00",
      "lastUseDate": "2026-01-02T00:00:00.000+00:00"
    }
  ]
}
```

**涉及接口：**

| Action              | 所需权限                     |
| ------------------- | ---------------------------- |
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` |
| `DescribeApiKeys`   | `agentbay:DescribeApiKeys`   |

```json
{
  "Action": ["agentbay:DescribeMcpApiKey", "agentbay:DescribeApiKeys"]
}
```

---

### `apikey concurrency set`

设置 API Key 的最大并发会话数。

```bash
# 使用用户可见的 API Key（推荐）
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10

# 使用内部 API Key ID
agentbay apikey concurrency set --api-key-id ak-xxx --concurrency 10
```

**参数：**

| 参数            | 类型   | 必填 | 说明                                                |
| --------------- | ------ | ---- | --------------------------------------------------- |
| `--api-key`     | string | 是\* | 用户可见的 API Key（akm-xxx 格式，推荐）            |
| `--api-key-id`  | string | 是\* | 内部 API Key ID（ak-xxx）。日常使用推荐 `--api-key` |
| `--concurrency` | int    | 是   | 最大并发会话数（必须 >= 1）                         |

\* 必须指定 `--api-key` 或 `--api-key-id` 其中之一，不可同时指定。

**涉及接口：**

| Action                  | 所需权限                         |
| ----------------------- | -------------------------------- |
| `DescribeMcpApiKey`     | `agentbay:DescribeMcpApiKey`     |
| `ModifyMcpApiKeyConfig` | `agentbay:ModifyMcpApiKeyConfig` |

```json
{
  "Action": ["agentbay:DescribeMcpApiKey", "agentbay:ModifyMcpApiKeyConfig"]
}
```

---

### `apikey describe-key-content`

根据 API Key ID（ak-xxx）查询对应的明文 API Key（akm-xxx 格式）。

```bash
# 根据内部 API Key ID 查询明文 API Key
agentbay apikey describe-key-content --api-key-id ak-xxxxxxxxxxxxxxxx
```

**参数：**

| 参数           | 类型   | 必填 | 说明                           |
| -------------- | ------ | ---- | ------------------------------ |
| `--api-key-id` | string | 是   | 内部 API Key ID（ak-xxx 格式） |

**输出：** 命令显示对应的明文 `ApiKey`（akm-xxx 格式）及查询时使用的 `ApiKeyId`。

**涉及接口：**

| Action               | 所需权限                      |
| -------------------- | ----------------------------- |
| `DescribeKeyContent` | `agentbay:DescribeKeyContent` |

```json
{
  "Action": ["agentbay:DescribeKeyContent"]
}
```
