// Code generated by MockGen. DO NOT EDIT.
// Source: park_2020/2020_2_tmp_name/api/chats (interfaces: ChatUsecase)

package mock

import (
	context "context"
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
func (_m *MockChatUsecase) EXPECT() *MockChatUsecaseMockRecorder {
	return _m.recorder
}

// Chat mocks base method
func (_m *MockChatUsecase) Chat(_param0 models.Chat) error {
	ret := _m.ctrl.Call(_m, "Chat", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chat indicates an expected call of Chat
func (_mr *MockChatUsecaseMockRecorder) Chat(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Chat", reflect.TypeOf((*MockChatUsecase)(nil).Chat), arg0)
}

// ChatID mocks base method
func (_m *MockChatUsecase) ChatID(_param0 models.User, _param1 int) (models.ChatData, error) {
	ret := _m.ctrl.Call(_m, "ChatID", _param0, _param1)
	ret0, _ := ret[0].(models.ChatData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChatID indicates an expected call of ChatID
func (_mr *MockChatUsecaseMockRecorder) ChatID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ChatID", reflect.TypeOf((*MockChatUsecase)(nil).ChatID), arg0, arg1)
}

// Chats mocks base method
func (_m *MockChatUsecase) Chats(_param0 models.User) (models.ChatModel, error) {
	ret := _m.ctrl.Call(_m, "Chats", _param0)
	ret0, _ := ret[0].(models.ChatModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chats indicates an expected call of Chats
func (_mr *MockChatUsecaseMockRecorder) Chats(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Chats", reflect.TypeOf((*MockChatUsecase)(nil).Chats), arg0)
}

// CheckSession mocks base method
func (_m *MockChatUsecase) CheckSession(_param0 context.Context, _param1 []*http.Cookie) (models.User, error) {
	ret := _m.ctrl.Call(_m, "CheckSession", _param0, _param1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession
func (_mr *MockChatUsecaseMockRecorder) CheckSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckSession", reflect.TypeOf((*MockChatUsecase)(nil).CheckSession), arg0, arg1)
}

// Dislike mocks base method
func (_m *MockChatUsecase) Dislike(_param0 models.User, _param1 models.Dislike) error {
	ret := _m.ctrl.Call(_m, "Dislike", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dislike indicates an expected call of Dislike
func (_mr *MockChatUsecaseMockRecorder) Dislike(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Dislike", reflect.TypeOf((*MockChatUsecase)(nil).Dislike), arg0, arg1)
}

// Like mocks base method
func (_m *MockChatUsecase) Like(_param0 models.User, _param1 models.Like) error {
	ret := _m.ctrl.Call(_m, "Like", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Like indicates an expected call of Like
func (_mr *MockChatUsecaseMockRecorder) Like(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Like", reflect.TypeOf((*MockChatUsecase)(nil).Like), arg0, arg1)
}

// MatchUser mocks base method
func (_m *MockChatUsecase) MatchUser(_param0 models.User, _param1 models.Like) (models.Chat, bool, error) {
	ret := _m.ctrl.Call(_m, "MatchUser", _param0, _param1)
	ret0, _ := ret[0].(models.Chat)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MatchUser indicates an expected call of MatchUser
func (_mr *MockChatUsecaseMockRecorder) MatchUser(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "MatchUser", reflect.TypeOf((*MockChatUsecase)(nil).MatchUser), arg0, arg1)
}

// Message mocks base method
func (_m *MockChatUsecase) Message(_param0 models.User, _param1 models.Message) error {
	ret := _m.ctrl.Call(_m, "Message", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Message indicates an expected call of Message
func (_mr *MockChatUsecaseMockRecorder) Message(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Message", reflect.TypeOf((*MockChatUsecase)(nil).Message), arg0, arg1)
}

// Msg mocks base method
func (_m *MockChatUsecase) Msg(_param0 models.User, _param1 models.Msg) error {
	ret := _m.ctrl.Call(_m, "Msg", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Msg indicates an expected call of Msg
func (_mr *MockChatUsecaseMockRecorder) Msg(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Msg", reflect.TypeOf((*MockChatUsecase)(nil).Msg), arg0, arg1)
}

// Partner mocks base method
func (_m *MockChatUsecase) Partner(_param0 models.User, _param1 int) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "Partner", _param0, _param1)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partner indicates an expected call of Partner
func (_mr *MockChatUsecaseMockRecorder) Partner(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Partner", reflect.TypeOf((*MockChatUsecase)(nil).Partner), arg0, arg1)
}

// Sessions mocks base method
func (_m *MockChatUsecase) Sessions(_param0 int) ([]string, error) {
	ret := _m.ctrl.Call(_m, "Sessions", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sessions indicates an expected call of Sessions
func (_mr *MockChatUsecaseMockRecorder) Sessions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Sessions", reflect.TypeOf((*MockChatUsecase)(nil).Sessions), arg0)
}

// Superlike mocks base method
func (_m *MockChatUsecase) Superlike(_param0 models.User, _param1 models.Superlike) error {
	ret := _m.ctrl.Call(_m, "Superlike", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Superlike indicates an expected call of Superlike
func (_mr *MockChatUsecaseMockRecorder) Superlike(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Superlike", reflect.TypeOf((*MockChatUsecase)(nil).Superlike), arg0, arg1)
}

// User mocks base method
func (_m *MockChatUsecase) User(_param0 string) (models.User, error) {
	ret := _m.ctrl.Call(_m, "User", _param0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User
func (_mr *MockChatUsecaseMockRecorder) User(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "User", reflect.TypeOf((*MockChatUsecase)(nil).User), arg0)
}

// UserFeed mocks base method
func (_m *MockChatUsecase) UserFeed(_param0 string) (models.UserFeed, error) {
	ret := _m.ctrl.Call(_m, "UserFeed", _param0)
	ret0, _ := ret[0].(models.UserFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserFeed indicates an expected call of UserFeed
func (_mr *MockChatUsecaseMockRecorder) UserFeed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UserFeed", reflect.TypeOf((*MockChatUsecase)(nil).UserFeed), arg0)
}
