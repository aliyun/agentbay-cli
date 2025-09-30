// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlaygroundConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPlaygroundConfigResponseBody
	GetCode() *string
	SetData(v *GetPlaygroundConfigResponseBodyData) *GetPlaygroundConfigResponseBody
	GetData() *GetPlaygroundConfigResponseBodyData
	SetHttpStatusCode(v int32) *GetPlaygroundConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPlaygroundConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPlaygroundConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPlaygroundConfigResponseBody
	GetSuccess() *bool
}

type GetPlaygroundConfigResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetPlaygroundConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPlaygroundConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPlaygroundConfigResponseBody) GetData() *GetPlaygroundConfigResponseBodyData {
	return s.Data
}

func (s *GetPlaygroundConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPlaygroundConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPlaygroundConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlaygroundConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPlaygroundConfigResponseBody) SetCode(v string) *GetPlaygroundConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetPlaygroundConfigResponseBody) SetData(v *GetPlaygroundConfigResponseBodyData) *GetPlaygroundConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetPlaygroundConfigResponseBody) SetHttpStatusCode(v int32) *GetPlaygroundConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPlaygroundConfigResponseBody) SetMessage(v string) *GetPlaygroundConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetPlaygroundConfigResponseBody) SetRequestId(v string) *GetPlaygroundConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlaygroundConfigResponseBody) SetSuccess(v bool) *GetPlaygroundConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetPlaygroundConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPlaygroundConfigResponseBodyData struct {
	CurrentTask    *GetPlaygroundConfigResponseBodyDataCurrentTask      `json:"CurrentTask,omitempty" xml:"CurrentTask,omitempty" type:"Struct"`
	CurrentTimes   *int32                                               `json:"CurrentTimes,omitempty" xml:"CurrentTimes,omitempty"`
	LimitTimes     *int32                                               `json:"LimitTimes,omitempty" xml:"LimitTimes,omitempty"`
	TemplateScenes []*GetPlaygroundConfigResponseBodyDataTemplateScenes `json:"TemplateScenes,omitempty" xml:"TemplateScenes,omitempty" type:"Repeated"`
}

func (s GetPlaygroundConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponseBodyData) GetCurrentTask() *GetPlaygroundConfigResponseBodyDataCurrentTask {
	return s.CurrentTask
}

func (s *GetPlaygroundConfigResponseBodyData) GetCurrentTimes() *int32 {
	return s.CurrentTimes
}

func (s *GetPlaygroundConfigResponseBodyData) GetLimitTimes() *int32 {
	return s.LimitTimes
}

func (s *GetPlaygroundConfigResponseBodyData) GetTemplateScenes() []*GetPlaygroundConfigResponseBodyDataTemplateScenes {
	return s.TemplateScenes
}

func (s *GetPlaygroundConfigResponseBodyData) SetCurrentTask(v *GetPlaygroundConfigResponseBodyDataCurrentTask) *GetPlaygroundConfigResponseBodyData {
	s.CurrentTask = v
	return s
}

func (s *GetPlaygroundConfigResponseBodyData) SetCurrentTimes(v int32) *GetPlaygroundConfigResponseBodyData {
	s.CurrentTimes = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyData) SetLimitTimes(v int32) *GetPlaygroundConfigResponseBodyData {
	s.LimitTimes = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyData) SetTemplateScenes(v []*GetPlaygroundConfigResponseBodyDataTemplateScenes) *GetPlaygroundConfigResponseBodyData {
	s.TemplateScenes = v
	return s
}

func (s *GetPlaygroundConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPlaygroundConfigResponseBodyDataCurrentTask struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetPlaygroundConfigResponseBodyDataCurrentTask) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponseBodyDataCurrentTask) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponseBodyDataCurrentTask) GetSessionId() *string {
	return s.SessionId
}

func (s *GetPlaygroundConfigResponseBodyDataCurrentTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPlaygroundConfigResponseBodyDataCurrentTask) SetSessionId(v string) *GetPlaygroundConfigResponseBodyDataCurrentTask {
	s.SessionId = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataCurrentTask) SetTaskId(v string) *GetPlaygroundConfigResponseBodyDataCurrentTask {
	s.TaskId = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataCurrentTask) Validate() error {
	return dara.Validate(s)
}

type GetPlaygroundConfigResponseBodyDataTemplateScenes struct {
	ExampleList  []*GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList  `json:"ExampleList,omitempty" xml:"ExampleList,omitempty" type:"Repeated"`
	SceneName    *string                                                          `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	SceneType    *string                                                          `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	TemplateList []*GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
}

func (s GetPlaygroundConfigResponseBodyDataTemplateScenes) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponseBodyDataTemplateScenes) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) GetExampleList() []*GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList {
	return s.ExampleList
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) GetSceneName() *string {
	return s.SceneName
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) GetSceneType() *string {
	return s.SceneType
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) GetTemplateList() []*GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList {
	return s.TemplateList
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) SetExampleList(v []*GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) *GetPlaygroundConfigResponseBodyDataTemplateScenes {
	s.ExampleList = v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) SetSceneName(v string) *GetPlaygroundConfigResponseBodyDataTemplateScenes {
	s.SceneName = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) SetSceneType(v string) *GetPlaygroundConfigResponseBodyDataTemplateScenes {
	s.SceneType = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) SetTemplateList(v []*GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) *GetPlaygroundConfigResponseBodyDataTemplateScenes {
	s.TemplateList = v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenes) Validate() error {
	return dara.Validate(s)
}

type GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList struct {
	ExampleId   *string `json:"ExampleId,omitempty" xml:"ExampleId,omitempty"`
	ExampleName *string `json:"ExampleName,omitempty" xml:"ExampleName,omitempty"`
}

func (s GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) GetExampleId() *string {
	return s.ExampleId
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) GetExampleName() *string {
	return s.ExampleName
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) SetExampleId(v string) *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList {
	s.ExampleId = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) SetExampleName(v string) *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList {
	s.ExampleName = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesExampleList) Validate() error {
	return dara.Validate(s)
}

type GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList struct {
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) String() string {
	return dara.Prettify(s)
}

func (s GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) GoString() string {
	return s.String()
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) SetTemplateId(v string) *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList {
	s.TemplateId = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) SetTemplateName(v string) *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList {
	s.TemplateName = &v
	return s
}

func (s *GetPlaygroundConfigResponseBodyDataTemplateScenesTemplateList) Validate() error {
	return dara.Validate(s)
}
