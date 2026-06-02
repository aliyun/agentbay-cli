---
trigger: always_on
---

# AgentBay CLI 开发规则

## 🔐 文档敏感信息脱敏规则

### 重要：文档中的真实账号 ID 必须脱敏

**规则**：在文档（包括 `docs/`、`README`、`CHANGELOG` 等所有对外可见的 Markdown 文件）中编写示例、输出样例或场景描述时，**禁止**直接展示真实的账号标识符（AliUID、用户 ID 等）。

**脱敏格式**：保留末尾 4 位字符，其余替换为 `****`。

| 原始值             | 脱敏后显示 |
| ------------------ | ---------- |
| `1730408327554214` | `****4214` |
| `1242716971377069` | `****7069` |
| `abc123456789`     | `****6789` |

**适用范围**：

- Aliyun UID（阿里云主账号 / RAM 子账号 UID）
- 任何以用户真实账号 ID 作为路径组成部分的 URL 或镜像地址（如 `/customer_cli/<uid>:tag`）
- 命令行示例中出现的 `--target-uid <uid>` 等参数值
- 输出示例中的 `PeerAliUid` 列内容

**豁免**：内部测试脚本（`test/`、`.aoneci/`）中的测试数据可使用占位符（如 `<YOUR_UID>`）或脱敏值，无需展示真实值。

**检查清单**：

- [ ] 新增或更新文档时，搜索是否包含完整的数字 UID（通常 10+ 位纯数字）
- [ ] 若有，按 `****<末尾4位>` 格式统一替换
- [ ] 同时检查 URL / 镜像路径中嵌入的 UID 段

---

## 🤖 LLM-facing 文档维护 SOP

### 重要：CLI / 文档变更必须同步 llms 文档

**规则**：`llms.txt` 和 `llms-full.txt` 是面向 AI 助手的对外文档入口。后续任何 CLI 功能、命令文档或 README 调整，只要影响用户可见说明，就必须检查并同步 llms 相关文件。

**触发条件**：

| 变更类型                                                                | 必须动作                                                                         |
| ----------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| 修改 `README.md`                                                        | 执行 `bash scripts/build-llms-full.sh`，提交更新后的 `llms-full.txt`             |
| 修改 `docs/en/**`                                                       | 执行 `bash scripts/build-llms-full.sh`，提交更新后的 `llms-full.txt`             |
| 新增 / 删除 / 重命名对外文档                                            | 更新 `llms.txt` 中的中英文导航链接；如涉及英文源文档，同步重建 `llms-full.txt`   |
| 新增 / 修改 / 删除 CLI 命令、参数、输出或用户可见行为                   | 若同步修改了 `README.md` 或 `docs/en/<group>.md`，必须同步重建 `llms-full.txt`   |
| 仅修改 `docs/zh/**`                                                     | 通常不需要重建 `llms-full.txt`；若文档结构发生变化，仍需检查 `llms.txt` 中文链接 |
| 修改 `docs/internal/**`、`test/**`、`scripts/README.md` 等内部/开发文档 | 不进入 `llms.txt` / `llms-full.txt`，通常无需同步                                |

**执行命令**：

```bash
bash scripts/build-llms-full.sh
```

**维护原则**：

- `llms.txt` 是静态导航索引，采用双语链接；仅在文档文件新增、删除、重命名、命令组索引变化时手动更新。
- `llms-full.txt` 是英文全文聚合文件，只由 `README.md` 和 `docs/en/**` 生成，不纳入 `docs/zh/**` 或 `docs/internal/**`。
- `scripts/build-llms-full.sh` 中的 `FILES` 顺序代表 LLM 学习路径；新增英文对外文档时必须评估是否加入该数组。
- 发版前若本次 release 包含 `README.md` 或 `docs/en/**` 变更，必须确认 `llms-full.txt` 已重新生成并提交。

**检查清单**：

- [ ] 本次是否修改了 `README.md` 或 `docs/en/**`
- [ ] 如是，已执行 `bash scripts/build-llms-full.sh`
- [ ] 新增 / 删除 / 重命名对外文档时，`llms.txt` 链接已同步
- [ ] `llms-full.txt` 未包含 `docs/internal/**` 内容
- [ ] `git diff --stat` 中包含应同步的 `llms-full.txt` / `llms.txt` 变更，或已明确说明无需更新

---

## 🔗 Skill 自动装配规则（Quest / 对话 / 任意入口通用）

**凡符合下列任一特征的任务，AI 必须主动加载并遵循对应的 `.qoder/skills/` 指南**（包括但不限于 Quest Design/Execute 阶段、直接对话、Execute Directly 模式）：

| 任务特征                                                          | 必须加载的 Skill                                                                                                                       | Skill 路径                                            |
| ----------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------- |
| 新增 / 修改 CLI 命令、参数、子命令，或将前端 API 封装为命令       | **create-cli-command**                                                                                                                 | `.qoder/skills/create-cli-command/SKILL.md`           |
| 涉及分支管理、commit、push、PR、变更档案（Quest/CR 目录）         | **feature-development-workflow**                                                                                                       | `.qoder/skills/feature-development-workflow/SKILL.md` |
| 更新/同步 CLI 命令文档（README、docs/、CHANGELOG readiness）      | **update-cli-command-docs**                                                                                                            | `.qoder/skills/update-cli-command-docs/SKILL.md`      |
| 发版、双语 CHANGELOG、GitHub Release notes、release-prep/backfill | **bilingual-changelog-release**                                                                                                        | `.qoder/skills/bilingual-changelog-release/SKILL.md`  |
| 新增 CLI 命令类需求（同时触发上述三条）                           | **三者组合使用**：先 workflow 拉分支/建档 → 再 create-cli-command 实现 → 再 update-cli-command-docs 同步文档 → 回到 workflow 提交/推送 | 同上                                                  |

