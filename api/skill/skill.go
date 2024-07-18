package skill

type CreateSkillRequest struct {
	Key         string   `json:"key" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Logo        string   `json:"logo" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
}

type ResponseSkill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

type UpdateSkillRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Logo        string   `json:"logo" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
}

type UpdateSkillNameRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateSkillDescriptionRequest struct {
	Description string `json:"description" binding:"required"`
}

type UpdateSkillLogoRequest struct {
	Logo string `json:"logo" binding:"required"`
}

type UpdateSkillTagsRequest struct {
	Tags []string `json:"tags" binding:"required"`
}
