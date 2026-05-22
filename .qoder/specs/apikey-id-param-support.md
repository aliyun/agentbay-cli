# apikey 命令组支持 `--api-key-id` 参数

## Context

测试自动化从 `apikey create` 流程开始时方便拿到 `ak-xxx`（API Key ID），但不方便拿 `akm-xxx`（API Key）。用户手动 CLI 操作则相反。需要让 apikey 命令组同时支持 `--api-key`（akm-xxx）和 `--api-key-id`（ak-xxx）两种传参方式，与 `concurrency set` 已有的双参数模式保持一致。

另外需要：修正文档中 apikey/api-key-id 术语混淆问题，以及 `apikey create` 输出的 `KeyId` 改名为 `ApiKeyId`。

## 关键设计决策

- **enable/disable/delete**：从位置参数改为 `--api-key` 和 `--api-key-id` 互斥 flag 模式，不保留位置参数向后兼容
- **list**：新增 `--api-key-id` flag，与 `--api-key` 互斥
- **核心原则：`DescribeMcpApiKey` 只在 `--api-key` 路径中使用**，其唯一用途是将 akm-xxx 转换为 ak-xxx。`--api-key-id` 已提供 ak-xxx，应跳过此查找
- **`--api-key-id` 路径的步骤数优化**：
  - **enable/disable**：1 步（直接调用 `ModifyApiKeyStatus(ak-xxx)`），与 `concurrency set` 模式一致
  - **delete**：需要 Status 来判断是否先禁用，用 `DescribeApiKeys(KeyIds=[ak-xxx])` 获取（2-3 步）
  - **list**：1 步（直接传 ak-xxx 给 `DescribeApiKeys`）
- **ModifyApiKeyStatus 的 ApiKey 参数**：虽然参数名叫 `ApiKey`，但实际传入的是 `ak-xxx`（已有代码如此）
- **DescribeApiKeys 成功判定**：使用 `isSuccessCode()`（该 API 返回 Code="200" 而非 "ok"）
- **create 输出**：仅改 `KeyId` → `ApiKeyId`，不同时展示 ApiKey（akm-xxx）

### 各命令 `--api-key-id` 路径流程对比

| 命令 | `--api-key` 路径 | `--api-key-id` 路径 | 需要额外查找？ |
|------|-------------------|----------------------|---------------|
| enable/disable | 2步：DescribeMcpApiKey → ModifyApiKeyStatus | 1步：ModifyApiKeyStatus | 不需要 |
| delete | 3步：DescribeMcpApiKey → (可选)禁用 → 删除 | 2-3步：DescribeApiKeys获取Status → (可选)禁用 → 删除 | 需要Status判断是否先禁用 |
| list | 2步：DescribeMcpApiKey → DescribeApiKeys | 1步：DescribeApiKeys | 不需要 |
| concurrency set | 2步：DescribeMcpApiKey → ModifyMcpApiKeyConfig | 1步：ModifyMcpApiKeyConfig | 不需要（已有实现） |

---

## 实现步骤

### Step 1: 修改 `cmd/apikey_status.go` — enable/disable 改为双 flag 模式

**文件**: `/Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/apikey_status.go`

1. 命令定义修改：
   - `Use`: `"enable <api-key>"` → `"enable"`，`"disable <api-key>"` → `"disable"`
   - 移除 `Args: cobra.ExactArgs(1)`
   - `RunE`: 从 `func(cmd, args) { return runApiKeyStatusChange(cmd, args[0], ...) }` 改为 `func(cmd, args) { return runApiKeyStatusChange(cmd, ...) }`

2. 新增包级变量：
   ```go
   var apikeyStatusApiKey string
   var apikeyStatusApiKeyId string
   ```

3. `init()` 中注册 flags（两个命令共用）：
   ```go
   apikeyEnableCmd.Flags().StringVar(&apikeyStatusApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
   apikeyEnableCmd.Flags().StringVar(&apikeyStatusApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")
   apikeyDisableCmd.Flags().StringVar(&apikeyStatusApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
   apikeyDisableCmd.Flags().StringVar(&apikeyStatusApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")
   ```
   注意：两个命令共享变量，因为不会同时执行。

