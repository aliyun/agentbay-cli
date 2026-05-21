// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDescribeWarmUpStatusOpenResponse_JSONInt32AsString(t *testing.T) {
	t.Parallel()
	body := `{
		"Code":"OK",
		"HttpStatusCode":"200",
		"RequestId":"R1",
		"Success":true,
		"Data":{
			"MaxSessionNumLimit":"10",
			"TotalUsedSessionQuota":"3",
			"AvailableSessionQuota":"7",
			"MaxImageCount":"5",
			"CurrentImageCount":"2",
			"Images":[
				{"ImageId":"imgc-xxx","TotalMaxSize":"100","GroupCount":"3","AvailableInstanceSize":"80"},
				{"ImageId":"imgc-yyy","TotalMaxSize":"50","GroupCount":"1","AvailableInstanceSize":"40"}
			]
		}
	}`
	res := map[string]interface{}{"body": body, "statusCode": 200}
	out, err := parseDescribeWarmUpStatusOpenResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R1", *out.Body.GetRequestId())
	require.True(t, *out.Body.GetSuccess())
	require.NotNil(t, out.Body.Data)
	require.Equal(t, int32(10), out.Body.Data.GetMaxSessionNumLimit())
	require.Equal(t, int32(3), out.Body.Data.GetTotalUsedSessionQuota())
	require.Equal(t, int32(7), out.Body.Data.GetAvailableSessionQuota())
	require.Equal(t, int32(5), out.Body.Data.GetMaxImageCount())
	require.Equal(t, int32(2), out.Body.Data.GetCurrentImageCount())
	require.Len(t, out.Body.Data.Images, 2)
	require.Equal(t, "imgc-xxx", out.Body.Data.Images[0].GetImageId())
	require.Equal(t, int32(100), out.Body.Data.Images[0].GetTotalMaxSize())
	require.Equal(t, int32(3), out.Body.Data.Images[0].GetGroupCount())
	require.Equal(t, int32(80), out.Body.Data.Images[0].GetAvailableInstanceSize())
	require.Equal(t, "imgc-yyy", out.Body.Data.Images[1].GetImageId())
	require.Equal(t, int32(50), out.Body.Data.Images[1].GetTotalMaxSize())
	require.Equal(t, int32(1), out.Body.Data.Images[1].GetGroupCount())
	require.Equal(t, int32(40), out.Body.Data.Images[1].GetAvailableInstanceSize())
}

func TestParseDescribeWarmUpStatusOpenResponse_JSONInt32AsNumber(t *testing.T) {
	t.Parallel()
	body := `{
		"Code":"OK",
		"HttpStatusCode":200,
		"RequestId":"R2",
		"Success":true,
		"Data":{
			"MaxSessionNumLimit":10,
			"TotalUsedSessionQuota":3,
			"AvailableSessionQuota":7,
			"MaxImageCount":5,
			"CurrentImageCount":2,
			"Images":[
				{"ImageId":"imgc-xxx","TotalMaxSize":100,"GroupCount":3,"AvailableInstanceSize":80}
			]
		}
	}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeWarmUpStatusOpenResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.Equal(t, "R2", *out.Body.GetRequestId())
	require.NotNil(t, out.Body.Data)
	require.Equal(t, int32(10), out.Body.Data.GetMaxSessionNumLimit())
	require.Equal(t, int32(3), out.Body.Data.GetTotalUsedSessionQuota())
	require.Equal(t, int32(7), out.Body.Data.GetAvailableSessionQuota())
	require.Equal(t, int32(5), out.Body.Data.GetMaxImageCount())
	require.Equal(t, int32(2), out.Body.Data.GetCurrentImageCount())
	require.Len(t, out.Body.Data.Images, 1)
	require.Equal(t, int32(100), out.Body.Data.Images[0].GetTotalMaxSize())
	require.Equal(t, int32(3), out.Body.Data.Images[0].GetGroupCount())
	require.Equal(t, int32(80), out.Body.Data.Images[0].GetAvailableInstanceSize())
}

func TestParseDescribeWarmUpStatusOpenResponse_XML(t *testing.T) {
	t.Parallel()
	body := `<DescribeWarmUpStatusOpenResponse>` +
		`<RequestId>R3</RequestId>` +
		`<HttpStatusCode>200</HttpStatusCode>` +
		`<Code>OK</Code>` +
		`<Success>true</Success>` +
		`<Data>` +
		`<MaxSessionNumLimit>10</MaxSessionNumLimit>` +
		`<TotalUsedSessionQuota>3</TotalUsedSessionQuota>` +
		`<AvailableSessionQuota>7</AvailableSessionQuota>` +
		`<MaxImageCount>5</MaxImageCount>` +
		`<CurrentImageCount>2</CurrentImageCount>` +
		`<Images>` +
		`<Image>` +
		`<ImageId>imgc-xxx</ImageId>` +
		`<TotalMaxSize>100</TotalMaxSize>` +
		`<GroupCount>3</GroupCount>` +
		`<AvailableInstanceSize>80</AvailableInstanceSize>` +
		`</Image>` +
		`</Images>` +
		`</Data>` +
		`</DescribeWarmUpStatusOpenResponse>`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeWarmUpStatusOpenResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Equal(t, "R3", *out.Body.GetRequestId())
	require.NotNil(t, out.Body.HttpStatusCode)
	require.Equal(t, int32(200), *out.Body.HttpStatusCode)
	require.True(t, *out.Body.GetSuccess())
	require.NotNil(t, out.Body.Data)
	require.Equal(t, int32(10), out.Body.Data.GetMaxSessionNumLimit())
	require.Equal(t, int32(3), out.Body.Data.GetTotalUsedSessionQuota())
	require.Equal(t, int32(7), out.Body.Data.GetAvailableSessionQuota())
	require.Equal(t, int32(5), out.Body.Data.GetMaxImageCount())
	require.Equal(t, int32(2), out.Body.Data.GetCurrentImageCount())
	require.Len(t, out.Body.Data.Images, 1)
	require.Equal(t, "imgc-xxx", out.Body.Data.Images[0].GetImageId())
	require.Equal(t, int32(100), out.Body.Data.Images[0].GetTotalMaxSize())
	require.Equal(t, int32(3), out.Body.Data.Images[0].GetGroupCount())
	require.Equal(t, int32(80), out.Body.Data.Images[0].GetAvailableInstanceSize())
}

func TestParseDescribeWarmUpStatusOpenResponse_JSONEmptyImages(t *testing.T) {
	t.Parallel()
	body := `{
		"Code":"OK",
		"HttpStatusCode":200,
		"RequestId":"R4",
		"Success":true,
		"Data":{
			"MaxSessionNumLimit":10,
			"TotalUsedSessionQuota":0,
			"AvailableSessionQuota":10,
			"MaxImageCount":5,
			"CurrentImageCount":0,
			"Images":[]
		}
	}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeWarmUpStatusOpenResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.NotNil(t, out.Body.Data)
	require.Empty(t, out.Body.Data.Images)
}

func TestParseDescribeWarmUpStatusOpenResponse_JSONNullData(t *testing.T) {
	t.Parallel()
	body := `{
		"Code":"OK",
		"HttpStatusCode":200,
		"RequestId":"R5",
		"Success":true,
		"Data":null
	}`
	res := map[string]interface{}{"body": body}
	out, err := parseDescribeWarmUpStatusOpenResponse(res)
	require.NoError(t, err)
	require.NotNil(t, out.Body)
	require.Nil(t, out.Body.Data)
}
