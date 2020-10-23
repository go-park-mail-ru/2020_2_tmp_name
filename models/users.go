package models

type User struct {
	Name       string   `json:"name"`
	Telephone  string   `json:"telephone"`
	Password   string   `json:"password"`
	Age        int      `json:"age"`
	Day        string   `json:"day"`
	Month      string   `json:"month"`
	Year       string   `json:"year"`
	Sex        string   `json:"sex"`
	AccountID  int      `json:"account_id"`
	LinkImages []string `json:"linkImages"`
	Job        string   `json:"job"`
	Education  string   `json:"education"`
	AboutMe    string   `json:"aboutMe"`
}

type UserSafe struct {
	Name       string   `json:"name"`
	Telephone  string   `json:"telephone"`
	Age        int      `json:"age"`
	Day        string   `json:"day"`
	Month      string   `json:"month"`
	Year       string   `json:"year"`
	Sex        string   `json:"sex"`
	AccountID  int      `json:"account_id"`
	LinkImages []string `json:"linkImages"`
	Job        string   `json:"job"`
	Education  string   `json:"education"`
	AboutMe    string   `json:"aboutMe"`
}

type UserFeed struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	LinkImages []string `json:"linkImages"`
	Job        string   `json:"job"`
	Education  string   `json:"education"`
	AboutMe    string   `json:"aboutMe"`
}
