package models

import "time"

type LoginData struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type Photo struct {
	ID   int    `json:"-"`
	Path string `json:"path"`
	UID  int    `json:"user_id"`
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
