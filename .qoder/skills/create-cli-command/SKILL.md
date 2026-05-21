---
name: create-cli-command
description: 将前端 API 能力封装成 agentbay-cli 命令的标准化流程
---

# Create CLI Command

## 🔗 前置约束（必读）

**本 skill 必须与 `feature-development-workflow` 配套使用**，两者分工如下：

| Skill                            | 负责                                                                                          | 触发时机       |
| -------------------------------- | --------------------------------------------------------------------------------------------- | -------------- |
| **feature-development-workflow** | 分支管理（从 `aliyun/master` 拉 feat 分支）、双远程推送（origin → aliyun）、PR 流程、变更档案 | 开发前、提交后 |
| **create-cli-command**（本文件） | SDK 模型 → Client 接口 → Cobra 命令 → mock 同步 → 单测 → 对客文档                             | 开发中         |

**执行铁律**：

1. **开发前**：必须先执行 `feature-development-workflow` 的 Phase 0（变更档案初始化）和分支创建，确认当前在 `feat-<name>` 分支且基于 `aliyun/master`，否则不得进入本 skill 的 Phase 1。
2. **开发中**：按本 skill 的 Phase 1-5 实现代码和测试。
3. **提交/推送**：切回 `feature-development-workflow` 的 Phase 4-6 完成 commit、双远程 push（origin 先、aliyun 后）、PR、trace.md 更新。
4. **禁止跳过**：不得在未拉 feat 分支时直接在 master 上开发，不得单远程推送，不得跳过变更档案。

> 若用户未提前走 `feature-development-workflow`，本 skill 执行第一步必须主动提醒并引导用户先建档、拉分支。

---

## 📋 职责

将 agent-bay 前端控制台中的 API 能力封装成 agentbay-cli 命令行工具，包括：

- 创建 CLI 命令（使用 Cobra 框架）
- 实现 SDK 客户端方法
- 编写单元测试
- 生成对客文档

## 🎯 触发场景

当用户提出以下需求时触发：

- "帮我把 XX 接口封装成 CLI"
- "新增 XX 功能的命令行工具"
- "把前端的 XX 能力做成 CLI 命令"
- "新增 agentbay xx 命令"

## 🚀 执行步骤

### Phase 1: 需求分析

1. **查找前端 API 实现**
   - 在 `agent-bay/src/http/api/` 目录下搜索接口
   - 确认接口的 Action、Version、Product ID
   - 分析请求参数和响应结构

2. **确认 API 信息**
   - Product ID: `xiaoying`（CLI 使用，不是前端的 xiaoying-double-centre）
   - API Version: `2025-05-01`
   - Endpoint: 根据环境自动选择

3. **与用户确认**
   - 命令名称和层级结构
   - 参数设计（使用命名参数 `--name` 而非位置参数）
   - 是否需要子命令

### Phase 2: 代码实现

#### 2.1 创建 SDK 层模型

```
internal/client/
├── {action}_request_model.go      # 请求模型
└── {action}_response_model.go     # 响应模型
```

**要求**:

- 请求模型必须有 `Validate()` 方法
- 响应模型提供 `GetXxx()` 辅助方法
- 注意后端实际返回的字段类型（可能是字符串而非对象）

#### 2.2 添加 SDK 客户端方法

在 `internal/client/client.go` 中添加：

- `{Action}WithOptions()` - 完整调用方法
- `{Action}()` - 简化调用方法
- `{Action}WithContext()` - 支持 context 的方法
- `parse{Action}Response()` - 响应解析函数

**⚠️ parser 必须放入 `internal/client/dual_format_responses.go`**（而非 `client.go`），且必须遵守下述容错规范：

