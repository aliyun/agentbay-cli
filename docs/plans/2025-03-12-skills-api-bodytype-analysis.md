# Skills 相关接口 BodyType 与「后端返回 XML」问题分析

**使用技能：** systematic-debugging（根因调查 + 模式对比）  
**结论先行：** 与 ListMarketGroupSkill 同网关且使用 `BodyType: "json"` 的接口存在相同风险，建议按需改为 `"string"` + 手动解析或确认后端实际返回格式。

**接口与响应格式的规范来源：** `docs/plans/2025-03-12-skills-cli-implementation.md`（实现方案）。其中与响应/BodyType 相关的约定摘要如下，便于与当前实现及实际网关行为对照：

| 接口 | 实现方案中的描述 |
|------|------------------|
| GetMarketSkillCredential | Response 含 Body/Data、OSS 上传 URL 或 Bucket+Path+Credential，与 GetDockerFileStoreCredential 结构类似（当前实现用 XML） |
| CreateMarketSkill | Response 含 SkillId 等（当前实现用 XML） |
| DescribeMarketSkillDetail | 查询 Skill 详情（当前实现用 XML） |
| CreateMarketSkillGroup | Task 4：响应 BodyType **可能为 JSON**，需与现有 XML 解析方式区分或统一 |
| ListMarketGroupSkill | Task 4：list_market_group_skill_response_model **与后端 JSON 一致**（如 Data: []{GroupId, GroupName}）；若后端返回 JSON 使用 bodyType JSON。**实际运行中该接口返回了 XML**，已按 BodyType "string" + 手动解析修复 |
| AddMarketGroupSkill | Task 4：同上，响应 BodyType 可能为 JSON |
| RemoveMarketGroupSkill | Task 4：同上，响应 BodyType 可能为 JSON |

---

## 1. 问题模式回顾（ListMarketGroupSkill 根因）

- **现象：** `invalid character '<' looking for beginning of value`，且 RequestId 无法打印。
- **根因：**
  - 使用 `BodyType: "json"` 时，SDK 在 `CallApi` 内部会调用 `dara.ReadAsJSON(response_.Body)`。
  - 若后端返回的是 **XML**（body 以 `<` 开头），ReadAsJSON 直接报错并 `return (_result, _err)`，**不会**再组装含 body/headers 的 response map。
  - 因此：拿不到 `_body`、无法从响应中取 RequestId，且返回的是 SDK 原始错误而非我们包装的 `ErrWithRequestID`。
- **修复方式：** 将 ListMarketGroupSkill 改为 `BodyType: "string"`，在业务层用 `parseListMarketGroupSkillResponse` 根据首字符区分 XML/JSON 再解析，并在错误路径统一附带 RequestId。

---

## 2. 新增 Skills 相关接口一览

| 接口 | BodyType | Accept Header | 响应处理方式 | 与 ListMarketGroupSkill 同网关 |
|------|----------|----------------|--------------|--------------------------------|
| GetMarketSkillCredential   | **xml**  | application/xml  | `dara.Convert(_body, &_result)` | 是（同一产品/预发） |
| CreateMarketSkill          | **xml**  | application/xml  | `dara.Convert(_body, &_result)` | 是 |
| DescribeMarketSkillDetail  | **xml**  | application/xml  | `dara.Convert(_body, &_result)` | 是 |
| CreateMarketSkillGroup     | **json** | application/json | `dara.Convert(_body, &_result)` | 是 |
| ListMarketGroupSkill       | **string**（已修复） | application/json | `parseListMarketGroupSkillResponse` | 是 |
| AddMarketGroupSkill        | **json** | application/json | `dara.Convert(_body, &_result)` | 是 |
| RemoveMarketGroupSkill     | **json** | application/json | `dara.Convert(_body, &_result)` | 是 |

---

## 3. 风险分析（仅分析，不修改代码）

### 3.1 高风险：BodyType 为 `"json"` 的接口

若网关对以下接口实际返回 **XML**（与 ListMarketGroupSkill 行为一致），会复现同一类问题：

- **CreateMarketSkillGroup**（创建技能组）
- **AddMarketGroupSkill**（组内添加技能）
- **RemoveMarketGroupSkill**（组内移除技能）

**表现：**  
CallApi 内部 ReadAsJSON 失败 → 报 `invalid character '<'`，且无法从响应中取 RequestId，错误也无法包装为 `ErrWithRequestID`。

**建议（后续若需修改时）：**  
与 ListMarketGroupSkill 相同策略：改为 `BodyType: "string"`，在业务层根据 body 首字符做 XML/JSON 分支解析，并在错误路径从 map/headers/body 中提取 RequestId 并返回 `ErrWithRequestID`。

### 3.2 中/低风险：BodyType 为 `"xml"` 的接口

- **GetMarketSkillCredential**、**CreateMarketSkill**、**DescribeMarketSkillDetail** 使用 `BodyType: "xml"` 且 `Accept: application/xml`。
- 若后端**确实返回 XML**：与当前配置一致，无「JSON 解析 XML」问题。
- 若后端**某日改为只返回 JSON**：SDK 对 `BodyType: "xml"` 在 doRPCRequest_opResponse 中走 default 分支，**不会**把 body 放进 map（仅 headers + statusCode），可能导致响应 Body 为空或解析失败；需再确认 SDK 对 xml 的实际行为及当前这些接口是否依赖其他路径（如 agentbay 层自定义解析）。

### 3.3 已修复

- **ListMarketGroupSkill**：已改为 `BodyType: "string"` + 手动解析，成功/失败均可输出 RequestId。

---

## 4. 建议的后续动作（仅规划，本次不实现）

1. **与后端/网关确认：**  
   CreateMarketSkillGroup、AddMarketGroupSkill、RemoveMarketGroupSkill 的**实际**响应格式（Content-Type 与 body 是 JSON 还是 XML）。若为 XML，则与 ListMarketGroupSkill 同风险。
2. **若确认上述三个接口返回 XML：**  
   对三者做与 ListMarketGroupSkill 相同的改造：`BodyType: "string"` + 专用解析函数 + 错误时 RequestId 提取与 `ErrWithRequestID`。
3. **若后端统一返回 JSON：**  
   当前三个接口保持 `BodyType: "json"` 即可；若仍出现解析错误，再按「后端实际返回内容」单独排查。
4. **GetMarketSkillCredential / CreateMarketSkill / DescribeMarketSkillDetail：**  
   若当前功能正常，可维持现状；若出现 body 为空或解析错误，再查 SDK 对 `BodyType: "xml"` 的 body 处理及 agentbay 层是否有单独 XML 解析。

---

## 5. 参考代码位置（便于后续改代码时定位）

- **响应格式与接口约定：** 以 `docs/plans/2025-03-12-skills-cli-implementation.md` 为准。
- BodyType 与 CallApi：`internal/client/client.go`（各 `*WithOptions` 中 `params.BodyType` 与 `dara.Convert`）。
- ListMarketGroupSkill 修复参考：同文件中的 `BodyType: "string"`、`parseListMarketGroupSkillResponse`、`extractRequestIDFromResponse`、`ErrWithRequestID`。
- 响应模型：`internal/client/add_market_group_skill_response_model.go`、`remove_market_group_skill_response_model.go`、`create_market_skill_group_response_model.go`。
