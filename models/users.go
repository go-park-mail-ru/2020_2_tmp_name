package models

import "time"

type User struct {
	ID         int       `json:"-"`
	Name       string    `json:"name"`
	Telephone  string    `json:"telephone"`
	Password   string    `json:"password"`
	DateBirth  time.Time `json:"date_birth"`
	Sex        string    `json:"sex"`
	LinkImages []string  `json:"linkImages"`
	Job        string    `json:"job"`
	Education  string    `json:"education"`
	AboutMe    string    `json:"aboutMe"`
}

type UserSafe struct {
	ID         int       `json:"-"`
	Name       string    `json:"name"`
	Telephone  string    `json:"telephone"`
	DateBirth  time.Time `json:"date_birth"`
	Sex        string    `json:"sex"`
	LinkImages []string  `json:"linkImages"`
	Job        string    `json:"job"`
	Education  string    `json:"education"`
	AboutMe    string    `json:"aboutMe"`
}

type UserFeed struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	DateBirth  time.Time `json:"date_birth"`
	LinkImages []string  `json:"linkImages"`
	Job        string    `json:"job"`
	Education  string    `json:"education"`
	AboutMe    string    `json:"aboutMe"`
}

type Feed struct {
	Data []UserFeed `json:"user_feed"`
}
