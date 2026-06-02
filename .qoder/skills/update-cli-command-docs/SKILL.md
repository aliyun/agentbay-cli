---
name: update-cli-command-docs
description: AgentBay CLI 命令文档同步流程（docs/、README、CHANGELOG readiness）
---

# Update CLI Command Docs

## 📋 职责

在 AgentBay CLI 需求开发完成后，同步更新以下文档，并确保变更可被发版阶段的双语 CHANGELOG 流程正确采集：

1. **docs/ 命令文档** — `docs/en/<group>.md` 和 `docs/zh/<group>.md`
2. **README Command Overview 表格** — `README.md` 和 `README.zh-CN.md`
3. **LLM-facing docs readiness** — 根据 `README.md` / `docs/en/**` 变更同步 `llms-full.txt`，并在文档结构变化时检查 `llms.txt`
4. **CHANGELOG readiness** — 校验 commit / scope / 文档摘要满足 `make release-prep` 采集要求；真正的版本段生成与翻译委托 `bilingual-changelog-release` skill

**本 skill 不涉及代码实现，仅负责文档层面的同步。**

对客文档（`cli-analysis/` 目录、钉钉文档）不在此 skill 范围内，需手动同步。

## 🎯 触发场景

当出现以下情况时触发：

- 新增 CLI 命令后需同步文档
- 修改已有命令的参数、默认值、输出格式后需同步文档
- 日常开发完成后需确认 LLM-facing docs readiness（`llms-full.txt` / `llms.txt` 是否需要同步）
- 日常开发完成后需确认 CHANGELOG readiness（commit type/scope/subject 是否能被 release-prep 正确采集）
- 用户明确提出"更新文档"、"同步文档"等需求
- 用户明确提出"更新 CHANGELOG"且上下文是日常命令文档同步；若目标是发版版本段或 GitHub Release notes，必须改用 `bilingual-changelog-release` skill
- `create-cli-command` skill 的 Phase 5 委托调用

## 🔗 前置约束

**本 skill 必须在代码变更完成后触发**，不得在代码尚未完成时提前生成文档。⚠️ **若涉及分支切换，必须先询问用户确认**，不得自动切换 feat 分支 —— 用户可能希望继续在当前分支开发。典型组合：

| 场景                           | 前置 skill                                      | 本 skill 时机             |
| ------------------------------ | ----------------------------------------------- | ------------------------- |
| 新增命令                       | `create-cli-command` Phase 1-4 完成后           | 替代其 Phase 5            |
| 修改命令                       | 直接修改代码后                                  | 代码修改完成即可          |
| 日常 PR 文档同步               | `feature-development-workflow` Phase 2-3 完成后 | 提交前                    |
| 发版 CHANGELOG / Release notes | `bilingual-changelog-release`                   | 不在本 skill 内生成版本段 |

## 🚀 执行步骤

### Phase 0: 变更分析

**目的**：识别本次变更影响了哪些命令组、哪些文档需要更新，避免遗漏或无谓修改。

1. **检查当前变更范围**

   ```bash
   # 最近一次提交的变更文件
   git diff --name-only HEAD~1
   # 工作区未提交的变更
   git diff --name-only
   ```

2. **识别受影响的命令组**

   根据变更的代码文件映射到命令组：

   | 代码文件                                            | 命令组  | docs 文件                                   |
   | --------------------------------------------------- | ------- | ------------------------------------------- |
   | `cmd/apikey*.go`, `cmd/concurrency.go`              | apikey  | `docs/en/apikey.md` / `docs/zh/apikey.md`   |
   | `cmd/image*.go`                                     | image   | `docs/en/image.md` / `docs/zh/image.md`     |
   | `cmd/network.go`                                    | network | `docs/en/network.md` / `docs/zh/network.md` |
   | `cmd/skills*.go`                                    | skills  | `docs/en/skills.md` / `docs/zh/skills.md`   |
   | `cmd/docker.go`                                     | docker  | `docs/en/docker.md` / `docs/zh/docker.md`   |
   | `cmd/login.go`, `cmd/logout.go`, `cmd/constants.go` | core    | `docs/en/core.md` / `docs/zh/core.md`       |

