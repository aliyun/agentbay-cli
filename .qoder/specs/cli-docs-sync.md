# AgentBay CLI README & docs 治理方案

## Context

当前项目的文档存在三个结构性问题：
1. **README 臃肿**：README.md 和 README.zh-CN.md 各 502 行，Command Reference 占了 300+ 行，内容与 docs/USER_GUIDE.md 大量重复
2. **docs/ 组织混乱**：用户文档和内部文档混在一起，无双语目录，文件命名不规范（`LINUX&MAC_INSTALL.md` 含特殊字符）
3. **CHANGELOG 仅英文**：无中文版本，变更未按命令组分类，Release Notes 缺少结构化展示

目标：参考 agentrun-cli 的 docs/en+zh 目录结构 + ant-design 的双语 Release Notes 格式，重建文档体系，并将文档同步规则固化到开发流程中。

---

## Phase 1: 创建 docs/ 新目录结构

**目标路径**：
```
docs/
├── en/                          # 英文文档
│   ├── README.md                # 索引页，链接到各命令组文档
│   ├── installation.md          # 合并 WINDOWS_INSTALL + LINUX&MAC_INSTALL
│   ├── authentication.md        # 认证 + 环境变量（从 README 搬出）
│   ├── core.md                  # version, login, logout
│   ├── image.md                 # 所有 image 子命令
│   ├── apikey.md                # 所有 apikey 子命令
│   ├── network.md               # 所有 network 子命令
│   ├── skills.md                # 所有 skills 子命令
│   ├── docker.md                # 所有 docker 子命令
│   └── faq.md                   # FAQ（从 USER_GUIDE 提取）
├── zh/                          # 中文文档（结构与 en/ 完全镜像）
│   ├── README.md
│   ├── installation.md
│   ├── authentication.md
│   ├── core.md
│   ├── image.md
│   ├── apikey.md
│   ├── network.md
│   ├── skills.md
│   ├── docker.md
│   └── faq.md
└── internal/                    # 内部文档（不对客）
    ├── cli-openapi-actions.md   # 需补齐 image delete/status + apikey/network 映射
    ├── test-international-prod.md
    └── oss-skill-vs-image-upload.md
```

**操作**：
1. 创建 `docs/en/`、`docs/zh/`、`docs/internal/` 目录
2. 每个文件顶部加双语切换链接：`[中文](../../zh/xxx.md) | **English**` / `[English](../../en/xxx.md) | **中文**`

## Phase 2: 填充 docs/en/ 文档

内容来源映射：

| 目标文件 | 内容来源 |
|---------|---------|
| `en/README.md` | 新建索引页 |
| `en/installation.md` | 合并 `WINDOWS_INSTALL.md` + `LINUX&MAC_INSTALL.md` |
| `en/authentication.md` | README 的 Authentication + Environment Variables 章节 + USER_GUIDE 的 Environment Switching |
| `en/core.md` | README 的 Core Commands + USER_GUIDE 的 Login/Logout |
| `en/image.md` | README 的 Image Management + USER_GUIDE 的 Image 相关章节(4-10) |
| `en/apikey.md` | README 的 API Key Management |
| `en/network.md` | README 的 Network Management |
| `en/skills.md` | README 的 Skills Management + USER_GUIDE 的 Skills 章节 |
| `en/docker.md` | README 的 Docker Operations |
| `en/faq.md` | USER_GUIDE 的 FAQ 章节 |

每个命令组文档的统一结构：
```markdown
[中文](../../zh/xxx.md) | **English**

# <Group Name> — `agentbay <command>`

<一段概述>

## Commands

### `agentbay <command> <subcommand>`

<描述>

**Usage:**
```bash
agentbay <command> <subcommand> [flags]
```

**Examples:**
```bash
<示例>
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--flag` | `-f` | string | Yes | 说明 |

**Output:**
```
<输出示例>
```

**Notes:**
- <注意事项>
```

## Phase 3: 填充 docs/zh/ 文档

- 与 docs/en/ 结构完全一致，内容翻译为中文
- 命令示例保留英文（`agentbay image list` 不翻译）
- 可从 README.zh-CN.md 中直接提取对应中文内容

## Phase 4: 精简 README

