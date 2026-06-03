// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCreateTagResponse_JSONHttpStatusCodeAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":"200","RequestId":"R1","Success":true,"Data":[{"TagName":"tag1","TagId":"1"}]}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseCreateTagResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", out.Body.GetRequestId())
	require.True(t, out.Body.GetSuccess())
	require.Len(t, out.Body.Data, 1)
	require.Equal(t, "tag1", *out.Body.Data[0].TagName)
	require.Equal(t, "1", *out.Body.Data[0].TagId)
}

func TestParseCreateTagResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":200,"RequestId":"R2","Success":true,"Data":[{"TagName":"tag2","TagId":"2"}]}`
	res := map[string]interface{}{"body": body}
	out, err := parseCreateTagResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Len(t, out.Body.Data, 1)
	require.Equal(t, "tag2", *out.Body.Data[0].TagName)
}

func TestParseCreateTagResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<CreateTagResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`<Data>` +
		`<TagName>xmltag</TagName>` +
		`<TagId>3</TagId>` +
		`</Data>` +
		`</CreateTagResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseCreateTagResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
	require.Len(t, out.Body.Data, 1)
	require.Equal(t, "xmltag", *out.Body.Data[0].TagName)
	require.Equal(t, "3", *out.Body.Data[0].TagId)
}
