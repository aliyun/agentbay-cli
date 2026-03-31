// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseGetDockerfileTemplateResponse_JSONNonEditLineNumAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":200,"Data":{"OssDownloadUrl":"https://x","NonEditLineNum":"5","DockerfileContent":"FROM a"},"RequestId":"R1","Success":true}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseGetDockerfileTemplateResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.NotNil(t, out.Body.Data.NonEditLineNum)
	require.Equal(t, int32(5), *out.Body.Data.NonEditLineNum)
	require.NotNil(t, out.Body.Data.OssDownloadUrl)
	require.Equal(t, "https://x", *out.Body.Data.OssDownloadUrl)
}

func TestParseGetDockerfileTemplateResponse_JSONNonEditLineNumAsNumber(t *testing.T) {
	t.Parallel()
	body := `{"Data":{"NonEditLineNum":12,"DockerfileContent":"FROM b"},"Success":true}`
	res := map[string]interface{}{"body": body}
	out, err := parseGetDockerfileTemplateResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.NotNil(t, out.Body.Data.NonEditLineNum)
	require.Equal(t, int32(12), *out.Body.Data.NonEditLineNum)
}

func TestParseGetDockerfileTemplateResponse_JSONNonEditLineNumNull(t *testing.T) {
	t.Parallel()
	body := `{"Data":{"NonEditLineNum":null,"DockerfileContent":"FROM c"}}`
	res := map[string]interface{}{"body": body}
	out, err := parseGetDockerfileTemplateResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.Nil(t, out.Body.Data.NonEditLineNum)
}
