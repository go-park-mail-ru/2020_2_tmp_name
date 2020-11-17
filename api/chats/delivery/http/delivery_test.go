package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/domain"
	"park_2020/2020_2_tmp_name/domain/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	userHttp "park_2020/2020_2_tmp_name/users/delivery/http"
)

func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	userHttp.NewUserHandler(router, mock)
}

func TestUserHandler_ChatHandlerSuccess(t *testing.T) {
	chat := models.Chat{
		Uid2:    10,
		LastMsg: "How are you",
	}
	var byteData = []byte(`{
		"user_id2":       10,
		"last_msg" : "How are you"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/chat", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Chat(chat).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_ChatHandlerFail(t *testing.T) {
	chat := models.Chat{
		Uid2:    10,
		LastMsg: "How are you",
	}
	var byteData = []byte(`{
		"user_id2":       10,
		"last_msg" : "How are you"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/chat", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Chat(chat).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_ChatHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/chat", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_MessageHandlerSuccess(t *testing.T) {
	message := models.Message{
		Text:   "How are you",
		ChatID: 2,
	}
	var byteData = []byte(`{
		"text" : "How are you",
		"chat_id" : 2 
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/message", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Message(sid, message).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MessageHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_MessageHandlerFail(t *testing.T) {
	message := models.Message{
		Text:   "How are you",
		ChatID: 2,
	}
	var byteData = []byte(`{
		"text" : "How are you",
		"chat_id" : 2 
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/message", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Message(sid, message).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MessageHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_MessageHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/message", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MessageHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_ChatsHandlerSuccess(t *testing.T) {
	sid := "something-like-this"

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	var chats []models.ChatData

	chat1 := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chat2 := models.ChatData{
		ID: 2,
		Partner: models.UserFeed{
			ID:         4,
			Name:       "Dasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}

	chats = append(chats, chat1, chat2)
	var byteData = []byte(`{
		"id" : 1 
	}`)
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/chats", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var chatModel models.ChatModel
	chatModel.Data = chats

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Chats(sid).Return(chatModel, nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_ChatsHandlerFail(t *testing.T) {
	sid := "something-like-this"

	msg1 := models.Msg{}

	msg2 := models.Msg{}

	var chats []models.ChatData

	chat1 := models.ChatData{
		ID:       1,
		Partner:  models.UserFeed{},
		Messages: []models.Msg{msg1, msg2},
	}

	chat2 := models.ChatData{
		ID:       2,
		Partner:  models.UserFeed{},
		Messages: []models.Msg{msg1, msg2},
	}

	chats = append(chats, chat1, chat2)
	var byteData = []byte(`{
		"id" : 1 
	}`)
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/chats", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var chatModel models.ChatModel
	chatModel.Data = chats

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Chats(sid).Return(chatModel, domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_ChatIDSuccess(t *testing.T) {
	sid := "something-like-this"
	var chid = 1

	msg1 := models.Msg{
		UserID:       1,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:40",
	}

	msg2 := models.Msg{
		UserID:       6,
		ChatID:       3,
		Message:      "Hi",
		TimeDelivery: "18:41",
	}

	chat := models.ChatData{
		ID: 1,
		Partner: models.UserFeed{
			ID:         6,
			Name:       "Natasha",
			DateBirth:  20,
			LinkImages: nil,
			Job:        "",
			Education:  "BMSTU",
			AboutMe:    "",
		},
		Messages: []models.Msg{msg1, msg2},
	}
	var byteData = []byte(`{
		"id" : 1 
	}`)

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/api/v1/chats/1", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().ChatID(sid, chid).Return(chat, nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_ChatIDFail(t *testing.T) {
	sid := "something-like-this"
	var chid = 1

	msg1 := models.Msg{}

	msg2 := models.Msg{}

	chat := models.ChatData{
		ID:       1,
		Partner:  models.UserFeed{},
		Messages: []models.Msg{msg1, msg2},
	}
	var byteData = []byte(`{
		"id" : 1 
	}`)

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/api/v1/chats/1", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().ChatID(sid, chid).Return(chat, domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.ChatIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_GochatFail(t *testing.T) {
	sid := "something-like-this"

	user := models.UserFeed{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("GET", "/gochat", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Gochat(sid).Return(user, domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GochatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}
