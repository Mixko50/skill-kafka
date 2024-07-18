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
