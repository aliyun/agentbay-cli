# Spec: agentbay apikey delete 命令

## Context

用户需要通过 CLI 删除 API Key。目前 CLI 支持 create / enable / disable，但缺少 delete 功能。

**业务规则**：
- 只允许删除处于 `DISABLED` 状态的 API Key（防止意外删除使用中的密钥）
- 若 API Key 为 `ENABLED`，询问用户是否先禁用（需用户确认）
- 若状态为其他值，报错退出
- 最终执行删除前需用户二次确认
- 接口 `DeleteApiKey` 是同步的，无报错即视为成功

---

## 实现方案（分层，自下而上）

### Layer 1：SDK 模型层（internal/client/）

**新增 3 个文件**：

#### `delete_api_key_request_model.go`
```go
// 请求参数：KeyIdListJson（JSON 数组字符串，如 `["ak-xxx"]`）
type DeleteApiKeyRequest struct {
    KeyIdListJson *string `json:"KeyIdListJson,omitempty" xml:"KeyIdListJson,omitempty"`
}
// 包含 Validate() / GetKeyIdListJson() / SetKeyIdListJson() 方法
```

#### `delete_api_key_response_model.go`
```go
type DeleteApiKeyResponse struct {
    Headers    map[string]*string         `json:"headers,omitempty"`
    StatusCode *int32                     `json:"statusCode,omitempty"`
    Body       *DeleteApiKeyResponseBody  `json:"body,omitempty"`
}
```

#### `delete_api_key_response_body_model.go`
```go
type DeleteApiKeyResponseBody struct {
    Code           *string `json:"Code,omitempty"`
    HttpStatusCode *int32  `json:"HttpStatusCode,omitempty"`   // ⚠️ 用 int32，但 parser 用 RawMessage 容错
    Message        *string `json:"Message,omitempty"`
    RequestId      *string `json:"RequestId,omitempty"`
    Success        *bool   `json:"Success,omitempty"`
}
// 包含 GetXxx() 系列方法
```

---

### Layer 2：SDK 客户端方法（internal/client/client.go）

在 `client.go` 中新增（紧随其他 delete 方法之后）：

```go
func (client *Client) DeleteApiKeyWithOptions(request *DeleteApiKeyRequest, runtime *dara.RuntimeOptions) (*DeleteApiKeyResponse, error)
func (client *Client) DeleteApiKey(request *DeleteApiKeyRequest) (*DeleteApiKeyResponse, error)
func (client *Client) DeleteApiKeyWithContext(ctx context.Context, request *DeleteApiKeyRequest, runtime *dara.RuntimeOptions) (*DeleteApiKeyResponse, error)
```

**API 参数**（body 方式传参，与 `{"KeyIdListJson":["ak-xxx"]}` 对应）：
```go
params := &openapiutil.Params{
    Action:      dara.String("DeleteApiKey"),
    Version:     dara.String("2025-05-01"),
    Protocol:    dara.String("HTTPS"),
    Pathname:    dara.String("/"),
    Method:      dara.String("POST"),
    AuthType:    dara.String("AK"),
    Style:       dara.String("RPC"),
    ReqBodyType: dara.String("formData"),
    BodyType:    dara.String("string"),
}
// body["KeyIdListJson"] = request.KeyIdListJson
```

---

### Layer 3：响应解析（internal/client/dual_format_responses.go）

**新增 `parseDeleteApiKeyResponse()`**，遵循 dual-format 模板：
- XML 分支：xml.Unmarshal → 映射字段
- JSON 分支：wire struct（`HttpStatusCode json.RawMessage`）+ `int32FromFlexibleJSON` 容错
- 解析失败：`&ErrWithRequestID{Err: ..., RequestID: extractRequestIDFromResponse(res)}`
- 调用 `applyMapHeadersAndStatus` 归一 headers/statusCode

