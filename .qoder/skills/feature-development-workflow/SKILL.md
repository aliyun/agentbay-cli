---
name: feature-development-workflow
description: AgentBay CLI 需求开发全流程规范(分支管理、推送、提交、PR),不含具体需求内容
---

# Feature Development Workflow

## 📋 职责

规定 agentbay-cli 在接到**任何新需求**时的标准开发流程,覆盖从拉分支、本地开发、构建验证、提交、推送到提 PR 的全链路操作规范。

**本 skill 不包含任何具体需求的代码实现细节**,仅约束流程本身。

## 🎯 触发场景

当用户提出以下诉求时触发:

- "开发一个 XX 功能/命令"
- "新增一个 XX 需求"
- "开始一个新需求"
- "按开发流程启动 XX"
- 任何需要新建 feat 分支开发的需求

## 🌐 远程仓库拓扑

本项目只保留两个远程,**不使用 upstream(个人 fork)**:

| Remote   | 地址                                                       | 用途                          |
| -------- | ---------------------------------------------------------- | ----------------------------- |
| `origin` | `git@gitlab.alibaba-inc.com:InnoArchClub/agentbay-cli.git` | 内网代码仓,归档/协作/CI       |
| `aliyun` | `git@github.com:aliyun/agentbay-cli.git`                   | 对客真实代码仓,合入主线的目标 |

验证命令:

```bash
git remote -v
# 必须只有 origin 和 aliyun,不允许出现 upstream
```

如出现多余 remote:

```bash
git remote remove <name>
```

## 📒 变更档案(Change Record)双路径

为让每个需求都有"需求 → 设计 → 代码 → 提交 → PR → 发布"的完整追溯链,启动任何新需求前必须**二选一**建立变更档案,两条路径都是强制的,目的一致 — 在 Git 仓库内留下可审阅的档案:

| 路径                    | 触发场景                                                                        | 产物位置                                                               | 备注                                                                                          |
| ----------------------- | ------------------------------------------------------------------------------- | ---------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| **A. Qoder Quest Mode** | 通过 `⌘ E` 进入 Quest 模式,在 Design 阶段让 AI 生成 Spec,Review 后点 `Download` | `.qoder/specs/<feature>.md`(Qoder 客户端 Download 默认写盘位置,零配置) | 推荐用于复杂需求;Spec 可与 AI 协同编辑;文件默认 untracked,需 `git add` 入库                   |
| **B. 手动 CR 目录**     | 未启用 Quest Mode / 轻量需求 / 已有外部设计稿需沉淀                             | `.qoder/changes/CR-<YYYY-MM-DD>-<feature-name>/`                       | 包含 spec.md / design.md / tasks.md / decisions.md / test-plan.md / rollback.md / trace.md 等 |

**核心规则**:

- 需求启动时二选一,**不允许两条都不走**
- Quest Mode 下 `.qoder/specs/<feature>.md` 由 Qoder 在点 Download 后自动落盘,**必须 `git add` 纳入版本库**,避免档案与代码脱节
- 非 Quest 场景必须手动在 `.qoder/changes/` 下建 CR 目录,命名严格遵守 `CR-<YYYY-MM-DD>-<feature-name>` 格式,与 feat 分支名强关联
- 两种路径产物都应最终在 `trace.md` / 对应 Spec 中登记:feat 分支名、关键 commit SHA、推送目标 remote、PR 链接、合并 commit、release tag

## 🚀 标准开发流程(Task 顺序不可颠倒)

### Phase 0: 变更档案初始化(追溯链起点)

> ⚠️ **铁律**:此 Phase 产出之前,**不允许进入 Phase 1 拉分支**。

选择一条路径建立档案:

**路径 A · Quest Mode(推荐)**

1. `⌘ E` 打开 Quest 面板
2. `New Task` → 勾选相关文件作为 Context → 粘贴需求描述 → 回车(**不要勾 Execute Directly**,否则跳过 Design 阶段不会生成 Spec)
3. 等 Qoder 在 Design 阶段生成 Spec(右侧 Spec Tab 流式显示)
4. 与 AI 协同在对话区迭代 Spec,满意后点 `Run Spec` 进入执行
5. **关键**:Spec Tab 右上角点 `Download` → Qoder 客户端会自动写入 `.qoder/specs/<feature>.md`(文件名基于 Spec 标题 slug 生成,零配置)
6. Spec 文件默认 untracked,后续 commit 中必须一并 `git add` 纳入版本库

