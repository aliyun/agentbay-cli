# Skills Update 命令实现计划

## Context

当前 `skills push` 命令使用 `CreateMarketSkill` API 创建新技能，但缺少更新已有技能的能力。需要新增 `skills update` 命令，使用 `UpdateMarketSkill` API 更新已存在的技能。该 API 与 `CreateMarketSkill` 传参和逻辑基本一致，但多两个参数：`SkillId`（必传）和 `Icon`（非必传 string）。与 push 不同的是 update 命令**所有参数均使用命名参数（flag）**，文件路径通过 `--file` 传入且为**必传**，`--tag` 和 `--icon` 为可选。

## 命令设计

```bash
# 完整更新：上传新文件 + 更新 icon/tags
agentbay skills update --skill-id <id> --file ./my-skill --tag "tag1" --icon "https://example.com/icon.png"

# 上传文件并更新标签
agentbay skills update --skill-id <id> --file ./my-skill.zip --tag "tag1" --tag "tag2"

# 仅上传新文件，不更新 icon/tags
agentbay skills update --skill-id <id> --file ./my-skill
```

**参数列表：**

| Flag         | 类型        | 必填 | 说明                       |
| ------------ | ----------- | ---- | -------------------------- |
| `--skill-id` | string      | 是   | 要更新的技能 ID            |
| `--file`     | string      | 是   | 技能目录或 `.zip` 文件路径 |
| `--tag`      | stringArray | 否   | 技能标签（可多次指定）     |
| `--icon`     | string      | 否   | 技能图标（URL 或标识）     |

**核心流程：**

- 验证路径 → 解析 SKILL.md → 处理 tags（如有）→ 获取凭证 → 上传 → 调用 UpdateMarketSkill

## 实现步骤

### Step 1: SDK 请求模型

**新建** `internal/client/update_market_skill_request_model.go`

- `UpdateMarketSkillRequest` 结构体，字段：
  - `SkillId *string` — 必传
  - `OssBucket *string` — 有文件时必传
  - `OssFilePath *string` — 有文件时必传
  - `Tags []string` — 可选
  - `Icon *string` — 可选
- 接口 `iUpdateMarketSkillRequest` 含所有字段的 getter/setter
- `Validate()` 方法
- 参照 `create_market_skill_request_model.go` 模式

### Step 2: SDK 响应模型

**新建** `internal/client/update_market_skill_response_model.go`

- `UpdateMarketSkillResponse` 结构体，字段：`Headers`, `StatusCode *int32`, `Body *CreateMarketSkillResponseBody`（复用），`RawBody string`
- Body 类型复用 `CreateMarketSkillResponseBody`，因为 Update 响应结构与 Create 完全一致
- 接口 `iUpdateMarketSkillResponse` 含 getter/setter

### Step 3: 响应解析器

**修改** `internal/client/dual_format_responses.go`

添加：

- `xmlUpdateMarketSkillResponse` XML 结构体（根元素 `UpdateMarketSkillResponse`）
- `parseUpdateMarketSkillResponse()` 函数，逻辑与 `parseCreateMarketSkillResponse` 一致：
  - XML 分支：使用 `xmlUpdateMarketSkillResponse`
  - JSON 分支：复用 `createMarketSkillJSONWire` 和 `parseCreateMarketSkillDataField`
  - 错误包装：`ErrWithRequestID`
  - 归一 headers/statusCode：`applyMapHeadersAndStatus`

**新建** `internal/client/update_market_skill_parse_test.go`

至少覆盖：JSON Data 为字符串、JSON Data 为对象、XML 格式

### Step 4: SDK 客户端方法

**修改** `internal/client/client.go`

添加 `UpdateMarketSkillWithOptions()`，参照 `CreateMarketSkillWithOptions`（line 405）：

- body 增加 `SkillId`、`Icon` 字段
- Action 改为 `"UpdateMarketSkill"`
- 调用 `parseUpdateMarketSkillResponse`

添加 `UpdateMarketSkill()` 便捷方法。

**修改** `internal/client/client_context_func.go`

添加 `UpdateMarketSkillWithContext()`，参照 `CreateMarketSkillWithContext`（line 32）。

### Step 5: Client 接口

**修改** `internal/agentbay/client.go`

- `Client` interface 添加：`UpdateMarketSkill(ctx context.Context, request *client.UpdateMarketSkillRequest) (*client.CreateMarketSkillResponse, error)`
- `clientWrapper` 实现：调用 `sdkClient.UpdateMarketSkillWithOptions`

### Step 6: Mock 类同步

**修改** `cmd/image_list_helper_test.go` — `mockImageListClient` 添加：

