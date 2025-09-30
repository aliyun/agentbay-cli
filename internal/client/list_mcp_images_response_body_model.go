// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpImagesResponseBody
	GetCode() *string
	SetData(v []*ListMcpImagesResponseBodyData) *ListMcpImagesResponseBody
	GetData() []*ListMcpImagesResponseBodyData
	SetHttpStatusCode(v int32) *ListMcpImagesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListMcpImagesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpImagesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpImagesResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListMcpImagesResponseBody
	GetPageSize() *int32
	SetPageStart(v int32) *ListMcpImagesResponseBody
	GetPageStart() *int32
	SetRequestId(v string) *ListMcpImagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpImagesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMcpImagesResponseBody
	GetTotalCount() *int32
}

type ListMcpImagesResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListMcpImagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize       *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart      *int32                           `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcpImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpImagesResponseBody) GetData() []*ListMcpImagesResponseBodyData {
	return s.Data
}

func (s *ListMcpImagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMcpImagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpImagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpImagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpImagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcpImagesResponseBody) GetPageStart() *int32 {
	return s.PageStart
}

func (s *ListMcpImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpImagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpImagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMcpImagesResponseBody) SetCode(v string) *ListMcpImagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetData(v []*ListMcpImagesResponseBodyData) *ListMcpImagesResponseBody {
	s.Data = v
	return s
}

func (s *ListMcpImagesResponseBody) SetHttpStatusCode(v int32) *ListMcpImagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetMaxResults(v int32) *ListMcpImagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetMessage(v string) *ListMcpImagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetNextToken(v string) *ListMcpImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetPageSize(v int32) *ListMcpImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetPageStart(v int32) *ListMcpImagesResponseBody {
	s.PageStart = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetRequestId(v string) *ListMcpImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetSuccess(v bool) *ListMcpImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpImagesResponseBody) SetTotalCount(v int32) *ListMcpImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpImagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcpImagesResponseBodyData struct {
	ImageApplyScene        *string                                              `json:"ImageApplyScene,omitempty" xml:"ImageApplyScene,omitempty"`
	ImageBuildInfo         *ListMcpImagesResponseBodyDataImageBuildInfo         `json:"ImageBuildInfo,omitempty" xml:"ImageBuildInfo,omitempty" type:"Struct"`
	ImageBuildType         *string                                              `json:"ImageBuildType,omitempty" xml:"ImageBuildType,omitempty"`
	ImageId                *string                                              `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageInfo              *ListMcpImagesResponseBodyDataImageInfo              `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	ImageIntro             *string                                              `json:"ImageIntro,omitempty" xml:"ImageIntro,omitempty"`
	ImageName              *string                                              `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageResourceGroupInfo *ListMcpImagesResponseBodyDataImageResourceGroupInfo `json:"ImageResourceGroupInfo,omitempty" xml:"ImageResourceGroupInfo,omitempty" type:"Struct"`
	ImageResourceStatus    *string                                              `json:"ImageResourceStatus,omitempty" xml:"ImageResourceStatus,omitempty"`
	ToolInfo               []*ListMcpImagesResponseBodyDataToolInfo             `json:"ToolInfo,omitempty" xml:"ToolInfo,omitempty" type:"Repeated"`
}

func (s ListMcpImagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBodyData) GetImageApplyScene() *string {
	return s.ImageApplyScene
}

func (s *ListMcpImagesResponseBodyData) GetImageBuildInfo() *ListMcpImagesResponseBodyDataImageBuildInfo {
	return s.ImageBuildInfo
}

func (s *ListMcpImagesResponseBodyData) GetImageBuildType() *string {
	return s.ImageBuildType
}

func (s *ListMcpImagesResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ListMcpImagesResponseBodyData) GetImageInfo() *ListMcpImagesResponseBodyDataImageInfo {
	return s.ImageInfo
}

func (s *ListMcpImagesResponseBodyData) GetImageIntro() *string {
	return s.ImageIntro
}

