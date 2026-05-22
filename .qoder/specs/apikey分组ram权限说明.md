# 计划：apikey 文档增加 RAM 账号接口权限说明

## Context

使用 AgentBay CLI 的用户，如果使用的是阿里云 RAM 子账号的 AK/SK，需要为该 RAM 账号授予对应 OpenAPI 接口的权限（Action 级别的 RAM Policy）。当前 `docs/en/apikey.md` 和 `docs/zh/apikey.md` 缺少这部分说明，导致 RAM 账号用户在使用 CLI 时可能遭遇权限报错。

主账号不需要配置，只有 RAM 子账号才需要在 RAM 控制台（https://ram.console.aliyun.com/users）配置接口权限。

---

## OpenAPI 接口统计（去重后，共 7 个）

| OpenAPI Action | RAM 权限字符串 | 被调用的命令 |
|---|---|---|
| `CreateApiKey` | `agentbay:CreateApiKey` | `apikey create` |
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` | `apikey enable`、`apikey disable`、`apikey delete`、`apikey list`、`apikey concurrency set` |
| `DescribeApiKeys` | `agentbay:DescribeApiKeys` | `apikey delete`、`apikey list` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` | `apikey enable`、`apikey disable`、`apikey delete` |
| `DeleteApiKey` | `agentbay:DeleteApiKey` | `apikey delete` |
| `ModifyMcpApiKeyConfig` | `agentbay:ModifyMcpApiKeyConfig` | `apikey concurrency set` |
| `DescribeKeyContent` | `agentbay:DescribeKeyContent` | `apikey describe-key-content` |

---

## 每个子命令涉及的接口（不区分分支，列举所有可能调用的接口）

| 命令 | 涉及的 OpenAPI Action |
|---|---|
| `apikey create` | `CreateApiKey` |
| `apikey enable` | `DescribeMcpApiKey`、`ModifyApiKeyStatus` |
| `apikey disable` | `DescribeMcpApiKey`、`ModifyApiKeyStatus` |
| `apikey delete` | `DescribeMcpApiKey`、`DescribeApiKeys`、`ModifyApiKeyStatus`、`DeleteApiKey` |
| `apikey list` | `DescribeMcpApiKey`、`DescribeApiKeys` |
| `apikey concurrency set` | `DescribeMcpApiKey`、`ModifyMcpApiKeyConfig` |
| `apikey describe-key-content` | `DescribeKeyContent` |

---

## 实现方案

### 文档结构变更

**变更 1：在 `README.md` 和 `README.zh-CN.md` 的 `## Authentication` 章节之后，新增 `## RAM Permissions` 章节**

包含内容：
- 说明主账号无需配置，RAM 子账号需要在 RAM 控制台授权
- RAM 控制台地址：https://ram.console.aliyun.com/users
- apikey 分组下所有接口的权限清单表格（Action + 权限字符串 + 使用命令）
- 完整授权 JSON Policy 示例

**变更 2：在 `docs/en/apikey.md` 和 `docs/zh/apikey.md` 每个子命令文档节点内，在 Flags 表格之后新增 `**Involved APIs:**` 接口表格 + JSON Action 数组**

格式如下（以 `apikey enable` 为例）：

```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` |

```json
{
  "Action": [
    "agentbay:DescribeMcpApiKey",
    "agentbay:ModifyApiKeyStatus"
  ]
}
```
```

### 涉及文件

- `docs/en/apikey.md`
- `docs/zh/apikey.md`

---

## 详细内容：英文文档 Required Permissions 章节

插入位置：`## Terminology` 之后、`## Commands` 之前。

