// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpImageResponseBody
	GetCode() *string
	SetData(v *CreateMcpImageResponseBodyData) *CreateMcpImageResponseBody
	GetData() *CreateMcpImageResponseBodyData
	SetHttpStatusCode(v int32) *CreateMcpImageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMcpImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMcpImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMcpImageResponseBody
	GetSuccess() *bool
}

type CreateMcpImageResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateMcpImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcpImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpImageResponseBody) GetData() *CreateMcpImageResponseBodyData {
	return s.Data
}

func (s *CreateMcpImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMcpImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpImageResponseBody) SetCode(v string) *CreateMcpImageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpImageResponseBody) SetData(v *CreateMcpImageResponseBodyData) *CreateMcpImageResponseBody {
	s.Data = v
	return s
}

func (s *CreateMcpImageResponseBody) SetHttpStatusCode(v int32) *CreateMcpImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMcpImageResponseBody) SetMessage(v string) *CreateMcpImageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpImageResponseBody) SetRequestId(v string) *CreateMcpImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpImageResponseBody) SetSuccess(v bool) *CreateMcpImageResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMcpImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcpImageResponseBodyData struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s CreateMcpImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMcpImageResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *CreateMcpImageResponseBodyData) SetImageId(v string) *CreateMcpImageResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *CreateMcpImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
