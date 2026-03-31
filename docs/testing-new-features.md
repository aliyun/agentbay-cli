# AgentBay CLI 全量回归测试步骤说明

本文档用于**手工执行**全功能回归，约定：

- 命令统一使用当前目录下的 **`./agentbay`**（在仓库根目录 `go build -o agentbay .` 后执行）。
- **线上**：中国大陆生产环境（`AGENTBAY_ENV` 未设置、或 `production` / `prod`）。
- **预发**：中国大陆预发环境（`AGENTBAY_ENV=prerelease` 或 `pre` / `staging`）。
- **平台**：**macOS** 与 **Linux** 各跑一遍相同步骤（命令一致；**鉴权覆盖不同**，差异见第 2 节）。
- **鉴权**：
  - **macOS**：在每个「平台 × 环境（线上/预发）」下须覆盖 **OAuth（`login`）** 与 **AccessKey / SecretKey（`AGENTBAY_ACCESS_KEY_*` 环境变量）** 两种路径（至少镜像 list/status/create、activate、skills 等）；不能只用其中一种就勾选整轮通过。
  - **Linux**：**仅使用 AK/SK（环境变量）** 做鉴权回归，**不执行** OAuth 登录、登出、刷新等小节（见 **5.2、5.8** 及勾选表说明）。

用户手册见 [USER_GUIDE.md](./USER_GUIDE.md)；自动化测试见 [test/README.md](../test/README.md)。

---

## 1. 当前 CLI 命令范围（须全部覆盖）

| 分组 | 命令 | 说明 |
|------|------|------|
| 核心 | `./agentbay version`、`./agentbay --help` | 版本与根帮助 |
| 认证 | `./agentbay login`、`./agentbay logout`；`AGENTBAY_ACCESS_KEY_*` | **macOS**：OAuth 与 AK/SK **都要测**；**Linux**：**只测 AK/SK** |
| 镜像 | `./agentbay image create`、`list`、`init`、`activate`、`deactivate`、`status` | 镜像全流程 |
| 技能 | `./agentbay skills push`、`skills show` | 目录或 zip |
| 全局 | `./agentbay -v ...` | 详细日志（按需） |

**鉴权方式**：本地 OAuth（`login` 后 `~/.config/agentbay/config.json` 等）；**仅**环境变量 **`AGENTBAY_ACCESS_KEY_ID` + `AGENTBAY_ACCESS_KEY_SECRET`**（可选 `AGENTBAY_ACCESS_KEY_SESSION_TOKEN`，即 AK/SK/STS）。同时存在时 CLI **优先 AccessKey**。

**回归要求**：**macOS**——除 **5.2 / 5.3** 外，凡写「**再次 OAuth 或 AK**」的小节，在同一环境下应**分别**用 OAuth 登录态与「仅 AK、无有效 OAuth」各执行一次关键命令。**Linux**——全程以 **5.3** 的 AK/SK 方式执行后续镜像与 skills 步骤，**跳过 5.2、5.8**；「再次 OAuth 或 AK」处**仅使用 AK/SK**。

---

## 2. 平台差异（macOS / Linux）

| 项 | macOS | Linux |
|----|--------|--------|
| 构建 | `go build -o agentbay .` | 同上 |
| 可执行权限 | 一般可直接 `./agentbay` | 若无执行位：`chmod +x agentbay` |
| 鉴权回归 | OAuth + AccessKey（AK/SK）**两种都测** | **仅 AK/SK**（环境变量），不测 OAuth 全流程 |
| `login` 回调 | 本机浏览器打开 `localhost:3001`（或配置端口） | 回归不要求跑 `login`；若手工抽查，无桌面时可按终端提示打开 URL |
| 路径 | 示例用 `/path/to/...` 或仓库内相对路径 | 同左 |
| 配置文件目录 | `~/.config/agentbay/`（XDG） | 通常相同 |

**回归要求**：下文章节中的步骤在 **macOS、Linux 各执行一轮**（或团队分工），**Linux 以 AK/SK 为唯一鉴权方式**；勾选表见第 6.2 节说明。

---

## 3. 环境差异（线上 / 预发）

| 项 | 线上（中国大陆生产） | 预发（中国大陆预发） |
|----|----------------------|----------------------|
| 典型变量 | 不设置或 `export AGENTBAY_ENV=production` | `export AGENTBAY_ENV=prerelease` |
| 验证 | `./agentbay version` 中 **Environment** 为 production，**Endpoint** 含 `cn-shanghai` | **Environment** 为 prerelease，**Endpoint** 含 `cn-hangzhou` 与 `pre` |
| 系统镜像 ID | 文档常见：`code-space-debian-12` 等 | **与生产不一定相同**；须 `./agentbay image list --system-only` 取**当前环境**下的 ID |
| `image create` 的 Dockerfile | **均为 5.5 节**固定 `FROM` | **均为 5.5 节**固定 `FROM` |
| OAuth Client ID | 与生产一致 | 与预发一致（见实现或 `version` 行为） |

