# Skills Push 增加 Tags 支持与 RequestId 输出

## Context

`skills push` 命令目前不支持标签（Tags）功能，且各 OpenAPI 调用未输出 RequestId，不利于问题排查。本次需求：

1. 为 `skills push` 新增 `--tag` 参数（可多次指定），支持传入标签名称
2. 当有 Tags 时，先调用 ListTag 确认标签存在，不存在则调用 CreateTag 创建
3. 所有 OpenAPI 调用始终以 `[INFO]` 输出 RequestId

---

## 修改文件清单

| 文件                                                   | 改动类型 | 说明                                                                       |
| ------------------------------------------------------ | -------- | -------------------------------------------------------------------------- |
| `internal/client/list_tag_request_model.go`            | 新增     | ListTag 请求模型（无参数）                                                 |
| `internal/client/list_tag_response_model.go`           | 新增     | ListTag 响应模型                                                           |
| `internal/client/list_tag_response_body_model.go`      | 新增     | ListTag 响应 Body 模型（含 Data 数组：TagName/TagId）                      |
| `internal/client/create_tag_request_model.go`          | 新增     | CreateTag 请求模型（TagNameList 数组，批量创建）                           |
| `internal/client/create_tag_response_model.go`         | 新增     | CreateTag 响应模型                                                         |
| `internal/client/create_tag_response_body_model.go`    | 新增     | CreateTag 响应 Body 模型                                                   |
| `internal/client/client.go`                            | 修改     | 新增 ListTag / CreateTag 的 WithOptions 方法                               |
| `internal/client/client_context_func.go`               | 修改     | 新增 ListTag / CreateTag 的 WithContext 方法                               |
| `internal/client/dual_format_responses.go`             | 修改     | 新增 parseListTagResponse / parseCreateTagResponse                         |
| `internal/client/list_tag_parse_test.go`               | 新增     | ListTag parser 单测                                                        |
| `internal/client/create_tag_parse_test.go`             | 新增     | CreateTag parser 单测                                                      |
| `internal/client/create_market_skill_request_model.go` | 修改     | 新增 Tags 字段（[]string）                                                 |
| `internal/agentbay/client.go`                          | 修改     | Client 接口新增 ListTag / CreateTag 方法 + clientWrapper 实现              |
| `cmd/skills.go`                                        | 修改     | 核心逻辑：--tag flag、ListTag/CreateTag 调用、RequestId 输出、动态步骤编号 |
| `test/unit/cmd/skills_cmd_test.go`                     | 修改     | 新增 --tag flag 相关测试                                                   |
| `cmd/image_list_helper_test.go`                        | 修改     | mock 新增 ListTag / CreateTag                                              |
| `cmd/image_status_helper_test.go`                      | 修改     | mock 新增 ListTag / CreateTag                                              |
| `docs/en/skills.md`                                    | 修改     | 更新 push 命令文档（新增 --tag、更新输出示例、更新涉及接口）               |
| `docs/zh/skills.md`                                    | 修改     | 同步更新中文文档                                                           |
| `README.md`                                            | 修改     | Command Overview 表格更新                                                  |
| `README.zh-CN.md`                                      | 修改     | Command Overview 表格更新                                                  |

---

## 实现步骤

### Step 1: SDK 层 — 新增 ListTag / CreateTag API 定义

#### 1.1 ListTag 请求/响应模型

**`list_tag_request_model.go`**（无参数，空请求体）

**`list_tag_response_body_model.go`**：

```go
// ListTag 响应 Data 数组元素
type ListTagResponseBodyDataItem struct {
    TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
    TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

// ListTag 响应 Body
type ListTagResponseBody struct {
    Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
    Data           []ListTagResponseBodyDataItem `json:"Data,omitempty" xml:"Data,omitempty"`
    HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
    Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
    RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
    Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}
```

**`list_tag_response_model.go`**：标准 Response 包装

#### 1.2 CreateTag 请求/响应模型

