// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketSkillDetailResponseBodyData interface {
	dara.Model
	String() string
	GoString() string
	GetSkillId() *string
	SetSkillId(v string) *DescribeMarketSkillDetailResponseBodyData
	GetName() *string
	SetName(v string) *DescribeMarketSkillDetailResponseBodyData
	GetDescription() *string
	SetDescription(v string) *DescribeMarketSkillDetailResponseBodyData
}

type DescribeMarketSkillDetailResponseBodyData struct {
	SkillId     *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeMarketSkillDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketSkillDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMarketSkillDetailResponseBodyData) GetSkillId() *string { return s.SkillId }
func (s *DescribeMarketSkillDetailResponseBodyData) SetSkillId(v string) *DescribeMarketSkillDetailResponseBodyData {
	s.SkillId = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBodyData) GetName() *string { return s.Name }
func (s *DescribeMarketSkillDetailResponseBodyData) SetName(v string) *DescribeMarketSkillDetailResponseBodyData {
	s.Name = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBodyData) GetDescription() *string { return s.Description }
func (s *DescribeMarketSkillDetailResponseBodyData) SetDescription(v string) *DescribeMarketSkillDetailResponseBodyData {
	s.Description = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type iDescribeMarketSkillDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMarketSkillDetailResponseBody
	GetCode() *string
	SetData(v *DescribeMarketSkillDetailResponseBodyData) *DescribeMarketSkillDetailResponseBody
	GetData() *DescribeMarketSkillDetailResponseBodyData
	SetHttpStatusCode(v int32) *DescribeMarketSkillDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeMarketSkillDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMarketSkillDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMarketSkillDetailResponseBody
	GetSuccess() *bool
}

type DescribeMarketSkillDetailResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeMarketSkillDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMarketSkillDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketSkillDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMarketSkillDetailResponseBody) GetCode() *string { return s.Code }
func (s *DescribeMarketSkillDetailResponseBody) SetCode(v string) *DescribeMarketSkillDetailResponseBody {
	s.Code = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBody) GetData() *DescribeMarketSkillDetailResponseBodyData {
	return s.Data
}
func (s *DescribeMarketSkillDetailResponseBody) SetData(v *DescribeMarketSkillDetailResponseBodyData) *DescribeMarketSkillDetailResponseBody {
	s.Data = v
	return s
}
func (s *DescribeMarketSkillDetailResponseBody) GetHttpStatusCode() *int32 { return s.HttpStatusCode }
func (s *DescribeMarketSkillDetailResponseBody) SetHttpStatusCode(v int32) *DescribeMarketSkillDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBody) GetMessage() *string { return s.Message }
func (s *DescribeMarketSkillDetailResponseBody) SetMessage(v string) *DescribeMarketSkillDetailResponseBody {
	s.Message = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBody) GetRequestId() *string { return s.RequestId }
func (s *DescribeMarketSkillDetailResponseBody) SetRequestId(v string) *DescribeMarketSkillDetailResponseBody {
	s.RequestId = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBody) GetSuccess() *bool { return s.Success }
func (s *DescribeMarketSkillDetailResponseBody) SetSuccess(v bool) *DescribeMarketSkillDetailResponseBody {
	s.Success = &v
	return s
}
func (s *DescribeMarketSkillDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
