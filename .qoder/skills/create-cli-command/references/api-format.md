# API 格式规范

## Product ID 映射

| 项目 | Product ID | 说明 |
|-----|-----------|------|
| agent-bay（前端） | `xiaoying-double-centre` | 前端控制台使用 |
| agentbay-cli（CLI） | `xiaoying` | CLI 工具使用 |

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

| 环境 | Endpoint |
|-----|----------|
| 生产（默认） | `xiaoying.cn-shanghai.aliyuncs.com` |
| 预发 | `xiaoying-pre.cn-hangzhou.aliyuncs.com` |
| 国际站 | 待确认 |

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

### ⚠️ Data 字段类型

后端返回的 `Data` 字段类型**可能不是对象**，需要根据实际响应调整：

**情况 1: Data 是字符串**
```json
{
  "Code": "ok",
  "Data": "ak-d06m2mftwy4jpasw8",  // 字符串，直接就是 KeyId
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
