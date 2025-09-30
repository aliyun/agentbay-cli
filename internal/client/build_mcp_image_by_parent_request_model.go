// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildMcpImageByParentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiskSize(v int32) *BuildMcpImageByParentRequest
	GetDataDiskSize() *int32
	SetDescription(v string) *BuildMcpImageByParentRequest
	GetDescription() *string
	SetImageName(v string) *BuildMcpImageByParentRequest
	GetImageName() *string
	SetLanguage(v string) *BuildMcpImageByParentRequest
	GetLanguage() *string
	SetOsType(v string) *BuildMcpImageByParentRequest
	GetOsType() *string
	SetParentImageId(v string) *BuildMcpImageByParentRequest
	GetParentImageId() *string
	SetPerformanceType(v string) *BuildMcpImageByParentRequest
	GetPerformanceType() *string
	SetSystemDiskSize(v int32) *BuildMcpImageByParentRequest
	GetSystemDiskSize() *int32
}

type BuildMcpImageByParentRequest struct {
	DataDiskSize    *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Language        *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OsType          *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ParentImageId   *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	PerformanceType *string `json:"PerformanceType,omitempty" xml:"PerformanceType,omitempty"`
	SystemDiskSize  *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s BuildMcpImageByParentRequest) String() string {
	return dara.Prettify(s)
}

func (s BuildMcpImageByParentRequest) GoString() string {
	return s.String()
}

func (s *BuildMcpImageByParentRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *BuildMcpImageByParentRequest) GetDescription() *string {
	return s.Description
}

func (s *BuildMcpImageByParentRequest) GetImageName() *string {
	return s.ImageName
}

func (s *BuildMcpImageByParentRequest) GetLanguage() *string {
	return s.Language
}

func (s *BuildMcpImageByParentRequest) GetOsType() *string {
	return s.OsType
}

func (s *BuildMcpImageByParentRequest) GetParentImageId() *string {
	return s.ParentImageId
}

func (s *BuildMcpImageByParentRequest) GetPerformanceType() *string {
	return s.PerformanceType
}

func (s *BuildMcpImageByParentRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *BuildMcpImageByParentRequest) SetDataDiskSize(v int32) *BuildMcpImageByParentRequest {
	s.DataDiskSize = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetDescription(v string) *BuildMcpImageByParentRequest {
	s.Description = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetImageName(v string) *BuildMcpImageByParentRequest {
	s.ImageName = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetLanguage(v string) *BuildMcpImageByParentRequest {
	s.Language = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetOsType(v string) *BuildMcpImageByParentRequest {
	s.OsType = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetParentImageId(v string) *BuildMcpImageByParentRequest {
	s.ParentImageId = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetPerformanceType(v string) *BuildMcpImageByParentRequest {
	s.PerformanceType = &v
	return s
}

func (s *BuildMcpImageByParentRequest) SetSystemDiskSize(v int32) *BuildMcpImageByParentRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *BuildMcpImageByParentRequest) Validate() error {
	return dara.Validate(s)
}
