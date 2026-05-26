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
- **同名限制：** `push` 是纯创建操作，服务端不允许同一用户下存在重名技能。若需更新已有技能的内容，请使用 `skills update` 命令。

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
- **技能名不允许修改：** 上传的新文件中 `SKILL.md` 的 `name` 字段必须与原技能的名称保持一致。若两者不同，服务端会返回错误。
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

### `skills list`

分页查询云端技能列表，支持按名称和标签筛选。

```bash
agentbay skills list
agentbay skills list --page 2
agentbay skills list --size 20
agentbay skills list --name "find"
agentbay skills list --tag test --tag aliyun
agentbay skills list --name "find" --tag aliyun --page 1 --size 5
```

**Flags：**

| 参数       | 短参数 | 类型        | 必填 | 默认值 | 说明                                                                                     |
| ---------- | ------ | ----------- | ---- | ------ | ---------------------------------------------------------------------------------------- |
| `--page`   |        | int         | 否   | 1      | 页码                                                                                     |
| `--size`   |        | int         | 否   | 10     | 每页条数                                                                                 |
| `--name`   |        | string      | 否   | （无） | 按技能名称筛选                                                                           |
| `--tag`    |        | stringArray | 否   | （无） | 按标签筛选（可多次指定）；多个标签之间为**或（OR）**关系，返回包含任意一个指定标签的技能 |
| `--output` | `-o`   | string      | 否   | （无） | 输出格式。使用 `json` 获取机器可读的完整数据（适合 AI/脚本使用）                         |

**输出：**

默认表格输出（根据终端宽度自适应列）：

```
[INFO] ListMarketSkillByPage Request ID: A4E9C0A5-7BD3-1B1C-A3C5-D54F9472F3AE
[PAGE] Page 1 of 1 (Page Size: 10, Total: 6)

SKILL NAME                      SKILL ID                          STATUS                 TAGS                                       MODIFIED
------------------------------  --------------------------------  ----------------------  ------------------------------------------  ------------------------------
lxy-find-skills                 skill-04p87enx9u4moq5fi           VERIFY_PASSED          哈哈, 阿里云, lxy, test-2, test-1           2026-05-26T02:37:59.000+00:00
stock-watcher                   skill-04p87lvcjt9o1o9uj           INIT                                                               2026-04-04T08:42:11.000+00:00
```

使用 `--output json`（或 `-o json`）输出完整 JSON，适合 AI/脚本使用：

```bash
agentbay skills list -o json
```

```json
{
  "totalCount": 2,
  "totalPage": 1,
  "pageSize": 10,
  "pageNumber": 1,
  "result": [
    {
      "skillId": "skill-04p87enx9u4moq5fi",
      "skillName": "lxy-find-skills",
      "description": "...",
      "status": "VERIFY_PASSED",
      "tags": ["哈哈", "阿里云"],
      "icon": "https://...",
      "gmtModified": "2026-05-26T02:37:59.000+00:00",
      "gmtCreate": "2026-05-22T08:23:04.000+00:00"
    }
  ]
}
```

当结果有多页时，末尾会显示下一页提示：

```
[TIP] Use --page 2 to view the next page.
```

**涉及接口：**

| Action                  | 所需权限                         |
| ----------------------- | -------------------------------- |
| `ListMarketSkillByPage` | `agentbay:ListMarketSkillByPage` |

```json
{
  "Action": ["agentbay:ListMarketSkillByPage"]
}
```

---

### `skills delete`

永久删除云端技能。

默认情况下，命令会先查询技能详情并展示，再提示确认是否删除。指定 `--yes` 时跳过详情查询和确认提示，直接执行删除，适合脚本/CI 场景。

```bash
# 交互式删除（展示技能信息并确认）
agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx

# 跳过详情查询和确认，直接删除（脚本/CI）
agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx --yes
agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx -y
```

**Flags：**

| 参数         | 短参数 | 类型   | 必填 | 默认值  | 说明                                            |
| ------------ | ------ | ------ | ---- | ------- | ----------------------------------------------- |
| `--skill-id` |        | string | 是   | （无）  | 要删除的技能 ID                                 |
| `--yes`      | `-y`   | bool   | 否   | `false` | 跳过详情查询和确认提示（适合非交互式/脚本场景） |

**输出（交互模式）：**

```
[STEP 1/2] Fetching skill details...
[INFO] DescribeMarketSkillDetail Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
  SkillId: skill-xxxxxxxxxxxxxxxx
  Name:    my-skill

Are you sure you want to permanently delete this skill? [y/N]: y
[INFO] DeleteMarketSkill Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

[SUCCESS] Skill has been deleted.
  SkillId: skill-xxxxxxxxxxxxxxxx
```

**输出（`--yes` 模式）：**

```
[INFO] --yes specified, skipping skill detail lookup.
[INFO] DeleteMarketSkill Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

[SUCCESS] Skill has been deleted.
  SkillId: skill-xxxxxxxxxxxxxxxx
```

**涉及接口：**

| Action                      | 所需权限                             |
| --------------------------- | ------------------------------------ |
| `DescribeMarketSkillDetail` | `agentbay:DescribeMarketSkillDetail` |
| `DeleteMarketSkill`         | `agentbay:DeleteMarketSkill`         |

```json
{
  "Action": ["agentbay:DescribeMarketSkillDetail", "agentbay:DeleteMarketSkill"]
}
```

> **注意：** 使用 `--yes` 时，仅需 `agentbay:DeleteMarketSkill` 权限。
