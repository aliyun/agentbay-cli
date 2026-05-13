---
trigger: always_on
---

# AgentBay CLI 开发规则

## 🔗 Skill 自动装配规则（Quest / 对话 / 任意入口通用）

**凡符合下列任一特征的任务，AI 必须主动加载并遵循对应的 `.qoder/skills/` 指南**（包括但不限于 Quest Design/Execute 阶段、直接对话、Execute Directly 模式）：

| 任务特征                                                    | 必须加载的 Skill                                                                                 | Skill 路径                                            |
| ----------------------------------------------------------- | ------------------------------------------------------------------------------------------------ | ----------------------------------------------------- |
| 新增 / 修改 CLI 命令、参数、子命令，或将前端 API 封装为命令 | **create-cli-command**                                                                           | `.qoder/skills/create-cli-command/SKILL.md`           |
| 涉及分支管理、commit、push、PR、变更档案（Quest/CR 目录）   | **feature-development-workflow**                                                                 | `.qoder/skills/feature-development-workflow/SKILL.md` |
| 新增 CLI 命令类需求（同时触发上述两条）                     | **两者组合使用**：先 workflow 拉分支/建档 → 再 create-cli-command 实现 → 回到 workflow 提交/推送 | 同上                                                  |

**执行铁律**:

1. **不得跳过**：即便用户未显式打 `/skill-name`，只要任务特征命中上表，AI 就必须主动阅读 skill 文件并按其 Phase 执行。
2. **前置检查**：进入实现阶段前，必须确认 `feature-development-workflow` 的 Phase 0（变更档案）和分支创建已完成，否则先引导用户补齐。
3. **Quest 场景**：Quest 生成 spec 后的 Execute 阶段，等同于"对话入口"，本规则照常生效，无需 spec 里额外声明。
4. **Execute Directly 场景**：即便跳过 Design 阶段，AI 也必须在动手前主动加载匹配的 skill。

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
<type>: <description>

[optional body]
```

**Type 类型**:

- `feat`: 新功能
- `fix`: 修复 bug
- `test`: 添加或修改测试
- `docs`: 文档更新
- `refactor`: 代码重构
- `style`: 代码格式调整（不影响功能）
- `chore`: 构建过程或辅助工具的变动

**示例**:

```bash
feat: add API key concurrency management CLI command

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

## 💻 Go 代码规范

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

### 新增或修改命令必须同步更新 README 和测试用例

**规则**: 每次**新增或修改** CLI 命令（包括新增参数、修改默认值、调整输出格式等）时，**必须**同步完成以下工作：

1. **更新 `agentbay-cli/README.md`**
   - 新增命令：在 Features 列表和 Quick Start 示例中补充新命令的使用说明
   - 修改命令：同步更新 README 中对应命令的参数说明、示例和注意事项
   - 保持与已有命令文档风格一致

2. **同步更新对外文档**
   - 钉钉文档（对外使用手册）和 `cli-analysis/Agentbay cli 使用手册.md` 需同步更新
   - 对外文档遵循精简原则：仅保留客户需要的功能说明，剔除内部实现细节
   - 文档内容包括：语法、参数、示例、输出说明、注意事项

3. **编写/更新单元测试**
   - 在 `test/unit/cmd/` 下创建或更新对应的测试文件
   - 测试内容必须覆盖：命令元数据、必填参数校验、子命令结构
   - 运行 `go test ./... -count=1` 确保全部通过

**检查清单**:

- [ ] 命令代码已完成（新增或修改）
- [ ] README.md 已更新，反映最新的命令用法
- [ ] 对外文档已同步（钉钉文档 / cli 使用手册）
- [ ] 单元测试已编写或更新并通过
- [ ] mock 类已同步更新（如有接口变更）
- [ ] `go build` 和 `go test ./...` 均通过

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
