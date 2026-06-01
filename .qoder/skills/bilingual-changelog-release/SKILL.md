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

默认使用方式 A（AI 对话翻译）：

1. 读取 `CHANGELOG.md` 顶部目标版本段。
2. 将 `### English` 下的分类与条目翻译到 `### 中文` 下。
3. 删除 `TRANSLATE_ME` 注释。
4. 如条目过细，可在不丢语义的前提下做粗粒度聚合。
5. 如需要 Highlights，在版本标题下添加 2-3 条用户视角亮点。

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
- [ ] 无残留 `TRANSLATE_ME` 或 `中文翻译待补充`
- [ ] 顶部仍保留空的 `## [Unreleased]`
- [ ] 无真实 UID、账号 ID 等敏感信息未脱敏
- [ ] PR 链接、commit 链接、author 不被翻译或破坏

可用检查命令：

```bash
grep -nE 'TRANSLATE_ME|中文翻译待补充' CHANGELOG.md
bash scripts/extract-changelog-section.sh X.Y.Z CHANGELOG.md >/tmp/release-notes.md
```

第一条应无输出；第二条应成功且 `/tmp/release-notes.md` 非空。

### Phase 4：提交、打 tag、推送（仅用户授权后）

用户明确要求后，按 Conventional Commits 提交：

```bash
git add CHANGELOG.md
git commit -m "docs: changelog for vX.Y.Z"
git tag vX.Y.Z
git push origin master vX.Y.Z
```

如果项目实际发布远程不是 `origin`，必须先向用户确认远程名；禁止猜测或直接 force push。

### Phase 5：workflow 发布验证

tag push 后，检查 `.github/workflows/homebrew.yml`：

- workflow 从 `CHANGELOG.md` 调用 `scripts/extract-changelog-section.sh` 抽取版本段
- 抽取失败应 fail-fast，并提示先跑 `make release-prep VERSION=X.Y.Z`
- workflow 不再调用 git-cliff 生成 release notes
- workflow 不再检查中文占位符
- workflow 不再 commit-back `CHANGELOG.md`

### Phase 6：已发布 Release 说明修订 / 历史回灌

若 release 已发布后要修订说明：

1. 先编辑 `CHANGELOG.md` 对应版本段。
2. 经用户授权后提交并推送 `CHANGELOG.md`。
3. 预览回灌：

```bash
bash scripts/backfill-release-notes.sh --dry-run --tag vX.Y.Z
```

4. 用户确认后执行：

```bash
bash scripts/backfill-release-notes.sh --tag vX.Y.Z
```

全量历史回灌必须先 `--dry-run`，再经用户确认后执行无参数脚本。

## ✅ 输出标准

完成发版准备时，必须向用户汇报：

- [ ] `CHANGELOG.md` 目标版本段已生成并翻译
- [ ] `TRANSLATE_ME` / 旧中文占位符已清理
- [ ] `extract-changelog-section.sh` 可成功抽取该版本段
- [ ] 已说明下一步是否需要用户授权 commit/tag/push
- [ ] 未自动执行未经授权的 git 或 gh 写操作

## 📚 参考

- [bilingual-changelog-proposal.md](../../docs/internal/bilingual-changelog-proposal.md)
- [release-checklist.md](../../docs/release-checklist.md)
- [development.md](../../rules/development.md)
- [update-cli-command-docs](../update-cli-command-docs/SKILL.md)