```markdown
## Required Permissions

> **Note:** The main Alibaba Cloud account does not require additional permission configuration.
> This section only applies to **RAM sub-accounts** using AK/SK authentication.

If you are using a RAM sub-account's AK/SK, grant the following permissions via the
[RAM console](https://ram.console.aliyun.com/users) before using `agentbay apikey` commands.

### All Permissions for `apikey` Command Group

| OpenAPI Action | Required Permission | Used By |
|---|---|---|
| `CreateApiKey` | `agentbay:CreateApiKey` | `apikey create` |
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` | `apikey enable`, `apikey disable`, `apikey delete`, `apikey list`, `apikey concurrency set` |
| `DescribeApiKeys` | `agentbay:DescribeApiKeys` | `apikey delete`, `apikey list` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` | `apikey enable`, `apikey disable`, `apikey delete` |
| `DeleteApiKey` | `agentbay:DeleteApiKey` | `apikey delete` |
| `ModifyMcpApiKeyConfig` | `agentbay:ModifyMcpApiKeyConfig` | `apikey concurrency set` |
| `DescribeKeyContent` | `agentbay:DescribeKeyContent` | `apikey describe-key-content` |

### RAM Policy Example (Full Access to apikey Commands)

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:CreateApiKey",
        "agentbay:DescribeMcpApiKey",
        "agentbay:DescribeApiKeys",
        "agentbay:ModifyApiKeyStatus",
        "agentbay:DeleteApiKey",
        "agentbay:ModifyMcpApiKeyConfig",
        "agentbay:DescribeKeyContent"
      ],
      "Resource": "*"
    }
  ]
}
```

> If you only use specific commands, refer to the **Involved APIs** section under each command
> and grant only the required subset of permissions.
```

---

## 详细内容：每个子命令的 Involved APIs 节点

### apikey create
```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `CreateApiKey` | `agentbay:CreateApiKey` |

```json
{
  "Action": [
    "agentbay:CreateApiKey"
  ]
}
```
```

### apikey enable
```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` |

```json
{
  "Action": [
    "agentbay:DescribeMcpApiKey",
    "agentbay:ModifyApiKeyStatus"
  ]
}
```
```

### apikey disable
（与 enable 相同）

### apikey delete
```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` |
| `DescribeApiKeys` | `agentbay:DescribeApiKeys` |
| `ModifyApiKeyStatus` | `agentbay:ModifyApiKeyStatus` |
| `DeleteApiKey` | `agentbay:DeleteApiKey` |

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
```

### apikey list
```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` |
| `DescribeApiKeys` | `agentbay:DescribeApiKeys` |

```json
{
  "Action": [
    "agentbay:DescribeMcpApiKey",
    "agentbay:DescribeApiKeys"
  ]
}
```
```

### apikey concurrency set
```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeMcpApiKey` | `agentbay:DescribeMcpApiKey` |
| `ModifyMcpApiKeyConfig` | `agentbay:ModifyMcpApiKeyConfig` |

```json
{
  "Action": [
    "agentbay:DescribeMcpApiKey",
    "agentbay:ModifyMcpApiKeyConfig"
  ]
}
```
```

### apikey describe-key-content
```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeKeyContent` | `agentbay:DescribeKeyContent` |

```json
{
  "Action": [
    "agentbay:DescribeKeyContent"
  ]
}
```
```

---

## 执行步骤

1. 修改 `docs/en/apikey.md`：
   - 在 `## Terminology` 与 `## Commands` 之间插入 `## Required Permissions` 章节
   - 在 7 个子命令各自的 Flags 表格之后、`---` 之前，插入对应的 `**Involved APIs:**` 表格 + JSON Action 数组

2. 修改 `docs/zh/apikey.md`：
   - 同步添加中文版内容（章节标题、说明文字翻译为中文，接口名、权限字符串、JSON 保持英文原文）
   - 保持与英文版结构完全对应

---

## 验证方式

纯文档修改，无代码变更，无需运行测试。

验证步骤：
1. 检查两个文档的 Markdown 结构正确（标题层级、表格、代码块）
2. 确认每个子命令下的接口与代码中实际调用的接口一致（已在探索阶段核对）
3. 确认 JSON Policy 示例格式合法，Action 数组无遗漏
