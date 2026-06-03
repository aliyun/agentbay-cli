// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseUpdateMarketSkillResponse_JSONDataAsString(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","HttpStatusCode":200,"Data":"sk-abc123","RequestId":"R1","Success":true}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseUpdateMarketSkillResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.NotNil(t, out.Body.Data.SkillId)
	require.Equal(t, "sk-abc123", *out.Body.Data.SkillId)
	require.NotNil(t, out.Body.RequestId)
	require.Equal(t, "R1", *out.Body.RequestId)
}

func TestParseUpdateMarketSkillResponse_JSONDataAsObject(t *testing.T) {
	t.Parallel()
	body := `{"Code":"OK","Data":{"SkillId":"sk-obj"},"Success":true}`
	res := map[string]interface{}{"body": body}
	out, err := parseUpdateMarketSkillResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.NotNil(t, out.Body.Data.SkillId)
	require.Equal(t, "sk-obj", *out.Body.Data.SkillId)
}

func TestParseUpdateMarketSkillResponse_XMLDataString(t *testing.T) {
	t.Parallel()
	body := `<?xml version="1.0"?><UpdateMarketSkillResponse><Data>sk-xml</Data><RequestId>XR</RequestId><Success>true</Success></UpdateMarketSkillResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseUpdateMarketSkillResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "sk-xml", *out.Body.Data.SkillId)
}
