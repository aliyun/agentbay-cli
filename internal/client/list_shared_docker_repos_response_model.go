// This file is auto-generated, don't edit it. Thanks.
package client

// ListSharedDockerReposResponseBodyDataItem is an item in the Data array of ListSharedDockerRepos response
type ListSharedDockerReposResponseBodyDataItem struct {
	PeerAliUid *int64  `json:"PeerAliUid,omitempty" xml:"PeerAliUid,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

// ListSharedDockerReposResponseBody is the response body struct for ListSharedDockerRepos
type ListSharedDockerReposResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Data           []*ListSharedDockerReposResponseBodyDataItem `json:"Data,omitempty" xml:"Data>object,omitempty"`
}

// ListSharedDockerReposResponse is the response struct for ListSharedDockerRepos
type ListSharedDockerReposResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *ListSharedDockerReposResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *ListSharedDockerReposResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetMessage returns the Message value or empty string if nil
func (s *ListSharedDockerReposResponseBody) GetMessage() string {
	if s == nil || s.Message == nil {
		return ""
	}
	return *s.Message
}

// GetRequestId returns the RequestId value or empty string if nil
func (s *ListSharedDockerReposResponseBody) GetRequestId() string {
	if s == nil || s.RequestId == nil {
		return ""
	}
	return *s.RequestId
}
