# API 格式规范

## Product ID 映射

| 项目                | Product ID               | 说明           |
| ------------------- | ------------------------ | -------------- |
| agent-bay（前端）   | `xiaoying-double-centre` | 前端控制台使用 |
| agentbay-cli（CLI） | `xiaoying`               | CLI 工具使用   |

**重要**: 两者调用的是相同的后端 API，只是 product ID 命名不同。

## API 通用配置

```go
Version:     "2025-05-01"
Protocol:    "HTTPS"
Pathname:    "/"
Method:      "POST"
AuthType:    "AK"
Style:       "RPC"
ReqBodyType: "formData"
BodyType:    "string"
```

## Endpoint 配置

| 环境         | Endpoint                                |
| ------------ | --------------------------------------- |
| 生产（默认） | `xiaoying.cn-shanghai.aliyuncs.com`     |
| 预发         | `xiaoying-pre.cn-hangzhou.aliyuncs.com` |
| 国际站       | 待确认                                  |

**环境切换**: 通过 `AGENTBAY_ENV` 环境变量

```bash
AGENTBAY_ENV=production   # 生产环境
AGENTBAY_ENV=prerelease   # 预发环境
```

## 请求格式

**Content-Type**: `application/x-www-form-urlencoded`

**请求示例** (curl):

```bash
curl -X POST 'https://xiaoying.cn-shanghai.aliyuncs.com/' \
  -d 'Action=CreateApiKey' \
  -d 'Version=2025-05-01' \
  -d 'Name=my-api-key'
```

## 响应格式注意事项

### ⚠️ 响应解析容错模板（必须使用）

**规则**: 所有 `parseXxxResponse` 统一放在 [internal/client/dual_format_responses.go](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/dual_format_responses.go)，**不**要再写在 `client.go` 里（它只保留 auto-generated 的 action 骨架）。

**强制要求**：

1. **数字字段不要直接 Unmarshal 到 `*int32`**：服务端常常把 `HttpStatusCode` / `NonEditLineNum` 等数字字段序列化为字符串，直接 `json.Unmarshal` 会报：

   ```text
   json: cannot unmarshal string into Go struct field Xxx.HttpStatusCode of type int32
   ```

   正确做法：在 JSON wire 类型里用 `json.RawMessage` 接盘，再用 `int32FromFlexibleJSON` 辅助函数解析，兼容数字与字符串。

2. **XML / JSON 双格式都要支持**：body 以 `<` 开头走 XML 分支、否则走 JSON 分支，两者都要调用 `applyMapHeadersAndStatus` 归一 headers / statusCode。

3. **解析失败包装**：任何解析失败都必须返回 `&ErrWithRequestID{Err: ..., RequestID: extractRequestIDFromResponse(res)}`，保证 RequestId 能透出到 CLI。

4. **最小单测**：每个 parser 配套一个 `internal/client/xxx_parse_test.go`，至少覆盖：
   - JSON 数字字段返回为字符串（`"HttpStatusCode":"200"`）
   - JSON 数字字段返回为数字（`"HttpStatusCode":200`）
   - XML 分支

**反面案例**（已修复，禁止复现）:

`BatchCreateHideResourceGroupsWithMaxSession` 早期 parser 直接 `json.Unmarshal` 到 `BatchCreateHideResourceGroupsWithMaxSessionResponseBody{HttpStatusCode *int32}`，服务端返回 `"HttpStatusCode":"200"` 导致 `agentbay image set-max-session` 整个命令直接失败。同类问题也在 `GetDockerfileTemplate.NonEditLineNum` 上出现过。

**模板（双格式 + 容错数字）**：

```go
type xxxJSONWire struct {
    Code           *string         `json:"Code"`
    Message        *string         `json:"Message"`
    RequestId      *string         `json:"RequestId"`
    HttpStatusCode json.RawMessage `json:"HttpStatusCode"`
    Success        *bool           `json:"Success"`
}

type xmlXxxResponse struct {
    XMLName        xml.Name `xml:"XxxResponse"`
    RequestId      string   `xml:"RequestId"`
    HttpStatusCode string   `xml:"HttpStatusCode"`
    Code           string   `xml:"Code"`
    Success        bool     `xml:"Success"`
    Message        string   `xml:"Message"`
}

func parseXxxResponse(res map[string]interface{}) (*XxxResponse, error) {
    bodyStr, err := rawBodyStringFromMap(res)
    if err != nil {
        return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
    }
    out := &XxxResponse{Headers: make(map[string]*string)}
    parsed := &XxxResponseBody{}
    trimmed := strings.TrimSpace(bodyStr)
    if bodyStr != "" {
        if len(trimmed) > 0 && trimmed[0] == '<' {
            var xr xmlXxxResponse
            if err := xml.Unmarshal([]byte(bodyStr), &xr); err != nil {
                return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
            }
            parsed.Code = dara.String(xr.Code)
            parsed.RequestId = dara.String(xr.RequestId)
            parsed.Success = dara.Bool(xr.Success)
            parsed.Message = dara.String(xr.Message)
            if s := strings.TrimSpace(xr.HttpStatusCode); s != "" {
                if n, perr := strconv.ParseInt(s, 10, 32); perr == nil {
                    parsed.HttpStatusCode = dara.Int32(int32(n))
                }
            }
        } else {
            var wire xxxJSONWire
            if err := json.Unmarshal([]byte(bodyStr), &wire); err != nil {
                return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
            }
            parsed.Code = wire.Code
            parsed.Message = wire.Message
            parsed.RequestId = wire.RequestId
            parsed.Success = wire.Success
            n, derr := int32FromFlexibleJSON(wire.HttpStatusCode)
            if derr != nil {
                return nil, &ErrWithRequestID{Err: fmt.Errorf("HttpStatusCode: %w", derr), RequestID: extractRequestIDFromResponse(res)}
            }
            parsed.HttpStatusCode = n
        }
    }
    out.Body = parsed
    applyMapHeadersAndStatus(&out.Headers, &out.StatusCode, res)
    return out, nil
}
```