4. 更新 Long 描述，添加双参数示例

5. `runApiKeyStatusChange` 重构：
   - 签名改为 `runApiKeyStatusChange(cmd *cobra.Command, targetStatus string) error`
   - 顶部加互斥验证
   - `--api-key` 路径（2步）：`DescribeMcpApiKey(akm-xxx)` → 获取 ak-xxx + Name + Status → 检查是否已为目标状态 → `ModifyApiKeyStatus(ak-xxx)`
   - `--api-key-id` 路径（1步）：直接 `ModifyApiKeyStatus(ak-xxx)`，跳过查找，与 `concurrency set` 的 `--api-key-id` 模式一致
   - `--api-key-id` 路径不检查"是否已为目标状态"，不显示 Name，更简洁

   ```go
   // 伪代码
   if apiKey != "" {
       // --api-key 路径：需要先查找转换 akm→ak
       fmt.Printf("[STEP 1/2] Looking up API key...\n")
       descResp := DescribeMcpApiKey(akm-xxx)  // 转换 akm→ak，同时获取 Name/Status
       apiKeyId = descResp.ApiKeyId
       // 检查是否已为目标状态，显示 Name 等
       fmt.Printf("[STEP 2/2] %s API key...\n", action)
       ModifyApiKeyStatus(ak-xxx, targetStatus)
   } else {
       // --api-key-id 路径：直接操作，1步完成
       apiKeyId = apikeyStatusApiKeyId
       fmt.Printf("[STEP 1/1] %s API key...\n", action)
       fmt.Printf("  ApiKeyId: %s\n", apiKeyId)
       ModifyApiKeyStatus(ak-xxx, targetStatus)
   }
   ```

### Step 2: 修改 `cmd/apikey_delete.go` — delete 改为双 flag 模式

**文件**: `/Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/apikey_delete.go`

1. 命令定义修改：
   - `Use`: `"delete <api-key>"` → `"delete"`
   - 移除 `Args: cobra.ExactArgs(1)`
   - `RunE`: 从 `func(cmd, args) { return runApiKeyDelete(cmd, args[0]) }` 改为 `func(cmd, args) { return runApiKeyDelete(cmd) }`

2. 新增包级变量：
   ```go
   var apikeyDeleteApiKey string
   var apikeyDeleteApiKeyId string
   ```

3. `init()` 中注册 flags（保留 `--yes/-y`）：
   ```go
   apikeyDeleteCmd.Flags().StringVar(&apikeyDeleteApiKey, "api-key", "", "User-visible API Key (akm-xxx format, recommended)")
   apikeyDeleteCmd.Flags().StringVar(&apikeyDeleteApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx). Prefer --api-key for normal usage")
   ```

4. 更新 Long 描述，添加双参数示例

5. `runApiKeyDelete` 重构：
   - 签名改为 `runApiKeyDelete(cmd *cobra.Command) error`
   - 顶部加互斥验证
   - `--api-key` 路径（2-3步）：`DescribeMcpApiKey(akm-xxx)` → 获取 ak-xxx + Name + Status → 若 ENABLED 则禁用(1步) → 确认 → 删除
   - `--api-key-id` 路径（2-3步）：`DescribeApiKeys(KeyIds=[ak-xxx])` → 获取 Name + Status → 若 ENABLED 则禁用(1步) → 确认 → 删除
   - 两条路径的区别仅在第一步的查找方式：`--api-key` 用 `DescribeMcpApiKey`（转换 akm→ak），`--api-key-id` 用 `DescribeApiKeys`（用 ak-xxx 查详情）
   - 后续步骤（禁用 → 确认 → 删除）两条路径完全一致
   - 保留 `--yes/-y` 确认跳过逻辑

### Step 3: 修改 `cmd/apikey_list.go` — 新增 `--api-key-id` flag

**文件**: `/Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/apikey_list.go`

1. 新增包级变量：`var apikeyListApiKeyId string`

2. `init()` 中注册 flag：
   ```go
   apikeyListCmd.Flags().StringVar(&apikeyListApiKeyId, "api-key-id", "", "Internal API Key ID (ak-xxx) to filter. Prefer --api-key for normal usage")
   ```

