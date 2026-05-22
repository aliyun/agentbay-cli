// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDescribeKeyContentResponse_JSONHttpStatusCodeAsString(t *testing.T) {
	// Simulates the actual server response format where httpStatusCode is a string
	body := `{"code":"200","data":{"ApiKey":"akm-81872b44-11d9-4a8f-9817-3df33512605f","RequestId":"EE6F8DBD-AB01-16C1-90D9-0B402345511F"},"httpStatusCode":"200","requestId":"EE6F8DBD-AB01-16C1-90D9-0B402345511F","successResponse":true}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeKeyContentResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "200", out.Body.GetCode())
	require.Equal(t, "EE6F8DBD-AB01-16C1-90D9-0B402345511F", out.Body.GetRequestId())
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "akm-81872b44-11d9-4a8f-9817-3df33512605f", out.Body.Data.GetApiKey())
}

func TestParseDescribeKeyContentResponse_JSONHttpStatusCodeAsNumber(t *testing.T) {
	// Simulates a response where httpStatusCode is a number
	body := `{"code":"200","data":{"ApiKey":"akm-test-key","RequestId":"AAAABBBB-1111-2222-3333-444455556666"},"httpStatusCode":200,"requestId":"AAAABBBB-1111-2222-3333-444455556666","successResponse":true}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeKeyContentResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "akm-test-key", out.Body.Data.GetApiKey())
	require.Equal(t, "AAAABBBB-1111-2222-3333-444455556666", out.Body.Data.GetRequestId())
}

func TestParseDescribeKeyContentResponse_XML(t *testing.T) {
	body := `<DescribeKeyContentResponse>` +
		`<RequestId>XML-REQID-1234</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>200</Code>` +
		`<Success>true</Success>` +
		`<Message></Message>` +
		`<Data>` +
		`<ApiKey>akm-xml-test-key</ApiKey>` +
		`<RequestId>XML-DATA-REQID-5678</RequestId>` +
		`</Data>` +
		`</DescribeKeyContentResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeKeyContentResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "XML-REQID-1234", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "akm-xml-test-key", out.Body.Data.GetApiKey())
	require.Equal(t, "XML-DATA-REQID-5678", out.Body.Data.GetRequestId())
}