**`create_tag_request_model.go`**（批量创建，传 TagNameList 数组）：

```go
type CreateTagRequest struct {
    TagNameList []string `json:"TagNameList,omitempty" xml:"TagNameList,omitempty"`
}
```

**`create_tag_response_body_model.go`**：

```go
// CreateTag 批量创建返回的每个 tag 项
type CreateTagResponseBodyDataItem struct {
    TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
    TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

type CreateTagResponseBody struct {
    Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
    Data           []CreateTagResponseBodyDataItem     `json:"Data,omitempty" xml:"Data,omitempty"`
    HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
    Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
    RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
    Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}
```

#### 1.3 client.go — 新增 WithOptions 方法

**`ListTagWithOptions`**（参考 `GetMarketSkillCredentialWithOptions` 的无参 GET 模式）：

- Action=`ListTag`，Method=GET，`BodyType="string"`
- query 为空 map，通过 `req.Query: openapiutil.Query(query)` 发送
- 使用 `Headers: map[string]*string{"Accept": dara.String("application/json")}`

**`CreateTagWithOptions`**（参考 `CreateMcpPolicyDataWithOptions` 的 POST+Body 模式）：

- Action=`CreateTag`，Method=POST，`ReqBodyType="formData"`，`BodyType="string"`
- **TagNameList 不使用 xxx.1 序列化格式**：在 body map 中将 TagNameList 先 `json.Marshal` 为 JSON 数组字符串，再放入 body，例如：
  ```go
  body := map[string]interface{}{}
  if len(request.TagNameList) > 0 {
      b, _ := json.Marshal(request.TagNameList)
      body["TagNameList"] = string(b)  // 值: `["哈哈","阿里云","lxy"]`
  }
  req := &openapiutil.OpenApiRequest{
      Body: openapiutil.ParseToMap(body),
      // ...
  }
  ```
- 这样发送的是 formData 格式 `TagNameList=["哈哈","阿里云","lxy"]`（单字段 JSON 数组字符串值），而非 `TagNameList.1=哈哈&TagNameList.2=阿里云`

#### 1.4 client_context_func.go — 新增 WithContext 方法

参考现有 `CreateMcpPolicyDataWithContext` 的模式（直接透传调用 WithOptions）：

```go
func (client *Client) ListTagWithContext(ctx context.Context, runtime *dara.RuntimeOptions) (*ListTagResponse, error) {
    return client.ListTagWithOptions(runtime)
}
func (client *Client) CreateTagWithContext(ctx context.Context, request *CreateTagRequest, runtime *dara.RuntimeOptions) (*CreateTagResponse, error) {
    return client.CreateTagWithOptions(request, runtime)
}
```

### Step 2: SDK 层 — 新增 dual-format parser

在 `dual_format_responses.go` 中新增 `parseListTagResponse` 和 `parseCreateTagResponse`。

遵循容错模板规则：

- 数字字段（HttpStatusCode）用 `json.RawMessage` + `int32FromFlexibleJSON`
- XML/JSON 双分支
- 错误用 `ErrWithRequestID` 包装

配套单测文件：

- `list_tag_parse_test.go`：JSON(数字为字符串) / JSON(数字为数字) / XML
- `create_tag_parse_test.go`：同上三种场景

### Step 3: SDK 层 — CreateMarketSkillRequest 新增 Tags 字段

修改 `create_market_skill_request_model.go`：

```go
type CreateMarketSkillRequest struct {
    OssBucket   *string  `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
    OssFilePath *string  `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
    Tags        []string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}
```

同步更新接口 `iCreateMarketSkillRequest` 和 getter/setter。

修改 `client.go` 和 `client_context_func.go` 中 `CreateMarketSkillWithOptions` 和 `CreateMarketSkillWithContext`，将 `CreateMarketSkill` 从 GET+query 改为 POST+Body，Tags 使用 `json.Marshal` 序列化为 JSON 数组字符串（与 `CreateTag` 的 TagNameList 传参方式一致）：

