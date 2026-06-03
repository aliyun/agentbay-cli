# Skills List 命令实现计划

## Context

当前 `skills list` 命令只是一个占位符，打印 "backend list API is not yet available..."。
现在需要对接 `ListMarketSkillByPage` 接口，实现完整的分页列表功能，支持按技能名称和标签筛选。

---

## 接口信息

- **Action**: `ListMarketSkillByPage`
- **Version**: `2025-05-01`
- **Product ID**: `xiaoying`
- **请求参数**: `PageNo`(int32)、`PageSize`(int32)、`SkillName`(string)、`TagNames`([]string，JSON 数组)
- **响应字段**: `TotalCount`、`TotalPage`、`PageSize`、`PageNumber`、`Result[]`（含 `SkillName`、`SkillId`、`TenantTags`、`SkillStatus`、`GmtModified`）

---

## 实现步骤

### Phase 1: SDK 层 — 请求/响应模型

**新建文件：**
- `internal/client/list_market_skill_by_page_request_model.go`
  - `ListMarketSkillByPageRequest` 结构体：`PageNo *int32`、`PageSize *int32`、`SkillName *string`、`TagNames []string`
  - `Validate()` 方法（无必填参数，直接返回 nil）
  - `GetXxx()` nil-safe getter 方法
  - `SetTagNamesJSON()` 辅助：将 `[]string` 序列化为 JSON 字符串供 query 传参

- `internal/client/list_market_skill_by_page_response_model.go`
  - `ListMarketSkillByPageResponse`：`Headers`、`StatusCode`、`Body`
  - `ListMarketSkillByPageResponseBody`：`Code`、`Data`、`HttpStatusCode`、`Message`、`RequestId`、`Success`
  - `ListMarketSkillByPageResponseBodyData`：`TotalCount`、`TotalPage`、`PageSize`、`PageNumber`、`Result []`
  - `ListMarketSkillByPageResponseBodyDataResult`：`SkillName`、`SkillId`、`TenantTags`、`SkillStatus`、`GmtModified`、`Description`、`Icon`
  - 所有字段的 nil-safe getter

### Phase 2: SDK 层 — client.go + dual_format_responses.go

**修改 `internal/client/client.go`：**
- 添加 `ListMarketSkillByPageWithOptions()`、`ListMarketSkillByPage()`、`ListMarketSkillByPageWithContext()` 三个方法
- API params 模板：`Action: "ListMarketSkillByPage"`，`Method: POST`，`BodyType: string`，`ReqBodyType: formData`
- `TagNames` 参数：序列化为 JSON 字符串后放入 query（如 `["test","aliyun"]`）

**修改 `internal/client/dual_format_responses.go`：**
- 新增 `parseListMarketSkillByPageResponse()` 函数，遵循 dual-format 规范：
  - XML/JSON 双分支
  - `HttpStatusCode` 使用 `int32FromFlexibleJSON` 处理
  - `TotalCount`、`TotalPage`、`PageSize`、`PageNumber` 均用 `int32FromFlexibleJSON`
  - 错误统一用 `ErrWithRequestID` 包装

**新建 `internal/client/list_market_skill_by_page_parse_test.go`：**
- 覆盖三种场景：JSON 数字字段为字符串、JSON 数字字段为数字、XML 分支

### Phase 3: 接口层 — internal/agentbay/client.go

**修改 `internal/agentbay/client.go`：**
- 在 `Client` interface 中添加：
  ```go
  ListMarketSkillByPage(ctx context.Context, request *client.ListMarketSkillByPageRequest) (*client.ListMarketSkillByPageResponse, error)
  ```
- 在 `clientWrapper` 中实现该方法

**同步更新所有 Mock 类（重要！）：**
- `cmd/image_status_helper_test.go` 的 `mockGetMcpImageInfoClient`
- `cmd/image_list_helper_test.go` 的 `mockImageListClient`
- 为每个 mock 类添加 `ListMarketSkillByPage` 方法（返回 `fmt.Errorf("not implemented")`）

### Phase 4: CLI 命令 — cmd/skills.go

**修改 `skillsListCmd` 定义：**
```go
var skillsListCmd = &cobra.Command{
    Use:   "list",
    Short: "List cloud skills",
    Long:  `List skills visible to you (yours and public), with optional filters for name and tags.`,
    Args:  cobra.NoArgs,
    RunE:  runSkillsList,
}
```

**在 `init()` 中注册 flags（skillsListCmd）：**
```go
skillsListCmd.Flags().Int("page", 1, "Page number (default: 1)")
skillsListCmd.Flags().Int("size", 10, "Page size (default: 10)")
skillsListCmd.Flags().String("name", "", "Filter by skill name (optional)")
skillsListCmd.Flags().StringArray("tag", nil, "Filter by tag name (can be specified multiple times, e.g. --tag test --tag aliyun)")
```

