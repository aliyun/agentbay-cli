[English](../en/ram-permissions.md) | **中文**

# RAM 账号接口权限汇总

> 阿里云**主账号**无需任何额外权限配置。
> 本节仅适用于使用 AK/SK 认证的 **RAM 子账号**。

如果使用 RAM 子账号的 AK/SK，可以到 [RAM 控制台](https://ram.console.aliyun.com/policies) 先新建或修改权限策略，再将策略授权给对应的 RAM 子账号。

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
