package models

type Account struct {
	AccountID int    `json:"account_id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
}

type LoginData struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type Photo struct {
	Telephone string `json:"telephone"`
	LinkImage string `json:"link_image"`
}

type Error struct {
	Message string `json:"message"`
}
