# image activate 增加 --region-id 参数支持

## Context

当前 `image activate` 命令的 region 信息完全从服务端 `DescribeMcpPolicyData` 返回的 `GroupSpec.RegionId` 获取，用户无法指定部署区域。需要新增 `--region-id` 参数，让用户可以显式指定区域，覆盖服务端默认值。同时 `CreateResourceGroup` 当前未设置 `BizRegionId`，需要补齐（默认用服务端值，用户指定时用指定值）。

完整性检查：activation 流程涉及的所有 API 调用中，只有以下四组有地域相关字段：

- `CreateMcpPolicyData` / `ModifyMcpPolicyData`（RegionName）
- `DescribeOfficeSites`（RegionName）
- `SaveMcpPolicyData`（GroupSpec.RegionName/RegionId + RegionId）
- `CreateResourceGroup`（BizRegionId）

其余接口（`DescribeInstanceTypes`、`GetMcpImageInfo`、`DescribeMcpPolicyData`）均无地域参数，不受影响。

`CreateResourceGroupRequest` 的 `RegionId` 字段为 API 路由用途，不需要设置。

---

## 参数命名决策

使用 `--region-id`（kebab-case），理由：

- 项目所有多词 flag 均使用 kebab-case（`--network-type`, `--lifecycle-mode`, `--biz-region-id` 等）
- 与 `network.go` 中已有的 `--biz-region-id` 风格一致
- 主流 CLI（AWS CLI, Azure CLI, gcloud, kubectl）均使用 kebab-case

---

## 行为规则

**未指定 `--region-id`：**

- `CreateMcpPolicyData`/`ModifyMcpPolicyData`、`DescribeOfficeSites`、`SaveMcpPolicyData`：保持当前行为（region 从服务端 `GroupSpec.RegionId` 获取）
- `CreateResourceGroup` 的 `BizRegionId`：使用 `policyData.GroupSpec.RegionId` 作为默认值（当前未设置，本次补齐）

**指定 `--region-id`：** 统一覆盖以下所有接口的 region 参数

---

## 影响的接口与字段

| 接口 | 字段 | 当前来源 | 改动后行为 |
|------|------|----------|------------|
| CreateMcpPolicyData / ModifyMcpPolicyData | RegionName | policyData.GroupSpec.RegionId | 用户指定时覆盖，否则保持 |
| DescribeOfficeSites (ADVANCED only) | RegionName | policyData.GroupSpec.RegionId | 用户指定时覆盖，否则保持 |
| CreateResourceGroup | BizRegionId | **当前未设置** | 始终设置：用户指定值 > policyData.GroupSpec.RegionId |
| SaveMcpPolicyData | GroupSpec.RegionName/RegionId + saveReq.RegionId | data.GroupSpec.RegionName/RegionId | 用户指定时覆盖，否则保持 |

注意：`CreateResourceGroupRequest` 还有一个 `RegionId` 字段（API 路由用途），不需要设置。

---

## 跨地域 Endpoint 说明（仅供回溯，无需代码调整）

使用非默认地域（如新加坡 `ap-southeast-1`）启用镜像时，不需要更换 endpoint。当前上海的 endpoint 可以直接启用其他地域的镜像，`BizRegionId` 决定资源实际部署位置，与 API endpoint 无关。

---

## 修改文件清单

### 1. `cmd/image.go` — 核心逻辑（主要改动）

#### 1a. 注册 flag — `init()` 函数（line 219 后）

```go
imageActivateCmd.Flags().String("region-id", "", "Region ID for resource deployment (optional, overrides server default)")
```

#### 1b. 读取 flag — `runImageActivate()` 函数（line 1147 后，lifecycle 参数解析之后）

```go
// Parse region parameter
regionId, _ := cmd.Flags().GetString("region-id")
```

#### 1c. 打印信息（line 1196 后，网络类型打印之后）

当 `regionId` 非空时打印：