将 README.md 从 ~500 行缩减到 ~150 行：

**保留章节**（适当精简）：
- Overview（2-3 句）
- Installation（3 行代码块 + 链接到 `docs/en/installation.md`）
- Authentication（3 行代码块 + 链接到 `docs/en/authentication.md`）
- Command Overview（**新表格**，替代原 Command Reference）

**Command Overview 表格**（新增，替代 300 行的 Command Reference）：

| Group | Commands | Description | Details |
|-------|----------|-------------|---------|
| Core | `version`, `login`, `logout` | Version & auth | [→](docs/en/core.md) |
| Image | `list`, `create`, `activate`, ... | Image lifecycle | [→](docs/en/image.md) |
| API Key | `create`, `enable`, `disable`, ... | Key management | [→](docs/en/apikey.md) |
| Network | `package list` | Network config | [→](docs/en/network.md) |
| Skills | `push`, `show` | Skill management | [→](docs/en/skills.md) |
| Docker | `login`, `tag`, `push` | Docker registry | [→](docs/en/docker.md) |

**移除章节**（已迁移到 docs/）：
- 完整 Environment Variables → `docs/en/authentication.md`
- 完整 Command Reference（各子命令详情）→ 各命令组文档
- 详细安装步骤 → `docs/en/installation.md`

同步更新 `README.zh-CN.md`，链接指向 `docs/zh/`。

## Phase 5: 处理内部文档

> 经过与当前代码逐一对比验证后的决策：

| 原文件 | 状态 | 操作 | 原因 |
|--------|------|------|------|
| `docs/cli-openapi-actions.md` | 部分过时 | 移至 `docs/internal/` 并**更新** | 核心结构有价值但缺少 image delete/status 和 apikey/network 的 Action 映射，需补齐 |
| `docs/testing-new-features.md` | **严重过时** | **删除** | 缺少 15+ 个新命令的回归步骤（apikey 全系列、image delete/set-max-session/warmup-status/create-from-template、network、docker tag/push），修补工作量接近重写，不如按需重建 |
| `docs/test-international-prod.md` | 准确 | 移至 `docs/internal/` | OAuth Client ID、endpoint、测试流程均与代码一致 |
| `docs/oss-skill-vs-image-upload.md` | 准确 | 移至 `docs/internal/` | 后端未改动，分析仍有效 |

**cli-openapi-actions.md 需补齐的内容**：
- `image delete` → `DeleteMcpImage`
- `image status` → `GetMcpImageInfo`（标注与 activate/set-max-session/deactivate 共用）
- `image create-from-template` 参数补充（确认是否需要 SourceImageId 等）
- `apikey create` → `CreateApiKey`
- `apikey enable/disable` → `ModifyMcpApiKeyConfig`（Action=EnableMcpApiKey / DisableMcpApiKey）
- `apikey delete` → `DeleteApiKey`
- `apikey list` → `DescribeMcpApiKey`
- `apikey concurrency set` → `ModifyMcpApiKeyConfig`（Action=SetMcpApiKeyConcurrency）
- `network package list` → `DescribeNetworkPackages`
- Action 总数从 15 更新到实际数量

**删除的文件**（内容已迁移到新结构）：
- `docs/USER_GUIDE.md`
- `docs/WINDOWS_INSTALL.md`
- `docs/LINUX&MAC_INSTALL.md`
- `docs/testing-new-features.md`（严重过时，缺 15+ 新命令）
- `docs/skills-cli-test-2026-03-14.md`（过时）
- `docs/plans/` 目录（历史规划）

## Phase 6: CHANGELOG 双语化

### 格式规范（参考 ant-design）

每个版本条目采用「英文 + `* * *` 分隔线 + 中文」格式，按命令组分类：

```markdown
## [0.3.0] - 2026-05-21

### 🚀 Features

- **apikey**: Add `apikey delete` command with multi-step confirmation
- **image**: Add `image warmup-status` command

### 🐞 Bug Fixes

- **client**: Tolerate stringified HttpStatusCode in response parsing

* * *

### 🚀 新功能

- **apikey**: 新增 `apikey delete` 命令，支持多步骤确认
- **image**: 新增 `image warmup-status` 命令

### 🐞 问题修复

- **client**: 兼容响应中 HttpStatusCode 字段为字符串类型的情况
```

