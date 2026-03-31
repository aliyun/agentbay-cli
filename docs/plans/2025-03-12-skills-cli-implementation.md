# Skills CLI 实现方案

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** 在 agentbay-cli 中新增 `agentbay skills` 与 `agentbay skills group` 子命令，实现 Skill 与 SkillGroup 的完整管理（推送、列表、详情、删除、元数据；组的创建/更新/列表/详情/删除及组内 Skill 的增删与启用/禁用），并对接现有后端 Market Skill 相关 API。

**Architecture:** 沿用现有 image 命令的 Cobra 子命令结构；在 `internal/client` 中新增或由代码生成器生成 Market Skill/Group 的 Request/Response 模型及 API 方法；在 `internal/agentbay` 中扩展 Client 接口与 clientWrapper 实现；在 `cmd/` 下新增 `skills.go`（及可选 `skills_group.go`）注册 `skills` / `skills group` 命令并调用 agentbay client。Skill push 流程参考 image create：先调 GetMarketSkillCredential 获取 OSS 上传凭证，打包 skill 目录为 zip 上传 OSS，再调 CreateMarketSkill 传入 OssBucket/OssFilePath。

**Tech Stack:** Go, Cobra, 现有 internal/client (Alibaba Cloud OpenAPI 风格), internal/agentbay, internal/config, internal/auth。

---

## 一、CLI 命令一览

以下为需求中约定的每条命令的用法与说明。

### 2.1 Skill 管理

| 需求描述 | 命令 | 说明 |
|----------|------|------|
| 推送本地 Skill 目录到云端 | `agentbay skills push <skill-dir>` | 读取 SKILL.md 的 name/description，打包 zip 上传，同名则更新 |
| 列出云端 Skills | `agentbay skills list` | 显示自己的 + public 的 |
| 查看 Skill 详情 | `agentbay skills show <skill-id>` | 展示单个 Skill 详情 |
| 删除 Skill | `agentbay skills delete <skill-id>` | 删除指定 Skill |
| 获取元数据（不启动沙箱） | `agentbay skills metadata [--group <group-id>...] [--include-content] [--format json\|table]` | 不传参数则返回当前用户所有可见 Skills；用于调试/预览注册到 LLM 的信息 |

**每条命令的 CLI 形式：**

```bash
# 推送（自动打包 + 上传 + 创建/更新）
agentbay skills push <skill-dir>

# 示例
agentbay skills push ./my-skill
agentbay skills push /path/to/my-skill/

# 列出云端 Skills
agentbay skills list

# 查看详情
agentbay skills show <skill-id>
# 示例
agentbay skills show sk-abc123

# 删除
agentbay skills delete <skill-id>
# 示例
agentbay skills delete sk-abc123

# 元数据（可选筛选与输出格式）
agentbay skills metadata
agentbay skills metadata --group grp-001 --group grp-002
agentbay skills metadata --include-content --format json
agentbay skills metadata --format table
```

### 2.2 SkillGroup 管理

所有 SkillGroup 相关命令均在 `agentbay skills group` 下。

| 需求描述 | 命令 | 说明 |
|----------|------|------|
| 创建 Group | `agentbay skills group create <name> [--description "..."]` | 创建技能组 |
| 更新 Group | `agentbay skills group update <group-id> [--description "..."]` | 更新描述等 |
| 列出 Groups | `agentbay skills group list` | 列出当前用户的技能组 |
| 查看 Group 详情 | `agentbay skills group show <group-id>` | 显示包含的 Skills 及启用状态 |
| 删除 Group | `agentbay skills group delete <group-id>` | 删除技能组 |
| 组内添加 Skill | `agentbay skills group add-skill <group-id> <skill-id>` | 将 Skill 加入组 |
| 组内移除 Skill | `agentbay skills group remove-skill <group-id> <skill-id>` | 从组中移除 Skill |
| 启用组内 Skill | `agentbay skills group enable-skill <group-id> <skill-id>` | 在组内启用该 Skill |
| 禁用组内 Skill | `agentbay skills group disable-skill <group-id> <skill-id>` | 在组内禁用该 Skill（不移除） |

