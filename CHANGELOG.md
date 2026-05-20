# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
## [Unreleased]

### Features

- Add image warmup-status command

### Documentation

- Add CLI command to OpenAPI action mapping reference

## [0.2.8] - 2026-05-16

### Features

- Support sudo docker

## [0.2.7] - 2026-05-15

### Features

- Add image set-max-session command for configuring max concurrent sessions
- Support specifying region for image activate

### Bug Fixes

- **client**: Tolerate stringified HttpStatusCode in BatchCreateHideResourceGroupsWithMaxSession response

### Documentation

- **qoder**: Codify response-parsing fault tolerance rules across rules/skill/references
- Clarify image type support for set-max-session command
- Add whitelist requirement for set-max-session command
- Add build verification rule and image activate region spec

## [0.2.6] - 2026-05-14

### Features

- Add docker login and image create-from-template commands
- Add source-image validation and path truncation for create-from-template
- Add UsePublicNetwork param to GetDockerfileTemplate and print RequestId in image init

## [0.2.5] - 2026-05-11

### Features

- **image**: Add 'agentbay image delete' command
- Add release-to-oss script and Makefile targets

## [0.2.4] - 2026-05-08

### Features

- Enhance image deactivate with RequestId logging and precise ListMcpImages query

## [0.2.3] - 2026-04-29

### Features

- Enhance image create file upload with per-file status, auto-retry and summary

## [0.2.2] - 2026-04-28

### Features

- Add sandbox lifecycle parameters to image activate command

## [0.2.1] - 2026-04-23

### Features

- Add network package list command (DescribeNetworkPackages)
- **network**: Remove UserAliUid param, add default BizRegionId and OfficeSiteId display

### Documentation

- Add network management to README and update development rules

## [0.2.0] - 2026-04-17

### Features

- Add DescribeOfficeSites to ADVANCED network activation flow with DNS default value support
- 调整 DEFAULT 网络激活流程，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

### Documentation

- Improve README with API key management and image activate examples

## [0.1.9] - 2026-04-10

### Features

- Add CreateApiKey CLI command
- Add API key concurrency management CLI command
- Validate add and copy file size

### Bug Fixes

- Correct CreateApiKey response Data field type from object to string
- Sync mock implementations after adding Client interface methods

### Documentation

- Add Qoder rules and skills for CLI development

### Other Changes

- Add 1 MiB COPY/ADD source limit and retries for OSS uploads and skill push APIs

## [0.1.8] - 2026-03-31

### Other Changes

- Update Makefile
- Update Makefile

## [0.1.3] - 2026-02-10

### Features

- Sync latest changes from internal repo

## [0.1.2] - 2025-12-26

### Bug Fixes

- Endpoint to xiaoying.cn-shanghai

## [0.1.1] - 2025-12-25

### Features

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

### Documentation

- Update installation and usage guides
- Add Linux & Mac installation guide and update docs structure

### Refactoring

- Add fixossendpoint to suitable return error ossendpoint

### Other Changes

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

## [0.1.0] - 2025-10-29

### Features

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

### Bug Fixes

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

### Documentation

- Update environment variable names in windows script
- Add Skills CLI usage (requirement a) to USER_GUIDE
- **plans**: Add Skills API BodyType analysis and backend format reference
- Align README and USER_GUIDE with skills CLI implementation
- Add skills output examples to USER_GUIDE, add manual test results

### Refactoring

- Add fixossendpoint to suitable return error ossendpoint

### Other Changes

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

<!-- generated by git-cliff -->
