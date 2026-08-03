// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDescribeImageReserveMinAmountResponse_JSONFullData(t *testing.T) {
	t.Parallel()
	body := `{
		"Code":"OK",
		"HttpStatusCode":200,
		"RequestId":"R1",
		"Success":true,
		"Data":{
			"Images":[
				{
					"ImageId":"imgc-test",
					"ResourceGroups":[
						{
							"ResourceGroupId":"rg-1",
							"AppInstanceGroupId":"pool-1",
							"ReserveMinAmount":5,
							"MaxAmount":1000,
							"ResourceGroupType":"DEFAULT",
							"Status":"PUBLISHED"
						}
					]
				}
			],
			"NextToken":"token-1"
		}
	}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeImageReserveMinAmountResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", out.Body.GetRequestId())
	require.True(t, out.Body.GetSuccess())
	require.Equal(t, "OK", out.Body.GetCode())

	data := out.Body.GetData()
	require.NotNil(t, data)
	require.Len(t, data.GetImages(), 1)
	require.Equal(t, "imgc-test", data.GetImages()[0].GetImageId())
	require.Len(t, data.GetImages()[0].GetResourceGroups(), 1)

	rg := data.GetImages()[0].GetResourceGroups()[0]
	require.Equal(t, "rg-1", rg.GetResourceGroupId())
	require.Equal(t, "pool-1", rg.GetAppInstanceGroupId())
	require.Equal(t, int32(5), rg.GetReserveMinAmount())
	require.Equal(t, int32(1000), rg.GetMaxAmount())
	require.Equal(t, "DEFAULT", rg.GetResourceGroupType())
	require.Equal(t, "PUBLISHED", rg.GetStatus())
	require.Equal(t, "token-1", data.GetNextToken())
}

func TestParseDescribeImageReserveMinAmountResponse_JSONFieldsAsString(t *testing.T) {
	t.Parallel()
	body := `{
		"Code":"OK",
		"HttpStatusCode":"200",
		"RequestId":"R2",
		"Success":true,
		"Data":{
			"Images":[
				{
					"ImageId":"imgc-test",
					"ResourceGroups":[
						{
							"ResourceGroupId":"rg-1",
							"AppInstanceGroupId":"pool-1",
							"ReserveMinAmount":"5",
							"MaxAmount":"1000",
							"ResourceGroupType":"DEFAULT",
							"Status":"PUBLISHED"
						}
					]
				}
			]
		}
	}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeImageReserveMinAmountResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R2", out.Body.GetRequestId())

	rg := out.Body.GetData().GetImages()[0].GetResourceGroups()[0]
	require.Equal(t, int32(5), rg.GetReserveMinAmount())
	require.Equal(t, int32(1000), rg.GetMaxAmount())
}

func TestParseDescribeImageReserveMinAmountResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<DescribeImageReserveMinAmountResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`<Data>` +
		`<Images>` +
		`<Image>` +
		`<ImageId>imgc-test</ImageId>` +
		`<ResourceGroups>` +
		`<ResourceGroup>` +
		`<ResourceGroupId>rg-1</ResourceGroupId>` +
		`<AppInstanceGroupId>pool-1</AppInstanceGroupId>` +
		`<ReserveMinAmount>5</ReserveMinAmount>` +
		`<MaxAmount>1000</MaxAmount>` +
		`<ResourceGroupType>DEFAULT</ResourceGroupType>` +
		`<Status>PUBLISHED</Status>` +
		`</ResourceGroup>` +
		`</ResourceGroups>` +
		`</Image>` +
		`</Images>` +
		`<NextToken>token-2</NextToken>` +
		`</Data>` +
		`</DescribeImageReserveMinAmountResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeImageReserveMinAmountResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, out.Body.GetSuccess())
	require.Equal(t, "OK", out.Body.GetCode())

	data := out.Body.GetData()
	require.NotNil(t, data)
	require.Equal(t, "token-2", data.GetNextToken())
	require.Len(t, data.GetImages(), 1)
	rg := data.GetImages()[0].GetResourceGroups()[0]
	require.Equal(t, int32(5), rg.GetReserveMinAmount())
	require.Equal(t, int32(1000), rg.GetMaxAmount())
}

func TestParseDescribeImageReserveMinAmountResponse_EmptyBody(t *testing.T) {
	t.Parallel()
	res := map[string]interface{}{"body": "", "statusCode": 200}
	out, err := parseDescribeImageReserveMinAmountResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "", out.Body.GetRequestId())
}
