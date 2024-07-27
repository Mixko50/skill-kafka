package skill

import (
	"encoding/json"
	"errors"
	"log"
)

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

var (
	ErrInvalidSkillAction = errors.New("invalid skill action")
	ErrorInvalidPayload   = errors.New("invalid payload")
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

func ConvertSkillType[t UpdateSkillTagsRequest | UpdateSkillLogoRequest | UpdateSkillDescriptionRequest | UpdateSkillNameRequest | UpdateSkillRequest | CreateSkillRequest](skill any) (*t, error) {
	byteData, err := json.Marshal(skill)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var payload *t
	err = json.Unmarshal(byteData, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
