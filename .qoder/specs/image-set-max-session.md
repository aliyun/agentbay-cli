# Spec: `agentbay image set-max-session` 命令

## Context

镜像组需要新增设置最大并发会话数的能力。用户激活镜像后，可以通过此命令配置镜像的最大并发会话数量。命令流程：校验镜像状态 -> 调用设置接口 -> 轮询等待生效。

### 业务约束说明

`agentbay image set-max-session` 实际只支持 **自定义镜像(User) + 已启用(RESOURCE_PUBLISHED) + 高级网络** 的场景。但 CLI 端仅校验前两个条件（User 镜像 + 已激活状态），**不检测是否为高级网络**。如果用户对非高级网络的镜像执行此命令，由服务端接口 `BatchCreateHideResourceGroupsWithMaxSession` 直接返回错误，CLI 将该错误透传给用户。

## 实施方案

### 文件变更总览

```
新建文件:
  internal/client/batch_create_hide_resource_groups_with_max_session_request_model.go
  internal/client/batch_create_hide_resource_groups_with_max_session_response_model.go
  cmd/image_set_max_session.go
  test/unit/cmd/image_set_max_session_test.go

修改文件:
  internal/client/get_mcp_image_info_response_body_model.go  (添加 ResourceGroupReady 字段)
  internal/client/client.go                                   (添加 SDK 方法)
  internal/client/client_context_func.go                      (添加 WithContext 方法)
  internal/agentbay/client.go                                 (接口 + wrapper 实现)
  cmd/image_status_helper.go                                  (轮询逻辑 + ImageInfo 扩展)
  cmd/image.go                                                (注册子命令)
  cmd/image_status_helper_test.go                             (mock 更新)
  cmd/image_list_helper_test.go                               (mock 更新)
```

### Step 1: SDK Request Model

**文件**: `internal/client/batch_create_hide_resource_groups_with_max_session_request_model.go`

```go
type BatchCreateHideResourceGroupsWithMaxSessionRequest struct {
    ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
    MaxSessionNum *int32  `json:"MaxSessionNum,omitempty" xml:"MaxSessionNum,omitempty"`
}
```

- 包含 `Validate()` 方法（两字段必填，MaxSessionNum >= 1）
- 包含 Get/Set 方法（遵循项目模式）

### Step 2: SDK Response Model

**文件**: `internal/client/batch_create_hide_resource_groups_with_max_session_response_model.go`

```go
type BatchCreateHideResourceGroupsWithMaxSessionResponseBody struct {
    Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
    HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
    Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
    RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
    Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

type BatchCreateHideResourceGroupsWithMaxSessionResponse struct {
    Headers    map[string]*string
    StatusCode *int32
    Body       *BatchCreateHideResourceGroupsWithMaxSessionResponseBody
}
```

### Step 3: GetMcpImageInfo 响应新增字段

**文件**: `internal/client/get_mcp_image_info_response_body_model.go`

在 `GetMcpImageInfoResponseBodyData` 结构体中添加:
```go
ResourceGroupReady *bool `json:"ResourceGroupReady,omitempty" xml:"ResourceGroupReady,omitempty"`
```

配套 Get/Set 方法。

### Step 4: SDK Client 方法

**文件**: `internal/client/client.go`

添加 3 个方法变体:
- `BatchCreateHideResourceGroupsWithMaxSession(request)` - 便捷方法
- `BatchCreateHideResourceGroupsWithMaxSessionWithOptions(request, runtime)` - 核心实现
- `parseBatchCreateHideResourceGroupsWithMaxSessionResponse(res)` - 响应解析

API 参数:
- Action: `"BatchCreateHideResourceGroupsWithMaxSession"`
- Version: `"2025-05-01"`
- Method: `"POST"`
- ReqBodyType: `"formData"`
- BodyType: `"string"`

**文件**: `internal/client/client_context_func.go`

添加 `BatchCreateHideResourceGroupsWithMaxSessionWithContext(ctx, request, runtime)` 方法。

### Step 5: Client Interface + Wrapper

**文件**: `internal/agentbay/client.go`

