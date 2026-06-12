# image activate 支持 CUSTOMIZED 自定义网络

## Context

`image activate` 当前支持 DEFAULT 和 ADVANCED 两种网络类型。需求要求新增 CUSTOMIZED（自定义网络）类型，允许用户使用自有 VPC 网络。CUSTOMIZED 流程比 ADVANCED 更复杂——需要调用 `DescribeOfficeSites`（传入 VpcId）查询办公网络，若无 OfficeSiteId 则需调用新 API `CreateSimpleOfficeSite` 创建办公网络。

## CUSTOMIZED 流程总览

```
STEP 1/8: DescribeInstanceTypes
STEP 2/8: DescribeMcpPolicyData
STEP 3/8: Create/ModifyMcpPolicyData
STEP 4/8: DescribeOfficeSites (OfficeSiteType="CUSTOMIZED", VpcId=用户输入)
           ├── 有 OfficeSiteId → 跳过 STEP 5
           └── 无 OfficeSiteId → 进入 STEP 5
STEP 5/8: CreateSimpleOfficeSite (条件步骤，创建办公网络)
STEP 6/8: SaveMcpPolicyData (NetworkData 含 VpcId, VSwitchId)
STEP 7/8: CreateResourceGroup (含 VpcId, VSwitchId, OfficeSiteId, OfficeSiteType="CUSTOMIZED")
STEP 8/8: Polling
```

**DnsAddress 优先级**: CLI `--dns-address` > DescribeOfficeSites 返回的 DnsAddress > 不传

## 实现步骤

### 1. SDK Model 层

#### 1.1 DescribeOfficeSitesRequest 增加 VpcId
**文件**: `internal/client/describe_office_sites_model.go`
- `DescribeOfficeSitesRequest` 增加 `VpcId *string` 字段 + getter/setter

#### 1.2 NetworkData 增加 VSwitchId
**文件**: `internal/client/describe_mcp_policy_data_model.go`
- `NetworkData` struct 增加 `VSwitchId *string` 字段（在 VpcName 和 SessionBandwidth 之间）+ getter/setter

#### 1.3 新建 CreateSimpleOfficeSite model
**新文件**: `internal/client/create_simple_office_site_model.go`

Request:
```go
type CreateSimpleOfficeSiteRequest struct {
    VpcType            *string  // 固定 "customized"
    OfficeSiteName     *string  // "AgentBay-YYYYMMDD-HH:mm:ss"
    VpcId              *string  // CLI 传入
    RegionId           *string  // CLI 或 API 获取
    RegionName         *string  // 同 RegionId
    DesktopAccessType  *string  // 固定 "INTERNET"
}
```

Response（Data 为字符串即 OfficeSiteId，与 CreateApiKey 模式一致）:
```go
type CreateSimpleOfficeSiteResponseBody struct {
    RequestId      string
    HttpStatusCode *int32
    Data           *string   // 直接是 OfficeSiteId
    Code           *string
    Success        *bool
    Message        *string
}
type CreateSimpleOfficeSiteResponse struct {
    Headers    map[string]*string
    StatusCode *int32
    Body       *CreateSimpleOfficeSiteResponseBody
    RawBody    string
}
```

### 2. SDK API 层

#### 2.1 DescribeOfficeSites 查询参数增加 VpcId
**文件**: `internal/client/advanced_network_api.go`
- `DescribeOfficeSitesWithOptions` 中增加:
  ```go
  if !dara.IsNil(request.VpcId) {
      query["VpcId"] = request.VpcId
  }
  ```

#### 2.2 SaveMcpPolicyData NetworkData 序列化增加 VSwitchId
**文件**: `internal/client/advanced_network_api.go`
- `SaveMcpPolicyDataWithOptions` 中 NetworkData marshalNested 增加:
  ```go
  "VSwitchId": request.NetworkData.VSwitchId,
  ```

#### 2.3 新建 CreateSimpleOfficeSite API 实现
**新文件**: `internal/client/create_simple_office_site_api.go`

- `CreateSimpleOfficeSiteWithOptions` — body 手动序列化每个字段，调用 `client.CallApi`
- `CreateSimpleOfficeSite` — 便捷方法
- `CreateSimpleOfficeSiteWithContext` — context 方法
- `parseCreateSimpleOfficeSiteResponse` — JSON parser（放在同文件，与 advanced_network_api.go 风格一致）
- 成功判定：使用 Code-based SOP（全新 API，Success 字段不确定）
- 请求头: `Accept: application/json`

