// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpConfigResponseBody
	GetCode() *string
	SetData(v *GetMcpConfigResponseBodyData) *GetMcpConfigResponseBody
	GetData() *GetMcpConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpConfigResponseBody
	GetSuccess() *bool
}

type GetMcpConfigResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMcpConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpConfigResponseBody) GetData() *GetMcpConfigResponseBodyData {
	return s.Data
}

func (s *GetMcpConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpConfigResponseBody) SetCode(v string) *GetMcpConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpConfigResponseBody) SetData(v *GetMcpConfigResponseBodyData) *GetMcpConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpConfigResponseBody) SetHttpStatusCode(v int32) *GetMcpConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpConfigResponseBody) SetMessage(v string) *GetMcpConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpConfigResponseBody) SetRequestId(v string) *GetMcpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpConfigResponseBody) SetSuccess(v bool) *GetMcpConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcpConfigResponseBodyData struct {
	McpConfigInfo *string `json:"McpConfigInfo,omitempty" xml:"McpConfigInfo,omitempty"`
}

func (s GetMcpConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpConfigResponseBodyData) GetMcpConfigInfo() *string {
	return s.McpConfigInfo
}

func (s *GetMcpConfigResponseBodyData) SetMcpConfigInfo(v string) *GetMcpConfigResponseBodyData {
	s.McpConfigInfo = &v
	return s
}

func (s *GetMcpConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