**执行铁律**:

1. **不得跳过**：即便用户未显式打 `/skill-name`，只要任务特征命中上表，AI 就必须主动阅读 skill 文件并按其 Phase 执行。
2. **前置检查**：进入实现阶段前，必须确认 `feature-development-workflow` 的 Phase 0（变更档案）和分支创建已完成，否则先引导用户补齐。
3. **Quest 场景**：Quest 生成 spec 后的 Execute 阶段，等同于"对话入口"，本规则照常生效，无需 spec 里额外声明。
4. **Execute Directly 场景**：即便跳过 Design 阶段，AI 也必须在动手前主动加载匹配的 skill。
5. **文档同步**：`create-cli-command` 的 Phase 5 已委托 `update-cli-command-docs` skill，文档操作不得在 `create-cli-command` 中内联执行。
6. **发版同步**：凡涉及 `CHANGELOG.md` 版本段、GitHub Release notes、tag 发布或 release notes 回灌，必须切换到 `bilingual-changelog-release` skill；不得继续使用旧的 `git-cliff -o CHANGELOG.md` 全量覆盖流程。

**目的**：让 Skill 指南在所有入口（slash command / Quest spec / 直接对话 / Execute Directly）下统一自动生效，避免重复约定。

---

## 🚫 Git 提交规则

### 重要：不要自动提交代码

**规则**: 除非用户明确要求（例如"帮我提交代码"、"commit 代码"、"push 到远程"等明确指令），否则**绝对不要**执行以下操作：

- ❌ 不要执行 `git commit`
- ❌ 不要执行 `git push`
- ❌ 不要执行 `git add`（除非是测试命令的一部分）

**正确做法**:

- ✅ 完成代码修改后，询问用户："需要我帮你提交代码吗？"
- ✅ 展示 `git status` 和 `git diff` 让用户确认
- ✅ 等待用户明确指示后再执行提交操作

**原因**:

- 用户可能需要先审查代码改动
- 用户可能需要手动调整 commit message
- 用户可能需要先测试代码
- 避免误提交未完成的代码

---

## 📝 代码提交规范

当用户明确要求提交代码时，遵循以下规范：

### Commit Message 格式

使用 Conventional Commits 规范：

```
<type>[(<scope>)][!]: <description>

[optional body]

[optional footer]
```

**Type 类型**:

- `feat`: 新功能（进入 CHANGELOG）
- `fix`: 修复 bug（进入 CHANGELOG）
- `docs`: 文档更新（进入 CHANGELOG）
- `refactor`: 代码重构（进入 CHANGELOG）
- `perf`: 性能优化（进入 CHANGELOG）
- `revert`: 回退变更（进入 CHANGELOG 兜底分类）
- `test`: 添加或修改测试（默认不进 CHANGELOG）
- `style`: 代码格式调整（不影响功能，默认不进 CHANGELOG）
- `chore`: 构建过程或辅助工具的变动（默认不进 CHANGELOG）
- `ci` / `build`: CI 或构建系统变更（默认不进 CHANGELOG）

**Scope 建议**：优先使用 CLI 命令组或模块名，例如 `apikey`、`image`、`network`、`skills`、`docker`、`core`、`client`。

**Breaking Changes**：不兼容变更必须在 subject 加 `!` 或在 footer 写 `BREAKING CHANGE:`，以便 git-cliff 归入 `⚠️ Breaking Changes`。

**示例**:

```bash
feat(apikey): add concurrency management command

- Add 'agentbay apikey concurrency set' command
- Use named parameters for better UX
- Add parameter validation
```

### 提交流程

1. **展示变更**

   ```bash
   git status
   git diff --stat
   ```

2. **执行提交**

   ```bash
   git add -A
   git commit -m "清晰的 commit message"
   ```

3. **确认结果**
   ```bash
   git log --oneline -3
   ```

---

## 🗒️ 双语 CHANGELOG / GitHub Release SOP

### 核心原则

- `CHANGELOG.md` 是 GitHub Release body 的**唯一事实源**。
- 发版前在本地执行 `make release-prep VERSION=X.Y.Z` 生成双语版本段；workflow 只抽取该版本段，不再生成、不再翻译、不再 commit-back `CHANGELOG.md`。
- 日常 CLI 功能开发**不再手动运行 `git-cliff -o CHANGELOG.md` 全量覆盖**；CHANGELOG 内容由符合 Conventional Commits 的提交在 release-prep 时统一生成。
- 中文翻译必须在 commit/tag 之前完成并人工 review，禁止发布带 `TRANSLATE_ME` 或旧 `中文翻译待补充` 占位符的版本段。

### 发版前标准流程

1. 在功能分支（默认如 `feat/dev-apikey`）准备发版内容，执行：

   ```bash
   make release-prep VERSION=X.Y.Z
   ```

2. 翻译 `CHANGELOG.md` 顶部 `## [X.Y.Z]` 版本段中的 `### English` 内容到 `### 中文`，删除 `TRANSLATE_ME` 占位；中文段分类标题必须使用中文，不得保留 `Bug Fixes` / `Documentation` 等英文标题。
3. 验证：

   ```bash
   grep -nE 'TRANSLATE_ME|中文翻译待补充' CHANGELOG.md
   bash scripts/extract-changelog-section.sh X.Y.Z CHANGELOG.md >/tmp/release-notes.md
   ```

   第一条应无输出，第二条应成功且输出非空。

