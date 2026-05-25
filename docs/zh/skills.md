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
```

**参数：**

| 参数     | 类型   | 必填 | 说明                       |
| -------- | ------ | ---- | -------------------------- |
| `<path>` | string | 是   | 技能目录或 `.zip` 文件路径 |

**Flags：**

| 参数    | 类型        | 必填 | 说明                                                         |
| ------- | ----------- | ---- | ------------------------------------------------------------ |
| `--tag` | stringArray | 否   | 技能标签名称（可多次指定，如 `--tag "标签1" --tag "标签2"`） |

**注意事项：**

- 目录必须包含 `SKILL.md`，且 YAML frontmatter 中有 `name` 和 `description`。
- 目录会自动打包为 `.zip` 后上传。
- 指定 `--tag` 时，CLI 会先检查每个标签是否已存在，不存在的标签会自动创建后再上传技能。
- 标签处理在获取上传凭证之前执行，以避免标签创建过程中凭证过期。

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