3. **判定文档更新类型**

   | 变更类型       | docs 更新             | README 更新           | CHANGELOG readiness                                    |
   | -------------- | --------------------- | --------------------- | ------------------------------------------------------ |
   | 新增命令       | ✅ 添加完整命令文档节 | ✅ 在表格中添加子命令 | ✅ 校验 commit 使用 `feat(<group>)`                    |
   | 新增参数       | ✅ 更新参数表格       | ✅ 如影响简述则更新   | ✅ 校验 commit 使用 `feat(<group>)` 或 `docs(<group>)` |
   | 修改参数默认值 | ✅ 更新默认值和说明   | ✅ 如影响简述则更新   | ✅ 校验 commit subject 能表达用户影响                  |
   | 修改输出格式   | ✅ 更新输出示例       | ❌ 通常不需要         | ✅ 校验 commit subject 能表达用户影响                  |
   | 修改命令行为   | ✅ 更新说明和注意事项 | ✅ 更新简述           | ✅ 校验 commit type/scope                              |
   | 仅内部重构     | ❌ 通常不需要         | ❌                    | 视情况                                                 |
   | 仅补发版翻译   | ❌                    | ❌                    | ❌ 改用 `bilingual-changelog-release`                  |

4. **判定 llms 更新类型**

   | 变更类型                                        | llms 动作                                                            |
   | ----------------------------------------------- | -------------------------------------------------------------------- |
   | 修改 `README.md`                                | 执行 `bash scripts/build-llms-full.sh`，同步 `llms-full.txt`         |
   | 修改 `docs/en/**`                               | 执行 `bash scripts/build-llms-full.sh`，同步 `llms-full.txt`         |
   | 新增 / 删除 / 重命名对外文档                    | 更新 `llms.txt` 导航链接；如涉及英文源文档，同步重建 `llms-full.txt` |
   | 仅修改 `docs/zh/**`                             | 通常不重建 `llms-full.txt`；若文档结构变化，检查 `llms.txt` 中文链接 |
   | 仅修改 `docs/internal/**` / 测试文档 / 脚本文档 | 不进入 llms 文档，通常无需同步                                       |

   CLI 命令变更通常会同步 `docs/en/<group>.md` 或 `README.md`，因此必须把 `llms-full.txt` 纳入本次文档同步范围。

5. **向用户确认更新范围**

   展示分析结果，例如：

   > 检测到以下变更需要文档同步：
   >
   > - **新增命令**: `apikey status` → 需更新 `docs/en/apikey.md`、`docs/zh/apikey.md`、README 表格、CHANGELOG readiness
   > - **修改参数**: `image create --os-type` 默认值变更 → 需更新 `docs/en/image.md`、`docs/zh/image.md`
   >
   > 是否按此范围执行文档更新？

---

### Phase 1: 更新 docs/ 命令文档

**原则**：中英文文档结构必须完全一致，内容互为翻译。先更新英文版，再同步中文版。

#### 1.1 读取现有文档结构

读取目标命令组的现有文档，理解当前格式和风格。

#### 1.2 新增命令 — 添加命令文档节

在对应命令组文档的 `## Commands` / `## 命令` 部分末尾追加新节。

**英文版模板**：

````markdown
---

### `<group> <subcommand>`

<一句话描述>

```bash
agentbay <group> <subcommand> [flags]
```
````

**Flags:**

| Flag      | Short | Type   | Required | Description |
| --------- | ----- | ------ | -------- | ----------- |
| `--param` | `-p`  | string | Yes      | 说明        |

**Notes:**

- 注意事项

````

**中文版模板**：

```markdown
---

### `<group> <subcommand>`

<一句话中文描述>

```bash
agentbay <group> <subcommand> [flags]
````

**参数：**

| 参数      | 短参数 | 类型   | 必填 | 说明     |
| --------- | ------ | ------ | ---- | -------- |
| `--param` | `-p`   | string | 是   | 中文说明 |

**注意事项：**

- 注意事项

````

**关键约束**：

- 命令示例保留英文（`agentbay apikey create` 不翻译为 `agentbay apikey 创建`）
- 表头英文版用 `Flag` / `Short` / `Type` / `Required` / `Description`，中文版用 `参数` / `短参数` / `类型` / `必填` / `说明`
- `Required` 列英文版用 `Yes`/`No`，中文版用 `是`/`否`
- 如果命令有破坏性操作，必须包含 `--yes` / `-y` 参数说明和注意事项
- 如果没有短参数（Short），中文版可省略"短参数"列（参考同文件已有命令的格式）
- 新增节前用 `---` 分隔线

#### 1.3 修改命令 — 更新已有命令文档节

定位到对应子命令的文档节，更新受影响的部分：

- **新增参数**：在参数表格中追加行
- **修改参数默认值**：更新 Type 列中的默认值标注和 Description/说明列
- **修改输出格式**：更新输出示例代码块
- **修改命令行为**：更新描述段落和注意事项

