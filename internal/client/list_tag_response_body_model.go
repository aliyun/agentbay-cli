// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type ListTagResponseBodyDataItem struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListTagResponseBodyDataItem) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponseBodyDataItem) GoString() string {
	return s.String()
}

func (s *ListTagResponseBodyDataItem) GetTagName() *string { return s.TagName }
func (s *ListTagResponseBodyDataItem) SetTagName(v string) *ListTagResponseBodyDataItem {
	s.TagName = &v
	return s
}
func (s *ListTagResponseBodyDataItem) GetTagId() *string { return s.TagId }
func (s *ListTagResponseBodyDataItem) SetTagId(v string) *ListTagResponseBodyDataItem {
	s.TagId = &v
	return s
}

type iListTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTagResponseBody
	GetCode() string
	SetHttpStatusCode(v int32) *ListTagResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTagResponseBody
	GetMessage() string
	SetRequestId(v string) *ListTagResponseBody
	GetRequestId() string
	SetSuccess(v bool) *ListTagResponseBody
	GetSuccess() bool
}

type ListTagResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []ListTagResponseBodyDataItem `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}
func (s *ListTagResponseBody) SetCode(v string) *ListTagResponseBody {
	s.Code = &v
	return s
}
func (s *ListTagResponseBody) GetHttpStatusCode() *int32 { return s.HttpStatusCode }
func (s *ListTagResponseBody) SetHttpStatusCode(v int32) *ListTagResponseBody {
	s.HttpStatusCode = &v
	return s
}
func (s *ListTagResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}
func (s *ListTagResponseBody) SetMessage(v string) *ListTagResponseBody {
	s.Message = &v
	return s
}
func (s *ListTagResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
func (s *ListTagResponseBody) SetRequestId(v string) *ListTagResponseBody {
	s.RequestId = &v
	return s
}
func (s *ListTagResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}
func (s *ListTagResponseBody) SetSuccess(v bool) *ListTagResponseBody {
	s.Success = &v
	return s
}
func (s *ListTagResponseBody) Validate() error {
	return dara.Validate(s)
}
