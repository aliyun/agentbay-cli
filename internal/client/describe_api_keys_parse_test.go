// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestParseDescribeApiKeysResponse_JSONRealServerFormat tests the actual SDK response format
// where the body IS the data payload directly (SDK strips the outer wrapper).
// Field names: ApiKeys (PascalCase), requestId (lowercase), Count (PascalCase), NextToken (PascalCase).
func TestParseDescribeApiKeysResponse_JSONRealServerFormat(t *testing.T) {
	t.Parallel()
	// Body is the direct data payload — SDK strips outer wrapper
	body := `{"ApiKeys":[{"Status":"ENABLED","GmtCreate":"2026-04-10T10:10:57+08:00","LastUseDate":"2026-05-19T10:03:05+08:00","ApiKey":"akm-7fdffcbe-******************837e","Concurrency":"21","KeyId":"ak-df1u29s116881nt6q","Name":"lxy-cli-3"},{"Status":"DISABLED","GmtCreate":"2026-03-15T08:30:00+08:00","LastUseDate":"2026-05-18T14:22:11+08:00","ApiKey":"akm-abc123","Concurrency":5,"KeyId":"ak-abc123xyz456def7","Name":"prod-key"}],"requestId":"R-real","Count":"2","NextToken":"NEXT_PAGE_TOKEN"}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeApiKeysResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	// Code/Success/HttpStatusCode are not in the body; they remain nil
	require.Equal(t, "", out.Body.GetCode())
	require.False(t, out.Body.GetSuccess()) // Success nil → GetSuccess returns false
	// RequestId comes from the data-level requestId field
	require.Equal(t, "R-real", out.Body.GetRequestId())
	require.NotNil(t, out.Body.Data)
	require.Len(t, out.Body.Data.ApiKeys, 2)

	key1 := out.Body.Data.ApiKeys[0]
	require.Equal(t, "ENABLED", key1.GetStatus())
	require.Equal(t, "ak-df1u29s116881nt6q", key1.GetKeyId())
	require.Equal(t, "lxy-cli-3", key1.GetName())
	require.Equal(t, int32(21), key1.GetConcurrency())
	require.Equal(t, "2026-04-10T10:10:57+08:00", key1.GetGmtCreate())

	key2 := out.Body.Data.ApiKeys[1]
	require.Equal(t, "DISABLED", key2.GetStatus())
	require.Equal(t, "ak-abc123xyz456def7", key2.GetKeyId())
	require.Equal(t, "prod-key", key2.GetName())
	require.Equal(t, int32(5), key2.GetConcurrency())

	require.Equal(t, "2", out.Body.Data.GetCount())
	require.Equal(t, "NEXT_PAGE_TOKEN", out.Body.Data.GetNextToken())
}

// TestParseDescribeApiKeysResponse_JSONWrappedFormat tests backward compatibility
// for the case where the SDK returns the full wrapped response with outer-level fields.
func TestParseDescribeApiKeysResponse_JSONWrappedFormat(t *testing.T) {
	t.Parallel()
	body := `{"code":"200","data":{"ApiKeys":[{"Status":"ENABLED","GmtCreate":"2026-04-10T10:10:57+08:00","LastUseDate":"2026-05-19T10:03:05+08:00","ApiKey":"akm-7fdffcbe-******************837e","Concurrency":"21","KeyId":"ak-df1u29s116881nt6q","Name":"lxy-cli-3"}],"requestId":"R1","Count":"1","NextToken":"AAAAA"},"httpStatusCode":"200","requestId":"R-outer","successResponse":true}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeApiKeysResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	// Outer-level requestId takes precedence
	require.Equal(t, "R-outer", out.Body.GetRequestId())
	require.True(t, out.Body.GetSuccess())
	require.Equal(t, "200", out.Body.GetCode())
	require.NotNil(t, out.Body.Data)
	require.Len(t, out.Body.Data.ApiKeys, 1)

	key := out.Body.Data.ApiKeys[0]
	require.Equal(t, "ENABLED", key.GetStatus())
	require.Equal(t, "ak-df1u29s116881nt6q", key.GetKeyId())
	require.Equal(t, "lxy-cli-3", key.GetName())
	require.Equal(t, int32(21), key.GetConcurrency())
	require.Equal(t, "2026-04-10T10:10:57+08:00", key.GetGmtCreate())
	require.Equal(t, "2026-05-19T10:03:05+08:00", key.GetLastUseDate())
	require.Equal(t, "1", out.Body.Data.GetCount())
	require.Equal(t, "AAAAA", out.Body.Data.GetNextToken())
}

// TestParseDescribeApiKeysResponse_JSONHttpStatusCodeAsNumber tests direct data payload
// with HttpStatusCode present at the body level (as JSON number).
func TestParseDescribeApiKeysResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"ApiKeys":[{"Status":"DISABLED","GmtCreate":"2026-03-15T08:30:00+08:00","LastUseDate":"2026-05-18T14:22:11+08:00","ApiKey":"akm-abc123","Concurrency":5,"KeyId":"ak-abc123xyz456def7","Name":"prod-key"}],"requestId":"R2","Count":"1","NextToken":""}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeApiKeysResponse(res)
	require.NoError(t, err)
	require.Equal(t, "DISABLED", out.Body.Data.ApiKeys[0].GetStatus())
	require.Equal(t, "ak-abc123xyz456def7", out.Body.Data.ApiKeys[0].GetKeyId())
	require.Equal(t, int32(5), out.Body.Data.ApiKeys[0].GetConcurrency())
}

func TestParseDescribeApiKeysResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<DescribeApiKeysResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`<Data>` +
		`<ApiKeys>` +
		`<ApiKey>` +
		`<Status>ENABLED</Status>` +
		`<GmtCreate>2026-04-10T10:10:57+08:00</GmtCreate>` +
		`<LastUseDate>2026-05-19T10:03:05+08:00</LastUseDate>` +
		`<ApiKey>akm-7fdffcbe-******************837e</ApiKey>` +
		`<Concurrency>21</Concurrency>` +
		`<KeyId>ak-df1u29s116881nt6q</KeyId>` +
		`<Name>lxy-cli-3</Name>` +
		`</ApiKey>` +
		`</ApiKeys>` +
		`<RequestId>R3</RequestId>` +
		`<Count>1</Count>` +
		`<NextToken>BBBBB</NextToken>` +
		`</Data>` +
		`</DescribeApiKeysResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeApiKeysResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
	require.NotNil(t, out.Body.Data)
	require.Len(t, out.Body.Data.ApiKeys, 1)

	key := out.Body.Data.ApiKeys[0]
	require.Equal(t, "ENABLED", key.GetStatus())
	require.Equal(t, "ak-df1u29s116881nt6q", key.GetKeyId())
	require.Equal(t, "lxy-cli-3", key.GetName())
	require.Equal(t, int32(21), key.GetConcurrency())
	require.Equal(t, "1", out.Body.Data.GetCount())
	require.Equal(t, "BBBBB", out.Body.Data.GetNextToken())
}
