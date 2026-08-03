// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

// DescribeImageReserveMinAmountResourceGroup represents a single resource group in the response
type DescribeImageReserveMinAmountResourceGroup struct {
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	ReserveMinAmount   *int32  `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	MaxAmount          *int32  `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	ResourceGroupType  *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

// GetResourceGroupId returns the ResourceGroupId value or empty string if nil
func (s *DescribeImageReserveMinAmountResourceGroup) GetResourceGroupId() string {
	if s == nil || s.ResourceGroupId == nil {
		return ""
	}
	return *s.ResourceGroupId
}

// GetAppInstanceGroupId returns the AppInstanceGroupId value or empty string if nil
func (s *DescribeImageReserveMinAmountResourceGroup) GetAppInstanceGroupId() string {
	if s == nil || s.AppInstanceGroupId == nil {
		return ""
	}
	return *s.AppInstanceGroupId
}

// GetReserveMinAmount returns the ReserveMinAmount value or 0 if nil
func (s *DescribeImageReserveMinAmountResourceGroup) GetReserveMinAmount() int32 {
	if s == nil || s.ReserveMinAmount == nil {
		return 0
	}
	return *s.ReserveMinAmount
}

// GetMaxAmount returns the MaxAmount value or 0 if nil
func (s *DescribeImageReserveMinAmountResourceGroup) GetMaxAmount() int32 {
	if s == nil || s.MaxAmount == nil {
		return 0
	}
	return *s.MaxAmount
}

// GetResourceGroupType returns the ResourceGroupType value or empty string if nil
func (s *DescribeImageReserveMinAmountResourceGroup) GetResourceGroupType() string {
	if s == nil || s.ResourceGroupType == nil {
		return ""
	}
	return *s.ResourceGroupType
}

// GetStatus returns the Status value or empty string if nil
func (s *DescribeImageReserveMinAmountResourceGroup) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}
	return *s.Status
}

// DescribeImageReserveMinAmountImage represents a single image in the response
type DescribeImageReserveMinAmountImage struct {
	ImageId        *string                                       `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ResourceGroups []*DescribeImageReserveMinAmountResourceGroup `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty"`
}

// GetImageId returns the ImageId value or empty string if nil
func (s *DescribeImageReserveMinAmountImage) GetImageId() string {
	if s == nil || s.ImageId == nil {
		return ""
	}
	return *s.ImageId
}

// GetResourceGroups returns the ResourceGroups slice
func (s *DescribeImageReserveMinAmountImage) GetResourceGroups() []*DescribeImageReserveMinAmountResourceGroup {
	if s == nil {
		return nil
	}
	return s.ResourceGroups
}

// DescribeImageReserveMinAmountResponseBodyData represents the Data field in the response
type DescribeImageReserveMinAmountResponseBodyData struct {
	Images    []*DescribeImageReserveMinAmountImage `json:"Images,omitempty" xml:"Images,omitempty"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

// GetImages returns the Images slice
func (s *DescribeImageReserveMinAmountResponseBodyData) GetImages() []*DescribeImageReserveMinAmountImage {
	if s == nil {
		return nil
	}
	return s.Images
}

// GetNextToken returns the NextToken value or empty string if nil
func (s *DescribeImageReserveMinAmountResponseBodyData) GetNextToken() string {
	if s == nil || s.NextToken == nil {
		return ""
	}
	return *s.NextToken
}

// DescribeImageReserveMinAmountResponseBody is the response body struct for DescribeImageReserveMinAmount
type DescribeImageReserveMinAmountResponseBody struct {
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeImageReserveMinAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

// GetCode returns the Code value or empty string if nil
func (s *DescribeImageReserveMinAmountResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DescribeImageReserveMinAmountResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}

// GetSuccess returns whether the request was successful
func (s *DescribeImageReserveMinAmountResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetMessage returns the Message pointer
func (s *DescribeImageReserveMinAmountResponseBody) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}

// GetData returns the Data field
func (s *DescribeImageReserveMinAmountResponseBody) GetData() *DescribeImageReserveMinAmountResponseBodyData {
	if s == nil {
		return nil
	}
	return s.Data
}

// DescribeImageReserveMinAmountResponse is the response struct for DescribeImageReserveMinAmount
type DescribeImageReserveMinAmountResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DescribeImageReserveMinAmountResponseBody
}

// GetHeaders returns the Headers map
func (s *DescribeImageReserveMinAmountResponse) GetHeaders() map[string]*string {
	if s == nil {
		return nil
	}
	return s.Headers
}

// GetStatusCode returns the StatusCode pointer
func (s *DescribeImageReserveMinAmountResponse) GetStatusCode() *int32 {
	if s == nil {
		return nil
	}
	return s.StatusCode
}

// GetBody returns the Body field
func (s *DescribeImageReserveMinAmountResponse) GetBody() *DescribeImageReserveMinAmountResponseBody {
	if s == nil {
		return nil
	}
	return s.Body
}
