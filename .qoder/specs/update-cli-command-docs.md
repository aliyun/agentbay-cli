# 新建 `update-cli-command-docs` Skill

## Context

当前 AgentBay CLI 的文档更新流程分散在 `create-cli-command` 的 Phase 5 中，存在以下问题：

1. **文档更新逻辑过于简略**：Phase 5 只有 5.1/5.2/5.3 三小段描述，缺少 CHANGELOG 更新流程、中文翻译质量保障、变更分析等关键步骤
2. **CHANGELOG 更新无标准流程**：fetch aliyun → 比对 master → git-cliff 生成 → 中文翻译的完整流程没有被任何 skill 覆盖
3. **文档更新无法独立触发**：修改参数默认值、发版前补翻译等场景不需要走 create-cli-command 全流程，但仍需文档同步
4. **双语一致性缺乏验证机制**：没有强制检查 docs/en/ 和 docs/zh/ 的结构对等

**目标**：创建独立的 `update-cli-command-docs` skill，覆盖 docs 命令文档、README 表格、CHANGELOG 三类文档的完整同步流程，并调整现有 skill 和规则实现联动。

---

## 实施步骤

### Step 1: 新建 `.qoder/skills/update-cli-command-docs/SKILL.md`

创建 skill 主文件，包含 4 个 Phase：

**Phase 0: 变更分析**
- 检查 `git diff --name-only` 识别变更文件
- 映射变更文件到命令组（apikey/image/network/skills/docker/core）
- 判定文档更新类型矩阵（新增命令/新增参数/修改默认值/修改输出/仅内部重构）
- 向用户确认更新范围

**Phase 1: 更新 docs/ 命令文档**
- 1.1 读取现有文档结构
- 1.2 新增命令 → 追加完整命令文档节（中英文模板，参数表格，示例）
- 1.3 修改命令 → 更新已有节（参数表格、说明、注意事项）
- 1.4 双语同步验证（命令数一致、参数数一致、必填列一致）
- 关键约束：命令名/参数名一律英文；表头英文 `Flag/Short/Type/Required/Description` vs 中文 `参数/短参数/类型/必填/说明`；破坏性操作必须包含 `--yes` 说明

**Phase 2: 更新 README Command Overview 表格**
- 定位 `## Command Overview` / `## 命令概览` 表格
- 新增子命令 → 在对应 Group 行按逻辑序追加
- 修改子命令 → 替换旧名
- 约束：命令名英文（中文版也不翻译）、两组表格内容一致

**Phase 3: 更新 CHANGELOG.md**
- 3.1 `git fetch aliyun`
- 3.2 `git log aliyun/master..HEAD --oneline` 比对差异
- 3.3 `make changelog` 或 `git-cliff -o CHANGELOG.md` 生成
- 3.4 检查 `[Unreleased]` 部分是否包含本次变更、分类是否正确
- 3.5 查找 `<!-- 中文翻译待补充 -->` 占位符，AI 生成翻译（参考已有版本风格，scope 保留英文）
- 3.6 **向用户展示翻译结果并确认**（铁律：翻译必须经用户确认才写入）
- 3.7 替换占位符并验证无残留

翻译规则：
| 英文 | 中文 |
|------|------|
| Features | 新功能 |
| Bug Fixes | 问题修复 |
| Documentation | 文档 |
| Refactoring | 重构 |
| Performance | 性能 |
| Other Changes | 其他变更 |

---

### Step 2: 修改 `.qoder/rules/development.md` — 自动装配规则表

在 Skill 自动装配规则表中新增一行，并修改组合使用行：

**新增行**：
```
| 更新/同步 CLI 命令文档（README、docs/、CHANGELOG） | **update-cli-command-docs** | `.qoder/skills/update-cli-command-docs/SKILL.md` |
```

**修改组合行**（原"两者组合"→"三者组合"）：
```
| 新增 CLI 命令类需求（同时触发上述三条） | **三者组合使用**：先 workflow 拉分支/建档 → 再 create-cli-command 实现 → 再 update-cli-command-docs 同步文档 → 回到 workflow 提交/推送 |
```

**执行铁律新增第 5 条**：
```
5. **文档同步**：create-cli-command 的 Phase 5 已委托 `update-cli-command-docs` skill，文档操作不得在 create-cli-command 中内联执行。
```

**"新增或修改命令必须同步更新文档和测试用例"章节**开头增加引用说明，检查清单增加：
```
- [ ] `update-cli-command-docs` skill 已执行（或已完成等效的手动文档同步）
```

---

### Step 3: 修改 `.qoder/skills/create-cli-command/SKILL.md` — Phase 5 完全委托

将 Phase 5（约第 328-373 行）的 5.1/5.2/5.3 详细步骤**全部删除**，替换为：

```markdown
### Phase 5: 文档生成与同步

**本阶段委托 `update-cli-command-docs` skill 执行**，不在本 skill 内展开。

加载并执行 `.qoder/skills/update-cli-command-docs/SKILL.md`，该 skill 将完成：
- 更新 `docs/en/<group>.md` 和 `docs/zh/<group>.md`
- 更新 `README.md` 和 `README.zh-CN.md` Command Overview 表格
- 更新 `CHANGELOG.md`（git-cliff 生成 + 中文翻译）

> ⚠️ 不得在本 Phase 中内联执行文档操作，必须遵循 `update-cli-command-docs` 的 Phase 0-3 完整流程。

对客文档（`cli-analysis/` 目录、钉钉文档）不在此 skill 范围内，需手动同步。
```

---

### Step 4: 修改 `.qoder/skills/feature-development-workflow/SKILL.md` — Phase 2 增加引用

在 Phase 2 本地开发规则列表（约第 148-152 行）追加一条：

```
- 文档同步遵循 update-cli-command-docs skill（由 create-cli-command Phase 5 自动委托，或独立触发）
```

---

## 涉及文件清单

| 操作 | 文件路径 |
|------|---------|
| **新建** | `.qoder/skills/update-cli-command-docs/SKILL.md` |
| **修改** | `.qoder/rules/development.md`（自动装配表 + 铁律 + 文档章节） |
| **修改** | `.qoder/skills/create-cli-command/SKILL.md`（Phase 5 → 委托） |
| **修改** | `.qoder/skills/feature-development-workflow/SKILL.md`（Phase 2 +1 行） |

---

## 验证方式

1. **语法验证**：检查 SKILL.md 的 YAML frontmatter 和 Markdown 格式是否正确
2. **交叉引用**：确认 development.md 中新增的 skill 路径与实际文件路径一致
3. **委托链完整性**：确认 create-cli-command Phase 5 → update-cli-command-docs 的引用关系正确
4. **内容一致性**：确认翻译规则表格与 CHANGELOG.md 中已有版本（v0.2.7, v0.2.9）的翻译风格匹配
5. **手动跑通**：在当前 `feat/dev-apikey` 分支上模拟执行 skill 流程，验证 `git fetch aliyun` + `git log aliyun/master..HEAD` + `make changelog` 能正常工作
