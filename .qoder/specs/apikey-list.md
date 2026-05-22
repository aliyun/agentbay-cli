# 实现 `agentbay apikey list` 命令

## Context

用户需要在 apikey 命令组下新增 `list` 子命令，调用 `DescribeApiKeys` OpenAPI 接口，以表格形式展示 API Key 列表。该接口是 CLI 中从未使用过的全新接口，需遵守 Code-based 成功判定 SOP。当用户传入 `--api-key`（akm-xxx）时，需先调用已有的 `DescribeMcpApiKey` 获取内部 KeyId（ak-xxx），再以 `KeyIds.1` 格式传给 DescribeApiKeys。

**关键发现**：DescribeApiKeys 接口的响应经 SDK（`BodyType "string"`）处理后，body 字符串是**直接的数据载荷**（`{"ApiKeys":[...],"requestId":"...",...}`），不含外层 `code`/`data`/`successResponse` 包装。Parser 需要双路径兼容（直接 data 载荷 + wrapped 格式）。由于 body 中无 Code/Success 字段，成功判定 SOP 的 nil 兼容逻辑尤为关键。

## 需要新建的文件（5 个）

| 文件 | 用途 |
|------|------|
| `internal/client/describe_api_keys_request_model.go` | 请求模型 |
| `internal/client/describe_api_keys_response_model.go` | 响应模型 + nil-safe getter |
| `internal/client/describe_api_keys_parse_test.go` | parser 单测（JSON string/number + XML） |
| `cmd/apikey_list.go` | CLI 命令实现 |
| `test/unit/cmd/apikey_list_cmd_test.go` | 命令元数据/flag 单测 |

## 需要修改的文件（5 个）

| 文件 | 改动 |
|------|------|
| `internal/client/dual_format_responses.go` | 新增 `parseDescribeApiKeysResponse` + JSON wire struct + XML struct |
| `internal/client/client.go` | 新增 `DescribeApiKeysWithOptions` / `DescribeApiKeys` / `DescribeApiKeysWithContext` |
| `internal/agentbay/client.go` | Client 接口新增 `DescribeApiKeys` 方法 + clientWrapper 实现 |
| `cmd/image_list_helper_test.go` | `mockImageListClient` 新增 `DescribeApiKeys` stub |
| `cmd/image_status_helper_test.go` | `mockGetMcpImageInfoClient` 新增 `DescribeApiKeys` stub |
| `README.md` | API Key Management 章节新增 `apikey list` |
| `README.zh-CN.md` | API Key 管理章节新增 `apikey list`（中文） |

---

## Step 1: SDK 请求模型

**文件**: `internal/client/describe_api_keys_request_model.go`

```go
type DescribeApiKeysRequest struct {
    MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
    NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
    KeyIds     []string  `json:"KeyIds,omitempty" xml:"KeyIds,omitempty"`
}
```

- `Validate()`: 所有参数均非必填，直接返回 nil
- Getter/Setter: `GetMaxResults()`, `SetMaxResults()`, `GetNextToken()`, `SetNextToken()`

## Step 2: SDK 响应模型

**文件**: `internal/client/describe_api_keys_response_model.go`

根据接口返回示例：

```go
type DescribeApiKeysResponse struct {
    Headers    map[string]*string
    StatusCode *int32
    Body       *DescribeApiKeysResponseBody
}

type DescribeApiKeysResponseBody struct {
    Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
    Data           *DescribeApiKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
    HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
    Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
    RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
    Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

type DescribeApiKeysResponseBodyData struct {
    ApiKeys   []*DescribeApiKeysResponseBodyDataApiKey `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty"`
    RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
    Count     *string                                  `json:"Count,omitempty" xml:"Count,omitempty"`
    NextToken *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

type DescribeApiKeysResponseBodyDataApiKey struct {
    Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
    GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
    LastUseDate  *string `json:"LastUseDate,omitempty" xml:"LastUseDate,omitempty"`
    ApiKey       *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
    Concurrency  *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
    KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
    Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
    // BoundPolicy 和 BoundResource 保留在结构体中但不在终端展示
    BoundPolicy   *DescribeApiKeysResponseBodyDataApiKeyBoundPolicy   `json:"BoundPolicy,omitempty" xml:"BoundPolicy,omitempty"`
    BoundResource *DescribeApiKeysResponseBodyDataApiKeyBoundResource `json:"BoundResource,omitempty" xml:"BoundResource,omitempty"`
}

