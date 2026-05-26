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
agentbay skills push ./my-skill --icon 'https://example.com/icon.png'
agentbay skills push ./my-skill --tag "tag1" --icon 'https://example.com/icon.png'
```

**Arguments:**

| Argument | Type   | Required | Description                            |
| -------- | ------ | -------- | -------------------------------------- |
| `<path>` | string | Yes      | Path to skill directory or `.zip` file |

**Flags:**

| Flag     | Short | Type        | Required | Default                   | Description                                                                                  |
| -------- | ----- | ----------- | -------- | ------------------------- | -------------------------------------------------------------------------------------------- |
| `--tag`  |       | stringArray | No       | (none)                    | Tag name for the skill (can be specified multiple times, e.g. `--tag "tag1" --tag "tag2"`)   |
| `--icon` |       | string      | No       | AgentBay default icon URL | Icon for the skill (URL or identifier). If not specified, the default AgentBay icon is used. |

**Notes:**

- Directory must contain `SKILL.md` with `name` and `description` in YAML frontmatter.
- Directory is automatically packed into a `.zip` before upload.
- When `--tag` is specified, the CLI first checks whether each tag already exists; missing tags are created automatically before the skill is uploaded.
- Tags are processed before obtaining the upload credential to avoid credential expiry during tag creation.
- If `--icon` is not specified, the default AgentBay icon is used automatically.
- **Shell quoting for `--icon`:** If the icon URL contains `!!` (e.g. Alibaba CDN URLs like `...!!6000000005528...`), wrap it in **single quotes** to prevent zsh history expansion: `--icon 'https://...'`.

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

Update an existing skill by ID. Upload a new skill package and optionally update tags or icon.

```bash
agentbay skills update --skill-id <id> --file ./my-skill
agentbay skills update --skill-id <id> --file ./my-skill.zip --tag "tag1" --tag "tag2"
agentbay skills update --skill-id <id> --file ./my-skill --icon 'https://example.com/icon.png'
```

**Flags:**

| Flag         | Type        | Required | Description                                                                                |
| ------------ | ----------- | -------- | ------------------------------------------------------------------------------------------ |
| `--skill-id` | string      | Yes      | Skill ID to update                                                                         |
| `--file`     | string      | Yes      | Path to skill directory or `.zip` file                                                     |
| `--tag`      | stringArray | No       | Tag name for the skill (can be specified multiple times, e.g. `--tag "tag1" --tag "tag2"`) |
| `--icon`     | string      | No       | Icon for the skill (e.g. URL or identifier)                                                |

**Notes:**

- When `--file` is a directory, it must contain `SKILL.md` with `name` and `description` in YAML frontmatter.
- When `--file` is a directory, it is automatically packed into a `.zip` before upload.
- When `--tag` is specified, the CLI first checks whether each tag already exists; missing tags are created automatically.
- Tags are processed before obtaining the upload credential to avoid credential expiry.
- **Shell quoting for `--icon`:** If the icon URL contains `!!` (e.g. Alibaba CDN URLs like `...!!6000000005528...`), wrap it in **single quotes** to prevent zsh history expansion: `--icon 'https://...'`. Double quotes or no quotes will cause zsh to expand `!!` into the previous command, resulting in a parse error.

**Output:**

```
[STEP 1/3] Getting upload credential...
[STEP 2/3] Uploading skill zip...
[STEP 3/3] Updating skill...
[INFO] UpdateMarketSkill RequestId: xxx
[SUCCESS] Skill updated successfully!
[RESULT] Skill ID: 35U2Ver2
```

> With `--tag`, a tag processing step is added before the credential step.

**Involved APIs:**

| Action                     | Required Permission                 |
| -------------------------- | ----------------------------------- |
| `ListTag`                  | `agentbay:ListTag`                  |
| `CreateTag`                | `agentbay:CreateTag`                |
| `GetMarketSkillCredential` | `agentbay:GetMarketSkillCredential` |
| `UpdateMarketSkill`        | `agentbay:UpdateMarketSkill`        |

> `ListTag` and `CreateTag` are only called when `--tag` is specified. `GetMarketSkillCredential` is only called when `--file` is specified.

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
