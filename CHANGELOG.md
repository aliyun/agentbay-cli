# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.0] - 2026-06-01

### English

#### 🚀 Features

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

- skills 命令组增强：`skills show` 展示 TenantTags，`skills push` 支持 `--tag` 与 `--icon`，新增 `skills update`、`skills delete`，并支持 `skills update --clear-tags` 与 `skills delete` 位置参数。
- 列表输出增强：为 skills/image/apikey list 增加 `--output json`，并优化 skills 与 image list 的终端自适应表格展示。
- 镜像与 docker 能力扩展：`image create-from-template` 支持短 source-image 路径，新增 docker repo share/unshare/list 与 docker shared repos list，并补充相关文档与 RAM 权限说明。
- 增加 skills-tester agent、分页测试规则，以及双语 CHANGELOG 发版流程。

#### 🐞 Bug Fixes

- OAuth 登录拒绝 RAM 身份，并优化 RAM OAuth 登录引导。
- Market skill detail 接口统一使用 JSON Accept 请求头。

#### 📖 Documentation

- 重构 README.md 与 README.zh-CN.md，提升结构清晰度和完整性。
- 对文档中的 AliUID 做脱敏处理，并补充敏感信息脱敏规则。
- 更新 image workflow 文档与 RAM 权限说明，并整理 RAM permissions 章节。

---

## [Unreleased]

## [0.3.0] - 2026-05-22

### 🚀 Features

- Enhance apikey commands with --api-key-id param and update CHANGELOG
- Add 'apikey describe-key-content' command

### 📖 Documentation

- Reorganize documentation into en/zh structure
- Soften 'agentbay login' wording from deprecated to not recommended
- Fix bilingual switch links and add docs governance spec
- Update CHANGELOG.md for v0.2.10
- Enrich docker login & image create-from-template usage notes
- Add RAM permission requirements for apikey commands
- Extend RAM permission docs to all command groups
- Backfill Chinese translations for all historical CHANGELOG versions
- Add update and uninstall instructions to install guides

---

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

---

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

---

## [0.2.8] - 2026-05-16

### 🚀 Features

- Support sudo docker

---

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

---

## [0.2.6] - 2026-05-14

### 🚀 Features

- Add docker login and image create-from-template commands
- Add source-image validation and path truncation for create-from-template
- Add UsePublicNetwork param to GetDockerfileTemplate and print RequestId in image init

---

## [0.2.5] - 2026-05-11

### 🚀 Features

- **image**: Add 'agentbay image delete' command
- Add release-to-oss script and Makefile targets

---

## [0.2.4] - 2026-05-08

### 🚀 Features

- Enhance image deactivate with RequestId logging and precise ListMcpImages query

---

## [0.2.3] - 2026-04-29

### 🚀 Features

- Enhance image create file upload with per-file status, auto-retry and summary

---

## [0.2.2] - 2026-04-28

### 🚀 Features

- Add sandbox lifecycle parameters to image activate command

---

## [0.2.1] - 2026-04-23

### 🚀 Features

- Add network package list command (DescribeNetworkPackages)
- **network**: Remove UserAliUid param, add default BizRegionId and OfficeSiteId display

### 📖 Documentation

- Add network management to README and update development rules

---

## [0.2.0] - 2026-04-17

### 🚀 Features

- Add DescribeOfficeSites to ADVANCED network activation flow with DNS default value support
- 调整 DEFAULT 网络激活流程，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

### 📖 Documentation

- Improve README with API key management and image activate examples

---

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

---

## [0.1.8] - 2026-03-31

### 📦 Other Changes

- Update Makefile
- Update Makefile

---

## [0.1.3] - 2026-02-10

### 🚀 Features

- Sync latest changes from internal repo

---

## [0.1.2] - 2025-12-26

### 🐞 Bug Fixes

- Endpoint to xiaoying.cn-shanghai

---

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

---

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

---

<!-- generated by git-cliff -->