**必须注意**：修改时不要破坏同一节中未变更的内容。

#### 1.4 双语同步验证

更新完成后，逐项对比中英文文档：

- [ ] 英文版和中文版的命令数量一致
- [ ] 每个命令的参数数量和名称一致
- [ ] Required/必填 列的值一致
- [ ] 注意事项的条目数一致
- [ ] 双语切换链接（文件第一行）正确

---

### Phase 1.5: 更新 RAM 接口权限说明

**目的**：当新增或删除 OpenAPI 调用时，同步维护各命令文档中的「涉及接口」章节，以及 README 中的 RAM 权限汇总表，并向用户输出权限变更摘要供人工核查。

#### 1.5.1 识别接口变更

基于 Phase 0 的变更分析，确认哪些命令新增/删除/修改了 OpenAPI 调用：

- 读取相关 `cmd/*.go` 文件，统计调用的接口名（函数名即 Action）
- 每个 Action 对应的权限字符串：`agentbay:<ActionName>`
- 不统计本地 docker CLI 封装命令（`docker tag`、`docker push`）

#### 1.5.2 在 docs/ 中维护「涉及接口」章节

**位置**：每个子命令的参数表格之后（或注意事项之后）、`---` 分隔线之前。

**英文版模板**：

```markdown
**Involved APIs:**

| Action | Required Permission |
|---|---|
| `XxxAction` | `agentbay:XxxAction` |

```json
{
  "Action": [
    "agentbay:XxxAction"
  ]
}
````

````

**中文版模板**：

```markdown
**涉及接口：**

| Action | 所需权限 |
|---|---|
| `XxxAction` | `agentbay:XxxAction` |

```json
{
  "Action": [
    "agentbay:XxxAction"
  ]
}
````

````

**规则**：

- 一个命令涉及多个接口时，表格按调用顺序列出所有 Action，JSON 数组同样全部列出
- 无需区分分支条件，只要命令可能调用的接口均需列出
- 本地 CLI 封装命令（无 AgentBay API 调用）不加「涉及接口」表格，改用提示：
  - 英文：`> **Note**: This is a native docker CLI wrapper — no AgentBay API calls are made. No additional RAM permissions required.`
  - 中文：`> **注意**：此命令是本地 docker CLI 的封装命令，不调用任何 AgentBay OpenAPI 接口，无需配置额外的 RAM 权限。`
- `skills list`（占位命令）等尚无 API 调用的命令不加「涉及接口」章节

#### 1.5.3 更新 README RAM 权限汇总表

`README.md` 和 `README.zh-CN.md` 中的 `## RAM Permissions` / `## RAM 账号接口权限` 章节包含各命令组的权限汇总表和 Policy JSON 示例。

**更新规则**：

- 新增接口：在对应命令组的汇总表中追加行，并在 Policy JSON 的 `Action` 数组中追加权限字符串
- 删除接口：从表格和 Policy JSON 中移除对应行/条目（需先确认该接口在该命令组的其他命令中已无引用）
- 接口名变更：同步修改表格和 Policy JSON

**Policy JSON 格式参考**（`README.md`）：

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:XxxAction"
      ],
      "Resource": "*"
    }
  ]
}
````

#### 1.5.4 终端输出接口变更摘要

**⚠️ 执行本 skill 时，必须在终端输出以下格式的权限变更摘要**，供用户核查是否有接口遗漏或误加：

```
========================================
  RAM 接口权限变更摘要
========================================

受影响命令组：<group>
受影响命令：<group> <subcommand>

新增接口权限：
  + agentbay:XxxAction     （<group> <subcommand>）

删除接口权限：
  - agentbay:YyyAction     （<group> <subcommand>）

无变更：（如无新增/删除则显示此行）

已同步位置：
  - docs/en/<group>.md — <subcommand> 涉及接口章节
  - docs/zh/<group>.md — <subcommand> 涉及接口章节
  - README.md — RAM Permissions > <group> 命令组
  - README.zh-CN.md — RAM 账号接口权限 > <group> 命令组

请核查上述权限变更是否正确，如有遗漏请告知。
========================================
```

**输出时机**：在完成所有文档更新后，统一输出一次摘要。如果本次变更无 OpenAPI 调用变化（纯参数修改、文档格式调整等），则输出「无接口权限变更」。

---

### Phase 2: 更新 README Command Overview 表格

**原则**：README 表格只做概要展示，不展示详细参数。更新时保持表格格式不变。

#### 2.1 定位 Command Overview 表格

