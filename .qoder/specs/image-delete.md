# Plan: `agentbay image delete` Command

## Context

AgentBay CLI 目前支持镜像的创建、激活、去激活和状态查询，但缺少**物理删除**能力。用户在镜像不再需要时，无法通过 CLI 清理已去激活的镜像资源。本次新增 `agentbay image delete <image-id> [--yes/-y]` 命令，调用后端 `DeleteMcpImage` 同步接口永久删除 User 类型镜像。

---

## Implementation Steps

### Phase 1: SDK Model Files (3 new files)

**1.1** `internal/client/delete_mcp_image_request_model.go` (NEW)
- `DeleteMcpImageRequest` struct，单字段 `ImageId *string`
- Setter/Getter/String/GoString/Validate 方法
- 参考模板: `delete_resource_group_request_model.go`

**1.2** `internal/client/delete_mcp_image_response_body_model.go` (NEW)
- `DeleteMcpImageResponseBody` struct: `Code`, `HttpStatusCode`, `Message`, `RequestId`, `Success`
- 无 Data 字段（同步删除接口无返回数据）
- 参考模板: `delete_resource_group_response_body_model.go`

**1.3** `internal/client/delete_mcp_image_response_model.go` (NEW)
- `DeleteMcpImageResponse` struct: `Headers`, `StatusCode`, `Body`
- 参考模板: `delete_resource_group_response_model.go`

### Phase 2: SDK Client Methods

**2.1** `internal/client/dual_format_responses.go` (MODIFY)
- 新增 `xmlDeleteMcpImageResponse` struct (XMLName: `"DeleteMcpImageResponse"`)
- 新增 `parseDeleteMcpImageResponse` function（XML/JSON 双格式解析）
- 放在 `parseDeleteResourceGroupResponse` 之后

**2.2** `internal/client/client.go` (MODIFY)
- 新增 `DeleteMcpImage(request) (*DeleteMcpImageResponse, error)` — 便捷版
- 新增 `DeleteMcpImageWithOptions(request, runtime)` — 核心实现
  - Action: `"DeleteMcpImage"`, Version: `"2025-05-01"`, Method: POST, Style: RPC
  - Query param: `ImageId`
  - 调用 `parseDeleteMcpImageResponse`

**2.3** `internal/client/client_context_func.go` (MODIFY)
- 新增 `DeleteMcpImageWithContext` — 委托给 WithOptions

### Phase 3: Interface Layer

**3.1** `internal/agentbay/client.go` (MODIFY)
- Client interface 新增:
  ```go
  DeleteMcpImage(ctx context.Context, request *client.DeleteMcpImageRequest) (*client.DeleteMcpImageResponse, error)
  ```
- clientWrapper 新增对应方法实现

### Phase 4: Status Helper

**4.1** `cmd/image_status_helper.go` (MODIFY)
- 新增常量: `StatusResourceMaintaining ImageResourceStatus = "RESOURCE_MAINTAINING"`
- `TranslateImageResourceStatus` 新增 case: `StatusResourceMaintaining` → `"Maintaining"`
- 新增函数 `IsDeletable(status string) bool`:
  - 返回 false 的状态（6 种禁止删除）:
    - `IMAGE_CREATING`, `RESOURCE_DEPLOYING`, `RESOURCE_DELETING`
    - `RESOURCE_PUBLISHED`, `RESOURCE_FAILED`, `RESOURCE_MAINTAINING`
  - 其他状态返回 true（如 `IMAGE_AVAILABLE`, `IMAGE_CREATE_FAILED`, `RESOURCE_CEASED`）

### Phase 5: Confirmation Prompt

**5.1** `cmd/confirm.go` (NEW)
- `ConfirmPrompt(prompt string, autoYes bool) (bool, error)`
  - `autoYes == true` → 直接返回 `(true, nil)`
  - 使用 `term.IsTerminal(int(os.Stdin.Fd()))` 检测 TTY
  - 非 TTY 且未传 --yes → 返回 error: `"non-interactive environment detected: use --yes to confirm"`
  - TTY: 打印 prompt，读取一行输入
  - 仅 `y`/`Y`/`yes`/`YES` 视为确认，其他（含空回车）视为拒绝返回 `(false, nil)`
- 提取 `isConfirmInput(s string) bool` 为独立函数便于测试
- 依赖: `golang.org/x/term`（已在 go.sum 传递依赖中，需 `go get` 加入 go.mod）

### Phase 6: Command Implementation

**6.1** `cmd/image.go` (MODIFY)

新增 `imageDeleteCmd` 变量:
```go
var imageDeleteCmd = &cobra.Command{
    Use:   "delete <image-id>",
    Short: "Delete a User image permanently",
    Long:  `...`, // 含 --yes 用法示例
    Args:  cobra.ExactArgs(1),
    RunE:  runImageDelete,
}
```

在 `init()` 中:
- 添加 `imageDeleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")`
- `ImageCmd.AddCommand(imageDeleteCmd)`

