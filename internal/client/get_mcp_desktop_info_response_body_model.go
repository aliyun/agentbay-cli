// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpDesktopInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpDesktopInfoResponseBody
	GetCode() *string
	SetData(v *GetMcpDesktopInfoResponseBodyData) *GetMcpDesktopInfoResponseBody
	GetData() *GetMcpDesktopInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpDesktopInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpDesktopInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpDesktopInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpDesktopInfoResponseBody
	GetSuccess() *bool
}

type GetMcpDesktopInfoResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMcpDesktopInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpDesktopInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpDesktopInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpDesktopInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpDesktopInfoResponseBody) GetData() *GetMcpDesktopInfoResponseBodyData {
	return s.Data
}

func (s *GetMcpDesktopInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpDesktopInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpDesktopInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpDesktopInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpDesktopInfoResponseBody) SetCode(v string) *GetMcpDesktopInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpDesktopInfoResponseBody) SetData(v *GetMcpDesktopInfoResponseBodyData) *GetMcpDesktopInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpDesktopInfoResponseBody) SetHttpStatusCode(v int32) *GetMcpDesktopInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpDesktopInfoResponseBody) SetMessage(v string) *GetMcpDesktopInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpDesktopInfoResponseBody) SetRequestId(v string) *GetMcpDesktopInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpDesktopInfoResponseBody) SetSuccess(v bool) *GetMcpDesktopInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpDesktopInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcpDesktopInfoResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMcpDesktopInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpDesktopInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpDesktopInfoResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetMcpDesktopInfoResponseBodyData) SetUrl(v string) *GetMcpDesktopInfoResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetMcpDesktopInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
