// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListResourceGroupsResponseBody
	GetCode() *string
	SetData(v []*ListResourceGroupsResponseBodyData) *ListResourceGroupsResponseBody
	GetData() []*ListResourceGroupsResponseBodyData
	SetHttpStatusCode(v int32) *ListResourceGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListResourceGroupsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListResourceGroupsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListResourceGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourceGroupsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListResourceGroupsResponseBody
	GetTotalCount() *int32
}

type ListResourceGroupsResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListResourceGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListResourceGroupsResponseBody) GetData() []*ListResourceGroupsResponseBodyData {
	return s.Data
}

func (s *ListResourceGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListResourceGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListResourceGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceGroupsResponseBody) SetCode(v string) *ListResourceGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetData(v []*ListResourceGroupsResponseBodyData) *ListResourceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetHttpStatusCode(v int32) *ListResourceGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetMaxResults(v int32) *ListResourceGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetMessage(v string) *ListResourceGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetNextToken(v string) *ListResourceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetSuccess(v bool) *ListResourceGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int32) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsResponseBodyData struct {
	BizRegionId      *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	OfficeSiteId     *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteType   *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	ResourceConfigId *string `json:"ResourceConfigId,omitempty" xml:"ResourceConfigId,omitempty"`
}

func (s ListResourceGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyData) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListResourceGroupsResponseBodyData) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *ListResourceGroupsResponseBodyData) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListResourceGroupsResponseBodyData) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *ListResourceGroupsResponseBodyData) GetResourceConfigId() *string {
	return s.ResourceConfigId
}

func (s *ListResourceGroupsResponseBodyData) SetBizRegionId(v string) *ListResourceGroupsResponseBodyData {
	s.BizRegionId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetNetworkPackageId(v string) *ListResourceGroupsResponseBodyData {
	s.NetworkPackageId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetOfficeSiteId(v string) *ListResourceGroupsResponseBodyData {
	s.OfficeSiteId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetOfficeSiteType(v string) *ListResourceGroupsResponseBodyData {
	s.OfficeSiteType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetResourceConfigId(v string) *ListResourceGroupsResponseBodyData {
	s.ResourceConfigId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
