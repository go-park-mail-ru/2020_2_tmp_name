package postgres

import (
	"database/sql"
	domain "park_2020/2020_2_tmp_name/api/photos"
	"park_2020/2020_2_tmp_name/models"
)

type postgresPhotoRepository struct {
	Conn *sql.DB
}

func NewPostgresPhotoRepository(Conn *sql.DB) domain.PhotoRepository {
	return &postgresPhotoRepository{Conn}
}

func (p *postgresPhotoRepository) SelectUserFeed(telephone string) (models.UserFeed, error) {
	var u models.UserFeed
	var tid int
	row := p.Conn.QueryRow(`SELECT id, name, date_birth, education, job, about_me, filter_id FROM users
						WHERE  telephone=$1;`, telephone)
	err := row.Scan(&u.ID, &u.Name, &u.DateBirth, &u.Education, &u.Job, &u.AboutMe, &tid)
	if err != nil {
		return u, err
	}

	u.LinkImages, err = p.SelectImages(u.ID)
	u.Target = models.IDToTarget(tid)
	return u, err
}

func (p *postgresPhotoRepository) SelectImages(uid int) ([]string, error) {
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

func (p *postgresPhotoRepository) InsertPhoto(path string, uid int) error {
	_, err := p.Conn.Exec(`INSERT INTO photo(path, user_id) VALUES ($1, $2);`, path, uid)
	return err
}

func (p *postgresPhotoRepository) DeletePhoto(path string, uid int) error {
	_, err := p.Conn.Exec(`DELETE FROM photo WHERE path=$1 AND user_id=$2;`, path, uid)
	return err
}

func (p *postgresPhotoRepository) SelectPhotoWithMask(path string) ([]string, error) {
	rows, err := p.Conn.Query(`SELECT path FROM photo WHERE path LIKE '$1_%';`, path)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []string
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
