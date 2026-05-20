# Warmup Status 命令

## 需求

在 CLI 中提供查询用户预热状态的命令。

## 命令设计

```bash
agentbay image warmup-status
```

- 无需传参
- 返回当前账号的会话配额、镜像配额和各预热镜像详情

## 接口

- Action: `DescribeWarmUpStatusOpen`
- Version: `2025-05-01`
- Product: `xiaoying`

## 返回数据结构

```
Data:
  MaxSessionNumLimit      int32
  TotalUsedSessionQuota   int32
  AvailableSessionQuota   int32
  MaxImageCount           int32
  CurrentImageCount       int32
  Images[]:
    ImageId      string
    TotalMaxSize int32
    GroupCount   int32
```

## 实现文件

### 新增
- `internal/client/describe_warm_up_status_open_request_model.go`
- `internal/client/describe_warm_up_status_open_response_model.go`
- `internal/client/describe_warm_up_status_open_parse_test.go`
- `cmd/image_warmup_status.go`
- `test/unit/cmd/image_warmup_status_test.go`

### 修改
- `internal/client/client.go`
- `internal/client/dual_format_responses.go`
- `internal/agentbay/client.go`
- `cmd/image.go`

### Mock 同步
- `cmd/image_list_helper_test.go`
- `cmd/image_status_helper_test.go`
