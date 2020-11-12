package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
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

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	userHandler := &userHttp.UserHandler{}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUserHandler_LoginHandlerSuccess(t *testing.T) {
	var byteData = []byte(`{
			"telephone" : "909-277-47-21",
			"password" : "qwerty"
		}`)

	login := models.LoginData{
		Telephone: "909-277-47-21",
		Password:  "qwerty",
	}
	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Login(login).Return("some uuid", nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 200, status)

}

func TestUserHandler_LoginHandlerFail(t *testing.T) {
	var byteData = []byte(``)
	var byteData2 = []byte(`{
			"telephone" : "909-277-47-21",
			"password" : "qwerty"
		}`)

	body := bytes.NewReader(byteData)
	body_sec := bytes.NewReader(byteData2)

	login := models.LoginData{
		Telephone: "909-277-47-21",
		Password:  "qwerty",
	}

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}
	req2, err := http.NewRequest("POST", "/login", body_sec)
	if err != nil {
		t.Fatal(err)
	}
	requests := make([]*http.Request, 0, 1)
	requests = append(requests, req)
	requests = append(requests, req2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	gomock.InOrder(
		mock.EXPECT().Login(login).Return("some uuid", errors.New("error uuid")),
	)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, requests[i])
		status := rr.Code

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}

}
func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	userHttp.NewUserHandler(router, mock)
}

func TestUserHandler_LogoutHandlerSuccess(t *testing.T) {
	var byteData = []byte(`{}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/logout", body)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "some uuid",
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Logout(cookie.Value).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LogoutHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 200, status)
}

func TestUserHandler_LogoutHandlerFail(t *testing.T) {
	var byteData = []byte(`{}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/logout", body)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "some uuid",
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Logout(cookie.Value).Return(errors.New("error"))

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LogoutHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, req)
		status := rr.Code
		req.AddCookie(cookie)

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}
}

func TestUserHandler_SignupHandlerSuccess(t *testing.T) {
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
		"id":         "0",
		"name":       "Misha",
		"telephone":  "909-277-47-21",
		"password":   "1234",
		"sex":        "male",
		"linkImages": null,
		"job":        "Fullstack",
		"education":  "BMSTU",
		"aboutMe":    ""
	}`)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Signup(user).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SignupHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_SignupHandlerFail(t *testing.T) {
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
		"id":         "0",
		"name":       "Misha",
		"telephone":  "909-277-47-21",
		"password":   "1234",
		"sex":        "male",
		"linkImages": null,
		"job":        "Fullstack",
		"education":  "BMSTU",
		"aboutMe":    ""
	}`)
	var ByteData = []byte(``)

	body2 := bytes.NewReader(ByteData)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/signup", body2)
	if err != nil {
		t.Fatal(err)
	}

	req2, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	requests := make([]*http.Request, 0, 1)
	requests = append(requests, req)
	requests = append(requests, req2)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Signup(user).Return(errors.New("error"))

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SignupHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, requests[i])
		status := rr.Code

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}

}

func TestUserHandler_SettingsHandlerSuccess(t *testing.T) {
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
		"name":       "Misha",
		"telephone":  "909-277-47-21",
		"password":   "1234",
		"sex":        "male",
		"linkImages": null,
		"job":        "Fullstack",
		"education":  "BMSTU",
		"aboutMe":    ""
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/settings", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Settings(sid, user).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SettingsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_SettingsHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/settings", body)
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
	handler := http.HandlerFunc(userHandler.SettingsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_SettingsHandlerFail(t *testing.T) {
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
		"name":       "Misha",
		"telephone":  "909-277-47-21",
		"password":   "1234",
		"sex":        "male",
		"linkImages": null,
		"job":        "Fullstack",
		"education":  "BMSTU",
		"aboutMe":    ""
	}`)
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/settings", body)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Settings(sid, user).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SettingsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)

}

func TestUserHandler_MeHandlerSuccess(t *testing.T) {
	user := models.UserFeed{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("GET", "/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Me(sid).Return(user, nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_MeHandlerFail(t *testing.T) {
	user := models.UserFeed{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("GET", "/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Me(sid).Return(user, domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_FeedHandlerSuccess(t *testing.T) {
	var users []models.UserFeed
	user1 := models.UserFeed{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	user2 := models.UserFeed{
		Name:       "Masha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	users = append(users, user1, user2)

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("GET", "/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Feed(sid).Return(users, nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.FeedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_FeedHandlerFail(t *testing.T) {
	var users []models.UserFeed
	user1 := models.UserFeed{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	user2 := models.UserFeed{
		Name:       "Masha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}
	users = append(users, user1, user2)

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("GET", "/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Feed(sid).Return(users, domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.FeedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_AddPhotoHandlerSuccess(t *testing.T) {
	photo := models.Photo{
		Path:      "./static/avatars/4.jpg",
		Telephone: "909-277-47-21",
	}

	var byteData = []byte(`{
		"linkImages": "./static/avatars/4.jpg",
		"telephone":  "909-277-47-21"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().AddPhoto(photo).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_AddPhotoHandlerFail(t *testing.T) {
	photo := models.Photo{
		Path:      "./static/avatars/4.jpg",
		Telephone: "909-277-47-21",
	}

	var byteData = []byte(`{
		"linkImages": "./static/avatars/4.jpg",
		"telephone":  "909-277-47-21"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().AddPhoto(photo).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_AddPhotoHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/add_photo", body)
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
	handler := http.HandlerFunc(userHandler.AddPhotoHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_UploadAvatarHandlerFail(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
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
	handler := http.HandlerFunc(userHandler.UploadAvatarHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Like(sid, like).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_LikeHandlerFail(t *testing.T) {
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Like(sid, like).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_LikeHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/like", body)
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
	handler := http.HandlerFunc(userHandler.LikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_DislikeHandlerSuccess(t *testing.T) {
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Dislike(sid, dislike).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.DislikeHandler)
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Dislike(sid, dislike).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.DislikeHandler)
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

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.DislikeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_CommentsByIdHandlerSuccess(t *testing.T) {
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().CommentsByID(12).Return(comments, nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CommentsByIdHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	err = json.NewDecoder(rr.Body).Decode(&outerComments)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, comments, outerComments)
	require.Equal(t, 200, status)
}

func TestUserHandler_CommentsByIdHandlerFail(t *testing.T) {
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().CommentsByID(12).Return(comments, errors.New("error"))

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CommentsByIdHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, requests[i])
		status := rr.Code

		require.NoError(t, err)
		require.Equal(t, 500, status)
	}
}

func TestUserHandler_CommentHandlerSuccess(t *testing.T) {
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Comment(sid, comment).Return(nil)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)

}

func TestUserHandler_CommentHandlerFail(t *testing.T) {
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

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Comment(sid, comment).Return(domain.ErrInternalServerError)

	userHandler := userHttp.UserHandler{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_CommentHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)
	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/comment", body)
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
	handler := http.HandlerFunc(userHandler.CommentHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
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
