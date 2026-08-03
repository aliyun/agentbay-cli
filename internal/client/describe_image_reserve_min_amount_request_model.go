// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

// DescribeImageReserveMinAmountRequest is the request struct for DescribeImageReserveMinAmount
type DescribeImageReserveMinAmountRequest struct {
	ImageIds   []string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	NextToken  *string  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

// Validate validates the DescribeImageReserveMinAmountRequest
func (s *DescribeImageReserveMinAmountRequest) Validate() error {
	return nil
}

// SetImageIds sets the ImageIds field
func (s *DescribeImageReserveMinAmountRequest) SetImageIds(v []string) *DescribeImageReserveMinAmountRequest {
	s.ImageIds = v
	return s
}

// GetImageIds returns the ImageIds slice
func (s *DescribeImageReserveMinAmountRequest) GetImageIds() []string {
	if s == nil {
		return nil
	}
	return s.ImageIds
}

// SetNextToken sets the NextToken field
func (s *DescribeImageReserveMinAmountRequest) SetNextToken(v string) *DescribeImageReserveMinAmountRequest {
	s.NextToken = &v
	return s
}

// GetNextToken returns the NextToken value or empty string if nil
func (s *DescribeImageReserveMinAmountRequest) GetNextToken() string {
	if s == nil || s.NextToken == nil {
		return ""
	}
	return *s.NextToken
}

// SetMaxResults sets the MaxResults field
func (s *DescribeImageReserveMinAmountRequest) SetMaxResults(v int32) *DescribeImageReserveMinAmountRequest {
	s.MaxResults = &v
	return s
}

// GetMaxResults returns the MaxResults value or 0 if nil
func (s *DescribeImageReserveMinAmountRequest) GetMaxResults() int32 {
	if s == nil || s.MaxResults == nil {
		return 0
	}
	return *s.MaxResults
}