```go
if regionId != "" {
    fmt.Printf("[REGION] Region ID: %s\n", regionId)
}
```

#### 1d. 更新三个函数签名

| 函数 | 基线位置 | 签名改动 |
|------|----------|----------|
| `handlePolicyDataCreateOrModify` | line 1910 | 增加 `regionId string` 参数（在 `osName string` 之后） |
| `handleAdvancedNetworkActivation` | line 2043 | 增加 `regionId string` 参数；返回值从 `(string, []string, error)` 改为 `(string, []string, string, error)`，新增返回 `serverRegionId` |
| `handleDefaultNetworkActivation` | line 2193 | 增加 `regionId string` 参数；返回值从 `error` 改为 `(string, error)`，新增返回 `serverRegionId` |

`serverRegionId` 为从 `policyResp.Body.Data.GroupSpec.RegionId` 解析出的原始服务端值，供 `CreateResourceGroup.BizRegionId` 使用。

#### 1e. `handlePolicyDataCreateOrModify` 内部（line 1946-1949）

当前：
```go
// RegionName from GroupSpec.RegionId
if policyData.GroupSpec != nil {
    req.RegionName = policyData.GroupSpec.RegionId
}
```

改为：
```go
// RegionName from user-specified regionId or GroupSpec.RegionId
if regionId != "" {
    req.RegionName = dara.String(regionId)
} else if policyData.GroupSpec != nil {
    req.RegionName = policyData.GroupSpec.RegionId
}
```

#### 1f. `handleAdvancedNetworkActivation` 内部

**regionName 解析 + serverRegionId 提取**（line 2066-2070）：

当前：
```go
// Get RegionName from GroupSpec.RegionId for DescribeOfficeSites
var regionName string
if policyResp.Body != nil && policyResp.Body.Data != nil && policyResp.Body.Data.GroupSpec != nil && policyResp.Body.Data.GroupSpec.RegionId != nil {
    regionName = *policyResp.Body.Data.GroupSpec.RegionId
}
```

改为：
```go
// Get RegionName from user-specified regionId or GroupSpec.RegionId for DescribeOfficeSites
var serverRegionId string
var regionName string
if policyResp.Body != nil && policyResp.Body.Data != nil && policyResp.Body.Data.GroupSpec != nil && policyResp.Body.Data.GroupSpec.RegionId != nil {
    serverRegionId = *policyResp.Body.Data.GroupSpec.RegionId
}
if regionId != "" {
    regionName = regionId
} else {
    regionName = serverRegionId
}
```

**SaveMcpPolicyData GroupSpec 覆盖**（line ~2137-2144，GroupSpec 初始化后）：

在 `saveReq.GroupSpec = &client.GroupSpec{...}` 之后添加：
```go
if regionId != "" {
    saveReq.GroupSpec.RegionName = dara.String(regionId)
    saveReq.GroupSpec.RegionId = dara.String(regionId)
}
```

**SaveMcpPolicyData RegionId 覆盖**（line ~2151-2153）：

当前：
```go
if data.GroupSpec != nil && data.GroupSpec.RegionId != nil {
    saveReq.RegionId = data.GroupSpec.RegionId
}
```

改为：
```go
if regionId != "" {
    saveReq.RegionId = dara.String(regionId)
} else if data.GroupSpec != nil && data.GroupSpec.RegionId != nil {
    saveReq.RegionId = data.GroupSpec.RegionId
}
```

**函数末尾返回值**：返回 `serverRegionId`（所有 return 语句从 3 元组改为 4 元组）。

#### 1g. `handleDefaultNetworkActivation` 内部

**serverRegionId 提取**（line 2210 后，DescribeMcpPolicyData 响应处理后）：

添加：
```go
// Extract serverRegionId from GroupSpec
var serverRegionId string
if policyResp.Body != nil && policyResp.Body.Data != nil && policyResp.Body.Data.GroupSpec != nil && policyResp.Body.Data.GroupSpec.RegionId != nil {
    serverRegionId = *policyResp.Body.Data.GroupSpec.RegionId
}
```

