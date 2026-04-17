// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// DescribeOfficeSitesRequest - 查询办公网络请求
type DescribeOfficeSitesRequest struct {
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	RegionName     *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s *DescribeOfficeSitesRequest) String() string {
	return dara.Prettify(s)
}

func (s *DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteType(v string) *DescribeOfficeSitesRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeOfficeSitesRequest) SetRegionName(v string) *DescribeOfficeSitesRequest {
	s.RegionName = &v
	return s
}

func (s *DescribeOfficeSitesRequest) Validate() error {
	return dara.Validate(s)
}

// DescribeOfficeSitesResponseBodyData - 响应数据
type DescribeOfficeSitesResponseBodyData struct {
	OfficeSiteId *string  `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	DnsAddress   []string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty"`
}

func (s *DescribeOfficeSitesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s *DescribeOfficeSitesResponseBodyData) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSitesResponseBodyData) GetDnsAddress() []string {
	return s.DnsAddress
}

// DescribeOfficeSitesResponseBody - 响应体
type DescribeOfficeSitesResponseBody struct {
	RequestId      string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *DescribeOfficeSitesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s *DescribeOfficeSitesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *DescribeOfficeSitesResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *DescribeOfficeSitesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeOfficeSitesResponseBody) GetData() *DescribeOfficeSitesResponseBodyData {
	return s.Data
}

func (s *DescribeOfficeSitesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeOfficeSitesResponseBody) GetSuccess() *bool {
	return s.Success
}

// DescribeOfficeSitesResponse - 响应
type DescribeOfficeSitesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOfficeSitesResponseBody      `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                               `json:"-"` // Store raw body for debugging
}

func (s DescribeOfficeSitesResponse) String() string {
	return dara.Prettify(s)
}

func (s *DescribeOfficeSitesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOfficeSitesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOfficeSitesResponse) GetBody() *DescribeOfficeSitesResponseBody {
	return s.Body
}