#### 2.4 CreateSimpleOfficeSite parser 测试
**新文件**: `internal/client/create_simple_office_site_parse_test.go`
- JSON 正常响应（Data 为字符串 OfficeSiteId）
- JSON 失败响应
- JSON Data 为空

### 3. Client Interface 层

#### 3.1 接口增加 CreateSimpleOfficeSite 方法
**文件**: `internal/agentbay/client.go`
- `Client` interface 增加:
  ```go
  CreateSimpleOfficeSite(ctx context.Context, request *client.CreateSimpleOfficeSiteRequest) (*client.CreateSimpleOfficeSiteResponse, error)
  ```
- `clientWrapper` 增加 wrapper 方法（与 DescribeOfficeSites 模式一致）

#### 3.2 更新所有 mock client
**文件**: `cmd/image_status_helper_test.go` — `mockGetMcpImageInfoClient`
**文件**: `cmd/image_list_helper_test.go` — `mockImageListClient`
- 每个增加:
  ```go
  func (m *mockXxxClient) CreateSimpleOfficeSite(ctx context.Context, request *client.CreateSimpleOfficeSiteRequest) (*client.CreateSimpleOfficeSiteResponse, error) {
      return nil, fmt.Errorf("not implemented")
  }
  ```

### 4. CLI Command 层

#### 4.1 Flag 注册
**文件**: `cmd/image.go` `init()` 函数

```go
imageActivateCmd.Flags().String("vpc-id", "", "VPC ID (required for CUSTOMIZED network type)")
imageActivateCmd.Flags().String("vswitch-id", "", "VSwitch ID (required for CUSTOMIZED network type)")
```

更新现有 flag 描述:
- `--network-type`: `"Network type: DEFAULT, ADVANCED, or CUSTOMIZED (default: DEFAULT)"`
- `--dns-address`: `"DNS addresses (for ADVANCED or CUSTOMIZED network, can be specified multiple times)"`

#### 4.2 Flag 解析
**文件**: `cmd/image.go` `runImageActivate` 函数

```go
vpcId, _ := cmd.Flags().GetString("vpc-id")
vswitchId, _ := cmd.Flags().GetString("vswitch-id")
```

#### 4.3 参数校验逻辑
**文件**: `cmd/image.go`

```
networkType 校验: 允许 "DEFAULT" | "ADVANCED" | "CUSTOMIZED"

DEFAULT: 拒绝 --session-bandwidth, --dns-address, --vpc-id, --vswitch-id
ADVANCED: 拒绝 --vpc-id, --vswitch-id
CUSTOMIZED: 必须传 --vpc-id 和 --vswitch-id; 拒绝 --session-bandwidth
```

#### 4.4 状态打印
```go
if networkType == "CUSTOMIZED" {
    fmt.Printf("[NETWORK] Type: CUSTOMIZED\n")
    fmt.Printf("[NETWORK] VPC ID: %s\n", vpcId)
    fmt.Printf("[NETWORK] VSwitch ID: %s\n", vswitchId)
    if len(dnsAddresses) > 0 {
        fmt.Printf("[NETWORK] DNS Addresses: %s\n", strings.Join(dnsAddresses, ", "))
    }
}
```

#### 4.5 CUSTOMIZED 流程分支
**文件**: `cmd/image.go` `runImageActivate`

在 ADVANCED 分支后、DEFAULT 分支前增加 CUSTOMIZED 分支:
```go
if networkType == "CUSTOMIZED" && shouldCreateResourceGroup {
    appInstanceType, err = getAppInstanceType(statusCtx, apiClient, imageId, cpu, memory, 8)
    _, officeSiteId, effectiveDnsAddresses, serverRegionId, err = handleCustomizedNetworkActivation(
        statusCtx, apiClient, imageId, appInstanceType, cpu, memory,
        vpcId, vswitchId, dnsAddresses,
        lifecycleParams, imageInfo.OsName, regionId)
}
```

需要在函数中声明 `var officeSiteId string` 并在 CUSTOMIZED 分支赋值。

