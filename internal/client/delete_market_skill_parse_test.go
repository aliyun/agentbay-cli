// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDeleteMarketSkillResponse_JSONHttpStatusCodeAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"ok","HttpStatusCode":"200","RequestId":"R1","Success":true,"Data":true}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDeleteMarketSkillResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", out.Body.GetRequestId())
	require.True(t, out.Body.GetSuccess())
	require.NotNil(t, out.Body.Data)
	require.True(t, *out.Body.Data)
}

func TestParseDeleteMarketSkillResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"Code":"ok","HttpStatusCode":200,"RequestId":"R2","Success":true,"Data":true}`
	res := map[string]interface{}{"body": body}
	out, err := parseDeleteMarketSkillResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
}

func TestParseDeleteMarketSkillResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<DeleteMarketSkillResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>ok</Code>` +
		`<Success>true</Success>` +
		`<Data>true</Data>` +
		`</DeleteMarketSkillResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseDeleteMarketSkillResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
}
