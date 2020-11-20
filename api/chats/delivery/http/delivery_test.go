package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/chats/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	chatHttp "park_2020/2020_2_tmp_name/api/chats/delivery/http"
)

func TestNewchatHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatUsecase(ctrl)
	chatHttp.NewChatHandler(router, mock)
}

func TestChatHandler_ChatHandlerSuccess(t *testing.T) {
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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().Chat(chat).Return(nil)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestChatHandler_ChatHandlerFail(t *testing.T) {
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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().Chat(chat).Return(models.ErrInternalServerError)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestChatHandler_ChatHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/chat", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatUsecase(ctrl)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestChatHandler_MessageHandlerSuccess(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Message(user, message).Return(nil)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.MessageHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestChatHandler_MessageHandlerFail(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Message(user, message).Return(models.ErrInternalServerError)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.MessageHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestChatHandler_MessageHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/message", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockChatUsecase(ctrl)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.MessageHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestChatHandler_ChatsHandlerSuccess(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Chats(user).Return(chatModel, nil)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestChatHandler_ChatsHandlerFail(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Chats(user).Return(chatModel, models.ErrInternalServerError)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestChatHandler_ChatIDSuccess(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().ChatID(user, chid).Return(chat, nil)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestChatHandler_ChatIDFail(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().ChatID(user, chid).Return(chat, models.ErrInternalServerError)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.ChatIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestChatHandler_GochatFail(t *testing.T) {
	sid := "something-like-this"

	user := models.User{}

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

	mock := mock.NewMockChatUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, models.ErrInternalServerError)

	chatHandler := chatHttp.ChatHandlerType{
		ChUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(chatHandler.GochatHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}
