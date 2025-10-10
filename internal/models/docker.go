// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package models

import "fmt"

// GetDockerFileStoreCredentialRequest represents the request for getting Docker file store credentials
type GetDockerFileStoreCredentialRequest struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s *GetDockerFileStoreCredentialRequest) GetSource() *string {
	return s.Source
}

func (s *GetDockerFileStoreCredentialRequest) SetSource(v string) *GetDockerFileStoreCredentialRequest {
	s.Source = &v
	return s
}

// GetDockerFileStoreCredentialResponse represents the response for getting Docker file store credentials
type GetDockerFileStoreCredentialResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty"`
	Body       *GetDockerFileStoreCredentialResponseBody `json:"body,omitempty"`
}

// GetDockerFileStoreCredentialResponseBody represents the response body
type GetDockerFileStoreCredentialResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDockerFileStoreCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDockerFileStoreCredentialResponseBody) String() string {
	return fmt.Sprintf("%#v", s)
}

func (s *GetDockerFileStoreCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDockerFileStoreCredentialResponseBody) GetData() *GetDockerFileStoreCredentialResponseBodyData {
	return s.Data
}

func (s *GetDockerFileStoreCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

// GetDockerFileStoreCredentialResponseBodyData represents the data part of the response
type GetDockerFileStoreCredentialResponseBodyData struct {
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s *GetDockerFileStoreCredentialResponseBodyData) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetDockerFileStoreCredentialResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

// CreateDockerImageTaskRequest represents the request for creating a Docker image task
type CreateDockerImageTaskRequest struct {
	ImageName     *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s *CreateDockerImageTaskRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateDockerImageTaskRequest) GetSource() *string {
	return s.Source
}

func (s *CreateDockerImageTaskRequest) GetSourceImageId() *string {
	return s.SourceImageId
}

func (s *CreateDockerImageTaskRequest) GetTaskId() *string {
	return s.TaskId
}

// CreateDockerImageTaskResponse represents the response for creating a Docker image task
type CreateDockerImageTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty"`
	Body       *CreateDockerImageTaskResponseBody `json:"body,omitempty"`
}

// CreateDockerImageTaskResponseBody represents the response body
type CreateDockerImageTaskResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateDockerImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s *CreateDockerImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDockerImageTaskResponseBody) GetData() *CreateDockerImageTaskResponseBodyData {
	return s.Data
}

func (s *CreateDockerImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

// CreateDockerImageTaskResponseBodyData represents the data part of the response
type CreateDockerImageTaskResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s *CreateDockerImageTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

// GetDockerImageTaskRequest represents the request for getting Docker image task status
type GetDockerImageTaskRequest struct {
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s *GetDockerImageTaskRequest) GetSource() *string {
	return s.Source
}

func (s *GetDockerImageTaskRequest) GetTaskId() *string {
	return s.TaskId
}

// GetDockerImageTaskResponse represents the response for getting Docker image task status
type GetDockerImageTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty"`
	Body       *GetDockerImageTaskResponseBody `json:"body,omitempty"`
}

// GetDockerImageTaskResponseBody represents the response body
type GetDockerImageTaskResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDockerImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s *GetDockerImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDockerImageTaskResponseBody) GetData() *GetDockerImageTaskResponseBodyData {
	return s.Data
}

func (s *GetDockerImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

// GetDockerImageTaskResponseBodyData represents the data part of the response
type GetDockerImageTaskResponseBodyData struct {
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskMsg *string `json:"TaskMsg,omitempty" xml:"TaskMsg,omitempty"`
}

func (s *GetDockerImageTaskResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *GetDockerImageTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDockerImageTaskResponseBodyData) GetTaskMsg() *string {
	return s.TaskMsg
}