**路径 B · 手动 CR 目录**

1. `mkdir -p .qoder/changes/CR-$(date +%F)-<feature-name>`
2. 参考 `.qoder/changes/TEMPLATE.md`(若存在)建立:
   - `spec.md`:需求背景 / 目标 / 非目标 / 范围
   - `design.md`:接口设计 / 流程图 / 状态机
   - `tasks.md`:任务分解(建议映射到 todo_write 的 tasklist)
   - `decisions.md`:关键决策(交互默认值 / 参数命名 / 分支策略 等)
   - `test-plan.md`:单测 / 集成 / 回归
   - `rollback.md`:回滚预案
   - `trace.md`:后续持续更新的追溯链(分支 / commits / push / PR / release)
3. 档案先建起来再拉 feat 分支,保证 Phase 1 的第一个 commit 可以带着档案进入仓库

### Phase 1: 需求启动(分支准备)

> ⚠️ **重要**:任何写代码动作前必须确定一条干净、起点正确的 feat 分支,禁止在 master 直接改动。

1. **同步两个远程最新状态**

   ```bash
   git fetch origin
   git fetch aliyun
   ```

2. **判断是否可复用当前 feat 分支**(避免每次都新切)

   先读取当前仓库状态:

   ```bash
   git branch --show-current           # 当前分支名
   git status --short                  # 工作区是否干净
   git log --oneline aliyun/master..HEAD -1   # 是否领先 aliyun/master
   git log --oneline HEAD..aliyun/master -1   # 是否落后 aliyun/master
   ```

   **可复用的充要条件**(必须全部满足):
   - [ ] 当前分支以 `feat-` 或 `feat/` 开头,且**不是** master / main
   - [ ] 工作区干净(`git status --short` 输出为空,允许未入库的变更档案除外,但需在后续一起提交)
   - [ ] 当前分支**不落后** `aliyun/master`(即 `HEAD..aliyun/master` 为空;若落后,需先 rebase 或用户确认处理策略)
   - [ ] 当前分支主题与新需求相关,用户明确同意复用

   **决策分发**:
   - ✅ 全部满足 → **主动询问用户**:
     > "检测到当前在 `feat/xxx` 分支,已与 `aliyun/master` 同步且工作区干净。是 **直接复用当前分支** 继续开发,还是 **从 `aliyun/master` 新切一条 feat 分支**?"
   - ❌ 任一不满足 → **必须新切**,不再询问,按第 3 步执行。

3. **新切 feat 分支(默认路径)**

   基于 `aliyun/master` 创建(以 aliyun 主线为真源):

   ```bash
   git checkout -b feat-<feature-name> aliyun/master
   ```

   分支命名规范: `feat-<短横线分隔的需求关键词>`,如 `feat-image-delete`、`feat-apikey-concurrency`。

4. **确认最终分支状态**:

   ```bash
   git log --oneline -1
   git status
   ```

   无论复用还是新切,后续 Phase 2-6 的所有动作都在这条 feat 分支上进行。

### Phase 2: 本地开发

严格遵守 [development.md](../../rules/development.md) 规则:

- 接口变更必须同步所有 mock 类
- 新增/修改命令必须同步 README 和对外文档
- 文档同步遵循 update-cli-command-docs skill（由 create-cli-command Phase 5 自动委托，或独立触发）
- CHANGELOG 日常开发只做 readiness；发版版本段生成/翻译/回灌遵循 bilingual-changelog-release skill
- 新增命令必须有单元测试
- 参数使用命名参数(`--name`),不使用位置参数

开发过程**禁止任何自动 push 动作**。

### Phase 3: 构建与测试验证

每次提交前必须本地全量验证:

```bash
go build -o agentbay .
go test ./... -count=1
```

两条命令全部通过后才能进入提交环节。`go build -o agentbay .` 必须生成项目根目录的可执行二进制，不能只用 `go build ./...` 替代。

