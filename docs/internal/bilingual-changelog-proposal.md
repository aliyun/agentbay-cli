# Bilingual Changelog & GitHub Release 方案设计

> Status: **Proposal — pending approval**
> Author: 调研整理（基于 ant-design / wuying-agentbay-sdk / 主流工具横评）
> Last updated: 2026-06-01 (**v3**: 生成与翻译阶段 shift-left 到本地，workflow 改为只读投影)
>
> 历史版本：
> - v1（2026-05-29）：单文件双语 + git-cliff 模板增强 + workflow 内 LLM 翻译
> - v2（2026-05-29）：补充 push→release 端到端流程与一致性论证
> - **v3（2026-06-01）**：本次升级 — 把生成 / 翻译阶段从 workflow 移到本地，workflow 仅从 CHANGELOG 抽取版本段

## 1. 目标

一句话：让本地 `CHANGELOG.md` 和 GitHub Releases 完全一致，**双语**（英文 + 中文），每个版本**按用户视角合理分类总结**。

具体要求：

- 本地 `CHANGELOG.md` 与 GitHub Release body 同源同步，避免出现双方漂移
- 中英文双语，中文质量足够发布给国内用户阅读
- 分类语义对用户友好（不是 commit 类型，而是"我升级后能用什么 / 会被什么坑到"）
- 在已有 `git-cliff` 基础设施上演进，不推翻重做
- 维护成本控制在每次发版 5–10 分钟以内

## 2. 现状盘点

### 已有基础设施

| 文件 | 作用 |
|---|---|
| [cliff.toml](../../cliff.toml) | git-cliff 配置：分组规则（emoji 一级分类）、scope 归一、commit 过滤 |
| [Makefile:280-286](../../Makefile) | `make changelog` / `make changelog-next` |
| [scripts/backfill-release-notes.sh](../../scripts/backfill-release-notes.sh) | 把 git-cliff 输出推到 GitHub Release（`gh release edit`） |
| [CHANGELOG.md](../../CHANGELOG.md) | 历史版本列表，emoji 分组（🚀/🐞/📖/🛠/⚡️/🔒） |

### 痛点

1. **只有英文**（cliff.toml 模板里有 TODO 注释 `中文翻译待补充 / Add Chinese translation before release`，从未落地）
2. **粒度太细**：每个 commit 一行 bullet，用户读 release notes 时信息过载
3. **缺信息**：没有 PR 链接、没有 author、没有 Breaking Changes 突出
4. **scope 没用上**：cliff.toml 已经定义了 `scope_parsers`（apikey/image/docker/...），但模板里没渲染

### 当前 release 触发与生效路径

完整发版流程在 [.github/workflows/homebrew.yml](../../.github/workflows/homebrew.yml) 实现。git-cliff 在 workflow 中**真正生效的有 2 次**调用（第 50 行的 `--version` 仅做安装校验）：

| 行号 | 步骤 | 调用 | 流向 |
|---|---|---|---|
| [:478](../../.github/workflows/homebrew.yml) | "Generate Release Notes" | `git-cliff --tag v$VERSION --unreleased --strip header` | `/tmp/release-changes.md` → 给 `gh release create --notes-file` 用 |
| [:555](../../.github/workflows/homebrew.yml) | "Update CHANGELOG.md" | `git-cliff -o CHANGELOG.md` | 仓库根 `CHANGELOG.md` → 紧跟 :559–561 commit + push 回 master |

**触发链**：开发者本地 `git tag v0.4.0 && git push origin v0.4.0` → tag push 触发 workflow → 上述两步在同一次 workflow run 内串行执行。

**当前一致性是"约定一致"，不是"接口一致"**：两次 git-cliff 调用输入相同（同一 git pool + 同一份 `cliff.toml`）+ git-cliff 是确定性的 + 间隔秒级 → 输出相同。但两次调用之间没有共享中间产物，理论上手工编辑 `CHANGELOG.md` 或事后 `gh release edit` 都会破坏一致性。

**关键设计漏洞**：cliff.toml footer 的 `中文翻译待补充` 占位符**永远落不了地**——开发者手编 `CHANGELOG.md` 会被下次发版的 `git-cliff -o` 全量覆盖，当前方案根本没有"中文翻译"的入口。

## 3. 主流做法横评

### 3.1 同 monorepo 上游：[wuying-agentbay-sdk](https://github.com/agentbay-ai/wuying-agentbay-sdk)

