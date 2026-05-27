// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseShareDockerRepoResponse_JSONWithNumericStatusCode(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":200,"Success":true,"Data":{"TargetAliUid":123456789,"OwnerAliUid":987654321,"AcrRepoName":"my-repo","Status":"Sharing"}}`
	res := map[string]interface{}{"body": body}
	resp, err := parseShareDockerRepoResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.Equal(t, "ok", resp.Body.GetCode())
	assert.Equal(t, "test-req-id", resp.Body.GetRequestId())
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.NotNil(t, resp.Body.Data)
	assert.Equal(t, int64(123456789), *resp.Body.Data.TargetAliUid)
	assert.Equal(t, int64(987654321), *resp.Body.Data.OwnerAliUid)
	assert.Equal(t, "my-repo", *resp.Body.Data.AcrRepoName)
	assert.Equal(t, "Sharing", *resp.Body.Data.Status)
}

func TestParseShareDockerRepoResponse_JSONWithStringStatusCode(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":"200","Success":true,"Data":{"TargetAliUid":123456789,"OwnerAliUid":987654321,"AcrRepoName":"my-repo","Status":"Sharing"}}`
	res := map[string]interface{}{"body": body}
	resp, err := parseShareDockerRepoResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.NotNil(t, resp.Body.Data)
	assert.Equal(t, int64(123456789), *resp.Body.Data.TargetAliUid)
}

func TestParseShareDockerRepoResponse_XML(t *testing.T) {
	body := `<ShareDockerRepoResponse><RequestId>xml-req-id</RequestId><HttpStatusCode>200</HttpStatusCode><Code>ok</Code><Success>true</Success><Message></Message><Data><TargetAliUid>123456789</TargetAliUid><OwnerAliUid>987654321</OwnerAliUid><AcrRepoName>my-repo</AcrRepoName><Status>Sharing</Status></Data></ShareDockerRepoResponse>`
	res := map[string]interface{}{"body": body}
	resp, err := parseShareDockerRepoResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.Equal(t, "ok", resp.Body.GetCode())
	assert.Equal(t, "xml-req-id", resp.Body.GetRequestId())
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.NotNil(t, resp.Body.Data)
	assert.Equal(t, int64(123456789), *resp.Body.Data.TargetAliUid)
	assert.Equal(t, "Sharing", *resp.Body.Data.Status)
}
