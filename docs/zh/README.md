[English](../en/README.md) | **中文**

# AgentBay CLI 文档

欢迎使用 AgentBay CLI 文档。通过以下链接查找各命令组的详细信息。

## 快速开始

- [安装](installation.md) — 在 Windows、macOS 和 Linux 上安装 AgentBay CLI
- [认证与环境](authentication.md) — AccessKey、STS、OAuth 及环境配置

## 教程

- [镜像创建与共享完整流程](image-workflow.md) — 从 Dockerfile 模板到跨账号共享的端到端教程

## 命令参考

| 分组    | 命令                                  | 说明                                       | 详情                      |
| ------- | ------------------------------------- | ------------------------------------------ | ------------------------- |
| 核心    | `agentbay version`, `login`, `logout` | 版本信息与认证                             | [核心命令](core.md)       |
| 镜像    | `agentbay image ...`                  | 创建、列出、激活、停用、删除镜像等         | [镜像管理](image.md)      |
| API Key | `agentbay apikey ...`                 | 创建、列出、启用、禁用、删除密钥及设置并发 | [API Key 管理](apikey.md) |
| 网络    | `agentbay network ...`                | 查询网络包及 EIP 绑定信息                  | [网络管理](network.md)    |
| 技能    | `agentbay skills ...`                 | 推送与查看技能                             | [技能管理](skills.md)     |
| Docker  | `agentbay docker ...`                 | 登录、打 tag、推送镜像到 ACR               | [Docker 操作](docker.md)  |

## 权限配置

- [RAM 账号接口权限汇总](ram-permissions.md) — 仅 RAM 子账号需要配置

## 其他

- [常见问题](faq.md) — 常见问题解答