### Phase 4: 本地提交(按用户指令)

> ⚠️ **铁律**:**在用户明确指示之前,不得执行 `git add` / `git commit`**。

用户指示提交后:

1. **先展示变更**,供用户审阅:

   ```bash
   git status
   git diff --stat
   ```

2. **使用 Conventional Commits 规范**（确保 release-prep 可正确生成 CHANGELOG）:

   ```bash
   git add -A
   git commit -m "<type>(<scope>): <description>

   - 具体改动点 1
   - 具体改动点 2
   - 具体改动点 3"
   ```

   `<type>` 从以下中选: `feat` / `fix` / `docs` / `refactor` / `perf` / `revert` / `test` / `style` / `chore` / `ci` / `build`。
   用户可感知变更优先使用会进入 CHANGELOG 的 `feat` / `fix` / `docs` / `refactor` / `perf` / `revert`；`<scope>` 优先使用 `apikey`、`image`、`network`、`skills`、`docker`、`core`、`client`。
   不兼容变更必须使用 `!` 或 footer `BREAKING CHANGE:`。

3. **确认提交结果**:

   ```bash
   git log --oneline -3
   ```

4. **更新变更档案**:
   - 路径 A:检查 `.qoder/specs/<feature>.md` 是否已加入本次 commit
   - 路径 B:将本次 commit SHA 与动机摘要追加到 `.qoder/changes/CR-xxx/trace.md`

### Phase 5: 推送到远程(按用户指令)

> ⚠️ **铁律**:**在用户明确指示推送之前,不得执行任何 `git push`**。

推送时机和目标由用户决定,**只允许推送到 `origin` 和 `aliyun`**,推送命令模板:

- **推送到 origin(内网归档)**:

  ```bash
  git push origin feat-<feature-name>
  ```

- **推送到 aliyun(对客主线,PR 源分支)**:

  ```bash
  git push aliyun feat-<feature-name>
  ```

- **两个远程同时推送**(如用户要求):
  ```bash
  git push origin feat-<feature-name>
  git push aliyun feat-<feature-name>
  ```

**禁止操作**:

- ❌ `git push --force` / `git push -f` 到任何远程(除非用户显式授权)
- ❌ 推送到 `master` / `main` 主分支
- ❌ 推送到已移除的 `upstream`
- ❌ 使用 `--no-verify` 跳过 hook

**推送完成后**:

- 把推送时间、目标 remote 记录到变更档案的 `trace.md`(路径 B)或 Spec 的"执行记录"章节(路径 A)

### Phase 6: 提 PR(按用户指令)

当 feat 分支已推送到 `aliyun` 后,由用户在 GitHub 网页操作或通过 `gh` CLI 提 PR 到 `aliyun/master`,PR 模板要求:

- 标题复用 commit message 标题,并保持 Conventional Commits 格式,供 release-prep 生成 CHANGELOG
- 描述包含:需求背景 / 改动点 / 测试情况 / 风险说明

PR 合并由用户(或指定 reviewer)审批,合并后的清理动作:

```bash
git checkout master
git fetch aliyun
git reset --hard aliyun/master
git branch -D feat-<feature-name>                # 本地删除
git push origin --delete feat-<feature-name>      # origin 删除
git push aliyun --delete feat-<feature-name>      # aliyun 删除(可选,或由 PR 合并页勾选)
```

**档案归档**:

- 在 CR 的 `trace.md`(路径 B)或 Quest Spec(路径 A)登记 PR 链接、合并 commit、对应的 release tag
- 档案目录保持在仓库中,**不要删除**,作为历史追溯资产

### Phase 7: 发版交接(仅发版需求触发)

如果用户要求继续发版、准备 tag、生成 GitHub Release notes 或回灌历史 release body，**立即加载 `bilingual-changelog-release` skill**，按其 SOP 执行：

```bash
make release-prep VERSION=X.Y.Z
```

注意：日常 feature 分支开发完成后不要在本 workflow 中运行 `git-cliff -o CHANGELOG.md` 或 `make changelog`；CHANGELOG 版本段由发版阶段统一生成和翻译。

## 🔒 关键原则