**回归要求**：**同一套步骤分别在「仅线上」「仅预发」下执行**（切换环境前可 `logout` 或换 `AGENTBAY_CLI_CONFIG_DIR` 避免串环境）。

> **国际站**（`AGENTBAY_ENV=international` 等）若团队需要，可在上述两套通过后再加一轮；步骤结构相同，Endpoint 与镜像 ID 以该环境为准。

---

## 4. 测试前准备（每条环境、每台机器）

1. 安装 Go、能访问对应 AgentBay / 阿里云网络。
2. 仓库根目录：`go build -o agentbay .`（Linux 按需 `chmod +x agentbay`）。
3. 准备测试资源：
   - 做 **`image create` 回归**时，**线上与预发** Dockerfile 内容均以 **5.5 节**固定 `FROM` 为准（`init` 模板仅作其他用途时可另备）。
   - 合法 **Skill** 目录（含 `SKILL.md` frontmatter）或 **zip**。
4. **勿**将 AccessKey、Token、OSS 带签 URL 写入仓库或公开渠道。

---

## 5. 分步回归（每个小节在「线上」「预发」各做一遍）

以下步骤中，**先执行** `export AGENTBAY_ENV=...`（或 `unset AGENTBAY_ENV` 表示线上），再按顺序测。

**鉴权**：**macOS**——每一轮须覆盖 **OAuth** 与 **AccessKey**，见 **5.2、5.3**；后续「再次 OAuth 或 AK」的小节，两种鉴权下各验证关键命令。**Linux**——只做 **5.3（AK/SK）**，跳过 **5.2、5.8**；后续小节在已导出 AK/SK 的前提下执行。

### 5.1 版本与帮助

```bash
./agentbay version
./agentbay --help
./agentbay image --help
./agentbay skills --help
./agentbay login --help
./agentbay logout --help
```

**预期**：无 panic；`version` 中 Environment / Endpoint 与当前 `AGENTBAY_ENV` 一致；`image` 含 create、list、init、activate、deactivate、status；`skills` 含 push、show。

---

### 5.2 OAuth：登录 → 列表 → 登出 → 未认证

**适用范围**：**仅 macOS** 回归必做。**Linux** 跳过本节。

```bash
./agentbay login
./agentbay image list
./agentbay -v image list    # 可选
./agentbay logout
./agentbay image list       # 未设 AK 时应失败并提示 login 或 AK 环境变量
```

**预期**：登录后列表成功；登出后失败且文案含认证引导。

可选：

```bash
./agentbay image status <任意镜像ID>
./agentbay skills show dummy
```

**预期**：未认证时失败（或 skills 因 API 不同略有差异，以不静默成功为准）。

---

### 5.3 AccessKey 鉴权（AK / SK，与 OAuth 互斥场景）

**适用范围**：**macOS 与 Linux 均必做**；Linux 以本节作为**唯一**鉴权入口（不测 5.2）。

```bash
./agentbay logout
unset AGENTBAY_CLI_CONFIG_DIR   # 若曾设置
# export AGENTBAY_CLI_CONFIG_DIR="$(mktemp -d)"   # 可选：彻底隔离 OAuth
export AGENTBAY_ACCESS_KEY_ID="..."
export AGENTBAY_ACCESS_KEY_SECRET="..."
# export AGENTBAY_ACCESS_KEY_SESSION_TOKEN="..."   # STS 时
./agentbay image list
./agentbay image status <真实存在的 image-id>
```

**预期**：无有效 OAuth 时仅凭 AK 可访问 API（权限足够前提下）。

**与 OAuth 同时存在**（**仅 macOS**）：先 `login` 再导出 AK，执行 `./agentbay -v image list`，确认行为符合「**优先 AK**」的说明。Linux 回归可跳过本段。

```bash
unset AGENTBAY_ACCESS_KEY_ID AGENTBAY_ACCESS_KEY_SECRET AGENTBAY_ACCESS_KEY_SESSION_TOKEN
```

---

### 5.4 镜像：`list` / `init` / `status`

在已认证前提下执行（**macOS**：OAuth 与 AK 两种各测一遍；**Linux**：**仅 AK/SK**）：

```bash
./agentbay image list
./agentbay image list --include-system
./agentbay image list --system-only
./agentbay image list --os-type Linux --size 5 -p 1    # 可选组合
```