```go
body := map[string]interface{}{}
if !dara.IsNil(request.OssBucket) {
    body["OssBucket"] = request.OssBucket
}
if !dara.IsNil(request.OssFilePath) {
    body["OssFilePath"] = request.OssFilePath
}
if len(request.Tags) > 0 {
    b, _ := json.Marshal(request.Tags)
    body["Tags"] = string(b)  // 值: `["tag1","tag2"]`
}
req := &openapiutil.OpenApiRequest{
    Body: openapiutil.ParseToMap(body),
    Headers: map[string]*string{
        "Accept": dara.String("application/json"),
    },
}
params := &openapiutil.Params{
    Action:      dara.String("CreateMarketSkill"),
    Method:      dara.String("POST"),
    // ... 其他字段不变
    ReqBodyType: dara.String("formData"),
    BodyType:    dara.String("string"),
}
```

这样发送的是 formData 格式 `Tags=["tag1","tag2"]`（单字段 JSON 数组字符串值），而非 `Tags.1=tag1&Tags.2=tag2`。

同步更新 `dual_format_responses.go` 中 `parseCreateMarketSkillResponse`，确保不因新增 Tags 而破坏（当前 Data 只解析 SkillId，Tags 在请求侧而非响应侧，无需改动 parser）。

### Step 4: agentbay 接口层 — Client 接口新增方法

在 `internal/agentbay/client.go` 的 `Client` interface 新增：

```go
ListTag(ctx context.Context) (*client.ListTagResponse, error)
CreateTag(ctx context.Context, request *client.CreateTagRequest) (*client.CreateTagResponse, error)
```

在 `clientWrapper` 中实现：

```go
func (cw *clientWrapper) ListTag(ctx context.Context) (*client.ListTagResponse, error) {
    sdkClient, err := cw.getClient()
    if err != nil { return nil, err }
    return sdkClient.ListTagWithOptions(&dara.RuntimeOptions{})
}
func (cw *clientWrapper) CreateTag(ctx context.Context, request *client.CreateTagRequest) (*client.CreateTagResponse, error) {
    sdkClient, err := cw.getClient()
    if err != nil { return nil, err }
    return sdkClient.CreateTagWithContext(ctx, request, cw.getRuntimeOptions())
}
```

### Step 5: 更新 Mock 类

在以下两个 mock 类中新增 `ListTag` 和 `CreateTag` 方法（返回 "not implemented"）：

- `cmd/image_list_helper_test.go` → `mockImageListClient`
- `cmd/image_status_helper_test.go` → `mockGetMcpImageInfoClient`

### Step 6: 命令层 — cmd/skills.go 核心改造

#### 6.1 新增 --tag flag

使用 `--tag`（单数）而非 `--tags`，支持多次指定：`agentbay skills push ./my-skill --tag "lxy test" --tag "测试"`。

使用 Cobra 的 `StringArray` 类型（不是 `StringSlice`），因为 `StringSlice` 会拆分逗号而 `StringArray` 将每个 `--tag` 视为一个独立值：

```go
func init() {
    // ...existing...
    skillsPushCmd.Flags().StringArray("tag", nil, "Tag name for the skill (can be specified multiple times, e.g. --tag \"tag1\" --tag \"tag2\")")
}
```

读取方式：

```go
tagsFlag, _ := cmd.Flags().GetStringArray("tag")
```

#### 6.2 动态步骤编号与执行顺序

**重要**：当有 Tags 时，标签处理必须在获取 OSS 凭证之前执行，因为 OSS 凭证有有效期，标签处理耗时可能导致凭证过期失效。

**无 Tags**：3 步

```
Step 1/3: Getting upload credential...
Step 2/3: Uploading skill zip...
Step 3/3: Creating skill...
```

**有 Tags**：4 步（标签处理在最前面）

```
Step 1/4: Processing tags...
  - 调用 ListTag，获取全量标签列表
  - 收集不存在的 tag 名，一次性批量调用 CreateTag（如有）
  - 若所有 tag 已存在，输出 [INFO] All tags already exist.
Step 2/4: Getting upload credential...
Step 3/4: Uploading skill zip...
Step 4/4: Creating skill...
```

