// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *CreateMcpImageRequest
	GetApiKeyId() *string
	SetCustomerAliUid(v int64) *CreateMcpImageRequest
	GetCustomerAliUid() *int64
	SetGroupId(v string) *CreateMcpImageRequest
	GetGroupId() *string
	SetImageName(v string) *CreateMcpImageRequest
	GetImageName() *string
	SetParentImageId(v string) *CreateMcpImageRequest
	GetParentImageId() *string
	SetPersistentId(v string) *CreateMcpImageRequest
	GetPersistentId() *string
}

type CreateMcpImageRequest struct {
	ApiKeyId       *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	CustomerAliUid *int64  `json:"CustomerAliUid,omitempty" xml:"CustomerAliUid,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ImageName      *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ParentImageId  *string `json:"ParentImageId,omitempty" xml:"ParentImageId,omitempty"`
	PersistentId   *string `json:"PersistentId,omitempty" xml:"PersistentId,omitempty"`
}

func (s CreateMcpImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpImageRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpImageRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *CreateMcpImageRequest) GetCustomerAliUid() *int64 {
	return s.CustomerAliUid
}

func (s *CreateMcpImageRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateMcpImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateMcpImageRequest) GetParentImageId() *string {
	return s.ParentImageId
}

func (s *CreateMcpImageRequest) GetPersistentId() *string {
	return s.PersistentId
}

func (s *CreateMcpImageRequest) SetApiKeyId(v string) *CreateMcpImageRequest {
	s.ApiKeyId = &v
	return s
}

func (s *CreateMcpImageRequest) SetCustomerAliUid(v int64) *CreateMcpImageRequest {
	s.CustomerAliUid = &v
	return s
}

func (s *CreateMcpImageRequest) SetGroupId(v string) *CreateMcpImageRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMcpImageRequest) SetImageName(v string) *CreateMcpImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateMcpImageRequest) SetParentImageId(v string) *CreateMcpImageRequest {
	s.ParentImageId = &v
	return s
}

func (s *CreateMcpImageRequest) SetPersistentId(v string) *CreateMcpImageRequest {
	s.PersistentId = &v
	return s
}

func (s *CreateMcpImageRequest) Validate() error {
	return dara.Validate(s)
}
