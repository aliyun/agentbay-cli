// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpResourceResponseBody
	GetCode() *string
	SetData(v *GetMcpResourceResponseBodyData) *GetMcpResourceResponseBody
	GetData() *GetMcpResourceResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpResourceResponseBody
	GetSuccess() *bool
}

type GetMcpResourceResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMcpResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpResourceResponseBody) GetData() *GetMcpResourceResponseBodyData {
	return s.Data
}

func (s *GetMcpResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpResourceResponseBody) SetCode(v string) *GetMcpResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetData(v *GetMcpResourceResponseBodyData) *GetMcpResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpResourceResponseBody) SetHttpStatusCode(v int32) *GetMcpResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetMessage(v string) *GetMcpResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetRequestId(v string) *GetMcpResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpResourceResponseBody) SetSuccess(v bool) *GetMcpResourceResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcpResourceResponseBodyData struct {
	ApiKeyName            *string                                              `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	Concurrency           *int32                                               `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	DesktopReleaseTimings *GetMcpResourceResponseBodyDataDesktopReleaseTimings `json:"DesktopReleaseTimings,omitempty" xml:"DesktopReleaseTimings,omitempty" type:"Struct"`
	ImageId               *string                                              `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageInfo             *GetMcpResourceResponseBodyDataImageInfo             `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	ImageName             *string                                              `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageProperties       *string                                              `json:"ImageProperties,omitempty" xml:"ImageProperties,omitempty"`
	ResourceType          *string                                              `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	UserName              *string                                              `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetMcpResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponseBodyData) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *GetMcpResourceResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetMcpResourceResponseBodyData) GetDesktopReleaseTimings() *GetMcpResourceResponseBodyDataDesktopReleaseTimings {
	return s.DesktopReleaseTimings
}

func (s *GetMcpResourceResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetMcpResourceResponseBodyData) GetImageInfo() *GetMcpResourceResponseBodyDataImageInfo {
	return s.ImageInfo
}

func (s *GetMcpResourceResponseBodyData) GetImageName() *string {
	return s.ImageName
}

func (s *GetMcpResourceResponseBodyData) GetImageProperties() *string {
	return s.ImageProperties
}

func (s *GetMcpResourceResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetMcpResourceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMcpResourceResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetMcpResourceResponseBodyData) SetApiKeyName(v string) *GetMcpResourceResponseBodyData {
	s.ApiKeyName = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetConcurrency(v int32) *GetMcpResourceResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetDesktopReleaseTimings(v *GetMcpResourceResponseBodyDataDesktopReleaseTimings) *GetMcpResourceResponseBodyData {
	s.DesktopReleaseTimings = v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetImageId(v string) *GetMcpResourceResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetImageInfo(v *GetMcpResourceResponseBodyDataImageInfo) *GetMcpResourceResponseBodyData {
	s.ImageInfo = v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetImageName(v string) *GetMcpResourceResponseBodyData {
	s.ImageName = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetImageProperties(v string) *GetMcpResourceResponseBodyData {
	s.ImageProperties = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetResourceType(v string) *GetMcpResourceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetStatus(v string) *GetMcpResourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) SetUserName(v string) *GetMcpResourceResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetMcpResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMcpResourceResponseBodyDataDesktopReleaseTimings struct {
	DesktopMaxRuntime *int32 `json:"DesktopMaxRuntime,omitempty" xml:"DesktopMaxRuntime,omitempty"`
	McpIdleTimeout    *int32 `json:"McpIdleTimeout,omitempty" xml:"McpIdleTimeout,omitempty"`
	UserIdleTimeout   *int32 `json:"UserIdleTimeout,omitempty" xml:"UserIdleTimeout,omitempty"`
}

func (s GetMcpResourceResponseBodyDataDesktopReleaseTimings) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResourceResponseBodyDataDesktopReleaseTimings) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) GetDesktopMaxRuntime() *int32 {
	return s.DesktopMaxRuntime
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) GetMcpIdleTimeout() *int32 {
	return s.McpIdleTimeout
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) GetUserIdleTimeout() *int32 {
	return s.UserIdleTimeout
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) SetDesktopMaxRuntime(v int32) *GetMcpResourceResponseBodyDataDesktopReleaseTimings {
	s.DesktopMaxRuntime = &v
	return s
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) SetMcpIdleTimeout(v int32) *GetMcpResourceResponseBodyDataDesktopReleaseTimings {
	s.McpIdleTimeout = &v
	return s
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) SetUserIdleTimeout(v int32) *GetMcpResourceResponseBodyDataDesktopReleaseTimings {
	s.UserIdleTimeout = &v
	return s
}

func (s *GetMcpResourceResponseBodyDataDesktopReleaseTimings) Validate() error {
	return dara.Validate(s)
}

type GetMcpResourceResponseBodyDataImageInfo struct {
	OsName     *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	OsVersion  *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetMcpResourceResponseBodyDataImageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResourceResponseBodyDataImageInfo) GoString() string {
	return s.String()
}

func (s *GetMcpResourceResponseBodyDataImageInfo) GetOsName() *string {
	return s.OsName
}

func (s *GetMcpResourceResponseBodyDataImageInfo) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetMcpResourceResponseBodyDataImageInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMcpResourceResponseBodyDataImageInfo) SetOsName(v string) *GetMcpResourceResponseBodyDataImageInfo {
	s.OsName = &v
	return s
}

func (s *GetMcpResourceResponseBodyDataImageInfo) SetOsVersion(v string) *GetMcpResourceResponseBodyDataImageInfo {
	s.OsVersion = &v
	return s
}

func (s *GetMcpResourceResponseBodyDataImageInfo) SetUpdateTime(v string) *GetMcpResourceResponseBodyDataImageInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetMcpResourceResponseBodyDataImageInfo) Validate() error {
	return dara.Validate(s)
}