**SaveMcpPolicyData GroupSpec 覆盖**（line ~2241-2249，GroupSpec 初始化后）：

在 `saveReq.GroupSpec = &client.GroupSpec{...}` 之后添加：
```go
if regionId != "" {
    saveReq.GroupSpec.RegionName = dara.String(regionId)
    saveReq.GroupSpec.RegionId = dara.String(regionId)
}
```

**SaveMcpPolicyData RegionId 覆盖**（line ~2256-2258）：

当前：
```go
if data.GroupSpec != nil && data.GroupSpec.RegionId != nil {
    saveReq.RegionId = data.GroupSpec.RegionId
}
```

改为：
```go
if regionId != "" {
    saveReq.RegionId = dara.String(regionId)
} else if data.GroupSpec != nil && data.GroupSpec.RegionId != nil {
    saveReq.RegionId = data.GroupSpec.RegionId
}
```

**函数末尾返回值**：从 `return nil` 改为 `return serverRegionId, nil`（所有 return 语句从 error 改为 (string, error)）。

#### 1h. CreateResourceGroup 调用处（line ~1348 后，网络参数设置后、Debug 打印前）

`BizRegionId` 始终设置，优先使用用户指定值，否则使用服务端默认值。

```go
// Set BizRegionId: user-specified regionId takes priority, otherwise use server default
effectiveBizRegionId := regionId
if effectiveBizRegionId == "" {
    effectiveBizRegionId = serverRegionId
}
if effectiveBizRegionId != "" {
    createReq.SetBizRegionId(effectiveBizRegionId)
}
```

方案说明：由于 `CreateResourceGroup` 在 `runImageActivate` 中调用（不在 handler 函数内部），需要从 handler 函数获取服务端的 `GroupSpec.RegionId`。因此：
- `handleAdvancedNetworkActivation` 返回值从 `(string, []string, error)` 改为 `(string, []string, string, error)`，第三个 string 为 serverRegionId
- `handleDefaultNetworkActivation` 返回值从 `error` 改为 `(string, error)`，string 为 serverRegionId

#### 1i. 更新 `runImageActivate` 中两处 handler 调用

在 `var appInstanceType string` 后添加 `var serverRegionId string` 声明。

**ADVANCED flow**（line ~1285）：

当前：
```go
_, effectiveDnsAddresses, err = handleAdvancedNetworkActivation(statusCtx, apiClient, imageId, appInstanceType, cpu, memory, sessionBandwidth, dnsAddresses, lifecycleParams, imageInfo.OsName)
```

改为：
```go
_, effectiveDnsAddresses, serverRegionId, err = handleAdvancedNetworkActivation(statusCtx, apiClient, imageId, appInstanceType, cpu, memory, sessionBandwidth, dnsAddresses, lifecycleParams, imageInfo.OsName, regionId)
```

**DEFAULT flow**（line ~1300）：

当前：
```go
err = handleDefaultNetworkActivation(statusCtx, apiClient, imageId, appInstanceType, cpu, memory, lifecycleParams, imageInfo.OsName)
```

改为：
```go
serverRegionId, err = handleDefaultNetworkActivation(statusCtx, apiClient, imageId, appInstanceType, cpu, memory, lifecycleParams, imageInfo.OsName, regionId)
```

两处 `handlePolicyDataCreateOrModify` 调用（在 handler 内部）也需传入 `regionId`。

#### 1j. 更新 `imageActivateCmd` 的 Long description（line ~91-110）

在 Examples 区域添加：
```
  # Activate with a specific region
  agentbay image activate imgc-xxxxxxxxxxxxxx --region-id cn-shanghai
```

---

### 2. `test/unit/cmd/image_activate_validation_test.go` — 测试

新增 `TestImageActivate_RegionIdResolution` 函数，验证 region 解析优先级逻辑：