1. **显式授权原则**:`commit` 和 `push` 都需要用户明确指示后才执行
2. **真源原则**:feat 分支从 `aliyun/master` 拉出,确保起点是对客主线
3. **双远程原则**:推送仅限 `origin` 和 `aliyun`,不引入额外 fork
4. **可回滚原则**:每次提交必须是本地验证通过的稳定状态
5. **分支隔离原则**:每个需求独立 feat 分支,禁止混合多个需求
6. **可追溯原则**:每个需求必须在 `.qoder/specs/`(Quest Mode Download 自动落盘)或 `.qoder/changes/`(手动 CR)下建立档案,档案随代码一并入库

## ✅ 流程检查清单

需求启动前:

- [ ] `git remote -v` 只有 origin 和 aliyun
- [ ] 已执行 `git fetch aliyun`
- [ ] 当前 HEAD 位于 feat 分支(新切 或 经用户确认复用的已有 feat 分支)
- [ ] 若复用:已满足"不落后 aliyun/master + 工作区干净 + 用户同意"三条件
- [ ] 变更档案已建立(`.qoder/specs/<feature>.md` Quest Spec 或 `.qoder/changes/CR-xxx/` 目录)

提交前:

- [ ] `go build -o agentbay .` 通过并生成根目录二进制
- [ ] `go test ./... -count=1` 通过
- [ ] README / 对外文档 / 单元测试 / mock 类均已同步
- [ ] 已完成 CHANGELOG readiness：commit/PR title 符合 Conventional Commits，必要时带命令组 scope
- [ ] `git status` 无预期外的改动
- [ ] 变更档案文件已 `git add`,与代码一同提交
- [ ] 已获得用户提交授权

推送前:

- [ ] 已获得用户推送授权
- [ ] 明确用户指定的目标 remote(origin / aliyun / 两者)
- [ ] 分支名以 `feat-` 开头,不是 master
- [ ] `trace.md` / Quest Spec 已更新本次推送意图

## 📚 关联规则

- [development.md](../../rules/development.md) - 代码规范 / Mock 同步 / Commit 规范
- [create-cli-command](../create-cli-command/SKILL.md) - CLI 命令封装的具体实现流程
- [update-cli-command-docs](../update-cli-command-docs/SKILL.md) - docs/README 文档同步与 CHANGELOG readiness
- [bilingual-changelog-release](../bilingual-changelog-release/SKILL.md) - 发版阶段双语 CHANGELOG 与 GitHub Release notes SOP

## ⚠️ 常见陷阱

| 陷阱                  | 表现                                                       | 规避                                                                    |
| --------------------- | ---------------------------------------------------------- | ----------------------------------------------------------------------- |
| 在 master 直接改      | `git status` 显示 HEAD 是 master 且有 diff                 | 开发前先 `git checkout -b feat-xxx aliyun/master`                       |
| 从 origin/master 起点 | feat 分支基点滞后于 aliyun 主线,PR 有冲突                  | 明确从 `aliyun/master` 拉分支                                           |
| 盲目复用旧 feat 分支  | 复用了落后于 aliyun/master 或工作区有残留的分支,污染新需求 | Phase 1 步骤 2 按充要条件严格校验,任一不满足必须新切                    |
| 未授权自动 push       | 用户还没审查就已推送                                       | 推送动作必须等用户说"推送"/"push"                                       |
| mock 漏改 CI 报错     | `*mockClient does not implement agentbay.Client`           | 接口变更后立即 `grep -r "type mock.*Client struct"` 全量补齐            |
| 残留 upstream         | `git remote -v` 仍有个人 fork                              | 执行 `git remote remove upstream`                                       |
| 缺失变更档案          | 需求完成但仓库无 Spec / CR 目录,后续无法追溯               | 启动需求时按 Phase 0 强制二选一建立档案                                 |
| Quest 跳过 Design     | 勾了 Execute Directly 导致 Spec 根本没生成                 | New Task 时确认未勾 Execute Directly,点 Download 写盘到 `.qoder/specs/` |
| 档案与代码脱节        | Spec 建好但没 `git add`,合并后档案丢失                     | 提交前检查 `git status` 确保 Spec/CR 文件已进入暂存区                   |
