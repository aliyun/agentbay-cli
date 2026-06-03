// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseUnshareDockerRepoResponse_JSONWithNumericStatusCode(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":200,"Success":true,"Data":{"Revoked":true}}`
	res := map[string]interface{}{"body": body}
	resp, err := parseUnshareDockerRepoResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.Equal(t, "ok", resp.Body.GetCode())
	assert.Equal(t, "test-req-id", resp.Body.GetRequestId())
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.NotNil(t, resp.Body.Data)
	require.NotNil(t, resp.Body.Data.Revoked)
	assert.True(t, *resp.Body.Data.Revoked)
}

func TestParseUnshareDockerRepoResponse_JSONWithStringStatusCode(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":"200","Success":true,"Data":{"Revoked":true}}`
	res := map[string]interface{}{"body": body}
	resp, err := parseUnshareDockerRepoResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.NotNil(t, resp.Body.Data)
	assert.True(t, *resp.Body.Data.Revoked)
}

func TestParseUnshareDockerRepoResponse_XML(t *testing.T) {
	body := `<UnshareDockerRepoResponse><RequestId>xml-req-id</RequestId><HttpStatusCode>200</HttpStatusCode><Code>ok</Code><Success>true</Success><Message></Message><Data><Revoked>true</Revoked></Data></UnshareDockerRepoResponse>`
	res := map[string]interface{}{"body": body}
	resp, err := parseUnshareDockerRepoResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.Equal(t, "ok", resp.Body.GetCode())
	assert.Equal(t, "xml-req-id", resp.Body.GetRequestId())
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.NotNil(t, resp.Body.Data)
	require.NotNil(t, resp.Body.Data.Revoked)
	assert.True(t, *resp.Body.Data.Revoked)
}
