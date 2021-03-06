package http

import (
	"context"
	"encoding/json"
	"net/http"
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	_authClientGRPC "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/client"
	"park_2020/2020_2_tmp_name/models"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AuthHandlerType struct {
	AUsecase   domain.AuthUsecase
	AuthClient _authClientGRPC.AuthClientInterface
}

func NewAuthHandler(r *mux.Router, us domain.AuthUsecase, client _authClientGRPC.AuthClientInterface) {
	handler := &AuthHandlerType{
		AUsecase:   us,
		AuthClient: client,
	}

	r.HandleFunc("/api/v1/login", handler.LoginHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/v1/logout", handler.LogoutHandler).Methods(http.MethodPost)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (a *AuthHandlerType) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginData := models.LoginData{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	if !loginData.IsLoggedIn {
		logrus.Error("have not firebase cookie")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sidString, err := a.AuthClient.Login(context.Background(), &loginData)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sidString,
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	body, err := json.Marshal(loginData)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (a *AuthHandlerType) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	err = a.AuthClient.Logout(context.Background(), session.Value)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal("logout success")
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
