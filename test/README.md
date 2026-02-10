# 测试文档

本项目采用分层测试策略，将测试分为单元测试和集成测试两个层次。

## 测试结构

```
test/
├── unit/                    # 单元测试 - 不依赖外部环境
│   ├── main/                # 主程序测试
│   │   └── main_test.go
│   ├── cmd/                 # 命令测试
│   │   ├── dockerfile_parser_test.go   # Dockerfile COPY/ADD 解析测试
│   │   ├── login_test.go
│   │   ├── logout_test.go
│   │   └── version_test.go
│   └── internal/            # 内部包测试
│       ├── auth/
│       │   └── auth_test.go
│       └── config/
│           └── config_test.go
├── integration/             # 集成测试 - 需要外部API环境
│   └── image_add_copy_upload_test.go  # image create ADD/COPY 上传流程测试
└── README.md               # 本文档
```

## 测试类型

### 单元测试 (Unit Tests)

- **位置**: `test/unit/`
- **特点**: 不依赖外部环境，使用mock服务器
- **覆盖范围**: 
  - 命令行接口逻辑
  - OAuth认证流程（使用mock）
  - 配置管理
  - 错误处理
  - Dockerfile COPY/ADD 解析（parseCOPYADDSources、expandSource、tokenizeInstruction 等）

### 集成测试 (Integration Tests)

- **位置**: `test/integration/`
- **特点**: 测试与真实API的交互
- **覆盖范围**:
  - 完整的OAuth认证流程
  - 真实API调用
  - 网络错误处理
  - image create 中 ADD/COPY 文件解析与上传流程（需设置 RUN_INTEGRATION_TESTS=1）

## 运行测试

### 使用Make命令

```bash
# 运行单元测试（默认）
make test

# 运行单元测试
make test-unit

# 运行集成测试
make test-integration

# 运行所有测试
make test-all

# 运行测试并生成覆盖率报告
make test-coverage
```

### 直接使用Go命令

```bash
# 运行单元测试
go test ./test/unit/...

# 运行集成测试（带构建标签）
go test -tags=integration ./test/integration/...

# 运行所有测试
go test ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./test/unit/... ./internal/... ./cmd/...
go tool cover -html=coverage.out -o coverage.html
```

## 环境变量

### 测试环境变量

| 变量名 | 描述 | 默认值 | 必需 |
|--------|------|--------|------|
| `AGENTBAY_CLI_CONFIG_DIR` | 测试配置目录 | 临时目录 | 否 |
| `SKIP_INTEGRATION_TESTS` | 跳过集成测试 | false | 否 |

## 测试原则

### 单元测试原则

1. **隔离性**: 每个测试独立运行，不依赖其他测试
2. **可重复性**: 测试结果应该是确定的和可重复的
3. **快速性**: 单元测试应该快速执行
4. **无外部依赖**: 使用mock和stub替代外部服务

### 集成测试原则

1. **真实环境**: 测试真实的API交互
2. **容错性**: 能够处理网络错误和API变更
3. **可配置**: 通过环境变量控制测试行为
4. **文档化**: 清楚说明测试的预期行为

## 测试覆盖率

项目目标是保持高测试覆盖率：

- **单元测试覆盖率**: 目标 > 80%
- **关键路径覆盖**: 100%
- **错误处理覆盖**: 100%

查看覆盖率报告：

```bash
make test-coverage
open coverage.html  # macOS
```

## 添加新测试

### 添加单元测试

1. 在 `test/unit/` 对应模块目录下创建测试文件
2. 使用 `httptest` 创建mock服务器
3. 测试所有成功和失败场景
4. 确保测试不依赖外部环境

### 添加集成测试

1. 在 `test/integration/` 目录下创建测试文件
2. 添加构建标签 `//go:build integration`
3. 使用环境变量控制测试行为
4. 处理网络错误和API不可用情况 