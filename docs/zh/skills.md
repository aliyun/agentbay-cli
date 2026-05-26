[English](../en/skills.md) | **中文**

# 技能管理 — `agentbay skills`

推送本地技能包，按 ID 查看技能详情。

## 命令

### `skills push`

推送本地技能（目录或 `.zip`）到云端。目录形式必须包含带 `name` / `description` frontmatter 的 `SKILL.md`，目录会被打包为 zip 后上传。

```bash
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
agentbay skills push ./my-skill --tag "标签1" --tag "标签2"
agentbay skills push ./my-skill --icon 'https://example.com/icon.png'
agentbay skills push ./my-skill --tag "标签1" --icon 'https://example.com/icon.png'
```

**参数：**

| 参数     | 类型   | 必填 | 说明                       |
| -------- | ------ | ---- | -------------------------- |
| `<path>` | string | 是   | 技能目录或 `.zip` 文件路径 |

**Flags：**

| 参数     | 类型        | 必填 | 默认値                | 说明                                                         |
| -------- | ----------- | ---- | --------------------- | ------------------------------------------------------------ |
| `--tag`  | stringArray | 否   | （无）                | 技能标签名称（可多次指定，如 `--tag "标签1" --tag "标签2"`） |
| `--icon` | string      | 否   | AgentBay 默认图标 URL | 技能图标（URL 或标识），不传则自动使用默认图标               |

**注意事项：**

- 目录必须包含 `SKILL.md`，且 YAML frontmatter 中有 `name` 和 `description`。
- 目录会自动打包为 `.zip` 后上传。
- 指定 `--tag` 时，CLI 会先检查每个标签是否已存在，不存在的标签会自动创建后再上传技能。
- 标签处理在获取上传凭证之前执行，以避免标签创建过程中凭证过期。
- 不指定 `--icon` 时，会自动使用默认的 AgentBay 图标。
- **`--icon` 的 Shell 引号问题：** 若图标 URL 中含有 `!!`（例如阿里云 CDN URL 中常见的 `...!!6000000005528...`），请用**单引号**包裹，以防止 zsh 将 `!!` 展开为上一条命令：`--icon 'https://...'`。

**输出：**

```
[STEP 1/4] Processing tags...
[INFO] All tags already exist.
[STEP 2/4] Getting upload credential...
[STEP 3/4] Uploading skill package...
[STEP 4/4] Creating skill...
[SUCCESS] Skill created successfully!
[RESULT] Skill ID: 35U2Ver2
```

> 不指定 `--tag` 时，步骤数为 3（无标签处理步骤）。

**涉及接口：**

| Action                     | 所需权限                            |
| -------------------------- | ----------------------------------- |
| `ListTag`                  | `agentbay:ListTag`                  |
| `CreateTag`                | `agentbay:CreateTag`                |
| `GetMarketSkillCredential` | `agentbay:GetMarketSkillCredential` |
| `CreateMarketSkill`        | `agentbay:CreateMarketSkill`        |

> `ListTag` 和 `CreateTag` 仅在指定 `--tag` 时调用。

```json
{
  "Action": [
    "agentbay:ListTag",
    "agentbay:CreateTag",
    "agentbay:GetMarketSkillCredential",
    "agentbay:CreateMarketSkill"
  ]
}
```

---

### `skills update`

按 ID 更新已有技能。上传新的技能包，并可选地更新标签或设置图标。

```bash
agentbay skills update --skill-id <id> --file ./my-skill
agentbay skills update --skill-id <id> --file ./my-skill.zip --tag "标签1" --tag "标签2"
agentbay skills update --skill-id <id> --file ./my-skill --icon 'https://example.com/icon.png'
```

**Flags：**

| 参数         | 类型        | 必填 | 说明                                                         |
| ------------ | ----------- | ---- | ------------------------------------------------------------ |
| `--skill-id` | string      | 是   | 要更新的技能 ID                                              |
| `--file`     | string      | 是   | 技能目录或 `.zip` 文件路径                                   |
| `--tag`      | stringArray | 否   | 技能标签名称（可多次指定，如 `--tag "标签1" --tag "标签2"`） |
| `--icon`     | string      | 否   | 技能图标（如 URL 或标识）                                    |

**注意事项：**

- `--file` 为目录时，必须包含带 `name` / `description` frontmatter 的 `SKILL.md`。
- `--file` 为目录时，会自动打包为 `.zip` 后上传。
- 指定 `--tag` 时，CLI 会先检查每个标签是否已存在，不存在的标签会自动创建。
- 标签处理在获取上传凭证之前执行，以避免标签创建过程中凭证过期。
- **`--icon` 的 Shell 引号问题：** 若图标 URL 中含有 `!!`（例如阿里云 CDN URL 中常见的 `...!!6000000005528...`），请用**单引号**包裹，以防止 zsh 将 `!!` 展开为上一条命令：`--icon 'https://...'`。使用双引号或不加引号会导致 zsh 历史扩展，造成命令解析错误。

**输出：**

```
[STEP 1/3] Getting upload credential...
[STEP 2/3] Uploading skill zip...
[STEP 3/3] Updating skill...
[INFO] UpdateMarketSkill RequestId: xxx
[SUCCESS] Skill updated successfully!
[RESULT] Skill ID: 35U2Ver2
```

> 指定 `--tag` 时会在凭证获取步骤之前增加标签处理步骤。

**涉及接口：**

| Action                     | 所需权限                            |
| -------------------------- | ----------------------------------- |
| `ListTag`                  | `agentbay:ListTag`                  |
| `CreateTag`                | `agentbay:CreateTag`                |
| `GetMarketSkillCredential` | `agentbay:GetMarketSkillCredential` |
| `UpdateMarketSkill`        | `agentbay:UpdateMarketSkill`        |

> `ListTag` 和 `CreateTag` 仅在指定 `--tag` 时调用。`GetMarketSkillCredential` 仅在指定 `--file` 时调用。

```json
{
  "Action": [
    "agentbay:ListTag",
    "agentbay:CreateTag",
    "agentbay:GetMarketSkillCredential",
    "agentbay:UpdateMarketSkill"
  ]
}
```

---

### `skills show`

按 ID 查看技能详情。

```bash
agentbay skills show <skill-id>
```

**参数：**

| 参数         | 类型   | 必填 | 说明    |
| ------------ | ------ | ---- | ------- |
| `<skill-id>` | string | 是   | 技能 ID |

**输出：**

```
SkillId:       <技能ID>
Name:          <技能名称>
Tags:          标签1, 标签2
Description:
  <描述文本>
```

> `Tags` 仅在技能存在用户自定义租户标签时显示。

**涉及接口：**

| Action                      | 所需权限                             |
| --------------------------- | ------------------------------------ |
| `DescribeMarketSkillDetail` | `agentbay:DescribeMarketSkillDetail` |

```json
{
  "Action": ["agentbay:DescribeMarketSkillDetail"]
}
```

---

### `skills list`（占位）

列出云端技能。后端 list 接口尚未提供，该命令目前为占位实现。