4. 提交并推送功能分支，经 PR 合入 `aliyun/master`。本项目默认发布路径是：

   ```text
   本地 feat/dev-apikey
     ↓
   push aliyun/feat-dev-apikey
     ↓
   PR 合入 aliyun/master
     ↓
   GitHub Actions 手动 Run workflow，输入 X.Y.Z（如 0.4.0）
     ↓
   workflow 从 master 的 CHANGELOG.md 抽取 X.Y.Z 段
     ↓
   gh release create vX.Y.Z --target "$GITHUB_SHA"
     ↓
   如果 vX.Y.Z tag 不存在，则自动创建 tag
     ↓
   创建 GitHub Release
   ```

5. **默认不在本地手动打 tag**：tag 由 `.github/workflows/homebrew.yml` 中的 `gh release create "v$VERSION" --target "$GITHUB_SHA"` 在 Release 创建时自动创建，并绑定到本次 workflow checkout/build 的 `master` commit。仅当用户明确选择 tag-driven release 时，才走本地/CI 预先推送 `vX.Y.Z` tag 的可选路径。
6. 如需修订已发布 Release 或统一整理历史 CHANGELOG：先改 `CHANGELOG.md`，历史版本也必须保留 `### English` 与 `### 中文` 双语结构；注意这不会自动更新 GitHub 上已存在的 Release body。若同时准备新版本和历史回灌，先完成新版本 Release，再用 `scripts/backfill-release-notes.sh --dry-run` 预览历史回灌内容，确认后再执行正式回灌。

### 双语段组织规则

- `### English` 与 `### 中文` 必须保持**结构对齐**：分类标题、命令组、子条目数量和顺序必须一一对应；中文段是英文段的等价翻译与用户视角表达，不得新增、遗漏或合并英文段中的独立条目。
- 中文分类标题必须翻译：`Features` → `功能`，`Bug Fixes` → `缺陷修复`，`Documentation` → `文档`，`Refactoring` → `重构`，`Performance` → `性能优化`，`Security` → `安全`，`Breaking Changes` → `不兼容变更`，`Other Changes` → `其他变更`。
- 对 CLI 命令相关条目，英文和中文都必须优先按命令组归类（如 `apikey`、`image`、`docker`、`skills`、`network`、`core/auth`），父条目使用加粗命令组名；格式视改动数量而定（见下方格式规范）。
- 无法归入具体命令组或属于全局能力 / 基础设施 / 发版流程的改动，英文可归为 `global`、`security/compliance`、`RAM permissions`、`release` 等；中文对应归为“全局”“安全合规”“RAM 权限”“发版”等用户可理解的主题。
- 允许在双语段做粗粒度聚合，但必须**双语同步聚合**：如果英文将多个 commit 聚合为一个用户可感知条目，中文必须使用同一个条目结构；如果中文拆成多个子条目，英文必须同步拆成对应子条目。

#### 📐 命令组条目格式规范（英文 & 中文段统一适用，多改动必须拆行）

**核心原则：`### English` 和 `### 中文` 段均须遵守。一行只描述一个子命令或功能点，禁止在同一行用顿号、分号并列多个改动。**

| 改动数量                          | 格式                                                                 |
| --------------------------------- | -------------------------------------------------------------------- |
| 命令组下有 **2 个及以上**独立改动 | 父条目只写加粗命令组名（不加冒号和说明），子条目两空格缩进，每行一项 |
| 命令组下仅有 **1 个**改动         | 单行：父条目直接加冒号说明                                           |

多改动示例（正确，英文/中文必须同构）：

```markdown
### English

- **skills**
  - `skills show`: Display FileUrl and TenantTags fields
  - `skills push`: Support `--tag` and `--icon` flags
  - `skills list`: Support `--output json` and improve terminal-adaptive display

### 中文

- **skills**
  - `skills show`：展示 FileUrl 与 TenantTags 字段
  - `skills push`：支持 `--tag` 与 `--icon` 参数
  - `skills list`：支持 `--output json`，优化终端自适应展示
```

单改动示例（正确，英文/中文必须同构）：

```markdown
### English

- **apikey**: `apikey list` supports `--output json` for structured script and AI consumption

### 中文

- **apikey**：`apikey list` 支持 `--output json`，便于脚本和 AI 场景读取结构化结果
```

子条目格式：

- 英文：`  - \`<subcommand>\`: <concise description>`
- 中文：`  - \`<子命令>\`：<简洁描述>`

两者均使用两个空格缩进 + 破折号。

#### ✅ 双语结构对齐检查清单

- [ ] 每个版本段同时包含 `### English` 与 `### 中文`
- [ ] 英文分类与中文分类数量一致、顺序一致（如 Features ↔ 功能）
- [ ] 每个分类下的命令组数量一致、顺序一致（如 `skills` ↔ `skills`、`global` ↔ `全局`）
- [ ] 每个命令组下的子条目数量一致、顺序一致
- [ ] 英文聚合了多个 commit 时，中文使用相同聚合粒度；中文拆行时，英文同步拆行
- [ ] commit / PR 链接可以只保留在 English 段；但链接对应的用户可感知改动必须在中文段有等价描述

### 翻译术语

- 保留英文不翻译：API Key、AK/SK、CLI、OSS、SDK、PR、Homebrew、OAuth、apikey、image、docker
- `image` → 镜像；`container` → 容器；`warmup` → 预热；`session` → 会话；`context` → 上下文
- `scope` 在权限语境译为“范围”，变量语境译为“作用域”
- `flag` 译为“参数”或“选项”，不要译为“标志”
- 命令名、参数名、scope、PR 链接、author 保持原样

### 对应 Skill

发版、翻译 CHANGELOG、tag 发布、GitHub Release notes 回灌必须遵循 `.qoder/skills/bilingual-changelog-release/SKILL.md`。

---

## 📋 `--output json` 输出格式 SOP

### 适用场景

凡命令返回**列表类数据**（即输出一个以上条目的表格），**必须**支持 `--output json` flag。

**判断标准**：

