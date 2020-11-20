package http_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"park_2020/2020_2_tmp_name/api/users/mock"
	"park_2020/2020_2_tmp_name/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	userHttp "park_2020/2020_2_tmp_name/api/users/delivery/http"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	userHandler := &userHttp.UserHandlerType{}

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

	userHandler := userHttp.UserHandlerType{
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
	var byteData = []byte(`{
			"telephone" : "909-277-47-21",
			"password" : "qwerty"
		}`)

	body := bytes.NewReader(byteData)

	login := models.LoginData{
		Telephone: "909-277-47-21",
		Password:  "qwerty",
	}

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	gomock.InOrder(
		mock.EXPECT().Login(login).Return("some uuid", models.ErrInternalServerError),
	)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	require.Equal(t, 500, status)
}

func TestUserHandler_LoginHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)

	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LoginHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
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

	userHandler := userHttp.UserHandlerType{
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

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.LogoutHandler)
	for i := 0; i < 2; i++ {
		handler.ServeHTTP(rr, req)
		status := rr.Code
		req.AddCookie(cookie)

		require.NoError(t, err)
		require.Equal(t, 400, status)
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

	userHandler := userHttp.UserHandlerType{
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

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Signup(user).Return(errors.New("error"))

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SignupHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 500, status)

}

func TestUserHandler_SignupHandlerFailDecode(t *testing.T) {
	var byteData = []byte(``)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SignupHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.NoError(t, err)
	require.Equal(t, 400, status)

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

	uid := 0
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
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Settings(uid, user).Return(nil)

	userHandler := userHttp.UserHandlerType{
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

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SettingsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
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

	uid := 0
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
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Settings(uid, user).Return(models.ErrInternalServerError)

	userHandler := userHttp.UserHandlerType{
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

	userHandler := userHttp.UserHandlerType{
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
	mock.EXPECT().Me(sid).Return(user, models.ErrInternalServerError)

	userHandler := userHttp.UserHandlerType{
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
	user := models.User{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	user1 := models.UserFeed{
		Name:       "Dasha",
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
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Feed(user).Return(users, nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.FeedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_FeedHandlerFail(t *testing.T) {
	var user models.User
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
	mock.EXPECT().User(sid).Return(user, nil)
	mock.EXPECT().Feed(user).Return(users, models.ErrInternalServerError)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.FeedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}
