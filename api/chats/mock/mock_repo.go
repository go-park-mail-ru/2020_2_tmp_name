// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/chats (interfaces: ChatRepository)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockChatRepository is a mock of ChatRepository interface
type MockChatRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatRepositoryMockRecorder
}

// MockChatRepositoryMockRecorder is the mock recorder for MockChatRepository
type MockChatRepositoryMockRecorder struct {
	mock *MockChatRepository
}

// NewMockChatRepository creates a new mock instance
func NewMockChatRepository(ctrl *gomock.Controller) *MockChatRepository {
	mock := &MockChatRepository{ctrl: ctrl}
	mock.recorder = &MockChatRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatRepository) EXPECT() *MockChatRepositoryMockRecorder {
	return m.recorder
}

// CheckChat mocks base method
func (m *MockChatRepository) CheckChat(arg0 models.Chat) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckChat", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckChat indicates an expected call of CheckChat
func (mr *MockChatRepositoryMockRecorder) CheckChat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckChat", reflect.TypeOf((*MockChatRepository)(nil).CheckChat), arg0)
}

// CheckUserBySession mocks base method
func (m *MockChatRepository) CheckUserBySession(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserBySession", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckUserBySession indicates an expected call of CheckUserBySession
func (mr *MockChatRepositoryMockRecorder) CheckUserBySession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserBySession", reflect.TypeOf((*MockChatRepository)(nil).CheckUserBySession), arg0)
}

// InsertChat mocks base method
func (m *MockChatRepository) InsertChat(arg0 models.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertChat", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertChat indicates an expected call of InsertChat
func (mr *MockChatRepositoryMockRecorder) InsertChat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertChat", reflect.TypeOf((*MockChatRepository)(nil).InsertChat), arg0)
}

// InsertMessage mocks base method
func (m *MockChatRepository) InsertMessage(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMessage indicates an expected call of InsertMessage
func (mr *MockChatRepositoryMockRecorder) InsertMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMessage", reflect.TypeOf((*MockChatRepository)(nil).InsertMessage), arg0, arg1, arg2)
}

// SelectChatByID mocks base method
func (m *MockChatRepository) SelectChatByID(arg0, arg1 int) (models.ChatData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectChatByID", arg0, arg1)
	ret0, _ := ret[0].(models.ChatData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectChatByID indicates an expected call of SelectChatByID
func (mr *MockChatRepositoryMockRecorder) SelectChatByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectChatByID", reflect.TypeOf((*MockChatRepository)(nil).SelectChatByID), arg0, arg1)
}

// SelectChatsByID mocks base method
func (m *MockChatRepository) SelectChatsByID(arg0 int) ([]models.ChatData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectChatsByID", arg0)
	ret0, _ := ret[0].([]models.ChatData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectChatsByID indicates an expected call of SelectChatsByID
func (mr *MockChatRepositoryMockRecorder) SelectChatsByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectChatsByID", reflect.TypeOf((*MockChatRepository)(nil).SelectChatsByID), arg0)
}

// SelectImages mocks base method
func (m *MockChatRepository) SelectImages(arg0 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectImages", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages
func (mr *MockChatRepositoryMockRecorder) SelectImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectImages", reflect.TypeOf((*MockChatRepository)(nil).SelectImages), arg0)
}

// SelectMessage mocks base method
func (m *MockChatRepository) SelectMessage(arg0, arg1 int) (models.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMessage", arg0, arg1)
	ret0, _ := ret[0].(models.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMessage indicates an expected call of SelectMessage
func (mr *MockChatRepositoryMockRecorder) SelectMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMessage", reflect.TypeOf((*MockChatRepository)(nil).SelectMessage), arg0, arg1)
}

// SelectMessages mocks base method
func (m *MockChatRepository) SelectMessages(arg0 int) ([]models.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMessages", arg0)
	ret0, _ := ret[0].([]models.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMessages indicates an expected call of SelectMessages
func (mr *MockChatRepositoryMockRecorder) SelectMessages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMessages", reflect.TypeOf((*MockChatRepository)(nil).SelectMessages), arg0)
}

// SelectUserByChat mocks base method
func (m *MockChatRepository) SelectUserByChat(arg0, arg1 int) (models.UserFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserByChat", arg0, arg1)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserByChat indicates an expected call of SelectUserByChat
func (mr *MockChatRepositoryMockRecorder) SelectUserByChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserByChat", reflect.TypeOf((*MockChatRepository)(nil).SelectUserByChat), arg0, arg1)
}

// SelectUserFeed mocks base method
func (m *MockChatRepository) SelectUserFeed(arg0 string) (models.UserFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserFeed", arg0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeed indicates an expected call of SelectUserFeed
func (mr *MockChatRepositoryMockRecorder) SelectUserFeed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserFeed", reflect.TypeOf((*MockChatRepository)(nil).SelectUserFeed), arg0)
}

// SelectUserFeedByID mocks base method
func (m *MockChatRepository) SelectUserFeedByID(arg0 int) (models.UserFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserFeedByID", arg0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeedByID indicates an expected call of SelectUserFeedByID
func (mr *MockChatRepositoryMockRecorder) SelectUserFeedByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserFeedByID", reflect.TypeOf((*MockChatRepository)(nil).SelectUserFeedByID), arg0)
}