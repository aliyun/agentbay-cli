// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDescribeMarketSkillDetailResponse_JSONWithTenantTags(t *testing.T) {
	t.Parallel()
	body := `{"Code":"ok","HttpStatusCode":200,"RequestId":"R1","Success":true,"Data":{"SkillId":"sk-1","Name":"My Skill","FileUrl":"https://example.com/skill.zip","Description":"A skill","TenantTags":["tag1","tag2"]}}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeMarketSkillDetailResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "sk-1", *out.Body.Data.SkillId)
	require.Equal(t, "My Skill", *out.Body.Data.Name)
	require.Equal(t, "https://example.com/skill.zip", *out.Body.Data.FileUrl)
	require.Equal(t, []string{"tag1", "tag2"}, out.Body.Data.TenantTags)
	require.Equal(t, "R1", *out.Body.RequestId)
}

func TestParseDescribeMarketSkillDetailResponse_JSONNoTenantTags(t *testing.T) {
	t.Parallel()
	body := `{"Code":"ok","HttpStatusCode":200,"RequestId":"R2","Success":true,"Data":{"SkillId":"sk-2","Name":"Skill2","Description":"Desc"}}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeMarketSkillDetailResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "sk-2", *out.Body.Data.SkillId)
	require.Empty(t, out.Body.Data.TenantTags)
}

func TestParseDescribeMarketSkillDetailResponse_JSONSingleTag(t *testing.T) {
	t.Parallel()
	// Single tag in TenantTags array.
	body := `{"Code":"ok","HttpStatusCode":200,"RequestId":"R3","Data":{"SkillId":"sk-3","TenantTags":["alpha"]}}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeMarketSkillDetailResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "sk-3", *out.Body.Data.SkillId)
	require.Equal(t, []string{"alpha"}, out.Body.Data.TenantTags)
}

func TestParseDescribeMarketSkillDetailResponse_JSONNoFileUrl(t *testing.T) {
	t.Parallel()
	// FileUrl absent in response - should be nil.
	body := `{"Code":"ok","HttpStatusCode":200,"RequestId":"R5","Data":{"SkillId":"sk-5","Name":"NoUrl Skill","Description":"desc"}}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeMarketSkillDetailResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "sk-5", *out.Body.Data.SkillId)
	require.Nil(t, out.Body.Data.FileUrl)
}

func TestParseDescribeMarketSkillDetailResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<?xml version="1.0"?><DescribeMarketSkillDetailResponse><RequestId>XR1</RequestId><Code>ok</Code><Success>true</Success><Data><SkillId>sk-xml</SkillId><Name>XML Skill</Name><FileUrl>https://example.com/xml.zip</FileUrl><Description>XML desc</Description><TenantTags>tagA</TenantTags><TenantTags>tagB</TenantTags></Data></DescribeMarketSkillDetailResponse>`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeMarketSkillDetailResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body.Data)
	require.Equal(t, "sk-xml", *out.Body.Data.SkillId)
	require.Equal(t, "https://example.com/xml.zip", *out.Body.Data.FileUrl)
	require.Equal(t, []string{"tagA", "tagB"}, out.Body.Data.TenantTags)
	require.Equal(t, "XR1", *out.Body.RequestId)
}