```go
func (m *mockImageListClient) UpdateMarketSkill(ctx context.Context, request *client.UpdateMarketSkillRequest) (*client.CreateMarketSkillResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

**修改** `cmd/image_status_helper_test.go` — `mockGetMcpImageInfoClient` 添加同上。

### Step 7: 命令层

**修改** `cmd/skills.go`

与 `skills push` 共享大部分逻辑，但由于 update 的参数结构和流程分支（文件可选）与 push 有较大差异，**不提取共享函数**，而是独立实现 `runSkillsUpdate`。理由：

- push 的文件路径是位置参数（存量不改），update 的文件路径是 `--file` flag
- push 必须有文件，update 可以没有文件
- update 有 `--skill-id`、`--icon` 等 push 没有的参数
- 强行抽取共享函数会导致参数传递复杂化，降低可读性

新增内容：

1. `skillsUpdateCmd` 定义：

```go
var skillsUpdateCmd = &cobra.Command{
    Use:   "update",
    Short: "Update an existing skill in the cloud",
    Long:  `Update an existing skill by ID. Upload a new zip file, update tags, or set an icon.`,
    Args:  cobra.NoArgs,  // 所有参数都是 flag
    RunE:  runSkillsUpdate,
}
```

2. `init()` 注册：

```go
SkillsCmd.AddCommand(skillsUpdateCmd)
skillsUpdateCmd.Flags().String("skill-id", "", "Skill ID to update (required)")
_ = skillsUpdateCmd.MarkFlagRequired("skill-id")
skillsUpdateCmd.Flags().String("file", "", "Path to skill directory or .zip file (required)")
_ = skillsUpdateCmd.MarkFlagRequired("file")
skillsUpdateCmd.Flags().StringArray("tag", nil, `Tag name for the skill (can be specified multiple times, e.g. --tag "tag1" --tag "tag2")`)
skillsUpdateCmd.Flags().String("icon", "", "Icon for the skill (e.g. URL or identifier)")
```

3. `runSkillsUpdate` 实现，流程：

```
1. 读取 --skill-id（必填，已由 MarkFlagRequired 校验）
2. 读取 --file（必填，已由 MarkFlagRequired 校验）、--tag、--icon
3. 验证路径（目录需含 SKILL.md，或 .zip 文件）
4. 处理 tags（ListTag → CreateTag 缺失标签）
5. 获取上传凭证 GetMarketSkillCredential
6. 打包上传到 OSS
7. 构建 UpdateMarketSkillRequest（含 SkillId, OssBucket, OssFilePath, Tags, Icon）
8. 调用 UpdateMarketSkill
9. 打印 RequestId
10. 打印成功消息
```

4. 更新 `SkillsCmd.Long` 描述包含 update。

**复用的辅助函数**（不需要改动）：

- `parseSkillFrontmatter()` — 解析 SKILL.md
- `skillDirToZipFileName()` — 目录名转 zip 名
- `parseBucketAndPathForCreate()` — 解析 bucket/path
- `zipSkillDir()` — 打包目录
- `uploadFileToOSS()` — 上传文件
- `withTransientRetry()` — 重试

### Step 8: 单元测试

**修改** `test/unit/cmd/skills_cmd_test.go`

添加：

- 子命令列表增加 `"update"`
- update 不接受位置参数（`cobra.NoArgs`）
- `--skill-id` flag 存在且必填
- `--file` flag 存在且**必填**
- `--tag` flag 存在
- `--icon` flag 存在且非必填

### Step 9: 文档更新（委托 update-cli-command-docs skill）

**修改** `docs/en/skills.md` — 新增 `skills update` 章节
**修改** `docs/zh/skills.md` — 新增 `skills update` 中文章节
**修改** `README.md` — Command Overview 表格、API 表格、RAM policy
**修改** `README.zh-CN.md` — 同步中文

## 关键文件清单

| 操作 | 文件路径                                                |
| ---- | ------------------------------------------------------- |
| 新建 | `internal/client/update_market_skill_request_model.go`  |
| 新建 | `internal/client/update_market_skill_response_model.go` |
| 新建 | `internal/client/update_market_skill_parse_test.go`     |
| 修改 | `internal/client/dual_format_responses.go`              |
| 修改 | `internal/client/client.go`                             |
| 修改 | `internal/client/client_context_func.go`                |
| 修改 | `internal/agentbay/client.go`                           |
| 修改 | `cmd/skills.go`                                         |
| 修改 | `cmd/image_list_helper_test.go`                         |
| 修改 | `cmd/image_status_helper_test.go`                       |
| 修改 | `test/unit/cmd/skills_cmd_test.go`                      |
| 修改 | `docs/en/skills.md`                                     |
| 修改 | `docs/zh/skills.md`                                     |
| 修改 | `README.md`                                             |
| 修改 | `README.zh-CN.md`                                       |

## 验证步骤

1. `go build -o agentbay .` — 编译通过
2. `go test ./... -count=1` — 全量测试通过
3. `./agentbay skills update --help` — 帮助信息正确，显示所有 flag
4. `./agentbay skills --help` — 子命令列表包含 update
5. `./agentbay skills update` — 缺少 --skill-id 报错
6. `./agentbay skills update --skill-id abc` — 缺少 --file 报错
7. `./agentbay skills update --skill-id abc --file ./my-skill` — 上传并更新
8. `./agentbay skills update --skill-id abc --file ./my-skill --icon "https://..."` — 上传并更新 icon
