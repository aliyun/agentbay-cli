// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListSessionLogsRequest
	GetEndTime() *string
	SetImageId(v string) *ListSessionLogsRequest
	GetImageId() *string
	SetImageType(v string) *ListSessionLogsRequest
	GetImageType() *string
	SetMaxResults(v int32) *ListSessionLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSessionLogsRequest
	GetNextToken() *string
	SetSessionId(v string) *ListSessionLogsRequest
	GetSessionId() *string
	SetStartTime(v string) *ListSessionLogsRequest
	GetStartTime() *string
	SetStatus(v string) *ListSessionLogsRequest
	GetStatus() *string
}

type ListSessionLogsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType  *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSessionLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionLogsRequest) GoString() string {
	return s.String()
}

func (s *ListSessionLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSessionLogsRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListSessionLogsRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListSessionLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSessionLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionLogsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSessionLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSessionLogsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSessionLogsRequest) SetEndTime(v string) *ListSessionLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSessionLogsRequest) SetImageId(v string) *ListSessionLogsRequest {
	s.ImageId = &v
	return s
}

func (s *ListSessionLogsRequest) SetImageType(v string) *ListSessionLogsRequest {
	s.ImageType = &v
	return s
}

func (s *ListSessionLogsRequest) SetMaxResults(v int32) *ListSessionLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSessionLogsRequest) SetNextToken(v string) *ListSessionLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSessionLogsRequest) SetSessionId(v string) *ListSessionLogsRequest {
	s.SessionId = &v
	return s
}

func (s *ListSessionLogsRequest) SetStartTime(v string) *ListSessionLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSessionLogsRequest) SetStatus(v string) *ListSessionLogsRequest {
	s.Status = &v
	return s
}

func (s *ListSessionLogsRequest) Validate() error {
	return dara.Validate(s)
}
