// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDesktopReleaseTimingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDesktopReleaseTimingsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDesktopReleaseTimingsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDesktopReleaseTimingsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDesktopReleaseTimingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDesktopReleaseTimingsResponseBody
	GetSuccess() *bool
}

type UpdateDesktopReleaseTimingsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDesktopReleaseTimingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDesktopReleaseTimingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDesktopReleaseTimingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDesktopReleaseTimingsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDesktopReleaseTimingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDesktopReleaseTimingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDesktopReleaseTimingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDesktopReleaseTimingsResponseBody) SetCode(v string) *UpdateDesktopReleaseTimingsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponseBody) SetHttpStatusCode(v int32) *UpdateDesktopReleaseTimingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponseBody) SetMessage(v string) *UpdateDesktopReleaseTimingsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponseBody) SetRequestId(v string) *UpdateDesktopReleaseTimingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponseBody) SetSuccess(v bool) *UpdateDesktopReleaseTimingsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDesktopReleaseTimingsResponseBody) Validate() error {
	return dara.Validate(s)
}
