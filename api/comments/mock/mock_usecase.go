// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/comments (interfaces: CommentUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockCommentUsecase is a mock of CommentUsecase interface
type MockCommentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCommentUsecaseMockRecorder
}

// MockCommentUsecaseMockRecorder is the mock recorder for MockCommentUsecase
type MockCommentUsecaseMockRecorder struct {
	mock *MockCommentUsecase
}

// NewMockCommentUsecase creates a new mock instance
func NewMockCommentUsecase(ctrl *gomock.Controller) *MockCommentUsecase {
	mock := &MockCommentUsecase{ctrl: ctrl}
	mock.recorder = &MockCommentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentUsecase) EXPECT() *MockCommentUsecaseMockRecorder {
	return m.recorder
}

// Comment mocks base method
func (m *MockCommentUsecase) Comment(arg0 models.User, arg1 models.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Comment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Comment indicates an expected call of Comment
func (mr *MockCommentUsecaseMockRecorder) Comment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comment", reflect.TypeOf((*MockCommentUsecase)(nil).Comment), arg0, arg1)
}

// CommentsByID mocks base method
func (m *MockCommentUsecase) CommentsByID(arg0 int) (models.CommentsData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommentsByID", arg0)
	ret0, _ := ret[0].(models.CommentsData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentsByID indicates an expected call of CommentsByID
func (mr *MockCommentUsecaseMockRecorder) CommentsByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentsByID", reflect.TypeOf((*MockCommentUsecase)(nil).CommentsByID), arg0)
}

// User mocks base method
func (m *MockCommentUsecase) User(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User
func (mr *MockCommentUsecaseMockRecorder) User(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockCommentUsecase)(nil).User), arg0)
}
