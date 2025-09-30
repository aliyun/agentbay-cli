// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBenefitDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBenefitLevel(v string) *DescribeUserBenefitDetailResponseBody
	GetBenefitLevel() *string
	SetCode(v string) *DescribeUserBenefitDetailResponseBody
	GetCode() *string
	SetExpireAt(v string) *DescribeUserBenefitDetailResponseBody
	GetExpireAt() *string
	SetHttpStatusCode(v int32) *DescribeUserBenefitDetailResponseBody
	GetHttpStatusCode() *int32
	SetImageSpec(v *DescribeUserBenefitDetailResponseBodyImageSpec) *DescribeUserBenefitDetailResponseBody
	GetImageSpec() *DescribeUserBenefitDetailResponseBodyImageSpec
	SetMessage(v string) *DescribeUserBenefitDetailResponseBody
	GetMessage() *string
	SetNetworkSpec(v *DescribeUserBenefitDetailResponseBodyNetworkSpec) *DescribeUserBenefitDetailResponseBody
	GetNetworkSpec() *DescribeUserBenefitDetailResponseBodyNetworkSpec
	SetRequestId(v string) *DescribeUserBenefitDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeUserBenefitDetailResponseBody
	GetSuccess() *bool
	SetUserBaseSpec(v *DescribeUserBenefitDetailResponseBodyUserBaseSpec) *DescribeUserBenefitDetailResponseBody
	GetUserBaseSpec() *DescribeUserBenefitDetailResponseBodyUserBaseSpec
}