### 实施方式

1. **回填所有历史版本的中文翻译**（v0.1.0–v0.2.8）
   - v0.2.5–v0.2.8：逐条精确翻译
   - v0.1.0–v0.2.4：概要式翻译（合并同类条目，提炼核心变更，不逐条翻译低质量 commit message）
2. 在 `cliff.toml` 的 body 模板末尾添加 `* * *` 分隔线和中文 placeholder：
   ```markdown
   * * *

   <!-- 中文翻译待补充 / Add Chinese translation before release -->
   ```
3. 发布前由开发者手动补充中文翻译

### README 中添加 CHANGELOG 入口

在 README 末尾 License 前添加：

```markdown
## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.
```

中文版：

```markdown
## 更新日志

查看 [CHANGELOG.md](CHANGELOG.md) 了解版本更新记录。
```

## Phase 7: 更新 cliff.toml

**修改文件**：`cliff.toml`

1. **body 模板**：添加 `* * *` 分隔线和中文 placeholder
2. **commit_parsers**：group 名称加 emoji 前缀
3. **新增 scope_parsers**：规范化 scope 到命令组名

```toml
# body 模板关键变更
body = """
{% if version %}\
## [{{ version | trim_start_matches(pat="v") }}] - {{ timestamp | date(format="%Y-%m-%d") }}
{% else %}\
## [Unreleased]
{% endif %}\
{% for group, commits in commits | group_by(attribute="group") %}
### {{ group | striptags | trim | upper_first }}
{% for commit in commits %}
- {% if commit.scope %}**{{ commit.scope }}**: {% endif %}\
{% if commit.breaking %}[**breaking**] {% endif %}\
{{ commit.message | upper_first }}\
{% endfor %}
{% endfor %}

* * *

<!-- 中文翻译待补充 / Add Chinese translation before release -->
"""

# group 名称加 emoji
{ message = "^feat", group = "<!-- 0 -->🚀 Features" },
{ message = "^fix", group = "<!-- 1 -->🐞 Bug Fixes" },
{ message = "^docs", group = "<!-- 2 -->📖 Documentation" },
{ message = "^refactor", group = "<!-- 3 -->🛠 Refactoring" },

# scope 规范化
scope_parsers = [
  { pattern = "api-key|apikey", scope = "apikey" },
  { pattern = "img|image", scope = "image" },
  { pattern = "net|network", scope = "network" },
  { pattern = "skill|skills", scope = "skills" },
  { pattern = "docker", scope = "docker" },
  { pattern = "core|auth|login|logout", scope = "core" },
  { pattern = "client|sdk", scope = "client" },
]
```

## Phase 7.5: 更新 CI Workflow — 中文翻译提醒

**修改文件**：`.github/workflows/homebrew.yml`

在 `Update CHANGELOG.md` 步骤之后，新增一个检查步骤（非阻塞，仅警告）：

```yaml
- name: Check Chinese translation in CHANGELOG
  run: |
    VERSION="${{ steps.setup-build-vars.outputs.version }}"
    # 检查当前版本区域是否仍有中文翻译占位符
    if git-cliff --tag "v$VERSION" --unreleased --strip header | grep -q "中文翻译待补充"; then
      echo "⚠️ WARNING: Chinese translation is missing in CHANGELOG for v$VERSION"
      echo "Please add Chinese translation below the '* * *' separator before final release."
      echo "This warning will not block the release, but Chinese translation is required."
    else
      echo "✅ Chinese translation found for v$VERSION"
    fi
```

在生成 GitHub Release Notes 的步骤中同样添加检查：

```yaml
- name: Check Chinese translation in Release Notes
  run: |
    if grep -q "中文翻译待补充" /tmp/release-changes.md; then
      echo "⚠️ WARNING: Release notes for v$VERSION are missing Chinese translation."
      echo "Consider adding Chinese section below '* * *' separator."
    fi
```

## Phase 8: 更新文档同步规则

### 8.1 更新 `.qoder/rules/development.md`

将现有的「新增或修改命令必须同步更新 README 和测试用例」章节扩展为：

