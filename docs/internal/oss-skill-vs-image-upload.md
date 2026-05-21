# Skill OSS 上传：为何对象名没有 .zip，与 Image 的差异

## 现象

- **Image 上传 Dockerfile**：OSS 里对象 key 是真实文件名（如 `taskId/Dockerfile`），Create 时用 TaskId 关联，无需 OssFilePath。
- **Skill 上传 zip**：OSS 里对象 key 是后端返回的随机 id（如 `1762926266827681/ZJvvEdFD`），没有 `.zip` 后缀；传给 CreateMarketSkill 的 OssFilePath 也是同一串，没有“.zip”。内容虽是 zip，但 key 不是“真实 zip 文件名”。

## 原因分析

### Image 流程（有真实文件名）

1. 请求凭证时**带上要存的文件路径**：
   - `GetDockerFileStoreCredentialRequest{ FilePath: "Dockerfile", IsDockerfile: "true" }`
   - ADD/COPY 时每个文件带 `FilePath: relPath`（如 `src/main.py`）
2. 后端按该 path 生成预签名 URL，对象 key 即该 path（如 `taskId/Dockerfile`、`taskId/src/main.py`）。
3. CLI 用该 URL 做 PUT，OSS 上就是“真实文件名”。

### Skill 流程（当前无 .zip）

1. 请求凭证时**没有传 path/filename**：
   - `GetMarketSkillCredentialRequest{}` 为空，无 FilePath 等字段。
2. 后端自己生成预签名 URL，path 为随机 id，**未加 .zip 后缀**（如 `1762926266827681/ZJvvEdFD`）。
3. CLI 只能原样用该 URL 上传，不能改 path（改 path 会破坏签名），所以：
   - OSS 对象 key = `1762926266827681/ZJvvEdFD`（无 .zip）
   - 解析出的 OssFilePath 也 = `1762926266827681/ZJvvEdFD`，传给 CreateMarketSkill 的同样是这个无后缀 path。

结论：**对象名和 filepath 没有 .zip，是因为后端下发的预签名 URL 的 path 本身就没有 .zip；客户端不能擅自改 path。**

## 正确做法（与 Image 对齐）

要让“OSS 里的对象”和“CreateMarketSkill 的 OssFilePath”都是**真实 zip 名称**（带 .zip），必须从**生成 URL 的一方**改：

- **方案 A（推荐）**  
  后端在生成 Skill 上传凭证时，把 object key 定为带 `.zip` 的路径，例如：
  - 当前：`prefix/{randomId}`  
  - 改为：`prefix/{randomId}.zip`  
  这样：
  - 预签名 URL 的 path 已是 `.../xxx.zip`
  - CLI 原样 PUT，OSS 对象即为 `xxx.zip`
  - 从 URL 解析出的 OssFilePath 也是 `prefix/xxx.zip`，传给 CreateMarketSkill 即可，无需改 CLI。

- **方案 B（与 Image 完全一致）**  
  若后端支持“按文件名生成 URL”（像 Image 的 FilePath）：
  - 在 `GetMarketSkillCredential` 请求中增加可选参数（如 `FilePath` 或 `ObjectName`），CLI 传 `"skill.zip"`（或后端约定的名字）。
  - 后端用该 path 生成预签名 URL（如 `prefix/skill.zip` 或 `prefix/{id}.zip`）。
  - 效果同方案 A：OSS 和 OssFilePath 都是真实 zip 名。

## CLI 侧现状

- 已按“与 Image 一致”的方式上传：临时文件 + `uploadFileToOSS`，并处理了 307/308 重定向的 body 重发。
- OssFilePath 来自凭证响应的 OssFilePath，或从预签名 URL 解析；**与实际上传的 URL path 一致**，不做篡改。
- 一旦后端改为返回 path 带 `.zip` 的 URL（方案 A 或 B），当前 CLI 无需改即可得到“OSS 对象名 = 真实 zip 名、OssFilePath = 真实 zip 名”的效果。