`README.md` 中搜索 `## Command Overview`。
`README.zh-CN.md` 中搜索 `## 命令概览`。

#### 2.2 更新表格

**新增子命令**：在对应命令组的 `Commands` 列中，按逻辑序追加子命令名（用反引号包裹，逗号分隔）。

例如新增 `apikey status` 命令后：

- 英文：`` `create`, `enable`, `disable`, `delete`, `list`, `status`, `concurrency set` ``
- 中文：`` `create`, `enable`, `disable`, `delete`, `list`, `status`, `concurrency set` ``

**修改子命令**：如果命令重命名，替换旧名为新名。

**关键约束**：

- README 表格中的命令名一律使用英文（中文版也不例外）
- 保持 `Details` 列的链接不变
- 两组 README 的表格内容必须完全一致（仅表头和列名不同）
- 不要改变表格列的顺序和宽度

---

### Phase 2.5: LLM-facing docs readiness

**原则**：`llms.txt` 和 `llms-full.txt` 是 AI 助手优先读取的对外文档入口。只要本 skill 修改了 `README.md` 或 `docs/en/**`，就必须同步生成 `llms-full.txt`。

#### 2.5.1 同步 `llms-full.txt`

触发条件：

- `README.md` 有变更
- 任意 `docs/en/**` 对外文档有变更
- 新增英文对外文档并已加入 `scripts/build-llms-full.sh` 的 `FILES` 数组

执行：

```bash
bash scripts/build-llms-full.sh
```

要求：

- 生成结果必须随同本次文档变更提交。
- 不要手写 `llms-full.txt`，它应由脚本生成。
- `llms-full.txt` 不应包含 `docs/internal/**` 的 Source 标记。

#### 2.5.2 检查 `llms.txt`

触发条件：

- 新增、删除、重命名对外文档
- 新增、删除、重命名命令组文档
- README 或文档入口结构发生变化，导致导航索引需要调整

要求：

- `llms.txt` 使用 GitHub `master` 分支绝对 URL。
- 英文和中文对外文档链接都要检查。
- 不收录 `docs/internal/**`、`test/**`、`.aoneci/**`、`scripts/README.md` 等内部/开发文档。

#### 2.5.3 readiness 输出模板

```text
LLM-facing docs readiness:
- llms-full.txt: Updated / Not needed（说明原因）
- llms.txt: Updated / Checked, no change needed（说明原因）
- 触发依据: README.md / docs/en/** / 文档结构变化 / 无
```

---

### Phase 3: CHANGELOG readiness（不生成版本段）

**原则**：日常命令文档同步阶段只确认本次变更能被发版阶段正确采集；不得再运行旧的 `git-cliff -o CHANGELOG.md` / `make changelog` 全量覆盖流程。真正的双语版本段生成、翻译、tag 发布和 GitHub Release notes 回灌，统一交给 `bilingual-changelog-release` skill。

#### 3.1 确认提交语义可被 git-cliff 采集

根据本次变更类型，为后续 commit/PR title 建议 Conventional Commits：

| 变更                | 推荐 commit / PR title                             |
| ------------------- | -------------------------------------------------- |
| 新增命令 / 新增参数 | `feat(<group>): add ...`                           |
| 修复用户可见问题    | `fix(<group>): ...`                                |
| 文档说明变更        | `docs(<group>): ...`                               |
| 不改变行为的重构    | `refactor(<group>): ...`                           |
| 不兼容变更          | `feat(<group>)!: ...` 或 footer `BREAKING CHANGE:` |

要求：

- `<group>` 优先使用 `apikey`、`image`、`network`、`skills`、`docker`、`core`、`client`
- subject 使用英文祈使句，简洁表达用户可感知变化
- 一次提交尽量只描述一类用户可感知变更，避免把 feature/fix/docs 混在一个 subject 里

#### 3.2 不再日常写入 CHANGELOG.md

以下操作只允许在 `bilingual-changelog-release` skill 中执行：

```bash
make release-prep VERSION=X.Y.Z
bash scripts/extract-changelog-section.sh X.Y.Z CHANGELOG.md
bash scripts/backfill-release-notes.sh --tag vX.Y.Z
```

以下旧流程禁止在本 skill 中使用，避免覆盖已人工修订的中文版本段：

```bash
make changelog
git-cliff -o CHANGELOG.md
```

`make changelog` 仅作为紧急修复历史 CHANGELOG 的 legacy target，必须经用户明确同意后使用。

#### 3.3 若用户要求“现在就更新 CHANGELOG”

先判断意图：

