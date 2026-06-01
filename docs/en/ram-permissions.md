[中文](../zh/ram-permissions.md) | **English**

# RAM Permissions Summary

> The main Alibaba Cloud account does **not** require any additional permission configuration.
> This section applies only to **RAM sub-accounts** using AK/SK authentication.

If you are using a RAM sub-account's AK/SK, go to the [RAM console](https://ram.console.aliyun.com/policies) to create or modify a permission policy first, and then attach the policy to the corresponding RAM sub-account.

---

## `apikey` Command Group

| OpenAPI Action              | Required Permission                 | Used By                                                                                                                                                       |
| --------------------------- | ----------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CreateApiKey`              | `agentbay:CreateApiKey`             | `apikey create`                                                                                                                                               |
| `DescribeApiKeys`           | `agentbay:DescribeApiKeys`          | `apikey list`, `apikey delete` (when `--api-key-id` is used)                                                                                                   |
| `DescribeMcpApiKey`         | `agentbay:DescribeMcpApiKey`        | `apikey list`, `apikey delete`, `apikey enable`, `apikey disable`, `apikey concurrency set` (when `--api-key` is used)                                        |
| `ModifyMcpApiKeyConfig`     | `agentbay:ModifyMcpApiKeyConfig`    | `apikey concurrency set`                                                                                                                                      |
| `ModifyApiKeyStatus`        | `agentbay:ModifyApiKeyStatus`       | `apikey enable`, `apikey disable`, `apikey delete` (when deleting an ENABLED API key, the command disables it first)                                           |
| `DeleteApiKey`              | `agentbay:DeleteApiKey`             | `apikey delete`                                                                                                                                               |
| `DescribeKeyContent`        | `agentbay:DescribeKeyContent`       | `apikey describe-key-content`                                                                                                                                 |

**RAM Policy example:**

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

## `core` Command Group

`login`, `logout`, and `version` do not call AgentBay OpenAPI directly. No additional RAM permissions are required.

---

## `docker` Command Group

| OpenAPI Action          | Required Permission              | Used By              |
| ----------------------- | -------------------------------- | -------------------- |
| `GetACRRepoCredential`  | `agentbay:GetACRRepoCredential`  | `docker login`       |
| `ShareDockerRepo`       | `agentbay:ShareDockerRepo`       | `docker share`       |
| `UnshareDockerRepo`     | `agentbay:UnshareDockerRepo`     | `docker unshare`     |
| `ListSharedDockerRepos` | `agentbay:ListSharedDockerRepos` | `docker list-shares` |

**RAM Policy example:**

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

> `docker tag` and `docker push` are wrappers around the native `docker` CLI and do not call any AgentBay API directly. No additional RAM permissions are required.

---

## `image` Command Group

| OpenAPI Action                                | Required Permission                                    | Used By                                                                                                       |
| --------------------------------------------- | ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------- |
| `ListMcpImages`                               | `agentbay:ListMcpImages`                               | `image list`, `image deactivate`                                                                              |
| `GetMcpImageInfo`                             | `agentbay:GetMcpImageInfo`                             | `image create`, `image activate`, `image deactivate`, `image delete`, `image status`, `image set-max-session` |
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

**RAM Policy example (full access to `image` commands):**

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

> If you only use specific commands, refer to the **Involved APIs** section in [Image docs](image.md) and grant only the required subset.

---

## `network` Command Group

| OpenAPI Action              | Required Permission                 | Used By                |
| --------------------------- | ----------------------------------- | ---------------------- |
| `DescribeNetworkPackages`   | `agentbay:DescribeNetworkPackages`  | `network package list` |

**RAM Policy example:**

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

## `skills` Command Group

| OpenAPI Action                 | Required Permission                    | Used By                                           |
| ------------------------------ | -------------------------------------- | ------------------------------------------------- |
| `ListTag`                      | `agentbay:ListTag`                     | `skills push`, `skills update` (when `--tag` is provided) |
| `CreateTag`                    | `agentbay:CreateTag`                   | `skills push`, `skills update` (when new tags are provided) |
| `GetMarketSkillCredential`     | `agentbay:GetMarketSkillCredential`    | `skills push`, `skills update` (`skills update` requires `--file`) |
| `CreateMarketSkill`            | `agentbay:CreateMarketSkill`           | `skills push`                                     |
| `UpdateMarketSkill`            | `agentbay:UpdateMarketSkill`           | `skills update`                                   |
| `ListMarketSkillByPage`        | `agentbay:ListMarketSkillByPage`       | `skills list`                                     |
| `DescribeMarketSkillDetail`    | `agentbay:DescribeMarketSkillDetail`   | `skills show`, `skills delete` (when `--yes` is not provided) |
| `DeleteMarketSkill`            | `agentbay:DeleteMarketSkill`           | `skills delete`                                   |

**RAM Policy example:**

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