| 命令类型                    | 是否需要      |
| --------------------------- | ------------- |
| `list` 命令（返回列表）     | ✅ 必须添加   |
| `show` / 查询单条详情       | ❌ 通常不需要 |
| 只读查询但返回单一字段      | ❌ 不需要     |
| 创建 / 修改 / 删除 等变更类 | ❌ 不需要     |

### Flag 设计规范

- Flag 名称统一为 `--output`
- **短参数**：
  - 区分是否已有其他短参数占用 `-o`：
    - 未被占用：添加 `-o` 短参数（如 `apikey list`、`skills list`）
    - 已被占用（如 `image list` 的 `--os-type -o`）：**不加短参数**
- 当前仅支持 `json` 一种输出格式

```go
// 有短参数的情况
cmd.Flags().StringP("output", "o", "", `Output format. Use "json" for machine-readable output (e.g. for AI/scripts)`)

// 无短参数的情况（-o 已被占用）
cmd.Flags().String("output", "", `Output format. Use "json" for machine-readable output (e.g. for AI/scripts)`)
```

### 实现模板

```go
func runXxxList(cmd *cobra.Command, args []string) error {
    outputFmt, _ := cmd.Flags().GetString("output")

    // ... 调用 API、获取数据 ...

    // JSON 输出分支：放在表格输出之前
    if strings.EqualFold(outputFmt, "json") {
        type itemJSON struct {
            // 输出全量字段，包括表格中被桓出的列
            Field1 string `json:"field1"`
            Field2 string `json:"field2"`
        }
        type outputJSON struct {
            TotalCount int        `json:"totalCount"`
            NextToken  string     `json:"nextToken,omitempty"`  // 分页命令需要
            Items      []itemJSON `json:"items"`
        }
        out := outputJSON{TotalCount: len(items)}
        for _, item := range items {
            // 填充字段
        }
        if out.Items == nil {
            out.Items = []itemJSON{} // 空数组用 [] 而非 null
        }
        b, err := json.MarshalIndent(out, "", "  ")
        if err != nil {
            return fmt.Errorf("json marshal: %w", err)
        }
        fmt.Println(string(b))
        return nil
    }

    // 默认表格输出
    printTable(items)
    return nil
}
```

### JSON 输出字段要求

1. **包含全量字段**：导出所有 API 返回的字段，包括表格中因横向空间限制而被栓略的列
2. **字段命名**：使用 camelCase（如 `skillId`、`gmtCreate`、`statusDisplay`）
3. **空数组**：永远输出 `[]` 而非 `null`，即 `if out.Items == nil { out.Items = []itemJSON{} }`
4. **可选字段**：使用 `omitempty`（如分页的 `nextToken`）
5. **不包含内部请求元信息**：`[INFO] Request ID:` 行仍打印到 stdout，但不包含在 JSON 输出中

### 文档要求

凡新增 `--output json` 支持，必须在对应的 `docs/zh/<group>.md` 和 `docs/en/<group>.md` 中记录：

1. 在 Flags 表格中添加 `--output` 行
2. 提供 JSON 输出示例
3. 如有短参数冲突（如 `image list` 的 `-o` 已被 `--os-type` 占用），要在文档中说明

### 参考实现

- [`cmd/skills.go`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/skills.go) — `runSkillsList`（`-o json` 短参数）
- [`cmd/image_list_helper.go`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/image_list_helper.go) — `printImagesAsJSON` 共享输出 helper
- [`cmd/apikey_list.go`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/apikey_list.go) — `runApikeyList`（`-o json` 短参数）

### 检查清单

- [ ] `--output` flag 已添加，有无短参数 `-o` 遵循冲突检查规则
- [ ] JSON 输出分支放在表格输出**之前**，而非之后
- [ ] 空数组输出 `[]` 而非 `null`
- [ ] 字段包含表格中因横向空间被栓略的列
- [ ] 对应的 `docs/zh/<group>.md` 和 `docs/en/<group>.md` 已更新

---

## 💻 Go 代码规范

### ⚠️ OpenAPI 接口请求头规范（重要！）

**规则**：所有 OpenAPI 接口请求**必须**在 `Headers` 中显式设置 `Accept: application/json`，禁止使用 `application/xml` 或省略该头。

**标准写法**：

```go
req := &openapiutil.OpenApiRequest{
    Query: openapiutil.Query(query),  // 或 Body 字段
    Headers: map[string]*string{
        "Accept": dara.String("application/json"),
    },
}
```

**背景**：`DescribeMarketSkillDetail` 接口历史上曾使用 `application/xml`（因服务端响应格式不稳定），现已统一改为 `application/json`。`dual_format_responses.go` 中的双格式 parser 仍保留 XML 分支作为容错兜底，但请求头统一为 JSON。

**检查清单**：

- [ ] 新增接口时，`Headers` 中设置 `"Accept": dara.String("application/json")`
- [ ] 禁止设置 `Accept: application/xml`
- [ ] 新增接口后，用 `grep -n 'application/xml' internal/client/` 确认无遗漏

---

### ⚠️ 接口变更必须同步更新 Mock（重要！）

**规则**: 当给接口（interface）添加新方法时，**必须立即更新所有实现该接口的 mock 类**！

**问题案例**:

```go
// 在 internal/agentbay/client.go 中添加新方法
type Client interface {
    CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error)
    ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error)
}
```

**❌ 错误做法**: 只更新接口，不更新 mock 类

- 导致 CI 编译失败
- 错误信息：`*mockClient does not implement agentbay.Client (missing method CreateApiKey)`

**✅ 正确做法**:

1. 找到所有实现该接口的 mock 类
2. 为每个 mock 类添加新方法（返回 "not implemented" 错误）

