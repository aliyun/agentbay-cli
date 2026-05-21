[中文](../../zh/faq.md) | **English**

# FAQ

**Q: How to view help?**

```bash
agentbay --help
agentbay image --help
```

**Q: Check CLI version?**

```bash
agentbay version
```

**Q: Enable detailed logs?**

```bash
agentbay -v image list
agentbay -v skills push ./my-skill
```

**Q: Login issues?**

- Check network connection
- Ensure browser can access `signin.aliyun.com` (or the international sign-in host if using `AGENTBAY_ENV=international`)
- Check firewall settings
- For non-interactive use, set `AGENTBAY_ACCESS_KEY_ID` and `AGENTBAY_ACCESS_KEY_SECRET` instead of `agentbay login`

**Q: Image build fails?**

- Verify Dockerfile syntax
- Check base image ID is valid (use `agentbay image list --include-system` to find valid system image IDs)
- Check if you modified the first N lines of the Dockerfile (N is shown when downloading the template)
- Use `agentbay image init -i <base-image-id>` to download a template Dockerfile
- Use `-v` option to view detailed error information

**Q: Which parts of the Dockerfile cannot be modified?**

- The first N lines of the Dockerfile template downloaded by `agentbay image init -i <image-id>` are system-defined and cannot be modified
- Only modify content after line N+1, otherwise the image build may fail

**Q: Where is config stored?**

- `~/.config/agentbay/config.json` (macOS/Linux) or `%APPDATA%\agentbay\config.json` (Windows)
- OAuth tokens are stored there; AccessKey credentials are **not** saved by the CLI — only read from environment variables

**Q: Supported OS types?**

Linux, Windows, Android

**Q: How to skip confirmation prompts in CI/CD?**

Use `--yes` / `-y` on destructive commands:

```bash
agentbay apikey delete akm-xxx --yes
agentbay image delete imgc-xxx --yes
```
