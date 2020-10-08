package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(NewServer().HealthHandler)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLogin(t *testing.T) {
	var byteData = []byte(`{
			"login" : "amavrin",
			"password" : "qwerty"
		}`)

	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/login", body)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(NewServer().Login)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusNotFound {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLogout(t *testing.T) {
	req, err := http.NewRequest("POST", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(NewServer().Logout)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusUnauthorized {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestSignup(t *testing.T) {
	var byteData = []byte(`{
		"name" : "andrey",
		"telephone" : "12345",
		"day": "01",
		"month" : "01",
		"year" : "2001",
		"sex": "male",
		"job" : "frontend-developer",
		"education" : "high"
	}`)

	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/signup", body)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(NewServer().Signup)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	resultData := []byte(`{"name":"andrey","telephone":"12345","password":"","age":19,"day":"01","month":"01","year":"2001","sex":"male","account_id":1,"linkImages":null,"job":"frontend-developer","education":"high","aboutMe":""}`)

	if !bytes.Equal(r.Body.Bytes(), resultData) {
		t.Errorf("wrong response: got %v want %v", string(r.Body.Bytes()), string(resultData))
	}
}

func TestSettings(t *testing.T) {
	var byteData = []byte(`{
		"name" : "andrey",
		"telephone" : "12345",
		"day": "01",
		"month" : "01",
		"year" : "2001",
		"sex": "male",
		"job" : "mobile-deleloper",
		"education" : "high"
	}`)

	body := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST", "/settings", body)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(NewServer().Settings)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestFeedHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/feed", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(NewServer().FeedHandler)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}
}
