// This file is auto-generated, don't edit it. Thanks.
package client

// DescribeKeyContentResponseBodyData is the data payload for DescribeKeyContent
type DescribeKeyContentResponseBodyData struct {
	ApiKey    *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

// DescribeKeyContentResponseBody is the response body struct for DescribeKeyContent
type DescribeKeyContentResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeKeyContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

// DescribeKeyContentResponse is the response struct for DescribeKeyContent
type DescribeKeyContentResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DescribeKeyContentResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *DescribeKeyContentResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DescribeKeyContentResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}

// GetSuccess returns the Success value or false if nil
func (s *DescribeKeyContentResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetMessage returns the Message value or empty string if nil
func (s *DescribeKeyContentResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}

// GetApiKey returns the ApiKey value or empty string if nil
func (s *DescribeKeyContentResponseBodyData) GetApiKey() string {
	if s == nil || s.ApiKey == nil {
		return ""
	}
	return *s.ApiKey
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DescribeKeyContentResponseBodyData) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
