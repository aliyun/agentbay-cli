# Skills 命令组真实数据测试 Agent

## Context

当前项目测试分两类：

- **单元测试**（`test/unit/cmd/`）：验证命令结构和参数配置，不调用真实 API
- **集成测试**（`test/integration/`）：仅覆盖本地参数校验场景，也不调用真实 API

缺少对 skills 命令组进行**端到端真实 API 调用**的测试手段。  
目标：在 `.qoder/agents/` 下创建第一个项目级 Subagent——`skills-tester.md`，让 AI 作为测试执行者，默认使用 prerelease 环境，动态发现当前所有 skills 子命令并全量测试，输出结构化报告。

---

## 关键设计决策

### 动态发现 + 通用策略（而非写死用例）

**不**在 agent 文件里写死子命令列表。Agent 每次运行时：

1. 执行 `./agentbay skills --help` 枚举所有当前子命令
2. 对每个子命令执行 `./agentbay skills <sub> --help` 获取参数列表
3. 按内置的**命令分类策略**（写类/读类/删除类）自动规划测试用例

**好处**：skills 命令组新增/删除子命令后，Agent 下次运行自动覆盖，无需手动维护测试列表。

### 测试技能包：运行时自动生成

Agent 在临时目录创建最小合法技能包（参考 `packages/find-skills.zip` 的 SKILL.md frontmatter 格式），测试完成后通过 delete 命令清理远程资源。其他用户克隆项目后无需额外准备。

### 认证：读取当前环境变量（不主动索取）

Agent 使用当前 shell 环境中已有的 AK/SK 或本地 OAuth token，不主动提示用户输入凭据。用户需确保在调用 Agent 前已配置好认证信息（见项目 README 认证章节）。

---

## Qoder Agent 文件规范

### 官方格式

Qoder 官方推荐的 Agent 格式是**单文件**：`.qoder/agents/<agent-name>.md`，直接放在 `agents` 目录下。

### AGENT.md 格式

AGENT.md 由两部分组成：**YAML frontmatter** + **Markdown 主体**。

**YAML frontmatter（文件顶部）**：

```yaml
---
name: <agent-name>
description: <描述>
tools: <可选，逗号分隔>
---
```

**字段规范**：

| 字段          | 必填 | 格式规范                                                            |
| ------------- | ---- | ------------------------------------------------------------------- |
| `name`        | 是   | kebab-case（小写字母+连字符），用于标识和任务路由                   |
| `description` | 是   | 自然语言描述，支持多行折行格式。`>` 折叠空白换行，`\|` 保留原始换行 |
| `tools`       | 否   | 逗号分隔的工具列表（默认 `*` 继承全部可用工具）                     |

**Markdown 主体**：

- **Phase 定义**：使用 `## Phase N: 标题` 格式组织执行阶段（如 `## Phase 1: 环境准备`）
- 每个 Phase 内包含具体的执行指令、命令示例、判定标准等
- 支持标准 Markdown 语法（代码块、表格、列表、加粗等）

### 关于 assets 目录

官方单文件格式**不支持**附属资源目录。如需模板、fixtures 等辅助文件，建议：

- 通过 shell heredoc（`cat <<EOF`）在运行时内联生成
- 或将资源放到项目其他目录（如 `test/fixtures/`）下引用

---

## 需要创建的文件

| 操作 | 文件路径                         | 说明                                         |
| ---- | -------------------------------- | -------------------------------------------- |
| 新建 | `.qoder/agents/skills-tester.md` | Skills 测试 Agent 定义文件（官方单文件格式） |

**无需修改任何现有代码或测试文件。**

### 测试技能包：运行时内联生成

测试所需的最小合法 SKILL.md 不再依赖外部模板文件，而是在 Phase 3 通过 shell heredoc 直接在 `/tmp/` 下内联生成。这样：

- 无需在 `assets/` 中维护 `skill-template.md` 模板文件
- 任何人 clone 项目后无需额外准备即可直接使用测试 Agent
- 减少维护负担，skill 模板格式可直接参考 `.qoder/skills/` 下的 `SKILL.md`

### Agent 内联生成 SKILL.md 的方式

Phase 3（创建测试技能包）：

