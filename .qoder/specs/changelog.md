# Changelog 管理方案

## Context

当前 agentbay-cli 项目的 GitHub Releases（v0.1.0 ~ v0.2.8，共 19 个版本）全部使用硬编码的模板文本作为 release notes，没有任何实际变更信息。项目没有 CHANGELOG.md 文件，也没有任何自动化 changelog 生成工具。

**目标**：
1. 在仓库中维护 CHANGELOG.md，从 conventional commits 自动生成
2. 修复 GitHub Actions 工作流，让 Release 页面自动显示实际变更
3. 回填所有历史版本的变更信息

## 方案：git-cliff

使用 [git-cliff](https://git-cliff.org/) 作为 changelog 生成工具。

**为什么选 git-cliff 而非 git-chglog**：
- 能同时生成 CHANGELOG.md 和 GitHub Release notes（git-chglog 只能生成 CHANGELOG.md）
- `commit_parsers` 可过滤噪音提交（如重复的 Homebrew formula 自动提交）
- 配置是可读的 TOML，不需要写 Go template
- 更活跃的维护（git-chglog 已不太活跃）

## 实现步骤

### Step 1: 创建 `cliff.toml` 配置文件

**文件**: `cliff.toml`（新建）

核心配置：
- `commit_parsers`：将 `feat` → Features，`fix` → Bug Fixes，`docs` → Documentation，`refactor` → Refactoring
- 过滤规则：隐藏 `chore`、`ci`、以及 `feat: add official Homebrew formula for v*` 等自动提交
- 非 conventional commits 归入 "Other Changes"
- `remote = "github"`, `owner = "aliyun"`, `repo = "agentbay-cli"`
- `tag_pattern = "v[0-9].*"`

### Step 2: 添加 Makefile target

**文件**: `Makefile`（修改）

新增 targets：
- `make changelog` — 生成/更新 CHANGELOG.md
- `make changelog-next` — 预览下一版本的变更（输出到 stdout）
- `make changelog-install` — 安装 git-cliff（`brew install git-cliff`）

更新 `.PHONY` 和 `help` target。

### Step 3: 回填历史 CHANGELOG.md

在本地执行：
```bash
git-cliff -o CHANGELOG.md
```

git-cliff 会读取所有 tag（v0.1.0 ~ v0.2.8）和对应的 conventional commits，一次性生成完整的 CHANGELOG.md。

手动审阅后提交：`docs: add CHANGELOG.md with backfilled history for v0.1.0–v0.2.8`

### Step 4: 修改 GitHub Actions 工作流

**文件**: `.github/workflows/homebrew.yml`（修改）

#### 4a: 新增 "Install git-cliff" 步骤
在 "Checkout code" 之后添加，下载 git-cliff Linux binary 到 `/usr/local/bin/`。

#### 4b: 新增 "Generate Release Notes" 步骤
在 "Create GitHub Release" 之前添加：
```bash
# 生成当前版本的变更内容（不含 changelog 顶标题）
git-cliff --tag "v$VERSION" --unreleased --strip header > /tmp/release-changes.md

# 拼接变更内容 + 安装说明
cat /tmp/release-changes.md > /tmp/release-notes-full.md
cat <<'EOF' >> /tmp/release-notes-full.md

## Installation
Once merged into Homebrew core, install with:
\`\`\`bash
brew install agentbay
\`\`\`

## Manual Installation
Download the appropriate binary for your platform from the assets below.
EOF
```

#### 4c: 修改 "Create GitHub Release" 步骤
将硬编码的 `--notes "..."` 替换为：
```bash
gh release create "v$VERSION" \
  --title "Agentbay CLI v$VERSION" \
  --notes-file /tmp/release-notes-full.md \
  "${all_files[@]}"
```

#### 4d: 新增 "Update CHANGELOG.md" 步骤
在 Release 创建之后：
```bash
git-cliff -o CHANGELOG.md
git add CHANGELOG.md
git commit -m "docs: update CHANGELOG.md for v$VERSION" || echo "No changelog changes"
git push
```

### Step 5: 回填 GitHub Release Notes（可选）

**文件**: `scripts/backfill-release-notes.sh`（新建）

一次性脚本，遍历 v0.1.0 ~ v0.2.8 的所有 tag：
```bash
for tag in $(git tag -l 'v*' --sort=v:refname); do
  git-cliff --tag "$tag" --strip header > /tmp/notes-$tag.md
  # 添加安装说明
  cat <<EOF >> /tmp/notes-$tag.md

## Installation
...
EOF
  gh release edit "$tag" --notes-file /tmp/notes-$tag.md
done
```

### Step 6: 更新项目文档

**文件**: `README.md`（修改）— 添加 Changelog 链接段落

## 涉及文件清单

| 文件 | 操作 | 说明 |
|------|------|------|
| `cliff.toml` | 新建 | git-cliff 配置（commit parsers、模板、remote） |
| `CHANGELOG.md` | 新建 | 回填的完整变更日志（v0.1.0 ~ v0.2.8） |
| `Makefile` | 修改 | 添加 changelog / changelog-next / changelog-install targets |
| `.github/workflows/homebrew.yml` | 修改 | 添加 git-cliff 安装、自动生成 release notes、更新 CHANGELOG.md |
| `scripts/backfill-release-notes.sh` | 新建 | 回填历史 GitHub Release notes 的一次性脚本 |
| `README.md` | 修改 | 添加 Changelog 段落 |

## 验证方式

1. **本地验证 git-cliff 配置**：
   ```bash
   brew install git-cliff
   git-cliff -o CHANGELOG.md          # 生成完整 changelog
   git-cliff --unreleased --bump      # 预览下一个版本变更
   ```

2. **检查 CHANGELOG.md 内容**：
   - 确认 19 个版本都有对应段落
   - 确认 feat/fix/docs 分类正确
   - 确认 Homebrew formula 自动提交被过滤

3. **CI 验证**（push tag 触发）：
   - 确认 GitHub Release 包含实际变更信息而非硬编码模板
   - 确认 CHANGELOG.md 被自动更新并推送到仓库

4. **回填验证**：
   ```bash
   bash scripts/backfill-release-notes.sh --dry-run
   ```
   逐版本确认生成的 release notes 内容合理

---

## 使用说明

### GitHub Release（自动）

当你 push `v*` tag 触发 GitHub Actions 时，CI 会自动：
1. 安装 git-cliff
2. 从 git 历史生成该版本的变更内容
3. 用它作为 Release notes（替代之前的硬编码模板）

只要你按 conventional commits 格式写提交消息（`feat:`, `fix:`, `docs:` 等），Release 页面就会自动显示变更信息。

### CHANGELOG.md（半自动）

CHANGELOG.md 有两个更新时机：

- **CI 自动**：发布 Release 后，CI 会自动运行 `git-cliff -o CHANGELOG.md` 并推送到仓库。但这只有在 CI 流程触发时才会发生。
- **本地手动**：如果你想在发版之前先看/更新 CHANGELOG.md，需要手动运行 `make changelog`。

日常开发中不需要做任何额外操作——只要 commit message 遵循 conventional commits 格式，changelog 内容就会从 git 历史自动提取。

### GitHub Release 与 CHANGELOG.md 的关系

核心变更条目一致，但格式有差异：

| | GitHub Release | CHANGELOG.md |
|---|---|---|
| **范围** | 只包含当前版本的变更 | 包含所有历史版本 |
| **额外内容** | 末尾追加安装说明（Homebrew、手动下载） | 无 |
| **生成时机** | CI 创建 Release 时（单版本） | CI 更新时（全量重新生成） |

### 本地常用命令

```bash
make changelog          # 生成/更新 CHANGELOG.md
make changelog-next     # 预览下一版本的变更（输出到 stdout）
make changelog-install  # 安装 git-cliff（brew install git-cliff）
```

### 回填历史 Release Notes

```bash
scripts/backfill-release-notes.sh --dry-run          # 预览
scripts/backfill-release-notes.sh                     # 执行更新
scripts/backfill-release-notes.sh --tag v0.2.8        # 只更新单个版本
```
