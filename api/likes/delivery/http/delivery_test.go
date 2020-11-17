package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/likes/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	likeHttp "park_2020/2020_2_tmp_name/api/likes/delivery/http"
)

func TestNewLikeHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)
	likeHttp.NewLikeHandler(router, mock)
}

func TestUserHandler_LikeHandlerSuccess(t *testing.T) {
	like := models.Like{
		Uid2: 10,
	}
	var byteData = []byte(`{
		"user_id2":       10
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/like", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)
	mock.EXPECT().Like(sid, like).Return(nil)

	likeHandler := likeHttp.LikeHandlerType{
		LUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(likeHandler.LikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestLikeHandler_LikeHandlerFail(t *testing.T) {
	like := models.Like{
		Uid2: 10,
	}
	var byteData = []byte(`{
		"user_id2":       10
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/like", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)
	mock.EXPECT().Like(sid, like).Return(models.ErrInternalServerError)

	likeHandler := likeHttp.LikeHandlerType{
		LUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(likeHandler.LikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestLikeHandler_LikeHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/like", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)

	likeHandler := likeHttp.LikeHandlerType{
		LUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(likeHandler.LikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestLikeHandler_DislikeHandlerSuccess(t *testing.T) {
	dislike := models.Dislike{
		Uid2: 10,
	}
	var byteData = []byte(`{
		"user_id2":       10
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/dislike", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)
	mock.EXPECT().Dislike(sid, dislike).Return(nil)

	likeHandler := likeHttp.LikeHandlerType{
		LUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(likeHandler.DislikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_DislikeHandlerFail(t *testing.T) {
	dislike := models.Dislike{
		Uid2: 10,
	}
	var byteData = []byte(`{
		"user_id2":       10
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/dislike", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)
	mock.EXPECT().Dislike(sid, dislike).Return(models.ErrInternalServerError)

	likeHandler := likeHttp.LikeHandlerType{
		LUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(likeHandler.DislikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_DisLikeHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/dislike", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLikeUsecase(ctrl)

	likeHandler := likeHttp.LikeHandlerType{
		LUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(likeHandler.DislikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}
