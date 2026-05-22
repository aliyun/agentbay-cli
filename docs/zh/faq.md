[English](../en/faq.md) | **中文**

# 常见问题

**Q: 如何查看帮助？**

```bash
agentbay --help
agentbay image --help
```

**Q: 如何查看版本？**

```bash
agentbay version
```

**Q: 如何启用详细日志？**

```bash
agentbay -v image list
agentbay -v skills push ./my-skill
```

**Q: 登录遇到问题？**

- 检查网络连接
- 确保浏览器能访问 `signin.aliyun.com`（国际站为 `signin.alibabacloud.com`）
- 检查防火墙设置
- 非交互式环境建议使用 `AGENTBAY_ACCESS_KEY_ID` 和 `AGENTBAY_ACCESS_KEY_SECRET` 替代 `agentbay login`

**Q: 镜像构建失败？**

- 检查 Dockerfile 语法
- 确认基础镜像 ID 有效（使用 `agentbay image list --include-system` 查看可用的系统镜像 ID）
- 检查是否修改了 Dockerfile 前 N 行（N 在下载模板时提示）
- 使用 `agentbay image init -i <base-image-id>` 下载模板 Dockerfile
- 使用 `-v` 选项查看详细错误信息

**Q: Dockerfile 哪些部分不能修改？**

- 通过 `agentbay image init -i <image-id>` 下载的 Dockerfile 前 N 行是系统定义的，不可修改
- 仅可编辑第 N+1 行之后的内容，否则镜像构建可能失败

**Q: 配置文件存储在哪里？**

- `~/.config/agentbay/config.json`（macOS/Linux）或 `%APPDATA%\agentbay\config.json`（Windows）
- OAuth Token 存储在配置文件中；AccessKey 凭证**不会**被 CLI 保存，仅从环境变量读取

**Q: 支持哪些 OS 类型？**

Linux、Windows、Android

**Q: CI/CD 中如何跳过确认提示？**

在破坏性命令上使用 `--yes` / `-y`：

```bash
agentbay apikey delete akm-xxx --yes
agentbay image delete imgc-xxx --yes
```
