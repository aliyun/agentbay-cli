// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCreateSimpleOfficeSiteResponse_JSONHttpStatusCodeAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":"200","RequestId":"R1","Success":true,"Data":"os-site-001"}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseCreateSimpleOfficeSiteResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", *out.Body.GetRequestId())
	require.NotNil(t, out.Body.Success)
	require.True(t, *out.Body.Success)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "os-site-001", *out.Body.Data)
}

func TestParseCreateSimpleOfficeSiteResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":200,"RequestId":"R2","Success":true,"Data":"os-site-002"}`
	res := map[string]interface{}{"body": body}
	out, err := parseCreateSimpleOfficeSiteResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "os-site-002", *out.Body.Data)
}

func TestParseCreateSimpleOfficeSiteResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<CreateSimpleOfficeSiteResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`<Data>os-site-003</Data>` +
		`</CreateSimpleOfficeSiteResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseCreateSimpleOfficeSiteResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", *out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.NotNil(t, out.Body.Success)
	require.True(t, *out.Body.Success)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "os-site-003", *out.Body.Data)
}