标签处理作为一个整体步骤输出，内部循环创建不存在的 tag 时不单独占步骤号。

#### 6.3 Tags 处理逻辑（在获取凭证之前执行）

CreateTag 支持批量创建，传 TagNameList 数组，不需要逐个循环调用。只需一次 ListTag + 一次 CreateTag（如有不存在的 tag）。**CreateTag 调用成功即代表所有传入的 tags 都已创建成功**，无需逐个检查返回结果。

```go
// 1. 先处理 tags（在获取凭证之前，避免凭证过期）
tagsFlag, _ := cmd.Flags().GetStringArray("tag")
var tags []string
for _, t := range tagsFlag {
    trimmed := strings.TrimSpace(t)
    if trimmed != "" {
        tags = append(tags, trimmed)
    }
}

stepIdx := 1
totalSteps := 3
if len(tags) > 0 {
    totalSteps = 4
}

if len(tags) > 0 {
    fmt.Printf("[STEP %d/%d] Processing tags...\n", stepIdx, totalSteps)
    stepIdx++

    // 1a. 调用 ListTag（获取全量标签列表）
    listResp, err := apiClient.ListTag(ctx)
    if err != nil { return err }
    if listResp.Body != nil && listResp.Body.RequestId != nil && *listResp.Body.RequestId != "" {
        fmt.Printf("[INFO] ListTag RequestId: %s\n", *listResp.Body.RequestId)
    }

    // 1b. 建立 tagName set（仅用于判断是否存在）
    existingTagNames := map[string]bool{}
    if listResp.Body != nil && listResp.Body.Data != nil {
        for _, item := range listResp.Body.Data {
            if item.TagName != nil {
                existingTagNames[*item.TagName] = true
            }
        }
    }

    // 1c. 收集所有不存在的 tag 名，一次性批量 CreateTag
    var missingTags []string
    for _, tagName := range tags {
        if !existingTagNames[tagName] {
            missingTags = append(missingTags, tagName)
        }
    }

    if len(missingTags) > 0 {
        fmt.Printf("[INFO] Tags not found: %s, creating...\n", strings.Join(missingTags, ", "))
        createReq := &client.CreateTagRequest{TagNameList: missingTags}
        createTagResp, err := apiClient.CreateTag(ctx, createReq)
        if err != nil { return err }
        if createTagResp.Body != nil && createTagResp.Body.RequestId != nil && *createTagResp.Body.RequestId != "" {
            fmt.Printf("[INFO] CreateTag RequestId: %s\n", *createTagResp.Body.RequestId)
        }
        fmt.Printf("[INFO] Tags created successfully.\n")
    } else {
        fmt.Printf("[INFO] All tags already exist.\n")
    }
}

// 2. 获取凭证
fmt.Printf("[STEP %d/%d] Getting upload credential...\n", stepIdx, totalSteps)
stepIdx++
// ... GetMarketSkillCredential logic ...

// 3. 上传
fmt.Printf("[STEP %d/%d] Uploading skill zip...\n", stepIdx, totalSteps)
stepIdx++
// ... upload logic ...

// 4. 创建 Skill
fmt.Printf("[STEP %d/%d] Creating skill...\n", stepIdx, totalSteps)
// ... CreateMarketSkill logic, 传入 tags ...
```

#### 6.4 RequestId 始终 [INFO] 输出

所有 OpenAPI 调用后，始终输出 RequestId（不依赖 verbose 模式）：

