// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// CreateSimpleOfficeSiteRequest - 创建简单办公网络请求
type CreateSimpleOfficeSiteRequest struct {
	VpcType           *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	OfficeSiteName    *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName        *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
}

func (s *CreateSimpleOfficeSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s *CreateSimpleOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteRequest) GetVpcType() *string {
	return s.VpcType
}

func (s *CreateSimpleOfficeSiteRequest) SetVpcType(v string) *CreateSimpleOfficeSiteRequest {
	s.VpcType = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *CreateSimpleOfficeSiteRequest) SetOfficeSiteName(v string) *CreateSimpleOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSimpleOfficeSiteRequest) SetVpcId(v string) *CreateSimpleOfficeSiteRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSimpleOfficeSiteRequest) SetRegionId(v string) *CreateSimpleOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *CreateSimpleOfficeSiteRequest) SetRegionName(v string) *CreateSimpleOfficeSiteRequest {
	s.RegionName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *CreateSimpleOfficeSiteRequest) SetDesktopAccessType(v string) *CreateSimpleOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) Validate() error {
	return dara.Validate(s)
}

// CreateSimpleOfficeSiteResponseBody - 响应体
// Note: Data is a string directly representing the OfficeSiteId
type CreateSimpleOfficeSiteResponseBody struct {
	RequestId      string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s *CreateSimpleOfficeSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *CreateSimpleOfficeSiteResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *CreateSimpleOfficeSiteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSimpleOfficeSiteResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSimpleOfficeSiteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSimpleOfficeSiteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSimpleOfficeSiteResponseBody) GetMessage() *string {
	return s.Message
}

// CreateSimpleOfficeSiteResponse - 响应
type CreateSimpleOfficeSiteResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimpleOfficeSiteResponseBody   `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                                `json:"-"` // Store raw body for debugging
}

func (s CreateSimpleOfficeSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s *CreateSimpleOfficeSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimpleOfficeSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimpleOfficeSiteResponse) GetBody() *CreateSimpleOfficeSiteResponseBody {
	return s.Body
}
