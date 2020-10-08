package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var uid int

type service struct {
	sessions map[string]string
	users    map[string]*models.User
}

func (s *service) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewServer() *service {
	return &service{
		sessions: make(map[string]string, 10),
		users:    make(map[string]*models.User, 10),
	}
}

func (s *service) Login(w http.ResponseWriter, r *http.Request) {
	loginData := models.LoginData{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, ok := s.users[loginData.Telephone]
	if !ok {
		http.Error(w, "no user", 404)
		return
	}

	if user.Password != loginData.Password {
		http.Error(w, "wrong password", 400)
		return
	}

	body, err := json.Marshal(loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	SID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	s.sessions[SID.String()] = user.Telephone

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID.String(),
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
	w.Write(body)
}

func (s *service) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println(err)
		http.Error(w, "no session", 401)
		return
	}

	if _, ok := s.sessions[session.Value]; !ok {
		http.Error(w, "no session", 401)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (s *service) Signup(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uid++
	user.AccountID = uid

	if _, ok := s.users[user.Telephone]; ok {
		http.Error(w, "user alredy exists", 401)
		return
	}

	year, err := strconv.Atoi(user.Year)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Age = time.Now().Year() - year

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.users[user.Telephone] = &user
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *service) Settings(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *service) MeHandler(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	var cookie string
	cookie = fmt.Sprintf("%s", cookies[0])
	cookie = strings.TrimPrefix(cookie, "session_id=")

	var res string
	var ok bool
	if res, ok = s.sessions[cookie]; !ok {
		http.Error(w, "no session", 401)
		return
	}

	body, err := json.Marshal(s.users[res])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *service) FeedHandler(w http.ResponseWriter, r *http.Request) {
	user := models.UserFeed{
		Name:       "andrey",
		Age:        39,
		LinkImages: []string{"/static/avatars/3.jpg", "/static/avatars/4.jpg", "/static/avatars/5.jpg"},
		Job:        "developer",
		Education:  "high",
		AboutMe:    "I am cool",
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *service) AddPhoto(w http.ResponseWriter, r *http.Request) {
	photo := models.Photo{}
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.users[photo.Name]; !ok {
		http.Error(w, "user doesn' t exists", 400)
		return
	}

	s.users[photo.Name].LinkImages = append(s.users[photo.Name].LinkImages, photo.LinkImage)

	w.WriteHeader(http.StatusOK)
}

func (s *service) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	r.FormValue("file_name")
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	w.WriteHeader(http.StatusOK)
}