```bash
# 1. 创建主测试技能包目录并内联生成 SKILL.md
mkdir -p /tmp/${TEST_ID}
cat > /tmp/${TEST_ID}/SKILL.md <<EOF
---
name: agentbay-e2e-test-${TEST_ID}
description: |
  Auto-generated test skill for agentbay CLI end-to-end testing.
  This skill is used by the skills-tester agent to validate CLI commands.
  Safe to delete if found in marketplace.
---

# AgentBay CLI E2E Test Skill

This skill is automatically used by the `skills-tester` agent for end-to-end testing of the `agentbay skills` command group.

It has no functional content — its sole purpose is to validate that push/update/list/show/delete commands work correctly.
EOF

# 2. 创建"改名"测试包目录（用于验证 update 不允许改名）
mkdir -p /tmp/${TEST_ID}-renamed
cat > /tmp/${TEST_ID}-renamed/SKILL.md <<EOF
---
name: agentbay-e2e-test-${TEST_ID}-changed
description: |
  Auto-generated test skill for agentbay CLI end-to-end testing (renamed variant).
  This skill is used to test that update command rejects name changes.
  Safe to delete if found in marketplace.
---

# AgentBay CLI E2E Test Skill (Renamed Variant)

This is a variant of the test skill with a different name, used to verify that the update command rejects skill name changes.
EOF

# 3. 记录技能包路径
SKILL_DIR=/tmp/${TEST_ID}
SKILL_DIR_RENAMED=/tmp/${TEST_ID}-renamed
```

---

## Agent 文件完整内容设计

### YAML frontmatter（官方规范格式）

```yaml
---
name: skills-tester
description: >
  对 agentbay skills 命令组执行端到端真实 API 测试。
  动态发现当前所有子命令，默认使用 prerelease 环境，
  自动创建测试技能包、全量跑通所有命令、输出结构化测试报告并清理测试数据。
---
```

> 注：官方仅支持 `name`（必填，小写字母+连字符）、`description`（必填）、`tools`（可选）三个字段。不使用 `id` 字段。

### 正文提示词结构（六个阶段）

#### Phase 0：声明角色与规则

- 本 Agent 是 agentbay skills 命令组的端到端测试执行者
- 默认环境：`AGENTBAY_ENV=prerelease`（用户明确说"生产环境"才不设置）
- 所有命令通过项目根目录的 `./agentbay` 二进制执行

#### Phase 1：环境准备

```
1. 确认 agentbay 二进制
   - 检查 ./agentbay 是否存在
   - 不存在则执行 go build -o agentbay .

2. 确认认证
   - 检查环境变量 AGENTBAY_ACCESS_KEY_ID / AGENTBAY_ACCESS_KEY_SECRET 是否已设置
   - 如未设置，检查本地 OAuth token（~/.agentbay/ 或 AGENTBAY_CLI_CONFIG_DIR）
   - 如两者都不存在，告知用户配置认证后重试，终止测试

3. 设置测试环境变量
   - 检查当前 shell 中 AGENTBAY_ENV 是否已设置（echo $AGENTBAY_ENV）
   - 如已设置：告知用户当前将使用 $AGENTBAY_ENV 环境，尊重用户配置，不覆盖
   - 如未设置，且用户本次未明确说"使用生产/线上环境"：自动 export AGENTBAY_ENV=prerelease
   - 如用户本次明确说"使用生产/线上环境"：不设置 AGENTBAY_ENV，让 CLI 默认使用 production

4. 生成测试唯一标识
   - TEST_ID = agentbay-test-$(date +%s)
   - 后续所有测试资源名称均包含此 ID，便于识别和清理
```

#### Phase 2：动态发现命令结构

```
1. 执行 ./agentbay skills --help，解析出所有子命令名称
2. 对每个子命令执行 ./agentbay skills <sub> --help，记录：
   - 参数名称和类型（string/bool/stringArray/int）
   - 必填参数（Required）
   - 可选参数及默认值
3. 按如下分类标记每个子命令：
   - 写类（CREATE）：push、update 等（需要上传文件的）
   - 读类（READ）：list、show 等（只读查询）
   - 删除类（DELETE）：delete 等（带 --yes 参数的）
```

#### Phase 3：创建测试技能包

```bash
# 1. 创建主测试技能包目录并内联生成 SKILL.md
mkdir -p /tmp/${TEST_ID}
cat > /tmp/${TEST_ID}/SKILL.md <<EOF
---
name: agentbay-e2e-test-${TEST_ID}
description: |
  Auto-generated test skill for agentbay CLI end-to-end testing.
  This skill is used by the skills-tester agent to validate CLI commands.
  Safe to delete if found in marketplace.
---

# AgentBay CLI E2E Test Skill

This skill is automatically used by the `skills-tester` agent for end-to-end testing of the `agentbay skills` command group.

It has no functional content — its sole purpose is to validate that push/update/list/show/delete commands work correctly.
EOF

# 2. 创建"改名"测试包目录（用于验证 update 不允许改名）
mkdir -p /tmp/${TEST_ID}-renamed
cat > /tmp/${TEST_ID}-renamed/SKILL.md <<EOF
---
name: agentbay-e2e-test-${TEST_ID}-changed
description: |
  Auto-generated test skill for agentbay CLI end-to-end testing (renamed variant).
  This skill is used to test that update command rejects name changes.
  Safe to delete if found in marketplace.
---

# AgentBay CLI E2E Test Skill (Renamed Variant)

This is a variant of the test skill with a different name, used to verify that the update command rejects skill name changes.
EOF

# 3. 记录技能包路径
SKILL_DIR=/tmp/${TEST_ID}
SKILL_DIR_RENAMED=/tmp/${TEST_ID}-renamed
```