- **纯手工**，无任何工具配置（无 cliff.toml / release-please / semantic-release / changesets）
- 单人维护，每次发版手写 CHANGELOG → 手工 `gh release create` 粘贴对应段
- 内容风格亮点（值得借鉴）：
  - **[Keep a Changelog](https://keepachangelog.com)** 标准分类：`Added` / `Changed` / `Fixed` / `Deprecated` / `Removed` / `Security` / `Breaking Changes`（**用户视角**）
  - **scope 加粗前缀**：`**Java**:` `**All SDKs**:` `**Context sync**:`
  - **粗粒度聚合**：一个 bullet = 一个用户可感知的 feature，子 bullet 展开实现细节
  - **Breaking Changes 独立成节**，第一时间提示
- 短板（不要学）：
  - 多份 CHANGELOG 不同步（顶层 vs `python/CHANGELOG.md` 差一版）
  - 完全英文
  - 完全手工，不可扩展

### 3.2 [Ant Design](https://github.com/ant-design/ant-design)

- 双语：维护两份 [CHANGELOG.en-US.md](https://github.com/ant-design/ant-design/blob/master/CHANGELOG.en-US.md) + [CHANGELOG.zh-CN.md](https://github.com/ant-design/ant-design/blob/master/CHANGELOG.zh-CN.md)
- 强制 PR 双语：每个 PR 的 description 必须填两段 changelog（🇨🇳/🇺🇸），机器人聚合
- 行格式：`emoji + 描述 + PR 链接 + @author`，**按组件名二级聚合**
- 适用场景：贡献者多、UI 组件库（按组件分组天然合理）
- **不适合本项目**：贡献者主要是内部，强制 PR 双语成本高

### 3.3 工具横向对比

| 工具 | 触发方式 | 双语支持 | 适合度 |
|---|---|---|---|
| **git-cliff**（现用） | 从 conventional commits 渲染 Tera 模板 | 模板可多份，但内容来自 commit | ✅ 已有基础 |
| release-please (Google) | PR 驱动，自动开 Release PR | 单语 | 改动大 |
| semantic-release (JS) | 全自动版本+发布 | 单语 | Go 项目少用 |
| changesets (Vue/Tauri) | 每个 PR 加 fragment 文件 | fragment 可写双语 | 改动大 |
| towncrier (Python) | 同 changesets | 同上 | 改动大 |

**结论**：保留 git-cliff，叠加 LLM 翻译层。这是改动最小、效果最接近上游的路径。**v3 进一步把翻译阶段下沉到本地**（业内称 shift-left changelog），CHANGELOG.md 成为唯一上游事实，workflow 退化为只读投影。Tauri / Rust ecosystem 多采用此风格。

## 4. 推荐方案（v3）

### 4.1 总体架构

```
本地（开发者）                                    GitHub
─────────────                                    ──────
make release-prep VERSION=0.4.0
  ├─ git-cliff --tag v0.4.0 --unreleased  → 英文片段
  ├─ AI 翻译（方式 A/B/C，见 4.5）        → 中文片段
  ├─ 拼装双语 + Highlights
  └─ prepend 写入 CHANGELOG.md 顶部

vim CHANGELOG.md          # 审核 / 修订（v3 关键优势）

git add CHANGELOG.md
git commit -m "docs: changelog for v0.4.0"
git tag v0.4.0
git push origin master v0.4.0  ──┐
                                 │
                                 └─→ workflow（homebrew.yml）
                                     ① 编译 binary（不变）
                                     ② extract-changelog-section.sh "v0.4.0" CHANGELOG.md
                                        → /tmp/release-notes.md
                                     ③ Fail-fast：抠不到段就 exit 1
                                        提示"先跑 make release-prep"
                                     ④ 拼 Installation 段
                                     ⑤ gh release create --notes-file ...
                                     ⑥ ❌ 不再 commit-back CHANGELOG.md
```

**核心原则**：
1. **CHANGELOG.md 是单一上游事实来源**，GitHub Release body 是它的投影
2. **生成 / 翻译只在本地发生**，开发者在 commit 之前可以 review、改翻译、调措辞
3. **workflow 退化为只读**：只抽取版本段、拼 Installation、`gh release create`，不再生成、不再翻译、不再写回仓库

### 4.2 CHANGELOG 格式

每个版本的 body 长这样（**单文件双语，上下分块**）：

```markdown
## [0.3.0] - 2026-05-22

> ✨ Highlights / 本次亮点
> - API Key 命令体系增强：新增 describe-key-content / 支持 --api-key-id 等参数
> - 文档结构按 en/zh 重组，补齐双语翻译

### English

#### 🚀 Features
- **apikey**: Add `describe-key-content` command ([#123](https://...) by @user)
- **apikey**: Support `--api-key-id` param for relevant commands ([#120](https://...))

#### 📖 Documentation
- Reorganize documentation into en/zh structure ([#118](https://...))
- Add RAM permission requirements for apikey commands ([#115](https://...))

#### 🛠 Refactoring
- Replace OAuth login hints with AK/SK env var guidance ([#112](https://...))

### 中文

#### 🚀 新功能
- **apikey**: 新增 `describe-key-content` 子命令（[#123](https://...) by @user）
- **apikey**: 相关命令支持 `--api-key-id` 参数（[#120](https://...)）

#### 📖 文档
- 文档结构按 en/zh 重组（[#118](https://...)）
- 补充 apikey 命令的 RAM 权限要求（[#115](https://...)）

#### 🛠 重构
- 登录提示从 OAuth 改为 AK/SK 环境变量指引（[#112](https://...)）

---
```

**为什么不分两个文件？**
- 单文件减少同步成本（避免上游 SDK 那种"双 CHANGELOG 落后一版"的尴尬）
- GitHub Release body 直接复用整段
- 文档站如果未来需要按语言路由再切分也来得及（content 不变只是渲染层）

### 4.3 分类策略（融合 emoji + Keep a Changelog 语义）

保留你现在的 emoji，但**改名标题对齐用户视角**，并新增 Breaking Changes：

| emoji | 英文标题 | 中文标题 | 触发条件 |
|---|---|---|---|
| ⚠️ | Breaking Changes | 不兼容变更 | commit 含 `!` 或 footer `BREAKING CHANGE:` |
| 🚀 | Features | 新功能 | `feat:` |
| 🐞 | Bug Fixes | 问题修复 | `fix:` |
| ⚡️ | Performance | 性能优化 | `perf:` |
| 🛠 | Refactoring | 重构 | `refactor:` |
| 📖 | Documentation | 文档 | `docs:` |
| 🔒 | Security | 安全 | body 含 security |
| 📦 | Other Changes | 其他变更 | 兜底 |

**额外规则**：

1. **Breaking Changes 永远排第一**（用户最该看到的）
2. **scope 加粗前缀**：`**apikey**: ...`（借鉴上游 SDK 风格，比按 scope 二级分组更紧凑）
3. **每条尾部加 PR 链接和 author**：`([#123](url) by @user)`，git-cliff 支持 `{{ commit.github.pr_number }}` `{{ commit.github.username }}`
4. **过滤更激进**：
   - 已过滤：chore/ci/build/style/test、merge、内部更新（你已经做了）
   - 建议加：单纯的 typo 修复（`docs.*typo`）、版本号 bump 提交

### 4.4 Highlights / 亮点段（可选）

每版顶部一段 blockquote（2–3 行），让用户 5 秒抓住核心。

**生成方式**（待确认，见第 7 节）：
- 选项 A：LLM 自动总结（翻译那步顺手做）
- 选项 B：人工填，无则不显示
- 选项 C：不要

### 4.5 翻译方式（三选一，开发者本机）

v3 的翻译发生在本地，开发者可任选：

| 方式 | 实现 | 适合场景 | 是否需要 API key |
|---|---|---|---|
| **A. Claude Code 对话翻译**（推荐） | `make release-prep` 输出英文片段到 `/tmp/release-en.md`，开发者在 Claude Code 里说"翻译这个"，AI 输出中文，写入 CHANGELOG.md | 发版频率低、追求质量、可对话调整、做粗粒度聚合 | ❌ 不需要 |
| **B. 本机脚本调 Claude API** | `make release-prep` 自动调 API，需开发者 shell 有 `ANTHROPIC_API_KEY` | 发版频繁、希望全自动、需要批量回灌历史 | ✅ 开发者本机 env |
| **C. 混合** | 默认调 API，无 key 时 fallback 到方式 A 的提示 | 多人发版、能力不一 | ⚠️ 可选 |

**推荐方式 A**，理由：
- 不需要任何凭据管理（既不需要 GitHub secret，也不需要本机 env）
- AI 翻译质量最高，可对话调整、可顺手做粗粒度聚合 + Highlights
- 与 Claude Code 已有工作流天然契合

#### 术语表（无论用哪种方式都要约束 AI）

```
保留英文不翻译：API Key、AK/SK、CLI、OSS、SDK、PR、Homebrew、OAuth、apikey、image、docker
统一译法：
  image → 镜像（指 OSS 镜像或 docker 镜像，看上下文）
  container → 容器
  warmup → 预热
  scope → 范围（权限语境）/ 作用域（变量语境）
  session → 会话
  context → 上下文
  flag → 参数 / 选项（不要用"标志"）
  feat / fix / docs / refactor → 不翻译类别名，只翻译描述本体
```

方式 A：术语表写在 Claude Code 的系统提示或对话 context 里。
方式 B/C：术语表硬编码在脚本的 prompt template 里。

#### 翻译缓存（v3 改为可选）

v3 下 CHANGELOG.md 本身就是事实来源，**不再强制需要 cache**。但仍可保留：

| 是否保留 cache | 利 | 弊 |
|---|---|---|
| **保留**（推荐方式 B/C 时） | 重生成不重复翻译；改 cliff.toml 模板回灌历史时省 API 调用 | 多一个目录要维护；rebase/squash 后可能产生孤儿条目 |
| **不保留**（推荐方式 A 时） | 简单；CHANGELOG.md 唯一上游 | 重生成历史需要再翻一遍；但方式 A 本来也是一次性手工 |

`.changelog-cache/` 目录如保留，提交进仓库（不是 gitignore），按 commit SHA 索引：

```
.changelog-cache/
├── abc123def.en.md  # 英文片段
└── abc123def.zh.md  # 译文
```

#### 译文人工覆盖

无论用哪种方式，发布前 review 时直接编辑 CHANGELOG.md 即可（v3 优势：commit 之前你都能看到、能改）。

### 4.6 本地生成管线（v3）

#### 新增脚本

**`scripts/release-prep.sh`**（本地调用，发版前必跑）
```
用法: bash scripts/release-prep.sh <VERSION>      # e.g. release-prep.sh 0.4.0
行为:
  1. 校验工作树干净 + 当前在 master 分支
  2. git pull --ff-only origin master
  3. 校验 v<VERSION> tag 不存在
  4. git-cliff --tag v<VERSION> --unreleased --strip header > /tmp/release-en.md
  5. 翻译（按方式 A/B/C）→ /tmp/release-zh.md
  6. 拼装双语片段 + Highlights → /tmp/release-<VERSION>.md
  7. 把片段 prepend 到 CHANGELOG.md（替换 [Unreleased] 段）
  8. 显示 git diff CHANGELOG.md，提示开发者审核
  9. 输出下一步指引：
     git add CHANGELOG.md
     git commit -m "docs: changelog for v<VERSION>"
     git tag v<VERSION>
     git push origin master v<VERSION>
```

**`scripts/extract-changelog-section.sh`**（本地 + workflow 共用）
```
用法: bash scripts/extract-changelog-section.sh <VERSION> <CHANGELOG_PATH>
行为: 从 CHANGELOG.md 抠出 ## [<VERSION>] 那一段（到下一个 ## 之前）
       输出到 stdout
       抠不到则 exit 1
实现: awk 一行解决
```

#### 改造现有

**`Makefile`**

```makefile
.PHONY: release-prep changelog-preview

release-prep:                                 # 发版前必跑
	@if [ -z "$(VERSION)" ]; then echo "Usage: make release-prep VERSION=0.4.0"; exit 1; fi
	@bash scripts/release-prep.sh $(VERSION)

changelog-preview:                            # 不写文件，只预览
	@git-cliff --unreleased --strip header

# 旧的 changelog target 保留，仅用于全量重生（紧急修复历史 CHANGELOG 时用）
changelog:
	@git-cliff -o CHANGELOG.md
```

**`scripts/backfill-release-notes.sh`**

逻辑反转：从"调 git-cliff 生成 → 推到 release"变为"从 CHANGELOG.md extract → 推到 release"。这样回灌历史 release 也走"CHANGELOG 是上游"的同一路径。

**`cliff.toml`**

模板增强（同 v2 不变）：
- 加 PR 链接和 author 渲染
- 加 Breaking Changes 独立分组
- scope 加粗前缀
- 标题改成产品化措辞

### 4.7 Push → Release 端到端流程与一致性保证（v3）

#### 时序

```
本地（开发者）                              GitHub Actions (homebrew.yml)
─────────────                              ─────────────────────────────
make release-prep VERSION=0.4.0
  ├─ 校验工作树干净 + 在 master
  ├─ git pull --ff-only
  ├─ git-cliff --tag v0.4.0 --unreleased  → 英文片段
  ├─ AI 翻译（方式 A/B/C）                 → 中文片段
  ├─ 拼装双语 + Highlights
  └─ prepend 到 CHANGELOG.md，显示 diff

vim CHANGELOG.md          # ← v3 关键优势：commit 前可改

git add CHANGELOG.md
git commit -m "docs: changelog for v0.4.0"
git tag v0.4.0
git push origin master v0.4.0  ──┐
                                 │
                                 └─→ ① 触发 workflow（push v* tag）
                                     ② 编译四平台 binary + bottle（不变）
                                     ③ Verify CHANGELOG section（:478 新逻辑）
                                        bash scripts/extract-changelog-section.sh "v0.4.0" CHANGELOG.md
                                          → /tmp/release-changes.md
                                        Fail-fast：empty → exit 1，提示开发者
                                        "先在本地跑 make release-prep VERSION=0.4.0"
                                     ④ 拼 Installation 段
                                          → /tmp/release-notes-full.md
                                     ⑤ gh release create --notes-file ...   ──→ GitHub Release v0.4.0
                                     ⑥ ❌ 删除 commit-back CHANGELOG.md 步骤
                                        （CHANGELOG 已经在开发者推送的 commit 里）
                                     ⑦ Homebrew formula / bottle / Pages（不变）
```

发版者实际感受：
- 多了一步本地命令（`make release-prep`），但**换来在 commit 前预览 + 修订翻译的能力**
- workflow 不再 commit 回仓库（master 上多余的"自动 commit"消失，git log 更干净）
- 不再看到"中文翻译占位符"警告（v2 已解决，v3 同样）

#### Workflow 具体改动点

只动 [.github/workflows/homebrew.yml](../../.github/workflows/homebrew.yml)，**改动比 v2 更少**：

| 行号 | 现状 | v3 新方案 |
|---|---|---|
| [:44–51](../../.github/workflows/homebrew.yml) Install git-cliff | 安装 git-cliff | **整步删除**（workflow 不再调 git-cliff） |
| [:478](../../.github/workflows/homebrew.yml) Generate Release Notes | `git-cliff --tag --unreleased --strip header > /tmp/release-changes.md` | `bash scripts/extract-changelog-section.sh "v$VERSION" CHANGELOG.md > /tmp/release-changes.md` + 紧跟 fail-fast 校验 |
| [:499–514](../../.github/workflows/homebrew.yml) Check Chinese Translation | 检查占位符 | **整步删除** |
| [:552–561](../../.github/workflows/homebrew.yml) Update CHANGELOG.md + commit-back | 跑 `git-cliff -o` 然后 commit + push | **整步删除** |
| Secrets | — | **不需要新增** secret（无 workflow 内 API 调用） |

发布脚本 [scripts/release-to-oss.sh](../../scripts/release-to-oss.sh) 仍不需要动。

> 对比 v2：v2 的 workflow 改动是"换两处脚本 + 加 secret + 加缓存提交"；v3 是"删两步 + 改一处提取逻辑 + 加 fail-fast"。v3 更倾向于让 workflow 退化为纯发布动作。

#### 三个核心入口（v3）

| 入口 | 当前 | v3 |
|---|---|---|
| **发版触发** | `git push v* tag` | 同；但**之前**必须先 `make release-prep` + commit |
| **CHANGELOG.md 写入** | workflow `:555` 自动 | **`make release-prep` 时由开发者写入并 commit**（workflow 不再写） |
| **中文翻译落地** | ❌ 不存在 | **本地 `make release-prep` 即时翻译，commit 之前可改** |

#### 一致性论证：CHANGELOG.md 是单一上游

```
                          CHANGELOG.md（仓库 master HEAD）
                          ↑
                          │ 由 make release-prep 写入
                          │
              ┌───────────┴───────────┐
              │                       │
    extract-changelog-section.sh   开发者本地 cat / 阅读
              │
              ↓
    GitHub Release body（投影）
```

| 维度 | v2（同脚本 + 同缓存） | **v3（CHANGELOG 单一上游）** |
|---|---|---|
| 单版本片段（Release body）来源 | `generate-bilingual-release.sh` 渲染 | **从 CHANGELOG.md extract** |
| CHANGELOG.md 来源 | workflow 内 `generate-bilingual-changelog.sh` | **本地 release-prep 写入，仓库即上游** |
| 一致性支撑 | "同脚本 + 同缓存" | **结构上不可能漂移**（Release 是 CHANGELOG 的子集投影） |
| 手工修订 release body 是否破坏一致性 | ⚠️ 需要走回流流程 | ✅ 拒绝直接改 release body；改源 → workflow_dispatch 重抽即可 |
| 翻译质量何时可见 | workflow 跑完，已发出去 | ✅ commit 前可见，可 vim 改 |

**结构性优势**：v3 不再需要 cache、也不再需要"回流流程"——因为 CHANGELOG.md 本身就是事实，Release 永远是它的子集投影。任何对 release 的修订**必须**先改 CHANGELOG，再用 backfill 抽一遍。

#### 发版后修订翻译的回流流程（v3 简化版）

```bash
# 1. 改 CHANGELOG.md 里 v0.4.0 那段（直接编辑）
vim CHANGELOG.md

# 2. commit & push
git add CHANGELOG.md
git commit -m "docs: refine zh translation for v0.4.0"
git push

# 3. 同步刷新已发 release 的 body（从 CHANGELOG 抽取）
bash scripts/backfill-release-notes.sh --tag v0.4.0
```

少一步（无需改缓存），更直观。

#### 失败模式（v3）

| 失败 | 行为 | 处理 |
|---|---|---|
| **开发者忘了 `make release-prep` 直接打 tag** | workflow `extract-changelog-section.sh` 抠不到段 → exit 1 | 删 tag → 本地补 release-prep → 重新打 tag、push |
| **CHANGELOG.md 节段格式被人手破坏**（如标题不再是 `## [0.4.0]`） | 同上，extract 失败 | 修 CHANGELOG.md 格式后 push，workflow_dispatch 重跑 |
| **AI 翻译质量差** | release-prep 输出后开发者已能看到 | commit 之前直接 vim 改；不进 git history |
| **Claude Code 不可用 / 本机无 API key** | release-prep 失败或退化为英文版 | 方式 A 时手工补翻译；方式 B 时切方式 A |
| **历史 backfill 误刷** | `scripts/backfill-release-notes.sh` 已支持 `--dry-run` 和 `--tag` 单版本 | 不变 |
| **cliff.toml 模板改坏** | 下次 release-prep 才会发现 | 改模板的 PR 必须本地跑 `make changelog-preview` 验证再合 |

### 4.8 历史 release backfill

`scripts/backfill-release-notes.sh` 复用即可（已经支持 `--tag` / `--dry-run` / 全量更新）。**v3 下逻辑反转**：脚本内部不再调 `git-cliff`，而是先确保 CHANGELOG.md 包含目标版本段，然后用 `extract-changelog-section.sh` 抠出来推到 release。

回灌历史的两个阶段：

1. **CHANGELOG.md 历史段补齐**：用 `make release-prep` 思路对 v0.1.0 ~ v0.3.0 共约 15 个版本逐个生成双语段并 prepend（一次性手工 + AI 翻译），最终一次 commit 把整份补齐版本的 CHANGELOG.md 推上去。
2. **release body 同步**：`bash scripts/backfill-release-notes.sh`（无参数，全量回灌）从 CHANGELOG.md extract 每个版本段，推到对应 release。

预估：方式 A（Claude Code 翻译）下，15 个版本约一两小时人工 + AI 协作；方式 B（API）下约 5 分钟，成本 < $0.5。

### 4.9 Commit 规范建议

git-cliff 的自动分类、scope 提取、Breaking Changes 识别**强依赖** [Conventional Commits](https://www.conventionalcommits.org/)。本项目当前 commit 规范度已经很高（最近 100 个 commit 仅 1 条不符合，且已被 cliff.toml 兜底规则覆盖），本节把这种"事实标准"明文化，作为后续团队成员的参考。

#### Commit message 格式

```
<type>[(<scope>)][!]: <subject>

[<body>]

[<footer>]
```

- `<type>` —— **必填**，下文表 1
- `<scope>` —— **可选但推荐**，下文表 2
- `!` —— **可选**，标记 Breaking Change（与 footer `BREAKING CHANGE:` 二选一即可）
- `<subject>` —— 必填，英文，祈使句，首字母小写，行末不加句号；推荐 ≤ 72 字符
- `<body>` —— 可选，多行描述「为什么/怎么实现」；空一行后写
- `<footer>` —— 可选，`BREAKING CHANGE: 详细说明` 或 `Closes #123`

**例子**：

```
feat(apikey): add describe-key-content command
fix(image): handle empty registry argument gracefully
refactor(core)!: rename Token to ApiKey

The old name conflicted with OAuth concepts and confused new users.
This is a breaking change for SDK consumers.

BREAKING CHANGE: Token type renamed to ApiKey across all command groups.
```

#### 表 1：合法的 type

| type | CHANGELOG 分类 | 何时用 | 是否进 changelog |
|---|---|---|---|
| `feat` | 🚀 Features / 新功能 | 用户可感知的新功能、新命令、新参数 | ✅ |
| `fix` | 🐞 Bug Fixes / 问题修复 | 修 bug | ✅ |
| `perf` | ⚡️ Performance / 性能优化 | 性能改进 | ✅ |
| `refactor` | 🛠 Refactoring / 重构 | 不改变行为的代码重构 | ✅ |
| `docs` | 📖 Documentation / 文档 | 文档变更（README、注释、guide） | ✅ |
| `style` | — | 代码格式化、空白、分号（不影响行为） | ❌ 过滤 |
| `test` | — | 加测试或调测试 | ❌ 过滤 |
| `chore` | — | 依赖升级、版本号 bump、构建脚本 | ❌ 过滤 |
| `ci` | — | CI 配置变更（GitHub Actions / OSS pipeline） | ❌ 过滤 |
| `build` | — | 构建系统（Makefile、Dockerfile） | ❌ 过滤 |
| `revert` | 📦 Other Changes / 其他变更 | 回退之前的 commit | ✅ 兜底 |

#### 表 2：本项目固定 scope（按 CLI 命令组）

[cliff.toml:78-86](../../cliff.toml) 已定义归一化规则。提交时优先用左列**规范名**，写成右列任一别名也会被自动归一：

| 规范名 | 接受的别名 | 含义 |
|---|---|---|
| `apikey` | `api-key`, `apikey` | API Key 命令组 |
| `image` | `img`, `image` | OSS 镜像命令组 |
| `network` | `net`, `network` | 网络命令组 |
| `skills` | `skill`, `skills` | Skills 命令组 |
| `docker` | `docker` | Docker 命令组 |
| `core` | `core`, `auth`, `login`, `logout` | 核心 / 认证 |
| `client` | `client`, `sdk` | SDK 客户端层 |

跨命令组的全局变更可以**省略 scope**：

```
feat: add --output json flag to all list commands
docs: restructure README with bilingual support
```

#### Breaking Changes 标记

两种等效写法（任选其一，**推荐 `!`**，更醒目）：

```
# 方式 A：subject 行加 !
refactor(apikey)!: rename --token to --api-key

# 方式 B：footer 写 BREAKING CHANGE:
refactor(apikey): rename --token to --api-key

BREAKING CHANGE: --token flag renamed to --api-key for consistency.
```

git-cliff 都能识别，CHANGELOG 里都会归到 ⚠️ Breaking Changes / 不兼容变更段。

#### Subject 行写作建议

| ❌ 不推荐 | ✅ 推荐 | 原因 |
|---|---|---|
| `feat: 添加列表命令` | `feat(apikey): add list command` | 英文一致性（与上游 SDK 对齐） |
| `feat: added the new list command for apikey.` | `feat(apikey): add list command` | 祈使句 + 首字母小写 + 无句号 |
| `fix: bugs` | `fix(image): handle empty registry argument` | 具体说"修了什么"，不是"修了 bug" |
| `feat: 增加新的子命令、修复登录问题、更新文档` | 拆成 3 个 commit | 一个 commit 一件事 |

#### 推荐工作流：PR squash merge

> 目前你项目已经在用，但写明确保未来不退化。

- PR 内可以有任意中间 commit（包括 `wip`, `address review` 等）
- merge 时用 **squash merge**，PR title 即最终 commit message
- PR title 必须符合 conventional commits 格式
- 这样所有进 master 的 commit 都规范，无需 commitlint

#### 可选强化：CI 校验 PR title

如果未来贡献者增多，加一个 GitHub Action 自动检查 PR title 格式：

```yaml
# .github/workflows/lint-pr-title.yml
name: Lint PR Title
on:
  pull_request:
    types: [opened, edited, synchronize]
jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: amannn/action-semantic-pull-request@v5
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

零维护成本。当前阶段（贡献者少 + PR review 把关）可不加。

#### 不规范 commit 的兜底

万一历史中混入了不规范 commit，cliff.toml 里两条机制兜底：

1. **正则识别口语化前缀**（[cliff.toml:65-66](../../cliff.toml)）：`^add ` `^support ` 也算 features
2. **万能兜底分组**（[cliff.toml:69](../../cliff.toml)）：`^.*` → 📦 Other Changes

**新规范不规范都不会让 git-cliff 崩溃**，但长期看仍建议保持规范。

#### 当前合规度（参考）

最近 100 个 commit：

```
  ~36 feat / docs / fix / refactor / perf  ✅ 标准 conventional
  ~14 chore / ci / build / style / test    ✅ 标准 + 自动过滤
  ~ 7 Merge ...                             ✅ 已被 cliff 过滤
   1 support sudo docker                    ⚠️ 已被 ^support 兜底
─────────
   0 真正"漏网"的不规范 commit
```

继续保持即可，无需任何改造。

## 5. 实施步骤（v3）

| # | 步骤 | 工作量 | 依赖 |
|---|---|---|---|
| 1 | 改 [cliff.toml](../../cliff.toml)：PR链接、author、Breaking 独立、scope 前缀；本地 dry-run 渲染 v0.3.0 验证英文模板 | 30 min | 无 |
| 2 | 写 `scripts/extract-changelog-section.sh`（awk 从 CHANGELOG.md 抠版本段） | 15 min | 无 |
| 3 | 写 `scripts/release-prep.sh`（git-cliff 英文 + 翻译占位 + prepend CHANGELOG） | 45 min | 步骤 1, 2 |
| 4 | 实现翻译方式：方式 A 仅写 prompt 模板和术语表文档（无脚本），或方式 B 写 `translate-changelog.sh` 调 Claude API | 30 min（A）/ 1 h（B） | 步骤 3 |
| 5 | 改 [Makefile](../../Makefile)：加 `release-prep` / `changelog-preview` target | 15 min | 步骤 3 |
| 6 | 改 [.github/workflows/homebrew.yml](../../.github/workflows/homebrew.yml)：删 `:44–51` install git-cliff、`:478` 改为 extract + fail-fast、删 `:499–514` 占位符检查、删 `:552–561` 整个 commit-back 步骤 | 30 min | 步骤 2 |
| 7 | 改 [scripts/backfill-release-notes.sh](../../scripts/backfill-release-notes.sh)：从 git-cliff 改为 extract-changelog-section | 20 min | 步骤 2 |
| 8 | 在 v0.3.0 上端到端验证：本地 release-prep → review → 模拟 push tag → workflow_dispatch 测试运行 | 30 min | 步骤 1–7 |
| 9 | 历史 backfill：补齐 CHANGELOG.md v0.1.0 ~ v0.3.0 中文段（方式 A 配 AI 协作）→ commit → 全量刷 release body | 1–2 h | 步骤 8 |
| 10 | 更新 [docs/release-checklist.md](release-checklist.md)：第 20 行 pre-flight 加 `make release-prep`；新增"翻译修订流程"小节 | 20 min | 步骤 9 |
| | **合计** | **半天到一天**（方式 A）/ **大半天**（方式 B） | |

## 6. 风险 & 回滚（v3）

| 风险 | 缓解 |
|---|---|
| **开发者忘了 `make release-prep` 直接打 tag** | workflow fail-fast；错误信息直接告诉开发者下一步命令；删 tag → 补流程 → 重打 tag |
| **CHANGELOG.md 节段格式被人手破坏** | extract 脚本失败提示具体行号；release-checklist 写明节段标题格式约定 |
| AI 翻译质量差 | commit 之前开发者可见可改；本地工作流的天然优势 |
| 模板改坏导致渲染异常 | 改 cliff.toml 的 PR 必须本地 `make changelog-preview` dry-run；workflow 不再依赖 cliff.toml |
| 历史 backfill 误刷 release | `--tag` 单版本 + `--dry-run` 双重保护 |
| 译文风格漂移（不同版本用词不一致） | 术语表（无论方式 A/B 都用）；首次回灌后形成"内部语料"基线 |
| 本机无 ANTHROPIC_API_KEY 又选了方式 B | release-prep.sh 在调 API 前显式 check key，缺则提示切换方式 A |
| 多人同时发版 | 罕见；按 git 正常冲突解决（push tag 时 master 已被先发者的 changelog commit 占了，rebase 即可） |

回滚：
- `git revert` cliff.toml 与 workflow 改动
- 删除 `scripts/release-prep.sh` `scripts/extract-changelog-section.sh`（以及方式 B 的 `translate-changelog.sh`）
- 还原 [.github/workflows/homebrew.yml](../../.github/workflows/homebrew.yml) 的 install git-cliff / generate / commit-back 步骤
- 历史 release body 用旧 backfill 脚本重刷英文版

## 7. 待确认的关键决策点（v3）

发起实施前，需要你拍板这几个：

| # | 决策点 | 我的建议 | 你的选择 |
|---|---|---|---|
| 1 | CHANGELOG 排版（单文件双语 / 内联双语 / 双文件） | **单文件 English/中文 上下分块** | ☐ |
| 2 | Highlights 段是否要 | 在 release-prep 时由 AI 顺手生成，开发者 review 时调整 | ☐ |
| 3 | 翻译方式（A: Claude Code 对话 / B: 本机调 API / C: 混合） | **方式 A**（与 Claude Code 工作流契合，无 key 管理） | ☐ |
| 4 | 是否保留 `.changelog-cache/` | 方式 A 下**不保留**；方式 B 下**保留**以省 API 调用 | ☐ |
| 5 | 是否做粗粒度聚合（合并相似 commit） | 在 release-prep 的 review 阶段顺手做（v3 优势：commit 前可改） | ☐ |
| 6 | 实施节奏（一次到位 / 先模板预览 / 分步） | **先做步骤 1–2**（cliff.toml + extract 脚本），本地试跑英文版；OK 再做翻译那部分 | ☐ |

## 8. 参考

- [Keep a Changelog v1.1.0](https://keepachangelog.com/en/1.1.0/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [git-cliff configuration](https://git-cliff.org/docs/configuration)
- 上游 SDK：[agentbay-ai/wuying-agentbay-sdk CHANGELOG](https://github.com/agentbay-ai/wuying-agentbay-sdk/blob/main/CHANGELOG.md)
- Ant Design 双语：[CHANGELOG.en-US.md](https://github.com/ant-design/ant-design/blob/master/CHANGELOG.en-US.md) / [CHANGELOG.zh-CN.md](https://github.com/ant-design/ant-design/blob/master/CHANGELOG.zh-CN.md)
- 项目内：[CHANGELOG.md](../../CHANGELOG.md) | [cliff.toml](../../cliff.toml) | [scripts/backfill-release-notes.sh](../../scripts/backfill-release-notes.sh) | [docs/release-checklist.md](release-checklist.md)
