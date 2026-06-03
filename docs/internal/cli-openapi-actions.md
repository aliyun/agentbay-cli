# AgentBay CLI 命令 OpenAPI Action 对照表

| CLI 命令                              | OpenAPI Action                                | 说明                                           |
| ------------------------------------- | --------------------------------------------- | ---------------------------------------------- |
| `agentbay image init`                 | `GetDockerfileTemplate`                       | 获取 Dockerfile 模板                           |
| `agentbay docker login`               | `GetACRRepoCredential`                        | 获取 ACR 仓库临时凭证                          |
| `agentbay image create-from-template` | `ListSharedDockerRepos`                       | 校验共享 Docker 仓库授权                       |
|                                       | `CreateImageFromTemplate`                     | 从模板创建镜像                                 |
| `agentbay image activate`             | `GetMcpImageInfo`                             | 获取镜像信息（前置检查）                       |
|                                       | `DescribeInstanceTypes`                       | 获取可用实例规格                               |
|                                       | `DescribeMcpPolicyData`                       | 获取策略数据                                   |
|                                       | `CreateMcpPolicyData` / `ModifyMcpPolicyData` | 创建或修改策略数据                             |
|                                       | `DescribeOfficeSites`（仅 ADVANCED 网络）     | 获取办公网络站点信息                           |
|                                       | `SaveMcpPolicyData`                           | 保存策略数据                                   |
|                                       | `CreateResourceGroup`                         | 创建资源组（实际激活）                         |
|                                       | `GetMcpImageInfo`（轮询）                     | 轮询激活状态                                   |
| `agentbay image set-max-session`      | `GetMcpImageInfo`                             | 校验镜像状态（前置检查）                       |
|                                       | `BatchCreateHideResourceGroupsWithMaxSession` | 设置最大会话数                                 |
|                                       | `GetMcpImageInfo`（轮询）                     | 等待资源组就绪                                 |
| `agentbay image deactivate`           | `GetMcpImageInfo`                             | 获取镜像信息（前置检查）                       |
|                                       | `ListMcpImages`                               | 获取 ResourceGroupId                           |
|                                       | `DeleteResourceGroup`                         | 删除资源组（停用镜像）                         |
|                                       | `GetMcpImageInfo`（轮询）                     | 等待停用完成                                   |
| `agentbay image warmup-status`        | `DescribeWarmUpStatusOpen`                    | 查询预热状态                                   |
| `agentbay image delete`               | `GetMcpImageInfo`                             | 获取镜像信息（前置检查）                       |
|                                       | `DeleteMcpImage`                              | 删除镜像                                       |
| `agentbay image status`               | `GetMcpImageInfo`                             | 查询镜像资源生命周期状态                       |
| `agentbay apikey create`              | `CreateApiKey`                                | 创建 API Key                                   |
| `agentbay apikey enable`              | `ModifyMcpApiKeyConfig`                       | 启用 API Key（Action=EnableMcpApiKey）         |
| `agentbay apikey disable`             | `ModifyMcpApiKeyConfig`                       | 禁用 API Key（Action=DisableMcpApiKey）        |
| `agentbay apikey delete`              | `DeleteApiKey`                                | 删除 API Key                                   |
| `agentbay apikey list`                | `DescribeMcpApiKey`                           | 查询 API Key 列表                              |
| `agentbay apikey concurrency set`     | `ModifyMcpApiKeyConfig`                       | 设置并发上限（Action=SetMcpApiKeyConcurrency） |
| `agentbay network package list`       | `DescribeNetworkPackages`                     | 查询网络包                                     |

## 详细说明

### 1. `agentbay image init`

- **Action**: `GetDockerfileTemplate`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **参数**: Source, SourceImageId, UsePublicNetwork

### 2. `agentbay docker login`

- **Action**: `GetACRRepoCredential`
- **调用方式**: POP RPC V1（原生 ACS HTTP 客户端）
- **参数**: 无

### 3. `agentbay image create-from-template`

- **Action**: `ListSharedDockerRepos`（当 `source-image` 不匹配本地 ACR 缓存时校验 Incoming 共享授权）
- **Action**: `CreateImageFromTemplate`
- **调用方式**: `ListSharedDockerRepos` 使用 OpenAPI SDK；`CreateImageFromTemplate` 使用 POP RPC V1（原生 ACS HTTP 客户端）
- **参数**: `ListSharedDockerRepos` 使用 Direction=Incoming, QueryAliUid, PageStart, PageSize；`CreateImageFromTemplate` 使用 PhysicalImageId, ImageName, TemplateImageId

### 4. `agentbay image activate`

该命令流程较复杂，涉及多个 Action 的编排：

