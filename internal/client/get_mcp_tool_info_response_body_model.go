// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpToolInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpToolInfoResponseBody
	GetCode() *string
	SetData(v []*GetMcpToolInfoResponseBodyData) *GetMcpToolInfoResponseBody
	GetData() []*GetMcpToolInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpToolInfoResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *GetMcpToolInfoResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *GetMcpToolInfoResponseBody
	GetMessage() *string
	SetNextToken(v string) *GetMcpToolInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetMcpToolInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpToolInfoResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetMcpToolInfoResponseBody
	GetTotalCount() *int32
}

type GetMcpToolInfoResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*GetMcpToolInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMcpToolInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpToolInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpToolInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpToolInfoResponseBody) GetData() []*GetMcpToolInfoResponseBodyData {
	return s.Data
}

func (s *GetMcpToolInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpToolInfoResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetMcpToolInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpToolInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMcpToolInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpToolInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpToolInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetMcpToolInfoResponseBody) SetCode(v string) *GetMcpToolInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetData(v []*GetMcpToolInfoResponseBodyData) *GetMcpToolInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetHttpStatusCode(v int32) *GetMcpToolInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetMaxResults(v int32) *GetMcpToolInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetMessage(v string) *GetMcpToolInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetNextToken(v string) *GetMcpToolInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetRequestId(v string) *GetMcpToolInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetSuccess(v bool) *GetMcpToolInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) SetTotalCount(v int32) *GetMcpToolInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetMcpToolInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcpToolInfoResponseBodyData struct {
	// 代表资源一级ID的资源属性字段
	McpServerId   *string                                   `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	McpServerName *string                                   `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
	ToolList      []*GetMcpToolInfoResponseBodyDataToolList `json:"ToolList,omitempty" xml:"ToolList,omitempty" type:"Repeated"`
}

func (s GetMcpToolInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpToolInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpToolInfoResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetMcpToolInfoResponseBodyData) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *GetMcpToolInfoResponseBodyData) GetToolList() []*GetMcpToolInfoResponseBodyDataToolList {
	return s.ToolList
}

func (s *GetMcpToolInfoResponseBodyData) SetMcpServerId(v string) *GetMcpToolInfoResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *GetMcpToolInfoResponseBodyData) SetMcpServerName(v string) *GetMcpToolInfoResponseBodyData {
	s.McpServerName = &v
	return s
}

func (s *GetMcpToolInfoResponseBodyData) SetToolList(v []*GetMcpToolInfoResponseBodyDataToolList) *GetMcpToolInfoResponseBodyData {
	s.ToolList = v
	return s
}

func (s *GetMcpToolInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMcpToolInfoResponseBodyDataToolList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Tool        *string `json:"Tool,omitempty" xml:"Tool,omitempty"`
}

func (s GetMcpToolInfoResponseBodyDataToolList) String() string {
	return dara.Prettify(s)
}

func (s GetMcpToolInfoResponseBodyDataToolList) GoString() string {
	return s.String()
}

func (s *GetMcpToolInfoResponseBodyDataToolList) GetDescription() *string {
	return s.Description
}

func (s *GetMcpToolInfoResponseBodyDataToolList) GetTool() *string {
	return s.Tool
}

func (s *GetMcpToolInfoResponseBodyDataToolList) SetDescription(v string) *GetMcpToolInfoResponseBodyDataToolList {
	s.Description = &v
	return s
}

func (s *GetMcpToolInfoResponseBodyDataToolList) SetTool(v string) *GetMcpToolInfoResponseBodyDataToolList {
	s.Tool = &v
	return s
}

func (s *GetMcpToolInfoResponseBodyDataToolList) Validate() error {
	return dara.Validate(s)
}
