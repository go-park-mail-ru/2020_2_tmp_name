package models

import "time"

type User struct {
	Name       string    `json:"name"`
	Telephone  string    `json:"telephone"`
	Password   string    `json:"password"`
	DateBirth  time.Time `json:"date_birth"`
	Sex        string    `json:"sex"`
	AccountID  int       `json:"account_id"`
	LinkImages []string  `json:"linkImages"`
	Job        string    `json:"job"`
	Education  string    `json:"education"`
	AboutMe    string    `json:"aboutMe"`
}

type UserSafe struct {
	Name       string    `json:"name"`
	Telephone  string    `json:"telephone"`
	DateBirth  time.Time `json:"date_birth"`
	Sex        string    `json:"sex"`
	AccountID  int       `json:"account_id"`
	LinkImages []string  `json:"linkImages"`
	Job        string    `json:"job"`
	Education  string    `json:"education"`
	AboutMe    string    `json:"aboutMe"`
}

type UserFeed struct {
	Name       string    `json:"name"`
	DateBirth  time.Time `json:"date_birth"`
	LinkImages []string  `json:"linkImages"`
	Job        string    `json:"job"`
	Education  string    `json:"education"`
	AboutMe    string    `json:"aboutMe"`
}
