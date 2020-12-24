package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/microservices/authorization"
	"park_2020/2020_2_tmp_name/models"
)

type postgresAuthRepository struct {
	Conn *sql.DB
}

func NewPostgresAuthRepository(Conn *sql.DB) domain.AuthRepository {
	return &postgresAuthRepository{Conn}
}

func (p *postgresAuthRepository) CheckUser(telephone string) bool {
	var count int
	err := p.Conn.QueryRow(`SELECT COUNT(telephone) FROM users WHERE telephone=$1;`, telephone).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (p *postgresAuthRepository) SelectUser(telephone string) (models.User, error) {
	var u models.User
	var tid int
	row := p.Conn.QueryRow(`SELECT id, name, telephone, password, date_birth, sex, job, education, about_me, filter_id FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.Telephone, &u.Password, &u.DateBirth, &u.Sex, &u.Education, &u.Job, &u.AboutMe, &tid)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	u.Target = models.IDToTarget(tid)
	return u, err
}

func (p *postgresAuthRepository) InsertSession(sid, telephone string) error {
	_, err := p.Conn.Exec(`INSERT INTO sessions(key, value) VALUES ($1, $2);`, sid, telephone)
	return err
}

func (p *postgresAuthRepository) DeleteSession(sid string) error {
	_, err := p.Conn.Exec(`DELETE FROM sessions WHERE key=$1;`, sid)
	return err
}

func (p *postgresAuthRepository) CheckUserBySession(sid string) string {
	var str string
	err := p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid).Scan(&str)
	if err != nil {
		return ""
	}
	return str
}

func (p *postgresAuthRepository) SelectUserBySession(sid string) (string, error) {
	var telephone string
	row := p.Conn.QueryRow(`SELECT value FROM sessions WHERE key=$1;`, sid)
	err := row.Scan(&telephone)
	return telephone, err
}

func (p *postgresAuthRepository) SelectImages(uid int) ([]string, error) {
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
			return images, err
		}
		images = append(images, image)
	}
	return images, nil
}
