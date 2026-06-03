---
name: push-aliyun
description: 将当前本地分支推送到 aliyun 远程的同名分支（禁止推送到 master/main）
---

# Push to Aliyun

## 📋 职责

将当前本地分支**一键推送**到 `aliyun` 远程的**同名分支**。

适用于将 feat 分支推送到对客 GitHub 仓库（`git@github.com:aliyun/agentbay-cli.git`），作为后续提 PR 的源分支。若 `aliyun` 远程尚未绑定，本 skill 会自动添加。

---

## 🚀 执行步骤

### Step 0：前置检查——确认 aliyun 远程已绑定

由于项目可能被其他人 clone，其本地仓库可能尚未绑定 `aliyun` 远程。需先检查并自动配置：

```bash
git remote get-url aliyun
```

- **若返回 `git@github.com:aliyun/agentbay-cli.git`**：远程已正确绑定，继续下一步。
- **若返回其他 URL**：说明 `aliyun` 已被用户用于其他仓库，**中止并提示**：

  ```
  [WARN] aliyun 远程已绑定到其他地址：<实际URL>
  请手动处理后再执行推送（避免覆盖您已有的 remote 配置）。
  ```

- **若报错 `No such remote`**：自动添加远程：

  ```bash
  git remote add aliyun git@github.com:aliyun/agentbay-cli.git
  ```

  输出提示：

  ```
  [INFO] 已自动添加 aliyun 远程 → git@github.com:aliyun/agentbay-cli.git
  ```

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

- 若工作区干净（无改动），跳过 Step 3 和 Step 4，直接执行 Step 5 同步与推送。

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

### Step 5：同步 aliyun/master

在推送前，先获取远程最新状态，并确保当前分支已包含 `aliyun/master` 的最新提交：

```bash
git fetch aliyun
```

检查当前分支是否已同步：

```bash
git merge-base --is-ancestor aliyun/master HEAD
```

- **退出码 0**：当前分支已包含 `aliyun/master` 的全部提交，无需额外操作，继续下一步。
- **退出码非 0**：当前分支落后于 `aliyun/master`，执行 merge：

  ```bash
  git merge aliyun/master --no-edit
  ```

  merge 完成后继续下一步。

> 若 merge 过程中产生冲突，skill 应**中止**并提示用户手动解决冲突后再重新执行推送。

### Step 6：执行推送

```bash
BRANCH=$(git branch --show-current)
git push aliyun "$BRANCH"
```

### Step 7：确认结果

推送成功后，输出提示：

```
[INFO] 已成功推送到 aliyun/$BRANCH
```

---

## ⚠️ 注意事项

- 不使用 `--force` / `-f`，除非用户**明确指示**
- 不使用 `--no-verify` 跳过 hook
- 推送目标仅为 `aliyun`，不操作 `origin`
- `git merge aliyun/master --no-edit` 使用默认 merge message，避免交互阻塞
- 若 merge 产生冲突，skill 应中止并提示用户手动解决
- 若推送失败（如远程分支不存在），自动加 `--set-upstream` 重试：

  ```bash
  git push --set-upstream aliyun "$BRANCH"
  ```
