// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/microservices/authorization (interfaces: AuthUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockAuthUsecase is a mock of AuthUsecase interface
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// CheckSession mocks base method
func (m *MockAuthUsecase) CheckSession(arg0 context.Context, arg1 string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession
func (mr *MockAuthUsecaseMockRecorder) CheckSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockAuthUsecase)(nil).CheckSession), arg0, arg1)
}

// Login mocks base method
func (m *MockAuthUsecase) Login(arg0 context.Context, arg1 models.LoginData) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthUsecaseMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthUsecase)(nil).Login), arg0, arg1)
}

// Logout mocks base method
func (m *MockAuthUsecase) Logout(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAuthUsecaseMockRecorder) Logout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthUsecase)(nil).Logout), arg0, arg1)
}