package http_test

import (
	"bytes"
	"context"
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
	authClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	mockClient "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client/mock"
	faceClient "park_2020/2020_2_tmp_name/microservices/face_features/delivery/grpc/client"
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

func TestNewUserHandler(t *testing.T) {
	router := mux.NewRouter()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClient := &authClient.AuthClient{}
	mock := mock.NewMockUserUsecase(ctrl)
	faceClient := &faceClient.FaceClient{}
	userHttp.NewUserHandler(router, mock, authClient, faceClient)
}

func TestUserHandler_UploadAvatarHandlerFail(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
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
	handler := http.HandlerFunc(userHandler.UploadAvatarHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().Settings(uid, user).Return(nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().Settings(uid, user).Return(models.ErrInternalServerError)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SettingsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_SettingsHandlerFailUser(t *testing.T) {
	user := models.User{}

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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.SettingsHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestUserHandler_IsPremiumHandlerSuccess(t *testing.T) {
	user := models.User{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	uid := 0

	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("GET", "/is_premium", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().IsPremium(uid).Return(true)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.IsPremiumHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_IsPremiumHandlerFail(t *testing.T) {
	user := models.User{
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

	req, err := http.NewRequest("GET", "/is_premium", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.IsPremiumHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestUserHandler_MeHandlerSuccess(t *testing.T) {
	user := models.User{
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_MeHandlerFail(t *testing.T) {
	user := models.User{
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.MeHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().Feed(user).Return(users, nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, nil)
	mock.EXPECT().Feed(user).Return(users, models.ErrInternalServerError)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.FeedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 500, status)
}

func TestUserHandler_FeedHandlerFailUser(t *testing.T) {
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
	clientMock := mockClient.NewMockAuthClientInterface(ctrl)
	clientMock.EXPECT().CheckSession(context.Background(), req.Cookies()).Return(user, models.ErrUnauthorized)

	userHandler := userHttp.UserHandlerType{
		UUsecase:   mock,
		AuthClient: clientMock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.FeedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 401, status)
}

func TestUserHandler_UserIDHandlerSuccess(t *testing.T) {
	user := models.UserFeed{
		Name:       "Misha",
		LinkImages: nil,
		Job:        "Fullstack",
		Education:  "BMSTU",
		AboutMe:    "",
	}

	uid := 1

	req, err := http.NewRequest("GET", "/api/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().UserID(uid).Return(user, nil)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.UserIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_UserIDHandlerFailUser(t *testing.T) {
	user := models.UserFeed{}
	uid := 1
	req, err := http.NewRequest("GET", "/api/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().UserID(uid).Return(user, models.ErrNotFound)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.UserIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 404, status)
}

func TestUserHandler_UserIDHandlerFailAtoi(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1", nil)
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
	handler := http.HandlerFunc(userHandler.UserIDHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestUserHandler_TelephoneHandlerSuccess(t *testing.T) {
	phone := models.Phone{
		Telephone: "telephone",
	}

	var byteData = []byte(`{
		"telephone" : "telephone"
	}`)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/api/v1/telephone/telephone", body)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)
	mock.EXPECT().Telephone(phone.Telephone).Return(true)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.TelephoneHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 200, status)
}

func TestUserHandler_TelephoneHandlerFail(t *testing.T) {
	var byteData = []byte(``)

	body := bytes.NewReader(byteData)
	req, err := http.NewRequest("POST", "/api/v1/telephone/telephone", body)
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
	handler := http.HandlerFunc(userHandler.TelephoneHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}

func TestUserHandler_GetPremiumHandlerFail(t *testing.T) {
	sid := "something-like-this"
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(10 * time.Hour),
	}

	req, err := http.NewRequest("POST", "/get_premium", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockUserUsecase(ctrl)

	userHandler := userHttp.UserHandlerType{
		UUsecase: mock,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetPremiumHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code

	require.Equal(t, 400, status)
}
