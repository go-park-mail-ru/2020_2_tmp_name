package http

import (
	"encoding/json"
	"log"
	"net/http"
	domain "park_2020/2020_2_tmp_name/api/users"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"

	"time"

	"github.com/gorilla/mux"
)

type UserHandlerType struct {
	UUsecase domain.UserUsecase
}

func NewUserHandler(r *mux.Router, us domain.UserUsecase) {
	handler := &UserHandlerType{
		UUsecase: us,
	}

	r.HandleFunc("/health", handler.HealthHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/login", handler.LoginHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/v1/logout", handler.LogoutHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/signup", handler.SignupHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/settings", handler.SettingsHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/me", handler.MeHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/feed", handler.FeedHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/{user_id}", handler.UserIDHandler).Methods(http.MethodGet)
}

func (u *UserHandlerType) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (u *UserHandlerType) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginData := models.LoginData{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	sidString, err := u.UUsecase.Login(loginData)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Logout(session.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal("logout success")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	err = u.UUsecase.Signup(user)
	if err != nil {
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) SettingsHandler(w http.ResponseWriter, r *http.Request) {
	userData := models.User{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError(err.Error()))
		return
	}

	cookie := r.Cookies()[0]
	err = u.UUsecase.Settings(cookie.Value, userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(userData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) MeHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]
	user, err := u.UUsecase.Me(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) FeedHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	cookie := r.Cookies()[0]
	var feed models.Feed
	feed.Data, err = u.UUsecase.Feed(cookie.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(feed)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (u *UserHandlerType) UserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/user/"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	user, err := u.UUsecase.UserID(userID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(models.GetStatusCode(err))
		w.Write(JSONError(err.Error()))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}