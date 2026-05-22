// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseModifyApiKeyStatusResponse_JSONHttpStatusCodeAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":"200","RequestId":"R1","Success":true}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseModifyApiKeyStatusResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", out.Body.GetRequestId())
	require.True(t, out.Body.GetSuccess())
}

func TestParseModifyApiKeyStatusResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":200,"RequestId":"R2","Success":true}`
	res := map[string]interface{}{"body": body}
	out, err := parseModifyApiKeyStatusResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
}

func TestParseModifyApiKeyStatusResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<ModifyApiKeyStatusResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`</ModifyApiKeyStatusResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseModifyApiKeyStatusResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
}