- 所有 `*int32` / `*int64` 字段用 `json.RawMessage` + `int32FromFlexibleJSON` 解析，兼容数字与字符串两种序列化形式（服务端常会以字符串返回 `HttpStatusCode` 等数字字段）。
- body 以 `<` 开头走 XML 分支、否则走 JSON 分支，两条路径都要调用 `applyMapHeadersAndStatus` 归一 headers / statusCode。
- 解析失败统一用 `&ErrWithRequestID{Err: ..., RequestID: extractRequestIDFromResponse(res)}` 包装。
- 必须在 `internal/client/` 下配套一个 `xxx_parse_test.go`，至少覆盖「JSON 数字字段为字符串 / 数字 / XML」三种场景。

反面案例：`BatchCreateHideResourceGroupsWithMaxSession` 早期直接用 `json.Unmarshal` 打到 `*int32`，遇到 `"HttpStatusCode":"200"` 直接报 `cannot unmarshal string into Go struct field ... of type int32`。详见 [references/api-format.md](references/api-format.md#响应解析容错模板必须使用) 与规则 [development.md 响应解析必须使用容错模板](../../rules/development.md)。

**API 配置模板**:

```go
params := &openapiutil.Params{
    Action:      dara.String("ActionName"),
    Version:     dara.String("2025-05-01"),
    Protocol:    dara.String("HTTPS"),
    Pathname:    dara.String("/"),
    Method:      dara.String("POST"),
    AuthType:    dara.String("AK"),
    Style:       dara.String("RPC"),
    ReqBodyType: dara.String("formData"),
    BodyType:    dara.String("string"),
}
```

#### 2.3 添加客户端接口

在 `internal/agentbay/client.go` 中：

- 在 `Client` interface 中添加方法定义
- 在 `clientWrapper` 中实现方法

**⚠️ 重要：更新接口后必须同步更新所有 Mock 类！**

```bash
# 1. 查找所有实现该接口的 mock 类
grep -r "type mock.*Client struct" cmd/ test/

# 2. 为每个 mock 类添加新方法
# 示例：在 cmd/image_status_helper_test.go 和 cmd/image_list_helper_test.go 中添加
```

```go
// 为每个 mock 类添加（返回 "not implemented"）
func (m *mockClient) NewMethod(ctx context.Context, request *client.NewRequest) (*client.NewResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

**检查清单**:

- [ ] 找到所有 mock 类（通常 2-3 个）
- [ ] 为每个 mock 类添加新方法
- [ ] 运行 `go test ./cmd/...` 确保编译通过

#### 2.4 创建 CLI 命令

在 `cmd/` 目录下创建命令文件：

- 使用命名参数（`--name`, `--api-key-id`）
- 标记必填参数：`MarkFlagRequired()`
- 提供清晰的帮助信息和示例
- 实现错误处理和友好提示
- **每个对外 API 请求必须默认打印 RequestId**（详见下方铁律）

##### 🔑 RequestId 打印铁律（强制）

**规则**: 每个 CLI 命令调用的**每一个对外接口请求**，无论成功还是失败、无论是否带 `-v / --verbose`，都**必须**在终端打印对应的 RequestId，便于客户在出现问题时直接复制日志给运维定位。

**❌ 禁止做法**（之前的旧规范）：

```go
// ❌ 不要再用 verbose 守卫保护 RequestId 打印
verbose, _ := cmd.Flags().GetBool("verbose")
if verbose && resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
    printRequestIDIfVerbose(cmd, *resp.Body.RequestId)
}
```

**✅ 正确做法**：

```go
// 1. 成功路径：直接打印（不依赖 verbose）
if resp != nil && resp.Body != nil {
    if reqId := resp.Body.GetRequestId(); reqId != nil && *reqId != "" {
        fmt.Printf("[INFO] {Action} Request ID: %s\n", *reqId)
    }
}