```bash
# 将 -i 换成本环境 system-only 列表中的真实 ID（预发勿照抄生产 ID）
./agentbay image init -i <本环境系统镜像ID>
```

```bash
./agentbay image status <用户镜像或系统镜像 ID>
./agentbay -v image status <同上>
```

参数错误：

```bash
./agentbay image status
./agentbay image status a b
```

**预期**：列表与 init、status 成功路径无报错；缺参时 Cobra 报错清晰。

---

### 5.5 镜像：`create`（上传 + 创建任务）

**`--imageId` 必须为本环境存在的 base image**（线上、预发分别用当前环境下的系统/基础镜像 ID，勿混用）。

**线上与预发**：做本小节回归时，Dockerfile 内容**统一**为下面一行（可整文件仅此一行）：

```dockerfile
FROM ai-container-registry-vpc.cn-hangzhou.cr.aliyuncs.com/qwen/qwen_claw:main-2026032914-ca9d486fbe
```

将上述内容写入例如 `./Dockerfile` 后再执行：

```bash
./agentbay image create <名称> --dockerfile ./Dockerfile --imageId <本环境base镜像ID> -v
```

**预期**：STEP 1 凭证、STEP 2 上传、STEP 4 创建任务均成功（或失败时错误可读，且与 JSON/XML 解析无关）。若构建较慢，允许在合理时间内保持 RUNNING 后由后续 `GetDockerImageTask` 轮询（见实现）。

---

### 5.6 镜像：`activate` / `deactivate`（需用户镜像）

仅对**用户镜像**（非仅系统可用镜像）：

```bash
./agentbay image activate <imgc-...>
./agentbay image deactivate <imgc-...>
```

**预期**：与文档一致；资源规格 `--cpu` / `--memory` 可选测一组合法组合。

---

### 5.7 Skills

```bash
./agentbay skills push /path/to/skill-dir
./agentbay skills push /path/to/skill.zip
./agentbay skills show <返回的 SkillId>
./agentbay skills push
./agentbay skills show
```

**预期**：成功路径返回 ID；`show` 有详情；无参报错。

---

### 5.8 OAuth 刷新（可选）

**适用范围**：**仅 macOS**（可选）。**Linux** 跳过。

在 access token 临近过期或正常使用一段时间后：

```bash
./agentbay -v image list
```

**预期**：不因 `expires_in` 为字符串等原因解码失败而整段退出登录（见历史问题修复说明）。

---

## 6. 总勾选表（建议打印或复制到工单）

### 6.1 平台 × 环境

| 平台 | 线上 | 预发 |
|------|------|------|
| macOS | ☐ | ☐ |
| Linux | ☐ | ☐ |

### 6.2 功能（每个「平台×环境」组合下勾选）

**说明**：**macOS** 勾选第 2、4、11 项（OAuth 相关）；**Linux** 将第 2、4、11 项视为 **N/A**，以第 3 项（5.3 AK/SK）作为鉴权主路径，其余项（5、6…10）在 AK/SK 下照常勾选。

| # | 项 | 通过 |
|---|-----|------|
| 1 | version / help | ☐ |
| 2 | （**macOS**）OAuth：`login` → list → `logout` → 未认证 list | ☐ |
| 3 | AccessKey（AK/SK）：list / status（**Linux 必做**；**macOS** 亦须做，且 macOS 还须完成第 2 项） | ☐ |
| 4 | （**macOS**）OAuth+AK 同时存在（优先 AK） | ☐ |
| 5 | image list（含 --system-only 等至少一种组合） | ☐ |
| 6 | image init（本环境系统镜像 ID） | ☐ |
| 7 | image status（含 -v 与缺参） | ☐ |
| 8 | image create（完整上传+建任务；**线上与预发** Dockerfile 固定 `FROM` 见 **5.5**） | ☐ |
| 9 | image activate / deactivate（有用户镜像时） | ☐ |
| 10 | skills push（目录或 zip）+ show + 缺参 | ☐ |
| 11 | （**macOS** 可选）`-v` 与 OAuth 刷新（**5.8**） | ☐ |

### 6.3 记录字段（每次发布建议填写）

- **日期**、**CLI 版本 / commit**、**执行人**、**平台与 OS 版本**、**AGENTBAY_ENV**、**鉴权方式**（如 macOS：OAuth+AK / 双存在；Linux：仅 AK/SK）、**网络（内网/VPN）**、**阻塞项与 RequestId**。

---

## 附录：自动化测试（可选）

不等同于上文手工回归，可用于提交前快速检查：

```bash
go test ./... -count=1
go test -tags=integration ./test/integration/... -count=1
```

真 API 相关用例见 `test/integration/` 内说明及环境变量（如 `RUN_INTEGRATION_TESTS`）。