#### 4.6 CreateResourceGroup CUSTOMIZED 分支
```go
if networkType == "CUSTOMIZED" {
    createReq.SetOfficeSiteType("CUSTOMIZED")
    createReq.SetVpcId(vpcId)
    createReq.SetVSwitchId(vswitchId)
    if officeSiteId != "" {
        createReq.SetOfficeSiteId(officeSiteId)
    }
    if appInstanceType != "" {
        createReq.SetAppInstanceType(appInstanceType)
    }
    if len(effectiveDnsAddresses) > 0 {
        createReq.SetDnsAddress(effectiveDnsAddresses)
    }
}
```

#### 4.7 实现 handleCustomizedNetworkActivation
**文件**: `cmd/image.go`

新函数，签名:
```go
func handleCustomizedNetworkActivation(
    ctx context.Context, apiClient agentbay.Client,
    imageId, appInstanceType string, cpu, memory int,
    vpcId, vswitchId string, dnsAddresses []string,
    lf *lifecycleFlags, osName, regionId string,
) (policyId, officeSiteId string, effectiveDnsAddresses []string, serverRegionId string, err error)
```

核心逻辑:
1. **STEP 2**: DescribeMcpPolicyData → 获取 policyId, serverRegionId
2. **STEP 3**: handlePolicyDataCreateOrModify（stepNum=3, totalSteps=8）
3. **STEP 4**: DescribeOfficeSites — `OfficeSiteType: "CUSTOMIZED"`, `VpcId: vpcId`, `RegionName: regionName`
   - 提取 `OfficeSiteId` 和 `DnsAddress`
4. **STEP 5**（条件）: 若 OfficeSiteId 为空，调用 CreateSimpleOfficeSite
   - `OfficeSiteName`: `fmt.Sprintf("AgentBay-%s", time.Now().Format("20060102-15:04:05"))`
   - `VpcType: "customized"`, `VpcId`, `DesktopAccessType: "INTERNET"`
   - **RegionId/RegionName 解析逻辑**（与 ADVANCED 中 DescribeOfficeSites 的 regionName 解析一致）:
     ```go
     // regionId 是函数入参（CLI --region-id），可能为空
     // serverRegionId 在 STEP 2 从 DescribeMcpPolicyData 响应的 GroupSpec.RegionId 获取
     effectiveRegionId := regionId
     if effectiveRegionId == "" {
         effectiveRegionId = serverRegionId
     }
     req.RegionId = effectiveRegionId
     req.RegionName = effectiveRegionId  // RegionName = RegionId（前端示例中 RegionName: regionId）
     ```
   - 使用 Code-based 成功判定
   - 若 STEP 4 已返回 OfficeSiteId，则跳过此步骤并打印 "Skipped — office site already exists"
5. **STEP 6**: SaveMcpPolicyData
   - `NetworkData`: `OfficeSiteType: "CUSTOMIZED"`, `VpcId: vpcId`, `VSwitchId: vswitchId`
   - DnsAddress: 有则设置（逗号分隔字符串），无则不设置
6. **返回**: `(policyId, officeSiteId, effectiveDnsAddresses, serverRegionId, error)`

DnsAddress 解析逻辑:
```go
effectiveDnsAddresses = dnsAddresses  // CLI 传入
if len(effectiveDnsAddresses) == 0 && len(describeOfficeSitesDns) > 0 {
    effectiveDnsAddresses = describeOfficeSitesDns
}
// 如果都为空，不设置 DnsAddress
```

### 5. 测试

#### 5.1 更新单元测试
**文件**: `test/unit/cmd/image_activate_validation_test.go`
- 更新 network type 校验测试（CUSTOMIZED 为合法值）
- 新增 CUSTOMIZED 参数校验测试:
  - 缺少 --vpc-id → error
  - 缺少 --vswitch-id → error
  - 带 --session-bandwidth → error
  - 正确参数 → pass
  - 带 --dns-address → pass
  - DEFAULT 带 --vpc-id → error
  - DEFAULT 带 --vswitch-id → error
  - ADVANCED 带 --vpc-id → error
  - ADVANCED 带 --vswitch-id → error

#### 5.2 parser 测试
**新文件**: `internal/client/create_simple_office_site_parse_test.go`

