package skill

import (
	"encoding/json"
	"log"
)

type SkillHandler interface {
	HandleSkill(payload *SkillQueuePayload)
}

type skillHandler struct {
	skillService skillService
}

func NewSkillHandler(skillService skillService) skillHandler {
	return skillHandler{
		skillService: skillService,
	}
}

func (h skillHandler) HandleSkill(payload *SkillQueuePayload) {
	if payload.Key == nil {
		log.Printf("cannot handle skill action: %s without key", payload.Action)
		return
	}
	var err error
	switch payload.Action {
	case CreateSkillAction:
		req, err := ConvertSkillType[CreateSkillRequest](payload.Payload)
		if err != nil {
			log.Printf("cannot convert payload to CreateSkillRequest with error: %v", err)
			return
		}
		err = h.skillService.CreateSkill(*req)
	case UpdateSkillAction:
		req, err := ConvertSkillType[UpdateSkillRequest](payload.Payload)
		if err != nil {
			log.Printf("cannot convert payload to UpdateSkillRequest with error: %v", err)
			return
		}
		err = h.skillService.UpdateSkill(*payload.Key, *req)
	case UpdateNameAction:
		req, err := ConvertSkillType[UpdateSkillNameRequest](payload.Payload)
		if err != nil {
			log.Printf("cannot convert payload to UpdateSkillNameRequest with error: %v", err)
			return
		}
		err = h.skillService.UpdateName(*payload.Key, *req)
	case UpdateDescAction:
		req, err := ConvertSkillType[UpdateSkillDescriptionRequest](payload.Payload)
		if err != nil {
			log.Printf("cannot convert payload to UpdateSkillDescriptionRequest with error: %v", err)
			return
		}
		err = h.skillService.UpdateDescription(*payload.Key, *req)
	case UpdateLogoAction:
		req, err := ConvertSkillType[UpdateSkillLogoRequest](payload.Payload)
		if err != nil {
			log.Printf("cannot convert payload to UpdateSkillLogoRequest with error: %v", err)
			return
		}
		err = h.skillService.UpdateLogo(*payload.Key, *req)
	case UpdateTagsAction:
		req, err := ConvertSkillType[UpdateSkillTagsRequest](payload.Payload)
		if err != nil {
			log.Printf("cannot convert payload to UpdateSkillTagsRequest with error: %v", err)
			return
		}
		err = h.skillService.UpdateTags(*payload.Key, *req)
	case DeleteSkillAction:
		err = h.skillService.DeleteSkill(*payload.Key)
	}

	if err != nil {
		log.Printf("cannot handle skill action: %s with error: %v", payload.Action, err)
	}
}

func ConvertSkillType[t UpdateSkillTagsRequest | UpdateSkillLogoRequest | UpdateSkillDescriptionRequest | UpdateSkillNameRequest | UpdateSkillRequest | CreateSkillRequest](skill any) (*t, error) {
	byteData, err := json.Marshal(skill)
	if err != nil {
		return nil, err
	}

	var payload *t

	err = json.Unmarshal(byteData, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
