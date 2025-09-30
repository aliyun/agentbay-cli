// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpResourcesResponseBody
	GetCode() *string
	SetData(v []*ListMcpResourcesResponseBodyData) *ListMcpResourcesResponseBody
	GetData() []*ListMcpResourcesResponseBodyData
	SetHttpStatusCode(v int32) *ListMcpResourcesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListMcpResourcesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpResourcesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpResourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMcpResourcesResponseBody
	GetTotalCount() *int32
}

type ListMcpResourcesResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListMcpResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcpResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpResourcesResponseBody) GetData() []*ListMcpResourcesResponseBodyData {
	return s.Data
}

func (s *ListMcpResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMcpResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMcpResourcesResponseBody) SetCode(v string) *ListMcpResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetData(v []*ListMcpResourcesResponseBodyData) *ListMcpResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListMcpResourcesResponseBody) SetHttpStatusCode(v int32) *ListMcpResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetMaxResults(v int32) *ListMcpResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetMessage(v string) *ListMcpResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetNextToken(v string) *ListMcpResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetRequestId(v string) *ListMcpResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetSuccess(v bool) *ListMcpResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpResourcesResponseBody) SetTotalCount(v int32) *ListMcpResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcpResourcesResponseBodyData struct {
	ApiKeyId     *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	ApiKeyName   *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName    *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListMcpResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMcpResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMcpResourcesResponseBodyData) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ListMcpResourcesResponseBodyData) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ListMcpResourcesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMcpResourcesResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ListMcpResourcesResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *ListMcpResourcesResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListMcpResourcesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListMcpResourcesResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *ListMcpResourcesResponseBodyData) SetApiKeyId(v string) *ListMcpResourcesResponseBodyData {
	s.ApiKeyId = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetApiKeyName(v string) *ListMcpResourcesResponseBodyData {
	s.ApiKeyName = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetCreateTime(v string) *ListMcpResourcesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetImageId(v string) *ListMcpResourcesResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetImageName(v string) *ListMcpResourcesResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetResourceType(v string) *ListMcpResourcesResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetStatus(v string) *ListMcpResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) SetUserName(v string) *ListMcpResourcesResponseBodyData {
	s.UserName = &v
	return s
}

func (s *ListMcpResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
