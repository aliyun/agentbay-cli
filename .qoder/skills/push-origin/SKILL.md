---
name: push-origin
description: 将当前本地分支推送到 origin 远程的同名分支（禁止推送到 master/main）
---

# Push to Origin

## 📋 职责

将当前本地分支**一键推送**到 `origin` 远程的**同名分支**。

适用于日常将工作内容归档到内网 GitLab 仓库（`git@gitlab.alibaba-inc.com:InnoArchClub/agentbay-cli.git`）。

---

## 🚀 执行步骤

### Step 1：获取当前分支名

```bash
git branch --show-current
```

### Step 2：安全检查

**禁止推送到以下分支**（遇到时必须中止并告知用户）：

- `master`
- `main`

如果当前分支是 `master` 或 `main`，**立即停止**，输出：

```
[ERROR] 当前分支是 master/main，禁止直接推送到远程主分支。
请切换到 feat 分支后再执行推送。
```

### Step 3：检查并暂存改动

```bash
git status --short
```

- 若有**未暂存/未追踪**的文件，执行：

```bash
git add .
```

- 若工作区干净（无改动），跳过 Step 3 和 Step 4，直接执行 Step 5 推送。

### Step 4：自动生成 commit 信息并提交

先分析改动内容：

```bash
git diff --cached --name-status
```

根据改动文件和内容，按 **Conventional Commits** 规范自动生成 commit message：

- 新增功能 → `feat: <描述>`
- 修复问题 → `fix: <描述>`
- 文档变更 → `docs: <描述>`
- 测试相关 → `test: <描述>`
- 重构代码 → `refactor: <描述>`
- 样式/格式 → `style: <描述>`
- 构建/工具 → `chore: <描述>`

描述须**简洁准确**，反映本次改动的核心内容（英文或中文均可，与项目既有风格保持一致）。

```bash
git commit -m "<自动生成的 commit message>"
```

### Step 5：执行推送

```bash
BRANCH=$(git branch --show-current)
git push origin "$BRANCH"
```

### Step 6：确认结果

推送成功后，输出提示：

```
[INFO] 已成功推送到 origin/$BRANCH
```

---

## ⚠️ 注意事项

- 不使用 `--force` / `-f`，除非用户**明确指示**
- 不使用 `--no-verify` 跳过 hook
- 推送目标仅为 `origin`，不操作 `aliyun`
- 若推送失败（如远程分支不存在），自动加 `--set-upstream` 重试：

  ```bash
  git push --set-upstream origin "$BRANCH"
  ```
