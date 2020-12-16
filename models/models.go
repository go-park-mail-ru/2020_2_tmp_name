package models

import "time"

type LoginData struct {
	Telephone  string `json:"telephone"`
	Password   string `json:"password"`
	IsLoggedIn bool   `json:"is_logged_in"`
}

type HasTelephone struct {
	Telephone bool `json:"telephone"`
}

type Photo struct {
	ID        int    `json:"-"`
	Path      string `json:"linkImages"`
	Telephone string `json:"telephone"`
}

type Comment struct {
	ID           int    `json:"-"`
	Uid1         int    `json:"user_id1"`
	Uid2         int    `json:"user_id2"`
	TimeDelivery string `json:"timeDelivery"`
	CommentText  string `json:"commentText"`
}

type CommentId struct {
	UserId       int    `json:"-"`
	CommentText  string `json:"commentText"`
	TimeDelivery string `json:"timeDelivery"`
}

type CommentsById struct {
	Comments []CommentById `json:"comments"`
}

type CommentsData struct {
	Data CommentsById `json:"data"`
}

type CommentById struct {
	User         UserFeed `json:"user"`
	CommentText  string   `json:"commentText"`
	TimeDelivery string   `json:"timeDelivery"`
}

type Like struct {
	ID   int `json:"-"`
	Uid1 int `json:"-"`
	Uid2 int `json:"user_id2"`
}

type Dislike struct {
	ID   int `json:"-"`
	Uid1 int `json:"-"`
	Uid2 int `json:"user_id2"`
}

type Chat struct {
	ID      int    `json:"-"`
	Uid1    int    `json:"user_id1"`
	Uid2    int    `json:"user_id2"`
	LastMsg string `json:"last_msg"`
}

type ChatData struct {
	ID       int      `json:"id"`
	Partner  UserFeed `json:"partner"`
	Messages []Msg    `json:"messages"`
}

type Msg struct {
	UserID       int    `json:"user_id"`
	ChatID       int    `json:"chat_id"`
	Message      string `json:"message"`
	TimeDelivery string `json:"timeDelivery"`
}

type ChatModel struct {
	Data []ChatData `json:"data"`
}

type Message struct {
	ID           int       `json:"-"`
	Text         string    `json:"text"`
	TimeDelivery time.Time `json:"-"`
	ChatID       int       `json:"chat_id"`
	UserID       int       `json:"-"`
}

type Phone struct {
	Telephone string `json:"telephone"`
}

type Image struct {
	LinkImage string `json:"link_image"`
}

type Premium struct {
	IsPremium bool `json:"is_premium"`
}

type Error struct {
	Message string `json:"message"`
}

type Superlike struct {
	ID   int `json:"-"`
	Uid1 int `json:"-"`
	Uid2 int `json:"user_id2"`
}
