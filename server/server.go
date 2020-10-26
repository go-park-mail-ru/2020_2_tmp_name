package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"park_2020/2020_2_tmp_name/models"

	"strings"
	"time"

	"github.com/google/uuid"
)

var uid int

type Service struct {
	sessions map[string]string
	users    map[string]*models.User
	DB       *sql.DB
}

func (s *Service) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewServer() *Service {
	return &Service{
		sessions: make(map[string]string, 10),
		users:    make(map[string]*models.User, 10),
	}
}

func JSONError(message string) []byte {
	jsonError, err := json.Marshal(models.Error{Message: message})
	if err != nil {
		return []byte("")
	}
	return jsonError
}

func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	loginData := models.LoginData{}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	var check bool

	if check = s.CheckUser(loginData.Telephone); !check {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("No user"))
		return
	}

	user, err := s.SelectUser(loginData.Telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	if user.Password != loginData.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("Wrong password"))
		return
	}

	SID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if check = s.CheckSession(loginData.Telephone); check {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User already authorized"))
		return
	}

	err = s.InsertSession(SID.String(), loginData.Telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	s.sessions[SID.String()] = user.Telephone

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID.String(),
		Expires: time.Now().Add(10 * time.Hour),
	}
	cookie.HttpOnly = false
	cookie.Secure = false

	body, err := json.Marshal(loginData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	http.SetCookie(w, cookie)
	w.Write(body)
}

func (s *Service) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("No session"))
		return
	}

	if _, ok := s.sessions[session.Value]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("No session"))
		return
	}

	err = s.DeleteSession(session.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("delete DB error"))
		return
	}

	body, err := json.Marshal("logout success")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	uid++
	var check bool

	if check = s.CheckUser(user.Telephone); check {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User alredy exists"))
		return
	}

	err = s.InsertUser(user, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	s.users[user.Telephone] = &user
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) Settings(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	err = s.UpdateUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Update user's params error"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Service) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var password string
	err := json.NewDecoder(r.Body).Decode(&password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}
}

func (s *Service) MeHandler(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	var cookie string
	cookie = fmt.Sprintf("%s", cookies[0])
	cookie = strings.TrimPrefix(cookie, "session_id=")

	var res string
	var ok bool
	if res, ok = s.sessions[cookie]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("No session"))
		return
	}

	body, err := json.Marshal(s.users[res])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) FeedHandler(w http.ResponseWriter, r *http.Request) {
	user := models.UserFeed{
		Name:       "andrey",
		LinkImages: []string{"/static/avatars/3.jpg", "/static/avatars/4.jpg", "/static/avatars/5.jpg"},
		Job:        "developer",
		Education:  "high",
		AboutMe:    "I am cool",
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) AddPhoto(w http.ResponseWriter, r *http.Request) {
	photo := models.Photo{}
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	if _, ok := s.users[photo.Telephone]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User alredy exists"))
		return
	}

	s.users[photo.Telephone].LinkImages = append(s.users[photo.Telephone].LinkImages, photo.LinkImage)

	w.WriteHeader(http.StatusOK)
}

func (s *Service) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)
	file, handler, err := r.FormFile("photo")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't parse form"))
		return
	}
	defer file.Close()

	str, err := os.Getwd()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't get directory"))
		return
	}

	r.FormValue("photo")
	os.Chdir("/home/ubuntu/go/src/2020_2_tmp_name/static/avatars")
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't open file"))
		return
	}
	defer f.Close()
	os.Chdir(str)

	body, err := json.Marshal("")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	io.Copy(f, file)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
