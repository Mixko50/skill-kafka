package skill

import (
	"database/sql"
	"github.com/lib/pq"
)

type Skill struct {
	Key         string
	Name        string
	Description string
	Logo        string
	Tags        []string
}

type skillStorage struct {
	db *sql.DB
}

func NewSkillStorage(db *sql.DB) skillStorage {
	return skillStorage{
		db: db,
	}
}

func (s skillStorage) CreateSkill(req CreateSkillRequest) error {
	qry := `INSERT INTO skill (key,name,description,logo,tags) VALUES($1,$2,$3,$4,$5);`
	_, err := s.db.Exec(qry, req.Key, req.Name, req.Description, req.Logo, pq.Array(req.Tags))
	if err != nil {
		return err
	}

	return nil
}
func (s skillStorage) UpdateSkill(id string, skill UpdateSkillRequest) error {
	qry := `UPDATE skill SET name = $1, description = $2, logo = $3, tags = $4 WHERE key = $5`
	_, err := s.db.Exec(qry, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags), id)
	if err != nil {
		return err
	}

	return nil
}
func (s skillStorage) UpdateName(key string, name string) error {
	qry := `UPDATE skill SET name = $1 WHERE key = $2`
	_, err := s.db.Exec(qry, name, key)
	if err != nil {
		return err
	}
	return nil
}
func (s skillStorage) UpdateDescription(key string, desc string) error {
	qry := `UPDATE skill SET description = $1 WHERE key = $2`
	_, err := s.db.Exec(qry, desc, key)
	if err != nil {
		return err
	}
	return nil
}
func (s skillStorage) UpdateLogo(key string, logo string) error {
	qry := `UPDATE skill SET logo = $1 WHERE key = $2`
	_, err := s.db.Exec(qry, logo, key)
	if err != nil {
		return err
	}
	return nil
}
func (s skillStorage) UpdateTags(key string, tag []string) error {
	qry := `UPDATE skill SET tags = $1 WHERE key = $2`
	_, err := s.db.Exec(qry, pq.Array(tag), key)
	if err != nil {
		return err
	}
	return nil
}

func (s skillStorage) DeleteSkill(key string) error {
	qry := `DELETE FROM skill WHERE key = $1`
	_, err := s.db.Exec(qry, key)
	if err != nil {
		return err
	}
	return nil
}