func (s *ListMcpImagesResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *ListMcpImagesResponseBodyData) GetImageResourceGroupInfo() *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	return s.ImageResourceGroupInfo
}

func (s *ListMcpImagesResponseBodyData) GetImageResourceStatus() *string {
	return s.ImageResourceStatus
}

func (s *ListMcpImagesResponseBodyData) GetToolInfo() []*ListMcpImagesResponseBodyDataToolInfo {
	return s.ToolInfo
}

func (s *ListMcpImagesResponseBodyData) SetImageApplyScene(v string) *ListMcpImagesResponseBodyData {
	s.ImageApplyScene = &v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageBuildInfo(v *ListMcpImagesResponseBodyDataImageBuildInfo) *ListMcpImagesResponseBodyData {
	s.ImageBuildInfo = v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageBuildType(v string) *ListMcpImagesResponseBodyData {
	s.ImageBuildType = &v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageId(v string) *ListMcpImagesResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageInfo(v *ListMcpImagesResponseBodyDataImageInfo) *ListMcpImagesResponseBodyData {
	s.ImageInfo = v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageIntro(v string) *ListMcpImagesResponseBodyData {
	s.ImageIntro = &v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageName(v string) *ListMcpImagesResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageResourceGroupInfo(v *ListMcpImagesResponseBodyDataImageResourceGroupInfo) *ListMcpImagesResponseBodyData {
	s.ImageResourceGroupInfo = v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetImageResourceStatus(v string) *ListMcpImagesResponseBodyData {
	s.ImageResourceStatus = &v
	return s
}

func (s *ListMcpImagesResponseBodyData) SetToolInfo(v []*ListMcpImagesResponseBodyDataToolInfo) *ListMcpImagesResponseBodyData {
	s.ToolInfo = v
	return s
}

func (s *ListMcpImagesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListMcpImagesResponseBodyDataImageBuildInfo struct {
	ApiKeyId      *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	InstanceReady *bool   `json:"InstanceReady,omitempty" xml:"InstanceReady,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListMcpImagesResponseBodyDataImageBuildInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBodyDataImageBuildInfo) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) GetInstanceReady() *bool {
	return s.InstanceReady
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) SetApiKeyId(v string) *ListMcpImagesResponseBodyDataImageBuildInfo {
	s.ApiKeyId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) SetInstanceReady(v bool) *ListMcpImagesResponseBodyDataImageBuildInfo {
	s.InstanceReady = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) SetTaskId(v string) *ListMcpImagesResponseBodyDataImageBuildInfo {
	s.TaskId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) SetVersionId(v string) *ListMcpImagesResponseBodyDataImageBuildInfo {
	s.VersionId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageBuildInfo) Validate() error {
	return dara.Validate(s)
}

type ListMcpImagesResponseBodyDataImageInfo struct {
	DataDiskSize   *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	FotaVersion    *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	OsName         *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	OsVersion      *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	PlatformName   *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDiskSize *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListMcpImagesResponseBodyDataImageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBodyDataImageInfo) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetFotaVersion() *string {
	return s.FotaVersion
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetOsName() *string {
	return s.OsName
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetOsVersion() *string {
	return s.OsVersion
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetStatus() *string {
	return s.Status
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ListMcpImagesResponseBodyDataImageInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetDataDiskSize(v int32) *ListMcpImagesResponseBodyDataImageInfo {
	s.DataDiskSize = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetFotaVersion(v string) *ListMcpImagesResponseBodyDataImageInfo {
	s.FotaVersion = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetOsName(v string) *ListMcpImagesResponseBodyDataImageInfo {
	s.OsName = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetOsVersion(v string) *ListMcpImagesResponseBodyDataImageInfo {
	s.OsVersion = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetPlatformName(v string) *ListMcpImagesResponseBodyDataImageInfo {
	s.PlatformName = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetStatus(v string) *ListMcpImagesResponseBodyDataImageInfo {
	s.Status = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetSystemDiskSize(v int32) *ListMcpImagesResponseBodyDataImageInfo {
	s.SystemDiskSize = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) SetUpdateTime(v string) *ListMcpImagesResponseBodyDataImageInfo {
	s.UpdateTime = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageInfo) Validate() error {
	return dara.Validate(s)
}

type ListMcpImagesResponseBodyDataImageResourceGroupInfo struct {
	BizRegionId         *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	OfficeSiteId        *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteType      *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	PolicyId            *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupStatus *string `json:"ResourceGroupStatus,omitempty" xml:"ResourceGroupStatus,omitempty"`
	SessionBandwidth    *int32  `json:"SessionBandwidth,omitempty" xml:"SessionBandwidth,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListMcpImagesResponseBodyDataImageResourceGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBodyDataImageResourceGroupInfo) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetResourceGroupStatus() *string {
	return s.ResourceGroupStatus
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetSessionBandwidth() *int32 {
	return s.SessionBandwidth
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetBizRegionId(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.BizRegionId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetOfficeSiteId(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.OfficeSiteId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetOfficeSiteType(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.OfficeSiteType = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetPolicyId(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.PolicyId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetResourceGroupId(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetResourceGroupStatus(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.ResourceGroupStatus = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetSessionBandwidth(v int32) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.SessionBandwidth = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetVSwitchId(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.VSwitchId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) SetVpcId(v string) *ListMcpImagesResponseBodyDataImageResourceGroupInfo {
	s.VpcId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataImageResourceGroupInfo) Validate() error {
	return dara.Validate(s)
}

type ListMcpImagesResponseBodyDataToolInfo struct {
	McpServerId   *string                                          `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	McpServerName *string                                          `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
	ToolList      []*ListMcpImagesResponseBodyDataToolInfoToolList `json:"ToolList,omitempty" xml:"ToolList,omitempty" type:"Repeated"`
}

func (s ListMcpImagesResponseBodyDataToolInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBodyDataToolInfo) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBodyDataToolInfo) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListMcpImagesResponseBodyDataToolInfo) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *ListMcpImagesResponseBodyDataToolInfo) GetToolList() []*ListMcpImagesResponseBodyDataToolInfoToolList {
	return s.ToolList
}

func (s *ListMcpImagesResponseBodyDataToolInfo) SetMcpServerId(v string) *ListMcpImagesResponseBodyDataToolInfo {
	s.McpServerId = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataToolInfo) SetMcpServerName(v string) *ListMcpImagesResponseBodyDataToolInfo {
	s.McpServerName = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataToolInfo) SetToolList(v []*ListMcpImagesResponseBodyDataToolInfoToolList) *ListMcpImagesResponseBodyDataToolInfo {
	s.ToolList = v
	return s
}

func (s *ListMcpImagesResponseBodyDataToolInfo) Validate() error {
	return dara.Validate(s)
}

type ListMcpImagesResponseBodyDataToolInfoToolList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Tool        *string `json:"Tool,omitempty" xml:"Tool,omitempty"`
}

func (s ListMcpImagesResponseBodyDataToolInfoToolList) String() string {
	return dara.Prettify(s)
}

func (s ListMcpImagesResponseBodyDataToolInfoToolList) GoString() string {
	return s.String()
}

func (s *ListMcpImagesResponseBodyDataToolInfoToolList) GetDescription() *string {
	return s.Description
}

func (s *ListMcpImagesResponseBodyDataToolInfoToolList) GetTool() *string {
	return s.Tool
}

func (s *ListMcpImagesResponseBodyDataToolInfoToolList) SetDescription(v string) *ListMcpImagesResponseBodyDataToolInfoToolList {
	s.Description = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataToolInfoToolList) SetTool(v string) *ListMcpImagesResponseBodyDataToolInfoToolList {
	s.Tool = &v
	return s
}

func (s *ListMcpImagesResponseBodyDataToolInfoToolList) Validate() error {
	return dara.Validate(s)
}