`runImageDelete` 函数流程:
1. 获取 `imageId` (args[0]) 和 `--yes` flag
2. 加载 config，鉴权校验
3. 创建 API client
4. 调用 `GetImageInfo` 获取镜像信息，打印 RequestId
5. 拒绝 System 镜像（打印提示并退出）
6. 调用 `IsDeletable(status)`:
   - false → 打印当前状态 + 不可删除原因 + 建议（如已激活建议先 deactivate）
7. 调用 `ConfirmPrompt`:
   - error → 打印 "use --yes in non-interactive mode"，exit 1
   - false → 打印 "Operation cancelled"，exit 0
8. 调用 `apiClient.DeleteMcpImage(ctx, request)`
9. 校验响应（Body 非空, Success == true）
10. 打印 DeleteMcpImage RequestId
11. 打印成功消息

### Phase 7: Mock Updates

**7.1** `cmd/image_status_helper_test.go` (MODIFY)
- `mockGetMcpImageInfoClient` 新增:
  ```go
  func (m *mockGetMcpImageInfoClient) DeleteMcpImage(ctx context.Context, request *client.DeleteMcpImageRequest) (*client.DeleteMcpImageResponse, error) {
      return nil, fmt.Errorf("not implemented")
  }
  ```

**7.2** `cmd/image_list_helper_test.go` (MODIFY)
- `mockImageListClient` 新增同样的 stub 方法

### Phase 8: Unit Tests

**8.1** `test/unit/cmd/image_delete_test.go` (NEW)
- `TestImageDeleteCommand`: 命令元数据验证 (Use, Short, Long)
- `TestImageDeleteArgs`: `cobra.ExactArgs(1)` 验证
- `TestImageDeleteYesFlag`: --yes/-y flag 存在且默认 false
- `TestIsDeletable`: 表驱动测试覆盖所有状态

**8.2** `test/unit/cmd/confirm_test.go` (NEW)
- `TestConfirmPromptAutoYes`: autoYes=true 直接返回 true
- `TestConfirmPromptNonTTY`: 测试环境 stdin 非 TTY，返回 error
- `TestIsConfirmInput`: 提取输入白名单判断为独立函数 `isConfirmInput(s string) bool` 并测试

**8.3** `test/unit/internal/client/delete_mcp_image_test.go` (NEW)
- Request model SetImageId/GetImageId 测试
- Response body model setter/getter 测试
- Validate 通过/失败测试

### Phase 9: Documentation

**9.1** `README.md` (MODIFY)
- Features 中 "Image Management" 描述补充 "delete"
- Quick Start 在 deactivate 之后添加:
  ```bash
  # 7. Delete an image permanently (irreversible)
  agentbay image delete imgc-xxxxx...xxx
  agentbay image delete imgc-xxxxx...xxx --yes  # Skip confirmation
  ```

**9.2** `docs/USER_GUIDE.md` (MODIFY)
- 在 deactivate 章节后新增 image delete 章节（语法、参数、示例、状态限制、注意事项）

---

## Files Modified/Created Summary

| File | Action | Description |
|------|--------|-------------|
| `internal/client/delete_mcp_image_request_model.go` | NEW | Request model |
| `internal/client/delete_mcp_image_response_body_model.go` | NEW | Response body model |
| `internal/client/delete_mcp_image_response_model.go` | NEW | Response model |
| `internal/client/dual_format_responses.go` | MODIFY | XML/JSON parser |
| `internal/client/client.go` | MODIFY | SDK methods |
| `internal/client/client_context_func.go` | MODIFY | Context wrapper |
| `internal/agentbay/client.go` | MODIFY | Interface + wrapper |
| `cmd/image_status_helper.go` | MODIFY | 新状态常量 + IsDeletable |
| `cmd/confirm.go` | NEW | 确认提示工具 |
| `cmd/image.go` | MODIFY | 命令注册 + runImageDelete |
| `cmd/image_status_helper_test.go` | MODIFY | Mock stub |
| `cmd/image_list_helper_test.go` | MODIFY | Mock stub |
| `test/unit/cmd/image_delete_test.go` | NEW | 命令测试 |
| `test/unit/cmd/confirm_test.go` | NEW | 确认工具测试 |
| `test/unit/internal/client/delete_mcp_image_test.go` | NEW | SDK model 测试 |
| `README.md` | MODIFY | 功能文档 |
| `docs/USER_GUIDE.md` | MODIFY | 用户手册 |

---

## Verification

```bash
# 1. 编译检查
go build ./...

# 2. 全量测试
go test ./... -count=1

# 3. 手动验证（需真实环境）
agentbay image delete imgc-xxx         # 交互式确认
agentbay image delete imgc-xxx --yes   # 跳过确认
echo "" | agentbay image delete imgc-xxx  # 非 TTY 报错

# 4. 状态校验边界
# 尝试删除已激活镜像 → 应提示不可删除
# 尝试删除 System 镜像 → 应拒绝
```
