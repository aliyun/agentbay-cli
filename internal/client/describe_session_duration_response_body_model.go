// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionDurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSessionDurationResponseBody
	GetCode() *string
	SetData(v *DescribeSessionDurationResponseBodyData) *DescribeSessionDurationResponseBody
	GetData() *DescribeSessionDurationResponseBodyData
	SetHttpStatusCode(v int32) *DescribeSessionDurationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeSessionDurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSessionDurationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSessionDurationResponseBody
	GetSuccess() *bool
}

type DescribeSessionDurationResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeSessionDurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSessionDurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionDurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSessionDurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSessionDurationResponseBody) GetData() *DescribeSessionDurationResponseBodyData {
	return s.Data
}

func (s *DescribeSessionDurationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeSessionDurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSessionDurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSessionDurationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSessionDurationResponseBody) SetCode(v string) *DescribeSessionDurationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSessionDurationResponseBody) SetData(v *DescribeSessionDurationResponseBodyData) *DescribeSessionDurationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSessionDurationResponseBody) SetHttpStatusCode(v int32) *DescribeSessionDurationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSessionDurationResponseBody) SetMessage(v string) *DescribeSessionDurationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSessionDurationResponseBody) SetRequestId(v string) *DescribeSessionDurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSessionDurationResponseBody) SetSuccess(v bool) *DescribeSessionDurationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSessionDurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSessionDurationResponseBodyData struct {
	SessionAverageDuration *int64                                                        `json:"SessionAverageDuration,omitempty" xml:"SessionAverageDuration,omitempty"`
	SessionDurationList    []*DescribeSessionDurationResponseBodyDataSessionDurationList `json:"SessionDurationList,omitempty" xml:"SessionDurationList,omitempty" type:"Repeated"`
	SessionTotalDuration   *int64                                                        `json:"SessionTotalDuration,omitempty" xml:"SessionTotalDuration,omitempty"`
}

func (s DescribeSessionDurationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionDurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSessionDurationResponseBodyData) GetSessionAverageDuration() *int64 {
	return s.SessionAverageDuration
}

func (s *DescribeSessionDurationResponseBodyData) GetSessionDurationList() []*DescribeSessionDurationResponseBodyDataSessionDurationList {
	return s.SessionDurationList
}

func (s *DescribeSessionDurationResponseBodyData) GetSessionTotalDuration() *int64 {
	return s.SessionTotalDuration
}

func (s *DescribeSessionDurationResponseBodyData) SetSessionAverageDuration(v int64) *DescribeSessionDurationResponseBodyData {
	s.SessionAverageDuration = &v
	return s
}

func (s *DescribeSessionDurationResponseBodyData) SetSessionDurationList(v []*DescribeSessionDurationResponseBodyDataSessionDurationList) *DescribeSessionDurationResponseBodyData {
	s.SessionDurationList = v
	return s
}

func (s *DescribeSessionDurationResponseBodyData) SetSessionTotalDuration(v int64) *DescribeSessionDurationResponseBodyData {
	s.SessionTotalDuration = &v
	return s
}

func (s *DescribeSessionDurationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSessionDurationResponseBodyDataSessionDurationList struct {
	SessionDuration *int64  `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	SessionTime     *string `json:"SessionTime,omitempty" xml:"SessionTime,omitempty"`
}

func (s DescribeSessionDurationResponseBodyDataSessionDurationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionDurationResponseBodyDataSessionDurationList) GoString() string {
	return s.String()
}

func (s *DescribeSessionDurationResponseBodyDataSessionDurationList) GetSessionDuration() *int64 {
	return s.SessionDuration
}

func (s *DescribeSessionDurationResponseBodyDataSessionDurationList) GetSessionTime() *string {
	return s.SessionTime
}

func (s *DescribeSessionDurationResponseBodyDataSessionDurationList) SetSessionDuration(v int64) *DescribeSessionDurationResponseBodyDataSessionDurationList {
	s.SessionDuration = &v
	return s
}

func (s *DescribeSessionDurationResponseBodyDataSessionDurationList) SetSessionTime(v string) *DescribeSessionDurationResponseBodyDataSessionDurationList {
	s.SessionTime = &v
	return s
}

func (s *DescribeSessionDurationResponseBodyDataSessionDurationList) Validate() error {
	return dara.Validate(s)
}
