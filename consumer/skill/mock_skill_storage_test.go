package skill

type mockSkillStorage struct {
	SkillStorage
	err error
}

func (m *mockSkillStorage) CreateSkill(req CreateSkillRequest) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockSkillStorage) UpdateSkill(id string, skill UpdateSkillRequest) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockSkillStorage) UpdateName(key string, name string) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockSkillStorage) UpdateDescription(key string, desc string) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockSkillStorage) UpdateLogo(key string, logo string) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockSkillStorage) UpdateTags(key string, tag []string) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m *mockSkillStorage) DeleteSkill(key string) error {
	if m.err != nil {
		return m.err
	}
	return nil
}
