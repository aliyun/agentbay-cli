// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketSkillDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMarketSkillDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMarketSkillDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMarketSkillDetailResponseBody) *DescribeMarketSkillDetailResponse
	GetBody() *DescribeMarketSkillDetailResponseBody
}

type DescribeMarketSkillDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMarketSkillDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMarketSkillDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketSkillDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeMarketSkillDetailResponse) GetHeaders() map[string]*string    { return s.Headers }
func (s *DescribeMarketSkillDetailResponse) SetHeaders(v map[string]*string) *DescribeMarketSkillDetailResponse {
	s.Headers = v
	return s
}
func (s *DescribeMarketSkillDetailResponse) GetStatusCode() *int32 { return s.StatusCode }
func (s *DescribeMarketSkillDetailResponse) SetStatusCode(v int32) *DescribeMarketSkillDetailResponse {
	s.StatusCode = &v
	return s
}
func (s *DescribeMarketSkillDetailResponse) GetBody() *DescribeMarketSkillDetailResponseBody {
	return s.Body
}
func (s *DescribeMarketSkillDetailResponse) SetBody(v *DescribeMarketSkillDetailResponseBody) *DescribeMarketSkillDetailResponse {
	s.Body = v
	return s
}
func (s *DescribeMarketSkillDetailResponse) Validate() error {
	return dara.Validate(s)
}