```markdown
### 新增或修改命令必须同步更新文档和测试用例

**规则**: 每次新增或修改 CLI 命令时，必须同步完成以下工作：

1. **更新 `README.md` 和 `README.zh-CN.md`** — Command Overview 表格
2. **更新 `docs/en/<command-group>.md` 和 `docs/zh/<command-group>.md`** — 详细命令文档（参数、示例、输出）
3. **更新 `CHANGELOG.md`** — 发布前补充中文翻译（`* * *` 分隔线下方）
4. **同步更新对外文档** — 钉钉文档 / cli 使用手册
5. **编写/更新单元测试**

**检查清单**:
- [ ] README Command Overview 表格已更新
- [ ] docs/en/<group>.md 已更新
- [ ] docs/zh/<group>.md 已更新（与英文版保持一致）
- [ ] CHANGELOG 中文翻译已补充（如有新版本条目）
- [ ] 对外文档已同步
- [ ] 单元测试已编写或更新并通过
- [ ] go build -o agentbay . 构建通过
- [ ] go test ./... -count=1 全部通过
```

### 8.2 更新 `.qoder/skills/create-cli-command/SKILL.md`

扩展 Phase 5（文档生成），添加 docs/ 目录更新步骤：

```markdown
### Phase 5: 文档生成

1. **更新 docs/en/<command-group>.md** — 添加新命令的详细文档（用法、参数、示例、输出）
2. **更新 docs/zh/<command-group>.md** — 同步中文翻译
3. **更新 README.md 和 README.zh-CN.md** — Command Overview 表格
4. **更新 cli-analysis/ 对客文档** — 按需
```

### 8.3 触发机制说明

由于 development.md 的 Skill 自动装配规则已经将「新增/修改 CLI 命令」映射到 `create-cli-command` skill，因此：
- **直接对话**：自动装配规则触发 → 加载 create-cli-command skill → Phase 5 执行文档更新
- **Quest spec 模式**：Execute 阶段等同对话入口，规则照常生效
- **Execute Directly**：AI 仍须在动手前主动加载匹配的 skill

无需新建独立 skill，通过更新现有 development.md + create-cli-command skill 即可覆盖所有入口。

---

## 关键文件清单

| 文件 | 操作 |
|------|------|
| `docs/en/*.md` (10 个文件) | 新建 |
| `docs/zh/*.md` (10 个文件) | 新建 |
| `docs/internal/cli-openapi-actions.md` | 移入 + 更新（补齐缺失命令映射） |
| `docs/internal/test-international-prod.md` | 移入 |
| `docs/internal/oss-skill-vs-image-upload.md` | 移入 |
| `docs/USER_GUIDE.md` | 删除（内容已迁移） |
| `docs/WINDOWS_INSTALL.md` | 删除（内容已迁移） |
| `docs/LINUX&MAC_INSTALL.md` | 删除（内容已迁移） |
| `docs/testing-new-features.md` | 删除（严重过时） |
| `docs/skills-cli-test-2026-03-14.md` | 删除 |
| `docs/plans/` | 删除整个目录 |
| `README.md` | 精简（~500→~150 行） |
| `README.zh-CN.md` | 精简（同上） |
| `CHANGELOG.md` | 回填中文翻译 |
| `cliff.toml` | 更新模板、emoji、scope_parsers |
| `.github/workflows/homebrew.yml` | 新增中文翻译检查步骤（非阻塞） |
| `.qoder/rules/development.md` | 扩展文档同步规则 |
| `.qoder/skills/create-cli-command/SKILL.md` | 扩展 Phase 5 |

## 验证

1. **链接完整性**：检查 README → docs/ 的所有链接、docs/en/ ↔ docs/zh/ 的双语切换链接
2. **内容覆盖**：对比当前 README + USER_GUIDE 的命令列表，确保所有 24 个命令/子命令在新 docs 中都有对应文档
3. **中英文一致**：docs/en/ 和 docs/zh/ 的结构、命令覆盖范围完全一致
4. **构建验证**：`go build -o agentbay .` 通过（本次不涉及代码变更，但作为习惯验证）
5. **CHANGELOG 格式**：确认 `* * *` 分隔线、emoji 分类、scope 规范化生效
6. **cliff.toml 试运行**：`git-cliff --unreleased` 验证模板输出正确
