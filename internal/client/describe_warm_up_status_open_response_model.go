// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

// This file is auto-generated, don't edit it. Thanks.
package client

// DescribeWarmUpStatusOpenResponseBodyDataImage represents a single image in the warm-up status response
type DescribeWarmUpStatusOpenResponseBodyDataImage struct {
	ImageId               *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	TotalMaxSize          *int32  `json:"TotalMaxSize,omitempty" xml:"TotalMaxSize,omitempty"`
	GroupCount            *int32  `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	AvailableInstanceSize *int32  `json:"AvailableInstanceSize,omitempty" xml:"AvailableInstanceSize,omitempty"`
}

// GetImageId returns the ImageId value or empty string if nil
func (s *DescribeWarmUpStatusOpenResponseBodyDataImage) GetImageId() string {
	if s == nil || s.ImageId == nil {
		return ""
	}
	return *s.ImageId
}

// GetTotalMaxSize returns the TotalMaxSize value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyDataImage) GetTotalMaxSize() int32 {
	if s == nil || s.TotalMaxSize == nil {
		return 0
	}
	return *s.TotalMaxSize
}

// GetGroupCount returns the GroupCount value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyDataImage) GetGroupCount() int32 {
	if s == nil || s.GroupCount == nil {
		return 0
	}
	return *s.GroupCount
}

// GetAvailableInstanceSize returns the AvailableInstanceSize value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyDataImage) GetAvailableInstanceSize() int32 {
	if s == nil || s.AvailableInstanceSize == nil {
		return 0
	}
	return *s.AvailableInstanceSize
}

// DescribeWarmUpStatusOpenResponseBodyData represents the Data field in the response
type DescribeWarmUpStatusOpenResponseBodyData struct {
	MaxSessionNumLimit    *int32                                         `json:"MaxSessionNumLimit,omitempty" xml:"MaxSessionNumLimit,omitempty"`
	TotalUsedSessionQuota *int32                                         `json:"TotalUsedSessionQuota,omitempty" xml:"TotalUsedSessionQuota,omitempty"`
	AvailableSessionQuota *int32                                         `json:"AvailableSessionQuota,omitempty" xml:"AvailableSessionQuota,omitempty"`
	MaxImageCount         *int32                                         `json:"MaxImageCount,omitempty" xml:"MaxImageCount,omitempty"`
	CurrentImageCount     *int32                                         `json:"CurrentImageCount,omitempty" xml:"CurrentImageCount,omitempty"`
	Images                []*DescribeWarmUpStatusOpenResponseBodyDataImage `json:"Images,omitempty" xml:"Images,omitempty"`
}

// GetMaxSessionNumLimit returns the MaxSessionNumLimit value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyData) GetMaxSessionNumLimit() int32 {
	if s == nil || s.MaxSessionNumLimit == nil {
		return 0
	}
	return *s.MaxSessionNumLimit
}

// GetTotalUsedSessionQuota returns the TotalUsedSessionQuota value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyData) GetTotalUsedSessionQuota() int32 {
	if s == nil || s.TotalUsedSessionQuota == nil {
		return 0
	}
	return *s.TotalUsedSessionQuota
}

// GetAvailableSessionQuota returns the AvailableSessionQuota value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyData) GetAvailableSessionQuota() int32 {
	if s == nil || s.AvailableSessionQuota == nil {
		return 0
	}
	return *s.AvailableSessionQuota
}

// GetMaxImageCount returns the MaxImageCount value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyData) GetMaxImageCount() int32 {
	if s == nil || s.MaxImageCount == nil {
		return 0
	}
	return *s.MaxImageCount
}

// GetCurrentImageCount returns the CurrentImageCount value or 0 if nil
func (s *DescribeWarmUpStatusOpenResponseBodyData) GetCurrentImageCount() int32 {
	if s == nil || s.CurrentImageCount == nil {
		return 0
	}
	return *s.CurrentImageCount
}

// GetImages returns the Images slice
func (s *DescribeWarmUpStatusOpenResponseBodyData) GetImages() []*DescribeWarmUpStatusOpenResponseBodyDataImage {
	if s == nil {
		return nil
	}
	return s.Images
}

// DescribeWarmUpStatusOpenResponseBody is the response body struct for DescribeWarmUpStatusOpen
type DescribeWarmUpStatusOpenResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeWarmUpStatusOpenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

// GetCode returns the Code value or empty string if nil
func (s *DescribeWarmUpStatusOpenResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId pointer
func (s *DescribeWarmUpStatusOpenResponseBody) GetRequestId() *string {
	if s == nil {
		return nil
	}
	return s.RequestId
}

// GetSuccess returns the Success pointer
func (s *DescribeWarmUpStatusOpenResponseBody) GetSuccess() *bool {
	if s == nil {
		return nil
	}
	return s.Success
}

// GetMessage returns the Message pointer
func (s *DescribeWarmUpStatusOpenResponseBody) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}

// GetData returns the Data field
func (s *DescribeWarmUpStatusOpenResponseBody) GetData() *DescribeWarmUpStatusOpenResponseBodyData {
	if s == nil {
		return nil
	}
	return s.Data
}

// DescribeWarmUpStatusOpenResponse is the response struct for DescribeWarmUpStatusOpen
type DescribeWarmUpStatusOpenResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DescribeWarmUpStatusOpenResponseBody
}

// GetHeaders returns the Headers map
func (s *DescribeWarmUpStatusOpenResponse) GetHeaders() map[string]*string {
	if s == nil {
		return nil
	}
	return s.Headers
}

// GetStatusCode returns the StatusCode pointer
func (s *DescribeWarmUpStatusOpenResponse) GetStatusCode() *int32 {
	if s == nil {
		return nil
	}
	return s.StatusCode
}

// GetBody returns the Body field
func (s *DescribeWarmUpStatusOpenResponse) GetBody() *DescribeWarmUpStatusOpenResponseBody {
	if s == nil {
		return nil
	}
	return s.Body
}
