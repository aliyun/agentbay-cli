---
name: bilingual-changelog-release
description: AgentBay CLI 双语 CHANGELOG 与 GitHub Release 发版 SOP。Use when preparing a release, updating bilingual CHANGELOG, translating release notes, creating tags, refreshing GitHub Release notes, or running make release-prep/backfill-release-notes.
---

# Bilingual Changelog Release

## 📋 职责

将 `docs/internal/bilingual-changelog-proposal.md` 的 v3 方案固化为可执行 SOP：

1. **本地生成**：发版前用 `make release-prep VERSION=X.Y.Z` 生成双语 CHANGELOG 版本段骨架。
2. **本地翻译**：在提交前完成 `### 中文` 段翻译，删除 `TRANSLATE_ME` 占位。
3. **单一事实源**：`CHANGELOG.md` 是 GitHub Release body 的唯一上游，workflow 只负责抽取版本段并发布。
4. **回灌修订**：已发布 release 的说明需要改时，先改 `CHANGELOG.md`，再用 `scripts/backfill-release-notes.sh` 同步。

本 skill 不负责 CLI 命令文档的参数、示例、README 表格更新；这些仍由 `update-cli-command-docs` 处理。

## 🎯 触发场景

当用户提出以下诉求时必须加载本 skill：

- “准备发版 / cut release / 发布 vX.Y.Z”
- “更新/翻译 CHANGELOG”且目标是发版版本段
- “根据 CHANGELOG 生成 GitHub Release notes”
- “运行 release-prep / backfill release notes”
- “修订已发布 release 的中英文说明”
- 修改 `cliff.toml`、`scripts/release-prep.sh`、`scripts/extract-changelog-section.sh`、`scripts/backfill-release-notes.sh`、`.github/workflows/homebrew.yml` 中与 release notes 相关的逻辑

## 🔒 核心原则

1. **CHANGELOG 单一事实源**：GitHub Release body 必须从 `CHANGELOG.md` 抽取，不直接手写 release body。
2. **生成/翻译左移到本地**：workflow 不做 git-cliff 生成、不做翻译、不 commit-back `CHANGELOG.md`。
3. **提交前可审阅**：中文翻译、Highlights、粗粒度聚合必须在 commit/tag 之前完成并人工 review。
4. **显式授权**：不得在用户明确授权前执行 `git add`、`git commit`、`git tag`、`git push`、`gh release edit`。
5. **不保留翻译缓存**：方式 A（AI 对话翻译）为默认路径，不新增 `.changelog-cache/`。

## 🚀 标准流程

### Phase 0：发版前确认

1. 确认版本号使用 `X.Y.Z`，执行命令中不带前缀 `v`。
2. 确认当前目标分支是发布主线（通常为 `master`）。
3. 确认工作区干净；如不干净，先提示用户提交或 stash，不能继续 release-prep。
4. 确认 tag 不存在：`vX.Y.Z` 本地和远程都不应存在。
5. 确认 `llms-full.txt` 已随英文对外文档同步：如果本次 release 包含 `README.md` 或 `docs/en/**` 变更，必须先执行 `bash scripts/build-llms-full.sh` 并提交生成结果；如新增 / 删除 / 重命名对外文档，还要确认 `llms.txt` 已同步。

### Phase 1：生成 CHANGELOG 版本段

运行：

```bash
make release-prep VERSION=X.Y.Z
```

该命令应完成：

- 校验工作区和 tag
- 拉取最新主线
- 用 git-cliff 生成 `## [X.Y.Z] - YYYY-MM-DD`
- 写入 `### English` 与 `### 中文` 双语结构
- 在中文段留下 `TRANSLATE_ME` 占位
- 重新插入空的 `## [Unreleased]` anchor
- 展示 `CHANGELOG.md` diff 和下一步提示

如 `git-cliff` 缺失，提示：

```bash
make changelog-install
```

### Phase 2：翻译与内容修订

默认使用方式 A（AI 对话翻译 + 结构对齐）：

