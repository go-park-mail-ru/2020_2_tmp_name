package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/microservices/comments/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	authMock "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client/mock"
	commentClientGRPC "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client"
	commentMock "park_2020/2020_2_tmp_name/microservices/comments/delivery/grpc/client/mock"
	commentHttp "park_2020/2020_2_tmp_name/microservices/comments/delivery/http"
)

func TestNewCommentHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := &authClient.AuthClient{}
	commentClient := &commentClientGRPC.CommentClient{}
	mock := mock.NewMockCommentUsecase(ctrl)
	commentHttp.NewCommentHandler(router, mock, commentClient, authClient)
}

func TestCommentHandler_CommentsByIdHandlerSuccess(t *testing.T) {
	user := models.User{
		ID:         12,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comments := models.CommentsData{}
	outerComments := models.CommentsData{}

	req, err := http.NewRequest("GET", "/api/v1/comments/12", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	commentMock.EXPECT().CommentsByID(context.Background(), user.ID).Return(comments, nil)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
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
	user := models.User{
		ID:         12,
		Name:       "Misha",
		Telephone:  "909-277-47-21",
		Password:   "1234",
		Sex:        "male",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	comments := models.CommentsData{}
	req, err := http.NewRequest("GET", "/api/v1/comments/12", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	commentMock.EXPECT().CommentsByID(context.Background(), user.ID).Return(comments, models.ErrNotFound)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentsByIdHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 404, status)
}

func TestCommentHandler_CommentsByIdHandlerFailAtoi(t *testing.T) {
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

	req, err := http.NewRequest("GET", "/api/v1/comments/nickName", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentsByIdHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestCommentHandler_CommentsByIdHandlerFailUser(t *testing.T) {
	user := models.User{}
	req, err := http.NewRequest("GET", "/api/v1/comments/12", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockCommentUsecase(ctrl)
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentsByIdHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)

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
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	commentMock.EXPECT().Comment(context.Background(), user, comment).Return(nil)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
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
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	commentMock.EXPECT().Comment(context.Background(), user, comment).Return(models.ErrInternalServerError)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestCommentHandler_CommentHandlerFailUser(t *testing.T) {
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
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)
	authMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestCommentHandler_CommentHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
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
	commentMock := commentMock.NewMockCommentClientInterface(ctrl)
	authMock := authMock.NewMockAuthClientInterface(ctrl)

	commentHandler := commentHttp.CommentHandlerType{
		CUsecase:      mock,
		AuthClient:    authMock,
		CommentClient: commentMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(commentHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}
