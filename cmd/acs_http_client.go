// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

// acs_http_client.go provides raw HTTP client for Alibaba Cloud POP RPC V1.
// Used by docker login and image create-from-template commands.

package cmd

import (
"crypto/hmac"
"crypto/rand"
"crypto/sha1"
"encoding/base64"
"fmt"
"io"
"net/http"
"net/url"
"sort"
"strings"
"time"

log "github.com/sirupsen/logrus"

"github.com/agentbay/agentbay-cli/internal/auth"
"github.com/agentbay/agentbay-cli/internal/config"
)

// ---------------------------------------------------------------------------
// Direct HTTP client — POP RPC V1 query-string signing
// ---------------------------------------------------------------------------

// acsClient sends requests to Alibaba Cloud OpenAPI (POP RPC style) using raw
// net/http. It supports two auth modes:
//   - BearerToken (OAuth): BearerToken + SignatureType=BEARERTOKEN in query
//   - AK/SK: HMAC-SHA1 V1 signature in query
//
// Protocol: the darabonba-openapi SDK for Style:"RPC" uses the legacy POP
// gateway protocol where ALL parameters (Action, Version, auth fields, API
// params) are encoded in the query string. This is NOT the newer ACS V3
// header-based signing.
type acsClient struct {
endpoint   string // e.g. "xiaoying.cn-shanghai.aliyuncs.com"
apiVersion string // e.g. "2025-05-01"

// Auth — only one set is populated.
bearerToken     string
accessKeyID     string
accessKeySecret string
securityToken   string
}

// newACSClientFromConfig builds an acsClient using the same precedence as the
// SDK wrapper: AK/SK env vars > OAuth token from config file.
func newACSClientFromConfig(cfg *config.Config) (*acsClient, error) {
apiCfg := config.LoadAPIConfig(nil)

c := &acsClient{
endpoint:   apiCfg.Endpoint,
apiVersion: "2025-05-01",
}

// Priority 1: AK/SK from env
if ak, sk, session, ok := config.AccessKeyFromEnv(); ok {
c.accessKeyID = ak
c.accessKeySecret = sk
c.securityToken = session
log.Debugf("[RAW-HTTP] Using AK/SK authentication (AK=%s...)", ak[:6])
return c, nil
}

// Priority 2: OAuth BearerToken
tokenCfgAdapter := auth.NewConfigAdapter(
func() (string, string, time.Time, error) { return cfg.GetTokens() },
cfg.RefreshTokens,
cfg.IsTokenExpired,
cfg.ClearTokens,
)
if err := auth.RefreshTokenIfNeeded(tokenCfgAdapter, config.GetClientID()); err != nil {
return nil, fmt.Errorf("failed to ensure valid token: %w", err)
}
token, err := cfg.GetToken()
if err != nil {
return nil, fmt.Errorf("failed to get token: %w", err)
}
c.bearerToken = token.AccessToken
log.Debugf("[RAW-HTTP] Using BearerToken authentication")
return c, nil
}

