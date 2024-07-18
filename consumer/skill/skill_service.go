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

func (s skillService) CreateSkill(req CreateSkillRequest) error {
	return s.skillStorage.CreateSkill(req)
}

func (s skillService) UpdateSkill(id string, skill UpdateSkillRequest) error {
	return s.skillStorage.UpdateSkill(id, skill)
}

func (s skillService) UpdateName(key string, skill UpdateSkillNameRequest) error {
	return s.skillStorage.UpdateName(key, skill.Name)
}

func (s skillService) UpdateDescription(key string, skill UpdateSkillDescriptionRequest) error {
	return s.skillStorage.UpdateDescription(key, skill.Description)
}

func (s skillService) UpdateLogo(key string, skill UpdateSkillLogoRequest) error {
	return s.skillStorage.UpdateLogo(key, skill.Logo)
}

func (s skillService) UpdateTags(key string, skill UpdateSkillTagsRequest) error {
	return s.skillStorage.UpdateTags(key, skill.Tags)
}

func (s skillService) DeleteSkill(key string) error {
	return s.skillStorage.DeleteSkill(key)
}
