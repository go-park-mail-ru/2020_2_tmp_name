package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/api/users"
	"park_2020/2020_2_tmp_name/models"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

func NewPostgresUserRepository(Conn *sql.DB) domain.UserRepository {
	return &postgresUserRepository{Conn}
}

func (p *postgresUserRepository) CheckUser(telephone string) bool {
	var count int
	p.Conn.QueryRow(`SELECT COUNT(telephone) FROM users WHERE telephone=$1;`, telephone).Scan(&count)
	return count > 0
}

func (p *postgresUserRepository) InsertUser(user models.User) error {
	password, err := models.HashPassword(user.Password)
	if err != nil {
		return err
	}

	age, err := models.Age(user.Day, user.Month, user.Year)
	if err != nil {
		return err
	}

	_, err = p.Conn.Exec(`INSERT INTO users(name, telephone, password, date_birth, sex, job, education, about_me)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`,
		user.Name,
		user.Telephone,
		password,
		age,
		user.Sex,
		user.Job,
		user.Education,
		user.AboutMe,
	)
	if err != nil {
		return err
	}

	var uid int
	row := p.Conn.QueryRow(`SELECT id FROM users WHERE  telephone=$1;`, user.Telephone)
	err = row.Scan(&uid)
	if err != nil {
		return err
	}

	for _, path := range user.LinkImages {
		_, err = p.Conn.Exec(`INSERT INTO photo(path, user_id) VALUES ($1, $2);`, path, uid)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *postgresUserRepository) SelectUser(telephone string) (models.User, error) {
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

func (p *postgresUserRepository) SelectUserMe(telephone string) (models.UserMe, error) {
	var u models.UserMe
	row := p.Conn.QueryRow(`SELECT id, name, telephone, date_birth, job, education, about_me FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserFeed(telephone string) (models.UserFeed, error) {
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

func (p *postgresUserRepository) SelectUserByID(uid int) (models.User, error) {
	var u models.User
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me FROM users
						WHERE  id=$1;`, uid)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	return u, err
}

func (p *postgresUserRepository) SelectUserFeedByID(uid int) (models.UserFeed, error) {
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

func (p *postgresUserRepository) Match(uid1, uid2 int) bool {
	var id1, id2 int
	row := p.Conn.QueryRow(`Select user_id1, user_id2 FROM likes 
							WHERE user_id1 = $1 AND user_id2 = $2;`, uid2, uid1)
	err := row.Scan(&id1, &id2)
	return err == nil
}

func (p *postgresUserRepository) SelectUsers(user models.User) ([]models.UserFeed, error) {
	var users []models.UserFeed
	rows, err := p.Conn.Query(`SELECT id, name, date_birth, education, job,  about_me FROM users WHERE sex != $1`, user.Sex)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.UserFeed
		err := rows.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe)
		if err != nil {
			continue
		}

		u.LinkImages, err = p.SelectImages(u.ID)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (p *postgresUserRepository) UpdateUser(user models.User, uid int) error {
	if user.Name != "" {
		_, err := p.Conn.Exec(`UPDATE users SET name=$1 WHERE id = $2;`, user.Name, uid)
		if err != nil {
			return err
		}
	}
	if user.Telephone != "" {
		_, err := p.Conn.Exec(`UPDATE users SET telephone=$1 WHERE id = $2;`, user.Telephone, uid)
		if err != nil {
			return err
		}
	}
	if user.Password != "" {
		password, err := models.HashPassword(user.Password)
		if err != nil {
			return err
		}
		_, err = p.Conn.Exec(`UPDATE users SET password=$1 WHERE id = $2;`, password, uid)
		if err != nil {
			return err
		}
	}
	if user.Job != "" {
		_, err := p.Conn.Exec(`UPDATE users SET job=$1 WHERE id = $2;`, user.Job, uid)
		if err != nil {
			return err
		}
	}
	if user.Education != "" {
		_, err := p.Conn.Exec(`UPDATE users SET education=$1 WHERE id = $2;`, user.Education, uid)
		if err != nil {
			return err
		}
	}
	if user.AboutMe != "" {
		_, err := p.Conn.Exec(`UPDATE users SET about_me=$1 WHERE id = $2;`, user.AboutMe, uid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *postgresUserRepository) InsertSession(sid, telephone string) error {
	_, err := p.Conn.Exec(`INSERT INTO sessions(key, value) VALUES ($1, $2);`, sid, telephone)
	return err
}

func (p *postgresUserRepository) DeleteSession(sid string) error {
	_, err := p.Conn.Exec(`DELETE FROM sessions WHERE key=$1;`, sid)
	return err
}

func (p *postgresUserRepository) CheckUserBySession(sid string) string {
	var count string
	p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&count)
	return count
}

func (p *postgresUserRepository) SelectImages(uid int) ([]string, error) {
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
