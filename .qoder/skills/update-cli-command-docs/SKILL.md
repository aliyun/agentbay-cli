---
name: update-cli-command-docs
description: AgentBay CLI 命令文档同步更新流程（docs/、README、CHANGELOG）
---

# Update CLI Command Docs

## 📋 职责

在 AgentBay CLI 需求开发完成后，同步更新以下三类文档：

1. **docs/ 命令文档** — `docs/en/<group>.md` 和 `docs/zh/<group>.md`
2. **README Command Overview 表格** — `README.md` 和 `README.zh-CN.md`
3. **CHANGELOG.md** — git-cliff 生成 + 中文翻译补充

**本 skill 不涉及代码实现，仅负责文档层面的同步。**

对客文档（`cli-analysis/` 目录、钉钉文档）不在此 skill 范围内，需手动同步。

## 🎯 触发场景

当出现以下情况时触发：

- 新增 CLI 命令后需同步文档
- 修改已有命令的参数、默认值、输出格式后需同步文档
- 发版前需更新 CHANGELOG 中文翻译
- 用户明确提出"更新文档"、"同步文档"、"更新 CHANGELOG"等需求
- `create-cli-command` skill 的 Phase 5 委托调用

## 🔗 前置约束

**本 skill 必须在代码变更完成后触发**，不得在代码尚未完成时提前生成文档。典型组合：

| 场景 | 前置 skill | 本 skill 时机 |
|------|-----------|--------------|
| 新增命令 | `create-cli-command` Phase 1-4 完成后 | 替代其 Phase 5 |
| 修改命令 | 直接修改代码后 | 代码修改完成即可 |
| 发版前 | `feature-development-workflow` Phase 4 完成后 | 在 Phase 5 推送之前 |
| 仅补翻译 | 无前置 | 任意时间 |

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

   | 代码文件 | 命令组 | docs 文件 |
   |---------|--------|----------|
   | `cmd/apikey*.go`, `cmd/concurrency.go` | apikey | `docs/en/apikey.md` / `docs/zh/apikey.md` |
   | `cmd/image*.go` | image | `docs/en/image.md` / `docs/zh/image.md` |
   | `cmd/network.go` | network | `docs/en/network.md` / `docs/zh/network.md` |
   | `cmd/skills*.go` | skills | `docs/en/skills.md` / `docs/zh/skills.md` |
   | `cmd/docker.go` | docker | `docs/en/docker.md` / `docs/zh/docker.md` |
   | `cmd/login.go`, `cmd/logout.go`, `cmd/constants.go` | core | `docs/en/core.md` / `docs/zh/core.md` |

3. **判定文档更新类型**

   | 变更类型 | docs 更新 | README 更新 | CHANGELOG 更新 |
   |---------|----------|------------|---------------|
   | 新增命令 | ✅ 添加完整命令文档节 | ✅ 在表格中添加子命令 | ✅ |
   | 新增参数 | ✅ 更新参数表格 | ✅ 如影响简述则更新 | ✅ |
   | 修改参数默认值 | ✅ 更新默认值和说明 | ✅ 如影响简述则更新 | ✅ |
   | 修改输出格式 | ✅ 更新输出示例 | ❌ 通常不需要 | ✅ |
   | 修改命令行为 | ✅ 更新说明和注意事项 | ✅ 更新简述 | ✅ |
   | 仅内部重构 | ❌ 通常不需要 | ❌ | 视情况 |
   | 仅补充翻译 | ❌ | ❌ | ✅ 仅 Phase 3 |

4. **向用户确认更新范围**

   展示分析结果，例如：
   > 检测到以下变更需要文档同步：
   > - **新增命令**: `apikey status` → 需更新 `docs/en/apikey.md`、`docs/zh/apikey.md`、README 表格、CHANGELOG
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

```markdown
---

### `<group> <subcommand>`

<一句话描述>

```bash
agentbay <group> <subcommand> [flags]
```

**Flags:**

| Flag | Short | Type | Required | Description |
|------|-------|------|----------|-------------|
| `--param` | `-p` | string | Yes | 说明 |

**Notes:**

- 注意事项
```

**中文版模板**：

```markdown
---

### `<group> <subcommand>`

<一句话中文描述>

```bash
agentbay <group> <subcommand> [flags]
```

**参数：**

| 参数 | 短参数 | 类型 | 必填 | 说明 |
|------|--------|------|------|------|
| `--param` | `-p` | string | 是 | 中文说明 |

**注意事项：**

- 注意事项
```

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
```
```

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
```
```

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
```

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

### Phase 3: 更新 CHANGELOG.md

**原则**：CHANGELOG 由 git-cliff 自动生成英文部分，中文翻译由 AI 生成后经用户确认写入。

#### 3.1 获取最新远程信息

```bash
git fetch aliyun
```

#### 3.2 比对当前分支与 aliyun/master 的差异

```bash
git log aliyun/master..HEAD --oneline
```

此步骤确认有哪些新 commit 将被纳入 CHANGELOG，验证 git-cliff 的输出是否完整。

#### 3.3 生成/更新 CHANGELOG

```bash
# 优先使用 Makefile target
make changelog