**配套测试文件 `delete_api_key_parse_test.go`**，覆盖 3 个场景：
1. JSON HttpStatusCode 为字符串（`"200"`）
2. JSON HttpStatusCode 为数字（`200`）
3. XML 格式

---

### Layer 4：接口层（internal/agentbay/client.go）

**在 `Client` interface 中添加**：
```go
DeleteApiKey(ctx context.Context, request *client.DeleteApiKeyRequest) (*client.DeleteApiKeyResponse, error)
```

**在 `clientWrapper` 中实现**：
```go
func (cw *clientWrapper) DeleteApiKey(ctx context.Context, request *client.DeleteApiKeyRequest) (*client.DeleteApiKeyResponse, error) {
    sdkClient, err := cw.getClient()
    if err != nil { return nil, err }
    return sdkClient.DeleteApiKeyWithContext(ctx, request, cw.getRuntimeOptions())
}
```

**⚠️ 立即同步所有 mock 类**（2 个文件各加一个 stub 方法）：
- `cmd/image_status_helper_test.go` → `mockGetMcpImageInfoClient`
- `cmd/image_list_helper_test.go` → `mockImageListClient`

---

### Layer 5：CLI 命令（cmd/apikey_delete.go，新文件）

**命令定义**：
```
agentbay apikey delete <api-key> [--yes|-y]
```
- 位置参数（与 enable/disable 保持一致风格）
- `cobra.ExactArgs(1)`
- `--yes` / `-y`：跳过所有二次确认（适用于脚本/CI 场景）

**Flag 注册**（参考 `image delete` 的实现）：
```go
apikeyDeleteCmd.Flags().BoolP("yes", "y", false, "Skip all confirmation prompts (for non-interactive use)")
```

**确认逻辑**：复用 `cmd/confirm.go` 中已有的 `ConfirmPrompt(prompt, autoYes bool)` 函数：
- `autoYes=true`（用户传了 `--yes`）→ 直接返回 true，跳过提示
- 非 TTY 且 `autoYes=false` → 返回错误，提示用户加 `--yes`
- 交互式终端 → 打印提示，读取用户输入（仅接受 y/Y/yes/YES）

**完整执行流程**（`runApiKeyDelete`）：

```
autoYes := cmd.Flags().GetBool("yes")

Step 1/3: 查询 API Key 信息
  → DescribeMcpApiKey(ApiKey=用户输入)
  → 打印 [INFO] DescribeMcpApiKey Request ID: xxx
  → 解析 Status、ApiKeyId、Name
  → 若 Status != ENABLED && Status != DISABLED：
      return error: API key status is '<status>', cannot delete

  → 若 Status == ENABLED：
      提示：This API key is currently ENABLED. It must be disabled before deletion.
      confirmed, err := ConfirmPrompt("Disable it now? [y/N]: ", autoYes)
      若未确认 → 打印 [INFO] Operation cancelled. return nil

Step 2/3（仅当 ENABLED 时执行）: 禁用 API Key
  → ModifyApiKeyStatus(ApiKey=ApiKeyId, Status=DISABLED)
  → 打印 [INFO] ModifyApiKeyStatus Request ID: xxx
  → 失败则 return error

Step 3/3: 删除确认 + 执行删除
  → 打印 API Key 信息（ApiKeyId、Name、当前 Status）
  → confirmed, err := ConfirmPrompt("Are you sure you want to delete this API key? [y/N]: ", autoYes)
  → 若未确认 → 打印 [INFO] Operation cancelled. return nil
  → DeleteApiKey(KeyIdListJson=["<ApiKeyId>"])
  → 打印 [INFO] DeleteApiKey Request ID: xxx
  → 成功打印:
      [SUCCESS] API key has been deleted.
        ApiKeyId: ak-xxx
        Name:     my-key
```

**在 `init()` 中注册**：
```go
ApiKeyCmd.AddCommand(apikeyDeleteCmd)
```

