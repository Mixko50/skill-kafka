package skill

type mockSkillService struct {
	SkillService
	err error
}

func (s mockSkillService) CreateSkill(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s mockSkillService) UpdateSkill(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s mockSkillService) UpdateName(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s mockSkillService) UpdateDescription(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s mockSkillService) UpdateLogo(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s mockSkillService) UpdateTags(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s mockSkillService) DeleteSkill(payload SkillQueuePayload) error {
	if s.err != nil {
		return s.err
	}
	return nil
}
