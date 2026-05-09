# Change Records (CR) — 变更档案索引

本目录用于**手动变更管理(路径 B)**,与 Qoder Quest Mode 的 `.qoder/quest/`(路径 A)互补,共同承担需求追溯链责任。

> 完整流程规则见: [.qoder/skills/feature-development-workflow/SKILL.md](../skills/feature-development-workflow/SKILL.md)
> 代码规范见: [.qoder/rules/development.md](../rules/development.md)

## 📒 什么时候用本目录?

- 未启用 Qoder Quest Mode 的轻量需求
- 有外部设计稿 / 需求文档需要落地到代码仓追溯
- 紧急修复、单点调整等不适合启动 Quest 任务的场景

否则**优先使用 Quest Mode**,让 Qoder 自动生成 Spec 归档到 `.qoder/quest/`。

## 🗂️ 目录命名规范

每个需求一个独立的 CR 目录,命名格式:

```
CR-<YYYY-MM-DD>-<feature-name>/
```

- `<YYYY-MM-DD>`:需求启动日期(ISO 格式)
- `<feature-name>`:与 feat 分支名保持一致的短横线关键词

**示例**:

- `CR-2026-05-09-image-delete/` ↔ feat 分支 `feat-image-delete`
- `CR-2026-06-01-apikey-concurrency/` ↔ feat 分支 `feat-apikey-concurrency`

## 📄 CR 目录内部结构

参考 [TEMPLATE.md](./TEMPLATE.md) 建立以下文件(按需裁剪,但 `spec.md` / `trace.md` 必有):

| 文件           | 作用                                        | 必选   |
| -------------- | ------------------------------------------- | ------ |
| `spec.md`      | 需求规格:背景 / 目标 / 非目标 / 范围        | ✅     |
| `design.md`    | 技术设计:接口 / 流程图 / 状态机             | 推荐   |
| `tasks.md`     | 任务分解,映射到 `todo_write` tasklist       | 推荐   |
| `decisions.md` | 关键决策记录(含对话沉淀)                    | 推荐   |
| `test-plan.md` | 测试计划:单测 / 集成 / 回归                 | 推荐   |
| `rollback.md`  | 回滚预案                                    | 视风险 |
| `trace.md`     | 追溯链:分支 / commits / push / PR / release | ✅     |

## 🔁 生命周期

```
启动需求 → 建 CR 目录 + 填 spec.md → 拉 feat 分支
       → 开发中持续更新 design/tasks/decisions
       → 每次 commit 后登记 trace.md
       → 每次 push 后登记 trace.md
       → PR 合并后登记 PR 链接 / 合并 commit / release tag
       → CR 目录永久保留,作为历史追溯资产(禁止删除)
```

## 📚 已归档 CR 列表

> 新增 CR 后请在此追加一行,便于检索。

| CR ID    | 需求 | feat 分支 | 状态 | PR  |
| -------- | ---- | --------- | ---- | --- |
| _(尚无)_ | —    | —         | —    | —   |