```go
// cmd/image_status_helper_test.go
func (m *mockGetMcpImageInfoClient) CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

// cmd/image_list_helper_test.go
func (m *mockImageListClient) CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

**查找所有 mock 类的方法**:

```bash
# 搜索所有 mock 类定义
grep -r "type mock.*Client struct" cmd/ test/

# 搜索所有实现接口的声明
grep -r "var _ agentbay.Client" cmd/ test/
```

**检查清单**:

- [ ] 接口添加新方法后，立即搜索所有 mock 类
- [ ] 为每个 mock 类添加对应的方法实现
- [ ] 运行 `go test ./...` 确保编译通过
- [ ] 本地验证后再提交代码

---

### ⚠️ 响应解析必须使用容错模板（重要！）

**规则**: 新增 / 修改任意 OpenAPI 接口的响应 parser 时，**必须**按 `internal/client/dual_format_responses.go` 的 dual-format 模板实现，杜绝直接用 `json.Unmarshal` 打到带 `*int32`/`*bool` 的强类型结构体。

**背景案例**（已发生，禁止复现）:

`BatchCreateHideResourceGroupsWithMaxSession` 早期 parser 直接把响应体 `json.Unmarshal` 到 `BatchCreateHideResourceGroupsWithMaxSessionResponseBody{HttpStatusCode *int32}`，但服务端实际返回 `"HttpStatusCode": "200"`（字符串），导致：

```text
Error: failed to set max session: json: cannot unmarshal string into Go struct field
  BatchCreateHideResourceGroupsWithMaxSessionResponseBody.HttpStatusCode of type int32
