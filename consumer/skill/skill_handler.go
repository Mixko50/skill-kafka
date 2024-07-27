package skill

import (
	"encoding/json"
	"errors"
	"log"
)

type SkillService interface {
	CreateSkill(payload SkillQueuePayload) error
	UpdateSkill(payload SkillQueuePayload) error
	UpdateName(payload SkillQueuePayload) error
	UpdateDescription(payload SkillQueuePayload) error
	UpdateLogo(payload SkillQueuePayload) error
	UpdateTags(payload SkillQueuePayload) error
	DeleteSkill(payload SkillQueuePayload) error
}

type SkillHandler interface {
	HandleSkill(payload *SkillQueuePayload) error
	ValidateSkillMessage(msg []byte) (*SkillQueuePayload, error)
}

type skillHandler struct {
	skillService SkillService
}

func NewSkillHandler(skillService SkillService) skillHandler {
	return skillHandler{
		skillService: skillService,
	}
}

func (h skillHandler) HandleSkill(payload *SkillQueuePayload) error {
	switch payload.Action {
	case CreateSkillAction:
		return h.skillService.CreateSkill(*payload)
	case UpdateSkillAction:
		return h.skillService.UpdateSkill(*payload)
	case DeleteSkillAction:
		return h.skillService.DeleteSkill(*payload)
	case UpdateNameAction:
		return h.skillService.UpdateName(*payload)
	case UpdateDescAction:
		return h.skillService.UpdateDescription(*payload)
	case UpdateLogoAction:
		return h.skillService.UpdateLogo(*payload)
	case UpdateTagsAction:
		return h.skillService.UpdateTags(*payload)
	default:
		return ErrInvalidSkillAction
	}
}

func (h skillHandler) ValidateSkillMessage(msg []byte) (*SkillQueuePayload, error) {
	var payload *SkillQueuePayload
	err := json.Unmarshal(msg, &payload)
	if err != nil {
		return nil, err
	}

	log.Println("payload", payload)

	if payload.Action == "" {
		return nil, errors.New("action is empty")
	}

	if payload.Key == nil {
		return nil, errors.New("key is nil")
	}

	return payload, nil
}
