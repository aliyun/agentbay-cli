// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPlaygroundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopPlaygroundTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopPlaygroundTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopPlaygroundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopPlaygroundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopPlaygroundTaskResponseBody
	GetSuccess() *bool
}

type StopPlaygroundTaskResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopPlaygroundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPlaygroundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopPlaygroundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopPlaygroundTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopPlaygroundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopPlaygroundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPlaygroundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopPlaygroundTaskResponseBody) SetCode(v string) *StopPlaygroundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopPlaygroundTaskResponseBody) SetHttpStatusCode(v int32) *StopPlaygroundTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopPlaygroundTaskResponseBody) SetMessage(v string) *StopPlaygroundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopPlaygroundTaskResponseBody) SetRequestId(v string) *StopPlaygroundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPlaygroundTaskResponseBody) SetSuccess(v bool) *StopPlaygroundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopPlaygroundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
