# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
## [Unreleased]

### Features

- Add AvailableInstanceSize field to warmup-status image output

### Documentation

- Update CHANGELOG.md for v0.2.9

## [0.2.9] - 2026-05-20

### Features

- Add image warmup-status command
- Integrate git-cliff for automated changelog generation

### Bug Fixes

- Correct git-cliff download URL in homebrew workflow

### Documentation

- Add CLI command to OpenAPI action mapping reference

* * *

### 新功能

- 新增 image warmup-status 命令
- 集成 git-cliff 自动生成 CHANGELOG

### 问题修复

- 修正 homebrew 工作流中 git-cliff 的下载 URL

### 文档

- 新增 CLI 命令与 OpenAPI Action 的映射对照表

## [0.2.8] - 2026-05-16

### Features

- Support sudo docker

* * *

### 新功能

- 支持 sudo docker

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

* * *

### 新功能

- 新增 `image set-max-session` 命令，用于配置最大并发会话数
- `image activate` 支持通过 `--region-id` 指定区域

### 问题修复

- **client**: 兼容 `BatchCreateHideResourceGroupsWithMaxSession` 响应中 `HttpStatusCode` 字段为字符串类型的情况

### 文档

- **qoder**: 将响应解析容错规则写入 rules/skill/references
- 说明 `set-max-session` 命令支持的镜像类型
- 补充 `set-max-session` 命令的白名单要求
- 新增构建验证规则及 `image activate` 区域参数说明

## [0.2.6] - 2026-05-14

### Features

- Add docker login and image create-from-template commands
- Add source-image validation and path truncation for create-from-template
- Add UsePublicNetwork param to GetDockerfileTemplate and print RequestId in image init

* * *

### 新功能

- 新增 `docker login` 和 `image create-from-template` 命令
- `create-from-template` 增加 source-image 校验和路径截断
- `GetDockerfileTemplate` 新增 `UsePublicNetwork` 参数，`image init` 打印 RequestId

## [0.2.5] - 2026-05-11

### Features

- **image**: Add 'agentbay image delete' command
- Add release-to-oss script and Makefile targets

* * *

### 新功能

- **image**: 新增 `agentbay image delete` 命令
- 新增 release-to-oss 脚本及 Makefile 目标

## [0.2.4] - 2026-05-08

### Features

- Enhance image deactivate with RequestId logging and precise ListMcpImages query

* * *

### 新功能

- 增强 `image deactivate`：打印 RequestId 并精确查询 ListMcpImages

## [0.2.3] - 2026-04-29

### Features

- Enhance image create file upload with per-file status, auto-retry and summary

* * *

### 新功能

- 增强 `image create` 文件上传：逐文件状态、自动重试和汇总

## [0.2.2] - 2026-04-28

### Features

- Add sandbox lifecycle parameters to image activate command

* * *

### 新功能

- `image activate` 命令新增沙箱生命周期参数

## [0.2.1] - 2026-04-23

### Features

- Add network package list command (DescribeNetworkPackages)
- **network**: Remove UserAliUid param, add default BizRegionId and OfficeSiteId display

### Documentation

- Add network management to README and update development rules

* * *

### 新功能

- 新增 `network package list` 命令（DescribeNetworkPackages）
- **network**: 移除 UserAliUid 参数，增加默认 BizRegionId 和 OfficeSiteId 显示

### 文档

- README 新增网络管理章节，更新开发规则

## [0.2.0] - 2026-04-17

### Features

- Add DescribeOfficeSites to ADVANCED network activation flow with DNS default value support
- 调整 DEFAULT 网络激活流程，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

### Documentation

- Improve README with API key management and image activate examples

* * *

### 新功能

- ADVANCED 网络激活流程增加 DescribeOfficeSites，DNS 支持默认值
- DEFAULT 网络激活流程调整，增加 DescribeMcpPolicyData 和 SaveMcpPolicyData 调用

### 文档

- README 补充 API Key 管理和 image activate 示例

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

* * *

### 新功能

- 新增 `CreateApiKey` CLI 命令
- 新增 API Key 并发管理 CLI 命令
- COPY/ADD 文件大小校验

### 问题修复

- 修正 `CreateApiKey` 响应 Data 字段类型（从对象改为字符串）
- 新增 Client 接口方法后同步更新所有 mock 实现

### 文档

- 新增 Qoder 开发规则和 Skills

### 其他变更

- COPY/ADD 源文件限制 1 MiB，OSS 上传和 skill push 增加重试

## [0.1.8] - 2026-03-31

### Other Changes

- Update Makefile
- Update Makefile

* * *

### 其他变更

- 更新 Makefile（两次）

## [0.1.3] - 2026-02-10

### Features

- Sync latest changes from internal repo

* * *

### 新功能

- 从内部仓库同步最新变更

## [0.1.2] - 2025-12-26

### Bug Fixes

- Endpoint to xiaoying.cn-shanghai

* * *

### 问题修复

- Endpoint 修正为 xiaoying.cn-shanghai

## [0.1.1] - 2025-12-25

* * *

本次更新为初始功能版本，主要包括：镜像管理全流程（list / init / create / activate / deactivate / status）、Skills CLI（push / show）、Dockerfile COPY/ADD 文件上传、OAuth 登录、Windows 安装支持、系统镜像列表、环境切换功能、端口冲突处理等。

## [0.1.0] - 2025-10-29

* * *

首个发布版本，包含基础的 CLI 框架、OAuth 认证、镜像管理（create / list / activate / deactivate）、Skills 推送、API 层 XML/JSON 双格式响应解析等核心功能。

<!-- generated by git-cliff -->