**每条命令的 CLI 形式：**

```bash
# 创建 Group（description 可选）
agentbay skills group create doc-processing
agentbay skills group create doc-processing --description "Document processing skills"

# 更新 Group
agentbay skills group update <group-id> --description "New description"

# 列出 Groups
agentbay skills group list

# 查看 Group 详情
agentbay skills group show <group-id>
# 示例
agentbay skills group show grp-001

# 删除 Group
agentbay skills group delete <group-id>

# 管理组内 Skill
agentbay skills group add-skill <group-id> <skill-id>
agentbay skills group remove-skill <group-id> <skill-id>
agentbay skills group enable-skill <group-id> <skill-id>
agentbay skills group disable-skill <group-id> <skill-id>

# 示例
agentbay skills group add-skill grp-001 sk-001
agentbay skills group remove-skill grp-001 sk-003
agentbay skills group enable-skill grp-001 sk-002
agentbay skills group disable-skill grp-001 sk-003
```

---

## 二、后端 API 与 CLI 映射

### 2.1 已提供的后端接口

| Method | Pop Action | 用途 | CLI 命令 |
|--------|------------|------|----------|
| GET | GetMarketSkillCredential | 获取 Skill 上传凭证（OSS） | `skills push`（上传前） |
| GET | CreateMarketSkill | 创建 Skill（入参 OssBucket, OssFilePath） | `skills push`（上传后） |
| GET | DescribeMarketSkillDetail | 查询 Skill 详情（SkillId） | `skills show <skill-id>` |
| POST | CreateMarketSkillGroup | 创建技能组（GroupName） | `skills group create <name>` |
| POST | ListMarketGroupSkill | 列表（当前为 Group 列表） | `skills group list` |
| POST | AddMarketGroupSkill | 组内添加技能（GroupId, SkillId） | `skills group add-skill` |
| POST | RemoveMarketGroupSkill | 组内移除技能（GroupId, SkillId） | `skills group remove-skill` |

### 2.2 需求与 API 缺口

- **skills list**：需后端提供「列出当前用户 + 可见 public Skills」的接口（如 ListMarketSkills）。若暂无，可先实现为调用 DescribeMarketSkillDetail 的占位或返回“接口待后端提供”的友好提示。
- **skills delete**：需后端提供 DeleteMarketSkill(SkillId)。若暂无，可先做参数校验与占位错误提示。
- **skills metadata**：可由 list + show 组合实现；`--group` 筛选依赖「按 Group 查 Skill 列表」或 ListMarketGroupSkill 返回组内 Skill 列表，需确认 ListMarketGroupSkill 响应结构（是否含组内 Skill 列表）。
- **skills group create**：后端仅支持 GroupName；CLI 的 `--description` 若后端暂无字段，可先保留参数并在请求中省略或后续扩展。
- **skills group update**：需后端提供 UpdateMarketSkillGroup(GroupId, Description?)。若暂无，可先做占位。
- **skills group show**：需后端提供 DescribeMarketSkillGroup(GroupId) 或由 ListMarketGroupSkill + 组内 Skill 列表接口组合。需确认 ListMarketGroupSkill 是否返回组详情及组内 Skills。
- **skills group delete**：需后端提供 DeleteMarketSkillGroup(GroupId)。若暂无，可先做占位。
- **skills group enable-skill / disable-skill**：需后端提供组内 Skill 启用/禁用接口（或同一接口带状态）。若暂无，可先做占位。

实现顺序建议：先实现与已有 API 一一对应的命令（push、show、group create、group list、add-skill、remove-skill），再为 list/delete/metadata/group update/show/delete/enable/disable 做占位或在后端就绪后补齐。

### 2.3 可实现与暂不可实现汇总

**当前可实现（后端已有对应 API）：**