1. 读取 `CHANGELOG.md` 顶部目标版本段。
2. 先整理 `### English`：按用户可感知的功能点做适度聚合，并按命令组归类；多改动命令组必须使用父条目 + 子条目结构。
3. 再将 `### English` 下的分类、命令组与条目逐项翻译到 `### 中文` 下；中文段的分类标题必须使用中文，不得保留 `Bug Fixes` / `Documentation` 等英文标题。
4. `### English` 与 `### 中文` 必须保持**结构对齐**：分类标题、命令组、子条目数量和顺序必须一一对应；中文段不得新增、遗漏或合并英文段中的独立条目。
5. 删除 `TRANSLATE_ME` 注释。
6. 如条目过细，可在不丢语义的前提下做粗粒度聚合，但必须**双语同步聚合**：英文聚合后中文使用同一粒度；中文拆行时英文同步拆行。
7. 对 CLI 命令相关条目，英文和中文都优先按命令组归类（如 `apikey`、`image`、`docker`、`skills`、`network`、`core/auth`）；无法归入命令组或属于全局能力 / 基础设施 / 发版流程的改动，英文可归为 `global`、`security/compliance`、`RAM permissions`、`release`，中文对应归为“全局”“安全合规”“RAM 权限”“发版”。
8. 如需要 Highlights，在版本标题下添加 2-3 条用户视角亮点。

中文分类标题建议：

| 英文分类           | 中文分类     |
| ------------------ | ------------ |
| `Breaking Changes` | `不兼容变更` |
| `Features`         | `功能`       |
| `Bug Fixes`        | `缺陷修复`   |
| `Performance`      | `性能优化`   |
| `Refactoring`      | `重构`       |
| `Documentation`    | `文档`       |
| `Security`         | `安全`       |
| `Other Changes`    | `其他变更`   |

术语约束：

- 保留英文不翻译：API Key、AK/SK、CLI、OSS、SDK、PR、Homebrew、OAuth、apikey、image、docker
- `image` 译为“镜像”（按上下文区分 OSS 镜像或 docker 镜像）
- `container` 译为“容器”
- `warmup` 译为“预热”
- `scope` 在权限语境译为“范围”，变量语境译为“作用域”
- `session` 译为“会话”
- `context` 译为“上下文”
- `flag` 译为“参数”或“选项”，不要译为“标志”
- 命令名、参数名、scope、PR 链接、author 保持英文/原样

### Phase 3：验证 CHANGELOG

必须检查：

- [ ] 存在 `## [X.Y.Z] - YYYY-MM-DD` 版本段
- [ ] 存在 `### English` 和 `### 中文`
- [ ] 中文段不是空内容
- [ ] `### English` 与 `### 中文` 结构对齐：分类数量/顺序一致、命令组数量/顺序一致、子条目数量/顺序一致
- [ ] 英文聚合和中文聚合粒度一致；不存在英文拆分但中文合并、或中文拆分但英文合并的情况
- [ ] 无残留 `TRANSLATE_ME` 或 `中文翻译待补充`
- [ ] 顶部仍保留空的 `## [Unreleased]`
- [ ] 无真实 UID、账号 ID 等敏感信息未脱敏
- [ ] 如本次 release 包含 `README.md` 或 `docs/en/**` 变更，`llms-full.txt` 已由 `bash scripts/build-llms-full.sh` 重新生成并提交
- [ ] 如本次 release 新增 / 删除 / 重命名对外文档，`llms.txt` 已同步
- [ ] PR 链接、commit 链接、author 不被翻译或破坏

可用检查命令：

```bash
grep -nE 'TRANSLATE_ME|中文翻译待补充' CHANGELOG.md
bash scripts/extract-changelog-section.sh X.Y.Z CHANGELOG.md >/tmp/release-notes.md
```

第一条应无输出；第二条应成功且 `/tmp/release-notes.md` 非空。

### Phase 4：提交、PR 合入与手动发布（仅用户授权后）

#### 本项目实际默认发布路径（必须优先遵循）

```text
本地 feat/dev-apikey
  ↓
push aliyun/feat-dev-apikey
  ↓
PR 合入 aliyun/master
  ↓
GitHub Actions 手动 Run workflow，输入 X.Y.Z（例如 0.4.0）
  ↓
workflow 从 master 的 CHANGELOG.md 抽取 X.Y.Z 段
  ↓
gh release create vX.Y.Z --target "$GITHUB_SHA"
  ↓
如果 vX.Y.Z tag 不存在，则自动创建 tag
  ↓
创建 GitHub Release
```

**关键约束**：本项目默认不要求在本地手动 `git tag`。tag 由 `.github/workflows/homebrew.yml` 中的 `gh release create "v$VERSION" --target "$GITHUB_SHA"` 在 Release 创建时自动创建，并绑定到本次 workflow checkout/build 的 `master` commit。

1. 在功能分支提交发版准备内容：

