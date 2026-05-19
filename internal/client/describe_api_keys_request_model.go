// This file is auto-generated, don't edit it. Thanks.
package client

// DescribeApiKeysRequest is the request struct for DescribeApiKeys
type DescribeApiKeysRequest struct {
	MaxResults *int32   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	KeyIds     []string `json:"KeyIds,omitempty" xml:"KeyIds,omitempty"`
}

// Validate validates the DescribeApiKeysRequest
func (s *DescribeApiKeysRequest) Validate() error {
	return nil
}

// GetMaxResults returns the MaxResults value or 0 if nil
func (s *DescribeApiKeysRequest) GetMaxResults() int32 {
	if s == nil || s.MaxResults == nil {
		return 0
	}
	return *s.MaxResults
}

// SetMaxResults sets the MaxResults value
func (s *DescribeApiKeysRequest) SetMaxResults(v int32) *DescribeApiKeysRequest {
	s.MaxResults = &v
	return s
}

// GetNextToken returns the NextToken value or empty string if nil
func (s *DescribeApiKeysRequest) GetNextToken() string {
	if s == nil || s.NextToken == nil {
		return ""
	}
	return *s.NextToken
}

// SetNextToken sets the NextToken value
func (s *DescribeApiKeysRequest) SetNextToken(v string) *DescribeApiKeysRequest {
	s.NextToken = &v
	return s
}

// GetKeyIds returns the KeyIds slice
func (s *DescribeApiKeysRequest) GetKeyIds() []string {
	if s == nil {
		return nil
	}
	return s.KeyIds
}

// SetKeyIds sets the KeyIds slice
func (s *DescribeApiKeysRequest) SetKeyIds(v []string) *DescribeApiKeysRequest {
	s.KeyIds = v
	return s
}
