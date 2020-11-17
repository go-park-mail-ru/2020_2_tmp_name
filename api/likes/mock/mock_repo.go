// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/likes (interfaces: LikeRepository)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockLikeRepository is a mock of LikeRepository interface
type MockLikeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLikeRepositoryMockRecorder
}

// MockLikeRepositoryMockRecorder is the mock recorder for MockLikeRepository
type MockLikeRepositoryMockRecorder struct {
	mock *MockLikeRepository
}

// NewMockLikeRepository creates a new mock instance
func NewMockLikeRepository(ctrl *gomock.Controller) *MockLikeRepository {
	mock := &MockLikeRepository{ctrl: ctrl}
	mock.recorder = &MockLikeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLikeRepository) EXPECT() *MockLikeRepositoryMockRecorder {
	return m.recorder
}

// CheckChat mocks base method
func (m *MockLikeRepository) CheckChat(arg0 models.Chat) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckChat", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckChat indicates an expected call of CheckChat
func (mr *MockLikeRepositoryMockRecorder) CheckChat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckChat", reflect.TypeOf((*MockLikeRepository)(nil).CheckChat), arg0)
}

// CheckUserBySession mocks base method
func (m *MockLikeRepository) CheckUserBySession(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserBySession", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckUserBySession indicates an expected call of CheckUserBySession
func (mr *MockLikeRepositoryMockRecorder) CheckUserBySession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserBySession", reflect.TypeOf((*MockLikeRepository)(nil).CheckUserBySession), arg0)
}

// InsertChat mocks base method
func (m *MockLikeRepository) InsertChat(arg0 models.Chat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertChat", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertChat indicates an expected call of InsertChat
func (mr *MockLikeRepositoryMockRecorder) InsertChat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertChat", reflect.TypeOf((*MockLikeRepository)(nil).InsertChat), arg0)
}

// InsertDislike mocks base method
func (m *MockLikeRepository) InsertDislike(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDislike", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDislike indicates an expected call of InsertDislike
func (mr *MockLikeRepositoryMockRecorder) InsertDislike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDislike", reflect.TypeOf((*MockLikeRepository)(nil).InsertDislike), arg0, arg1)
}

// InsertLike mocks base method
func (m *MockLikeRepository) InsertLike(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertLike", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertLike indicates an expected call of InsertLike
func (mr *MockLikeRepositoryMockRecorder) InsertLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertLike", reflect.TypeOf((*MockLikeRepository)(nil).InsertLike), arg0, arg1)
}

// Match mocks base method
func (m *MockLikeRepository) Match(arg0, arg1 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match
func (mr *MockLikeRepositoryMockRecorder) Match(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockLikeRepository)(nil).Match), arg0, arg1)
}

// SelectImages mocks base method
func (m *MockLikeRepository) SelectImages(arg0 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectImages", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages
func (mr *MockLikeRepositoryMockRecorder) SelectImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectImages", reflect.TypeOf((*MockLikeRepository)(nil).SelectImages), arg0)
}

// SelectUserFeed mocks base method
func (m *MockLikeRepository) SelectUserFeed(arg0 string) (models.UserFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserFeed", arg0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeed indicates an expected call of SelectUserFeed
func (mr *MockLikeRepositoryMockRecorder) SelectUserFeed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserFeed", reflect.TypeOf((*MockLikeRepository)(nil).SelectUserFeed), arg0)
}
