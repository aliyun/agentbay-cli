[English](../en/ram-permissions.md) | **中文**

# RAM 账号接口权限汇总

> 阿里云**主账号**无需任何额外权限配置。
> 本节仅适用于使用 AK/SK 认证的 **RAM 子账号**。

如果使用 RAM 子账号的 AK/SK，可以到 [RAM 控制台](https://ram.console.aliyun.com/policies) 先新建或修改权限策略，再将策略授权给对应的 RAM 子账号。

---

## `apikey` 命令分组

| OpenAPI Action              | 所需权限                             | 调用命令                                                                                                                                           |
| --------------------------- | ------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CreateApiKey`              | `agentbay:CreateApiKey`              | `apikey create`                                                                                                                                    |
| `DescribeApiKeys`           | `agentbay:DescribeApiKeys`           | `apikey list`、`apikey delete`（使用 `--api-key-id` 时）                                                                                            |
| `DescribeMcpApiKey`         | `agentbay:DescribeMcpApiKey`         | `apikey list`、`apikey delete`、`apikey enable`、`apikey disable`、`apikey concurrency set`（使用 `--api-key` 时）                                  |
| `ModifyMcpApiKeyConfig`     | `agentbay:ModifyMcpApiKeyConfig`     | `apikey concurrency set`                                                                                                                           |
| `ModifyApiKeyStatus`        | `agentbay:ModifyApiKeyStatus`        | `apikey enable`、`apikey disable`、`apikey delete`（删除 ENABLED 状态 API Key 时会先禁用）                                                          |
| `DeleteApiKey`              | `agentbay:DeleteApiKey`              | `apikey delete`                                                                                                                                    |
| `DescribeKeyContent`        | `agentbay:DescribeKeyContent`        | `apikey describe-key-content`                                                                                                                      |

**RAM Policy 示例：**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:CreateApiKey",
        "agentbay:DescribeApiKeys",
        "agentbay:DescribeMcpApiKey",
        "agentbay:ModifyMcpApiKeyConfig",
        "agentbay:ModifyApiKeyStatus",
        "agentbay:DeleteApiKey",
        "agentbay:DescribeKeyContent"
      ],
      "Resource": "*"
    }
  ]
}
```

---

## `core` 命令分组

`login`、`logout` 和 `version` 不直接调用 AgentBay OpenAPI 接口，无需配置额外的 RAM 权限。

---

## `docker` 命令分组

| OpenAPI Action          | 所需权限                         | 调用命令             |
| ----------------------- | -------------------------------- | -------------------- |
| `GetACRRepoCredential`  | `agentbay:GetACRRepoCredential`  | `docker login`       |
| `ShareDockerRepo`       | `agentbay:ShareDockerRepo`       | `docker share`       |
| `UnshareDockerRepo`     | `agentbay:UnshareDockerRepo`     | `docker unshare`     |
| `ListSharedDockerRepos` | `agentbay:ListSharedDockerRepos` | `docker list-shares` |

**RAM Policy 示例：**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:GetACRRepoCredential",
        "agentbay:ShareDockerRepo",
        "agentbay:UnshareDockerRepo",
        "agentbay:ListSharedDockerRepos"
      ],
      "Resource": "*"
    }
  ]
}
```

> `docker tag` 和 `docker push` 是对原生 `docker` CLI 的封装，不直接调用任何 AgentBay API，无需额外 RAM 权限。

---

## `image` 命令分组

| OpenAPI Action                                | 所需权限                                               | 调用命令                                                                                                      |
| --------------------------------------------- | ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------- |
| `ListMcpImages`                               | `agentbay:ListMcpImages`                               | `image list`、`image deactivate`                                                                              |
| `GetMcpImageInfo`                             | `agentbay:GetMcpImageInfo`                             | `image create`、`image activate`、`image deactivate`、`image delete`、`image status`、`image set-max-session` |
| `GetDockerFileStoreCredential`                | `agentbay:GetDockerFileStoreCredential`                | `image create`                                                                                                |
| `CreateDockerImageTask`                       | `agentbay:CreateDockerImageTask`                       | `image create`                                                                                                |
| `GetDockerImageTask`                          | `agentbay:GetDockerImageTask`                          | `image create`                                                                                                |
| `CreateImageFromTemplate`                     | `agentbay:CreateImageFromTemplate`                     | `image create-from-template`                                                                                  |
| `DescribeInstanceTypes`                       | `agentbay:DescribeInstanceTypes`                       | `image activate`                                                                                              |
| `DescribeMcpPolicyData`                       | `agentbay:DescribeMcpPolicyData`                       | `image activate`                                                                                              |
| `CreateMcpPolicyData`                         | `agentbay:CreateMcpPolicyData`                         | `image activate`                                                                                              |
| `ModifyMcpPolicyData`                         | `agentbay:ModifyMcpPolicyData`                         | `image activate`                                                                                              |
| `DescribeOfficeSites`                         | `agentbay:DescribeOfficeSites`                         | `image activate`                                                                                              |
| `SaveMcpPolicyData`                           | `agentbay:SaveMcpPolicyData`                           | `image activate`                                                                                              |
| `CreateResourceGroup`                         | `agentbay:CreateResourceGroup`                         | `image activate`                                                                                              |
| `DeleteResourceGroup`                         | `agentbay:DeleteResourceGroup`                         | `image deactivate`                                                                                            |
| `DeleteMcpImage`                              | `agentbay:DeleteMcpImage`                              | `image delete`                                                                                                |
| `GetDockerfileTemplate`                       | `agentbay:GetDockerfileTemplate`                       | `image init`                                                                                                  |
| `BatchCreateHideResourceGroupsWithMaxSession` | `agentbay:BatchCreateHideResourceGroupsWithMaxSession` | `image set-max-session`                                                                                       |
| `DescribeWarmUpStatusOpen`                    | `agentbay:DescribeWarmUpStatusOpen`                    | `image warmup-status`                                                                                         |

**RAM Policy 示例（`image` 命令完整授权）：**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:ListMcpImages",
        "agentbay:GetMcpImageInfo",
        "agentbay:GetDockerFileStoreCredential",
        "agentbay:CreateDockerImageTask",
        "agentbay:GetDockerImageTask",
        "agentbay:CreateImageFromTemplate",
        "agentbay:DescribeInstanceTypes",
        "agentbay:DescribeMcpPolicyData",
        "agentbay:CreateMcpPolicyData",
        "agentbay:ModifyMcpPolicyData",
        "agentbay:DescribeOfficeSites",
        "agentbay:SaveMcpPolicyData",
        "agentbay:CreateResourceGroup",
        "agentbay:DeleteResourceGroup",
        "agentbay:DeleteMcpImage",
        "agentbay:GetDockerfileTemplate",
        "agentbay:BatchCreateHideResourceGroupsWithMaxSession",
        "agentbay:DescribeWarmUpStatusOpen"
      ],
      "Resource": "*"
    }
  ]
}
```

