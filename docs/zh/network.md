[English](../../en/network.md) | **中文**

# 网络管理 — `agentbay network`

按区域查询网络包及其 EIP 绑定信息。

## 命令

### `network package list`

按区域列出网络包。

```bash
agentbay network package list                              # 默认区域 cn-hangzhou
agentbay network package list --biz-region-id cn-shanghai  # 指定区域
```

**参数：**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `--biz-region-id` | string | 否 | 区域 ID（默认：`cn-hangzhou`） |
