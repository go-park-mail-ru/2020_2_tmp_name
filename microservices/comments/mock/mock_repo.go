// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/microservices/comments (interfaces: CommentRepository)

package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "park_2020/2020_2_tmp_name/models"
	reflect "reflect"
)

// MockCommentRepository is a mock of CommentRepository interface
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return _m.recorder
}

// InsertComment mocks base method
func (_m *MockCommentRepository) InsertComment(_param0 models.Comment, _param1 int) error {
	ret := _m.ctrl.Call(_m, "InsertComment", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertComment indicates an expected call of InsertComment
func (_mr *MockCommentRepositoryMockRecorder) InsertComment(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertComment", reflect.TypeOf((*MockCommentRepository)(nil).InsertComment), arg0, arg1)
}

// SelectComments mocks base method
func (_m *MockCommentRepository) SelectComments(_param0 int) (models.CommentsById, error) {
	ret := _m.ctrl.Call(_m, "SelectComments", _param0)
	ret0, _ := ret[0].(models.CommentsById)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectComments indicates an expected call of SelectComments
func (_mr *MockCommentRepositoryMockRecorder) SelectComments(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectComments", reflect.TypeOf((*MockCommentRepository)(nil).SelectComments), arg0)
}

// SelectImages mocks base method
func (_m *MockCommentRepository) SelectImages(_param0 int) ([]string, error) {
	ret := _m.ctrl.Call(_m, "SelectImages", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages
func (_mr *MockCommentRepositoryMockRecorder) SelectImages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectImages", reflect.TypeOf((*MockCommentRepository)(nil).SelectImages), arg0)
}

// SelectUserFeed mocks base method
func (_m *MockCommentRepository) SelectUserFeed(_param0 string) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "SelectUserFeed", _param0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeed indicates an expected call of SelectUserFeed
func (_mr *MockCommentRepositoryMockRecorder) SelectUserFeed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectUserFeed", reflect.TypeOf((*MockCommentRepository)(nil).SelectUserFeed), arg0)
}

// SelectUserFeedByID mocks base method
func (_m *MockCommentRepository) SelectUserFeedByID(_param0 int) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "SelectUserFeedByID", _param0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeedByID indicates an expected call of SelectUserFeedByID
func (_mr *MockCommentRepositoryMockRecorder) SelectUserFeedByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectUserFeedByID", reflect.TypeOf((*MockCommentRepository)(nil).SelectUserFeedByID), arg0)
}
