// This file is auto-generated, don't edit it. Thanks.
package client

// DescribeMcpApiKeyResponseBodyData contains the API key detail
type DescribeMcpApiKeyResponseBodyData struct {
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ApiKeyId *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliUid   *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
}

// DescribeMcpApiKeyResponseBody is the response body struct for DescribeMcpApiKey
type DescribeMcpApiKeyResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeMcpApiKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

// DescribeMcpApiKeyResponse is the response struct for DescribeMcpApiKey
type DescribeMcpApiKeyResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DescribeMcpApiKeyResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *DescribeMcpApiKeyResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DescribeMcpApiKeyResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}

// GetSuccess returns the Success value or false if nil
func (s *DescribeMcpApiKeyResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetData returns the Data or nil
func (s *DescribeMcpApiKeyResponseBody) GetData() *DescribeMcpApiKeyResponseBodyData {
	if s == nil {
		return nil
	}
	return s.Data
}

// GetStatus returns the Status value or empty string if nil
func (s *DescribeMcpApiKeyResponseBodyData) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}
	return *s.Status
}

// GetApiKeyId returns the ApiKeyId value or empty string if nil
func (s *DescribeMcpApiKeyResponseBodyData) GetApiKeyId() string {
	if s == nil || s.ApiKeyId == nil {
		return ""
	}
	return *s.ApiKeyId
}

// GetName returns the Name value or empty string if nil
func (s *DescribeMcpApiKeyResponseBodyData) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}
