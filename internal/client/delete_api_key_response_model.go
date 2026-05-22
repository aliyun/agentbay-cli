// This file is auto-generated, don't edit it. Thanks.
package client

// DeleteApiKeyResponseBody is the response body struct for DeleteApiKey
type DeleteApiKeyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

// DeleteApiKeyResponse is the response struct for DeleteApiKey
type DeleteApiKeyResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DeleteApiKeyResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *DeleteApiKeyResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DeleteApiKeyResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}

// GetSuccess returns the Success value or false if nil
func (s *DeleteApiKeyResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetMessage returns the Message value or empty string if nil
func (s *DeleteApiKeyResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}
