package skill

type mockSkillStorage struct {
	SkillStorage
	skill                 *Skill
	skills                []Skill
	errGet                error
	errUpdateCreateDelete error
}

func (m *mockSkillStorage) GetSkill(key string) (*Skill, error) {
	if m.errGet != nil {
		return nil, m.errGet
	}
	return m.skill, nil
}

func (m *mockSkillStorage) GetSkills() ([]Skill, error) {
	skills := make([]Skill, 0)
	if m.errGet != nil {
		return make([]Skill, 0), m.errGet
	}

	for _, skill := range m.skills {
		skills = append(skills, skill)
	}
	return skills, nil
}