| CLI 命令 | 依赖的后端 API | 说明 |
|----------|----------------|------|
| `agentbay skills push <skill-dir>` | GetMarketSkillCredential + CreateMarketSkill | 获取 OSS 凭证 → 打包 zip 上传 → 创建/更新 Skill |
| `agentbay skills show <skill-id>` | DescribeMarketSkillDetail | 查询单个 Skill 详情 |
| `agentbay skills group create <name> [--description "..."]` | CreateMarketSkillGroup | 创建技能组（description 若后端未支持可先不传） |
| `agentbay skills group list` | ListMarketGroupSkill | 列出当前用户技能组 |
| `agentbay skills group add-skill <group-id> <skill-id>` | AddMarketGroupSkill | 将 Skill 加入组 |
| `agentbay skills group remove-skill <group-id> <skill-id>` | RemoveMarketGroupSkill | 从组中移除 Skill |

**暂不可实现（需后端补齐接口后实现或先做占位）：**

| CLI 命令 | 缺失的后端 API | 建议 |
|----------|----------------|------|
| `agentbay skills list` | ListMarketSkills（列出用户 + 可见 public Skills） | 先做占位提示「列表接口即将支持」 |
| `agentbay skills delete <skill-id>` | DeleteMarketSkill(SkillId) | 先做参数校验 + 占位提示 |
| `agentbay skills metadata [--group ...] [--include-content] [--format ...]` | 依赖 list + 按 group 筛选能力；若 ListMarketGroupSkill 返回组内 Skill 可部分实现 | 有 list 后可组合实现；否则占位 |
| `agentbay skills group update <group-id> [--description "..."]` | UpdateMarketSkillGroup(GroupId, Description?) | 占位 |
| `agentbay skills group show <group-id>` | DescribeMarketSkillGroup(GroupId) 或 ListMarketGroupSkill 含组详情与组内 Skills | 需确认 ListMarketGroupSkill 响应结构；否则占位 |
| `agentbay skills group delete <group-id>` | DeleteMarketSkillGroup(GroupId) | 占位 |
| `agentbay skills group enable-skill <group-id> <skill-id>` | 组内 Skill 启用接口 | 占位 |
| `agentbay skills group disable-skill <group-id> <skill-id>` | 组内 Skill 禁用接口 | 占位 |

**小结：** 共 6 个 CLI 命令/子命令可立即实现，8 个需占位或等后端接口就绪后补齐。

### 2.4 需求 a 符合性

**需求 a：** CLI — Skill 新增@（upload，字段可缺省）、查询；Skill Group 新增、查询、绑定解绑 Skill。

| 子项 | 需求 | 对应 CLI | 是否满足 | 备注 |
|------|------|----------|----------|------|
| Skill 上传 | upload，字段可缺省 | `agentbay skills push <skill-dir>` | ✅ 满足 | 仅必填 `<skill-dir>`，name/description 从 SKILL.md 读取（可缺省），无额外必填字段 |
| Skill 查询 | 查询 | `agentbay skills show <skill-id>`、`agentbay skills list` | ✅ 满足 | 单条查询用 show（已可实现）；列表用 list（后端接口就绪后可实现，当前可做占位） |
| Skill Group 新增 | 新增 | `agentbay skills group create <name> [--description "..."]` | ✅ 满足 | 已有对应 API，description 可选 |
| Skill Group 查询 | 查询 | `agentbay skills group list`、`agentbay skills group show <group-id>` | ✅ 满足 | 列表用 list（已可实现）；单组详情用 show（后端就绪或 List 响应含组内 Skills 后可实现，当前可占位） |
| Skill Group 绑定/解绑 Skill | 绑定解绑 | `agentbay skills group add-skill <group-id> <skill-id>`、`agentbay skills group remove-skill <group-id> <skill-id>` | ✅ 满足 | 已有 AddMarketGroupSkill / RemoveMarketGroupSkill，可立即实现 |

**结论：** 需求 a 已覆盖。当前方案中与 a 相关的 6 个能力均有对应 CLI 设计；其中 6 项（push、show、group create、group list、add-skill、remove-skill）可立即实现，list / group show 可先做占位，待后端接口就绪后补齐。

---

## 三、实现任务（Bite-Sized）