```go
// GetMarketSkillCredential
if credResp.Body != nil && credResp.Body.RequestId != nil && *credResp.Body.RequestId != "" {
    fmt.Printf("[INFO] GetMarketSkillCredential RequestId: %s\n", *credResp.Body.RequestId)
}

// CreateMarketSkill
if createResp.Body != nil && createResp.Body.RequestId != nil && *createResp.Body.RequestId != "" {
    fmt.Printf("[INFO] CreateMarketSkill RequestId: %s\n", *createResp.Body.RequestId)
}

// ListTag
if listResp.Body != nil && listResp.Body.RequestId != nil && *listResp.Body.RequestId != "" {
    fmt.Printf("[INFO] ListTag RequestId: %s\n", *listResp.Body.RequestId)
}

// CreateTag
if createTagResp.Body != nil && createTagResp.Body.RequestId != nil && *createTagResp.Body.RequestId != "" {
    fmt.Printf("[INFO] CreateTag RequestId: %s\n", *createTagResp.Body.RequestId)
}
```

#### 6.5 CreateMarketSkill 传 Tags

```go
createReq := &client.CreateMarketSkillRequest{
    OssBucket:   &createBucket,
    OssFilePath: &createOssPath,
}
if len(tags) > 0 {
    createReq.Tags = tags
}
```

### Step 7: 更新单元测试

`test/unit/cmd/skills_cmd_test.go` 新增：

- 验证 `--tag` flag 存在且类型为 StringArray
- 验证 `--tag` 可以多次指定（`--tag "tag1" --tag "tag2"`）
- 验证无 `--tag` 时不影响原有行为

### Step 8: 更新文档

#### docs/en/skills.md — push 命令新增

**参数表新增**：
| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--tag` | string | No (repeatable) | Tag name for the skill. Can be specified multiple times, e.g. `--tag "tag1" --tag "tag2"` |

**输出示例更新**（有 tags 时）：

```
[STEP 1/4] Processing tags...
[INFO] ListTag RequestId: xxx
[INFO] Tags not found: new-tag, creating...
[INFO] CreateTag RequestId: xxx
[INFO] Tags created successfully.
[STEP 2/4] Getting upload credential...
[INFO] GetMarketSkillCredential RequestId: xxx
[STEP 3/4] Uploading skill zip...
[STEP 4/4] Creating skill...
[INFO] CreateMarketSkill RequestId: xxx

[SUCCESS] Skill created successfully!
[RESULT] Skill ID: 35U2Ver2
```

所有 tags 都已存在时的输出：

```
[STEP 1/4] Processing tags...
[INFO] ListTag RequestId: xxx
[INFO] All tags already exist.
[STEP 2/4] Getting upload credential...
...
```

**用法示例**：

```bash
# No tags
agentbay skills push ./my-skill

# With tags
agentbay skills push ./my-skill --tag "lxy test" --tag "测试"
```

**涉及接口新增**：ListTag / CreateTag

#### docs/zh/skills.md — 同步更新中文版

#### README.md / README.zh-CN.md — Command Overview 表格

Skills 行的子命令列可能需要更新（若 push 的参数变化影响到表格），否则无需改动。

---

## 成功判定规则

ListTag 和 CreateTag 是全新接口，按开发规则 SOP，采用 `Code` 判定：

```go
code := resp.Body.GetCode()
successPtr := resp.Body.Success
if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
    msg := resp.Body.GetMessage()
    return fmt.Errorf("[ERROR] Failed to xxx: Code=%s, Message=%s", code, msg)
}
```

GetMarketSkillCredential 和 CreateMarketSkill 是存量接口，沿用原有判定逻辑不变。

---

## 验证清单

1. `go test ./internal/client/... -count=1` — parser 单测通过
2. `go test ./... -count=1` — 全部测试通过
3. `go build -o agentbay .` — 构建成功
4. 手动测试场景：
   - `agentbay skills push ./my-skill`（无 --tag，行为不变，但 RequestId 始终输出）
   - `agentbay skills push ./my-skill --tag "lxy test" --tag "测试"`（tags 已存在）
   - `agentbay skills push ./my-skill --tag "new tag" --tag "tag1"`（部分 tag 不存在，自动创建）
   - `agentbay skills push ./my-skill --tag "  tag1  " --tag "  tag2  "`（验证前后空格去除）
