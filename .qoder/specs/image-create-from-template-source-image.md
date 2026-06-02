# Shared Source Image Validation

## Task 1: 明确目标行为

- **现状问题**：[`runImageCreateFromTemplate`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/image_create_from_template.go#L62-L93) 只允许 `--source-image` 匹配本地 `agentbay docker login` 缓存的 `RegistryURL/Namespace/RepoName`，导致被共享的 `/customer_cli/<ownerAliUid>:tag` 无法通过本地前缀校验。
- **目标行为**：
  - 自有仓库镜像：继续允许当前账号 `docker login` 缓存匹配的 `source-image`。
  - 共享仓库镜像：从 `--source-image` 提取 AliUID，调用 `ListSharedDockerRepos`，确认当前账号收到该 AliUID 的共享授权后继续创建。
  - 未授权：在调用 `CreateImageFromTemplate` 前报明确错误。

## Task 2: 设计 `source-image` 解析规则

- 在 [`cmd/image_create_from_template.go`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/image_create_from_template.go) 中新增小型 helper，例如：
  - `parseSourceImageRef(sourceImage string) (registry string, namespace string, repoAliUID int64, tag string, physicalImageId string, err error)`
- 推荐支持格式：
  - 短格式：`/customer_cli/<aliuid>:<tag>`
  - 完整格式：`<registry>/customer_cli/<aliuid>:<tag>`
- 必须校验：
  - 必须包含 tag。
  - namespace 当前先限定为 `customer_cli`，除非后端明确还有其他 namespace。
  - `<aliuid>` 必须可解析为 int64。
- 生成 `PhysicalImageId` 时统一使用短格式：`/customer_cli/<aliuid>:<tag>`，避免共享场景依赖本地缓存 registry。

## Task 3: 重构授权判断逻辑

- 将现有“加载 ACR 缓存失败即报错”的逻辑改成“尽量识别自有仓库；否则走共享授权查询”。
- 建议判断顺序：
  1. 先解析 `--source-image`，得到 `repoAliUID` 和 `physicalImageId`。
  2. 尝试读取本地 ACR 缓存 [`loadACRCredential`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/docker.go#L74-L88)。
  3. 如果缓存存在且 `Namespace/RepoName` 与 `source-image` 匹配：按自有仓库通过。
  4. 如果缓存不存在或不匹配：调用 `ListSharedDockerRepos` 查询共享授权。
- 共享查询复用现有请求模型：[`ListSharedDockerReposRequest`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/internal/client/list_shared_docker_repos_request_model.go#L9-L15)。
- 查询参数建议：
  - `Direction = "Incoming"`
  - `QueryAliUid = repoAliUID`
  - `PageStart = 1`
  - `PageSize = 10`
- 成功判定复用 [`runDockerListShares`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/docker.go#L630-L635) 的规范：
  - `Success` 显式 false 算失败。
  - `Code` 非空且不是 `ok` 算失败。
- 授权判定：按你当前要求，`Data` 非空即认为授权存在；可选增强是只接受 `Status` 为 `ACTIVE` / `Sharing` 的记录，需要你确认是否要加状态白名单。

## Task 3.5: 设计终端输出展示

- `PhysicalImageId` 统一展示短格式：`/customer_cli/<aliuid>:<tag>`，这也是最终传给 `CreateImageFromTemplate` 的值。
- `SourceImage` 建议展示“归一化后的业务源镜像”，规则如下：
  - 自有仓库且用户传完整路径：展示用户传入的完整路径。
  - 自有仓库且用户传短路径：可继续用本地缓存 `RegistryURL` 补全后展示完整路径，保持当前体验。
  - 共享仓库且用户传短路径：不要用当前账号本地缓存的 `RegistryURL` 强行补全，直接展示短路径，避免误导为当前账号 ACR 地址。
  - 共享仓库且用户传完整路径：展示用户传入的完整路径，但仍以解析出的 `/customer_cli/<aliuid>:<tag>` 作为 `PhysicalImageId`。
- 额外新增一行 `SourceType` 或 `Authorization`，用于说明授权来源：
  - 自有仓库：`SourceType: Own repository`
  - 共享仓库：`SourceType: Shared repository (owner AliUID: ****<last4>)`
- 推荐输出示例：
  - 自有短路径输入：
    ```text
    SourceImage:      ai-container-registry.cn-hangzhou.cr.aliyuncs.com/customer_cli/****4214:cli-test-0.0.1
    SourceType:       Own repository
    PhysicalImageId:  /customer_cli/****4214:cli-test-0.0.1
    ```
  - 共享短路径输入：
    ```text
    SourceImage:      /customer_cli/****4214:cli-test-0.0.1
    SourceType:       Shared repository (owner AliUID: ****4214)
    PhysicalImageId:  /customer_cli/****4214:cli-test-0.0.1
    ```

## Task 4: 统一 API Client 使用方式

- 当前 `CreateImageFromTemplate` 使用 raw ACS client：[`newACSClientFromConfig`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/cmd/acs_http_client.go#L54-L88)。
- `ListSharedDockerRepos` 当前通过 SDK wrapper：[`agentbay.Client.ListSharedDockerRepos`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/internal/agentbay/client.go#L467-L471)。
- 方案上可以在 `runImageCreateFromTemplate` 中同时创建：
  - `apiClient := agentbay.NewClientFromConfig(cfg)` 用于 `ListSharedDockerRepos`
  - `acsClient := newACSClientFromConfig(cfg)` 用于 `CreateImageFromTemplate`
- 注意所有 OpenAPI 调用都必须输出 RequestID，且不依赖 `--verbose`：
  - `ListSharedDockerRepos` 成功时打印：`[INFO] ListSharedDockerRepos Request ID: <request-id>`。
  - `ListSharedDockerRepos` 失败时从 `ErrWithRequestID` 中提取并打印 RequestID。
  - `CreateImageFromTemplate` 成功时继续打印响应里的 `RequestId`。
  - `CreateImageFromTemplate` 失败时，如响应 body 或错误中包含 `RequestId`，也要尽量提取并打印，避免失败链路不可追踪。
  - 如果一个执行链路先查共享授权再创建镜像，应输出两个 RequestID，分别标注 Action 名称。

## Task 5: 调整错误提示

- 未授权时错误建议改为类似：
  - `source-image '/customer_cli/<aliuid>:<tag>' is not owned by the current ACR cache and no incoming Docker repo sharing authorization was found for AliUID <aliuid>.`
- 对用户给出修复建议：
  - 如果是自己的镜像：执行 `agentbay docker login` 并使用返回的 registry path/tag。
  - 如果是别人共享的镜像：让共享方执行 `agentbay docker share --target-uid <your-uid>`，接收方可用 `agentbay docker list-shares --direction Incoming --aliuid <owner-uid>` 验证。

## Task 6: 单元测试与回归测试

- 新增/更新测试重点：
  - 短格式 `/customer_cli/<ownUid>:tag` 命中本地缓存，跳过共享查询。
  - 完整格式 `<registry>/customer_cli/<ownUid>:tag` 命中本地缓存。
  - 短格式 `/customer_cli/<sharedOwnerUid>:tag` 本地缓存不匹配，但 `ListSharedDockerRepos` 返回非空，继续创建。
  - `ListSharedDockerRepos` 返回空 `Data`，拒绝创建。
  - `--source-image` 缺 tag、UID 非数字、namespace 不支持时提前失败。
  - 涉及 OpenAPI 的路径必须验证 RequestID 输出：共享查询路径输出 `ListSharedDockerRepos Request ID`，创建镜像路径输出 `CreateImageFromTemplate` 的 `RequestId`，失败路径也要覆盖可提取 RequestID 的场景。
- 若需要 mock `agentbay.Client`，同步更新实现该接口的 mock，避免编译失败。
- 验证命令：
  - `go test ./... -count=1`
  - `go build -o agentbay .`

## Task 7: 文档与权限同步

- 更新文档：
  - [`docs/zh/image.md`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/docs/zh/image.md#L150-L190)
  - [`docs/en/image.md`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/docs/en/image.md#L162-L203)
  - [`docs/zh/image-workflow.md`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/docs/zh/image-workflow.md#L161-L189)
  - [`docs/en/image-workflow.md`](file:///Users/lxy/work/project/ai/agentbay/cli/agentbay-cli/docs/en/image-workflow.md)
- 更新权限说明：`image create-from-template` 现在可能调用两个接口，文档也要说明两个接口都会输出 RequestID：
  - `ListSharedDockerRepos`
  - `CreateImageFromTemplate`
- 因为命令结构没有变化，README Command Overview 通常无需改；但 RAM 权限汇总如果按命令列接口，需要补充 `ListSharedDockerRepos` 对 `image create-from-template` 的影响。
- CHANGELOG readiness 建议标题：`feat(image): support creating images from shared Docker repos`

## Task 8: 开发流程前置

- 正式改代码前，需要先按项目流程确认：
  - 是否复用当前分支，还是从 `aliyun/master` 新建 `feat-image-shared-source`。
  - 使用 Quest Spec 或 `.qoder/changes/CR-<date>-image-shared-source/` 建立变更档案。
- 在你确认方案后，我再进入实现阶段。