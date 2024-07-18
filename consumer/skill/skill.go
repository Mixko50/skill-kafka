package skill

type SkillAction string

const (
	CreateSkillAction SkillAction = "create"
	UpdateSkillAction SkillAction = "update"
	DeleteSkillAction SkillAction = "delete"
	UpdateNameAction  SkillAction = "update_name"
	UpdateDescAction  SkillAction = "update_desc"
	UpdateLogoAction  SkillAction = "update_logo"
	UpdateTagsAction  SkillAction = "update_tags"
)

type SkillQueuePayload struct {
	Action  SkillAction `json:"action"`
	Key     *string     `json:"key"`
	Payload any         `json:"payload"`
}

type UpdateSkillRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

type CreateSkillRequest struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

type UpdateSkillNameRequest struct {
	Name string `json:"name"`
}

type UpdateSkillDescriptionRequest struct {
	Description string `json:"description"`
}

type UpdateSkillLogoRequest struct {
	Logo string `json:"logo"`
}

type UpdateSkillTagsRequest struct {
	Tags []string `json:"tags"`
}
