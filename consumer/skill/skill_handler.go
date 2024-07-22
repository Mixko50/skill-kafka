package skill

import (
	"errors"
	"fmt"
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
	var err error
	if payload.Key == nil {
		err = errors.New(fmt.Sprintf("cannot handle skill action without key at action: %s", payload.Action))
		return err
	}

	if payload.Action == "" {
		err = errors.New(fmt.Sprintf("cannot handle skill action key: %s without action", *payload.Key))
		return err
	}

	switch payload.Action {
	case CreateSkillAction:
		err = h.skillService.CreateSkill(*payload)
	case UpdateSkillAction:
		err = h.skillService.UpdateSkill(*payload)
	case DeleteSkillAction:
		err = h.skillService.DeleteSkill(*payload)
	case UpdateNameAction:
		err = h.skillService.UpdateName(*payload)
	case UpdateDescAction:
		err = h.skillService.UpdateDescription(*payload)
	case UpdateLogoAction:
		err = h.skillService.UpdateLogo(*payload)
	case UpdateTagsAction:
		err = h.skillService.UpdateTags(*payload)
	default:
		return errors.New(fmt.Sprintf("unknown skill action: %s", payload.Action))
	}

	if err != nil {
		return errors.New(fmt.Sprintf("failed to handle skill action: %s, key: %s, error: %s", payload.Action, *payload.Key, err))
	}

	return nil
}
