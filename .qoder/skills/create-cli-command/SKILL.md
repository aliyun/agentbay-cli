---
name: create-cli-command
description: 将前端 API 能力封装成 agentbay-cli 命令的标准化流程
---

# Create CLI Command

## 📋 职责

将 agent-bay 前端控制台中的 API 能力封装成 agentbay-cli 命令行工具，包括：
- 创建 CLI 命令（使用 Cobra 框架）
- 实现 SDK 客户端方法
- 编写单元测试
- 生成对客文档

## 🎯 触发场景

当用户提出以下需求时触发：
- "帮我把 XX 接口封装成 CLI"
- "新增 XX 功能的命令行工具"
- "把前端的 XX 能力做成 CLI 命令"
- "新增 agentbay xx 命令"

## 🚀 执行步骤

### Phase 1: 需求分析

1. **查找前端 API 实现**
   - 在 `agent-bay/src/http/api/` 目录下搜索接口
   - 确认接口的 Action、Version、Product ID
   - 分析请求参数和响应结构

2. **确认 API 信息**
   - Product ID: `xiaoying`（CLI 使用，不是前端的 xiaoying-double-centre）
   - API Version: `2025-05-01`
   - Endpoint: 根据环境自动选择

3. **与用户确认**
   - 命令名称和层级结构
   - 参数设计（使用命名参数 `--name` 而非位置参数）
   - 是否需要子命令

### Phase 2: 代码实现

#### 2.1 创建 SDK 层模型

```
internal/client/
├── {action}_request_model.go      # 请求模型
└── {action}_response_model.go     # 响应模型
```

**要求**:
- 请求模型必须有 `Validate()` 方法
- 响应模型提供 `GetXxx()` 辅助方法
- 注意后端实际返回的字段类型（可能是字符串而非对象）

#### 2.2 添加 SDK 客户端方法

在 `internal/client/client.go` 中添加：
- `{Action}WithOptions()` - 完整调用方法
- `{Action}()` - 简化调用方法
- `{Action}WithContext()` - 支持 context 的方法
- `parse{Action}Response()` - 响应解析函数

**API 配置模板**:
```go
params := &openapiutil.Params{
    Action:      dara.String("ActionName"),
    Version:     dara.String("2025-05-01"),
    Protocol:    dara.String("HTTPS"),
    Pathname:    dara.String("/"),
    Method:      dara.String("POST"),
    AuthType:    dara.String("AK"),
    Style:       dara.String("RPC"),
    ReqBodyType: dara.String("formData"),
    BodyType:    dara.String("string"),
}
```

#### 2.3 添加客户端接口

在 `internal/agentbay/client.go` 中：
- 在 `Client` interface 中添加方法定义
- 在 `clientWrapper` 中实现方法

**⚠️ 重要：更新接口后必须同步更新所有 Mock 类！**

```bash
# 1. 查找所有实现该接口的 mock 类
grep -r "type mock.*Client struct" cmd/ test/

# 2. 为每个 mock 类添加新方法
# 示例：在 cmd/image_status_helper_test.go 和 cmd/image_list_helper_test.go 中添加
```

```go
// 为每个 mock 类添加（返回 "not implemented"）
func (m *mockClient) NewMethod(ctx context.Context, request *client.NewRequest) (*client.NewResponse, error) {
    return nil, fmt.Errorf("not implemented")
}
```

**检查清单**:
- [ ] 找到所有 mock 类（通常 2-3 个）
- [ ] 为每个 mock 类添加新方法
- [ ] 运行 `go test ./cmd/...` 确保编译通过

#### 2.4 创建 CLI 命令

在 `cmd/` 目录下创建命令文件：
- 使用命名参数（`--name`, `--api-key-id`）
- 标记必填参数：`MarkFlagRequired()`
- 提供清晰的帮助信息和示例
- 实现错误处理和友好提示

**命令层级**:
```
agentbay
└── apikey                          # 命令组
    ├── create --name <名称>        # 子命令
    └── concurrency                 # 子命令组
        └── set --api-key-id <ID> --concurrency <数值>
```

#### 2.5 注册命令

在 `main.go` 的 `init()` 中注册命令：
```go
rootCmd.AddCommand(cmd.XxxCmd)
```

### Phase 3: 单元测试

在 `test/unit/cmd/` 目录创建测试文件：

**测试覆盖要求**:
1. 命令元数据测试（Use, Short, Long, GroupID）
2. 子命令结构测试
3. 必填参数验证测试
4. 参数默认值测试