type DescribeUserBenefitDetailResponseBody struct {
	BenefitLevel   *string                                            `json:"BenefitLevel,omitempty" xml:"BenefitLevel,omitempty"`
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ExpireAt       *string                                            `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	HttpStatusCode *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ImageSpec      *DescribeUserBenefitDetailResponseBodyImageSpec    `json:"ImageSpec,omitempty" xml:"ImageSpec,omitempty" type:"Struct"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	NetworkSpec    *DescribeUserBenefitDetailResponseBodyNetworkSpec  `json:"NetworkSpec,omitempty" xml:"NetworkSpec,omitempty" type:"Struct"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	UserBaseSpec   *DescribeUserBenefitDetailResponseBodyUserBaseSpec `json:"UserBaseSpec,omitempty" xml:"UserBaseSpec,omitempty" type:"Struct"`
}

func (s DescribeUserBenefitDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBenefitDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBenefitDetailResponseBody) GetBenefitLevel() *string {
	return s.BenefitLevel
}

func (s *DescribeUserBenefitDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeUserBenefitDetailResponseBody) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *DescribeUserBenefitDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeUserBenefitDetailResponseBody) GetImageSpec() *DescribeUserBenefitDetailResponseBodyImageSpec {
	return s.ImageSpec
}

func (s *DescribeUserBenefitDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeUserBenefitDetailResponseBody) GetNetworkSpec() *DescribeUserBenefitDetailResponseBodyNetworkSpec {
	return s.NetworkSpec
}

func (s *DescribeUserBenefitDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBenefitDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeUserBenefitDetailResponseBody) GetUserBaseSpec() *DescribeUserBenefitDetailResponseBodyUserBaseSpec {
	return s.UserBaseSpec
}

func (s *DescribeUserBenefitDetailResponseBody) SetBenefitLevel(v string) *DescribeUserBenefitDetailResponseBody {
	s.BenefitLevel = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetCode(v string) *DescribeUserBenefitDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetExpireAt(v string) *DescribeUserBenefitDetailResponseBody {
	s.ExpireAt = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetHttpStatusCode(v int32) *DescribeUserBenefitDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetImageSpec(v *DescribeUserBenefitDetailResponseBodyImageSpec) *DescribeUserBenefitDetailResponseBody {
	s.ImageSpec = v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetMessage(v string) *DescribeUserBenefitDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetNetworkSpec(v *DescribeUserBenefitDetailResponseBodyNetworkSpec) *DescribeUserBenefitDetailResponseBody {
	s.NetworkSpec = v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetRequestId(v string) *DescribeUserBenefitDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetSuccess(v bool) *DescribeUserBenefitDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) SetUserBaseSpec(v *DescribeUserBenefitDetailResponseBodyUserBaseSpec) *DescribeUserBenefitDetailResponseBody {
	s.UserBaseSpec = v
	return s
}

func (s *DescribeUserBenefitDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserBenefitDetailResponseBodyImageSpec struct {
	EnableCustomized       *bool  `json:"EnableCustomized,omitempty" xml:"EnableCustomized,omitempty"`
	MaxCustomizedBootCount *int32 `json:"MaxCustomizedBootCount,omitempty" xml:"MaxCustomizedBootCount,omitempty"`
}

func (s DescribeUserBenefitDetailResponseBodyImageSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBenefitDetailResponseBodyImageSpec) GoString() string {
	return s.String()
}

func (s *DescribeUserBenefitDetailResponseBodyImageSpec) GetEnableCustomized() *bool {
	return s.EnableCustomized
}

func (s *DescribeUserBenefitDetailResponseBodyImageSpec) GetMaxCustomizedBootCount() *int32 {
	return s.MaxCustomizedBootCount
}

func (s *DescribeUserBenefitDetailResponseBodyImageSpec) SetEnableCustomized(v bool) *DescribeUserBenefitDetailResponseBodyImageSpec {
	s.EnableCustomized = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBodyImageSpec) SetMaxCustomizedBootCount(v int32) *DescribeUserBenefitDetailResponseBodyImageSpec {
	s.MaxCustomizedBootCount = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBodyImageSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeUserBenefitDetailResponseBodyNetworkSpec struct {
	AdvancedBandwidthLimitPerInstance *int32 `json:"AdvancedBandwidthLimitPerInstance,omitempty" xml:"AdvancedBandwidthLimitPerInstance,omitempty"`
	FreeBandwidthLimitPerInstance     *int32 `json:"FreeBandwidthLimitPerInstance,omitempty" xml:"FreeBandwidthLimitPerInstance,omitempty"`
}

func (s DescribeUserBenefitDetailResponseBodyNetworkSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBenefitDetailResponseBodyNetworkSpec) GoString() string {
	return s.String()
}

func (s *DescribeUserBenefitDetailResponseBodyNetworkSpec) GetAdvancedBandwidthLimitPerInstance() *int32 {
	return s.AdvancedBandwidthLimitPerInstance
}

func (s *DescribeUserBenefitDetailResponseBodyNetworkSpec) GetFreeBandwidthLimitPerInstance() *int32 {
	return s.FreeBandwidthLimitPerInstance
}

func (s *DescribeUserBenefitDetailResponseBodyNetworkSpec) SetAdvancedBandwidthLimitPerInstance(v int32) *DescribeUserBenefitDetailResponseBodyNetworkSpec {
	s.AdvancedBandwidthLimitPerInstance = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBodyNetworkSpec) SetFreeBandwidthLimitPerInstance(v int32) *DescribeUserBenefitDetailResponseBodyNetworkSpec {
	s.FreeBandwidthLimitPerInstance = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBodyNetworkSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeUserBenefitDetailResponseBodyUserBaseSpec struct {
	ConcurrencyInstanceLimit *int32 `json:"ConcurrencyInstanceLimit,omitempty" xml:"ConcurrencyInstanceLimit,omitempty"`
	ContextCapacityMB        *int64 `json:"ContextCapacityMB,omitempty" xml:"ContextCapacityMB,omitempty"`
}

func (s DescribeUserBenefitDetailResponseBodyUserBaseSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBenefitDetailResponseBodyUserBaseSpec) GoString() string {
	return s.String()
}

func (s *DescribeUserBenefitDetailResponseBodyUserBaseSpec) GetConcurrencyInstanceLimit() *int32 {
	return s.ConcurrencyInstanceLimit
}

func (s *DescribeUserBenefitDetailResponseBodyUserBaseSpec) GetContextCapacityMB() *int64 {
	return s.ContextCapacityMB
}

func (s *DescribeUserBenefitDetailResponseBodyUserBaseSpec) SetConcurrencyInstanceLimit(v int32) *DescribeUserBenefitDetailResponseBodyUserBaseSpec {
	s.ConcurrencyInstanceLimit = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBodyUserBaseSpec) SetContextCapacityMB(v int64) *DescribeUserBenefitDetailResponseBodyUserBaseSpec {
	s.ContextCapacityMB = &v
	return s
}

func (s *DescribeUserBenefitDetailResponseBodyUserBaseSpec) Validate() error {
	return dara.Validate(s)
}
