package server

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"park_2020/2020_2_tmp_name/chat"
	"park_2020/2020_2_tmp_name/models"
	"strconv"
	"strings"

	"time"

	"github.com/google/uuid"
)

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

	if !CheckPasswordHash(loginData.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("Wrong password"))
		return
	}

	SID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// if check = s.CheckSession(loginData.Telephone); check {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write(JSONError("User already authorized"))
	// 	return
	// }

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

	var check bool

	if check = s.CheckUser(user.Telephone); check {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(JSONError("User alredy exists"))
		return
	}

	err = s.InsertUser(user)
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

func (s *Service) MeHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUserFeed(telephone)
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

func (s *Service) Feed(w http.ResponseWriter, r *http.Request) {
	var err error
	var user models.User
	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err = s.SelectUser(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}
	var feed models.Feed
	feed.Data, err = s.SelectUsers(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can`t select users"))
		return
	}

	body, err := json.Marshal(feed)
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

	user, err := s.SelectUserFeed(photo.Telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	user.LinkImages = append(user.LinkImages, photo.Path)

	err = s.InsertPhoto(photo.Path, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	body, err := json.Marshal(photo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024)
	file, _, err := r.FormFile("photo")
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
	photoID, err := uuid.NewRandom()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	f, err := os.OpenFile(photoID.String(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't open file"))
		return
	}
	defer f.Close()
	os.Chdir(str)

	body, err := json.Marshal(photoID.String())
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

	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUserFeed(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	err = s.InsertLike(user.ID, like.Uid2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	body, err := json.Marshal(like)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
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

	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUserFeed(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	err = s.InsertDislike(user.ID, dislike.Uid2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	body, err := json.Marshal(dislike)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
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

	body, err := json.Marshal(comment)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) Chat(w http.ResponseWriter, r *http.Request) {
	chat := models.Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	err = s.InsertChat(chat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) Message(w http.ResponseWriter, r *http.Request) {
	message := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONError("Can't decode data"))
		return
	}

	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUserFeed(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	err = s.InsertMessage(message.Text, message.ChatID, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("insert DB error"))
		return
	}

	body, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) Chats(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUserFeed(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	chats, err := s.SelectChatsByID(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("select DB error"))
		return
	}

	var chatModel models.ChatModel
	chatModel.Data = chats

	body, err := json.Marshal(chatModel)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) ChatID(w http.ResponseWriter, r *http.Request) {
	chid, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/v1/chats/"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("data transform error"))
		return
	}

	cookie := r.Cookies()[0]
	telephone := s.CheckUserBySession(cookie.Value)
	user, err := s.SelectUserFeed(telephone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Can't select user"))
		return
	}

	chat, err := s.SelectChatByID(user.ID, chid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("select DB error"))
		return
	}

	body, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONError("Marshal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Service) Gochat(w http.ResponseWriter, r *http.Request) {
	server := chat.NewServer("/entry")
	go server.Listen()
	w.WriteHeader(http.StatusOK)
}
