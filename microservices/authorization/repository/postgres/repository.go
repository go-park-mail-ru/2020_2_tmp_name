package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	"park_2020/2020_2_tmp_name/models"
	"time"
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

func (p *postgresUserRepository) InsertSession(sid, telephone string) error {
	_, err := p.Conn.Exec(`INSERT INTO sessions(key, value) VALUES ($1, $2);`, sid, telephone)
	return err
}

func (p *postgresUserRepository) DeleteSession(sid string) error {
	_, err := p.Conn.Exec(`DELETE FROM sessions WHERE key=$1;`, sid)
	return err
}

func (p *postgresUserRepository) CheckUserBySession(sid string) bool {
	var count int
	p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&count)
	return count > 0
}

func (p *postgresUserRepository) SelectUserBySession(sid string) (string, error) {
	var telephone string
	row := p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid)
	err := row.Scan(&telephone)
	return telephone, err
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

func (p *postgresUserRepository) InsertPremium(uid int, dateFrom time.Time, dateTo time.Time) error {
	_, err := p.Conn.Exec(`INSERT INTO premium_accounts(user_id, date_to, date_from) 
								 VALUES ($1, $2, $3);`, uid, dateTo, dateFrom)

	return err
}