```bash
git add CHANGELOG.md
git commit -m "docs: changelog for vX.Y.Z"
git push <remote> <feature-branch>
```

2. 创建 PR 并合入上游 `aliyun/master`。合入前必须确认：
   - `CHANGELOG.md` 已包含 `## [X.Y.Z]` 版本段；
   - `### English` 与 `### 中文` 均已完成且结构对齐；
   - `TRANSLATE_ME` / `中文翻译待补充` 已清理。

3. PR 合入后，在 GitHub Actions 页面执行：
   - Actions → **Agentbay CLI Official Homebrew Release** → **Run workflow**
   - 选择 `master` 分支
   - 输入版本号 `X.Y.Z`（不带 `v` 前缀，如 `0.4.0`）

4. workflow 会执行 `gh release create "v$VERSION" --target "$GITHUB_SHA" ...`：
   - 如果 `vX.Y.Z` tag 不存在，`gh release create` 会自动创建；
   - `--target "$GITHUB_SHA"` 保证 tag 指向本次 workflow checkout/build 的 master commit；
   - GitHub Release body 从 `CHANGELOG.md` 中抽取 `## [X.Y.Z]` 段。

如果项目实际使用 tag-driven release（本地/CI 预先推送 `vX.Y.Z` tag），必须先向用户确认；禁止猜测远程名或直接 force push。

### Phase 5：workflow 发布验证

手动 Run workflow 后，检查 `.github/workflows/homebrew.yml`：

- workflow 从 `CHANGELOG.md` 调用 `scripts/extract-changelog-section.sh` 抽取目标版本段
- 抽取失败应 fail-fast，并提示先跑 `make release-prep VERSION=X.Y.Z`
- workflow 不再调用 git-cliff 生成 release notes
- workflow 不再检查中文占位符
- workflow 不再 commit-back `CHANGELOG.md`
- `gh release create` 必须带 `--target "$GITHUB_SHA"`，确保自动创建的 tag 指向本次构建 commit
- Release 创建成功后，确认 `vX.Y.Z` tag、Release body、构建资产均正确

### Phase 6：已发布 Release 说明修订 / 历史回灌

若 release 已发布后要修订说明（包括统一整理历史 CHANGELOG）：

> **回灌作用范围**：`scripts/backfill-release-notes.sh` 通过 `gh release edit --notes-file` 写入，**只替换 Release body（描述正文）**；不会修改 assets（二进制附件）、tag、title、draft / prerelease 状态。安全可重入。

> **前置检查（CHANGELOG 缺段）**：先跑 `bash scripts/backfill-release-notes.sh --dry-run` 查看输出。如果某个已发布 release 被报 `SKIP: CHANGELOG.md has no section for vX.Y.Z`，说明 CHANGELOG 里缺该版本段，必须先补齐再回灌：
>
> 1. 用 `git log v<prev>..vX.Y.Z --no-merges` 查看该版本实际包含的提交
> 2. 在 `CHANGELOG.md` 对应位置（按版本号倒序）插入 `## [X.Y.Z] - YYYY-MM-DD` 双语段，结构与同时期版本对齐
> 3. 所有缺段版本补齐后，再次跑 dry-run，确认 `Skipped: 0`，然后继续下面步骤 1

1. 先编辑 `CHANGELOG.md` 对应版本段；历史版本也必须保持双语结构，即每个版本包含 `### English` 与 `### 中文`，不得为了回灌只保留中文段。
2. 注意：仅修改并提交 `CHANGELOG.md` **不会自动更新 GitHub 上已存在的 Release body**；必须在变更推送到 GitHub 后执行 backfill，才能让历史 Release 说明同步。
3. 经用户授权后提交并推送 `CHANGELOG.md`。
4. 若本次同时准备新版本（如 `vX.Y.Z`）和历史回灌，推荐顺序是：先推送包含 CHANGELOG 的代码 → 先完成新版本 Release / tag workflow → 确认新 Release 正常 → 再对历史 Release 执行 backfill dry-run → 最后正式 backfill。
5. 预览单版本回灌：

```bash
bash scripts/backfill-release-notes.sh --dry-run --tag vX.Y.Z
```

6. 用户确认后执行单版本回灌：

```bash
bash scripts/backfill-release-notes.sh --tag vX.Y.Z
```

全量历史回灌必须先执行 `bash scripts/backfill-release-notes.sh --dry-run` 预览，再经用户确认后执行 `bash scripts/backfill-release-notes.sh` 无参数脚本；不得在未预览和未授权的情况下直接调用 `gh release edit`。

