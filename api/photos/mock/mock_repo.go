// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/photos (interfaces: PhotoRepository)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockPhotoRepository is a mock of PhotoRepository interface
type MockPhotoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoRepositoryMockRecorder
}

// MockPhotoRepositoryMockRecorder is the mock recorder for MockPhotoRepository
type MockPhotoRepositoryMockRecorder struct {
	mock *MockPhotoRepository
}

// NewMockPhotoRepository creates a new mock instance
func NewMockPhotoRepository(ctrl *gomock.Controller) *MockPhotoRepository {
	mock := &MockPhotoRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPhotoRepository) EXPECT() *MockPhotoRepositoryMockRecorder {
	return m.recorder
}

// CheckUserBySession mocks base method
func (m *MockPhotoRepository) CheckUserBySession(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserBySession", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckUserBySession indicates an expected call of CheckUserBySession
func (mr *MockPhotoRepositoryMockRecorder) CheckUserBySession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserBySession", reflect.TypeOf((*MockPhotoRepository)(nil).CheckUserBySession), arg0)
}

// DeletePhoto mocks base method
func (m *MockPhotoRepository) DeletePhoto(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoto", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePhoto indicates an expected call of DeletePhoto
func (mr *MockPhotoRepositoryMockRecorder) DeletePhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).DeletePhoto), arg0, arg1)
}

// InsertPhoto mocks base method
func (m *MockPhotoRepository) InsertPhoto(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPhoto", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPhoto indicates an expected call of InsertPhoto
func (mr *MockPhotoRepositoryMockRecorder) InsertPhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPhoto", reflect.TypeOf((*MockPhotoRepository)(nil).InsertPhoto), arg0, arg1)
}

// SelectImages mocks base method
func (m *MockPhotoRepository) SelectImages(arg0 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectImages", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages
func (mr *MockPhotoRepositoryMockRecorder) SelectImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectImages", reflect.TypeOf((*MockPhotoRepository)(nil).SelectImages), arg0)
}

// SelectUser mocks base method
func (m *MockPhotoRepository) SelectUser(arg0 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUser", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUser indicates an expected call of SelectUser
func (mr *MockPhotoRepositoryMockRecorder) SelectUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUser", reflect.TypeOf((*MockPhotoRepository)(nil).SelectUser), arg0)
}

// SelectUserFeed mocks base method
func (m *MockPhotoRepository) SelectUserFeed(arg0 string) (models.UserFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserFeed", arg0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeed indicates an expected call of SelectUserFeed
func (mr *MockPhotoRepositoryMockRecorder) SelectUserFeed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserFeed", reflect.TypeOf((*MockPhotoRepository)(nil).SelectUserFeed), arg0)
}
