# API Key Enable/Disable Feature

## Context

用户需要通过 CLI 启用/禁用 API key。当前 `apikey` 命令组只有 `create` 和 `concurrency set` 两个子命令，缺少状态管理能力。

**业务流程**：
1. 用户输入自己的 API key（如 `akm-xxx`）
2. CLI 调用 `DescribeMcpApiKey(ApiKey=akm-xxx)` 查询当前状态，获取内部 `ApiKeyId`（如 `ak-xxx`）
3. CLI 调用 `ModifyApiKeyStatus(ApiKey=ak-xxx, Status=ENABLED/DISABLED)` 修改状态

**命令设计**：新增 `apikey enable` 和 `apikey disable` 两个子命令。

---

## Implementation Plan

### Step 1: SDK Request/Response Models (2 个新 API)

#### 1.1 DescribeMcpApiKey

**新建** `internal/client/describe_mcp_api_key_request_model.go`
```
DescribeMcpApiKeyRequest { ApiKey *string }
```
- 标准 interface、String()、GoString()、Getter/Setter、Validate()

**新建** `internal/client/describe_mcp_api_key_response_model.go`
```
DescribeMcpApiKeyResponseBodyData {
    Status   *string  // ENABLED / DISABLED
    ApiKeyId *string  // ak-xxx (内部 ID)
    Name     *string
    AliUid   *string
}
DescribeMcpApiKeyResponseBody {
    Code           *string
    Data           *DescribeMcpApiKeyResponseBodyData
    HttpStatusCode *int32   // 需要容错解析
    Message        *string
    RequestId      *string
    Success        *bool
}
DescribeMcpApiKeyResponse { Headers, StatusCode, Body }
```

#### 1.2 ModifyApiKeyStatus

**新建** `internal/client/modify_api_key_status_request_model.go`
```
ModifyApiKeyStatusRequest { ApiKey *string, Status *string }
```
- Validate(): ApiKey 非空，Status 必须是 "ENABLED" 或 "DISABLED"

**新建** `internal/client/modify_api_key_status_response_model.go`
```
ModifyApiKeyStatusResponseBody {
    Code           *string
    HttpStatusCode *int32   // 需要容错解析
    Message        *string
    RequestId      *string
    Success        *bool
}
ModifyApiKeyStatusResponse { Headers, StatusCode, Body }
```

### Step 2: SDK Client Methods

**修改** `internal/client/client.go` - 添加 6 个方法（每个 API 3 个变体）：
- `DescribeMcpApiKeyWithOptions()` / `DescribeMcpApiKey()` / `DescribeMcpApiKeyWithContext()`
- `ModifyApiKeyStatusWithOptions()` / `ModifyApiKeyStatus()` / `ModifyApiKeyStatusWithContext()`

API 参数配置：
- Action: `"DescribeMcpApiKey"` / `"ModifyApiKeyStatus"`
- Version: `"2025-05-01"`, Protocol: `"HTTPS"`, Method: `"POST"`
- AuthType: `"AK"`, Style: `"RPC"`, BodyType: `"string"`

### Step 3: Response Parsers (dual_format_responses.go)

**修改** `internal/client/dual_format_responses.go` - 添加两个 parser：

#### parseDescribeMcpApiKeyResponse
- JSON 分支：用 wire struct，`HttpStatusCode` 走 `int32FromFlexibleJSON`，`Data` 直接解析为 struct
- XML 分支：定义 `xmlDescribeMcpApiKeyResponse`，手动映射字段
- 错误包装：`ErrWithRequestID`

#### parseModifyApiKeyStatusResponse
- JSON 分支：用 wire struct，`HttpStatusCode` 走 `int32FromFlexibleJSON`
- XML 分支：定义 `xmlModifyApiKeyStatusResponse`
- 错误包装：`ErrWithRequestID`
- 模式参考：`parseBatchCreateHideResourceGroupsWithMaxSessionResponse`（简单的无 Data 响应）

### Step 4: Parser Unit Tests

**新建** `internal/client/describe_mcp_api_key_parse_test.go`
- JSON HttpStatusCode 为字符串
- JSON HttpStatusCode 为数字
- XML 分支

**新建** `internal/client/modify_api_key_status_parse_test.go`
- JSON HttpStatusCode 为字符串
- JSON HttpStatusCode 为数字
- XML 分支

### Step 5: Client Interface & Wrapper

**修改** `internal/agentbay/client.go`：
- 在 `Client` interface 的 `// API Key` 注释区块下添加：
  ```go
  DescribeMcpApiKey(ctx, *client.DescribeMcpApiKeyRequest) (*client.DescribeMcpApiKeyResponse, error)
  ModifyApiKeyStatus(ctx, *client.ModifyApiKeyStatusRequest) (*client.ModifyApiKeyStatusResponse, error)
  ```
- 添加 `clientWrapper` 的两个方法实现（委托给 SDK 的 WithContext 方法）