## 🧾 常用命令速查

> 复制前把 `X.Y.Z` / `vX.Y.Z` 替换成真实版本号。所有 `git add` / `git commit` / `git tag` / `git push` / `gh release edit` 相关动作都必须在用户明确授权后执行。

### 1. 生成新版本 CHANGELOG 骨架

```bash
make release-prep VERSION=X.Y.Z
```

### 2. 验证目标版本段可用于 GitHub Release

```bash
grep -nE 'TRANSLATE_ME|中文翻译待补充' CHANGELOG.md
bash scripts/extract-changelog-section.sh X.Y.Z CHANGELOG.md >/tmp/release-notes.md
test -s /tmp/release-notes.md
```

期望结果：第一条无输出；后两条成功退出。

### 3. 提交 CHANGELOG（功能分支）

适用于“先在功能分支准备 CHANGELOG，PR 合入发布主线后再手动执行 release workflow”的默认场景。

```bash
git add CHANGELOG.md
git commit -m "docs: changelog for vX.Y.Z"
git push <remote> <feature-branch>
```

### 4. 默认发布：PR 合入 master 后手动 Run workflow

```text
1. 将包含 CHANGELOG.md 的功能分支 PR 合入 upstream/master
2. GitHub → Actions → Agentbay CLI Official Homebrew Release → Run workflow
3. 选择 master 分支
4. 输入版本号 X.Y.Z（不带 v 前缀）
5. workflow 执行 gh release create "v$VERSION" --target "$GITHUB_SHA" ...
```

`gh release create` 会在 tag 不存在时自动创建 `vX.Y.Z`，并通过 `--target "$GITHUB_SHA"` 将 tag 绑定到本次 workflow 构建的 master commit。

### 5. 可选发布：预先创建 tag 并推送

仅在用户明确选择 tag-driven release 时使用：

```bash
git checkout master
git pull
git tag vX.Y.Z
git push <remote> vX.Y.Z
```

如发布远程不是当前明确目标，必须先确认远程名和分支名，不得猜测。

### 6. 已发布版本：预览单版本回灌

```bash
bash scripts/backfill-release-notes.sh --dry-run --tag vX.Y.Z
```

### 7. 已发布版本：正式单版本回灌

```bash
bash scripts/backfill-release-notes.sh --tag vX.Y.Z
```

### 8. 历史 Release：预览全量回灌

```bash
bash scripts/backfill-release-notes.sh --dry-run
```

### 9. 历史 Release：正式全量回灌

```bash
bash scripts/backfill-release-notes.sh
```

脚本启动时会显示 `Target repo` 并要求二次确认（输入 `y` 才继续）。若在 CI 或脚本中调用需要跳过提示，加 `--yes`：

```bash
bash scripts/backfill-release-notes.sh --yes
```

### 10. 推荐顺序：新版本 + 历史回灌同批处理

```bash
# 1. 推送包含 CHANGELOG.md 的代码到 GitHub 发布分支
# 2. 先完成新版本 Release / tag workflow，并确认新 Release 正常
# 3. 再预览历史回灌
bash scripts/backfill-release-notes.sh --dry-run

# 4. 人工确认 dry-run 输出无误后，再正式全量回灌
bash scripts/backfill-release-notes.sh
```

### 11. 检查 gh 登录状态

```bash
gh auth status
```

如果未登录，先执行：

```bash
gh auth login
```

## ✅ 输出标准

完成发版准备时，必须向用户汇报：

- [ ] `CHANGELOG.md` 目标版本段已生成并翻译
- [ ] `TRANSLATE_ME` / 旧中文占位符已清理
- [ ] `extract-changelog-section.sh` 可成功抽取该版本段
- [ ] 如涉及 `README.md` 或 `docs/en/**` 变更，`llms-full.txt` 已同步；如涉及对外文档结构变化，`llms.txt` 已检查
- [ ] 已说明下一步是否需要用户授权 commit/tag/push
- [ ] 未自动执行未经授权的 git 或 gh 写操作

## 📚 参考

- [bilingual-changelog-proposal.md](../../docs/internal/bilingual-changelog-proposal.md)
- [release-checklist.md](../../docs/internal/release-checklist.md)
- [development.md](../../rules/development.md)
- [update-cli-command-docs](../update-cli-command-docs/SKILL.md)