type DescribeApiKeysResponseBodyDataApiKeyBoundPolicy struct {
    PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
    Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

type DescribeApiKeysResponseBodyDataApiKeyBoundResource struct{}
```

为每个 struct 提供 nil-safe getter 方法（`GetCode()`, `GetData()`, `GetName()`, `GetStatus()`, `GetKeyId()`, `GetNextToken()` 等）。

## Step 3: 响应 Parser

**文件**: `internal/client/dual_format_responses.go`（追加内容）

### JSON wire struct（双路径兼容：wrapped 响应 与 直接 data 载荷）

> **重要**：Alibaba Cloud SDK 的 `CallApi` 配合 `BodyType "string"` 时，实际返回的 body 字符串是**直接的数据载荷**（`{"ApiKeys":[...],"requestId":"...",...}`），不是包裹在 `{"code":"200","data":{...}}` 里的完整响应。Parser 必须双路径兼容：
> - 路径 A：若检测到 `wire.Data` 非空 → 走传统两阶段解析（wrapped 格式兼容）
> - 路径 B：若 `wire.Data` 为空 → 将 body 直接解析为 data 载荷（实际 SDK 行为）

外层字段名（wrapped 格式）实际为小写：

```go
type describeApiKeysJSONWire struct {
    Code           *string         `json:"code"`
    Data           json.RawMessage `json:"data"`
    HttpStatusCode json.RawMessage `json:"httpStatusCode"`
    Message        *string         `json:"message"`
    RequestId      *string         `json:"requestId"`
    Success        *bool           `json:"successResponse"`
}
```

Data 层级字段名（混合大小写，`requestId` 为小写）：

```go
type describeApiKeysDataJSONWire struct {
    ApiKeys   []describeApiKeysApiKeyJSONWire `json:"ApiKeys"`
    RequestId *string                         `json:"requestId"`
    Count     *string                         `json:"Count"`
    NextToken *string                         `json:"NextToken"`
}

type describeApiKeysApiKeyJSONWire struct {
    Status      *string         `json:"Status"`
    GmtCreate   *string         `json:"GmtCreate"`
    LastUseDate *string         `json:"LastUseDate"`
    ApiKey      *string         `json:"ApiKey"`
    Concurrency json.RawMessage `json:"Concurrency"`
    KeyId       *string         `json:"KeyId"`
    Name        *string         `json:"Name"`
    BoundPolicy json.RawMessage `json:"BoundPolicy"`
    BoundResource json.RawMessage `json:"BoundResource"`
}
```

### XML struct

```go
type xmlDescribeApiKeysResponse struct {
    XMLName        xml.Name `xml:"DescribeApiKeysResponse"`
    RequestId      string   `xml:"RequestId"`
    HttpStatusCode string   `xml:"HttpStatusCode"`
    Code           string   `xml:"Code"`
    Success        bool     `xml:"Success"`
    Message        string   `xml:"Message"`
    Data           struct {
        ApiKeys []struct {
            Status      string `xml:"Status"`
            GmtCreate   string `xml:"GmtCreate"`
            LastUseDate string `xml:"LastUseDate"`
            ApiKey      string `xml:"ApiKey"`
            Concurrency string `xml:"Concurrency"`
            KeyId       string `xml:"KeyId"`
            Name        string `xml:"Name"`
        } `xml:"ApiKeys"`
        RequestId string `xml:"RequestId"`
        Count     string `xml:"Count"`
        NextToken string `xml:"NextToken"`
    } `xml:"Data"`
}
```

### Parser 函数 `parseDescribeApiKeysResponse`

- **JSON 分支（双路径）**:
  1. 先 unmarshal 到 `describeApiKeysJSONWire`
  2. 若 `wire.Data` 非空 → 传统两阶段解析（wrapped 格式，提取 Code/Success/HttpStatusCode 等外层字段）
  3. 若 `wire.Data` 为空 → **将 body 直接 unmarshal 到 `describeApiKeysDataJSONWire`**（实际 SDK 行为：body 就是 data 载荷）
  4. 每条 ApiKey 的 Concurrency 用 `int32FromFlexibleJSON` 兼容字符串/数字
- XML 分支: Concurrency 从 string 用 `strconv.ParseInt` 解析
- 所有错误用 `&ErrWithRequestID{Err: ..., RequestID: extractRequestIDFromResponse(res)}` 包装
- 最后调用 `applyMapHeadersAndStatus`

> **调试经验**：最初误以为 body 是完整 wrapped 响应（`{"code":"200","data":{...}}`），导致 `wire.Data` 始终为空。通过实际打印 body 前缀发现 SDK 返回的是直接数据载荷（`{"ApiKeys":[...],"requestId":"...",...}`），确认 SDK 已剥离外层包装。

## Step 4: Parser 单测

**文件**: `internal/client/describe_api_keys_parse_test.go`

4 个测试用例：
1. `TestParseDescribeApiKeysResponse_JSONRealServerFormat` — **直接 data 载荷格式**（实际 SDK 行为）：body 为 `{"ApiKeys":[...],"requestId":"...","Count":"2","NextToken":"..."}`。Code/Success/HttpStatusCode 均为 nil，SOP 按成功处理。
2. `TestParseDescribeApiKeysResponse_JSONWrappedFormat` — **wrapped 响应格式**（兼容场景）：body 为 `{"code":"200","data":{"ApiKeys":[...]},"httpStatusCode":"200","requestId":"...","successResponse":true}`。
3. `TestParseDescribeApiKeysResponse_JSONHttpStatusCodeAsNumber` — **直接 data 载荷**，Concurrency 为数字 `5`（非字符串）
4. `TestParseDescribeApiKeysResponse_XML` — 完整 XML body

验证字段：Data.ApiKeys[0] 的各字段, Data.Count, Data.NextToken。对于 wrapped 格式额外验证 Code/Success/HttpStatusCode。

## Step 5: SDK Client 方法

**文件**: `internal/client/client.go`（追加）

```go
func (client *Client) DescribeApiKeysWithOptions(request *DescribeApiKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiKeysResponse, _err error) {
    _err = request.Validate()
    // ...
    query := map[string]interface{}{}
    if !dara.IsNil(request.MaxResults) {
        query["MaxResults"] = request.MaxResults
    }
    if request.NextToken != nil && *request.NextToken != "" {
        query["NextToken"] = request.NextToken
    }
    // KeyIds.1, KeyIds.2, ... 格式（与 ImageIds 同模式）
    for i, id := range request.KeyIds {
        key := fmt.Sprintf("KeyIds.%d", i+1)
        query[key] = id
    }
    // Action: "DescribeApiKeys", BodyType: "string"
    // ...
    _result, _err = parseDescribeApiKeysResponse(_body)
    return _result, _err
}
```

`DescribeApiKeys` 和 `DescribeApiKeysWithContext` 同已有模式。

## Step 6: Client 接口层

**文件**: `internal/agentbay/client.go`

1. Client interface 新增:
```go
DescribeApiKeys(ctx context.Context, request *client.DescribeApiKeysRequest) (*client.DescribeApiKeysResponse, error)
```

2. clientWrapper 新增实现:
```go
func (cw *clientWrapper) DescribeApiKeys(ctx context.Context, request *client.DescribeApiKeysRequest) (*client.DescribeApiKeysResponse, error) {
    sdkClient, err := cw.getClient()
    if err != nil {
        return nil, err
    }
    return sdkClient.DescribeApiKeysWithContext(ctx, request, cw.getRuntimeOptions())
}
```

## Step 7: Mock 类更新

两个 mock 类各新增一行 stub:

- `cmd/image_list_helper_test.go` → `mockImageListClient.DescribeApiKeys`
- `cmd/image_status_helper_test.go` → `mockGetMcpImageInfoClient.DescribeApiKeys`

```go
func (m *mockXxxClient) DescribeApiKeys(ctx context.Context, request *client.DescribeApiKeysRequest) (*client.DescribeApiKeysResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

## Step 8: CLI 命令实现

**文件**: `cmd/apikey_list.go`

### 命令定义

```
Use:   "list"
Short: "List API keys"
Long:  多行描述含示例
RunE:  runApikeyList
```

### Flags（在 init() 中注册）

| Flag | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `--max-results` | int32 | 10 | 每页返回的最大条数 |
| `--api-key` | string | "" | 用户可见 API Key（akm-xxx），用于筛选特定 key |
| `--next-token` | string | "" | 分页 Token |

注册到父命令: `ApiKeyCmd.AddCommand(apikeyListCmd)`

### `runApikeyList` 逻辑

1. 读取 flags: `maxResults`, `apiKey`, `nextToken`
2. 加载 config, 创建 apiClient
3. **如果 `--api-key` 非空**:
   - 打印 `[STEP 1/2] Looking up API key...`
   - 调用 `DescribeMcpApiKey(ctx, {ApiKey: &apiKey})`
   - 打印 `[INFO] DescribeMcpApiKey Request ID: xxx`
   - 用 `GetSuccess()` 判定（已有接口，沿用模板）
   - 提取 `apiKeyId = data.GetApiKeyId()`
4. **调用 DescribeApiKeys**:
   - 打印 `[STEP 2/2]` 或 `[STEP 1/1]`
   - 构建 `DescribeApiKeysRequest`:
     - `MaxResults = &maxResults`
     - `NextToken = &nextToken`（非空时）
     - `KeyIds = []string{apiKeyId}`（有 --api-key 时）
   - 调用 `apiClient.DescribeApiKeys(ctx, req)`
   - 打印 `[INFO] DescribeApiKeys Request ID: xxx`
5. **成功判定**（Code-based SOP — 全新接口）:
   ```go
   code := resp.Body.GetCode()
   successPtr := resp.Body.Success
   if (successPtr != nil && !*successPtr) || (code != "" && !isSuccessCode(code)) {
       msg := resp.Body.GetMessage()
       return fmt.Errorf("[ERROR] Failed to list API keys: Code=%s, Message=%s", code, msg)
   }
   ```

   其中 `isSuccessCode` 同时兼容 `"ok"` 和 HTTP 2xx 状态码（如 `"200"`）：
   ```go
   func isSuccessCode(code string) bool {
       if strings.EqualFold(code, "ok") { return true }
       if len(code) == 3 && code[0] == '2' { return true } // 2xx
       return false
   }
   ```

   > **注意**：由于 SDK 返回的 body 直接是 data 载荷（无外层 Code/Success），解析后 `resp.Body.Code` 和 `resp.Body.Success` 均为 nil。SOP 的兼容逻辑（nil Success 按成功处理、空 Code 按成功处理）在此场景下至关重要，否则会出现"假失败"。
6. **表格输出**:
   - 空: `[EMPTY] No API keys found.`
   - 非空: `printApiKeyTable(data.ApiKeys)`
   - 有 NextToken: `[INFO] More results available. Use --next-token <token> to fetch the next page.`

### `printApiKeyTable` 表格格式

```
NAME                 STATUS      CONCURRENCY    KEY ID                   CREATED                  LAST USED
----                 ------      ------------    ------                   -------                  ---------
lxy-cli-3            ENABLED               21    ak-df1u29s116881nt6q    2026-04-10T10:10:57      2026-05-19T10:03:05
```

- 使用已有的 `padString`, `truncateString`, `displayWidth` 辅助函数
- 不展示 BoundPolicy、BoundResource、ApiKey(akm-xxx) 字段
- Concurrency 显示整数，nil 时显示 "-"
- 日期字段截断时区偏移以更清晰

## Step 9: 命令单测

**文件**: `test/unit/cmd/apikey_list_cmd_test.go`

package `cmd_test`，参考 `apikey_cmd_test.go` 模式：

- `TestApiKeyListCmd`:
  - `list command has correct metadata` — Use="list", Short, Long
  - `list command has --max-results flag with default 10`
  - `list command has --api-key flag`
  - `list command has --next-token flag`
  - `list command is registered under apikey` — "list" in `ApiKeyCmd.Commands()`

## Step 10: README 更新

**文件**: `README.md`

在 `apikey delete` 和 `apikey concurrency set` 之间插入：

```markdown
#### `apikey list`

List API keys with optional filtering and pagination.

```bash
agentbay apikey list                                        # List up to 10 API keys
agentbay apikey list --max-results 20                       # List up to 20 API keys
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx         # Query a specific API key
agentbay apikey list --next-token <token>                   # Fetch the next page
```
```

同时更新 Overview 中的 API Key Management 描述，添加 "list"。

**文件**: `README.zh-CN.md`

在 `apikey delete` 和 `apikey concurrency set` 之间插入对应中文内容：

```markdown
#### `apikey list`

列出 API Key，支持筛选和分页。

```bash
agentbay apikey list                                        # 最多列出 10 个 API Key
agentbay apikey list --max-results 20                       # 最多列出 20 个 API Key
agentbay apikey list --api-key akm-xxxxxxxxxxxxxxxx         # 查询指定的 API Key
agentbay apikey list --next-token <token>                   # 获取下一页
```
```

---

## 实现顺序

1. 创建 `describe_api_keys_request_model.go`
2. 创建 `describe_api_keys_response_model.go`
3. 在 `dual_format_responses.go` 中添加 parser
4. 创建 `describe_api_keys_parse_test.go`，运行 `go test ./internal/client/... -run TestParseDescribeApiKeys -count=1`
5. 在 `client.go` 中添加 SDK client 方法
6. 在 `internal/agentbay/client.go` 中添加接口 + 实现
7. 更新两个 mock 类
8. 运行 `go test ./... -count=1` 确保编译通过
9. 创建 `cmd/apikey_list.go`，在 `apikey.go` 的 init() 中注册
10. 创建 `test/unit/cmd/apikey_list_cmd_test.go`
11. 更新 `README.md`
12. 运行 `go test ./... -count=1` + `go build -o agentbay .`

## 验证

```bash
# 1. Parser 单测
go test ./internal/client/... -run TestParseDescribeApiKeys -count=1 -v

# 2. 全量测试
go test ./... -count=1

# 3. 构建二进制
go build -o agentbay .

# 4. 手动验证帮助信息
./agentbay apikey list --help
./agentbay apikey --help   # 应看到 list 子命令

# 5. 有凭证时端到端测试
./agentbay apikey list
./agentbay apikey list --max-results 5
./agentbay apikey list --api-key akm-xxx
```