- **发版版本段 / release notes / tag 前准备**：停止本 skill 的 Phase 3，加载并执行 `bilingual-changelog-release`。
- **日常 PR 希望留下 changelog 依据**：不要直接编辑 `CHANGELOG.md`，而是给出符合规范的 commit/PR title 建议，并确认 docs/README 已同步。
- **修订已发布 release 的翻译**：加载 `bilingual-changelog-release`，先改 `CHANGELOG.md` 源，再 backfill。

#### 3.4 readiness 输出模板

完成本 Phase 后输出：

```text
CHANGELOG readiness:
- 推荐 commit/PR title: <type>(<group>): <subject>
- 是否进入 CHANGELOG: Yes/No（说明原因）
- 发版时执行: make release-prep VERSION=X.Y.Z
- 当前未运行 git-cliff -o CHANGELOG.md，避免覆盖双语版本段
```

---

## 📤 输出标准

### docs/ 命令文档

- [ ] `docs/en/<group>.md` 已更新（新增节或修改已有节）
- [ ] `docs/zh/<group>.md` 已更新（与英文版结构一致）
- [ ] 双语切换链接正确
- [ ] 参数表格完整且类型/必填标注准确
- [ ] 涉及接口章节已更新（新增/删除/修改 OpenAPI 调用时）

### README RAM 权限汇总

- [ ] `README.md` RAM Permissions 表格已更新（新增/删除接口时）
- [ ] `README.zh-CN.md` RAM 账号接口权限表格已更新
- [ ] Policy JSON 中的 Action 数组已同步

### README Command Overview

- [ ] `README.md` Command Overview 表格已更新
- [ ] `README.zh-CN.md` 命令概览表格已更新
- [ ] 两个表格的命令列表一致

### LLM-facing docs readiness

- [ ] 如修改 `README.md` 或 `docs/en/**`，已执行 `bash scripts/build-llms-full.sh`
- [ ] `llms-full.txt` 已随源文档同步更新，或已明确说明无需更新
- [ ] 如新增 / 删除 / 重命名对外文档，`llms.txt` 导航链接已同步
- [ ] `llms.txt` / `llms-full.txt` 未收录 `docs/internal/**` 内容

### CHANGELOG readiness

- [ ] 已给出推荐 commit/PR title（Conventional Commits + 合理 scope）
- [ ] 已判断本次变更是否应进入 CHANGELOG
- [ ] 未在日常文档同步中运行 `git-cliff -o CHANGELOG.md` / `make changelog`
- [ ] 如用户目标是发版或 release notes，已切换到 `bilingual-changelog-release` skill

## 📚 参考

- [development.md](../../rules/development.md) — 文档同步规则
- [create-cli-command](../create-cli-command/SKILL.md) — CLI 命令封装流程（Phase 5 委托本 skill）
- [feature-development-workflow](../feature-development-workflow/SKILL.md) — 开发流程规范
- [bilingual-changelog-release](../bilingual-changelog-release/SKILL.md) — 双语 CHANGELOG、GitHub Release notes、release-prep/backfill SOP
- [llms-txt spec](../../specs/llms-txt.md) — AgentBay CLI llms.txt / llms-full.txt 方案

## ⚠️ 注意事项

1. **先代码后文档**：文档更新必须在代码变更完成后执行，避免文档与代码不一致
2. **双语必须同步**：docs/en/ 和 docs/zh/ 的结构必须完全一致，不得只更新一方
3. **CHANGELOG 不在日常文档同步中全量生成**：本 skill 只做 readiness；发版版本段生成和中文翻译必须使用 `bilingual-changelog-release`
4. **不要修改已有版本的翻译**：已发布版本的翻译修订必须先改 `CHANGELOG.md` 源，再通过 backfill 同步 GitHub Release
5. **README 表格格式**：不得改变表格列的顺序、宽度或格式风格
6. **命令名一律英文**：文档中所有命令名、参数名均使用英文，不翻译
7. **参考已有风格**：更新 docs/ 时参考同文件中已有命令的文档风格，保持一致
8. **接口变更必须同步权限文档**：只要新增或删除了 OpenAPI 调用，涉及接口章节和 README RAM 权限汇总表必须同步更新，不得遗漏
9. **llms 文档必须随英文对外文档同步**：只要修改 `README.md` 或 `docs/en/**`，必须执行 `bash scripts/build-llms-full.sh` 并同步 `llms-full.txt`；文档结构变化时检查 `llms.txt`
10. **必须输出权限变更摘要**：无论有无接口变更，都必须在终端输出 Phase 1.5.4 要求的摘要，方便用户人工核查