3. 更新 Long 描述

4. `runApikeyList` 修改：
   - 添加 `--api-key` 和 `--api-key-id` 互斥验证
   - `--api-key` 路径不变：DescribeMcpApiKey → 获取 ak-xxx → DescribeApiKeys(KeyIds=[ak-xxx])
   - `--api-key-id` 路径（1步）：直接设置 `apiKeyId = apikeyListApiKeyId`，然后 `DescribeApiKeys(KeyIds=[ak-xxx])`

### Step 4: 修改 `cmd/apikey.go` — create 输出改名

**文件**: `/Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/apikey.go`

- `fmt.Printf("%-*s %s\n", apikeyDetailLabelW, "KeyId:", keyId)` → `"ApiKeyId:"`
- 错误消息 `"missing KeyId"` → `"missing ApiKeyId"`（保持一致性）

### Step 5: 更新单元测试

**文件**: `/Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/test/unit/cmd/apikey_cmd_test.go`

- `TestApiKeyEnableCmd`：`Use` 检查改为 `"enable"`，验证 `--api-key` 和 `--api-key-id` flags 存在
- `TestApiKeyDisableCmd`：同上
- `TestApikeyDeleteCmd`：`Use` 检查改为 `"delete"`，验证 `--api-key` 和 `--api-key-id` flags 存在，保留 `--yes/-y` 测试
- `TestApiKeyCreateCmd`：验证 Long 描述或输出格式提到 `ApiKeyId`

**文件**: `/Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/test/unit/cmd/apikey_list_cmd_test.go`

- 新增 `--api-key-id` flag 验证

### Step 6: 构建和测试验证

```bash
go test ./... -count=1
go build -o agentbay .
./agentbay apikey enable --help
./agentbay apikey disable --help
./agentbay apikey delete --help
./agentbay apikey list --help
./agentbay apikey create --help
```

### Step 7: 文档更新（使用 `update-cli-command-docs` skill）

需要更新的文档文件：

1. **`docs/zh/apikey.md`** 和 **`docs/en/apikey.md`**：
   - 在页面顶部添加术语说明：API Key (akm-xxx) = 用户可见密钥值，API Key ID (ak-xxx) = 内部密钥标识符
   - enable/disable：从位置参数改为 `--api-key` / `--api-key-id` 双 flag
   - delete：从位置参数改为 `--api-key` / `--api-key-id` 双 flag
   - list：新增 `--api-key-id` flag
   - concurrency set：修正 `--api-key` 描述（当前错误写为 "API Key ID"），补充 `--api-key-id` flag

2. **`README.md`** 和 **`README.zh-CN.md`**：
   - Quick Start 示例更新：`apikey enable akm-xxx` → `apikey enable --api-key akm-xxx` 等

3. **`CHANGELOG.md`**：添加变更记录

---

## 不需要修改的文件

- `internal/agentbay/client.go` — Client 接口无变化
- `internal/client/` 目录 — API 接口无变化
- `cmd/concurrency.go` — 已是参考实现，无需改动
- Mock 类 — 无接口变更，无需更新

## 注意事项

- **核心原则：`DescribeMcpApiKey` 只在 `--api-key` 路径中使用**。其唯一用途是将 akm-xxx 转换为 ak-xxx。`--api-key-id` 已提供 ak-xxx，应完全跳过此步骤
- `--api-key-id` 路径中调用 `DescribeApiKeys` 时，使用 `isSuccessCode()` 判定成功（该 API 返回 Code="200"）
- `--api-key-id` 路径中，若 `DescribeApiKeys` 返回空列表，应报错 "API key not found for the given API Key ID"
- enable/disable 的 `--api-key-id` 路径跳过"已为目标状态"检查，直接执行 ModifyApiKeyStatus，与 concurrency set 模式一致
- delete 的 `--api-key-id` 路径仍需 `DescribeApiKeys` 获取 Status（判断是否先禁用），但不用 `DescribeMcpApiKey`
- enable/disable 是存量命令，对 `ModifyApiKeyStatus` 继续使用 `GetSuccess()` 判定
- delete 对 `DeleteApiKey` 继续使用 Code-based SOP