#### Phase 4：执行测试（按固定顺序）

**测试顺序固定**：写类命令（获取 skill-id）→ 读类命令（验证数据）→ 更新类命令 → 删除类命令 → 再次读类（验证删除）

对每个发现的子命令，按其分类执行对应的测试策略：

**写类（CREATE）测试策略**

- 基础用例：使用必填参数执行，期望成功并拿到资源 ID
- 可选参数用例：依次测试每个可选参数（单独使用）
- 特殊用例：stringArray 类型参数测试多值情况（如 `--tag a --tag b`）
- 失败用例：缺少必填参数，期望明确的错误提示

**读类（READ）测试策略**

- list 类：默认参数、JSON 输出（`-o json` 或 `--output json`）、过滤参数、分页参数
- show 类：查看已创建的资源（用写类拿到的 ID）、查看不存在的 ID（期望错误含 requestId）

**删除类（DELETE）测试策略**

- 使用 `--yes` 跳过确认，直接删除写类创建的测试资源
- 删除后用读类命令验证资源已不存在

**通用测试判定标准**

- PASS：命令退出码为 0，且输出符合期望（含预期字段/不含错误信息）
- FAIL：命令退出码非 0，或输出包含 `[ERROR]`，或期望字段缺失
- 失败时记录：完整错误信息、输出中的 RequestID（格式 `[INFO] Request ID: xxx`）、错误分析

#### Phase 5：生成测试报告

**输出方式**：

1. **终端输出**：按如下格式在对话中输出完整结构化测试报告
2. **文件输出**：将同样的报告内容写入 `/tmp/agentbay-skills-test-report-${TEST_ID}.md`
   - 文件内容与终端输出保持一致
   - 在文件末尾追加注释：`<!-- Report generated at <ISO 8601> -->`
3. **路径标注**：在终端输出的报告末尾，单独追加一行：`[INFO] 测试报告已保存至: /tmp/agentbay-skills-test-report-${TEST_ID}.md`

报告格式：

```
════════════════════════════════════════════
  AgentBay Skills 命令组端到端测试报告
  环境:    prerelease
  时间:    <ISO 8601 时间>
  Test ID: <TEST_ID>
════════════════════════════════════════════

【发现的子命令】
  push    [写类]
  list    [读类]
  show    [读类]
  update  [写类]
  delete  [删除类]

【测试结果】

▶ skills push
  [PASS] 基础推送（无标签）              skill-id: skill-xxxxx
  [PASS] 推送带单个 --tag
  [PASS] 推送带多个 --tag
  [FAIL] 缺少位置参数                   错误: accepts 1 arg(s), received 0
                                        分析: 参数校验正常，符合预期（此为预期失败）

▶ skills list
  [PASS] 默认表格输出
  [PASS] JSON 输出（-o json）           totalCount 字段存在
  [PASS] 按名称过滤（--name）
  [PASS] 分页参数（--page --size）

▶ skills show
  [PASS] 查看已创建的技能               SkillId/Name 字段正常展示
  [PASS] 查看不存在的 ID               错误信息含 RequestID

▶ skills update
  [PASS] 更新技能包（--file）
  [PASS] 更新标签（--tag）

▶ skills delete
  [PASS] 删除技能（--yes）
  [PASS] 删除后列表验证                 技能已不在列表中

════════════════════════════════════════════
  汇总: 14 通过 / 0 失败
  测试数据已清理
════════════════════════════════════════════
```

#### Phase 6：清理

- 确认测试创建的 skill-id 已通过 delete 命令删除
- 删除临时目录 `/tmp/${TEST_ID}`
- 删除测试报告文件 `/tmp/agentbay-skills-test-report-${TEST_ID}.md`
- 如有任何 FAIL 用例，在报告末尾单独列出失败用例的完整错误信息和 RequestID，便于排查

---

## 使用方式

1. 在 Qoder 界面切换到 `Skills 命令组端到端测试` Agent
2. 发送「开始测试」或「测试 skills 命令组」
3. 如需测试生产环境，说明「使用生产环境」
4. Agent 全自动执行，完成后输出测试报告

**前提条件**（用户需提前配置）：

- 已设置 `AGENTBAY_ACCESS_KEY_ID` + `AGENTBAY_ACCESS_KEY_SECRET`，或已执行 `agentbay login`

---

## 验证方式

创建文件后：

1. Qoder 界面确认 `Skills 命令组端到端测试` Agent 出现在 Agent 列表
2. 激活 Agent，发送「开始测试」
3. 观察 Agent 是否：
   - 先执行 `--help` 动态发现子命令
   - 自动创建 `/tmp/agentbay-test-xxx/` 临时技能包
   - 按顺序跑完所有用例
   - 输出带 PASS/FAIL 标记的结构化报告
   - 测试结束后删除测试资源