**使用示例**：
```bash
# 交互式删除（默认，带确认提示）
agentbay apikey delete akm-xxxxxxxxxxxxxxxx

# 非交互式删除（脚本/CI 场景，跳过所有确认）
agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes
agentbay apikey delete akm-xxxxxxxxxxxxxxxx -y
```

---

### Layer 6：单元测试（test/unit/cmd/apikey_cmd_test.go）

在现有测试文件中**新增测试**：
- `TestApikeyDeleteCmd`：验证 `delete` 子命令元数据（Use, Short, Args）
- 验证 `ApiKeyCmd` 的子命令列表包含 `delete`
- 验证 `--yes` / `-y` flag 存在且默认值为 `false`、shorthand 为 `y`

```go
func TestApikeyDeleteCmd(t *testing.T) {
    // 1. 子命令存在
    deleteCmd := findSubCommand(cmd.ApiKeyCmd, "delete")
    require.NotNil(t, deleteCmd)

    // 2. 元数据
    assert.Equal(t, "delete <api-key>", deleteCmd.Use)

    // 3. --yes flag
    yesFlag := deleteCmd.Flags().Lookup("yes")
    require.NotNil(t, yesFlag)
    assert.Equal(t, "false", yesFlag.DefValue)
    assert.Equal(t, "y", yesFlag.Shorthand)
}
```

---

### Layer 7：文档更新

1. **README.md**：在 apikey 命令章节补充 `delete` 子命令说明
2. **cli-analysis/Agentbay cli 使用手册.md**：补充 `agentbay apikey delete` 的语法、参数、示例、注意事项

---

## 关键文件清单

| 操作 | 文件路径 |
|------|---------|
| 新增 | `internal/client/delete_api_key_request_model.go` |
| 新增 | `internal/client/delete_api_key_response_model.go` |
| 新增 | `internal/client/delete_api_key_response_body_model.go` |
| 新增 | `internal/client/delete_api_key_parse_test.go` |
| 修改 | `internal/client/client.go`（添加 3 个方法） |
| 修改 | `internal/client/dual_format_responses.go`（添加 parser） |
| 修改 | `internal/agentbay/client.go`（接口 + wrapper） |
| 修改 | `cmd/image_status_helper_test.go`（mock stub） |
| 修改 | `cmd/image_list_helper_test.go`（mock stub） |
| 新增 | `cmd/apikey_delete.go` |
| 修改 | `test/unit/cmd/apikey_cmd_test.go`（新增测试） |
| 修改 | `README.md` |
| 修改 | `cli-analysis/Agentbay cli 使用手册.md` |

---

## 验证步骤

```bash
# 1. parser 单测（Layer 3）
go test ./internal/client/... -count=1 -run TestDeleteApiKey

# 2. 全量测试
go test ./... -count=1

# 3. 构建二进制
go build -o agentbay .

# 4. 帮助信息验证
./agentbay apikey delete --help

# 5. 手动测试流程（需真实 API Key）

# 场景A：删除 DISABLED 状态的 key（交互式，确认后删除）
./agentbay apikey delete akm-xxxxxxxxxxxxxxxx

# 场景B：删除 DISABLED 状态的 key（非交互式，--yes 跳过确认）
./agentbay apikey delete akm-xxxxxxxxxxxxxxxx --yes

# 场景C：删除 ENABLED 状态的 key（先询问禁用，再询问删除）
./agentbay apikey delete akm-xxxxxxxxxxxxxxxx

# 场景D：删除 ENABLED 状态的 key（--yes 跳过所有确认，自动禁用并删除）
./agentbay apikey delete akm-xxxxxxxxxxxxxxxx -y

# 场景E：取消操作（在确认提示处输入 n，打印 Operation cancelled）
./agentbay apikey delete akm-xxxxxxxxxxxxxxxx

# 场景F：状态异常的 key（非 ENABLED/DISABLED），应报错退出
# 场景G：非 TTY 环境未指定 --yes，应报错提示用户加 --yes
echo "" | ./agentbay apikey delete akm-xxxxxxxxxxxxxxxx
```
