package skill

import (
	"errors"
	"fmt"
)

type SkillHandler interface {
	HandleSkill(payload *SkillQueuePayload) error
}

type skillHandler struct {
	skillService skillService
}

func NewSkillHandler(skillService skillService) skillHandler {
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
		err = errors.New(fmt.Sprintf("unknown skill action: %s", payload.Action))
	}

	return err
}
