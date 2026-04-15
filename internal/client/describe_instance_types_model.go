// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// DescribeInstanceTypesRequest - 查询实例规格请求
type DescribeInstanceTypesRequest struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s *DescribeInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s *DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstanceTypesRequest) SetImageId(v string) *DescribeInstanceTypesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstanceTypesRequest) Validate() error {
	return dara.Validate(s)
}

// DescribeInstanceTypesResponseBodyDataInstanceType - 实例规格项
type DescribeInstanceTypesResponseBodyDataInstanceType struct {
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	Cpu             *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Memory          *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	IsSelected      *bool   `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
}

func (s *DescribeInstanceTypesResponseBodyDataInstanceType) String() string {
	return dara.Prettify(s)
}

func (s *DescribeInstanceTypesResponseBodyDataInstanceType) GetAppInstanceType() *string {
	return s.AppInstanceType
}

func (s *DescribeInstanceTypesResponseBodyDataInstanceType) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstanceTypesResponseBodyDataInstanceType) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeInstanceTypesResponseBodyDataInstanceType) GetIsSelected() *bool {
	return s.IsSelected
}

// DescribeInstanceTypesResponseBody - 响应体
// Note: Data is directly an array of instance types
type DescribeInstanceTypesResponseBody struct {
	RequestId      string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           []*DescribeInstanceTypesResponseBodyDataInstanceType `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s *DescribeInstanceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s *DescribeInstanceTypesResponseBody) GetRequestId() *string {
	return &s.RequestId
}

func (s *DescribeInstanceTypesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeInstanceTypesResponseBody) GetData() []*DescribeInstanceTypesResponseBodyDataInstanceType {
	return s.Data
}

func (s *DescribeInstanceTypesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

// DescribeInstanceTypesResponse - 响应
type DescribeInstanceTypesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
	RawBody    string                           `json:"-"` // Store raw body for debugging
}

func (s DescribeInstanceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s *DescribeInstanceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTypesResponse) GetBody() *DescribeInstanceTypesResponseBody {
	return s.Body
}
