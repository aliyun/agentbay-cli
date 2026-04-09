# CLI 设计规范

## 命令层级设计原则

### ✅ 推荐做法

将相关功能组织为子命令：

```
agentbay
└── apikey                              # 命令组
    ├── create --name <名称>            # 子命令
    └── concurrency                     # 子命令组
        └── set --api-key-id <ID> --concurrency <数值>
```

### ❌ 不推荐做法

创建过多顶级命令：

```
agentbay
├── apikey-create                       # ❌ 不要这样
├── apikey-concurrency-set              # ❌ 不要这样
└── image-list                          # ❌ 不要这样
```

## 参数设计规范

### ✅ 使用命名参数

```bash
# ✅ 推荐：语义清晰
agentbay apikey create --name "my-api-key"
agentbay apikey concurrency set --api-key-id "ak-xxx" --concurrency 10

# ❌ 不推荐：位置参数，不清楚每个参数的含义
agentbay apikey create "my-api-key"
agentbay concurrency set "ak-xxx" 10
```

### 命名参数优势

1. **语义清晰** - 自解释，不需要记忆顺序
2. **顺序无关** - 可以交换参数位置
3. **行业标准** - 符合主流 CLI 工具规范（git, docker, kubectl）
4. **降低学习成本** - 看命令就知道用途

### 参数定义示例

```go
var createName string

func init() {
    // 定义参数
    apikeyCreateCmd.Flags().StringVar(&createName, "name", "", "API key name (required)")
    // 标记为必填
    apikeyCreateCmd.MarkFlagRequired("name")
}
```

## 命令元数据规范

### Use 字段

```go
// ✅ 命令组：简洁的名词
var ApiKeyCmd = &cobra.Command{
    Use:   "apikey",
    // ...
}

// ✅ 子命令：动词或动宾短语
var apikeyCreateCmd = &cobra.Command{
    Use:   "create",  // 不需要写 "create <name>"
    // ...
}
```

### Short 和 Long 字段

```go
var apikeyCreateCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new API key",  // 一句话简介
    Long: `Create a new API key with the specified name.

The API key is used to authenticate requests to AgentBay services.
Each key must have a unique name.

Examples:
  # Create an API key
  agentbay apikey create --name "my-api-key"
  
  # Create with verbose output
  agentbay apikey create --name "production-key" -v`,
    // ...
}
```

### GroupID

```go
// 核心命令
GroupID: "core"

// 管理命令（如 apikey, image 等）
GroupID: "management"
```

## 错误处理规范

### 参数验证

```go
func runApikeyCreate(cmd *cobra.Command, args []string) error {
    name := apikeyCreateName
    
    // 验证参数
    if name == "" {
        return fmt.Errorf("[ERROR] API key name cannot be empty")
    }
    
    // ...
}
```

### API 错误处理

```go
resp, err := apiClient.CreateApiKey(ctx, req)
if err != nil {
    // 打印 RequestId（verbose 模式）
    printRequestIDFromErrIfVerbose(cmd, err)
    
    // 特殊错误处理
    if resp != nil && resp.Body != nil {
        if code := resp.Body.GetCode(); code == "ApiKey.CreateError.NeedCertified" {
            return fmt.Errorf("[ERROR] Failed to create API key: account requires real-name verification")
        }
    }
    
    return fmt.Errorf("[ERROR] Failed to create API key: %w", err)
}
```

## 输出格式规范

### 成功输出

```go
fmt.Println()
fmt.Printf("[SUCCESS] ✅ API key created successfully!\n")
fmt.Printf("%-*s %s\n", 14, "KeyId:", keyId)
fmt.Printf("%-*s %s\n", 14, "Name:", name)
```

**输出示例**:
```
[SUCCESS] ✅ API key created successfully!
KeyId:         ak-d06m2mftwy4jpasw8
Name:          my-api-key
```

### 步骤输出

```go
fmt.Printf("[STEP 1/1] Creating API key...\n")
```

### Verbose 模式

```go
verbose, _ := cmd.Flags().GetBool("verbose")
if verbose && resp.Body.RequestId != nil && *resp.Body.RequestId != "" {
    printRequestIDIfVerbose(cmd, *resp.Body.RequestId)
}
```

## 命令注册规范

在 `main.go` 中注册：

```go
func init() {
    // ...
    rootCmd.AddCommand(cmd.ApiKeyCmd)
}
```

**注意**: 子命令在命令文件的 `init()` 中注册：

```go
func init() {
    ApiKeyCmd.AddCommand(apikeyCreateCmd)
    ApiKeyCmd.AddCommand(ApiKeyConcurrencyCmd)
}
```

## 完整示例

参考文件：
- `cmd/apikey.go` - API Key 创建命令
- `cmd/concurrency.go` - 并发设置命令
