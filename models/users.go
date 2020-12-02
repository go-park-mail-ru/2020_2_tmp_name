package models

type User struct {
	ID         int      `json:"-"`
	Name       string   `json:"name"`
	Telephone  string   `json:"telephone"`
	Password   string   `json:"password"`
	DateBirth  int      `json:"date_birth"`
	Day        string   `json:"day"`
	Month      string   `json:"month"`
	Year       string   `json:"year"`
	Sex        string   `json:"sex"`
	LinkImages []string `json:"linkImages"`
	Job        string   `json:"job"`
	Education  string   `json:"education"`
	AboutMe    string   `json:"aboutMe"`
}

type UserMe struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Telephone  string   `json:"telephone"`
	DateBirth  int      `json:"date_birth"`
	LinkImages []string `json:"linkImages"`
	Job        string   `json:"job"`
	Education  string   `json:"education"`
	AboutMe    string   `json:"aboutMe"`
}

type UserFeed struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	DateBirth   int      `json:"date_birth"`
	LinkImages  []string `json:"linkImages"`
	Job         string   `json:"job"`
	Education   string   `json:"education"`
	AboutMe     string   `json:"aboutMe"`
	IsSuperlike bool     `json:"is_superlike"`
}

type Feed struct {
	Data []UserFeed `json:"user_feed"`
}