| 步骤               | Action                                         | 用途                           |
| ------------------ | ---------------------------------------------- | ------------------------------ |
| 前置               | `GetMcpImageInfo`                              | 检查镜像类型和当前状态         |
| STEP 1             | `DescribeInstanceTypes`                        | 获取可用实例规格列表           |
| STEP 2             | `DescribeMcpPolicyData`                        | 获取当前策略数据               |
| STEP 3             | `CreateMcpPolicyData` 或 `ModifyMcpPolicyData` | 创建或修改策略                 |
| STEP 4（ADVANCED） | `DescribeOfficeSites`                          | 获取办公网络站点（仅高级网络） |
| STEP 4/5           | `SaveMcpPolicyData`                            | 保存策略数据                   |
| STEP 5/6           | `CreateResourceGroup`                          | 创建资源组，触发激活           |
| STEP 6/7           | `GetMcpImageInfo`（轮询）                      | 轮询等待激活完成               |

> DEFAULT 网络：共 6 步（跳过 DescribeOfficeSites）
> ADVANCED 网络：共 7 步

### 5. `agentbay image set-max-session`

该命令涉及 3 步：

| 步骤   | Action                                        | 用途                                       |
| ------ | --------------------------------------------- | ------------------------------------------ |
| Step 1 | `GetMcpImageInfo`                             | 校验镜像类型和状态（必须为 User 且已激活） |
| Step 2 | `BatchCreateHideResourceGroupsWithMaxSession` | 设置最大会话数                             |
| Step 3 | `GetMcpImageInfo`（轮询）                     | 等待资源组就绪                             |

- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ImageId, MaxSessionNum

### 6. `agentbay image deactivate`

该命令涉及 4 步：

| 步骤   | Action                    | 用途                               |
| ------ | ------------------------- | ---------------------------------- |
| Step 1 | `GetMcpImageInfo`         | 检查镜像类型和当前状态             |
| Step 2 | `ListMcpImages`           | 查询镜像列表获取 ResourceGroupId   |
| Step 3 | `DeleteResourceGroup`     | 删除资源组，触发停用               |
| Step 4 | `GetMcpImageInfo`（轮询） | 等待停用完成（状态变为 Available） |

- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ImageId, ResourceGroupId

### 7. `agentbay image warmup-status`

- **Action**: `DescribeWarmUpStatusOpen`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **参数**: 无

### 8. `agentbay image delete`

- **Action**: `GetMcpImageInfo`（前置） → `DeleteMcpImage`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ImageId

### 9. `agentbay image status`

- **Action**: `GetMcpImageInfo`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ImageId
- **说明**: 与 image activate / set-max-session / deactivate 共用同一 Action，但仅查询状态

### 10. `agentbay apikey create`

- **Action**: `CreateApiKey`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ApiKeyName

### 11. `agentbay apikey enable`

- **Action**: `ModifyMcpApiKeyConfig`（Action=EnableMcpApiKey）
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ApiKeyId

### 12. `agentbay apikey disable`

- **Action**: `ModifyMcpApiKeyConfig`（Action=DisableMcpApiKey）
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ApiKeyId

### 13. `agentbay apikey delete`

- **Action**: `DeleteApiKey`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ApiKeyId

### 14. `agentbay apikey list`

- **Action**: `DescribeMcpApiKey`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: MaxResults, ApiKeyId, NextToken

### 15. `agentbay apikey concurrency set`

- **Action**: `ModifyMcpApiKeyConfig`（Action=SetMcpApiKeyConcurrency）
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: ApiKeyId, MaxConcurrency

### 16. `agentbay network package list`

- **Action**: `DescribeNetworkPackages`
- **调用方式**: OpenAPI SDK（版本 2025-05-01）
- **主要参数**: BizRegionId

## Action 汇总（去重）

共涉及 **20 个** 不同的 OpenAPI Action：

| #   | Action                                        | 涉及命令                                                        |
| --- | --------------------------------------------- | --------------------------------------------------------------- |
| 1   | `GetDockerfileTemplate`                       | image init                                                      |
| 2   | `GetACRRepoCredential`                        | docker login                                                    |
| 3   | `CreateImageFromTemplate`                     | image create-from-template                                      |
| 4   | `GetMcpImageInfo`                             | image activate / set-max-session / deactivate / delete / status |
| 5   | `DescribeInstanceTypes`                       | image activate                                                  |
| 6   | `DescribeMcpPolicyData`                       | image activate                                                  |
| 7   | `CreateMcpPolicyData`                         | image activate                                                  |
| 8   | `ModifyMcpPolicyData`                         | image activate                                                  |
| 9   | `DescribeOfficeSites`                         | image activate (ADVANCED)                                       |
| 10  | `SaveMcpPolicyData`                           | image activate                                                  |
| 11  | `CreateResourceGroup`                         | image activate                                                  |
| 12  | `ListMcpImages`                               | image deactivate                                                |
| 13  | `DeleteResourceGroup`                         | image deactivate                                                |
| 14  | `BatchCreateHideResourceGroupsWithMaxSession` | image set-max-session                                           |
| 15  | `DescribeWarmUpStatusOpen`                    | image warmup-status                                             |
| 16  | `DeleteMcpImage`                              | image delete                                                    |
| 17  | `CreateApiKey`                                | apikey create                                                   |
| 18  | `ModifyMcpApiKeyConfig`                       | apikey enable / disable / concurrency set                       |
| 19  | `DeleteApiKey`                                | apikey delete                                                   |
| 20  | `DescribeMcpApiKey`                           | apikey list                                                     |
