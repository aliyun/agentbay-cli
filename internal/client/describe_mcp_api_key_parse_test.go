// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDescribeMcpApiKeyResponse_JSONHttpStatusCodeAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":"200","RequestId":"R1","Success":true,"Data":{"Status":"DISABLED","ApiKeyId":"ak-xxx","Name":"test-key","AliUid":"123456"}}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeMcpApiKeyResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", out.Body.GetRequestId())
	require.True(t, out.Body.GetSuccess())
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "DISABLED", out.Body.Data.GetStatus())
	require.Equal(t, "ak-xxx", out.Body.Data.GetApiKeyId())
	require.Equal(t, "test-key", out.Body.Data.GetName())
}

func TestParseDescribeMcpApiKeyResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":200,"RequestId":"R2","Success":true,"Data":{"Status":"ENABLED","ApiKeyId":"ak-yyy","Name":"prod-key","AliUid":"789"}}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeMcpApiKeyResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "ENABLED", out.Body.Data.GetStatus())
	require.Equal(t, "ak-yyy", out.Body.Data.GetApiKeyId())
}

func TestParseDescribeMcpApiKeyResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<DescribeMcpApiKeyResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`<Data>` +
		`<Status>DISABLED</Status>` +
		`<ApiKeyId>ak-zzz</ApiKeyId>` +
		`<Name>xml-key</Name>` +
		`<AliUid>456</AliUid>` +
		`</Data>` +
		`</DescribeMcpApiKeyResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeMcpApiKeyResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "DISABLED", out.Body.Data.GetStatus())
	require.Equal(t, "ak-zzz", out.Body.Data.GetApiKeyId())
	require.Equal(t, "xml-key", out.Body.Data.GetName())
}
