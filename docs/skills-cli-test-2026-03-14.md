# Skills CLI 手动测试结果 (2026-03-14)

## 测通接口

| 命令 | 说明 |
|------|------|
| `skills push <dir>` | 打包上传并创建 skill，返回 Skill ID |
| `skills show <skill-id>` | 查询 skill 详情（SkillId / Name / Description） |
| `skills group create <name>` | 创建技能组，返回 Group ID |
| `skills group list` | 列出当前用户技能组 |
| `skills group add-skill <group-id> <skill-id>` | 组内添加技能 |
| `skills group remove-skill <group-id> <skill-id>` | 组内移除技能 |

## 当前数据（测试时）

- **Skills**：如 `35U2Ver2`（name: xlsx）、`8dKWvDK3`（name: xlsx）等。
- **Groups**：6 个，包括 `NddfVFfd`(default)、`mKnRYMyM`/`J2ce2caN`/`f2AbaWwB`/`2ustVwmL`/`k4RS99VC`(test_skill_group) 等。

## 未测（占位）

- `skills list`、`skills group show` 后端 API 未就绪，仅占位输出。
