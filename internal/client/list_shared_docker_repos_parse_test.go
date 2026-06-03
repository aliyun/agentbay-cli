// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseListSharedDockerReposResponse_JSONWithNumericStatusCode(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":200,"Success":true,"Data":[{"PeerAliUid":123456789,"Status":"Sharing"},{"PeerAliUid":987654321,"Status":"Pending"}]}`
	res := map[string]interface{}{"body": body}
	resp, err := parseListSharedDockerReposResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.Equal(t, "ok", resp.Body.GetCode())
	assert.Equal(t, "test-req-id", resp.Body.GetRequestId())
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.Len(t, resp.Body.Data, 2)
	assert.Equal(t, int64(123456789), *resp.Body.Data[0].PeerAliUid)
	assert.Equal(t, "Sharing", *resp.Body.Data[0].Status)
	assert.Equal(t, int64(987654321), *resp.Body.Data[1].PeerAliUid)
}

func TestParseListSharedDockerReposResponse_JSONWithStringStatusCode(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":"200","Success":true,"Data":[{"PeerAliUid":123456789,"Status":"Sharing"}]}`
	res := map[string]interface{}{"body": body}
	resp, err := parseListSharedDockerReposResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.Len(t, resp.Body.Data, 1)
	assert.Equal(t, int64(123456789), *resp.Body.Data[0].PeerAliUid)
}

func TestParseListSharedDockerReposResponse_EmptyData(t *testing.T) {
	body := `{"Code":"ok","Message":"","RequestId":"test-req-id","HttpStatusCode":200,"Success":true,"Data":[]}`
	res := map[string]interface{}{"body": body}
	resp, err := parseListSharedDockerReposResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.NotNil(t, resp.Body.Data)
	assert.Len(t, resp.Body.Data, 0)
}

func TestParseListSharedDockerReposResponse_XML(t *testing.T) {
	body := `<ListSharedDockerReposResponse><RequestId>xml-req-id</RequestId><HttpStatusCode>200</HttpStatusCode><Code>ok</Code><Success>true</Success><Message></Message><Data><object><PeerAliUid>123456789</PeerAliUid><Status>Sharing</Status></object></Data></ListSharedDockerReposResponse>`
	res := map[string]interface{}{"body": body}
	resp, err := parseListSharedDockerReposResponse(res)
	require.NoError(t, err)
	require.NotNil(t, resp.Body)
	assert.Equal(t, "ok", resp.Body.GetCode())
	assert.Equal(t, "xml-req-id", resp.Body.GetRequestId())
	require.NotNil(t, resp.Body.HttpStatusCode)
	assert.Equal(t, int32(200), *resp.Body.HttpStatusCode)
	require.Len(t, resp.Body.Data, 1)
	assert.Equal(t, int64(123456789), *resp.Body.Data[0].PeerAliUid)
	assert.Equal(t, "Sharing", *resp.Body.Data[0].Status)
}
