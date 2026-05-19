// This file is auto-generated, don't edit it. Thanks.
package client

// DescribeApiKeysResponse is the response struct for DescribeApiKeys
type DescribeApiKeysResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DescribeApiKeysResponseBody
}

// DescribeApiKeysResponseBody is the response body struct for DescribeApiKeys
type DescribeApiKeysResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeApiKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

// DescribeApiKeysResponseBodyData contains the API key list
type DescribeApiKeysResponseBodyData struct {
	ApiKeys   []*DescribeApiKeysResponseBodyDataApiKey `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Count     *string                                  `json:"Count,omitempty" xml:"Count,omitempty"`
	NextToken *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

// DescribeApiKeysResponseBodyDataApiKey contains a single API key entry
type DescribeApiKeysResponseBodyDataApiKey struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	LastUseDate   *string `json:"LastUseDate,omitempty" xml:"LastUseDate,omitempty"`
	ApiKey        *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Concurrency   *int32  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	KeyId         *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BoundPolicy   *DescribeApiKeysResponseBodyDataApiKeyBoundPolicy   `json:"BoundPolicy,omitempty" xml:"BoundPolicy,omitempty"`
	BoundResource *DescribeApiKeysResponseBodyDataApiKeyBoundResource `json:"BoundResource,omitempty" xml:"BoundResource,omitempty"`
}

// DescribeApiKeysResponseBodyDataApiKeyBoundPolicy contains the bound policy info
type DescribeApiKeysResponseBodyDataApiKeyBoundPolicy struct {
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

// DescribeApiKeysResponseBodyDataApiKeyBoundResource is a placeholder for bound resource info
type DescribeApiKeysResponseBodyDataApiKeyBoundResource struct{}

// --- Nil-safe getters ---

// GetCode returns the Code value or empty string if nil
func (s *DescribeApiKeysResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *DescribeApiKeysResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}

// GetSuccess returns the Success value or false if nil
func (s *DescribeApiKeysResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetMessage returns the Message value or empty string if nil
func (s *DescribeApiKeysResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}

// GetData returns the Data or nil
func (s *DescribeApiKeysResponseBody) GetData() *DescribeApiKeysResponseBodyData {
	if s == nil {
		return nil
	}
	return s.Data
}

// GetApiKeys returns the ApiKeys slice
func (s *DescribeApiKeysResponseBodyData) GetApiKeys() []*DescribeApiKeysResponseBodyDataApiKey {
	if s == nil {
		return nil
	}
	return s.ApiKeys
}

// GetCount returns the Count value or empty string if nil
func (s *DescribeApiKeysResponseBodyData) GetCount() string {
	if s == nil || s.Count == nil {
		return ""
	}
	return *s.Count
}

// GetNextToken returns the NextToken value or empty string if nil
func (s *DescribeApiKeysResponseBodyData) GetNextToken() string {
	if s == nil || s.NextToken == nil {
		return ""
	}
	return *s.NextToken
}

// GetName returns the Name value or empty string if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

// GetStatus returns the Status value or empty string if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}
	return *s.Status
}

// GetKeyId returns the KeyId value or empty string if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetKeyId() string {
	if s == nil || s.KeyId == nil {
		return ""
	}
	return *s.KeyId
}

// GetApiKey returns the ApiKey value or empty string if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetApiKey() string {
	if s == nil || s.ApiKey == nil {
		return ""
	}
	return *s.ApiKey
}

// GetGmtCreate returns the GmtCreate value or empty string if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetGmtCreate() string {
	if s == nil || s.GmtCreate == nil {
		return ""
	}
	return *s.GmtCreate
}

// GetLastUseDate returns the LastUseDate value or empty string if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetLastUseDate() string {
	if s == nil || s.LastUseDate == nil {
		return ""
	}
	return *s.LastUseDate
}

// GetConcurrency returns the Concurrency value or 0 if nil
func (s *DescribeApiKeysResponseBodyDataApiKey) GetConcurrency() int32 {
	if s == nil || s.Concurrency == nil {
		return 0
	}
	return *s.Concurrency
}
