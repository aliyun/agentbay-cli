# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
## [Unreleased]

### 🚀 Features

- Enhance apikey commands with --api-key-id param and update CHANGELOG

### 📖 Documentation

- Reorganize documentation into en/zh structure
- Soften 'agentbay login' wording from deprecated to not recommended
- Fix bilingual switch links and add docs governance spec
- Update CHANGELOG.md for v0.2.10
- Enrich docker login & image create-from-template usage notes


* * *

### 🚀 新功能

- **apikey**：扩展 apikey 命令组，新增 `--api-key-id` 参数支持及 `describe-key-content` 子命令，可通过 API Key ID 查询明文 API Key

### 📖 文档

- 将文档重组为英中双语（`docs/en/` / `docs/zh/`）目录结构
- 将 `agentbay login` 的文案从"已废弃"降级为"不推荐使用"
- 修复双语切换链接，新增文档治理规范
- 更新 v0.2.10 版本 CHANGELOG
- 丰富 `docker login` 和 `image create-from-template` 使用说明


* * *

## [0.2.10] - 2026-05-21

### 🚀 Features

- Support positional arg for apikey create and improve auth error hint
- Add AvailableInstanceSize field to warmup-status image output

### 📖 Documentation

- Update CHANGELOG.md for v0.2.9
- Add image warmup-status command to README (EN & CN)
- Refine Quick Start to focus on API Key lifecycle with list step

### 🛠 Refactoring

- Replace OAuth login hints with AK/SK env var guidance


* * *

### 🚀 新功能

- `apikey create` 支持位置参数 `name`，并优化未登录 / 未配置 AK/SK 时的错误提示
- `image warmup-status` 输出新增 `AvailableInstanceSize` 字段

### 📖 文档

- 更新 v0.2.9 版本 CHANGELOG
- 在中英文 README 中补充 `image warmup-status` 命令说明
- 优化 Quick Start 流程，聚焦 API Key 全生命周期，新增 list 步骤

### 🛠 重构

- 将 OAuth 登录提示替换为 AK/SK 环境变量配置指引


* * *

## [0.2.9] - 2026-05-20

### 🚀 Features

- Add apikey enable/disable commands
- Add apikey delete command with robust success check
- Support --api-key for apikey concurrency set command
- Add apikey list command
- Add image warmup-status command
- Integrate git-cliff for automated changelog generation

### 🐞 Bug Fixes

- Correct git-cliff download URL in homebrew workflow

### 📖 Documentation

- Restructure README with bilingual support, env vars, and command groups
- Add CLI command to OpenAPI action mapping reference


* * *

### 🚀 新功能

- 新增 `apikey enable` / `apikey disable` 命令
- 新增 `apikey delete` 命令，并采用容错的成功判定逻辑
- `apikey concurrency set` 支持通过 `--api-key` 指定密钥
- 新增 `apikey list` 命令
- 新增 `image warmup-status` 命令
- 引入 git-cliff 实现 CHANGELOG 自动生成

### 🐞 修复

- 修复 homebrew workflow 中 git-cliff 下载链接错误

### 📖 文档

- 重构 README，支持中英文双语，补充环境变量与命令分组说明
- 新增 CLI 命令与 OpenAPI Action 映射参考文档


* * *

## [0.2.8] - 2026-05-16

### 🚀 Features

- Support sudo docker


* * *

### 🚀 新功能

- 支持 sudo docker 调用


* * *

## [0.2.7] - 2026-05-15

### 🚀 Features

- Add image set-max-session command for configuring max concurrent sessions
- Support specifying region for image activate

### 🐞 Bug Fixes

- **client**: Tolerate stringified HttpStatusCode in BatchCreateHideResourceGroupsWithMaxSession response

### 📖 Documentation

- **qoder**: Codify response-parsing fault tolerance rules across rules/skill/references
- Clarify image type support for set-max-session command
- Add whitelist requirement for set-max-session command
- Add build verification rule and image activate region spec


* * *

### 🚀 新功能

- 新增 `image set-max-session` 命令，用于配置镜像最大并发会话数
- `image activate` 支持指定 region

### 🐞 修复

- **client**：兼容 `BatchCreateHideResourceGroupsWithMaxSession` 响应中 `HttpStatusCode` 为字符串的场景

### 📖 文档

- **qoder**：在规则 / skill / 参考文档中沉淀响应解析的容错规范
- 明确 `set-max-session` 命令支持的镜像类型
- 补充 `set-max-session` 命令的白名单要求
- 新增构建验证规则与 `image activate` region 规范


* * *

## [0.2.6] - 2026-05-14

### 🚀 Features

- Add docker login and image create-from-template commands
- Add source-image validation and path truncation for create-from-template
- Add UsePublicNetwork param to GetDockerfileTemplate and print RequestId in image init


