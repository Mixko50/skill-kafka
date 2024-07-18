package skill

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type Skill struct {
	Key         string
	Name        string
	Description string
	Logo        string
	Tags        pq.StringArray
}

type skillStorage struct {
	db *sql.DB
}

func NewSkillStrage(db *sql.DB) skillStorage {
	return skillStorage{db: db}
}

func (s skillStorage) GetSkill(key string) (*Skill, error) {
	var skill Skill
	result := s.db.QueryRow("SELECT key,name,description,logo,tags from skill where key = $1", key)
	err := result.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)
	if err != nil {
		return nil, err
	}

	return &skill, nil
}

func (s skillStorage) GetSkills() ([]Skill, error) {
	skills := make([]Skill, 0)
	result, err := s.db.Query("SELECT key,name,description,logo,tags from skill")
	if err != nil {
		return make([]Skill, 0), err
	}
	for result.Next() {
		var skill Skill
		err := result.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)
		if err != nil {
			return make([]Skill, 0), errors.New("fail scaning")
		}
		skills = append(skills, skill)
	}

	return skills, nil
}
