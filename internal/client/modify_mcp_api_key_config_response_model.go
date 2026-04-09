// This file is auto-generated, don't edit it. Thanks.
package client

// ModifyMcpApiKeyConfigResponseBody is the response body struct for ModifyMcpApiKeyConfig
type ModifyMcpApiKeyConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

// ModifyMcpApiKeyConfigResponse is the response struct for ModifyMcpApiKeyConfig
type ModifyMcpApiKeyConfigResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *ModifyMcpApiKeyConfigResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *ModifyMcpApiKeyConfigResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetSuccess returns whether the request was successful
func (s *ModifyMcpApiKeyConfigResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}