### 6. 文档

#### 6.1 英文文档
**文件**: `docs/en/image.md`
- `--network-type` 描述增加 CUSTOMIZED
- Flags 表格增加 `--vpc-id` 和 `--vswitch-id`
- `--dns-address` 描述更新
- 增加 CUSTOMIZED 示例
- 增加 CUSTOMIZED 注意事项

#### 6.2 中文文档
**文件**: `docs/zh/image.md`
- 与英文文档同步

#### 6.3 README
- 如 Command Overview 表格需要更新则同步

#### 6.4 LLM 文档
- 执行 `bash scripts/build-llms-full.sh`

## 关键修改文件清单

| 文件 | 操作 | 说明 |
|------|------|------|
| `internal/client/describe_office_sites_model.go` | 修改 | Request 增加 VpcId |
| `internal/client/describe_mcp_policy_data_model.go` | 修改 | NetworkData 增加 VSwitchId |
| `internal/client/create_simple_office_site_model.go` | 新建 | 请求/响应模型 |
| `internal/client/create_simple_office_site_api.go` | 新建 | API 实现 + parser |
| `internal/client/create_simple_office_site_parse_test.go` | 新建 | parser 测试 |
| `internal/client/advanced_network_api.go` | 修改 | DescribeOfficeSites 增 VpcId 序列化; SaveMcpPolicyData 增 VSwitchId |
| `internal/agentbay/client.go` | 修改 | 接口+wrapper 增加 CreateSimpleOfficeSite |
| `cmd/image_status_helper_test.go` | 修改 | mock 增加 CreateSimpleOfficeSite |
| `cmd/image_list_helper_test.go` | 修改 | mock 增加 CreateSimpleOfficeSite |
| `cmd/image.go` | 修改 | flag、校验、CUSTOMIZED 流程、handleCustomizedNetworkActivation |
| `test/unit/cmd/image_activate_validation_test.go` | 修改 | 新增 CUSTOMIZED 测试 |
| `docs/en/image.md` | 修改 | 英文文档 |
| `docs/zh/image.md` | 修改 | 中文文档 |

## 回归安全性分析

**对 DEFAULT 和 ADVANCED 网络无影响**，所有改动均为纯增量：

| 改动点 | 影响分析 |
|--------|----------|
| `DescribeOfficeSitesRequest` 增加 `VpcId` | 新增可选 `*string` 字段，默认 `nil`。序列化 `if !dara.IsNil(request.VpcId)` 不输出 nil 字段。ADVANCED 流程不设置该字段，行为不变 |
| `NetworkData` 增加 `VSwitchId` | 新增可选 `*string` 字段，默认 `nil`。DEFAULT/ADVANCED 流程不设置该字段。`marshalNested` 中 `nil` 输出 `"VSwitchId": null`，与现有 `VpcId`/`VpcName` 等字段为 `nil` 时行为一致 |
| 校验逻辑扩展 | 仅增加 DEFAULT 拒绝 `--vpc-id`/`--vswitch-id`（新 flag）和 ADVANCED 拒绝 `--vpc-id`/`--vswitch-id`。现有 DEFAULT 拒绝 `--session-bandwidth`/`--dns-address` 逻辑不变 |
| `handleCustomizedNetworkActivation` 新函数 | 独立新函数，不修改 `handleAdvancedNetworkActivation` 或 `handleDefaultNetworkActivation` |
| `CreateSimpleOfficeSite` 新 API | 仅在 CUSTOMIZED 流程中调用，DEFAULT/ADVANCED 不涉及 |
| `CreateResourceGroup` CUSTOMIZED 分支 | 通过 `if networkType == "CUSTOMIZED"` 隔离，不影响现有 ADVANCED/DEFAULT 分支 |

## 验证

1. `go test ./... -count=1` — 确保全部测试通过
2. `go build -o agentbay .` — 编译成功
3. `./agentbay image activate --help` — 确认新 flag 出现
4. 手动验证参数校验: CUSTOMIZED 缺少必填参数、DEFAULT 带 vpc-id 等场景
5. 预发环境端到端: `AGENTBAY_ENV=prerelease agentbay image activate <image-id> --network-type CUSTOMIZED --vpc-id <vpc-id> --vswitch-id <vswitch-id>`
