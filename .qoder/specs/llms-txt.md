# AgentBay CLI 引入 llms.txt 方案

## Context

LLMs.txt 是 Jeremy Howard(Answer.AI)2024 年 9 月提出的提案,目标是让网站/项目对 LLM 友好,类似 `robots.txt` 之于爬虫、`sitemap.xml` 之于搜索引擎。规范要点:

- **位置**:项目/站点根目录 `/llms.txt`,可选搭配 `/llms-full.txt`
- **格式**:Markdown,人和机器都能读
- **`llms.txt`**:精炼的导航索引,帮 LLM 在有限上下文窗口内快速理解项目全貌
- **`llms-full.txt`**:全部核心文档拼接成单文件,LLM 可一次性吃完整个知识库

业界采用情况:Anthropic、Stripe、Cloudflare、FastAPI、Hugging Face、Vercel 等已发布。

### 为什么要在本项目加

1. **文档结构已经匹配 llms.txt 范式**——本项目已有完整的双语用户文档(`docs/en/` + `docs/zh/`,各 12 篇),按主题切分清晰,只需做索引即可,几乎零成本
2. **CLI 类项目特别受益**——用户经常通过 LLM 询问 "agentbay 怎么创建 image"、"怎么用 apikey 登录" 等问题,llms.txt 让 Cursor / Claude / ChatGPT 在抓取仓库时能快速定位正确文档,而不是盲读 README
3. **趋势所向**——Anthropic、Stripe、FastAPI 等同类工具/SDK 项目均已采用,正在形成事实标准

### 项目现状(2026/06/02)

- 仓库地址:`github.com/aliyun/agentbay-cli`(主分支 `master`)
- 双语 README:`README.md`(英文)、`README.zh-CN.md`(中文)
- 用户文档:`docs/en/` 和 `docs/zh/` 各 12 篇,1:1 翻译对应
- 内部文档:`docs/internal/`(release-checklist、cli-openapi-actions 等,**不对客**)
- 英文文档总量:2506 行 / 10,757 words / **~95 KB / ~16K tokens**
- 中文文档总量:2494 行,体量与英文相当

---

## 决策与权衡

### 决策 1:`llms.txt`(索引)采用**双语**

`llms.txt` 体量极小(只是链接列表),把中英文档都列出来不会带来负担,反而能让 LLM 感知到中文版本的存在,在用户特别需要原汁原味中文文档时引导他们直接打开 `docs/zh/xxx.md`。

`docs/internal/` 目录**不收录**——内部文档不应暴露给 LLM。

### 决策 2:`llms-full.txt`(全文拼接)**仅英文**

| 方案     | 体量                  | 评估                          |
| -------- | --------------------- | ----------------------------- |
| 仅英文   | ~95 KB / ~16K tokens  | ✅ 推荐                       |
| 仅中文   | ~120 KB / ~24K tokens | ❌ 海外用户体验下降           |
| 中英都收 | ~215 KB / ~40K tokens | ❌ 同份内容翻倍,对 LLM 是噪声 |

理由:

1. **`llms-full.txt` 设计初衷是「单一信息源」**——同份内容的两个语言版本对 LLM 是噪声,不是信号(它本就具备双向翻译能力)
2. **跨语言推理已经够用**——主流模型用英文文档回答中文问题质量足够,业界(Anthropic / Stripe / FastAPI / Cloudflare)的 llms-full.txt 都只发英文
3. **术语一致性更好**——CLI 命令、参数名本身是英文(`agentbay image create-from-template --source-image`),英文文档与实际命令对齐;基于中文文档反而可能产生「源镜像 / 源 image」这种混译
4. **体量友好**——~16K tokens 在所有现代 LLM 窗口里都毫无压力(Claude 200K 占 ~9%,GPT-4 128K 占 ~14%)

### 决策 3:可发现性

- 在 `README.md` 和 `README.zh-CN.md` 顶部各加一行 LLM 友好提示,链接到 `llms.txt`
- 不影响 README 主体阅读体验,但能让仓库访问者知道这个文件的存在

---

## Phase 1: 生成 `llms.txt`(双语索引)

**目标**:项目根目录 `llms.txt`

**结构**:

```
# AgentBay CLI
> 项目一句话介绍

可选的简短背景说明(2-3 句)

## Getting Started
- README、Installation、Authentication

## Tutorials
- Image Creation & Sharing(端到端教程)

## Command Reference
- Core / Image / API Key / Network / Skills / Docker

## Permissions
- RAM Permissions

## 中文文档
- 12 篇中文文档完整链接

## Optional
- FAQ、CHANGELOG、LICENSE
```

**链接规则**:

- 全部使用绝对 URL:`https://github.com/aliyun/agentbay-cli/blob/master/...`
- 不收录 `docs/internal/`
- 不收录 `test/README.md`、`scripts/README.md`、`.aoneci/*` 等开发者内部内容

**状态**:✅ 已完成(本次对话已生成 [llms.txt](../../llms.txt))

---

## Phase 2: 生成 `llms-full.txt`(英文全文)

**目标**:项目根目录 `llms-full.txt`

**拼接顺序**(按用户学习路径组织):

1. **顶部元信息**(从 `llms.txt` 复用项目简介)
2. **快速入门**:`README.md` → `docs/en/installation.md` → `docs/en/authentication.md`
3. **教程**:`docs/en/image-workflow.md`
4. **命令参考**(按使用频率排序):
   - `docs/en/core.md`
   - `docs/en/image.md`
   - `docs/en/apikey.md`
   - `docs/en/docker.md`
   - `docs/en/network.md`
   - `docs/en/skills.md`