# 或直接使用 git-cliff
git-cliff -o CHANGELOG.md
```

**如果 git-cliff 未安装**，提示用户安装：

```bash
make changelog-install
# 或
brew install git-cliff
```

#### 3.4 检查生成的 CHANGELOG

读取更新后的 `CHANGELOG.md`，检查：

- `[Unreleased]` 部分是否包含本次变更
- 分类是否正确（Features / Bug Fixes / Documentation / Other Changes）
- scope 标注是否正确（`**apikey**:`、`**image**:` 等）

#### 3.5 补充中文翻译

**查找待翻译占位符**：

在 `CHANGELOG.md` 中搜索 `<!-- 中文翻译待补充` 或 `<!-- 中文翻译待补充 / Add Chinese translation before release -->`。

**翻译规则**：

| 英文分类 | 中文分类 | 条目动词示例 |
|---------|---------|------------|
| Features | 新功能 | 新增、支持 |
| Bug Fixes | 问题修复 | 修正、修复、兼容 |
| Documentation | 文档 | 补充、更新、新增 |
| Refactoring | 重构 | 重构、简化 |
| Performance | 性能 | 优化、提升 |
| Other Changes | 其他变更 | 更新、移除、新增 |

**翻译格式**：

在 `* * *` 分隔线下方，镜像英文结构的中文版本：

```markdown
* * *

### 新功能

- **apikey**: 新增 `apikey delete` 命令，支持多步骤确认
- **image**: 新增 `image warmup-status` 命令

### 问题修复

- **client**: 兼容响应中 `HttpStatusCode` 字段为字符串类型的情况
```

**翻译要点**：

- scope 保留英文（`**apikey**:`、`**image**:`、`**client**:`）
- 命令名保留英文（`` `apikey delete` `` 不翻译）
- 遵循已有版本的翻译风格（参考 `CHANGELOG.md` 中 v0.2.7、v0.2.9 等已有翻译）
- 合并同类条目时保持语义不丢失

#### 3.6 向用户确认翻译

将生成的中文翻译完整展示给用户：

> 已生成以下 CHANGELOG 中文翻译，请确认：
>
> ```
> ### 新功能
> - **apikey**: 新增 `apikey status` 命令，用于查询 API Key 状态
> ```
>
> 是否需要修改？确认后将写入 CHANGELOG.md。

**⚠️ 铁律：翻译必须经用户确认后才可写入，不得自动替换占位符。**

#### 3.7 写入并验证

将翻译替换占位符，写入 `CHANGELOG.md`。验证：

- `<!-- 中文翻译待补充 -->` 占位符已被替换
- `* * *` 分隔线在英文和中文之间
- 没有残留的占位符

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

### CHANGELOG

- [ ] `git fetch aliyun` 已执行
- [ ] `git-cliff -o CHANGELOG.md` 已执行（或 `make changelog`）
- [ ] 中文翻译已补充（占位符已替换）
- [ ] 翻译已经用户确认
- [ ] 无残留 `<!-- 中文翻译待补充 -->` 占位符

## 📚 参考

- [development.md](../../rules/development.md) — 文档同步规则
- [create-cli-command](../create-cli-command/SKILL.md) — CLI 命令封装流程（Phase 5 委托本 skill）
- [feature-development-workflow](../feature-development-workflow/SKILL.md) — 开发流程规范

## ⚠️ 注意事项

1. **先代码后文档**：文档更新必须在代码变更完成后执行，避免文档与代码不一致
2. **双语必须同步**：docs/en/ 和 docs/zh/ 的结构必须完全一致，不得只更新一方
3. **CHANGELOG 翻译需确认**：AI 生成的中文翻译必须经用户确认后才写入，不得自动替换
4. **不要修改已有版本的翻译**：CHANGELOG 中已发布版本的翻译不得修改，只更新 `[Unreleased]` 或最新版本
5. **README 表格格式**：不得改变表格列的顺序、宽度或格式风格
6. **命令名一律英文**：文档中所有命令名、参数名均使用英文，不翻译
7. **参考已有风格**：更新 docs/ 时参考同文件中已有命令的文档风格，保持一致
8. **接口变更必须同步权限文档**：只要新增或删除了 OpenAPI 调用，涉及接口章节和 README RAM 权限汇总表必须同步更新，不得遗漏
9. **必须输出权限变更摘要**：无论有无接口变更，都必须在终端输出 Phase 1.5.4 要求的摘要，方便用户人工核查
