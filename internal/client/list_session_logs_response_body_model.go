// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSessionLogsResponseBody
	GetCode() *string
	SetData(v []*ListSessionLogsResponseBodyData) *ListSessionLogsResponseBody
	GetData() []*ListSessionLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListSessionLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListSessionLogsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListSessionLogsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSessionLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSessionLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSessionLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSessionLogsResponseBody
	GetTotalCount() *int32
}

type ListSessionLogsResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListSessionLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSessionLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSessionLogsResponseBody) GetData() []*ListSessionLogsResponseBodyData {
	return s.Data
}

func (s *ListSessionLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSessionLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSessionLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSessionLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSessionLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSessionLogsResponseBody) SetCode(v string) *ListSessionLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetData(v []*ListSessionLogsResponseBodyData) *ListSessionLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionLogsResponseBody) SetHttpStatusCode(v int32) *ListSessionLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetMaxResults(v int32) *ListSessionLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetMessage(v string) *ListSessionLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetNextToken(v string) *ListSessionLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetRequestId(v string) *ListSessionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetSuccess(v bool) *ListSessionLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSessionLogsResponseBody) SetTotalCount(v int32) *ListSessionLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSessionLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSessionLogsResponseBodyData struct {
	ApiKeyName      *string   `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	ApikeyCount     *int64    `json:"ApikeyCount,omitempty" xml:"ApikeyCount,omitempty"`
	AppInstanceId   *string   `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ContextIds      []*string `json:"ContextIds,omitempty" xml:"ContextIds,omitempty" type:"Repeated"`
	EndTime         *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ImageId         *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType       *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	SessionDuration *int64    `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	SessionId       *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionStatus   *string   `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	StartTime       *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSessionLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSessionLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionLogsResponseBodyData) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ListSessionLogsResponseBodyData) GetApikeyCount() *int64 {
	return s.ApikeyCount
}

func (s *ListSessionLogsResponseBodyData) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *ListSessionLogsResponseBodyData) GetContextIds() []*string {
	return s.ContextIds
}

func (s *ListSessionLogsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSessionLogsResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ListSessionLogsResponseBodyData) GetImageType() *string {
	return s.ImageType
}

func (s *ListSessionLogsResponseBodyData) GetSessionDuration() *int64 {
	return s.SessionDuration
}

func (s *ListSessionLogsResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSessionLogsResponseBodyData) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *ListSessionLogsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSessionLogsResponseBodyData) SetApiKeyName(v string) *ListSessionLogsResponseBodyData {
	s.ApiKeyName = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetApikeyCount(v int64) *ListSessionLogsResponseBodyData {
	s.ApikeyCount = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetAppInstanceId(v string) *ListSessionLogsResponseBodyData {
	s.AppInstanceId = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetContextIds(v []*string) *ListSessionLogsResponseBodyData {
	s.ContextIds = v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetEndTime(v string) *ListSessionLogsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetImageId(v string) *ListSessionLogsResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetImageType(v string) *ListSessionLogsResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetSessionDuration(v int64) *ListSessionLogsResponseBodyData {
	s.SessionDuration = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetSessionId(v string) *ListSessionLogsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetSessionStatus(v string) *ListSessionLogsResponseBodyData {
	s.SessionStatus = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) SetStartTime(v string) *ListSessionLogsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListSessionLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
