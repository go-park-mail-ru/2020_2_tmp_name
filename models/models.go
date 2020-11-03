package models

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
	Uid1 int `json:"user_id1"`
	Uid2 int `json:"user_id2"`
}

type Dislike struct {
	ID   int `json:"-"`
	Uid1 int `json:"user_id1"`
	Uid2 int `json:"user_id2"`
}

type Chat struct {
	ID   int `json:"-"`
	Uid1 int `json:"user_id1"`
	Uid2 int `json:"user_id2"`
}

type Error struct {
	Message string `json:"message"`
}
