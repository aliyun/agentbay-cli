[English](../../en/apikey.md) | **中文**

# API Key 管理 — `agentbay apikey`

管理 API Key：创建、列出、启用、禁用、删除以及设置每个密钥的并发上限。

> 创建 API Key 前账户必须完成实名认证，且每个 API Key 名称必须唯一。

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

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `--name` | string | 是 | API Key 名称（必须唯一） |

---

### `apikey enable`

重新启用已禁用的 API Key。

```bash
agentbay apikey enable akm-xxxxxxxxxxxxxxxx
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `<api-key-id>` | string | 是 | API Key ID（如 `akm-xxxxxxxxxxxxxxxx`） |

---

### `apikey disable`

禁用 API Key（禁用后无法用于认证）。

```bash
agentbay apikey disable akm-xxxxxxxxxxxxxxxx
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `<api-key-id>` | string | 是 | API Key ID |

---

### `apikey delete`

永久删除 API Key。仅 `DISABLED` 状态的 Key 可直接删除；若处于 `ENABLED` 状态会先提示禁用。

```bash
agentbay apikey delete akm-xxxxxxxxxxxxxxxx          # 交互式（带确认）
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes    # 跳过所有提示（脚本 / CI）
```

**参数：**

| 参数 | 短参数 | 类型 | 必填 | 说明 |
|------|--------|------|------|------|
| `--yes` | `-y` | | 否 | 跳过所有确认提示（非交互模式使用） |

**注意事项：**

- 如果 Key 处于 `ENABLED` 状态，命令会先提示禁用再删除。
- 非交互环境中请使用 `--yes` / `-y` 跳过所有提示。

---

### `apikey list`

列出 API Key，支持筛选和分页。

```bash
agentbay apikey list                                        # 最多列出 10 个 API Key
agentbay apikey list --max-results 20                       # 最多列出 20 个 API Key
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx         # 查询指定的 API Key
agentbay apikey list --next-token <token>                   # 获取下一页
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `--max-results` | int | 否 | 最大返回数量（默认：10） |
| `--api-key` | string | 否 | 查询指定 API Key |
| `--next-token` | string | 否 | 分页 Token |

---

### `apikey concurrency set`

设置 API Key 的最大并发会话数。

```bash
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `--api-key` | string | 是 | API Key ID |
| `--concurrency` | int | 是 | 最大并发会话数 |