**测试命名规范**:
```go
Test<命令组>Cmd           // 测试命令组
Test<子命令>Cmd           // 测试子命令
```

### Phase 4: 测试验证

1. **编译测试**
   ```bash
   go build -o agentbay .
   ```

2. **帮助信息测试**
   ```bash
   ./agentbay <command> --help
   ./agentbay <command> <subcommand> --help
   ```

3. **参数校验测试**
   - 缺少必填参数
   - 参数值验证

4. **运行单元测试**
   ```bash
   go test -v ./test/unit/cmd/ -run TestXxx -count=1
   ```

### Phase 5: 文档生成

创建对客功能文档（放在 `cli-analysis/` 目录）：

**文档要求**:
- 面向客户，不包含代码实现细节
- 包含完整的使用示例
- 参数说明表格
- 错误处理和 FAQ
- 认证方式和环境配置

### Phase 6: 代码提交（需用户确认）

**⚠️ 重要**: 必须询问用户是否提交，不要自动提交！

提交前展示：
```bash
git status
git diff --stat
```

询问用户："需要我帮你提交代码吗？"

用户确认后，使用规范的 commit message：
```bash
git add -A
git commit -m "feat: add <功能描述> CLI command

- 具体改动点 1
- 具体改动点 2
- 具体改动点 3"
```

## 📤 输出标准

### 代码输出

✅ **必须包含**:
- [ ] 请求/响应模型文件
- [ ] SDK 客户端方法（3 个变体 + 解析函数）
- [ ] 客户端接口定义和实现
- [ ] CLI 命令文件（使用命名参数）
- [ ] 命令注册（main.go）
- [ ] 单元测试文件

✅ **代码质量**:
- [ ] 编译无错误
- [ ] 所有测试通过
- [ ] 参数校验完整
- [ ] 错误处理友好

### 文档输出

✅ **对客文档**（cli-analysis/）:
- [ ] 功能概述
- [ ] 使用方法（含示例）
- [ ] 参数说明表格
- [ ] 输出示例
- [ ] 错误处理和 FAQ
- [ ] 认证方式说明
- [ ] 环境配置说明

### Git 提交

✅ **提交规范**:
- [ ] 询问用户是否提交
- [ ] 使用 Conventional Commits 格式
- [ ] 包含详细的 commit body
- [ ] 展示提交结果

## 📚 参考资料

- [references/api-format.md](references/api-format.md) - API 格式规范
- [references/cli-design.md](references/cli-design.md) - CLI 设计规范
- [references/test-requirements.md](references/test-requirements.md) - 测试要求
- [references/mock-sync-guide.md](references/mock-sync-guide.md) - Mock 同步更新规范（重要！）
- [references/document-template.md](references/document-template.md) - 文档模板

## 🛠️ 工具脚本

- [scripts/run-tests.sh](scripts/run-tests.sh) - 运行测试
- [scripts/build-and-test.sh](scripts/build-and-test.sh) - 编译并测试
- [scripts/show-diff.sh](scripts/show-diff.sh) - 展示改动

## 📝 模板文件

- [assets/request-model.go](assets/request-model.go) - 请求模型模板
- [assets/response-model.go](assets/response-model.go) - 响应模型模板
- [assets/cli-command.go](assets/cli-command.go) - CLI 命令模板
- [assets/test-file.go](assets/test-file.go) - 测试文件模板
- [assets/customer-doc.md](assets/customer-doc.md) - 对客文档模板

## ⚠️ 注意事项

1. **Product ID**: CLI 使用 `xiaoying`，不是前端的 `xiaoying-double-centre`
2. **响应格式**: 后端返回的字段类型可能与预期不同，需要实际测试确认
3. **参数设计**: 始终使用命名参数（`--name`），不使用位置参数
4. **命令层级**: 相关功能组织为子命令，不要创建顶级命令
5. **测试覆盖**: 必须有单元测试，且所有测试通过
6. **文档面向客户**: 对客文档不包含代码实现细节
7. **不要自动提交**: 必须用户明确要求才执行 git commit
8. **⚠️ 接口变更必须同步 Mock**: 给 `agentbay.Client` 接口添加新方法后，**必须立即更新所有 mock 类**！
   - 查找所有 mock 类：`grep -r "type mock.*Client struct" cmd/ test/`
   - 为每个 mock 类添加新方法（返回 `fmt.Errorf("not implemented")`）
   - 常见 mock 类：`mockGetMcpImageInfoClient`, `mockImageListClient`
   - 否则 CI 会报错：`*mockClient does not implement agentbay.Client (missing method Xxx)`
