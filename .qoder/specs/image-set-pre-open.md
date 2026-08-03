# CLI 调整镜像预开值（image set-pre-open）

## 变更日期

2026-07-28

## 背景与目标

用户需要通过 agentbay-cli 直接调整某个 ACS 镜像的预开值（reserveMinAmount）。该能力需要账号加白后才能使用，且仅支持 ACS 镜像。

本需求在 CLI 侧新增 `agentbay image set-pre-open` 命令，调用后端 `UpdateImageReserveMinAmount` OpenAPI 端点完成预开值设置。

## 方案设计

### CLI 命令设计

新增子命令：

```bash
agentbay image set-pre-open --image-id <image-id> --pre-open <value>
```

**参数：**

| 参数名 | 类型 | 是否必填 | 说明 |
|--------|------|---------|------|
| `--image-id` | String | 是 | 镜像 ID（如 `imgc-xxxxxxxxxxxxxx`） |
| `--pre-open` | Integer | 是 | 预开值（reserveMinAmount），最小值为 1 |

**前置校验：**

- 命令执行前，CLI 校验镜像类型为 User 镜像且处于已激活状态（`RESOURCE_PUBLISHED`）。
- `--pre-open` 小于 1 时，CLI 直接拒绝，不发送请求。
- 预开值上限由服务端根据账号白名单配置决定，CLI 不做客户端上限校验。

**调用流程：**

1. 调用 `GetMcpImageInfo` 获取镜像当前状态。
2. 校验镜像类型与状态。
3. 调用 `UpdateImageReserveMinAmount` 设置预开值。
4. 输出请求 ID 与成功提示；资源扩缩容为异步过程，用户可通过 `agentbay image warmup-status` 查看实际状态。

### 成功判定

`UpdateImageReserveMinAmount` 为新增 OpenAPI 接口，响应中 `Success` 字段可能缺失。命令层采用如下兼容判定：

- `Success` 显式为 `false` 时视为失败。
- `Code` 非空且不等于 `"ok"`（不区分大小写）时视为失败。
- 其余情况视为成功。

失败信息包含 `Code`、`Message` 以及请求 ID，便于排障。

### 响应解析

新增 `parseUpdateImageReserveMinAmountResponse`，按项目 dual-format 模板实现：

- 兼容 JSON 与 XML 两种响应格式。
- `HttpStatusCode` 使用 `int32FromFlexibleJSON` 中转，兼容数字和字符串两种序列化形式。
- 解析失败统一包装为 `ErrWithRequestID`，确保请求 ID 透出到 CLI。

## 代码改动清单

### agentbay-cli

#### 1. SDK 请求/响应模型（新建 2 个文件）

- `internal/client/update_image_reserve_min_amount_request_model.go`
- `internal/client/update_image_reserve_min_amount_response_model.go`

#### 2. SDK Client 方法

- `internal/client/client.go`：新增 `UpdateImageReserveMinAmountWithOptions` 等方法
- `internal/client/dual_format_responses.go`：新增 `parseUpdateImageReserveMinAmountResponse`

#### 3. Client 接口与 Wrapper

- `internal/agentbay/client.go`：接口新增 `UpdateImageReserveMinAmount` 方法签名 + wrapper 实现

#### 4. CLI Cobra 命令

- `cmd/image_set_pre_open.go`（新建）：`set-pre-open` 命令，flags `--image-id` + `--pre-open`
- `cmd/image.go`：注册子命令

#### 5. 单元测试

- `test/unit/cmd/image_set_pre_open_test.go`（新建）：覆盖命令元数据、flag 存在性、子命令注册
- `cmd/image_list_helper_test.go`、`cmd/image_status_helper_test.go`：为变更后的 `Client` 接口新增 `UpdateImageReserveMinAmount` mock 实现

#### 6. 配套文档

- `docs/en/image.md`、`docs/zh/image.md`：补充 `image set-pre-open` 命令说明
- `docs/en/ram-permissions.md`、`docs/zh/ram-permissions.md`：补充 `UpdateImageReserveMinAmount` 所需 RAM 权限
- `docs/internal/cli-openapi-actions.md`：补充 `set-pre-open` 的 OpenAPI Action 映射与调用说明
- `README.md`、`README.zh-CN.md`：Command Overview 表格补充 `set-pre-open`
- `llms.txt`：命令索引补充 `set-pre-open`
- `llms-full.txt`：由 `bash scripts/build-llms-full.sh` 重新生成

## 兼容性说明

- 新增 CLI 命令，不影响现有命令。
- 该功能需要账号加白后才能使用，未加白账号调用会收到服务端错误。
- 仅支持 ACS 镜像，非 ACS 镜像调用会返回错误。
- `reserveMinAmount` 最小值为 1，CLI 与 SDK request 模型均做校验。

## 验证方法

1. 未加白账号调用 `agentbay image set-pre-open` 返回权限错误。
2. 加白账号调用成功更新预开值。
3. CLI `--pre-open 0` 被客户端拒绝（最小值 1）。
4. CLI `--pre-open 41` 被服务端拒绝（超过默认上限 40，除非账号已配置更高的自定义上限）。
5. CLI `agentbay image set-pre-open --image-id imgc-xxx --pre-open 10` 执行成功。
6. 通过 `agentbay image warmup-status` 查看更新后的预开值。

## 修复记录

### 2026-07-29：新增 reserveMinAmount 最小值校验

**变更：**

- `--pre-open` 最小值从 0 调整为 1（不允许设为 0）。
- CLI `image_set_pre_open.go` 校验从 `< 0` 改为 `< 1`，帮助文本同步更新。
- SDK request 模型 `Validate()` 从 `< 0` 改为 `< 1`。
- CLI 文档（`docs/en/image.md`、`docs/zh/image.md`、`llms-full.txt`）中 `≥ 0` 更新为 `≥ 1`。

**涉及文件：**

- `cmd/image_set_pre_open.go`
- `internal/client/update_image_reserve_min_amount_request_model.go`
- `docs/en/image.md`
- `docs/zh/image.md`
- `llms-full.txt`
