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
	ID        int    `json:"id"`
	Telephone string `json:"telephone"`
	LinkImage string `json:"link_image"`
}

type Comment struct {
	ID      int    `json:"id"`
	PhotoID int    `json:"photo_id"`
	Text    string `json:"text"`
}

type Like struct {
	ID   int `json:"id"`
	Uid1 int `json:"user_id1"`
	Uid2 int `json:"user_id2"`
}

type Dislike struct {
	ID   int `json:"id"`
	Uid1 int `json:"user_id1"`
	Uid2 int `json:"user_id2"`
}

type Error struct {
	Message string `json:"message"`
}
