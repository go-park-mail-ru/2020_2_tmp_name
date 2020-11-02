package server

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"park_2020/2020_2_tmp_name/models"

	"time"

	"github.com/google/uuid"
)

var uid int

type Service struct {
	DB *sql.DB
}

func (s *Service) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewServer() *Service {
	return &Service{}
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

	if CheckPasswordHash(user.Password, loginData.Password) {
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
	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUser(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
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

	user, err := s.SelectUser(photo.Telephone)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't select user"))
		return
	}

	user.LinkImages = append(user.LinkImages, photo.LinkImage)

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

func (s *Service) Like(w http.ResponseWriter, r *http.Request) {
	like := models.Like{}
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	err = s.InsertLike(like)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}
}

func (s *Service) Dislike(w http.ResponseWriter, r *http.Request) {
	dislike := models.Dislike{}
	err := json.NewDecoder(r.Body).Decode(&dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	err = s.InsertDislike(dislike)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}
}

func (s *Service) Comment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	err = s.InsertComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}
}