接口新增:
```go
BatchCreateHideResourceGroupsWithMaxSession(ctx context.Context, request *client.BatchCreateHideResourceGroupsWithMaxSessionRequest) (*client.BatchCreateHideResourceGroupsWithMaxSessionResponse, error)
```

clientWrapper 实现委托到 SDK 层。

### Step 6: 轮询逻辑

**文件**: `cmd/image_status_helper.go`

1. `ImageInfo` 结构体新增 `ResourceGroupReady bool` 字段
2. `GetImageInfo()` 中从 `resp.Body.Data.ResourceGroupReady` 读取并填充
3. 新增 `DefaultSetMaxSessionPollingConfig()`:
   ```go
   PollingConfig{
       MaxAttempts:     60,
       InitialInterval: 5 * time.Second,
       MaxInterval:     30 * time.Second,
       Timeout:         30 * time.Minute,
   }
   ```
4. 新增 `PollForResourceGroupReady(ctx, apiClient, imageId, config)`:
   - 与 `pollForStatus` 逻辑类似（指数退避 1.5x）
   - 每次轮询调用 `GetImageInfo`
   - 判断 `info.ResourceGroupReady == true` → 成功
   - 判断 `IsFailed(info.ResourceStatus)` → 失败
   - 打印进度和 RequestId

### Step 7: 命令实现

**文件**: `cmd/image_set_max_session.go`

```go
var imageSetMaxSessionCmd = &cobra.Command{
    Use:   "set-max-session",
    Short: "Set the maximum concurrent session count for an activated User image",
    Long:  `...`, // 含 Examples
    RunE:  runImageSetMaxSession,
}
```

Flags:
- `--image-id` (string, required)
- `--max-session-num` (int, required)

`runImageSetMaxSession` 逻辑:

1. **校验镜像** - 调用 `GetImageInfo`，确认 `IsUserImage` 且 `IsActivated`；打印 GetMcpImageInfo 的 RequestId
2. **调用 API** - `BatchCreateHideResourceGroupsWithMaxSession(ImageId, MaxSessionNum)`；打印该接口的 RequestId（无论成功或失败）
3. **轮询等待** - `PollForResourceGroupReady` 直到 `ResourceGroupReady == true`；每次轮询迭代打印 GetMcpImageInfo 的 RequestId（复用 pollForStatus 模式）

> **RequestId 输出原则**：每个节点的接口请求都必须在终端打印对应的 RequestId，便于用户追踪和排查问题。

### Step 8: 注册命令

**文件**: `cmd/image.go` (init 函数)

添加: `ImageCmd.AddCommand(imageSetMaxSessionCmd)`

### Step 9: Mock 更新

**文件**: `cmd/image_status_helper_test.go` 和 `cmd/image_list_helper_test.go`

两个 mock 类都添加:
```go
func (m *mockXxxClient) BatchCreateHideResourceGroupsWithMaxSession(ctx context.Context, request *client.BatchCreateHideResourceGroupsWithMaxSessionRequest) (*client.BatchCreateHideResourceGroupsWithMaxSessionResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

### Step 10: 单元测试

**文件**: `test/unit/cmd/image_set_max_session_test.go`

- `TestImageSetMaxSessionCmd` - 命令元数据验证 (Use, Short, Long)
- `TestImageSetMaxSessionCmd_RequiredFlags` - `--image-id` 和 `--max-session-num` 必填
- `TestImageSetMaxSessionCmd_SubcommandRegistration` - 注册在 ImageCmd 下

## 实施顺序

1. SDK Request/Response Model (Step 1-2)
2. GetMcpImageInfo 响应扩展 (Step 3)
3. SDK Client 方法 (Step 4)
4. Client Interface + Wrapper (Step 5)
5. Mock 更新 (Step 9) — 紧跟接口变更
6. 轮询逻辑 (Step 6)
7. 命令实现 + 注册 (Step 7-8)
8. 单元测试 (Step 10)

## 验证方式

```bash
# 构建二进制到项目根目录（必须用 -o agentbay，不要仅用 go build ./...）
go build -o agentbay .

# 单元测试
go test ./... -count=1

# 命令帮助
./agentbay image set-max-session --help

# 参数校验
./agentbay image set-max-session  # 应报错缺少必填参数
```
