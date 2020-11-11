package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

// func TestLoginHandler(t *testing.T) {
// 	telephone := "944-739-32-28"
// 	loginData := &models.LoginData{
// 		Telephone: telephone,
// 		Password:  "password",
// 	}

// 	retModel := &models.LoginData{
// 		Telephone: telephone,
// 		Password:  "password",
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mock := mock.NewMockUserUsecase(ctrl)
// 	mock.EXPECT().Login(loginData).Times(1).Return(retModel, nil)

// 	uh := userHttp.UserHandler{
// 		UUsecase: mock,
// 	}

// 	models, err := uh.UUsecase.Login(*retModel)

// 	require.NoError(t, err)
// 	require.NotEqual(t, models, "")

// }

// func TestLogout(t *testing.T) {
// 	req, err := http.NewRequest("POST", "/logout", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(NewServer().Logout)
// 	handler.ServeHTTP(r, req)
// 	if status := r.Code; status != http.StatusUnauthorized {
// 		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// }

// func TestSignup(t *testing.T) {
// 	var byteData = []byte(`{
// 		"name" : "andrey",
// 		"telephone" : "12345",
// 		"day": "01",
// 		"month" : "01",
// 		"year" : "2001",
// 		"sex": "male",
// 		"job" : "frontend-developer",
// 		"education" : "high"
// 	}`)

// 	body := bytes.NewReader(byteData)

// 	req, err := http.NewRequest("POST", "/signup", body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(NewServer().Signup)
// 	handler.ServeHTTP(r, req)
// 	if status := r.Code; status != http.StatusOK {
// 		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	resultData := []byte(`{"name":"andrey","telephone":"12345","password":"","age":19,"day":"01","month":"01","year":"2001","sex":"male","linkImages":null,"job":"frontend-developer","education":"high","aboutMe":""}`)

// 	if !bytes.Equal(r.Body.Bytes(), resultData) {
// 		t.Errorf("wrong response: got %v want %v", string(r.Body.Bytes()), string(resultData))
// 	}
// }

// func TestSettings(t *testing.T) {
// 	var byteData = []byte(`{
// 		"name" : "andrey",
// 		"telephone" : "12345",
// 		"day": "01",
// 		"month" : "01",
// 		"year" : "2001",
// 		"sex": "male",
// 		"job" : "mobile-deleloper",
// 		"education" : "high"
// 	}`)

// 	body := bytes.NewReader(byteData)

// 	req, err := http.NewRequest("POST", "/settings", body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(NewServer().Settings)
// 	handler.ServeHTTP(r, req)
// 	if status := r.Code; status != http.StatusOK {
// 		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// }

// func TestFeed(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/feed", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(NewServer().Feed)
// 	handler.ServeHTTP(r, req)
// 	if status := r.Code; status != http.StatusOK {
// 		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// }
