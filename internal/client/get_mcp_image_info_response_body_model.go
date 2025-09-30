// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpImageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpImageInfoResponseBody
	GetCode() *string
	SetData(v *GetMcpImageInfoResponseBodyData) *GetMcpImageInfoResponseBody
	GetData() *GetMcpImageInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpImageInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpImageInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpImageInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpImageInfoResponseBody
	GetSuccess() *bool
}

type GetMcpImageInfoResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMcpImageInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpImageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpImageInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpImageInfoResponseBody) GetData() *GetMcpImageInfoResponseBodyData {
	return s.Data
}

func (s *GetMcpImageInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpImageInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpImageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpImageInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpImageInfoResponseBody) SetCode(v string) *GetMcpImageInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpImageInfoResponseBody) SetData(v *GetMcpImageInfoResponseBodyData) *GetMcpImageInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpImageInfoResponseBody) SetHttpStatusCode(v int32) *GetMcpImageInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpImageInfoResponseBody) SetMessage(v string) *GetMcpImageInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpImageInfoResponseBody) SetRequestId(v string) *GetMcpImageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpImageInfoResponseBody) SetSuccess(v bool) *GetMcpImageInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpImageInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcpImageInfoResponseBodyData struct {
	ImageApplyScene *string                                        `json:"ImageApplyScene,omitempty" xml:"ImageApplyScene,omitempty"`
	ImageBuildInfo  *GetMcpImageInfoResponseBodyDataImageBuildInfo `json:"ImageBuildInfo,omitempty" xml:"ImageBuildInfo,omitempty" type:"Struct"`
	ImageBuildType  *string                                        `json:"ImageBuildType,omitempty" xml:"ImageBuildType,omitempty"`
	ImageId         *string                                        `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageInfo       *GetMcpImageInfoResponseBodyDataImageInfo      `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	ImageName       *string                                        `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s GetMcpImageInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpImageInfoResponseBodyData) GetImageApplyScene() *string {
	return s.ImageApplyScene
}

func (s *GetMcpImageInfoResponseBodyData) GetImageBuildInfo() *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	return s.ImageBuildInfo
}

func (s *GetMcpImageInfoResponseBodyData) GetImageBuildType() *string {
	return s.ImageBuildType
}

func (s *GetMcpImageInfoResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetMcpImageInfoResponseBodyData) GetImageInfo() *GetMcpImageInfoResponseBodyDataImageInfo {
	return s.ImageInfo
}

func (s *GetMcpImageInfoResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *GetMcpImageInfoResponseBodyData) SetImageApplyScene(v string) *GetMcpImageInfoResponseBodyData {
	s.ImageApplyScene = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyData) SetImageBuildInfo(v *GetMcpImageInfoResponseBodyDataImageBuildInfo) *GetMcpImageInfoResponseBodyData {
	s.ImageBuildInfo = v
	return s
}

func (s *GetMcpImageInfoResponseBodyData) SetImageBuildType(v string) *GetMcpImageInfoResponseBodyData {
	s.ImageBuildType = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyData) SetImageId(v string) *GetMcpImageInfoResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyData) SetImageInfo(v *GetMcpImageInfoResponseBodyDataImageInfo) *GetMcpImageInfoResponseBodyData {
	s.ImageInfo = v
	return s
}

func (s *GetMcpImageInfoResponseBodyData) SetImageName(v string) *GetMcpImageInfoResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMcpImageInfoResponseBodyDataImageBuildInfo struct {
	AndroidMobileGroupId    *string `json:"AndroidMobileGroupId,omitempty" xml:"AndroidMobileGroupId,omitempty"`
	AndroidMobileInstanceId *string `json:"AndroidMobileInstanceId,omitempty" xml:"AndroidMobileInstanceId,omitempty"`
	ApiKeyId                *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	InstanceReady           *bool   `json:"InstanceReady,omitempty" xml:"InstanceReady,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VersionId               *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetMcpImageInfoResponseBodyDataImageBuildInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageInfoResponseBodyDataImageBuildInfo) GoString() string {
	return s.String()
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) GetAndroidMobileGroupId() *string {
	return s.AndroidMobileGroupId
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) GetAndroidMobileInstanceId() *string {
	return s.AndroidMobileInstanceId
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) GetInstanceReady() *bool {
	return s.InstanceReady
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) GetVersionId() *string {
	return s.VersionId
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) SetAndroidMobileGroupId(v string) *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	s.AndroidMobileGroupId = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) SetAndroidMobileInstanceId(v string) *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	s.AndroidMobileInstanceId = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) SetApiKeyId(v string) *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	s.ApiKeyId = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) SetInstanceReady(v bool) *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	s.InstanceReady = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) SetTaskId(v string) *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	s.TaskId = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) SetVersionId(v string) *GetMcpImageInfoResponseBodyDataImageBuildInfo {
	s.VersionId = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageBuildInfo) Validate() error {
	return dara.Validate(s)
}

type GetMcpImageInfoResponseBodyDataImageInfo struct {
	DataDiskSize   *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	OsName         *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	OsVersion      *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	PlatformName   *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDiskSize *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetMcpImageInfoResponseBodyDataImageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageInfoResponseBodyDataImageInfo) GoString() string {
	return s.String()
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetOsName() *string {
	return s.OsName
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetPlatformName() *string {
	return s.PlatformName
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetStatus() *string {
	return s.Status
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetDataDiskSize(v int32) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.DataDiskSize = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetOsName(v string) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.OsName = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetOsVersion(v string) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.OsVersion = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetPlatformName(v string) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.PlatformName = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetStatus(v string) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.Status = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetSystemDiskSize(v int32) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.SystemDiskSize = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) SetUpdateTime(v string) *GetMcpImageInfoResponseBodyDataImageInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetMcpImageInfoResponseBodyDataImageInfo) Validate() error {
	return dara.Validate(s)
}
