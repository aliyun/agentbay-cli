// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserBenefitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeScene(v string) *ModifyUserBenefitRequest
	GetChangeScene() *string
	SetNewBenefitLevel(v string) *ModifyUserBenefitRequest
	GetNewBenefitLevel() *string
	SetOrderBenefitParam(v string) *ModifyUserBenefitRequest
	GetOrderBenefitParam() *string
	SetResourceGroupId(v string) *ModifyUserBenefitRequest
	GetResourceGroupId() *string
	SetSessionBandwidth(v int32) *ModifyUserBenefitRequest
	GetSessionBandwidth() *int32
}

type ModifyUserBenefitRequest struct {
	ChangeScene       *string `json:"ChangeScene,omitempty" xml:"ChangeScene,omitempty"`
	NewBenefitLevel   *string `json:"NewBenefitLevel,omitempty" xml:"NewBenefitLevel,omitempty"`
	OrderBenefitParam *string `json:"OrderBenefitParam,omitempty" xml:"OrderBenefitParam,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SessionBandwidth  *int32  `json:"SessionBandwidth,omitempty" xml:"SessionBandwidth,omitempty"`
}

func (s ModifyUserBenefitRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserBenefitRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserBenefitRequest) GetChangeScene() *string {
	return s.ChangeScene
}

func (s *ModifyUserBenefitRequest) GetNewBenefitLevel() *string {
	return s.NewBenefitLevel
}

func (s *ModifyUserBenefitRequest) GetOrderBenefitParam() *string {
	return s.OrderBenefitParam
}

func (s *ModifyUserBenefitRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyUserBenefitRequest) GetSessionBandwidth() *int32 {
	return s.SessionBandwidth
}

func (s *ModifyUserBenefitRequest) SetChangeScene(v string) *ModifyUserBenefitRequest {
	s.ChangeScene = &v
	return s
}

func (s *ModifyUserBenefitRequest) SetNewBenefitLevel(v string) *ModifyUserBenefitRequest {
	s.NewBenefitLevel = &v
	return s
}

func (s *ModifyUserBenefitRequest) SetOrderBenefitParam(v string) *ModifyUserBenefitRequest {
	s.OrderBenefitParam = &v
	return s
}

func (s *ModifyUserBenefitRequest) SetResourceGroupId(v string) *ModifyUserBenefitRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyUserBenefitRequest) SetSessionBandwidth(v int32) *ModifyUserBenefitRequest {
	s.SessionBandwidth = &v
	return s
}

func (s *ModifyUserBenefitRequest) Validate() error {
	return dara.Validate(s)
}
