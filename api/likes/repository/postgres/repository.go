package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/api/likes"
	"park_2020/2020_2_tmp_name/models"
)

type postgresLikeRepository struct {
	Conn *sql.DB
}

func NewPostgresLikeRepository(Conn *sql.DB) domain.LikeRepository {
	return &postgresLikeRepository{Conn}
}

func (p *postgresLikeRepository) SelectUserFeed(telephone string) (models.UserFeed, error) {
	var u models.UserFeed
	row := p.Conn.QueryRow(`SELECT id, name, date_birth, education, job, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresLikeRepository) CheckUserBySession(sid string) string {
	var count string
	p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&count)
	return count
}

func (p *postgresLikeRepository) SelectUser(telephone string) (models.User, error) {
	var u models.User
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresLikeRepository) SelectChatID(uid1, uid2 int) (int, error) {
	var chid int
	row := p.Conn.QueryRow(`SELECT id FROM chat WHERE user_id1=$1 AND user_id2=$2;`, uid1, uid2)
	err := row.Scan(&chid)
	return chid, err
}

func (p *postgresLikeRepository) CheckChat(chat models.Chat) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`SELECT user_id1, user_id2 FROM chat 
							WHERE user_id1 = $1 AND user_id2 = $2 
							OR user_id1 = $2 AND user_id2 = $1`, chat.Uid1, chat.Uid2)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresLikeRepository) InsertChat(chat models.Chat) error {
	_, err := p.Conn.Exec(`INSERT INTO chat(user_id1, user_id2) VALUES ($1, $2);`, chat.Uid1, chat.Uid2)
	return err
}

func (p *postgresLikeRepository) Match(uid1, uid2 int) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`Select user_id1, user_id2 FROM likes 
							WHERE user_id1 = $1 AND user_id2 = $2;`, uid2, uid1)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresLikeRepository) SelectUserByChat(uid, chid int) (models.UserFeed, error) {
	var user models.UserFeed
	var id1, id2, id int

	row := p.Conn.QueryRow(`SELECT user_id1, user_id2 FROM chat WHERE id=$1;`, chid)

	err := row.Scan(&id1, &id2)
	if err != nil {
		return user, err
	}

	if id1 != uid {
		id = id1
	} else {
		id = id2
	}

	user, err = p.SelectUserFeedByID(id)
	return user, err
}

func (p *postgresLikeRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
	var u models.UserFeed
	row := p.Conn.QueryRow(`SELECT name, date_birth, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.Name, &u.DateBirth, &u.Job, &u.Education, &u.AboutMe)
	if err != nil {
		return u, err
	}
	u.ID = uid

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresLikeRepository) InsertLike(uid1, uid2 int) error {
	_, err := p.Conn.Exec(`INSERT INTO likes(user_id1, user_id2) VALUES ($1, $2);`, uid1, uid2)
	return err
}

func (p *postgresLikeRepository) InsertDislike(uid1, uid2 int) error {
	_, err := p.Conn.Exec(`INSERT INTO dislikes(user_id1, user_id2) VALUES ($1, $2);`, uid1, uid2)
	return err
}

func (p *postgresLikeRepository) SelectImages(uid int) ([]string, error) {
	var images []string
	rows, err := p.Conn.Query(`SELECT path FROM photo WHERE  user_id=$1;`, uid)
	if err != nil {
		return images, err
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		err := rows.Scan(&image)
		if err != nil {
			continue
		}
		images = append(images, image)
	}
	return images, nil
}
