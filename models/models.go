package models

import "time"

type LoginData struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type Photo struct {
	ID        int    `json:"-"`
	Path      string `json:"linkImages"`
	Telephone string `json:"telephone"`
}

type Comment struct {
	ID      int    `json:"-"`
	PhotoID int    `json:"photo_id"`
	Text    string `json:"text"`
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
	UserID       int       `json:"user_id"`
	Message      string    `json:"message"`
	TimeDelivery time.Time `json:"timeDelivery"`
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

type Error struct {
	Message string `json:"message"`
}
