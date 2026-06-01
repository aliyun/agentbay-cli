# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.0] - 2026-06-01

### English

#### 🚀 Features

- Add FileUrl field display to skills show command ([09c2914](https://github.com/aliyun/agentbay-cli/commit/09c2914505ff51de839441f1eb323859da1f0049))
- Add TenantTags display to skills show command ([8efaf9b](https://github.com/aliyun/agentbay-cli/commit/8efaf9ba0160a3a550099221ee031086abb0cb04))
- Add --tag flag to skills push command ([47c2ce4](https://github.com/aliyun/agentbay-cli/commit/47c2ce4cfe25bf992668be467386bc017a04b408))
- Add skills update command and skills push --icon support ([cb30097](https://github.com/aliyun/agentbay-cli/commit/cb3009725d2441c2d46b773014e3baa28b4451ec))
- Add --output json support to skills/image/apikey list commands ([d15ae5f](https://github.com/aliyun/agentbay-cli/commit/d15ae5f991477a7b6f6d4150fd23660894984f08))
- Add skills delete command and sync docs/tests ([94c558f](https://github.com/aliyun/agentbay-cli/commit/94c558f0fb73de0082824b9834b646949f56b5e5))
- Add skills-tester agent and pagination test rules ([4d53866](https://github.com/aliyun/agentbay-cli/commit/4d538664c183b5bb1c91b93d62335b385756b2cb))
- Support positional argument for skills delete command ([d147f9c](https://github.com/aliyun/agentbay-cli/commit/d147f9c896f3e0211856f3d019ef4eae93c58857))
- Add --clear-tags flag to skills update command ([9f3bdf9](https://github.com/aliyun/agentbay-cli/commit/9f3bdf9aeaf238395c4582d9539f678fd8f83a63))
- Support short source-image path in image create-from-template; enhance skills list with terminal adaptation and JSON output ([5e1075a](https://github.com/aliyun/agentbay-cli/commit/5e1075a05f104e1ca7fbc491dc33ca6c8a4c775d))
- Add --output json support and terminal-adaptive table to image list command ([f751e5c](https://github.com/aliyun/agentbay-cli/commit/f751e5c7151cb41ad766d5cca93b389b8431b0b4))
- Add docker repo share/unshare/list commands and related docs, RAM permissions guide ([bf8cdf5](https://github.com/aliyun/agentbay-cli/commit/bf8cdf5e2c9ce16ac6ff72cfe2a3162124d7153b))
- Add docker shared repos list command with --output json support ([639234d](https://github.com/aliyun/agentbay-cli/commit/639234d4c9d236642143608e84bf697589cded0b))
- **release**: Add bilingual changelog workflow ([3373cce](https://github.com/aliyun/agentbay-cli/commit/3373cceecd4c39fd2606057531615e6b9643851d))

#### 🐞 Bug Fixes

- Reject RAM identities during OAuth login ([e4453de](https://github.com/aliyun/agentbay-cli/commit/e4453de6190d7a505b5b089f137cf96a68aad500))
- Improve RAM OAuth login guidance ([063ebca](https://github.com/aliyun/agentbay-cli/commit/063ebcadfedf8995b536c2ca2a155986a6805e55))
- Use json accept header for market skill detail ([ab9aac9](https://github.com/aliyun/agentbay-cli/commit/ab9aac9f96283a5fa339cddd55344185f076d7c5))

#### 📖 Documentation

- Restructure README.md and README.zh-CN.md for clarity and completeness ([59789cb](https://github.com/aliyun/agentbay-cli/commit/59789cb601e58be13b47f9cd247b0dce335db45d))
- Mask AliUID in docs and add sensitive info redaction rule ([9604e55](https://github.com/aliyun/agentbay-cli/commit/9604e550c2c8c4bf91591c01a4a46b8048eb21ed))
- Update image workflow docs in zh and en ([76989e1](https://github.com/aliyun/agentbay-cli/commit/76989e106810dc39f4fd146c346a41104d032051))
- Update image workflow guide ([f8ec1e3](https://github.com/aliyun/agentbay-cli/commit/f8ec1e3d80d66bbe95a8259ccdc27dd811786995))
- Update RAM permissions guidance ([96ce1bb](https://github.com/aliyun/agentbay-cli/commit/96ce1bb061c978ebcaa4ad519ee4420366d2c312))
- Organize RAM permissions sections ([c400cf0](https://github.com/aliyun/agentbay-cli/commit/c400cf00dda2f4860f63641f69c1b7569b60f83f))

### 中文

#### 🚀 功能

- **skills**
  - `skills show`：新增 FileUrl 与 TenantTags 字段展示
  - `skills push`：支持 `--tag` 与 `--icon` 参数
  - `skills list`：支持 `--output json`，优化终端自适应展示
  - `skills update`：新增命令，支持 `--clear-tags` 参数
  - `skills delete`：新增命令，支持位置参数
- **image**
  - `image create-from-template`：支持短 source-image 路径
  - `image list`：新增 `--output json`，优化终端自适应表格展示
- **apikey**：`apikey list` 支持 `--output json`，便于脚本和 AI 场景读取结构化结果
- **docker**
  - `docker repo share/unshare/list`：新增仓库分享、取消分享、列表命令，并补充 RAM 权限说明
  - `docker shared repos list`：支持 `--output json`
- **全局**
  - 新增 skills-tester agent 与分页测试规则
  - 新增双语 CHANGELOG 发版流程

#### 🐞 缺陷修复

- **core/auth**：OAuth 登录拒绝 RAM 身份，并优化 RAM OAuth 登录引导。
- **skills**：Market skill detail 接口统一使用 JSON Accept 请求头，提升响应解析稳定性。

#### 📖 文档

- **全局**：重构 README.md 与 README.zh-CN.md，提升结构清晰度和完整性。
- **安全合规**：对文档中的 AliUID 做脱敏处理，并补充敏感信息脱敏规则。
- **image**：更新 image workflow 中英文文档和使用指南。
- **RAM 权限**：更新 RAM 权限说明，并整理 RAM permissions 章节。

---

## [Unreleased]

## [0.3.0] - 2026-05-22

### English

#### 🚀 Features

- Enhance apikey commands with --api-key-id param and update CHANGELOG
- Add 'apikey describe-key-content' command

#### 📖 Documentation

- Reorganize documentation into en/zh structure
- Soften 'agentbay login' wording from deprecated to not recommended
- Fix bilingual switch links and add docs governance spec
- Update CHANGELOG.md for v0.2.10
- Enrich docker login & image create-from-template usage notes
- Add RAM permission requirements for apikey commands
- Extend RAM permission docs to all command groups
- Backfill Chinese translations for all historical CHANGELOG versions
- Add update and uninstall instructions to install guides

### 中文

#### 🚀 功能

- **apikey**
  - 相关命令支持 `--api-key-id` 参数
  - 新增 `apikey describe-key-content` 命令

#### 📖 文档

- **全局**
  - 将文档重组为 en/zh 双语结构
  - 修复双语切换链接，补充文档治理规范
- **core/auth**：将 `agentbay login` 文案从“已废弃”调整为“不推荐”，弱化误导性表达。
- **docker/image**：补充 docker login 与 `image create-from-template` 使用说明。
- **apikey/RAM 权限**：补充 apikey 命令 RAM 权限要求，并扩展到所有命令组。
- **发版**
  - 回填历史 CHANGELOG 中文翻译
  - 补充安装文档中的更新与卸载说明

---

## [0.2.10] - 2026-05-21

### English

#### 🚀 Features

- Support positional arg for apikey create and improve auth error hint
- Add AvailableInstanceSize field to warmup-status image output

#### 📖 Documentation

- Update CHANGELOG.md for v0.2.9
- Add image warmup-status command to README (EN & CN)
- Refine Quick Start to focus on API Key lifecycle with list step

#### 🛠 Refactoring

- Replace OAuth login hints with AK/SK env var guidance

### 中文

#### 🚀 功能

- **apikey**：`apikey create` 支持位置参数，并优化认证错误提示。
- **image**：`image warmup-status` 输出新增 AvailableInstanceSize 字段。

#### 📖 文档

- **发版**：更新 v0.2.9 CHANGELOG。
- **image**：在中英文 README 中补充 `image warmup-status` 命令。
- **core**：重构 Quick Start，使其聚焦 API Key 生命周期并加入 list 步骤。

#### 🛠 重构

- **core/auth**：将 OAuth 登录提示替换为 AK/SK 环境变量指引。

---

## [0.2.9] - 2026-05-20

### English

#### 🚀 Features

- Add apikey enable/disable commands
- Add apikey delete command with robust success check
- Support --api-key for apikey concurrency set command
- Add apikey list command
- Add image warmup-status command
- Integrate git-cliff for automated changelog generation

#### 🐞 Bug Fixes

- Correct git-cliff download URL in homebrew workflow

#### 📖 Documentation

- Restructure README with bilingual support, env vars, and command groups
- Add CLI command to OpenAPI action mapping reference

### 中文

#### 🚀 功能

- **apikey**
  - 新增 enable/disable 命令
  - 新增 delete 命令
  - 新增 list 命令
  - `apikey concurrency set` 支持 `--api-key` 参数
- **image**：新增 `image warmup-status` 命令。
- **发版**：集成 git-cliff 自动生成 CHANGELOG。

#### 🐞 缺陷修复

- **CI/CD**：修复 homebrew workflow 中 git-cliff 下载 URL。

#### 📖 文档

- **全局**：重构 README，补充双语、环境变量与命令组说明。
- **内部参考**：新增 CLI 命令到 OpenAPI Action 的映射文档。

---

## [0.2.8] - 2026-05-16

### English

#### 🚀 Features

- Support sudo docker

### 中文

#### 🚀 功能

- **docker**：支持 sudo docker 场景。

---

## [0.2.7] - 2026-05-15

### English

#### 🚀 Features

- Add image set-max-session command for configuring max concurrent sessions
- Support specifying region for image activate

#### 🐞 Bug Fixes

- **client**: Tolerate stringified HttpStatusCode in BatchCreateHideResourceGroupsWithMaxSession response

#### 📖 Documentation

- **qoder**: Codify response-parsing fault tolerance rules across rules/skill/references
- Clarify image type support for set-max-session command
- Add whitelist requirement for set-max-session command
- Add build verification rule and image activate region spec

### 中文

#### 🚀 功能

- **image**
  - 新增 `image set-max-session`，用于配置最大并发会话数
  - `image activate` 支持指定地域

#### 🐞 缺陷修复

- **client**：BatchCreateHideResourceGroupsWithMaxSession 响应解析兼容字符串形式的 HttpStatusCode。

#### 📖 文档

- **qoder**：在规则、skill 与参考文档中固化响应解析容错规范。
- **image**
  - 说明 `set-max-session` 支持的镜像类型和白名单要求
  - 补充 `image activate` 地域规范
- **构建**：补充构建验证规则。

---

## [0.2.6] - 2026-05-14

### English

#### 🚀 Features

- Add docker login and image create-from-template commands
- Add source-image validation and path truncation for create-from-template
- Add UsePublicNetwork param to GetDockerfileTemplate and print RequestId in image init

### 中文

#### 🚀 功能

- **docker**：新增 `docker login` 命令。
- **image**
  - 新增 `image create-from-template` 命令，支持 source-image 校验与路径截断
  - `image init` 新增 UsePublicNetwork 参数，并打印 RequestId

---

## [0.2.5] - 2026-05-11

### English

#### 🚀 Features

- **image**: Add 'agentbay image delete' command
- Add release-to-oss script and Makefile targets

### 中文

#### 🚀 功能

- **image**：新增 `agentbay image delete` 命令。
- **发版**：新增 release-to-oss 脚本和 Makefile 发布目标。

---

## [0.2.4] - 2026-05-08

### English

#### 🚀 Features

- Enhance image deactivate with RequestId logging and precise ListMcpImages query

### 中文

#### 🚀 功能

- **image**：增强 `image deactivate`，增加 RequestId 日志并精准查询 ListMcpImages。

---

## [0.2.3] - 2026-04-29

### English

#### 🚀 Features

- Enhance image create file upload with per-file status, auto-retry and summary

### 中文

#### 🚀 功能

- **image**：增强 `image create` 文件上传，支持逐文件状态展示、自动重试和汇总结果。

---

## [0.2.2] - 2026-04-28

### English

#### 🚀 Features

- Add sandbox lifecycle parameters to image activate command

### 中文

#### 🚀 功能

- **image**：`image activate` 支持沙箱生命周期参数。

---

## [0.2.1] - 2026-04-23

### English

#### 🚀 Features

- Add network package list command (DescribeNetworkPackages)
- **network**: Remove UserAliUid param, add default BizRegionId and OfficeSiteId display

#### 📖 Documentation

- Add network management to README and update development rules

### 中文

#### 🚀 功能

- **network**
  - 新增 network package list 命令（DescribeNetworkPackages）
  - 移除 UserAliUid 参数，增加默认 BizRegionId，并展示 OfficeSiteId

#### 📖 文档

- **network**：在 README 中补充网络管理说明，并更新开发规则。

---

## [0.2.0] - 2026-04-17

### English

#### 🚀 Features

- Add DescribeOfficeSites to ADVANCED network activation flow with DNS default value support
- 调整 DEFAULT 网络激活流程，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

#### 📖 Documentation

- Improve README with API key management and image activate examples

### 中文

#### 🚀 功能

- **network**
  - ADVANCED 网络激活流程新增 DescribeOfficeSites，并支持 DNS 默认值
  - 调整 DEFAULT 网络激活流程，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

#### 📖 文档

- **全局**：改进 README，补充 API Key 管理和 `image activate` 示例。

---

## [0.1.9] - 2026-04-10

### English

#### 🚀 Features

- Add CreateApiKey CLI command
- Add API key concurrency management CLI command
- Validate add and copy file size

#### 🐞 Bug Fixes

- Correct CreateApiKey response Data field type from object to string
- Sync mock implementations after adding Client interface methods

#### 📖 Documentation

- Add Qoder rules and skills for CLI development

#### 📦 Other Changes

- Add 1 MiB COPY/ADD source limit and retries for OSS uploads and skill push APIs

### 中文

#### 🚀 功能

- **apikey**
  - 新增 `apikey create` 命令
  - 新增 API Key 并发管理命令
- **docker/image/skills**：为 OSS 上传和 skill push 增加 COPY/ADD 文件大小校验。

#### 🐞 缺陷修复

- **apikey**：修正 CreateApiKey 响应 Data 字段类型，从对象改为字符串。
- **测试**：新增 Client 接口方法后同步 mock 实现，修复编译问题。

#### 📖 文档

- **qoder**：新增 CLI 开发规则与 skills。

#### 📦 其他变更

- **docker/image/skills**：为 OSS 上传和 skill push API 增加 1 MiB COPY/ADD 源文件限制和重试机制。

---

## [0.1.8] - 2026-03-31

### English

#### 📦 Other Changes

- Update Makefile
- Update Makefile

### 中文

#### 📦 其他变更

- **构建**：更新 Makefile。

---

## [0.1.3] - 2026-02-10

### English

#### 🚀 Features

- Sync latest changes from internal repo

### 中文

#### 🚀 功能

- **全局**：同步内部仓库最新变更。

---

## [0.1.2] - 2025-12-26

### English

#### 🐞 Bug Fixes

- Endpoint to xiaoying.cn-shanghai

### 中文

#### 🐞 缺陷修复

- **core**：将默认 Endpoint 调整为 xiaoying.cn-shanghai。

---

## [0.1.1] - 2025-12-25

### English

#### 🚀 Features

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

#### 📖 Documentation

- Update installation and usage guides
- Add Linux & Mac installation guide and update docs structure

#### 🛠 Refactoring

- Add fixossendpoint to suitable return error ossendpoint

#### 📦 Other Changes

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

### 中文

#### 🚀 功能

- **image**
  - 增强 image list，支持系统镜像展示及 `--include-system` / `--system-only` 参数
  - 新增 Dockerfile demo 与 init Dockerfile template
  - 新增 image init 用户指南
- **core/auth**
  - 登录命令增加端口可用性检测与重试
  - 端口被占用时给出使用建议
- **安装**：新增 Windows 安装脚本和 Linux/Mac 安装指南。
- **CI/CD**：更新 workflow 配置。

#### 🐞 缺陷修复

- **image**：修复 image validation 单元测试错误。
- **OSS**：调整 fixossendpoint 以返回合适的 oss endpoint 错误信息。

#### 📖 文档

- **安装/使用**：更新安装与使用指南，并调整文档结构。

#### 📦 其他变更

- **image**：补充 image list、getdockerfile、image init、镜像修改信息展示等内部实现更新。
- **core/auth**：增加端口退避策略。
- **测试/CI**：修复测试错误信息大小写问题并更新 homebrew workflow。

---

## [0.1.0] - 2025-10-29

### English

#### 🚀 Features

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

#### 🐞 Bug Fixes

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

#### 📖 Documentation

- Update environment variable names in windows script
- Add Skills CLI usage (requirement a) to USER_GUIDE
- **plans**: Add Skills API BodyType analysis and backend format reference
- Align README and USER_GUIDE with skills CLI implementation
- Add skills output examples to USER_GUIDE, add manual test results

#### 🛠 Refactoring

- Add fixossendpoint to suitable return error ossendpoint

#### 📦 Other Changes

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

### 中文

#### 🚀 功能

- **core/auth**
  - 新增 refresh token 与环境切换能力
  - 支持国际站（prod/pre）默认 Endpoint 与 OAuth
  - 改进认证错误识别和处理
- **image**
  - 新增系统镜像列表展示及 `--include-system` / `--system-only` 参数
  - 新增 Dockerfile demo、init Dockerfile template 和 image init 指南
  - 新增 image status 命令，create image verbose 模式输出 RequestId
  - 新增 Dockerfile COPY/ADD 文件上传能力
  - 支持 sourceImageId 必填、默认环境/地域映射、镜像初始化信息和修改信息展示
- **skills**
  - 实现 skills push/list/show 命令
  - 实现 skills group create/list/show/add-skill/remove-skill 命令
  - 注册 skills 命令并完善测试 mock
- **client/agentbay**
  - 新增 CreateMarketSkill、DescribeMarketSkillDetail API
  - 新增 Market SkillGroup 系列 API（create/list/add-skill/remove-skill）
  - 扩展 agentbay client 支持以上 API
- **安装**：新增 Windows 安装脚本。
- **CI/CD**：新增 CI/CD 配置、发布流程与成功页。

#### 🐞 缺陷修复

- **skills**
  - 修复 ListMarketGroupSkill 预发环境 403 UnsupportedHTTPMethod 问题
  - 修复 CreateMarketSkillGroup Data 字段解析
  - 修复 DescribeMarketSkillDetail / AddMarketGroupSkill / RemoveMarketGroupSkill XML/JSON 解析
- **skills/image**：修复 OSS 上传和 API 响应处理问题。
- **image**
  - 适配 DeleteResourceGroup 接口，修复 deactivate 停止镜像失败问题
  - 修复 image init 解析与单元测试问题
- **core/auth**
  - 修复 logout warning、OAuth 登录解析、refresh token client 复用
  - 修复 AK/SK skill push 解析问题
- **client**
  - 对齐 Skills API HTTP 方法（GET/POST）
  - 修复编译与响应解析相关问题
- **CI/CD**：更新 Aone Makefile 以兼容最新 glibc。

#### 📖 文档

- **skills**
  - 新增 Skills CLI 使用说明与输出示例
  - 补充手动测试结果，将 README/USER_GUIDE 与 skills CLI 实现对齐
- **内部参考**：新增 Skills API BodyType 分析和后端格式参考。
- **环境配置**：更新 Windows 脚本中的环境变量名称。

#### 🛠 重构

- **OSS**：调整 fixossendpoint，以返回更合适的 oss endpoint 错误信息。

#### 📦 其他变更

- **image**
  - 整理 image list、getdockerfile、image init 历史实现更新
  - 补充 sourceImageId 必填与默认 Endpoint 更新
  - 展示原始错误信息以替代自定义错误文案
- **CI/CD**：重构 GitHub Homebrew workflow。
- **发版**：发布 v0.1.0。

---

<!-- generated by git-cliff -->