// 2. 错误路径：err 中携带的 RequestId 也必须打印
resp, err := apiClient.{Action}(ctx, req)
if err != nil {
    // 不再使用 printRequestIDFromErrIfVerbose（带 verbose 守卫的版本）
    if reqId := extractRequestIDFromErr(err); reqId != "" {
        fmt.Printf("[INFO] {Action} Request ID: %s\n", reqId)
    }
    return fmt.Errorf("[ERROR] Failed to {action}: %w", err)
}
```

**多接口命令**: 命令体内若调用多个接口（如先 GetMcpImageInfo 再 BatchCreateXxx），**每一个**接口的 RequestId 都要分别打印，并在前缀里标注接口名以便区分：

```
[INFO] GetMcpImageInfo Request ID: 1A2B3C4D-...
[INFO] BatchCreateHideResourceGroupsWithMaxSession Request ID: 5E6F7G8H-...
```

**参考实现**: [cmd/image_set_max_session.go](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/cmd/image_set_max_session.go#L85-L121)

**verbose / `-v` 的真正用途**: 仅控制额外的调试信息（请求体、响应体 JSON、堆栈等），**不再**控制 RequestId 是否打印。

##### 🛡️ 破坏性操作的二次确认设计（强制）

**规则**：命令涉及**不可逆操作**（删除、永久停用等）时，**必须**同时实现二次确认提示和 `--yes` / `-y` 跳过参数。

**哪些情况触发**：

| 操作类型 | 示例 | 是否需要 |
|---------|------|---------|
| 永久删除资源 | `apikey delete`, `image delete` | ✅ 必须 |
| 多步骤前置依赖（如先禁用才能删除）| 每步都提示 | ✅ 每步 |
| 可逆状态变更 | `enable`, `disable` | ❌ 不需要 |
| 查询/只读操作 | `list`, `status` | ❌ 不需要 |

**标准实现**（复用 `cmd/confirm.go` 中已有的 `ConfirmPrompt`）：

```go
// init() 中注册 flag
apikeyDeleteCmd.Flags().BoolP("yes", "y", false, "Skip all confirmation prompts (for non-interactive use)")

// RunE 中使用
autoYes, _ := cmd.Flags().GetBool("yes")

// 每个确认点调用 ConfirmPrompt，autoYes 透传
confirmed, err := ConfirmPrompt("Are you sure you want to delete? [y/N]: ", autoYes)
if err != nil {
    return fmt.Errorf("[ERROR] %w", err)
}
if !confirmed {
    fmt.Printf("[INFO] Operation cancelled.\n")
    return nil
}
```

**`ConfirmPrompt` 三种行为**：
- `--yes` 传入 → 直接 true，无任何输出
- 交互式 TTY → 打印提示，读取输入（仅 y/Y/yes/YES 通过）
- 非 TTY 且无 `--yes` → 返回错误，提示用户加 `--yes`

**多步骤命令**：每步单独调用 `ConfirmPrompt(prompt, autoYes)`，一个 `--yes` 跳过全部步骤。

**参考实现**：
- `cmd/apikey_delete.go` —— 多步骤（禁用确认 + 删除确认）
- `cmd/image.go` `runImageDelete` —— 单步骤确认

**命令层级**:

```
agentbay
└── apikey                          # 命令组
    ├── create --name <名称>        # 子命令
    └── concurrency                 # 子命令组
        └── set --api-key-id <ID> --concurrency <数值>
