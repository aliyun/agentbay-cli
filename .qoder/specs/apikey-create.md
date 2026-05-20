# 规划：`apikey create` 同时支持位置参数和 `--name` 命名参数

## Context

用户希望将 `agentbay apikey create --name "my-api-key"` 修改为支持更简洁的位置参数语法 `agentbay apikey create "my-api-key"`，同时保持向后兼容——已在使用 `--name` 写法的用户不受影响。

## 方案：同时支持两种语法

Cobra 支持同时使用 flag 和 positional args，可以实现两种语法并存：

- `agentbay apikey create "my-api-key"`（新增，位置参数）
- `agentbay apikey create --name "my-api-key"`（保留，命名参数，向后兼容）

优先级：位置参数 > `--name` flag（若同时提供，取位置参数）。

---

## 需要修改的文件

### 1. `cmd/apikey.go`

**修改 `apikeyCreateCmd` 定义**：

- `Use` 字段：从 `"create"` 改为 `"create [name]"` 以反映位置参数可选
- `Args`：从无约束改为 `cobra.MaximumNArgs(1)`，允许 0 或 1 个位置参数
- 移除 `apikeyCreateCmd.MarkFlagRequired("name")`（`--name` 从必填改为可选）
- 更新 `Long` 里的示例，补充位置参数用法
- 修改 `runApikeyCreate` 函数逻辑：

```go
func runApikeyCreate(cmd *cobra.Command, args []string) error {
    var name string
    if len(args) > 0 {
        name = args[0]  // 优先使用位置参数
    } else {
        name = apikeyCreateName  // 回退到 --name flag
    }
    if name == "" {
        return fmt.Errorf("[ERROR] API key name is required. Use: agentbay apikey create <name> or --name <name>")
    }
    // ... 后续逻辑不变
}
```

### 2. `test/unit/cmd/apikey_cmd_test.go`

**更新 `TestApiKeyCreateCmd`**：

- 更新 `Use` 字段断言：期望值从 `"create"` 改为 `"create [name]"`
- 新增测试：`--name` flag 不再是 required（`nameFlag.DefValue == ""`，但不再有 required annotation）
- 新增测试：验证 `Args` 可接受位置参数（通过检查 `Args` 不为 nil 或通过命令行调用验证）

### 3. `README.md` 和 `README.zh-CN.md`

更新 `apikey create` 命令的使用示例，同时展示两种语法：

```
# 使用位置参数（推荐，简洁）
agentbay apikey create "my-api-key"

# 使用 --name 参数（向后兼容）
agentbay apikey create --name "my-api-key"
```

---

## 实现步骤

1. **修改 `cmd/apikey.go`**
   - 更新 `apikeyCreateCmd.Use` 为 `"create [name]"`
   - 添加 `Args: cobra.MaximumNArgs(1)`
   - 移除 `MarkFlagRequired("name")`
   - 更新 Long 描述和示例
   - 修改 `runApikeyCreate` 优先读取位置参数

2. **更新测试 `test/unit/cmd/apikey_cmd_test.go`**
   - 更新 Use 字段断言
   - 验证 `--name` 不再是 required flag
   - 添加 Args 可选位置参数的验证

3. **运行测试和构建**
   ```bash
   go test ./... -count=1
   go build -o agentbay .
   ```

4. **更新 README.md / README.zh-CN.md**
   - 补充位置参数语法示例

---

## 验证方式

- `go test ./... -count=1` 全部通过
- `go build -o agentbay .` 构建成功
- 手动验证：
  - `agentbay apikey create "test-key"` 正常执行
  - `agentbay apikey create --name "test-key"` 正常执行（向后兼容）
  - `agentbay apikey create`（不带参数）返回清晰的错误提示
  - `agentbay apikey create "a" --name "b"` 以位置参数为准（使用 "a"）