以下按「先 client/API，再 skills 命令，再 group 命令」拆分为小步，每步可独立运行测试并提交。

### Task 1: 添加 Market Skill 相关 Request/Response 模型与 GetMarketSkillCredential

**Files:**
- Create: `internal/client/get_market_skill_credential_request_model.go`
- Create: `internal/client/get_market_skill_credential_response_model.go`（含 Body/Data，含 OSS 上传 URL 或 Bucket+Path+Credential，与现有 GetDockerFileStoreCredential 结构类似）
- Modify: `internal/client/client.go` — 新增 `GetMarketSkillCredentialWithOptions` / `GetMarketSkillCredential`，Action 为 `GetMarketSkillCredential`
- Modify: `internal/client/client_context_func.go` — 新增带 Context 的 GetMarketSkillCredential 方法（若项目统一用 Context）

**Step 1:** 参考 `get_docker_file_store_credential_*` 的字段与校验，新增 GetMarketSkillCredential 的 Request（若无入参则为空或仅登录态）/Response（Body.Data 含 OssUrl 或等价上传信息）。

**Step 2:** 在 client.go 中增加 CallApi，Action 为 `GetMarketSkillCredential`，Method/Protocol/Pathname 与现有接口一致（如 POST、HTTPS、"/"）。

**Step 3:** 运行 `go build ./...`，确认编译通过。

**Step 4:** Commit: `feat(client): add GetMarketSkillCredential request/response and API method`

---

### Task 2: 添加 CreateMarketSkill API

**Files:**
- Create: `internal/client/create_market_skill_request_model.go`（OssBucket, OssFilePath string）
- Create: `internal/client/create_market_skill_response_model.go`（含 SkillId 等）
- Modify: `internal/client/client.go` — 新增 CreateMarketSkillWithOptions / CreateMarketSkill，Action `CreateMarketSkill`

**Step 1:** 定义 Request（OssBucket, OssFilePath）、Response（SkillId 等），与后端文档一致。

**Step 2:** 在 client.go 中实现 CreateMarketSkill 的 CallApi。

**Step 3:** `go build ./...`

**Step 4:** Commit: `feat(client): add CreateMarketSkill API`

---

### Task 3: 添加 DescribeMarketSkillDetail API

**Files:**
- Create: `internal/client/describe_market_skill_detail_request_model.go`（SkillId）
- Create: `internal/client/describe_market_skill_detail_response_model.go`
- Modify: `internal/client/client.go` — 新增 DescribeMarketSkillDetail

**Step 1–3:** 同前，Action 为 `DescribeMarketSkillDetail`。

**Step 4:** Commit: `feat(client): add DescribeMarketSkillDetail API`

---

### Task 4: 添加 SkillGroup 相关 API（Create/List/Add/Remove）

**Files:**
- Create: `internal/client/create_market_skill_group_request_model.go`（GroupName）
- Create: `internal/client/create_market_skill_group_response_model.go`（GroupId）
- Create: `internal/client/list_market_group_skill_request_model.go`（若需分页再扩展）
- Create: `internal/client/list_market_group_skill_response_model.go`（与后端 JSON 一致，如 Data: []{GroupId, GroupName}）
- Create: `internal/client/add_market_group_skill_request_model.go`（GroupId, SkillId）
- Create: `internal/client/add_market_group_skill_response_model.go`
- Create: `internal/client/remove_market_group_skill_request_model.go`（GroupId, SkillId）
- Create: `internal/client/remove_market_group_skill_response_model.go`
- Modify: `internal/client/client.go` — 新增 CreateMarketSkillGroup, ListMarketGroupSkill, AddMarketGroupSkill, RemoveMarketGroupSkill（POST，响应 BodyType 可能为 JSON，需与现有 XML 解析方式区分或统一）

**Step 1:** 按后端文档定义各 Request/Response 字段。

**Step 2:** 在 client.go 中实现上述四个 Action 的 CallApi；若后端返回 JSON，使用 bodyType JSON 并反序列化到对应 Response。

**Step 3:** `go build ./...`