```

整个命令直接失败，用户无法使用。同类问题也在 `GetDockerfileTemplate.NonEditLineNum` 上出现过。

**✅ 强制要求**：

1. **parser 位置**：`parseXxxResponse` 统一放在 `internal/client/dual_format_responses.go`，**不**再写在 `client.go` 里（`client.go` 只保留「This file is auto-generated」的 action 调用骨架）。
2. **数字字段**：所有 `*int32` / `*int64` 字段在 JSON 路径一律用 `json.RawMessage` 中转 + 复用 `int32FromFlexibleJSON` 辅助解析，兼容数字与字符串两种序列化形式。
3. **XML/JSON 双格式**：body 以 `<` 开头走 XML 分支、否则走 JSON 分支，两条路径都要通过 `applyMapHeadersAndStatus` 归一 headers / statusCode。
4. **错误包装**：任何解析失败都必须用 `&ErrWithRequestID{Err: ..., RequestID: extractRequestIDFromResponse(res)}` 包装，保证 RequestId 能透出到 CLI。
5. **最小单测**：每个 parser **必须**在 `internal/client/` 下配套一个 `xxx_parse_test.go`，至少覆盖：
   - JSON 数字字段返回为字符串（`"HttpStatusCode":"200"`）
   - JSON 数字字段返回为数字（`"HttpStatusCode":200`）
   - XML 分支
6. **回归验证**：`go test ./internal/client/... -count=1` 与 `go test ./... -count=1` 均通过，再进入命令层。

**参考实现**:

- [parseBatchCreateHideResourceGroupsWithMaxSessionResponse](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/dual_format_responses.go)
- [parseGetDockerfileTemplateResponse](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/dual_format_responses.go)
- [batch_create_hide_resource_groups_with_max_session_parse_test.go](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/batch_create_hide_resource_groups_with_max_session_parse_test.go)

**检查清单**:

- [ ] parser 写在 `dual_format_responses.go`，没有放回 `client.go`
- [ ] 数字字段走 `int32FromFlexibleJSON`，不直接 Unmarshal 到 `*int32`
- [ ] JSON / XML 两条分支齐备
- [ ] `xxx_parse_test.go` 覆盖「数字字段为字符串 / 数字 / XML」三种场景
- [ ] 解析失败全部走 `ErrWithRequestID` 包装

---

### ⚠️ 命令层接口成功判定 SOP（仅约束新增接口，存量不改造）

**背景**：OpenAPI 各接口对响应体的字段下发并不一致——有的接口在 `data` 里返回 `Success: true`，有的接口（例如 `DeleteApiKey`）只返回 `{"RequestId":"...","HttpStatusCode":200,"Code":"ok"}`，**根本不下发 `Success`**。如果命令层用 `if !resp.Body.GetSuccess()` 判失败，对后者会因为 `Success` 是 `*bool` 且为 `nil`，`GetSuccess()` 返回 `false`，直接给用户报出 `Code=ok, Message=` 的"假错误"。

> 真实事故：`apikey delete` 服务端实际删除成功，但 CLI 因为依赖 `Success` 字段导致命令以错误码退出。

**适用范围**：

- ✅ **对接全新 OpenAPI 接口的命令**：必须遵守本 SOP。
- ❌ **存量已上线且工作正常的命令**（如 `apikey enable/disable`、`describe-mcp-api-key` 等）：**不改造**，避免引入回归风险。
- ❌ **新增 CLI 命令但复用已有接口**：如果该接口在已有命令中已用 `GetSuccess()` 判定且工作正常，新命令中**沿用相同写法即可**，不要为了统一而改。
- 📌 **`create-cli-command` skill 模板保持现状**（仍使用 `GetSuccess()` 写法），新增接口若发现服务端不下发 `Success`，按本 SOP 切换写法即可。

**简言之**：只有当你为一个**从未在 CLI 中使用过的全新 OpenAPI 接口**编写命令时，才需要按本 SOP 决定成功判定写法。

**判定规则**（新增接口必须遵守）：

1. **以 `Code` 字段为主依据**：约定 `"ok"`（不区分大小写）= 成功，其它非空值 = 失败。
2. **`Success` 兼容**：仅当 `Success` **显式为 `false`** 时才视为失败；为 `nil`（未下发）按成功处理。
3. **统一写法模板**：

```go
code := resp.Body.GetCode()
successPtr := resp.Body.Success // *bool
if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
    msg := resp.Body.GetMessage()
    return fmt.Errorf("[ERROR] Failed to xxx: Code=%s, Message=%s", code, msg)
}
```

4. **RequestID 仍要先打印**：`[INFO] XxxRequestID: ...` 在判定之前输出，便于排障。

**如何判断当前接口是否需要切换到本 SOP**：

- 真机调用一次接口（开发环境或预发），观察响应 body 是否包含 `Success` 字段。
- 若**不包含**或**不稳定**（部分场景缺失），按本 SOP 写；
- 若**稳定包含**，沿用 skill 模板 `GetSuccess()` 写法亦可。

**参考实现**：

- [cmd/apikey_delete.go](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/apikey_delete.go) —— `Success` 字段缺失场景下的标准成功判定写法

**检查清单（仅对接全新接口时适用）**：

- [ ] 已确认服务端响应是否稳定下发 `Success`
- [ ] 不下发的：采用 `Code != "ok"` 为失败的统一写法
- [ ] 下发的：可继续使用 `GetSuccess()`，但建议同时兼容 `Code` 判定
- [ ] 失败信息包含 `Code` 与 `Message`，并先打印 `RequestID`

---

### 单元测试

- 所有新增的 CLI 命令都必须有对应的单元测试
- 测试文件放在 `test/unit/cmd/` 目录下
- 测试函数命名：`Test<命令名称>Cmd`
- 测试内容至少包括：
  - 命令元数据验证（Use, Short, Long）
  - 必填参数验证
  - 子命令结构验证

### CLI 命令设计

- 使用命名参数（`--name`, `--api-key-id`）而非位置参数
- 将相关功能组织为子命令（如 `apikey create`, `apikey concurrency set`）
- 提供清晰的错误提示和使用示例

### 📋 分页参数 Flag 设计规范

**适用场景**：凡接口支持分页（页码式或游标式），CLI 层**必须**暴露分页参数，不得返回全量数据。

#### 页码式分页（PageNumber / PageStart 类接口）

| Flag     | 类型 | 默认值 | 短参数 | API 字段映射              |
| -------- | ---- | ------ | ------ | ------------------------- |
| `--page` | int  | 1      | 不加   | `PageNo` / `PageStart` 等 |
| `--size` | int  | 10     | 不加   | `PageSize`                |

统一使用 `Int`（不带短参数）注册：

```go
cmd.Flags().Int("page", 1, "Page number (default: 1)")
cmd.Flags().Int("size", 10, "Page size (default: 10)")
```

传参逻辑：

```go
page, _ := cmd.Flags().GetInt("page")
size, _ := cmd.Flags().GetInt("size")
if size > 0 {
    sizeInt32 := int32(size)
    req.PageSize = &sizeInt32
}
if page > 0 {
    pageInt32 := int32(page)
    req.PageStart = &pageInt32  // 或 req.PageNo = &pageInt32，视接口而定
}
```

> **历史遗留说明**：`image list` 使用了 `-p`/`-s` 短参数（`IntP`），属于早期不一致写法，**新命令不跟进，统一不加短参数**。

#### 游标式分页（Token 类接口）

| Flag            | 类型   | 默认值 | 说明                     |
| --------------- | ------ | ------ | ------------------------ |
| `--max-results` | int32  | 10     | 每次返回条数             |
| `--next-token`  | string | —      | 上次返回的游标，首次不传 |

参考实现：`cmd/apikey_list.go`

#### JSON 输出字段要求

支持分页的命令在 `--output json` 时，**必须**在顶层输出分页元信息：

- 页码式：`pageNumber`、`pageSize`
- 游标式：`nextToken`（`omitempty`）

```go
type outputJSON struct {
    TotalCount int        `json:"totalCount"`
    PageNumber int        `json:"pageNumber"`
    PageSize   int        `json:"pageSize"`
    Items      []itemJSON `json:"items"`
}
```

#### 检查清单

- [ ] 页码式接口使用 `--page`（默认 1）+ `--size`（默认 10），不加短参数
- [ ] 游标式接口使用 `--max-results`（默认 10）+ `--next-token`
- [ ] JSON 输出包含分页元信息字段（`pageNumber`/`pageSize` 或 `nextToken`）
- [ ] 单元测试验证 flag 默认值（`page` → `"1"`，`size` → `"10"`）
- [ ] **必查 SDK 序列化**：阅读 `internal/client/client.go` 中对应的 `XxxWithOptions` 函数，确认新增字段已加入 `body["FieldName"] = request.FieldName` 手动赋值（本项目 SDK **不**自动反射序列化，request model 加字段不等于会被发出去）

### ⚠️ 破坏性操作必须设计二次确认与 --yes 跳过（重要！）

**规则**：凡命令会导致**不可逆的数据变更**（删除、永久停用、批量覆盖等），**必须**同时实现：

1. **二次确认提示**（默认启用）：在执行前向用户展示操作对象信息并要求输入确认
2. **`--yes` / `-y` 跳过参数**：允许脚本/CI 场景绕过所有交互提示

**判断标准 —— 以下情况必须加确认**：

| 场景                           | 示例命令                        | 是否需要确认  |
| ------------------------------ | ------------------------------- | ------------- |
| 永久删除资源                   | `apikey delete`, `image delete` | ✅ 必须       |
| 状态前置依赖（如先禁用再删除） | `apikey delete` 遇到 ENABLED    | ✅ 每步都需要 |
| 批量覆盖/清空                  | 未来的批量删除命令              | ✅ 必须       |
| 可逆的状态变更                 | `apikey enable/disable`         | ❌ 不需要     |
| 只读查询                       | `image list`, `image status`    | ❌ 不需要     |

**实现规范**：

```go
// 1. 注册 flag（在 init() 中）
cmd.Flags().BoolP("yes", "y", false, "Skip all confirmation prompts (for non-interactive use)")