> 如果只使用特定命令，请参考 [镜像文档](image.md) 中各命令的**涉及接口**章节，仅授予所需的最小权限。

---

## `network` 命令分组

| OpenAPI Action              | 所需权限                             | 调用命令               |
| --------------------------- | ------------------------------------ | ---------------------- |
| `DescribeNetworkPackages`   | `agentbay:DescribeNetworkPackages`   | `network package list` |

**RAM Policy 示例：**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:DescribeNetworkPackages"
      ],
      "Resource": "*"
    }
  ]
}
```

---

## `skills` 命令分组

| OpenAPI Action                 | 所需权限                                | 调用命令                                      |
| ------------------------------ | --------------------------------------- | --------------------------------------------- |
| `ListTag`                      | `agentbay:ListTag`                      | `skills push`、`skills update`（提供 `--tag` 时） |
| `CreateTag`                    | `agentbay:CreateTag`                    | `skills push`、`skills update`（提供新标签时）     |
| `GetMarketSkillCredential`     | `agentbay:GetMarketSkillCredential`     | `skills push`、`skills update`（`skills update` 需提供 `--file`） |
| `CreateMarketSkill`            | `agentbay:CreateMarketSkill`            | `skills push`                                  |
| `UpdateMarketSkill`            | `agentbay:UpdateMarketSkill`            | `skills update`                                |
| `ListMarketSkillByPage`        | `agentbay:ListMarketSkillByPage`        | `skills list`                                  |
| `DescribeMarketSkillDetail`    | `agentbay:DescribeMarketSkillDetail`    | `skills show`、`skills delete`（未提供 `--yes` 时） |
| `DeleteMarketSkill`            | `agentbay:DeleteMarketSkill`            | `skills delete`                                |

**RAM Policy 示例：**

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "agentbay:ListTag",
        "agentbay:CreateTag",
        "agentbay:GetMarketSkillCredential",
        "agentbay:CreateMarketSkill",
        "agentbay:UpdateMarketSkill",
        "agentbay:ListMarketSkillByPage",
        "agentbay:DescribeMarketSkillDetail",
        "agentbay:DeleteMarketSkill"
      ],
      "Resource": "*"
    }
  ]
}
```