| 测试场景 | userRegionId | serverRegionId | 预期结果 |
|----------|-------------|----------------|----------|
| user_specified_overrides_server | cn-shanghai | cn-hangzhou | cn-shanghai |
| no_user_input_uses_server | "" | cn-hangzhou | cn-hangzhou |
| user_specified_with_empty_server | cn-beijing | "" | cn-beijing |
| both_empty | "" | "" | "" |
| user_specifies_overseas_region | ap-southeast-1 | cn-hangzhou | ap-southeast-1 |

测试逻辑模拟 `runImageActivate` 中的 `effectiveBizRegionId` 解析：
```go
effectiveBizRegionId := tt.userRegionId
if effectiveBizRegionId == "" {
    effectiveBizRegionId = tt.serverRegionId
}
assert.Equal(t, tt.expectedRegionId, effectiveBizRegionId)
```

---

### 3. `README.md` — 文档

在 activate 示例区域（Quick Start 部分）添加：
```bash
# Activate with a specific region
agentbay image activate imgc-xxxxx...xxx --region-id cn-shanghai
```

---

### 4. `docs/USER_GUIDE.md` — 用户文档

在 activate 命令的 Options 列表添加：
```
- `--region-id`: Region ID for resource deployment (optional, overrides server default)
```

在 Examples 区域添加：
```bash
# Activate with a specific region
agentbay image activate imgc-xxxxx...xxx --region-id cn-shanghai
```

---

## 数据流总览

```
runImageActivate
├── regionId = cmd.Flags().GetString("region-id")
│
├── [ADVANCED path]
│   └── policyId, dnsAddrs, serverRegionId, err = handleAdvancedNetworkActivation(..., regionId)
│       ├── serverRegionId = *GroupSpec.RegionId（保存原始服务端值）
│       ├── effectiveRegion = regionId || serverRegionId
│       ├── handlePolicyDataCreateOrModify(..., regionId)
│       │   └── req.RegionName = regionId (非空时) / GroupSpec.RegionId (fallback)
│       ├── DescribeOfficeSites(RegionName: effectiveRegion)
│       ├── SaveMcpPolicyData:
│       │   ├── GroupSpec.RegionName/RegionId = regionId (非空时) / 服务端原值
│       │   └── saveReq.RegionId = regionId (非空时) / GroupSpec.RegionId
│       └── return serverRegionId
│
├── [DEFAULT path]
│   └── serverRegionId, err = handleDefaultNetworkActivation(..., regionId)
│       ├── serverRegionId = *GroupSpec.RegionId（保存原始服务端值）
│       ├── handlePolicyDataCreateOrModify(..., regionId)
│       │   └── req.RegionName = regionId (非空时) / GroupSpec.RegionId (fallback)
│       ├── SaveMcpPolicyData:
│       │   ├── GroupSpec.RegionName/RegionId = regionId (非空时) / 服务端原值
│       │   └── saveReq.RegionId = regionId (非空时) / GroupSpec.RegionId
│       └── return serverRegionId
│
└── CreateResourceGroup
    ├── effectiveBizRegionId = regionId || serverRegionId
    └── createReq.SetBizRegionId(effectiveBizRegionId)  ← 始终设置（非空时）
```

---

## 验证方式

1. **编译检查**：`go build ./...` 确保无编译错误
2. **单元测试**：`go test ./... -count=1` 全部通过
3. **手动测试**：
   - `agentbay image activate imgc-xxx` — 不带 region-id，行为不变，BizRegionId 使用服务端默认值
   - `agentbay image activate imgc-xxx --region-id cn-shanghai` — region 覆盖全部四组接口
   - `agentbay image activate imgc-xxx --region-id cn-shanghai --network-type ADVANCED` — ADVANCED 模式下 DescribeOfficeSites 也使用指定 region

---

## 当前实现状态

本需求已在 `feat/image-activate-support-region` 分支上完整实现（4 个未提交文件），验证结果：
- `go build ./...` — 通过
- `go test ./... -count=1` — 全部通过（13 个包）
