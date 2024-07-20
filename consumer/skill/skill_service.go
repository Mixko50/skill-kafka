package skill

type SkillStorage interface {
	CreateSkill(req CreateSkillRequest) error
	UpdateSkill(id string, skill UpdateSkillRequest) error
	UpdateName(key string, name string) error
	UpdateDescription(key string, desc string) error
	UpdateLogo(key string, logo string) error
	UpdateTags(key string, tag []string) error
	DeleteSkill(key string) error
}

type skillService struct {
	skillStorage SkillStorage
}

func NewSkillService(skillStorage SkillStorage) skillService {
	return skillService{
		skillStorage: skillStorage,
	}
}

func (s skillService) CreateSkill(payload SkillQueuePayload) error {
	data, err := ConvertSkillType[CreateSkillRequest](payload.Payload)
	if err != nil {
		return err
	}

	err = s.skillStorage.CreateSkill(*data)
	if err != nil {
		return err
	}

	return nil
}

func (s skillService) UpdateSkill(payload SkillQueuePayload) error {
	data, err := ConvertSkillType[UpdateSkillRequest](payload.Payload)
	if err != nil {
		return err
	}

	err = s.skillStorage.UpdateSkill(*payload.Key, *data)
	if err != nil {
		return err
	}

	return nil
}

func (s skillService) UpdateName(payload SkillQueuePayload) error {
	data, err := ConvertSkillType[UpdateSkillNameRequest](payload.Payload)
	if err != nil {
		return err
	}

	err = s.skillStorage.UpdateName(*payload.Key, data.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s skillService) UpdateDescription(payload SkillQueuePayload) error {
	data, err := ConvertSkillType[UpdateSkillDescriptionRequest](payload.Payload)
	if err != nil {
		return err
	}

	err = s.skillStorage.UpdateDescription(*payload.Key, data.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s skillService) UpdateLogo(payload SkillQueuePayload) error {
	data, err := ConvertSkillType[UpdateSkillLogoRequest](payload.Payload)
	if err != nil {
		return err
	}

	err = s.skillStorage.UpdateLogo(*payload.Key, data.Logo)
	if err != nil {
		return err
	}

	return nil
}

func (s skillService) UpdateTags(payload SkillQueuePayload) error {
	data, err := ConvertSkillType[UpdateSkillTagsRequest](payload.Payload)
	if err != nil {
		return err
	}

	err = s.skillStorage.UpdateTags(*payload.Key, data.Tags)
	if err != nil {
		return err
	}

	return nil
}

func (s skillService) DeleteSkill(payload SkillQueuePayload) error {
	err := s.skillStorage.DeleteSkill(*payload.Key)
	if err != nil {
		return err
	}
	return nil
}
