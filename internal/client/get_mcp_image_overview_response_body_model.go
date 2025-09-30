// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpImageOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpImageOverviewResponseBody
	GetCode() *string
	SetData(v *GetMcpImageOverviewResponseBodyData) *GetMcpImageOverviewResponseBody
	GetData() *GetMcpImageOverviewResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpImageOverviewResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpImageOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpImageOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpImageOverviewResponseBody
	GetSuccess() *bool
}

type GetMcpImageOverviewResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMcpImageOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcpImageOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpImageOverviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpImageOverviewResponseBody) GetData() *GetMcpImageOverviewResponseBodyData {
	return s.Data
}

func (s *GetMcpImageOverviewResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpImageOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpImageOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpImageOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpImageOverviewResponseBody) SetCode(v string) *GetMcpImageOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpImageOverviewResponseBody) SetData(v *GetMcpImageOverviewResponseBodyData) *GetMcpImageOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpImageOverviewResponseBody) SetHttpStatusCode(v int32) *GetMcpImageOverviewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpImageOverviewResponseBody) SetMessage(v string) *GetMcpImageOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpImageOverviewResponseBody) SetRequestId(v string) *GetMcpImageOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpImageOverviewResponseBody) SetSuccess(v bool) *GetMcpImageOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpImageOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcpImageOverviewResponseBodyData struct {
	BrowserUserImageCount  *int32 `json:"BrowserUserImageCount,omitempty" xml:"BrowserUserImageCount,omitempty"`
	CodeSpaceImageCount    *int32 `json:"CodeSpaceImageCount,omitempty" xml:"CodeSpaceImageCount,omitempty"`
	ComputerUserImageCount *int32 `json:"ComputerUserImageCount,omitempty" xml:"ComputerUserImageCount,omitempty"`
	MobileUserImageCount   *int32 `json:"MobileUserImageCount,omitempty" xml:"MobileUserImageCount,omitempty"`
	TotalImageCount        *int32 `json:"TotalImageCount,omitempty" xml:"TotalImageCount,omitempty"`
}

func (s GetMcpImageOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpImageOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpImageOverviewResponseBodyData) GetBrowserUserImageCount() *int32 {
	return s.BrowserUserImageCount
}

func (s *GetMcpImageOverviewResponseBodyData) GetCodeSpaceImageCount() *int32 {
	return s.CodeSpaceImageCount
}

func (s *GetMcpImageOverviewResponseBodyData) GetComputerUserImageCount() *int32 {
	return s.ComputerUserImageCount
}

func (s *GetMcpImageOverviewResponseBodyData) GetMobileUserImageCount() *int32 {
	return s.MobileUserImageCount
}

func (s *GetMcpImageOverviewResponseBodyData) GetTotalImageCount() *int32 {
	return s.TotalImageCount
}

func (s *GetMcpImageOverviewResponseBodyData) SetBrowserUserImageCount(v int32) *GetMcpImageOverviewResponseBodyData {
	s.BrowserUserImageCount = &v
	return s
}

func (s *GetMcpImageOverviewResponseBodyData) SetCodeSpaceImageCount(v int32) *GetMcpImageOverviewResponseBodyData {
	s.CodeSpaceImageCount = &v
	return s
}

func (s *GetMcpImageOverviewResponseBodyData) SetComputerUserImageCount(v int32) *GetMcpImageOverviewResponseBodyData {
	s.ComputerUserImageCount = &v
	return s
}

func (s *GetMcpImageOverviewResponseBodyData) SetMobileUserImageCount(v int32) *GetMcpImageOverviewResponseBodyData {
	s.MobileUserImageCount = &v
	return s
}

func (s *GetMcpImageOverviewResponseBodyData) SetTotalImageCount(v int32) *GetMcpImageOverviewResponseBodyData {
	s.TotalImageCount = &v
	return s
}

func (s *GetMcpImageOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
