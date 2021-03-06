// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/chats (interfaces: ChatRepository)

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
func (_m *MockChatRepository) EXPECT() *MockChatRepositoryMockRecorder {
	return _m.recorder
}

// CheckChat mocks base method
func (_m *MockChatRepository) CheckChat(_param0 models.Chat) bool {
	ret := _m.ctrl.Call(_m, "CheckChat", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckChat indicates an expected call of CheckChat
func (_mr *MockChatRepositoryMockRecorder) CheckChat(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckChat", reflect.TypeOf((*MockChatRepository)(nil).CheckChat), arg0)
}

// CheckDislike mocks base method
func (_m *MockChatRepository) CheckDislike(_param0 int, _param1 int) bool {
	ret := _m.ctrl.Call(_m, "CheckDislike", _param0, _param1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckDislike indicates an expected call of CheckDislike
func (_mr *MockChatRepositoryMockRecorder) CheckDislike(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckDislike", reflect.TypeOf((*MockChatRepository)(nil).CheckDislike), arg0, arg1)
}

// CheckLike mocks base method
func (_m *MockChatRepository) CheckLike(_param0 int, _param1 int) bool {
	ret := _m.ctrl.Call(_m, "CheckLike", _param0, _param1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckLike indicates an expected call of CheckLike
func (_mr *MockChatRepositoryMockRecorder) CheckLike(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckLike", reflect.TypeOf((*MockChatRepository)(nil).CheckLike), arg0, arg1)
}

// CheckUserBySession mocks base method
func (_m *MockChatRepository) CheckUserBySession(_param0 string) string {
	ret := _m.ctrl.Call(_m, "CheckUserBySession", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CheckUserBySession indicates an expected call of CheckUserBySession
func (_mr *MockChatRepositoryMockRecorder) CheckUserBySession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckUserBySession", reflect.TypeOf((*MockChatRepository)(nil).CheckUserBySession), arg0)
}

// DeleteDislike mocks base method
func (_m *MockChatRepository) DeleteDislike(_param0 int, _param1 int) error {
	ret := _m.ctrl.Call(_m, "DeleteDislike", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDislike indicates an expected call of DeleteDislike
func (_mr *MockChatRepositoryMockRecorder) DeleteDislike(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteDislike", reflect.TypeOf((*MockChatRepository)(nil).DeleteDislike), arg0, arg1)
}

// DeleteLike mocks base method
func (_m *MockChatRepository) DeleteLike(_param0 int, _param1 int) error {
	ret := _m.ctrl.Call(_m, "DeleteLike", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLike indicates an expected call of DeleteLike
func (_mr *MockChatRepositoryMockRecorder) DeleteLike(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteLike", reflect.TypeOf((*MockChatRepository)(nil).DeleteLike), arg0, arg1)
}

// InsertChat mocks base method
func (_m *MockChatRepository) InsertChat(_param0 models.Chat) error {
	ret := _m.ctrl.Call(_m, "InsertChat", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertChat indicates an expected call of InsertChat
func (_mr *MockChatRepositoryMockRecorder) InsertChat(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertChat", reflect.TypeOf((*MockChatRepository)(nil).InsertChat), arg0)
}

// InsertDislike mocks base method
func (_m *MockChatRepository) InsertDislike(_param0 int, _param1 int, _param2 int) error {
	ret := _m.ctrl.Call(_m, "InsertDislike", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDislike indicates an expected call of InsertDislike
func (_mr *MockChatRepositoryMockRecorder) InsertDislike(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertDislike", reflect.TypeOf((*MockChatRepository)(nil).InsertDislike), arg0, arg1, arg2)
}

// InsertLike mocks base method
func (_m *MockChatRepository) InsertLike(_param0 int, _param1 int, _param2 int) error {
	ret := _m.ctrl.Call(_m, "InsertLike", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertLike indicates an expected call of InsertLike
func (_mr *MockChatRepositoryMockRecorder) InsertLike(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertLike", reflect.TypeOf((*MockChatRepository)(nil).InsertLike), arg0, arg1, arg2)
}

// InsertMessage mocks base method
func (_m *MockChatRepository) InsertMessage(_param0 string, _param1 int, _param2 int) error {
	ret := _m.ctrl.Call(_m, "InsertMessage", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMessage indicates an expected call of InsertMessage
func (_mr *MockChatRepositoryMockRecorder) InsertMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertMessage", reflect.TypeOf((*MockChatRepository)(nil).InsertMessage), arg0, arg1, arg2)
}

// InsertSuperlike mocks base method
func (_m *MockChatRepository) InsertSuperlike(_param0 int, _param1 int, _param2 int) error {
	ret := _m.ctrl.Call(_m, "InsertSuperlike", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertSuperlike indicates an expected call of InsertSuperlike
func (_mr *MockChatRepositoryMockRecorder) InsertSuperlike(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "InsertSuperlike", reflect.TypeOf((*MockChatRepository)(nil).InsertSuperlike), arg0, arg1, arg2)
}

// Match mocks base method
func (_m *MockChatRepository) Match(_param0 int, _param1 int, _param2 int) bool {
	ret := _m.ctrl.Call(_m, "Match", _param0, _param1, _param2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match
func (_mr *MockChatRepositoryMockRecorder) Match(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Match", reflect.TypeOf((*MockChatRepository)(nil).Match), arg0, arg1, arg2)
}

// SelectChatByID mocks base method
func (_m *MockChatRepository) SelectChatByID(_param0 int, _param1 int) (models.ChatData, error) {
	ret := _m.ctrl.Call(_m, "SelectChatByID", _param0, _param1)
	ret0, _ := ret[0].(models.ChatData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectChatByID indicates an expected call of SelectChatByID
func (_mr *MockChatRepositoryMockRecorder) SelectChatByID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectChatByID", reflect.TypeOf((*MockChatRepository)(nil).SelectChatByID), arg0, arg1)
}

// SelectChatID mocks base method
func (_m *MockChatRepository) SelectChatID(_param0 int, _param1 int) (int, error) {
	ret := _m.ctrl.Call(_m, "SelectChatID", _param0, _param1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectChatID indicates an expected call of SelectChatID
func (_mr *MockChatRepositoryMockRecorder) SelectChatID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectChatID", reflect.TypeOf((*MockChatRepository)(nil).SelectChatID), arg0, arg1)
}

// SelectChatsByID mocks base method
func (_m *MockChatRepository) SelectChatsByID(_param0 int) ([]models.ChatData, error) {
	ret := _m.ctrl.Call(_m, "SelectChatsByID", _param0)
	ret0, _ := ret[0].([]models.ChatData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectChatsByID indicates an expected call of SelectChatsByID
func (_mr *MockChatRepositoryMockRecorder) SelectChatsByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectChatsByID", reflect.TypeOf((*MockChatRepository)(nil).SelectChatsByID), arg0)
}

// SelectImages mocks base method
func (_m *MockChatRepository) SelectImages(_param0 int) ([]string, error) {
	ret := _m.ctrl.Call(_m, "SelectImages", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages
func (_mr *MockChatRepositoryMockRecorder) SelectImages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectImages", reflect.TypeOf((*MockChatRepository)(nil).SelectImages), arg0)
}

// SelectMessage mocks base method
func (_m *MockChatRepository) SelectMessage(_param0 int, _param1 int) (models.Msg, error) {
	ret := _m.ctrl.Call(_m, "SelectMessage", _param0, _param1)
	ret0, _ := ret[0].(models.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMessage indicates an expected call of SelectMessage
func (_mr *MockChatRepositoryMockRecorder) SelectMessage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectMessage", reflect.TypeOf((*MockChatRepository)(nil).SelectMessage), arg0, arg1)
}

// SelectMessages mocks base method
func (_m *MockChatRepository) SelectMessages(_param0 int) ([]models.Msg, error) {
	ret := _m.ctrl.Call(_m, "SelectMessages", _param0)
	ret0, _ := ret[0].([]models.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMessages indicates an expected call of SelectMessages
func (_mr *MockChatRepositoryMockRecorder) SelectMessages(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectMessages", reflect.TypeOf((*MockChatRepository)(nil).SelectMessages), arg0)
}

// SelectSessions mocks base method
func (_m *MockChatRepository) SelectSessions(_param0 int) ([]string, error) {
	ret := _m.ctrl.Call(_m, "SelectSessions", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectSessions indicates an expected call of SelectSessions
func (_mr *MockChatRepositoryMockRecorder) SelectSessions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectSessions", reflect.TypeOf((*MockChatRepository)(nil).SelectSessions), arg0)
}

// SelectUserByChat mocks base method
func (_m *MockChatRepository) SelectUserByChat(_param0 int, _param1 int) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "SelectUserByChat", _param0, _param1)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserByChat indicates an expected call of SelectUserByChat
func (_mr *MockChatRepositoryMockRecorder) SelectUserByChat(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectUserByChat", reflect.TypeOf((*MockChatRepository)(nil).SelectUserByChat), arg0, arg1)
}

// SelectUserByID mocks base method
func (_m *MockChatRepository) SelectUserByID(_param0 int) (models.User, error) {
	ret := _m.ctrl.Call(_m, "SelectUserByID", _param0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserByID indicates an expected call of SelectUserByID
func (_mr *MockChatRepositoryMockRecorder) SelectUserByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectUserByID", reflect.TypeOf((*MockChatRepository)(nil).SelectUserByID), arg0)
}

// SelectUserFeed mocks base method
func (_m *MockChatRepository) SelectUserFeed(_param0 string) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "SelectUserFeed", _param0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeed indicates an expected call of SelectUserFeed
func (_mr *MockChatRepositoryMockRecorder) SelectUserFeed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectUserFeed", reflect.TypeOf((*MockChatRepository)(nil).SelectUserFeed), arg0)
}

// SelectUserFeedByID mocks base method
func (_m *MockChatRepository) SelectUserFeedByID(_param0 int) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "SelectUserFeedByID", _param0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserFeedByID indicates an expected call of SelectUserFeedByID
func (_mr *MockChatRepositoryMockRecorder) SelectUserFeedByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SelectUserFeedByID", reflect.TypeOf((*MockChatRepository)(nil).SelectUserFeedByID), arg0)
}
