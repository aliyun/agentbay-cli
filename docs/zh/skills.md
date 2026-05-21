[English](../en/skills.md) | **中文**

# 技能管理 — `agentbay skills`

推送本地技能包，按 ID 查看技能详情。

## 命令

### `skills push`

推送本地技能（目录或 `.zip`）到云端。目录形式必须包含带 `name` / `description` frontmatter 的 `SKILL.md`，目录会被打包为 zip 后上传。

```bash
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `<path>` | string | 是 | 技能目录或 `.zip` 文件路径 |

**注意事项：**

- 目录必须包含 `SKILL.md`，且 YAML frontmatter 中有 `name` 和 `description`。
- 目录会自动打包为 `.zip` 后上传。

**输出：**

```
[SUCCESS] Skill created successfully!
[RESULT] Skill ID: 35U2Ver2
```

---

### `skills show`

按 ID 查看技能详情。

```bash
agentbay skills show <skill-id>
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `<skill-id>` | string | 是 | 技能 ID |

---

### `skills list`（占位）

列出云端技能。后端 list 接口尚未提供，该命令目前为占位实现。
