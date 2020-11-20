package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/comments/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	commentHttp "park_2020/2020_2_tmp_name/api/comments/delivery/http"
)

func TestNewCommentHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	commentHttp.NewCommentHandler(router, mock)
}

func TestCommentHandler_CommentsByIdHandlerSuccess(t *testing.T) {
	comments := models.CommentsData{}
	outerComments := models.CommentsData{}
	var ByteData = []byte(`{}`)
	body := bytes.NewReader(ByteData)
	req, err := http.NewRequest("GET", "/api/v1/comments/12", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	mock.EXPECT().CommentsByID(12).Return(comments, nil)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentsByIdHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	err = json.NewDecoder(rr.Body).Decode(&outerComments)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, comments, outerComments)
	require.Equal(t, 200, status)
}

func TestCommentHandler_CommentsByIdHandlerFail(t *testing.T) {
	comments := models.CommentsData{}
	var ByteData = []byte(`{}`)
	body := bytes.NewReader(ByteData)
	req, err := http.NewRequest("GET", "/api/v1/comments/nickName", body)
	if err != nil {
		t.Fatal(err)
	}
	req2, err := http.NewRequest("GET", "/api/v1/comments/12", body)
	if err != nil {
		t.Fatal(err)
	}
	requests := make([]*http.Request, 0, 1)
	requests = append(requests, req)
	requests = append(requests, req2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	mock.EXPECT().CommentsByID(12).Return(comments, errors.New("error"))

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentsByIdHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, requests[i])
		status := rr.Code

		require.NoError(t, err)
		require.Equal(t, 400, status)
	}
}

func TestCommentHandler_CommentHandlerSuccess(t *testing.T) {
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

	comment := models.Comment{
		Uid2:         10,
		TimeDelivery: "18:54",
		CommentText:  "How are you",
	}
	var byteData = []byte(`{
		"user_id2":       10,
		"timeDelivery" : "18:54",
		"CommentText" : "How are you"
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/comment", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Comment(user, comment).Return(nil)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestCommentHandler_CommentHandlerFail(t *testing.T) {
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
	comment := models.Comment{
		Uid2:         10,
		TimeDelivery: "18:54",
		CommentText:  "How are you",
	}
	var byteData = []byte(`{
		"user_id2":       10,
		"timeDelivery" : "18:54",
		"CommentText" : "How are you"
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/comment", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Comment(user, comment).Return(models.ErrInternalServerError)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestCommentHandler_CommentHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/comment", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}
