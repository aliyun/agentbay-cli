// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type CreateTagResponseBodyDataItem struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateTagResponseBodyDataItem) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBodyDataItem) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBodyDataItem) GetTagName() *string { return s.TagName }
func (s *CreateTagResponseBodyDataItem) SetTagName(v string) *CreateTagResponseBodyDataItem {
	s.TagName = &v
	return s
}
func (s *CreateTagResponseBodyDataItem) GetTagId() *string { return s.TagId }
func (s *CreateTagResponseBodyDataItem) SetTagId(v string) *CreateTagResponseBodyDataItem {
	s.TagId = &v
	return s
}

type iCreateTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTagResponseBody
	GetCode() string
	SetHttpStatusCode(v int32) *CreateTagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTagResponseBody
	GetMessage() string
	SetRequestId(v string) *CreateTagResponseBody
	GetRequestId() string
	SetSuccess(v bool) *CreateTagResponseBody
	GetSuccess() bool
}

type CreateTagResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []CreateTagResponseBodyDataItem `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}
func (s *CreateTagResponseBody) SetCode(v string) *CreateTagResponseBody {
	s.Code = &v
	return s
}
func (s *CreateTagResponseBody) GetHttpStatusCode() *int32 { return s.HttpStatusCode }
func (s *CreateTagResponseBody) SetHttpStatusCode(v int32) *CreateTagResponseBody {
	s.HttpStatusCode = &v
	return s
}
func (s *CreateTagResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}
func (s *CreateTagResponseBody) SetMessage(v string) *CreateTagResponseBody {
	s.Message = &v
	return s
}
func (s *CreateTagResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
func (s *CreateTagResponseBody) SetRequestId(v string) *CreateTagResponseBody {
	s.RequestId = &v
	return s
}
func (s *CreateTagResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}
func (s *CreateTagResponseBody) SetSuccess(v bool) *CreateTagResponseBody {
	s.Success = &v
	return s
}
func (s *CreateTagResponseBody) Validate() error {
	return dara.Validate(s)
}
