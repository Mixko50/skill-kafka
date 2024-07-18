package skill

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
