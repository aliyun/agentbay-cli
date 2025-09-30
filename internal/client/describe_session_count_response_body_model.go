// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSessionCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSessionCountResponseBody
	GetCode() *string
	SetData(v *DescribeSessionCountResponseBodyData) *DescribeSessionCountResponseBody
	GetData() *DescribeSessionCountResponseBodyData
	SetHttpStatusCode(v int32) *DescribeSessionCountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeSessionCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSessionCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSessionCountResponseBody
	GetSuccess() *bool
}

type DescribeSessionCountResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeSessionCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSessionCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSessionCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSessionCountResponseBody) GetData() *DescribeSessionCountResponseBodyData {
	return s.Data
}

func (s *DescribeSessionCountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeSessionCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSessionCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSessionCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSessionCountResponseBody) SetCode(v string) *DescribeSessionCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSessionCountResponseBody) SetData(v *DescribeSessionCountResponseBodyData) *DescribeSessionCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSessionCountResponseBody) SetHttpStatusCode(v int32) *DescribeSessionCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSessionCountResponseBody) SetMessage(v string) *DescribeSessionCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSessionCountResponseBody) SetRequestId(v string) *DescribeSessionCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSessionCountResponseBody) SetSuccess(v bool) *DescribeSessionCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSessionCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSessionCountResponseBodyData struct {
	SessionCountList    []*DescribeSessionCountResponseBodyDataSessionCountList `json:"SessionCountList,omitempty" xml:"SessionCountList,omitempty" type:"Repeated"`
	SessionFinishCount  *int64                                                  `json:"SessionFinishCount,omitempty" xml:"SessionFinishCount,omitempty"`
	SessionRunningCount *int64                                                  `json:"SessionRunningCount,omitempty" xml:"SessionRunningCount,omitempty"`
	SessionTotalCount   *int64                                                  `json:"SessionTotalCount,omitempty" xml:"SessionTotalCount,omitempty"`
}

func (s DescribeSessionCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSessionCountResponseBodyData) GetSessionCountList() []*DescribeSessionCountResponseBodyDataSessionCountList {
	return s.SessionCountList
}

func (s *DescribeSessionCountResponseBodyData) GetSessionFinishCount() *int64 {
	return s.SessionFinishCount
}

func (s *DescribeSessionCountResponseBodyData) GetSessionRunningCount() *int64 {
	return s.SessionRunningCount
}

func (s *DescribeSessionCountResponseBodyData) GetSessionTotalCount() *int64 {
	return s.SessionTotalCount
}

func (s *DescribeSessionCountResponseBodyData) SetSessionCountList(v []*DescribeSessionCountResponseBodyDataSessionCountList) *DescribeSessionCountResponseBodyData {
	s.SessionCountList = v
	return s
}

func (s *DescribeSessionCountResponseBodyData) SetSessionFinishCount(v int64) *DescribeSessionCountResponseBodyData {
	s.SessionFinishCount = &v
	return s
}

func (s *DescribeSessionCountResponseBodyData) SetSessionRunningCount(v int64) *DescribeSessionCountResponseBodyData {
	s.SessionRunningCount = &v
	return s
}

func (s *DescribeSessionCountResponseBodyData) SetSessionTotalCount(v int64) *DescribeSessionCountResponseBodyData {
	s.SessionTotalCount = &v
	return s
}

func (s *DescribeSessionCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSessionCountResponseBodyDataSessionCountList struct {
	SessionCount *int64  `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	SessionTime  *string `json:"SessionTime,omitempty" xml:"SessionTime,omitempty"`
}

func (s DescribeSessionCountResponseBodyDataSessionCountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSessionCountResponseBodyDataSessionCountList) GoString() string {
	return s.String()
}

func (s *DescribeSessionCountResponseBodyDataSessionCountList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeSessionCountResponseBodyDataSessionCountList) GetSessionTime() *string {
	return s.SessionTime
}

func (s *DescribeSessionCountResponseBodyDataSessionCountList) SetSessionCount(v int64) *DescribeSessionCountResponseBodyDataSessionCountList {
	s.SessionCount = &v
	return s
}

func (s *DescribeSessionCountResponseBodyDataSessionCountList) SetSessionTime(v string) *DescribeSessionCountResponseBodyDataSessionCountList {
	s.SessionTime = &v
	return s
}

func (s *DescribeSessionCountResponseBodyDataSessionCountList) Validate() error {
	return dara.Validate(s)
}