// 2. 读取 flag
autoYes, _ := cmd.Flags().GetBool("yes")

// 3. 所有确认点均复用 cmd/confirm.go 的 ConfirmPrompt()
confirmed, err := ConfirmPrompt("Are you sure? [y/N]: ", autoYes)
if err != nil {
    return fmt.Errorf("[ERROR] %w", err)  // 非 TTY 且未传 --yes 时报错
}
if !confirmed {
    fmt.Printf("[INFO] Operation cancelled.\n")
    return nil
}
```

**`ConfirmPrompt` 行为**（`cmd/confirm.go`）：

| 条件                     | 行为                                                                   |
| ------------------------ | ---------------------------------------------------------------------- |
| `autoYes=true`           | 直接返回 true，跳过提示                                                |
| 交互式终端（TTY）        | 打印提示，读取输入（仅 y/Y/yes/YES 确认）                              |
| 非 TTY + `autoYes=false` | 返回错误：`non-interactive environment detected: use --yes to confirm` |

**多步骤命令**（如先禁用再删除）：每个关键步骤单独调用 `ConfirmPrompt`，`autoYes` 透传，**一个 `--yes` 跳过全部**。

**单元测试要求**：在 `test/unit/cmd/` 的测试中必须验证：

```go
// 验证 --yes flag 存在且配置正确
yesFlag := deleteCmd.Flags().Lookup("yes")
assert.NotNil(t, yesFlag)
assert.Equal(t, "false", yesFlag.DefValue)
assert.Equal(t, "y", yesFlag.Shorthand)
```

**参考实现**：

- `cmd/apikey_delete.go` —— 多步骤确认（先禁用确认 + 最终删除确认）
- `cmd/image.go` `runImageDelete` —— 单步骤确认
- `cmd/confirm.go` `ConfirmPrompt` —— 可复用的确认函数

### 新增、修改或删除命令必须同步更新文档和测试用例

> 文档更新的具体操作流程参见 `update-cli-command-docs` skill。以下为规则概要和检查清单。

> ⚡ **强制前置动作**：任何涉及 CLI 命令的需求（新增 / 修改 / 删除，包括仅调整输出字段、参数名称、默认值等细微改动），
> **必须在创建 todo 列表时就把以下两条文档任务纳入**，不得等到代码写完后才想起来：
>
> ```
> - [ ] 更新 docs/en/<command-group>.md 和 docs/zh/<command-group>.md
> - [ ] 视需要更新 README.md 和 README.zh-CN.md Command Overview 表格
> - [ ] 如修改 README.md 或 docs/en/**，执行 bash scripts/build-llms-full.sh 并同步 llms-full.txt
> ```
>
> **禁止在文档任务完成前宣告需求开发完成。**

**规则**: 每次**新增、修改或删除** CLI 命令（包括新增参数、修改默认值、调整输出格式、删除命令/子命令等）时，**必须**同步完成以下工作：

**各场景文档更新范围速查**：

| 变更类型                           | README 命令表格 | docs/<group>.md 输出/参数说明 |
| ---------------------------------- | :-------------: | :---------------------------: |
| 新增命令 / 子命令                  |   ✅ 必须更新   |      ✅ 必须新增完整说明      |
| 修改参数名 / 默认值 / 必填性       |     视情况      |        ✅ 必须同步修改        |
| 调整命令输出（新增/删除/改名字段） |  ❌ 通常不需要  |    ✅ 必须更新 Output 示例    |
| 删除命令 / 子命令                  | ✅ 必须删除条目 |      ✅ 必须删除对应章节      |
| 仅修改内部实现，用户无感知         |    ❌ 不需要    |           ❌ 不需要           |

1. **更新 `README.md` 和 `README.zh-CN.md`**
   - 更新 Command Overview 表格，添加或修改对应命令的说明
   - 保持中英文 README 表格内容一致

2. **更新 `docs/en/<command-group>.md` 和 `docs/zh/<command-group>.md`**
   - 命令组与文件对应关系：`core` / `image` / `apikey` / `network` / `skills` / `docker`
   - 新增命令：在对应命令组文件中添加完整的语法、参数、示例和输出说明
   - 修改命令：同步更新参数说明、示例和注意事项
   - **中英文文档必须同步更新**，保持结构一致

3. **CHANGELOG readiness** — 发布前由 release-prep 统一生成
   - 日常 CLI 需求开发不再手动运行 `git-cliff -o CHANGELOG.md` 全量覆盖
   - 本次变更必须使用可被 git-cliff 识别的 Conventional Commit（如 `feat(apikey): ...`、`fix(image): ...`）
   - 发版前按 `bilingual-changelog-release` skill 执行 `make release-prep VERSION=X.Y.Z`，生成双语版本段并补齐中文翻译
   - 发布前必须确保 `CHANGELOG.md` 无 `TRANSLATE_ME` / `中文翻译待补充` 占位符残留

4. **同步更新对外文档**
   - 钉钉文档（对外使用手册）和 `cli-analysis/Agentbay cli 使用手册.md` 需同步更新
   - 对外文档遵循精简原则：仅保留客户需要的功能说明，剔除内部实现细节
   - 文档内容包括：语法、参数、示例、输出说明、注意事项

5. **编写/更新单元测试**
   - 在 `test/unit/cmd/` 下创建或更新对应的测试文件
   - 测试内容必须覆盖：命令元数据、必填参数校验、子命令结构
   - 运行 `go test ./... -count=1` 确保全部通过

**检查清单（任务结束前逐项核对，全部完成才能宣告需求完成）**:

- [ ] 命令代码已完成（新增、修改或删除）
- [ ] `docs/en/<command-group>.md` 已更新（输出字段 / 参数 / 示例）
- [ ] `docs/zh/<command-group>.md` 已更新（与英文版保持结构一致）
- [ ] `README.md` Command Overview 表格已更新（仅命令结构变化时）
- [ ] `README.zh-CN.md` Command Overview 表格已更新（仅命令结构变化时）
- [ ] 如修改 `README.md` 或 `docs/en/**`，已执行 `bash scripts/build-llms-full.sh` 并同步 `llms-full.txt`
- [ ] 如新增 / 删除 / 重命名对外文档，已同步 `llms.txt` 导航链接
- [ ] 对外文档已同步（钉钉文档 / cli 使用手册）
- [ ] 单元测试已编写或更新并通过
- [ ] mock 类已同步更新（如有接口变更）
- [ ] `go build -o agentbay .` 构建出可执行二进制并通过
- [ ] `go test ./... -count=1` 全部通过
- [ ] commit message 符合 Conventional Commits，确保 release-prep 可正确生成 CHANGELOG
- [ ] 如本次任务包含发版：已执行 `bilingual-changelog-release` skill，`CHANGELOG.md` 无 `TRANSLATE_ME` / `中文翻译待补充` 残留
- [ ] `update-cli-command-docs` skill 已执行（或已完成等效的手动文档同步）

---

## 🔨 构建验证规则

### 重要：需求开发完成后必须构建二进制

**规则**: 每次完成需求开发（新增功能、修复 bug、修改命令等）后，**必须**执行以下构建命令生成可执行二进制：

```bash
go build -o agentbay .
```

**执行时机**:

- ✅ 代码修改完成、单元测试通过之后
- ✅ 在询问用户是否提交代码**之前**
- ✅ 确保构建成功后再告知用户开发完成

**注意事项**:

- 不要只用 `go build ./...` 做编译检查，必须用 `go build -o agentbay .` 生成实际的可执行文件
- 构建产物 `agentbay` 已在 `.gitignore` 中，不会被提交
- 如果构建失败，必须先修复问题再继续

**完整的开发完成验证流程**:

```bash
# 1. 单元测试
go test ./... -count=1

# 2. 构建二进制
go build -o agentbay .

# 3. 确认构建产物
ls -lh agentbay
```

---

## 📂 项目结构

```
agentbay-cli/
├── cmd/                              # CLI 命令定义
│   ├── apikey.go                     # API Key 相关命令
│   ├── concurrency.go                # 并发设置命令
│   ├── network.go                    # 网络管理命令
│   └── ...
├── internal/
│   ├── agentbay/                     # 客户端接口层
│   │   └── client.go
│   └── client/                       # SDK 层
│       ├── client.go
│       ├── create_api_key_*.go
│       └── modify_mcp_api_key_config_*.go
├── test/
│   └── unit/
│       └── cmd/                      # 命令单元测试
│           └── apikey_cmd_test.go
└── .qoder/
    └── rules/                        # Qoder 规则
        └── development.md            # 本文件
```

---

## 🔐 认证方式

CLI 支持两种认证方式：

1. **OAuth 登录**（推荐本地开发）

   ```bash
   agentbay login
   ```

2. **AK/SK 环境变量**（推荐脚本和 CI/CD）
   ```bash
   export AGENTBAY_ACCESS_KEY_ID="your-access-key-id"
   export AGENTBAY_ACCESS_KEY_SECRET="your-access-key-secret"
   ```

**优先级**: AK/SK > OAuth Token

---

## 🌍 环境配置

通过 `AGENTBAY_ENV` 环境变量切换环境：

- `production`（默认）: 生产环境
- `prerelease`: 预发环境
- `international`: 国际站

```bash
# 生产环境
agentbay apikey create --name "my-key"

# 预发环境
AGENTBAY_ENV=prerelease agentbay apikey create --name "my-key"
```

---

## ⚠️ 注意事项

1. **Product ID 映射**:
   - agent-bay（前端）: `xiaoying-double-centre`
   - agentbay-cli（CLI）: `xiaoying`

2. **API 配置**:
   - Version: `2025-05-01`
   - Endpoint (生产): `xiaoying.cn-shanghai.aliyuncs.com`
   - Endpoint (预发): `xiaoying-pre.cn-hangzhou.aliyuncs.com`

3. **CreateApiKey 响应格式**:
   - `Data` 字段是**字符串类型**（直接就是 KeyId），不是对象

---

## 📔 分页接口测试规范（适用于所有 Subagent）

**规则**：凡涉及分页参数（`--page` / `--size` 或类似参数）的测试用例，**必须同时测试第一页和至少一个后续页**（第二页或第 N 页）。

**判定要求**：

| 页码   | 必须验证的内容                                                            |
| ------ | ------------------------------------------------------------------------- |
| 第一页 | `pageNumber=1`，返回数据符合 pageSize，`totalPage >= 2`（确认确实有多页） |
| 第二页 | `pageNumber=2`，返回非空数据，与第一页数据 **不重复**                     |

**预先条件**：

- 如果测试前总数据量 < 2 页，必须先创建足够的测试数据确保有第二页（例如推送够多的技能让 `size=2` 时有至少 3 条以上的数据）
- 不得用已就存在的数据凑幸（即：用例本身要保证走到第二页的条件可控）

**FAIL 判定：**

- 第二页返回空列表 / 报错
- 第二页返回的任意一条数据与第一页重复
- `pageNumber` 字段值与请求不一致

**适用范围**：包括但不限于 `skills list`、`apikey list`、`image list` 等任何支持分页的命令的 Subagent 测试。
