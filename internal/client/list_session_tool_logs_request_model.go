// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionToolLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListSessionToolLogsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListSessionToolLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSessionToolLogsRequest
	GetNextToken() *string
	SetSessionId(v string) *ListSessionToolLogsRequest
	GetSessionId() *string
	SetStartTime(v string) *ListSessionToolLogsRequest
	GetStartTime() *string
}

type ListSessionToolLogsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSessionToolLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionToolLogsRequest) GoString() string {
	return s.String()
}

func (s *ListSessionToolLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSessionToolLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSessionToolLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionToolLogsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSessionToolLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSessionToolLogsRequest) SetEndTime(v string) *ListSessionToolLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSessionToolLogsRequest) SetMaxResults(v int32) *ListSessionToolLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSessionToolLogsRequest) SetNextToken(v string) *ListSessionToolLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSessionToolLogsRequest) SetSessionId(v string) *ListSessionToolLogsRequest {
	s.SessionId = &v
	return s
}

func (s *ListSessionToolLogsRequest) SetStartTime(v string) *ListSessionToolLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSessionToolLogsRequest) Validate() error {
	return dara.Validate(s)
}