**参考实现**：

- [parseBatchCreateHideResourceGroupsWithMaxSessionResponse](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/dual_format_responses.go)
- [parseGetDockerfileTemplateResponse](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/dual_format_responses.go)
- [batch_create_hide_resource_groups_with_max_session_parse_test.go](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/batch_create_hide_resource_groups_with_max_session_parse_test.go)
- [get_dockerfile_template_parse_test.go](file:///Users/lxy/work/project/ai/agentbay/agentbay-cli/internal/client/get_dockerfile_template_parse_test.go)

**检查清单**：

- [ ] parser 写在 `dual_format_responses.go`，没有放回 `client.go`
- [ ] 数字字段走 `int32FromFlexibleJSON`，不直接 Unmarshal 到 `*int32`
- [ ] JSON / XML 两条分支齐备
- [ ] `xxx_parse_test.go` 覆盖「数字字段为字符串 / 数字 / XML」三种场景
- [ ] 解析失败全部走 `ErrWithRequestID` 包装

### ⚠️ Data 字段类型

后端返回的 `Data` 字段类型**可能不是对象**，需要根据实际响应调整：

**情况 1: Data 是字符串**

```json
{
  "Code": "ok",
  "Data": "ak-d06m2mftwy4jpasw8", // 字符串，直接就是 KeyId
  "RequestId": "xxx"
}
```

对应模型：

```go
type CreateApiKeyResponseBody struct {
    Data *string `json:"Data,omitempty"`
}

func (s *CreateApiKeyResponseBody) GetData() string {
    if s == nil || s.Data == nil {
        return ""
    }
    return *s.Data
}
```

**情况 2: Data 是对象**

```json
{
  "Code": "ok",
  "Data": {
    "KeyId": "ak-xxx",
    "Name": "my-key"
  },
  "RequestId": "xxx"
}
```

对应模型：

```go
type CreateApiKeyResponseBody struct {
    Data *CreateApiKeyResponseBodyData `json:"Data,omitempty"`
}

type CreateApiKeyResponseBodyData struct {
    KeyId *string `json:"KeyId,omitempty"`
    Name  *string `json:"Name,omitempty"`
}
```

**调试方法**: 遇到 JSON 解析错误时，添加调试代码打印原始响应：

```go
fmt.Printf("[DEBUG] Raw response: %s\n", bodyStr)
```

## SDK 客户端方法模板

```go
// {Action}WithOptions 完整调用方法
func (client *Client) {Action}WithOptions(request *{Action}Request, runtime *dara.RuntimeOptions) (_result *{Action}Response, _err error) {
    _err = request.Validate()
    if _err != nil {
        return _result, _err
    }
    query := map[string]interface{}{}
    if !dara.IsNil(request.{Field1}) {
        query["{Field1}"] = request.{Field1}
    }
    if !dara.IsNil(request.{Field2}) {
        query["{Field2}"] = request.{Field2}
    }

    req := &openapiutil.OpenApiRequest{
        Query: openapiutil.Query(query),
        Headers: map[string]*string{
            "Accept": dara.String("application/json"),
        },
    }
    params := &openapiutil.Params{
        Action:      dara.String("{Action}"),
        Version:     dara.String("2025-05-01"),
        Protocol:    dara.String("HTTPS"),
        Pathname:    dara.String("/"),
        Method:      dara.String("POST"),
        AuthType:    dara.String("AK"),
        Style:       dara.String("RPC"),
        ReqBodyType: dara.String("formData"),
        BodyType:    dara.String("string"),
    }
    _result = &{Action}Response{}
    _body, _err := client.CallApi(params, req, runtime)
    if _err != nil {
        reqID := ""
        if _body != nil {
            reqID = extractRequestIDFromResponse(_body)
        }
        return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
    }
    _result, _err = parse{Action}Response(_body)
    return _result, _err
}

// {Action} 简化调用方法
func (client *Client) {Action}(request *{Action}Request) (_result *{Action}Response, _err error) {
    runtime := &dara.RuntimeOptions{}
    return client.{Action}WithOptions(request, runtime)
}

// {Action}WithContext 支持 context 的调用方法
func (client *Client) {Action}WithContext(ctx context.Context, request *{Action}Request, runtime *dara.RuntimeOptions) (_result *{Action}Response, _err error) {
    // 类似实现
}

// parse{Action}Response 响应解析函数
func parse{Action}Response(res map[string]interface{}) (*{Action}Response, error) {
    out := &{Action}Response{}
    bodyStr := ""
    switch v := res["body"].(type) {
    case string:
        bodyStr = v
    case []byte:
        bodyStr = string(v)
    default:
        return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
    }
    parsed := &{Action}ResponseBody{}
    if bodyStr != "" {
        if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
            return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
        }
    }
    out.Body = parsed
    // ... 解析 headers 和 statusCode
    return out, nil
}
```