```

#### 2.5 注册命令

在 `main.go` 的 `init()` 中注册命令：

```go
rootCmd.AddCommand(cmd.XxxCmd)
```

### Phase 3: 单元测试

在 `test/unit/cmd/` 目录创建测试文件：

**测试覆盖要求**:

1. 命令元数据测试（Use, Short, Long, GroupID）
2. 子命令结构测试
3. 必填参数验证测试
4. 参数默认值测试

**测试命名规范**:

```go
Test<命令组>Cmd           // 测试命令组
Test<子命令>Cmd           // 测试子命令
```

### Phase 4: 测试验证

1. **编译测试**

   ```bash
   go build -o agentbay .
   ```

2. **帮助信息测试**

   ```bash
   ./agentbay <command> --help
   ./agentbay <command> <subcommand> --help
   ```

3. **参数校验测试**
   - 缺少必填参数
   - 参数值验证

4. **运行新命令的单元测试**

   ```bash
   go test -v ./test/unit/cmd/ -run TestXxx -count=1
   ```

5. **🔁 全量回归测试（强制）**

   新增 / 修改 CLI 命令后，**必须**运行全量测试，确保没有任何**已有命令**的单测因接口变更、mock 缺失或公共代码改动而被破坏：

   ```bash
   # 全量测试（包括 cmd / internal / test/unit/...）
   go test ./... -count=1

   # 或更严格：跑 race 检测
   go test ./... -count=1 -race
   ```

   **通过标准**：
   - [ ] `go build -o agentbay .` 无错误且**产出的二进制保留在项目根目录**供用户直接使用
   - [ ] `go test ./... -count=1` **全部 PASS**
   - [ ] 旧命令的单测**一个都没有**因为本次改动而失败
   - [ ] 若有 mock 类，全部已同步新方法（参见 [mock-sync-guide.md](references/mock-sync-guide.md)）

   > **重要**：每次新增或修改命令后，必须执行 `go build -o agentbay .` 重新构建二进制到项目根目录。不要仅用 `go build ./...`（只验证编译不输出文件），否则用户运行 `./agentbay` 时仍是旧版本。

   ⚠️ **隔离原则**：新增命令不得修改其它命令的公共行为。如果必须改公共代码（如 `internal/agentbay/client.go`、`config`、`auth`），必须在 PR/变更档案里明确列出影响范围，并跑完所有相关命令的回归用例。

### Phase 5: 文档生成与同步

#### 5.1 更新 docs/ 命令文档（必须）

根据新命令所属的命令组，更新对应的双语文档文件：

| 命令组 | 文件路径 |
|--------|----------|
| core（version/login/logout） | `docs/en/core.md` / `docs/zh/core.md` |
| image | `docs/en/image.md` / `docs/zh/image.md` |
| apikey | `docs/en/apikey.md` / `docs/zh/apikey.md` |
| network | `docs/en/network.md` / `docs/zh/network.md` |
| skills | `docs/en/skills.md` / `docs/zh/skills.md` |
| docker | `docs/en/docker.md` / `docs/zh/docker.md` |

**更新内容**：
- 添加命令语法（`agentbay <group> <subcommand> [flags]`）
- 参数说明表格（参数名、类型、必填、说明）
- 使用示例（至少 1 个基本示例）
- 输出说明
- 注意事项（如有）

**要求**：
- 中英文文档**必须同步更新**，结构保持一致
- 命令示例保持英文（如 `agentbay image list`）
- 参考同文件中已有命令的文档风格

#### 5.2 更新 README（必须）

更新 `README.md` 和 `README.zh-CN.md` 的 Command Overview 表格：
- 新增命令：在对应命令组行添加新子命令的简短说明
- 修改命令：更新对应行的描述
- 保持中英文表格内容一致

#### 5.3 更新对客文档

创建对客功能文档（放在 `cli-analysis/` 目录）：

**文档要求**:

- 面向客户，不包含代码实现细节
- 包含完整的使用示例
- 参数说明表格
- 错误处理和 FAQ
- 认证方式和环境配置

### Phase 6: 代码提交（需用户确认）

**⚠️ 重要**: 必须询问用户是否提交，不要自动提交！

提交前展示：

```bash
git status
git diff --stat
```

询问用户："需要我帮你提交代码吗？"

用户确认后，使用规范的 commit message：

```bash
git add -A
git commit -m "feat: add <功能描述> CLI command

