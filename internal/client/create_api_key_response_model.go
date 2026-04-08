// This file is auto-generated, don't edit it. Thanks.
package client

// CreateApiKeyResponseBody is the response body struct for CreateApiKey
type CreateApiKeyResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateApiKeyResponseBodyData    `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

// CreateApiKeyResponseBodyData is the data struct for CreateApiKey response
type CreateApiKeyResponseBodyData struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

// CreateApiKeyResponse is the response struct for CreateApiKey
type CreateApiKeyResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *CreateApiKeyResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *CreateApiKeyResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetKeyId returns the KeyId value or empty string if nil
func (s *CreateApiKeyResponseBodyData) GetKeyId() string {
	if s == nil || s.KeyId == nil {
		return ""
	}
	return *s.KeyId
}
