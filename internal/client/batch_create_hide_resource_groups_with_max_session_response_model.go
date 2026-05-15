// This file is auto-generated, don't edit it. Thanks.
package client

// BatchCreateHideResourceGroupsWithMaxSessionResponseBody is the response body struct for BatchCreateHideResourceGroupsWithMaxSession
type BatchCreateHideResourceGroupsWithMaxSessionResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

// BatchCreateHideResourceGroupsWithMaxSessionResponse is the response struct for BatchCreateHideResourceGroupsWithMaxSession
type BatchCreateHideResourceGroupsWithMaxSessionResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *BatchCreateHideResourceGroupsWithMaxSessionResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *BatchCreateHideResourceGroupsWithMaxSessionResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetSuccess returns whether the request was successful
func (s *BatchCreateHideResourceGroupsWithMaxSessionResponseBody) GetSuccess() bool {
	if s == nil || s.Success == nil {
		return false
	}
	return *s.Success
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *BatchCreateHideResourceGroupsWithMaxSessionResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