* * *

### 🚀 新功能

- 新增 `docker login` 和 `image create-from-template` 命令
- `create-from-template` 新增 source-image 校验及路径截断
- `GetDockerfileTemplate` 新增 `UsePublicNetwork` 参数，并在 `image init` 中打印 RequestId


* * *

## [0.2.5] - 2026-05-11

### 🚀 Features

- **image**: Add 'agentbay image delete' command
- Add release-to-oss script and Makefile targets


* * *

### 🚀 新功能

- **image**：新增 `agentbay image delete` 命令
- 新增 `release-to-oss` 脚本及对应的 Makefile 目标


* * *

## [0.2.4] - 2026-05-08

### 🚀 Features

- Enhance image deactivate with RequestId logging and precise ListMcpImages query


* * *

### 🚀 新功能

- 优化 `image deactivate`：增加 RequestId 日志并精准查询 `ListMcpImages`


* * *

## [0.2.3] - 2026-04-29

### 🚀 Features

- Enhance image create file upload with per-file status, auto-retry and summary


* * *

### 🚀 新功能

- 优化 `image create` 文件上传：支持按文件展示状态、自动重试与汇总输出


* * *

## [0.2.2] - 2026-04-28

### 🚀 Features

- Add sandbox lifecycle parameters to image activate command


* * *

### 🚀 新功能

- `image activate` 命令新增沙箱生命周期相关参数


* * *

## [0.2.1] - 2026-04-23

### 🚀 Features

- Add network package list command (DescribeNetworkPackages)
- **network**: Remove UserAliUid param, add default BizRegionId and OfficeSiteId display

### 📖 Documentation

- Add network management to README and update development rules


* * *

### 🚀 新功能

- 新增 `network package list` 命令（`DescribeNetworkPackages`）
- **network**：移除 `UserAliUid` 参数，新增默认 `BizRegionId` 与 `OfficeSiteId` 展示

### 📖 文档

- 在 README 中新增网络管理章节，并更新开发规则


* * *

## [0.2.0] - 2026-04-17

### 🚀 Features

- Add DescribeOfficeSites to ADVANCED network activation flow with DNS default value support
- 调整 DEFAULT 网络激活流程，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

### 📖 Documentation

- Improve README with API key management and image activate examples


* * *

### 🚀 新功能

- ADVANCED 网络激活流程接入 `DescribeOfficeSites`，并支持 DNS 默认值
- 调整 DEFAULT 网络激活流程，增加 `DescribeMcpPolicyData` 与 `SaveMcpPolicyData` 调用

### 📖 文档

- README 增加 API Key 管理与 `image activate` 的使用示例


* * *

## [0.1.9] - 2026-04-10

### 🚀 Features

- Add CreateApiKey CLI command
- Add API key concurrency management CLI command
- Validate add and copy file size

### 🐞 Bug Fixes

- Correct CreateApiKey response Data field type from object to string
- Sync mock implementations after adding Client interface methods

### 📖 Documentation

- Add Qoder rules and skills for CLI development

### 📦 Other Changes

- Add 1 MiB COPY/ADD source limit and retries for OSS uploads and skill push APIs


* * *

### 🚀 新功能

- 新增 `CreateApiKey` CLI 命令
- 新增 API Key 并发管理 CLI 命令
- 校验 add / copy 文件大小

### 🐞 修复

- 修正 `CreateApiKey` 响应中 `Data` 字段的类型（object → string）
- 新增 Client 接口方法后，同步更新 mock 实现

### 📖 文档

- 新增 CLI 开发的 Qoder 规则与 skill

### 📦 其他变更

- 为 OSS 上传与 skill push API 引入 1 MiB 的 COPY/ADD 源大小限制及自动重试


* * *

## [0.1.8] - 2026-03-31

### 📦 Other Changes

- Update Makefile
- Update Makefile


* * *

### 📦 其他变更

- 更新 Makefile
- 更新 Makefile


* * *

## [0.1.3] - 2026-02-10

### 🚀 Features

- Sync latest changes from internal repo


* * *

### 🚀 新功能

- 同步内部仓库的最新变更


* * *

## [0.1.2] - 2025-12-26

### 🐞 Bug Fixes

- Endpoint to xiaoying.cn-shanghai


* * *

### 🐞 修复

- 将默认 endpoint 调整为 `xiaoying.cn-shanghai`


* * *

## [0.1.1] - 2025-12-25

### 🚀 Features

- Add api
- Add windows installation scripts and docs
- Remove unuse files
- Update yml
- Add include system image listing with separeted display sections
- Add dockerfile demo
- Enhance image list with system image support
- Add init dockerfile template
- Add image init guide in user guide
- Add port availability check with retry for login command
- Add advise if port is occupied
- Add unit test for image validation error
- Add image management and port handling improvements