- 具体改动点 1
- 具体改动点 2
- 具体改动点 3"
```

## 📤 输出标准

### 代码输出

✅ **必须包含**:

- [ ] 请求/响应模型文件
- [ ] SDK 客户端方法（3 个变体 + 解析函数）
- [ ] 客户端接口定义和实现
- [ ] CLI 命令文件（使用命名参数）
- [ ] 命令注册（main.go）
- [ ] 单元测试文件

✅ **代码质量**:

- [ ] 编译无错误
- [ ] 所有测试通过
- [ ] 参数校验完整
- [ ] 错误处理友好

### 文档输出

✅ **docs/ 命令文档**（必须）:

- [ ] `docs/en/<command-group>.md` 已更新（语法、参数、示例、输出）
- [ ] `docs/zh/<command-group>.md` 已更新（与英文版结构一致）
- [ ] `README.md` Command Overview 表格已更新
- [ ] `README.zh-CN.md` Command Overview 表格已更新

✅ **对客文档**（cli-analysis/）:

- [ ] 功能概述
- [ ] 使用方法（含示例）
- [ ] 参数说明表格
- [ ] 输出示例
- [ ] 错误处理和 FAQ
- [ ] 认证方式说明
- [ ] 环境配置说明

### Git 提交

✅ **提交规范**:

- [ ] 询问用户是否提交
- [ ] 使用 Conventional Commits 格式
- [ ] 包含详细的 commit body
- [ ] 展示提交结果

## 📚 参考资料

- [references/api-format.md](references/api-format.md) - API 格式规范
- [references/cli-design.md](references/cli-design.md) - CLI 设计规范
- [references/test-requirements.md](references/test-requirements.md) - 测试要求
- [references/mock-sync-guide.md](references/mock-sync-guide.md) - Mock 同步更新规范（重要！）
- [references/document-template.md](references/document-template.md) - 文档模板

## 🛠️ 工具脚本

- [scripts/run-tests.sh](scripts/run-tests.sh) - 运行测试
- [scripts/build-and-test.sh](scripts/build-and-test.sh) - 编译并测试
- [scripts/show-diff.sh](scripts/show-diff.sh) - 展示改动

## 📝 模板文件

- [assets/request-model.go](assets/request-model.go) - 请求模型模板
- [assets/response-model.go](assets/response-model.go) - 响应模型模板
- [assets/cli-command.go](assets/cli-command.go) - CLI 命令模板
- [assets/test-file.go](assets/test-file.go) - 测试文件模板
- [assets/customer-doc.md](assets/customer-doc.md) - 对客文档模板

## ⚠️ 注意事项

1. **Product ID**: CLI 使用 `xiaoying`，不是前端的 `xiaoying-double-centre`
2. **响应格式**: 后端返回的字段类型可能与预期不同，需要实际测试确认。数字字段可能被返回为字符串（如 `"HttpStatusCode":"200"`），parser 必须用 `dual_format_responses.go` 中的 `int32FromFlexibleJSON` 容错，绝不直接 `json.Unmarshal` 打到 `*int32`。详见 [references/api-format.md 响应解析容错模板](references/api-format.md#响应解析容错模板必须使用)。
3. **参数设计**: 始终使用命名参数（`--name`），不使用位置参数
4. **命令层级**: 相关功能组织为子命令，不要创建顶级命令
5. **测试覆盖**: 必须有单元测试，且所有测试通过
6. **文档面向客户**: 对客文档不包含代码实现细节
7. **不要自动提交**: 必须用户明确要求才执行 git commit
8. **⚠️ 接口变更必须同步 Mock**: 给 `agentbay.Client` 接口添加新方法后，**必须立即更新所有 mock 类**！
   - 查找所有 mock 类：`grep -r "type mock.*Client struct" cmd/ test/`
   - 为每个 mock 类添加新方法（返回 `fmt.Errorf("not implemented")`）
   - 常见 mock 类：`mockGetMcpImageInfoClient`, `mockImageListClient`
   - 否则 CI 会报错：`*mockClient does not implement agentbay.Client (missing method Xxx)`