// callRPC sends a POP RPC-style request and returns the raw response body.
//
// POP RPC V1 protocol (used by darabonba-openapi SDK for Style:"RPC"):
//
// All parameters — system params + API business params — are encoded in the
// query string of a POST request with empty body.
//
// System params always present:
//
//Action, Version, Format, Timestamp, SignatureNonce
//
// BearerToken mode adds:
//
//BearerToken=<token>, SignatureType=BEARERTOKEN
//
// AK/SK mode adds:
//
//AccessKeyId, SignatureMethod=HMAC-SHA1, SignatureVersion=1.0, Signature=<sig>
//(+ SecurityToken if STS)
//
// Signature algorithm (V1, HMAC-SHA1):
//  1. Collect all query params (excluding Signature), sort by key
//  2. canonicalized = key1=percentEncode(val1)&key2=percentEncode(val2)&...
//  3. stringToSign = "POST&%2F&" + percentEncode(canonicalized)
//  4. Signature = Base64(HMAC-SHA1(AccessKeySecret + "&", stringToSign))
func (c *acsClient) callRPC(action string, apiParams map[string]string) ([]byte, int, error) {
// 1. Build system query params
now := time.Now().UTC().Format("2006-01-02T15:04:05Z")
nonce := generateNonce()

params := map[string]string{
"Action":         action,
"Version":        c.apiVersion,
"Format":         "JSON",
"Timestamp":      now,
"SignatureNonce": nonce,
}

// 2. Add API business params
for k, v := range apiParams {
params[k] = v
}

// 3. Add auth params
if c.bearerToken != "" {
// BearerToken mode: no signature computation needed
params["BearerToken"] = c.bearerToken
params["SignatureType"] = "BEARERTOKEN"
} else {
// AK/SK mode: add AK fields, then compute V1 signature
params["AccessKeyId"] = c.accessKeyID
params["SignatureMethod"] = "HMAC-SHA1"
params["SignatureVersion"] = "1.0"
if c.securityToken != "" {
params["SecurityToken"] = c.securityToken
}
// Compute signature over all params (Signature itself is excluded)
params["Signature"] = c.signV1(params)
}

// 4. Encode query string
q := url.Values{}
for k, v := range params {
q.Set(k, v)
}
rawURL := fmt.Sprintf("https://%s/?%s", c.endpoint, q.Encode())

// 5. Build HTTP request (POST, empty body)
req, err := http.NewRequest(http.MethodPost, rawURL, nil)
if err != nil {
return nil, 0, fmt.Errorf("failed to create request: %w", err)
}
req.Header.Set("Accept", "application/json")
req.Header.Set("User-Agent", "AgentBay-CLI/1.0 (raw-http)")

// 6. Debug log
if log.GetLevel() >= log.DebugLevel {
log.Debugf("[RAW-HTTP] %s https://%s/", req.Method, c.endpoint)
// Print params sorted for readability
keys := make([]string, 0, len(params))
for k := range params {
keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
v := params[k]
if k == "BearerToken" && len(v) > 20 {
v = v[:10] + "..." + v[len(v)-10:]
}
log.Debugf("[RAW-HTTP]   %s = %s", k, v)
}
}

// 7. Send
httpClient := &http.Client{Timeout: 60 * time.Second}
resp, err := httpClient.Do(req)
if err != nil {
return nil, 0, fmt.Errorf("HTTP request failed: %w", err)
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
if err != nil {
return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
}

if log.GetLevel() >= log.DebugLevel {
log.Debugf("[RAW-HTTP] Response status: %d", resp.StatusCode)
if len(body) < 2048 {
log.Debugf("[RAW-HTTP] Response body: %s", string(body))
} else {
log.Debugf("[RAW-HTTP] Response body (%d bytes): %s...", len(body), string(body[:512]))
}
}

return body, resp.StatusCode, nil
}

// signV1 computes the Alibaba Cloud POP RPC V1 signature (HMAC-SHA1).
//
// Algorithm:
//  1. Sort all param keys (Signature key must NOT be in the map)
//  2. canonicalized = popPercentEncode(k1)=popPercentEncode(v1)&...
//  3. stringToSign = "POST&" + popPercentEncode("/") + "&" + popPercentEncode(canonicalized)
//  4. signing key = AccessKeySecret + "&" (the trailing "&" is required per spec)
//  5. Signature = Base64(HMAC-SHA1(signingKey, stringToSign))
func (c *acsClient) signV1(params map[string]string) string {
// 1. Collect and sort keys
keys := make([]string, 0, len(params))
for k := range params {
if k == "Signature" {
continue
}
keys = append(keys, k)
}
sort.Strings(keys)

// 2. Build canonicalized query string with POP percent-encoding
var parts []string
for _, k := range keys {
parts = append(parts, popPercentEncode(k)+"="+popPercentEncode(params[k]))
}
canonicalized := strings.Join(parts, "&")

// 3. StringToSign
stringToSign := "POST&" + popPercentEncode("/") + "&" + popPercentEncode(canonicalized)

if log.GetLevel() >= log.DebugLevel {
log.Debugf("[RAW-HTTP] StringToSign: %s", stringToSign)
}

// 4. Compute HMAC-SHA1
signingKey := c.accessKeySecret + "&"
mac := hmac.New(sha1.New, []byte(signingKey))
mac.Write([]byte(stringToSign))
signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

return signature
}

// popPercentEncode implements the Alibaba Cloud POP-style percent encoding.
// It is similar to url.QueryEscape but:
//   - Space is encoded as %20 (not +)
//   - Asterisk (*) is encoded as %2A
//   - Tilde (~) is NOT encoded (kept as-is)
func popPercentEncode(s string) string {
encoded := url.QueryEscape(s)
// Fix differences from POP spec
encoded = strings.ReplaceAll(encoded, "+", "%20")
encoded = strings.ReplaceAll(encoded, "*", "%2A")
encoded = strings.ReplaceAll(encoded, "%7E", "~")
return encoded
}

// ---------------------------------------------------------------------------
// Crypto helpers
// ---------------------------------------------------------------------------

func generateNonce() string {
b := make([]byte, 16)
_, _ = rand.Read(b)
return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// ---------------------------------------------------------------------------
// General helpers
// ---------------------------------------------------------------------------

func ptrStr(p *string) string {
if p == nil {
return ""
}
return *p
}

func truncateBody(b []byte, max int) string {
if len(b) <= max {
return string(b)
}
return string(b[:max]) + "..."
}
