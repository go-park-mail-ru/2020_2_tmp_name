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
			"email" : "amavrin@mail.ru",
			"login" : "amavrin",
			"password" : "qwerty",
			"cellphone" : "12345"
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
}

func TestSettings(t *testing.T) {
	var byteData = []byte(`{
			"email" : "amavrin@mail.ru",
			"login" : "amavrin",
			"password" : "qwerty",
			"cellphone" : "12345"
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