### Step 6: Mock Classes Update

给以下两个 mock 类各添加 2 个 stub 方法（返回 `fmt.Errorf("not implemented")`）：

**修改** `cmd/image_status_helper_test.go` - `mockGetMcpImageInfoClient`：
- `DescribeMcpApiKey()`
- `ModifyApiKeyStatus()`

**修改** `cmd/image_list_helper_test.go` - `mockImageListClient`：
- `DescribeMcpApiKey()`
- `ModifyApiKeyStatus()`

### Step 7: CLI Commands

**新建** `cmd/apikey_status.go` - 包含 `enable` 和 `disable` 两个子命令：

```
agentbay apikey enable --api-key "akm-xxx"
agentbay apikey disable --api-key "akm-xxx"
```

**命令逻辑**（enable 和 disable 共享核心函数）：

```go
func runApiKeyStatusChange(cmd *cobra.Command, targetStatus string) error {
    // 1. 获取 --api-key 参数
    // 2. 加载 config，创建 apiClient
    // 3. [STEP 1/2] 调用 DescribeMcpApiKey(ApiKey=akm-xxx)
    //    - 打印 RequestId（无条件）
    //    - 获取当前 Status 和 ApiKeyId(ak-xxx)
    //    - 如果已经是目标状态，提示无需操作并退出
    // 4. [STEP 2/2] 调用 ModifyApiKeyStatus(ApiKey=ak-xxx, Status=targetStatus)
    //    - 打印 RequestId（无条件）
    //    - 检查 Success
    // 5. 打印成功信息
}
```

**RequestId 打印**（遵循新规范，不用 verbose 守卫）：
```go
// 成功时
fmt.Printf("[INFO] DescribeMcpApiKey Request ID: %s\n", requestId)
// 错误时
if reqId := extractReqIDFromErr(err); reqId != "" {
    fmt.Printf("[INFO] DescribeMcpApiKey Request ID: %s\n", reqId)
}
```

需要在 `cmd/verbose.go` 或 `cmd/apikey_status.go` 中添加 `extractReqIDFromErr` 辅助函数（从 `client.ErrWithRequestID` 中提取 RequestID）。

**注册子命令**：在 `cmd/apikey.go` 的 `init()` 中添加：
```go
ApiKeyCmd.AddCommand(apikeyEnableCmd)
ApiKeyCmd.AddCommand(apikeyDisableCmd)
```

**不需要**修改 `main.go`（`ApiKeyCmd` 已注册）。

### Step 8: Unit Tests

**修改** `test/unit/cmd/apikey_cmd_test.go` - 添加测试：

- `TestApiKeyEnableCmd`: 命令元数据 (Use, Short, Long)、required --api-key flag
- `TestApiKeyDisableCmd`: 命令元数据 (Use, Short, Long)、required --api-key flag
- 更新 `TestApiKeyCmd` 的子命令检查：确认包含 `enable` 和 `disable`

### Step 9: README & Docs Update

**修改** `README.md`：
- Features 描述更新：添加 enable/disable 能力
- Quick Start 区域添加示例命令

**修改** `cli-analysis/Agentbay cli 使用手册.md`（如果存在）：
- 添加 enable/disable 命令的语法、参数、示例、输出说明

---

## Files Summary

| 操作 | 文件路径 |
|------|----------|
| 新建 | `internal/client/describe_mcp_api_key_request_model.go` |
| 新建 | `internal/client/describe_mcp_api_key_response_model.go` |
| 新建 | `internal/client/modify_api_key_status_request_model.go` |
| 新建 | `internal/client/modify_api_key_status_response_model.go` |
| 新建 | `internal/client/describe_mcp_api_key_parse_test.go` |
| 新建 | `internal/client/modify_api_key_status_parse_test.go` |
| 新建 | `cmd/apikey_status.go` |
| 修改 | `internal/client/client.go` (添加 6 个 SDK 方法) |
| 修改 | `internal/client/dual_format_responses.go` (添加 2 个 parser) |
| 修改 | `internal/agentbay/client.go` (接口 + wrapper 各 2 个方法) |
| 修改 | `cmd/apikey.go` (注册 enable/disable 子命令) |
| 修改 | `cmd/image_status_helper_test.go` (mock +2 方法) |
| 修改 | `cmd/image_list_helper_test.go` (mock +2 方法) |
| 修改 | `test/unit/cmd/apikey_cmd_test.go` (添加测试) |
| 修改 | `README.md` (文档更新) |

---

## Verification

```bash
# 1. Parser 单测
go test ./internal/client/... -count=1

# 2. 全量单测
go test ./... -count=1

# 3. 构建二进制
go build -o agentbay .

# 4. 帮助信息检查
./agentbay apikey --help
./agentbay apikey enable --help
./agentbay apikey disable --help

# 5. 参数校验（应报错 --api-key required）
./agentbay apikey enable
./agentbay apikey disable
```