5. **权限**:`docs/en/ram-permissions.md`
6. **FAQ**:`docs/en/faq.md`

**拼接格式约定**:

每篇文档前加分隔符和原始路径标注,方便 LLM 引用来源:

```markdown
---

# === Source: docs/en/image.md ===

[原文档内容,跳过开头的语言切换链接]
```

**跳过项**:

- 文档开头的 `[中文](../zh/xxx.md) | **English**` 切换链接(对 LLM 无意义)
- `docs/en/README.md`(只是索引页,内容已在 `llms.txt` 体现)
- `docs/internal/*`

**生成方式**:写一个 shell 脚本 `scripts/build-llms-full.sh`,后续文档更新时跑一下重新生成。

---

## Phase 3: 生成脚本 `scripts/build-llms-full.sh`

**目标路径**:`scripts/build-llms-full.sh`

**职责**:

1. 按 Phase 2 定义的顺序拼接英文文档
2. 自动跳过每篇开头的语言切换链接行
3. 在每篇前插入 `# === Source: <相对路径> ===` 标注
4. 输出到项目根目录 `llms-full.txt`
5. 输出生成统计(总行数 / 字节数 / 估算 token 数)

**调用方式**:

```bash
./scripts/build-llms-full.sh
```

**集成考虑**(可选):

- 在 `Makefile` 中加 `make llms-full` 目标
- 不集成到 release 流程——`llms-full.txt` 与文档变更同步即可,不需要每次发版重生成

---

## Phase 4: README 加入发现入口

在两个 README 顶部 `[中文版/English]` 切换行下方,加一行轻提示:

**`README.md`**:

```markdown
[中文版](README.zh-CN.md) | **English**

> 🤖 **LLM-friendly**: This project provides [llms.txt](llms.txt) and [llms-full.txt](llms-full.txt) for AI assistants.
```

**`README.zh-CN.md`**:

```markdown
**中文** | [English](README.md)

> 🤖 **LLM 友好**:本项目提供 [llms.txt](llms.txt) 与 [llms-full.txt](llms-full.txt),便于 AI 助手理解。
```

**注意**:这是项目里**唯一允许使用 emoji 的位置**——因为 🤖 已经是 LLM/AI 相关项目的视觉惯例(Anthropic、Cursor 等都这么做),且只用一处不会蔓延。

---

## Phase 5: 维护机制

**触发条件**：`docs/en/` 或 `README.md` 发生变更时，需重新生成 `llms-full.txt`。

**当前落地方式**：手动 + checklist + Qoder SOP 固化。

1. `docs/internal/release-checklist.md` 已加入发版前检查：如本次涉及 `docs/en/` 或 `README.md` 变更，执行 `bash scripts/build-llms-full.sh` 并提交 `llms-full.txt`。
2. `.qoder/rules/development.md` 已加入 LLM-facing 文档维护 SOP，作为所有开发入口的全局规则。
3. `.qoder/skills/update-cli-command-docs/SKILL.md` 已加入 llms readiness，CLI 命令文档同步时会检查并更新 `llms-full.txt` / `llms.txt`。
4. `.qoder/skills/create-cli-command/SKILL.md` 已将 llms readiness 作为 Phase 5 文档交接的一部分。
5. `.qoder/skills/bilingual-changelog-release/SKILL.md` 已将 llms 检查纳入发版 pre-flight。

**维护规则**：

- 修改 `README.md` 或 `docs/en/**`：执行 `bash scripts/build-llms-full.sh`。
- 新增、删除、重命名对外文档：同步更新 `llms.txt`；如涉及英文源文档，同步重建 `llms-full.txt`。
- `docs/internal/**`、测试文档、脚本文档不进入 `llms.txt` / `llms-full.txt`。

**后续可选增强**：如果仍频繁遗漏，可在 `.github/workflows/` 增加 CI 校验：PR 中如果 `README.md` 或 `docs/en/**` 改了但 `llms-full.txt` 没改，则报警告或失败。

---

## 不在本方案范围内

以下内容明确**不做**,避免 scope creep:

- ❌ 中文版 `llms-full.txt` —— 决策 2 已说明理由
- ❌ 把 `docs/internal/` 暴露到 llms.txt —— 内部文档不对客
- ❌ 多版本 / 历史版本的 llms.txt —— 当前项目仅维护 master 一份文档
- ❌ 自动从源代码生成命令参考再合并到 llms-full.txt —— 文档已是手写精炼版,自动生成反而更糟
- ❌ 集成到 release 工作流 —— 文档驱动,不是 release 驱动

---

## 验收标准

- [x] 项目根目录存在 `llms.txt`,符合 llms.txt spec(单 H1 + blockquote + H2 章节 + 链接列表)
- [x] 项目根目录存在 `llms-full.txt`,英文文档全文,~95 KB
- [x] 存在 `scripts/build-llms-full.sh`,可重复生成 `llms-full.txt`
- [x] 两个 README 顶部均有 LLM 友好提示
- [x] `docs/internal/release-checklist.md` 包含更新 `llms-full.txt` 的条目
- [x] `llms.txt` 中所有链接可点击访问(GitHub master 分支)
- [x] `llms-full.txt` 不包含 `docs/internal/` 内容
- [x] `.qoder/rules/development.md` / 相关 skills 已固化 llms 后续维护 SOP
