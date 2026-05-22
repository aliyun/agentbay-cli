# 为 `apikey concurrency set` 增加 `--api-key` 参数

## Context

当前 `agentbay apikey concurrency set` 只支持 `--api-key-id`（ak-xxx 格式），但用户实际只能拿到 akm-xxx 格式的 API Key，ak-xxx 是内部 ID 用户拿不到。需要增加 `--api-key` 参数，通过 DescribeMcpApiKey 接口查询到 apiKeyId。

**设计决策**（用户已确认）：
- `--api-key` 和 `--api-key-id` 互斥，只允许传一个
- 传 `--api-key` 时动态显示 STEP 1/2 查询 + STEP 2/2 设置；传 `--api-key-id` 时仍为 STEP 1/1

---

## 修改文件清单

### 1. `cmd/concurrency.go` — 主要修改

**变量添加**：
- 新增 `var apiKeyConcurrencySetApiKey string`

**`init()` 修改**：
- 添加 `--api-key` flag：`apiKeyConcurrencySetCmd.Flags().StringVar(&apiKeyConcurrencySetApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")`
- 移除 `apiKeyConcurrencySetCmd.MarkFlagRequired("api-key-id")`（两者都不是单独 required）
- 更新 `--api-key-id` 的 usage 为 `"Internal API Key ID (ak-xxx). Prefer --api-key for normal usage"`
- 保留 `MarkFlagRequired("concurrency")`

**`Long` 描述更新**：
```
Set the maximum number of concurrent sessions for an API key.

The concurrency limit controls how many sessions can run simultaneously
for the specified API key.

You can identify the API key either by its user-visible value (--api-key, akm-xxx)
or by its internal ID (--api-key-id, ak-xxx). Using --api-key is recommended.

Examples:
  # Set concurrency using the user-visible API Key (recommended)
  agentbay apikey concurrency set --api-key "akm-xxx" --concurrency 10

  # Set concurrency using the internal API Key ID
  agentbay apikey concurrency set --api-key-id "ak-xxx" --concurrency 10

  # Set with verbose output
  agentbay apikey concurrency set --api-key "akm-xxx" --concurrency 5 -v
```

**`runApiKeyConcurrencySet` 重写**，逻辑流程：

1. **互斥校验**：
   - 两者都为空 → `"[ERROR] Either --api-key or --api-key-id must be specified. Using --api-key is recommended"`
   - 两者都有值 → `"[ERROR] --api-key and --api-key-id are mutually exclusive; please specify only one"`
2. **并发校验**：`concurrency >= 1`（不变）
3. **初始化客户端**（不变）
4. **分支处理**：
   - **`--api-key` 分支**（两步）：
     - `[STEP 1/2] Looking up API key...`
     - 调用 `DescribeMcpApiKey` → 打印 RequestId → 提取 `apiKeyId`
     - 参考 `apikey_status.go` 中 `runApiKeyStatusChange` 的调用模式
     - `[STEP 2/2] Setting concurrency for API key...`
   - **`--api-key-id` 分支**（单步）：
     - `[STEP 1/1] Setting concurrency for API key...`
5. **调用 `ModifyMcpApiKeyConfig`**（共享）：
   - 始终打印 RequestId（不再依赖 verbose）：`fmt.Printf("[INFO] ModifyMcpApiKeyConfig Request ID: %s\n", ...)`
   - 使用 `printReqIDFromErr(err)` 替换 `printRequestIDFromErrIfVerbose`
6. **成功输出**：打印 ApiKeyId、Concurrency；如果通过 `--api-key` 查询，也打印 ApiKey

### 2. `README.md` — 文档修改

将 `--api-key-id ak-xxx` 改为 `--api-key akm-xxx`，不暴露 `--api-key-id`：

```markdown
#### `apikey concurrency set`

Set the maximum concurrent session limit for an API key.

```bash
agentbay apikey concurrency set --api-key akm-xxx --concurrency 10
```
```

### 3. `README.zh-CN.md` — 中文文档修改

同上，将示例改为 `--api-key akm-xxx`。

### 4. `test/unit/cmd/apikey_cmd_test.go` — 测试修改

**更新 `TestApiKeyConcurrencySetCmd`**：
- 验证 `--api-key` flag 存在，默认值为空，usage 包含 "recommended"
- 验证 `--api-key-id` flag 仍存在，但 usage 不再包含 "required"（改为包含 "Prefer --api-key"）
- 验证 `--concurrency` 仍为 required
- 更新 `Long` 描述断言

---

## 不需要修改的文件

- `internal/agentbay/client.go` — DescribeMcpApiKey 接口已存在
- `internal/client/` — 无新增接口，无需新模型或 parser
- mock 类 — 接口未变更，无需更新

---

## 验证步骤

1. `go test ./... -count=1` — 全量测试通过
2. `go build -o agentbay .` — 构建成功
3. 手动功能测试：
   - `agentbay apikey concurrency set --api-key akm-xxx --concurrency 5` → 显示 STEP 1/2 + STEP 2/2
   - `agentbay apikey concurrency set --api-key-id ak-xxx --concurrency 5` → 显示 STEP 1/1
   - `agentbay apikey concurrency set --concurrency 5` → 报错缺少参数
   - `agentbay apikey concurrency set --api-key akm-xxx --api-key-id ak-xxx --concurrency 5` → 报错互斥
4. `agentbay apikey concurrency set --help` → 同时展示 `--api-key` 和 `--api-key-id`
