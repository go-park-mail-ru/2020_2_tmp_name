// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/chats (interfaces: ChatUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockChatUsecase is a mock of ChatUsecase interface
type MockChatUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockChatUsecaseMockRecorder
}

// MockChatUsecaseMockRecorder is the mock recorder for MockChatUsecase
type MockChatUsecaseMockRecorder struct {
	mock *MockChatUsecase
}

// NewMockChatUsecase creates a new mock instance
func NewMockChatUsecase(ctrl *gomock.Controller) *MockChatUsecase {
	mock := &MockChatUsecase{ctrl: ctrl}
	mock.recorder = &MockChatUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatUsecase) EXPECT() *MockChatUsecaseMockRecorder {
	return m.recorder
}

// Chat mocks base method
func (m *MockChatUsecase) Chat(arg0 models.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chat", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chat indicates an expected call of Chat
func (mr *MockChatUsecaseMockRecorder) Chat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockChatUsecase)(nil).Chat), arg0)
}

// ChatID mocks base method
func (m *MockChatUsecase) ChatID(arg0 string, arg1 int) (models.ChatData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChatID", arg0, arg1)
	ret0, _ := ret[0].(models.ChatData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChatID indicates an expected call of ChatID
func (mr *MockChatUsecaseMockRecorder) ChatID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChatID", reflect.TypeOf((*MockChatUsecase)(nil).ChatID), arg0, arg1)
}

// Chats mocks base method
func (m *MockChatUsecase) Chats(arg0 string) (models.ChatModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chats", arg0)
	ret0, _ := ret[0].(models.ChatModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chats indicates an expected call of Chats
func (mr *MockChatUsecaseMockRecorder) Chats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chats", reflect.TypeOf((*MockChatUsecase)(nil).Chats), arg0)
}

// Gochat mocks base method
func (m *MockChatUsecase) Gochat(arg0 string) (models.UserFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gochat", arg0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Gochat indicates an expected call of Gochat
func (mr *MockChatUsecaseMockRecorder) Gochat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gochat", reflect.TypeOf((*MockChatUsecase)(nil).Gochat), arg0)
}

// Message mocks base method
func (m *MockChatUsecase) Message(arg0 string, arg1 models.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Message indicates an expected call of Message
func (mr *MockChatUsecaseMockRecorder) Message(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockChatUsecase)(nil).Message), arg0, arg1)
}

// ServeWs mocks base method
func (m *MockChatUsecase) ServeWs(arg0 *models.Hub, arg1 http.ResponseWriter, arg2 *http.Request, arg3 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeWs", arg0, arg1, arg2, arg3)
}

// ServeWs indicates an expected call of ServeWs
func (mr *MockChatUsecaseMockRecorder) ServeWs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeWs", reflect.TypeOf((*MockChatUsecase)(nil).ServeWs), arg0, arg1, arg2, arg3)
}