// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMarketSkillCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMarketSkillCredentialResponseBody
	GetCode() *string
	SetData(v *GetMarketSkillCredentialResponseBodyData) *GetMarketSkillCredentialResponseBody
	GetData() *GetMarketSkillCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetMarketSkillCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMarketSkillCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMarketSkillCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMarketSkillCredentialResponseBody
	GetSuccess() *bool
}

type GetMarketSkillCredentialResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMarketSkillCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMarketSkillCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMarketSkillCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetMarketSkillCredentialResponseBody) GetCode() *string                    { return s.Code }
func (s *GetMarketSkillCredentialResponseBody) SetCode(v string) *GetMarketSkillCredentialResponseBody {
	s.Code = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBody) GetData() *GetMarketSkillCredentialResponseBodyData {
	return s.Data
}
func (s *GetMarketSkillCredentialResponseBody) SetData(v *GetMarketSkillCredentialResponseBodyData) *GetMarketSkillCredentialResponseBody {
	s.Data = v
	return s
}
func (s *GetMarketSkillCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}
func (s *GetMarketSkillCredentialResponseBody) SetHttpStatusCode(v int32) *GetMarketSkillCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBody) GetMessage() *string  { return s.Message }
func (s *GetMarketSkillCredentialResponseBody) SetMessage(v string) *GetMarketSkillCredentialResponseBody {
	s.Message = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBody) GetRequestId() *string { return s.RequestId }
func (s *GetMarketSkillCredentialResponseBody) SetRequestId(v string) *GetMarketSkillCredentialResponseBody {
	s.RequestId = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBody) GetSuccess() *bool { return s.Success }
func (s *GetMarketSkillCredentialResponseBody) SetSuccess(v bool) *GetMarketSkillCredentialResponseBody {
	s.Success = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMarketSkillCredentialResponseBodyData struct {
	// OssUrl is the legacy field name; backend may return Url instead (pre-release returns Data.Url).
	OssUrl      *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePath *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
}

func (s GetMarketSkillCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMarketSkillCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMarketSkillCredentialResponseBodyData) GetOssUrl() *string { return s.OssUrl }
func (s *GetMarketSkillCredentialResponseBodyData) SetOssUrl(v string) *GetMarketSkillCredentialResponseBodyData {
	s.OssUrl = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBodyData) GetUrl() *string { return s.Url }
func (s *GetMarketSkillCredentialResponseBodyData) SetUrl(v string) *GetMarketSkillCredentialResponseBodyData {
	s.Url = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBodyData) GetOssBucket() *string   { return s.OssBucket }
func (s *GetMarketSkillCredentialResponseBodyData) SetOssBucket(v string) *GetMarketSkillCredentialResponseBodyData {
	s.OssBucket = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBodyData) GetOssFilePath() *string  { return s.OssFilePath }
func (s *GetMarketSkillCredentialResponseBodyData) SetOssFilePath(v string) *GetMarketSkillCredentialResponseBodyData {
	s.OssFilePath = &v
	return s
}
func (s *GetMarketSkillCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
