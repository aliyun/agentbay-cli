// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionToolLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSessionToolLogsResponseBody
	GetCode() *string
	SetData(v []*ListSessionToolLogsResponseBodyData) *ListSessionToolLogsResponseBody
	GetData() []*ListSessionToolLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListSessionToolLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListSessionToolLogsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListSessionToolLogsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSessionToolLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSessionToolLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSessionToolLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSessionToolLogsResponseBody
	GetTotalCount() *int32
}

type ListSessionToolLogsResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListSessionToolLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSessionToolLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionToolLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionToolLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSessionToolLogsResponseBody) GetData() []*ListSessionToolLogsResponseBodyData {
	return s.Data
}

func (s *ListSessionToolLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSessionToolLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSessionToolLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSessionToolLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionToolLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionToolLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSessionToolLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSessionToolLogsResponseBody) SetCode(v string) *ListSessionToolLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetData(v []*ListSessionToolLogsResponseBodyData) *ListSessionToolLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetHttpStatusCode(v int32) *ListSessionToolLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetMaxResults(v int32) *ListSessionToolLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetMessage(v string) *ListSessionToolLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetNextToken(v string) *ListSessionToolLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetRequestId(v string) *ListSessionToolLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetSuccess(v bool) *ListSessionToolLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) SetTotalCount(v int32) *ListSessionToolLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSessionToolLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSessionToolLogsResponseBodyData struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestData  *string `json:"RequestData,omitempty" xml:"RequestData,omitempty"`
	ResponseData *string `json:"ResponseData,omitempty" xml:"ResponseData,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ToolName     *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolResult   *string `json:"ToolResult,omitempty" xml:"ToolResult,omitempty"`
}

func (s ListSessionToolLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSessionToolLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionToolLogsResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSessionToolLogsResponseBodyData) GetRequestData() *string {
	return s.RequestData
}

func (s *ListSessionToolLogsResponseBodyData) GetResponseData() *string {
	return s.ResponseData
}

func (s *ListSessionToolLogsResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSessionToolLogsResponseBodyData) GetToolName() *string {
	return s.ToolName
}

func (s *ListSessionToolLogsResponseBodyData) GetToolResult() *string {
	return s.ToolResult
}

func (s *ListSessionToolLogsResponseBodyData) SetEndTime(v string) *ListSessionToolLogsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListSessionToolLogsResponseBodyData) SetRequestData(v string) *ListSessionToolLogsResponseBodyData {
	s.RequestData = &v
	return s
}

func (s *ListSessionToolLogsResponseBodyData) SetResponseData(v string) *ListSessionToolLogsResponseBodyData {
	s.ResponseData = &v
	return s
}

func (s *ListSessionToolLogsResponseBodyData) SetStartTime(v string) *ListSessionToolLogsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListSessionToolLogsResponseBodyData) SetToolName(v string) *ListSessionToolLogsResponseBodyData {
	s.ToolName = &v
	return s
}

func (s *ListSessionToolLogsResponseBodyData) SetToolResult(v string) *ListSessionToolLogsResponseBodyData {
	s.ToolResult = &v
	return s
}

func (s *ListSessionToolLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
