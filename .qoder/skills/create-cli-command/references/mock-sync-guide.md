# Mock 同步更新规范

## ⚠️ 重要规则

**当给接口添加新方法时，必须立即更新所有实现该接口的 mock 类！**

这是 CI/CD 流水线中最常见的编译错误来源。

---

## 问题场景

### 触发条件

当在 `internal/agentbay/client.go` 的 `Client` 接口中添加新方法时：

```go
type Client interface {
    // ... 现有方法
    CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error)
    ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error)
}
```

### 错误表现

如果忘记更新 mock 类，CI 会报编译错误：

```
# github.com/agentbay/agentbay-cli/cmd [github.com/agentbay/agentbay-cli/cmd.test]
cmd/image_status_helper_test.go:67:25: cannot use (*mockGetMcpImageInfoClient)(nil) 
  (value of type *mockGetMcpImageInfoClient) as agentbay.Client value in variable declaration: 
  *mockGetMcpImageInfoClient does not implement agentbay.Client (missing method CreateApiKey)
```

### 错误原因

Go 的接口是**隐式实现**的：
- 如果一个类型实现了接口的所有方法，它就实现了该接口
- 如果接口添加了新方法，所有实现类都必须添加该方法
- Mock 类也必须实现接口的所有方法，即使是测试用不到的方法

---

## 解决步骤

### Step 1: 查找所有 mock 类

```bash
# 方法 1: 搜索 mock 类定义
grep -r "type mock.*Client struct" cmd/ test/

# 方法 2: 搜索接口实现声明
grep -r "var _ agentbay.Client" cmd/ test/
```

**当前项目中的 mock 类**（截至 2026-04-09）：
1. `cmd/image_status_helper_test.go` - `mockGetMcpImageInfoClient`
2. `cmd/image_list_helper_test.go` - `mockImageListClient`

### Step 2: 为每个 mock 类添加新方法

#### 示例 1: mockGetMcpImageInfoClient

文件：`cmd/image_status_helper_test.go`

```go
func (m *mockGetMcpImageInfoClient) CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

func (m *mockGetMcpImageInfoClient) ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

#### 示例 2: mockImageListClient

文件：`cmd/image_list_helper_test.go`

```go
func (m *mockImageListClient) CreateApiKey(ctx context.Context, request *client.CreateApiKeyRequest) (*client.CreateApiKeyResponse, error) {
    return nil, fmt.Errorf("not implemented")
}

func (m *mockImageListClient) ModifyMcpApiKeyConfig(ctx context.Context, request *client.ModifyMcpApiKeyConfigRequest) (*client.ModifyMcpApiKeyConfigResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

**注意**: 
- 方法签名必须与接口定义**完全一致**
- 返回 `nil, fmt.Errorf("not implemented")` 即可
- 放在其他 mock 方法之后，`var _ agentbay.Client` 声明之前

### Step 3: 验证编译通过

```bash
# 运行特定包的测试
go test ./cmd/...

# 运行所有测试
make test-unit

# 或者完整测试
go test ./...
```

**预期结果**: 所有测试通过，无编译错误

---

## 检查清单

每次给接口添加新方法后，必须完成以下检查：

- [ ] **查找 mock 类**: 运行 `grep -r "type mock.*Client struct" cmd/ test/`
- [ ] **更新所有 mock**: 为找到的每个 mock 类添加新方法
- [ ] **编译验证**: 运行 `go test ./cmd/...` 确保无编译错误
- [ ] **完整测试**: 运行 `make test-unit` 确保所有测试通过
- [ ] **本地验证**: 确认无误后再提交代码

---

## 最佳实践

### 1. 立即更新，不要拖延

```go
// ❌ 错误：先提交接口改动，稍后再更新 mock
// 1. 更新接口
// 2. 提交代码
// 3. CI 报错
// 4. 再修复 mock

// ✅ 正确：一次性完成所有更新
// 1. 更新接口
// 2. 更新 mock 类
// 3. 本地测试
// 4. 提交代码
```

### 2. 使用脚本辅助

创建快捷脚本查找 mock 类：

```bash
#!/bin/bash
# scripts/find-mocks.sh
echo "=== Mock Client Definitions ==="
grep -rn "type mock.*Client struct" cmd/ test/

echo ""
echo "=== Interface Implementation Declarations ==="
grep -rn "var _ agentbay.Client" cmd/ test/
```

### 3. 在提交前运行测试

```bash
# 提交前必做
make test-unit

# 或者至少运行
go test ./cmd/...
```

### 4. 记录在文档中

每次发现新的 mock 类，更新本文档的"当前项目中的 mock 类"列表。

---

## 常见问题

### Q1: 为什么 mock 方法要返回 "not implemented"？

因为这些方法在特定测试中**不会被调用**，只是为了满足接口契约。返回错误可以：
- 明确标识这是占位实现
- 如果意外调用，会立即发现错误
- 不会干扰正常的测试逻辑

### Q2: 如何知道有哪些 mock 类需要更新？

每次添加新方法后，运行：
```bash
grep -r "type mock.*Client struct" cmd/ test/
```

这个命令会列出所有 mock 类的定义位置。

### Q3: 如果 mock 类很多，有没有自动化方法？

可以编写脚本自动生成 mock 方法，但当前项目 mock 类不多（2-3 个），手动更新更可靠。

### Q4: 能不能不加 mock 方法？

**不行！** Go 编译器会检查接口实现，缺少任何方法都会导致编译失败。

---

## 历史案例

### 案例 1: CreateApiKey 和 ModifyMcpApiKeyConfig（2026-04-09）

**问题**: 在 `Client` 接口中添加了两个新方法，但忘记更新 mock 类。

**影响**: CI 流水线失败，10 个编译错误。

**修复**: 
- 更新 `mockGetMcpImageInfoClient`（添加 2 个方法）
- 更新 `mockImageListClient`（添加 2 个方法）
- 本地验证通过后重新提交

**教训**: 接口变更后必须立即更新所有 mock 类。

---

## 参考链接

- [Go 接口官方文档](https://go.dev/doc/effective_go#interfaces)
- [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces)
- 项目规则：`.qoder/rules/development.md`
- Skill 文档：`.qoder/skills/create-cli-command/SKILL.md`