**Step 4:** Commit: `feat(client): add Market SkillGroup APIs (create, list, add-skill, remove-skill)`

---

### Task 5: 在 agentbay Client 中扩展 Skill 与 Group 接口

**Files:**
- Modify: `internal/agentbay/client.go` — 在 Client interface 中增加 GetMarketSkillCredential, CreateMarketSkill, DescribeMarketSkillDetail, CreateMarketSkillGroup, ListMarketGroupSkill, AddMarketGroupSkill, RemoveMarketGroupSkill；在 clientWrapper 中实现上述方法（调用 internal/client 的对应方法，需处理 XML/JSON 与 fallback 解析，参考现有 GetDockerFileStoreCredential）。

**Step 1:** 扩展 Client interface。

**Step 2:** 实现 clientWrapper 各方法（getClient + 请求构造 + 调用 SDK + 响应解析）。

**Step 3:** 若有现有单元测试依赖 Client 接口，补充 mock 或跳过 Skill 方法；`go build ./...` 通过。

**Step 4:** Commit: `feat(agentbay): extend client with Market Skill and Group APIs`

---

### Task 6: 实现 skills push

**Files:**
- Create: `cmd/skills.go` — 定义 `SkillsCmd`，子命令 `push`；init 中注册 `push`。
- Modify: `main.go` — `rootCmd.AddCommand(cmd.SkillsCmd)`

**Step 1:** 实现 push 逻辑：解析 `<skill-dir>`，读取 SKILL.md 的 name/description（可参考 Cursor skills 的 frontmatter）；校验目录存在且含 SKILL.md。

**Step 2:** 调用 GetMarketSkillCredential 获取 OSS 上传凭证；将 skill 目录打包为 zip（仅包含 SKILL.md 及需上传的文件，避免无关文件）；使用凭证上传 zip 到 OSS（参考 cmd/image.go 的 uploadFileToOSS）。

**Step 3:** 调用 CreateMarketSkill，传入返回的 OssBucket、OssFilePath；若后端支持「同名更新」，则根据 name 先查是否已存在（若已有 ListMarketSkills 或 Describe 按 name 查），再决定创建或走更新逻辑；否则先仅实现「创建」路径。

**Step 4:** 输出成功信息，如：`Skill "my-skill" created (skill-id: sk-abc123)`。

**Step 5:** 单元测试：至少测试 push 的参数校验（无目录、无 SKILL.md、SKILL.md 无 name）导致失败。`go test ./cmd/... -run SkillsPush -v`

**Step 6:** Commit: `feat(cmd): add agentbay skills push`

---

### Task 7: 实现 skills list（含占位）

**Files:**
- Modify: `cmd/skills.go` — 新增子命令 `list`，RunE 中若后端暂无 ListMarketSkills，则打印友好提示“列表接口即将支持”并 exit 0 或 1（按产品约定）；若已有接口，则调用并表格/列表输出。

**Step 1:** 添加 `skills list` 子命令，Short/Long 描述与需求一致。

**Step 2:** 若后端已提供列表 API，则调用并格式化输出；否则实现占位提示。

**Step 3:** `go build ./...`，手动执行 `agentbay skills list` 验证。

**Step 4:** Commit: `feat(cmd): add agentbay skills list (or placeholder)`

---

### Task 8: 实现 skills show 与 skills delete

**Files:**
- Modify: `cmd/skills.go` — 新增 `show <skill-id>`、`delete <skill-id>`；show 调用 DescribeMarketSkillDetail 并输出详情；delete 若后端有 DeleteMarketSkill 则调用，否则占位提示。

**Step 1:** show：Args ExactArgs(1)，调用 DescribeMarketSkillDetail(args[0])，打印表格或 JSON。

**Step 2:** delete：Args ExactArgs(1)，若后端有接口则调用并提示已删除；否则占位。

**Step 3:** `go build ./...`，手动验证。

**Step 4:** Commit: `feat(cmd): add agentbay skills show and delete`

---

### Task 9: 实现 skills metadata

