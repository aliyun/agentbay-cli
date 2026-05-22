[中文](../zh/skills.md) | **English**

# Skills Management — `agentbay skills`

Push local skills and inspect details by ID.

## Commands

### `skills push`

Push a local skill (directory or `.zip`) to the cloud. A directory must contain `SKILL.md` with `name` / `description` frontmatter; a directory is packed into a zip and uploaded.

```bash
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
```

**Arguments:**

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<path>` | string | Yes | Path to skill directory or `.zip` file |

**Notes:**

- Directory must contain `SKILL.md` with `name` and `description` in YAML frontmatter.
- Directory is automatically packed into a `.zip` before upload.

**Output:**

```
[SUCCESS] Skill created successfully!
[RESULT] Skill ID: 35U2Ver2
```

**Involved APIs:**

| Action | Required Permission |
|---|---|
| `GetMarketSkillCredential` | `agentbay:GetMarketSkillCredential` |
| `CreateMarketSkill` | `agentbay:CreateMarketSkill` |

```json
{
  "Action": [
    "agentbay:GetMarketSkillCredential",
    "agentbay:CreateMarketSkill"
  ]
}
```

---

### `skills show`

Show skill details by ID.

```bash
agentbay skills show <skill-id>
```

**Arguments:**

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<skill-id>` | string | Yes | Skill ID |

**Involved APIs:**

| Action | Required Permission |
|---|---|
| `DescribeMarketSkillDetail` | `agentbay:DescribeMarketSkillDetail` |

```json
{
  "Action": [
    "agentbay:DescribeMarketSkillDetail"
  ]
}
```

---

### `skills list` _(placeholder)_

Lists cloud skills. Backend list API is not yet available; this command currently acts as a placeholder.
