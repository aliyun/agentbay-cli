[中文](../zh/skills.md) | **English**

# Skills Management — `agentbay skills`

Push local skills and inspect details by ID.

## Commands

### `skills push`

Push a local skill (directory or `.zip`) to the cloud. A directory must contain `SKILL.md` with `name` / `description` frontmatter; a directory is packed into a zip and uploaded.

```bash
agentbay skills push ./my-skill
agentbay skills push ./my-skill.zip
agentbay skills push ./my-skill --tag "tag1" --tag "tag2"
```

**Arguments:**

| Argument | Type   | Required | Description                            |
| -------- | ------ | -------- | -------------------------------------- |
| `<path>` | string | Yes      | Path to skill directory or `.zip` file |

**Flags:**

| Flag   | Short | Type        | Required | Description                                                                                   |
| ------ | ----- | ----------- | -------- | --------------------------------------------------------------------------------------------- |
| `--tag` |       | stringArray | No       | Tag name for the skill (can be specified multiple times, e.g. `--tag "tag1" --tag "tag2"`) |

**Notes:**

- Directory must contain `SKILL.md` with `name` and `description` in YAML frontmatter.
- Directory is automatically packed into a `.zip` before upload.
- When `--tag` is specified, the CLI first checks whether each tag already exists; missing tags are created automatically before the skill is uploaded.
- Tags are processed before obtaining the upload credential to avoid credential expiry during tag creation.

**Output:**

```
[STEP 1/4] Processing tags...
[INFO] All tags already exist.
[STEP 2/4] Getting upload credential...
[STEP 3/4] Uploading skill package...
[STEP 4/4] Creating skill...
[SUCCESS] Skill created successfully!
[RESULT] Skill ID: 35U2Ver2
```

> Without `--tag`, the step count is 3 (no tag processing step).

**Involved APIs:**

| Action                     | Required Permission                 |
| -------------------------- | ----------------------------------- |
| `ListTag`                  | `agentbay:ListTag`                  |
| `CreateTag`                | `agentbay:CreateTag`                |
| `GetMarketSkillCredential` | `agentbay:GetMarketSkillCredential` |
| `CreateMarketSkill`        | `agentbay:CreateMarketSkill`        |

> `ListTag` and `CreateTag` are only called when `--tag` is specified.

```json
{
  "Action": ["agentbay:ListTag", "agentbay:CreateTag", "agentbay:GetMarketSkillCredential", "agentbay:CreateMarketSkill"]
}
```

---

### `skills show`

Show skill details by ID.

```bash
agentbay skills show <skill-id>
```

**Arguments:**

| Argument     | Type   | Required | Description |
| ------------ | ------ | -------- | ----------- |
| `<skill-id>` | string | Yes      | Skill ID    |

**Output:**

```
SkillId:       <skill-id>
Name:          <skill-name>
Tags:          tag1, tag2
Description:
  <description text>
```

> `Tags` is only displayed when the skill has user-defined tenant tags.

**Involved APIs:**

| Action                      | Required Permission                  |
| --------------------------- | ------------------------------------ |
| `DescribeMarketSkillDetail` | `agentbay:DescribeMarketSkillDetail` |

```json
{
  "Action": ["agentbay:DescribeMarketSkillDetail"]
}
```

---

### `skills list` _(placeholder)_

Lists cloud skills. Backend list API is not yet available; this command currently acts as a placeholder.