**Files:**
- Modify: `cmd/skills.go` — 新增子命令 `metadata`，Flags: `--group` ([]string)，`--include-content` (bool)，`--format` (string, default "table", values json|table)。

**Step 1:** 不传 `--group` 时，等价于「当前用户所有可见 Skills」：若有 list 接口则用 list，否则用占位或空列表；传 `--group` 时，按 group-id 筛选（依赖 ListMarketGroupSkill 返回组内 Skill 或单独接口）。

**Step 2:** `--include-content` 控制是否在输出中包含 Skill 正文内容（若 DescribeMarketSkillDetail 返回）。

**Step 3:** `--format json|table` 控制输出格式。

**Step 4:** 单元测试：metadata 的 flag 解析与默认值。Commit: `feat(cmd): add agentbay skills metadata`

---

### Task 10: 实现 skills group 子命令（create, list, add-skill, remove-skill）

**Files:**
- Create 或 Modify: `cmd/skills.go` 或 `cmd/skills_group.go` — 定义 `SkillsGroupCmd`，挂到 `SkillsCmd` 下：`SkillsCmd.AddCommand(SkillsGroupCmd)`；子命令：`group create <name> [--description "..."]`，`group list`，`group add-skill <group-id> <skill-id>`，`group remove-skill <group-id> <skill-id>`。

**Step 1:** group create：Args ExactArgs(1)，name=args[0]，可选 --description；调用 CreateMarketSkillGroup(GroupName)；若后端暂不支持 description，请求中不传或忽略。

**Step 2:** group list：无 args，调用 ListMarketGroupSkill，输出 GroupId、GroupName 等。

**Step 3:** group add-skill / remove-skill：Args ExactArgs(2)，GroupId=args[0]，SkillId=args[1]，分别调用 AddMarketGroupSkill、RemoveMarketGroupSkill。

**Step 4:** `go build ./...`，`agentbay skills group --help` 与各子命令手动验证。

**Step 5:** Commit: `feat(cmd): add agentbay skills group create/list/add-skill/remove-skill`

---

### Task 11: 实现 skills group show / update / delete / enable-skill / disable-skill（占位或补齐）

**Files:**
- Modify: `cmd/skills_group.go` 或 `cmd/skills.go` — 新增子命令 `group show`，`group update`，`group delete`，`group enable-skill`，`group disable-skill`。

**Step 1:** 若后端已有 DescribeMarketSkillGroup、UpdateMarketSkillGroup、DeleteMarketSkillGroup、Enable/Disable 组内 Skill，则在 client 与 agentbay 中补充 API 后在此调用；否则每个子命令先做参数校验并打印“该功能即将支持”类提示。

**Step 2:** 保证 `agentbay skills group --help` 中列出全部子命令，行为与需求一致。

**Step 3:** Commit: `feat(cmd): add skills group show/update/delete/enable-skill/disable-skill (or placeholders)`

---

### Task 12: 文档与集成

**Files:**
- Modify: `docs/USER_GUIDE.md` 或新建 `docs/SKILLS_CLI.md` — 增加 2.1 / 2.2 的典型工作流示例（创建并推送 Skill、组织 Skills 到 Group）。
- Modify: `main.go` — 确认 SkillsCmd 已挂到 root，且 GroupID 归属（如 "management"）。

**Step 1:** 在用户文档中粘贴「典型工作流」示例（与本文第一节一致）。

**Step 2:** 运行 `go test ./...`，修复因新增命令导致的测试失败（如 main_test 需注册 SkillsCmd）。

**Step 3:** Commit: `docs: add Skills CLI usage and workflow`

---

## 四、执行选择

方案写完后可选两种执行方式：

1. **Subagent-Driven（本会话）** — 按任务分发给子 agent，每步完成后你做 review，再进入下一步。
2. **Parallel Session（新会话）** — 在新会话中打开本仓库（或 worktree），使用 superpowers:executing-plans，按任务批量执行并在检查点做 review。

若选 Subagent-Driven，请使用 superpowers:subagent-driven-development；若选新会话执行，请在新会话中引用本计划并采用 executing-plans。