### 📖 Documentation

- Update installation and usage guides
- Add Linux & Mac installation guide and update docs structure

### 🛠 Refactoring

- Add fixossendpoint to suitable return error ossendpoint

### 📦 Other Changes

- Add --include-system and --system-only flags to image list
- Debug --include-system for image list
- Include-system image list fixed
- Update scripts/readme for image list
- Update getdockerfile api
- Set sourceimageid accoding to environment and region
- Update image init info
- Update image to show modify info
- Add port backoff policy
- Fix test error message case sensitivity
- Update homebrew.yml


* * *

### 🚀 新功能

- 新增 API 基础能力
- 新增 Windows 安装脚本与文档
- 删除无用文件并更新 yml
- `image list` 支持系统镜像，按分区分别展示
- 新增 dockerfile 示例与初始化模板
- 在用户指南中补充 `image init` 说明
- `login` 命令新增端口可用性检查与重试
- 端口被占用时给出建议
- 为镜像校验错误新增单元测试
- 完善镜像管理与端口处理

### 📖 文档

- 更新安装与使用指南
- 新增 Linux / Mac 安装指南并调整文档结构

### 🛠 重构

- 新增 `fixossendpoint` 以适配返回的错误 ossendpoint

### 📦 其他变更

- `image list` 新增 `--include-system` 与 `--system-only` 参数
- 调试并修复 `image list` 的 `--include-system` 行为
- 更新 `image list` 的 scripts/readme
- 更新 `getdockerfile` API
- 根据环境与 region 设置 `sourceimageid`
- 更新 `image init` 信息与 `image` 修改展示
- 新增端口退避策略
- 修复测试中错误信息大小写敏感问题
- 更新 `homebrew.yml`


* * *

## [0.1.0] - 2025-10-29

### 🚀 Features

- Add api
- Add api
- Add cicd yml
- Update cicd yml
- Update yml
- Update yml
- Update success html
- Add refresh token
- Add doc
- Add env switch feature
- Add cpu and memory params
- Add windows installation scripts and docs
- Remove unuse files
- Update yml
- Update doc
- Add retry
- Improve authentication error detection and handling
- Add include system image listing with separeted display sections
- Add dockerfile demo
- Enhance image list with system image support
- Add init dockerfile template
- Add image init guide in user guide
- Add port availability check with retry for login command
- Add advise if port is occupied
- Add unit test for image validation error
- 新增 Dockerfile COPY/ADD 文件上传功能
- International env (prod/pre) with default endpoint and OAuth
- **skills**: Skills CLI and API implementation
- Remove skills list and group cmd
- **client**: Add CreateMarketSkill API
- **client**: Add DescribeMarketSkillDetail API
- **client**: Add Market SkillGroup APIs (create, list, add-skill, remove-skill)
- **agentbay**: Extend client with Market Skill and Group APIs
- **cmd**: Add agentbay skills push, list (placeholder), show
- **cmd**: Add agentbay skills group create, list, show (placeholder), add-skill, remove-skill
- **cmd**: Register skills command and extend client mock in tests
- **skills**: Skills CLI and API implementation
- Remove skills list and group cmd
- Add image status cmd
- Fix init parse
- Print requestid when create image verbose

### 🐞 Bug Fixes

- Unit test error
- Unit test error
- Unit test error
- Logout warning
- Cmd issue
- Add sourceimageid
- Image init test
- 适配DeleteResourceGroup接口， 修复deactivate 停止镜像失败问题
- Compile error
- Refresh token using same client
- **skills**: ListMarketGroupSkill XML response and RequestId in -v
- **client**: Align Skills API HTTP method with backend (GET/POST)
- **skills**: Use GET for ListMarketGroupSkill to fix 403 UnsupportedHTTPMethod on pre-release API
- **skills**: Parse CreateMarketSkillGroup Data as string and related client updates
- **skills**: Parse CreateMarketSkillGroup Data correctly and add -v raw response
- **skills,image**: OSS upload and API response handling
- **skills**: Parse XML/JSON for DescribeMarketSkillDetail and AddMarketGroupSkill
- **skills**: Parse XML/JSON for RemoveMarketGroupSkill response
- Compile error
- Parse nug
- Oauth login parse
- Dup error and parse
- Ak-sk skill push parse
- Aone makefile with latest glibc

### 📖 Documentation

- Update environment variable names in windows script
- Add Skills CLI usage (requirement a) to USER_GUIDE
- **plans**: Add Skills API BodyType analysis and backend format reference
- Align README and USER_GUIDE with skills CLI implementation
- Add skills output examples to USER_GUIDE, add manual test results

### 🛠 Refactoring

- Add fixossendpoint to suitable return error ossendpoint

### 📦 Other Changes

