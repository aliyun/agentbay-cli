# 单元测试要求

## 测试文件位置

```
test/unit/cmd/
└── <command>_cmd_test.go
```

**示例**:
- `apikey_cmd_test.go` - API Key 相关命令测试
- `image_cmd_test.go` - Image 相关命令测试

## 测试函数命名

```go
// 命令组测试
Test<命令组>Cmd

// 子命令测试
Test<子命令>Cmd

// 示例
TestApiKeyCmd              // apikey 命令组
TestApiKeyCreateCmd        // apikey create 子命令
TestApiKeyConcurrencyCmd   // apikey concurrency 命令组
TestApiKeyConcurrencySetCmd  // apikey concurrency set 子命令
```

## 必需要测试的内容

### 1. 命令元数据测试

```go
t.Run("<command> command has correct metadata", func(t *testing.T) {
    assert.Equal(t, "create", createCmd.Use)
    assert.Equal(t, "Create a new API key", createCmd.Short)
    assert.True(t, strings.Contains(createCmd.Long, "API key"))
})
```

### 2. 子命令结构测试

```go
t.Run("<parent> has subcommands <child1> and <child2>", func(t *testing.T) {
    children := cmd.ParentCmd.Commands()
    names := make([]string, len(children))
    for i, c := range children {
        names[i] = c.Name()
    }
    assert.Contains(t, names, "child1")
    assert.Contains(t, names, "child2")
})
```

### 3. 必填参数测试

```go
t.Run("<command> has required <flag> flag", func(t *testing.T) {
    flag := cmd.Flags().Lookup("flag-name")
    assert.NotNil(t, flag)
    assert.Equal(t, "", flag.DefValue)  // 无默认值
    assert.True(t, strings.Contains(flag.Usage, "required"))
})
```

### 4. 参数验证测试

```go
t.Run("<command> fails without required flags", func(t *testing.T) {
    // 验证参数默认值为空
    flag := cmd.Flags().Lookup("flag-name")
    assert.NotNil(t, flag)
    assert.Equal(t, "", flag.DefValue)
})
```

## 测试示例

### 完整的测试文件结构

```go
// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
    "strings"
    "testing"

    "github.com/spf13/cobra"
    "github.com/stretchr/testify/assert"

    "github.com/agentbay/agentbay-cli/cmd"
)

func TestApiKeyCmd(t *testing.T) {
    t.Run("apikey command has correct metadata", func(t *testing.T) {
        assert.Equal(t, "apikey", cmd.ApiKeyCmd.Use)
        assert.Equal(t, "Manage AgentBay API keys", cmd.ApiKeyCmd.Short)
        assert.Equal(t, "management", cmd.ApiKeyCmd.GroupID)
        assert.True(t, strings.Contains(cmd.ApiKeyCmd.Long, "Create"))
    })

    t.Run("apikey has subcommands create and concurrency", func(t *testing.T) {
        children := cmd.ApiKeyCmd.Commands()
        names := make([]string, len(children))
        for i, c := range children {
            names[i] = c.Name()
        }
        assert.Contains(t, names, "create")
        assert.Contains(t, names, "concurrency")
    })
}

func TestApiKeyCreateCmd(t *testing.T) {
    t.Run("create command has correct metadata", func(t *testing.T) {
        var createCmd *cobra.Command
        for _, c := range cmd.ApiKeyCmd.Commands() {
            if c.Name() == "create" {
                createCmd = c
                break
            }
        }
        
        assert.NotNil(t, createCmd)
        assert.Equal(t, "create", createCmd.Use)
        assert.Equal(t, "Create a new API key", createCmd.Short)
    })

    t.Run("create command has required name flag", func(t *testing.T) {
        var createCmd *cobra.Command
        for _, c := range cmd.ApiKeyCmd.Commands() {
            if c.Name() == "create" {
                createCmd = c
                break
            }
        }
        
        assert.NotNil(t, createCmd)
        
        nameFlag := createCmd.Flags().Lookup("name")
        assert.NotNil(t, nameFlag)
        assert.Equal(t, "", nameFlag.DefValue)
        assert.True(t, strings.Contains(nameFlag.Usage, "required"))
    })
}
```

## 运行测试

### 运行所有测试

```bash
# 运行所有单元测试
go test -v ./test/unit/...

# 运行特定包的测试
go test -v ./test/unit/cmd/
```

### 运行特定测试

```bash
# 运行包含 "ApiKey" 的测试
go test -v ./test/unit/cmd/ -run TestApiKey

# 运行单个测试函数
go test -v ./test/unit/cmd/ -run TestApiKeyCreateCmd

# 运行特定子测试
go test -v ./test/unit/cmd/ -run "TestApiKeyCmd/apikey_has_subcommands"
```

### 强制重新运行（禁用缓存）

```bash
go test -v ./test/unit/cmd/ -run TestApiKey -count=1
```

### 查看测试覆盖率

```bash
go test -v ./test/unit/cmd/ -run TestApiKey -cover
```

## 测试通过标准

✅ **必须满足**:
- 所有测试用例通过
- 无 panic 或 runtime error
- 测试覆盖率合理（至少覆盖命令结构和参数验证）

## 常见断言

```go
// 相等断言
assert.Equal(t, expected, actual)

// 包含断言
assert.Contains(t, names, "create")
assert.True(t, strings.Contains(long, "keyword"))

// 非空断言
assert.NotNil(t, cmd)
assert.NotNil(t, flag)

// 错误断言
assert.Error(t, err)
assert.NoError(t, err)
assert.True(t, strings.Contains(err.Error(), "required"))
```

## 参考文件

- `test/unit/cmd/apikey_cmd_test.go` - 完整的测试示例
- `test/unit/cmd/skills_cmd_test.go` - 另一个测试示例
