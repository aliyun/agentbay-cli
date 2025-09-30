// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *UpdateConcurrencyRequest
	GetApiKeyId() *string
	SetConcurrency(v string) *UpdateConcurrencyRequest
	GetConcurrency() *string
}

type UpdateConcurrencyRequest struct {
	ApiKeyId    *string `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	Concurrency *string `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s UpdateConcurrencyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *UpdateConcurrencyRequest) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *UpdateConcurrencyRequest) GetConcurrency() *string {
	return s.Concurrency
}

func (s *UpdateConcurrencyRequest) SetApiKeyId(v string) *UpdateConcurrencyRequest {
	s.ApiKeyId = &v
	return s
}

func (s *UpdateConcurrencyRequest) SetConcurrency(v string) *UpdateConcurrencyRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