- V0.1.0 (#1)
- Add --include-system and --system-only flags to image list
- Debug --include-system for image list
- Include-system image list fixed
- Update scripts/readme for image list
- Update getdockerfile api
- Set sourceimageid accoding to environment and region
- Update image init info
- Update image to show modify info
- Add port backoff policy
- Fix test error message case sensitivity
- Refactor github workflows homebrew
- Require sourceImageId for image init and improve error handling

- Make sourceImageId a required parameter for image init command
- Show raw error info instead of custom error messages
- Update default endpoint to xiaoying.cn-shanghai
- Require sourceImageId for image init and improve error handling

- Make sourceImageId a required parameter for image init command
- Show raw error info instead of custom error messages
- Update default endpoint to xiaoying.cn-shanghai
- 更新user guide
- Parse RPC responses as XML or JSON in the SDK client and drop wrapper-side XML caching


* * *

### 🚀 新功能

- 项目初始化：搭建 CLI 基础能力（API、CI/CD、文档）
- 新增 OAuth 登录与 refresh token 支持
- 新增环境切换能力（生产 / 预发 / 国际站，含默认 endpoint 与 OAuth）
- 新增 CPU 与内存参数
- 新增 Windows 安装脚本与文档
- 新增重试机制
- 提升认证错误检测与处理能力
- `image list` 支持系统镜像，按分区分别展示
- 新增 dockerfile 示例与初始化模板
- 在用户指南中补充 `image init` 说明
- `login` 命令新增端口可用性检查、重试与建议
- 为镜像校验错误新增单元测试
- 新增 Dockerfile COPY/ADD 文件上传功能
- **skills**：实现 Skills CLI 与 API
- **client**：新增 `CreateMarketSkill`、`DescribeMarketSkillDetail` API
- **client**：新增 Market SkillGroup APIs（create、list、add-skill、remove-skill）
- **agentbay**：扩展 client，新增 Market Skill 与 Group 相关 API
- **cmd**：新增 `agentbay skills push` / `list` / `show` 命令
- **cmd**：新增 `agentbay skills group` 子命令组（create、list、show、add-skill、remove-skill）
- **cmd**：注册 skills 命令并在测试中扩展 client mock
- 新增 `image status` 命令
- `image create` 在 verbose 模式下打印 RequestId

### 🐞 修复

- 修复多项单元测试错误
- 修复 logout 警告与命令相关问题
- 补充 `sourceimageid`、修复 `image init` 测试
- 适配 `DeleteResourceGroup` 接口，修复 deactivate 停止镜像失败问题
- 修复编译错误与 refresh token 复用 client 问题
- **skills**：修复 `ListMarketGroupSkill` XML 响应与 verbose 模式下的 RequestId
- **client**：对齐 Skills API HTTP 方法（GET / POST）
- **skills**：使用 GET 调用 `ListMarketGroupSkill` 以修复预发 API 的 403 错误
- **skills**：正确解析 `CreateMarketSkillGroup` 的 `Data` 字段（string），并在 verbose 模式下打印原始响应
- **skills, image**：修复 OSS 上传与 API 响应处理
- **skills**：兼容 XML / JSON 解析 `DescribeMarketSkillDetail`、`AddMarketGroupSkill`、`RemoveMarketGroupSkill` 响应
- 修复 OAuth 登录解析、重复错误及 AK/SK skill push 解析问题
- 修复 Aone Makefile 在最新 glibc 下的构建问题

### 📖 文档

- 更新 Windows 脚本中的环境变量名
- 在 USER_GUIDE 中新增 Skills CLI 用法
- **plans**：新增 Skills API BodyType 分析与后端格式参考
- 对齐 README 与 USER_GUIDE 中的 Skills CLI 实现
- 在 USER_GUIDE 中补充 Skills 输出示例与手动测试结果

### 🛠 重构

- 新增 `fixossendpoint` 以适配返回的错误 ossendpoint

### 📦 其他变更

- 发布 v0.1.0（#1）
- `image list` 新增 `--include-system` 与 `--system-only` 参数
- 调试并修复 `image list` 的 `--include-system` 行为
- 更新 `image list` 的 scripts/readme
- 更新 `getdockerfile` API
- 根据环境与 region 设置 `sourceimageid`
- 更新 `image init` 与 `image` 展示信息
- 新增端口退避策略
- 修复测试中错误信息大小写敏感问题
- 重构 GitHub workflows homebrew
- `image init` 要求必填 `sourceImageId` 并优化错误处理
- 直接展示原始错误信息
- 默认 endpoint 更新为 `xiaoying.cn-shanghai`
- 更新 user guide
- SDK client 将 RPC 响应同时按 XML / JSON 解析，去除 wrapper 端的 XML 缓存


* * *

<!-- generated by git-cliff -->