**实现 `runSkillsList()` 函数：**
1. 读取 flags：`page`、`size`、`name`、`tag`
2. 加载配置，认证检查
3. 构建 `ListMarketSkillByPageRequest`（`TagNames` 序列化为 JSON 数组字符串）
4. 调用 API（30s 超时上下文）
5. 无条件打印 `[INFO] ListMarketSkillByPage Request ID: ...`
6. 成功判定（以 `Code != "ok"` 为主，参照 SOP）：
   ```go
   code := resp.Body.GetCode()
   successPtr := resp.Body.Success
   if (successPtr != nil && !*successPtr) || (code != "" && !strings.EqualFold(code, "ok")) {
       return fmt.Errorf("[ERROR] Failed to list skills: Code=%s, Message=%s", code, msg)
   }
   ```
7. 分页信息打印：`[PAGE] Page X of Y (Page Size: Z, Total: N)`
8. 表格输出：列宽固定，列名 `SKILL ID`(30)、`SKILL NAME`(30)、`STATUS`(20)、`TAGS`(40)、`MODIFIED`(25)
   - `TenantTags` 拼接为逗号分隔字符串显示

**输出样例（有下一页）：**
```
[INFO] ListMarketSkillByPage Request ID: A4E9C0A5-...
[PAGE] Page 1 of 3 (Page Size: 10, Total: 25)

SKILL ID                          SKILL NAME                     STATUS               TAGS                                     MODIFIED
------------------------------    ----------------------------   ------------------   ---------------------------------------- -------------------------
skill-04p87enx9u4moq5fi           lxy-find-skills                VERIFY_PASSED        哈哈, 阿里云, lxy, test-2, test-1        2026-05-26T02:37:59.000+00:00
skill-04p87lvcjt9o1o9uj           stock-watcher                  INIT                                                          2026-04-04T08:42:11.000+00:00
...

[TIP] Use --page 2 to view the next page.
```

**输出样例（最后一页，无提示）：**
```
[INFO] ListMarketSkillByPage Request ID: A4E9C0A5-...
[PAGE] Page 3 of 3 (Page Size: 10, Total: 25)

SKILL ID                          SKILL NAME                     STATUS               TAGS                                     MODIFIED
...
```

**分页提示逻辑：**
```go
// 打印表格后，判断是否还有下一页
currentPage := resp.Body.GetData().GetPageNumber()  // 当前页码
totalPage   := resp.Body.GetData().GetTotalPage()   // 总页数
if currentPage != nil && totalPage != nil && *currentPage < *totalPage {
    fmt.Printf("\n[TIP] Use --page %d to view the next page.\n", *currentPage+1)
}
```

### Phase 5: 单元测试 — test/unit/cmd/skills_cmd_test.go

在现有测试文件中新增以下测试：
- `skills list has --page flag with default 1`
- `skills list has --size flag with default 10`
- `skills list has --name flag`
- `skills list has --tag flag`
- 保持现有所有测试不变

### Phase 6: 文档同步

**修改 `docs/en/skills.md` 和 `docs/zh/skills.md`：**
- 将 `skills list` 从占位符更新为完整命令说明
- 新增语法、参数表、示例、输出说明章节

**修改 `README.md` 和 `README.zh-CN.md`：**
- 在 RAM 权限表中添加：`ListMarketSkillByPage` → `agentbay:ListMarketSkillByPage` → `skills list`

---

## 关键文件清单

| 操作 | 文件路径 |
|------|---------|
| 新建 | `internal/client/list_market_skill_by_page_request_model.go` |
| 新建 | `internal/client/list_market_skill_by_page_response_model.go` |
| 新建 | `internal/client/list_market_skill_by_page_parse_test.go` |
| 修改 | `internal/client/client.go`（添加 3 个方法） |
| 修改 | `internal/client/dual_format_responses.go`（添加 parser） |
| 修改 | `internal/agentbay/client.go`（接口 + wrapper） |
| 修改 | `cmd/image_status_helper_test.go`（mock 同步） |
| 修改 | `cmd/image_list_helper_test.go`（mock 同步） |
| 修改 | `cmd/skills.go`（skillsListCmd + runSkillsList + flags） |
| 修改 | `test/unit/cmd/skills_cmd_test.go`（新增 list 参数测试） |
| 修改 | `docs/en/skills.md` |
| 修改 | `docs/zh/skills.md` |
| 修改 | `README.md` |
| 修改 | `README.zh-CN.md` |

---

## 验证步骤

```bash
# 1. 全量单元测试
go test ./... -count=1

# 2. 构建二进制
go build -o agentbay .

# 3. 帮助验证
./agentbay skills list --help

# 4. 实机测试（需认证）
./agentbay skills list
./agentbay skills list --page 1 --size 5
./agentbay skills list --name "find"
./agentbay skills list --tag test --tag aliyun
```

---

## 特殊注意事项

1. **TagNames 传参格式**：服务端期望 JSON 数组字符串，需在 SDK 层将 `[]string` 序列化为 `["test","aliyun"]` 后放入 query 参数
2. **成功判定**：接口返回的是 `"code":"200"` 而非 `success:true`，按 SOP 使用 `Code != "ok"` 为主的判定逻辑
3. **分页字段命名**：此接口用 `PageNo`/`PageNumber`（非 `PageStart`），CLI flag 用 `--page`/`--size` 与 image list 保持一致
4. **mock 类同步**：新增 `ListMarketSkillByPage` 方法后必须同步更新 `mockGetMcpImageInfoClient` 和 `mockImageListClient`
