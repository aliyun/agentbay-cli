// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMarketSkillResponseBodyData interface {
	dara.Model
	String() string
	GoString() string
	SetSkillId(v string) *CreateMarketSkillResponseBodyData
	GetSkillId() *string
}

type CreateMarketSkillResponseBodyData struct {
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s CreateMarketSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillResponseBodyData) GetSkillId() *string { return s.SkillId }
func (s *CreateMarketSkillResponseBodyData) SetSkillId(v string) *CreateMarketSkillResponseBodyData {
	s.SkillId = &v
	return s
}
func (s *CreateMarketSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type iCreateMarketSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMarketSkillResponseBody
	GetCode() *string
	SetData(v *CreateMarketSkillResponseBodyData) *CreateMarketSkillResponseBody
	GetData() *CreateMarketSkillResponseBodyData
	SetHttpStatusCode(v int32) *CreateMarketSkillResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMarketSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMarketSkillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMarketSkillResponseBody
	GetSuccess() *bool
}

type CreateMarketSkillResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateMarketSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMarketSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMarketSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMarketSkillResponseBody) GetCode() *string                    { return s.Code }
func (s *CreateMarketSkillResponseBody) SetCode(v string) *CreateMarketSkillResponseBody {
	s.Code = &v
	return s
}
func (s *CreateMarketSkillResponseBody) GetData() *CreateMarketSkillResponseBodyData {
	return s.Data
}
func (s *CreateMarketSkillResponseBody) SetData(v *CreateMarketSkillResponseBodyData) *CreateMarketSkillResponseBody {
	s.Data = v
	return s
}
func (s *CreateMarketSkillResponseBody) GetHttpStatusCode() *int32 { return s.HttpStatusCode }
func (s *CreateMarketSkillResponseBody) SetHttpStatusCode(v int32) *CreateMarketSkillResponseBody {
	s.HttpStatusCode = &v
	return s
}
func (s *CreateMarketSkillResponseBody) GetMessage() *string  { return s.Message }
func (s *CreateMarketSkillResponseBody) SetMessage(v string) *CreateMarketSkillResponseBody {
	s.Message = &v
	return s
}
func (s *CreateMarketSkillResponseBody) GetRequestId() *string { return s.RequestId }
func (s *CreateMarketSkillResponseBody) SetRequestId(v string) *CreateMarketSkillResponseBody {
	s.RequestId = &v
	return s
}
func (s *CreateMarketSkillResponseBody) GetSuccess() *bool { return s.Success }
func (s *CreateMarketSkillResponseBody) SetSuccess(v bool) *CreateMarketSkillResponseBody {
	s.Success = &v
	return s
}
func (s *CreateMarketSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
