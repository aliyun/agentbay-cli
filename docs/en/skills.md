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
- **Duplicate name restriction:** `push` is a pure create operation; the platform does not allow duplicate skill names under the same user account. To update an existing skill's content, use `skills update`.

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
- **Skill name cannot be changed:** The `name` field in `SKILL.md` of the new file must match the original skill's name exactly. If they differ, the server will return an error.
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

### `skills list`

List cloud skills with pagination, supporting optional filters by name and tags.

```bash
agentbay skills list
agentbay skills list --page 2
agentbay skills list --size 20
agentbay skills list --name "find"
agentbay skills list --tag test --tag aliyun
agentbay skills list --name "find" --tag aliyun --page 1 --size 5
```

**Flags:**

| Flag       | Short | Type        | Required | Default | Description                                                                                                                                             |
| ---------- | ----- | ----------- | -------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--page`   |       | int         | No       | 1       | Page number                                                                                                                                             |
| `--size`   |       | int         | No       | 10      | Number of results per page                                                                                                                              |
| `--name`   |       | string      | No       | (none)  | Filter by skill name                                                                                                                                    |
| `--tag`    |       | stringArray | No       | (none)  | Filter by tag name (can be specified multiple times); multiple `--tag` values use **OR** logic — skills matching any of the specified tags are returned |
| `--output` | `-o`  | string      | No       | (none)  | Output format. Use `json` for machine-readable complete data (e.g. for AI/scripts)                                                                      |

**Output:**

Default table output (columns adapt to terminal width):

```
[INFO] ListMarketSkillByPage Request ID: A4E9C0A5-7BD3-1B1C-A3C5-D54F9472F3AE
[PAGE] Page 1 of 1 (Page Size: 10, Total: 6)

SKILL NAME                      SKILL ID                          STATUS                 TAGS                                       MODIFIED
------------------------------  --------------------------------  ----------------------  ------------------------------------------  ------------------------------
lxy-find-skills                 skill-04p87enx9u4moq5fi           VERIFY_PASSED          tag1, tag2                                  2026-05-26T02:37:59.000+00:00
stock-watcher                   skill-04p87lvcjt9o1o9uj           INIT                                                               2026-04-04T08:42:11.000+00:00
```

Use `--output json` (or `-o json`) for complete JSON output, suitable for AI/scripts:

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
      "tags": ["tag1", "tag2"],
      "icon": "https://...",
      "gmtModified": "2026-05-26T02:37:59.000+00:00",
      "gmtCreate": "2026-05-22T08:23:04.000+00:00"
    }
  ]
}
```

When there are more pages, a tip is shown at the end:

```
[TIP] Use --page 2 to view the next page.
```

**Involved APIs:**

| Action                  | Required Permission              |
| ----------------------- | -------------------------------- |
| `ListMarketSkillByPage` | `agentbay:ListMarketSkillByPage` |

```json
{
  "Action": ["agentbay:ListMarketSkillByPage"]
}
```

---

### `skills delete`

Permanently delete a skill from the cloud.

By default, the command fetches the skill details and displays them before prompting for confirmation. With `--yes`, both the detail lookup and confirmation prompt are skipped and the deletion is performed directly — suitable for scripts/CI.

The skill ID can be passed as a positional argument or via the `--skill-id` flag.

```bash
# Delete using positional argument (interactive, shows skill info and prompts for confirmation)
agentbay skills delete skill-xxxxxxxxxxxxxxxx

# Delete using positional argument, skip confirmation (scripts/CI)
agentbay skills delete skill-xxxxxxxxxxxxxxxx --yes
agentbay skills delete skill-xxxxxxxxxxxxxxxx -y

# Delete using named flag (compatible)
agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx
agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx --yes
agentbay skills delete --skill-id skill-xxxxxxxxxxxxxxxx -y
```

**Flags:**

| Flag         | Short | Type   | Required | Default | Description                                                          |
| ------------ | ----- | ------ | -------- | ------- | -------------------------------------------------------------------- |
| `--skill-id` |       | string | No\*     | (none)  | Skill ID to delete (alternative to positional argument)              |
| `--yes`      | `-y`  | bool   | No       | `false` | Skip detail lookup and confirmation prompt (for non-interactive use) |

> \* The skill ID must be provided either as a positional argument or via `--skill-id`.

**Output (interactive mode):**

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

**Output (`--yes` mode):**

```
[INFO] --yes specified, skipping skill detail lookup.
[INFO] DeleteMarketSkill Request ID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

[SUCCESS] Skill has been deleted.
  SkillId: skill-xxxxxxxxxxxxxxxx
```

**Involved APIs:**

| Action                      | Required Permission                  |
| --------------------------- | ------------------------------------ |
| `DescribeMarketSkillDetail` | `agentbay:DescribeMarketSkillDetail` |
| `DeleteMarketSkill`         | `agentbay:DeleteMarketSkill`         |

```json
{
  "Action": ["agentbay:DescribeMarketSkillDetail", "agentbay:DeleteMarketSkill"]
}
```

> **Note:** When using `--yes`, only the `agentbay:DeleteMarketSkill` permission is required.
