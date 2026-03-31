// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/dara"
	"github.com/stretchr/testify/require"
)

func TestParseGetDockerFileStoreCredentialResponse_XML(t *testing.T) {
	xmlBody := `<?xml version='1.0' encoding='UTF-8'?><GetDockerFileStoreCredentialResponse>` +
		`<RequestId>rid-1</RequestId><HttpStatusCode>200</HttpStatusCode>` +
		`<Data><OssUrl>https://oss.example/presign</OssUrl><TaskId>task-9</TaskId></Data>` +
		`<Code>ok</Code><Success>true</Success></GetDockerFileStoreCredentialResponse>`
	res := map[string]interface{}{
		"body":       xmlBody,
		"statusCode": 200,
	}
	out, err := parseGetDockerFileStoreCredentialResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "https://oss.example/presign", dara.StringValue(out.Body.Data.OssUrl))
	require.Equal(t, "task-9", dara.StringValue(out.Body.Data.TaskId))
	require.True(t, dara.BoolValue(out.Body.Success))
}

func TestParseCreateDockerImageTaskResponse_XML(t *testing.T) {
	xmlBody := `<?xml version='1.0'?><CreateDockerImageTaskResponse>` +
		`<RequestId>r1</RequestId><HttpStatusCode>200</HttpStatusCode>` +
		`<Data><TaskId>task-x</TaskId></Data><Code>ok</Code><Success>true</Success></CreateDockerImageTaskResponse>`
	res := map[string]interface{}{"body": xmlBody, "statusCode": 200}
	out, err := parseCreateDockerImageTaskResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "task-x", dara.StringValue(out.Body.Data.TaskId))
}

func TestParseGetDockerFileStoreCredentialResponse_JSON(t *testing.T) {
	jsonBody := `{"RequestId":"r2","HttpStatusCode":200,"Code":"ok","Success":true,"Data":{"OssUrl":"https://x","TaskId":"t1"}}`
	res := map[string]interface{}{
		"body":       jsonBody,
		"statusCode": 200,
	}
	out, err := parseGetDockerFileStoreCredentialResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "https://x", dara.StringValue(out.Body.Data.OssUrl))
	require.Equal(t, "t1", dara.StringValue(out.Body.Data.TaskId))
}
