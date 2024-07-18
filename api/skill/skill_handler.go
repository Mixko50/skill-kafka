package skill

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"skill-api-kafka/types"
)

type SkillStorage interface {
	GetSkill(key string) (*Skill, error)
	GetSkills() ([]Skill, error)
}

type SkillQueue interface {
	PublishSkill(action SkillAction, key *string, skillPayload interface{}) error
}

type skillHandler struct {
	skillStorage SkillStorage
	skillQueue   SkillQueue
}

func NewSkillHandler(skillStorage SkillStorage, skillQueue SkillQueue) skillHandler {
	return skillHandler{
		skillStorage: skillStorage,
		skillQueue:   skillQueue,
	}
}

func (h skillHandler) GetSkill(c *gin.Context) {
	idParams := c.Param("key")
	skill, err := h.skillStorage.GetSkill(idParams)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, types.ErrorResponse("Skill not found"))
		return
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse(ResponseSkill{
		Key:         skill.Key,
		Name:        skill.Name,
		Description: skill.Description,
		Logo:        skill.Logo,
		Tags:        skill.Tags,
	}))
}

func (h skillHandler) GetSkills(c *gin.Context) {
	skills, err := h.skillStorage.GetSkills()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skills"))
		return
	}

	skillsMap := make([]ResponseSkill, 0)
	for _, skill := range skills {
		skillsMap = append(skillsMap, ResponseSkill{
			Key:         skill.Key,
			Name:        skill.Name,
			Description: skill.Description,
			Logo:        skill.Logo,
			Tags:        skill.Tags,
		})
	}

	c.JSON(http.StatusOK, types.SuccessResponse(skillsMap))
}

func (h skillHandler) CreateSkill(c *gin.Context) {
	var req CreateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse("not be able to create skill"))
		return
	}

	skill, err := h.skillStorage.GetSkill(req.Key)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if skill != nil {
		c.JSON(http.StatusConflict, types.ErrorResponse("Skill already exists"))
		return
	}

	if err := h.skillQueue.PublishSkill(CreateSkillAction, nil, req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to create skill"))
		return
	}

	c.JSON(http.StatusCreated, types.MessageResponse("creating skill already in progress"))
}

func (h skillHandler) UpdateSkill(c *gin.Context) {
	key := c.Param("key")

	var req UpdateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse("not be able to update skill"))
		return
	}

	skill, err := h.skillStorage.GetSkill(key)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse("skill not found"))
		return
	}

	if err := h.skillQueue.PublishSkill(UpdateSkillAction, &key, req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to update skill"))
		return
	}

	c.JSON(http.StatusOK, types.MessageResponse("updating skill already in progress"))
}

func (h skillHandler) UpdateName(c *gin.Context) {
	key := c.Param("key")

	var req UpdateSkillNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse("not be able to update name"))
		return
	}

	skill, err := h.skillStorage.GetSkill(key)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse("skill not found"))
		return
	}

	if err := h.skillQueue.PublishSkill(UpdateNameAction, &key, req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to update name"))
		return
	}

	c.JSON(http.StatusOK, types.MessageResponse("updating name already in progress"))
}

func (h skillHandler) UpdateDescription(c *gin.Context) {
	key := c.Param("key")

	var req UpdateSkillDescriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse("not be able to update description"))
		return
	}

	skill, err := h.skillStorage.GetSkill(key)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse("skill not found"))
		return
	}

	if err := h.skillQueue.PublishSkill(UpdateDescAction, &key, req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to update description"))
		return
	}

	c.JSON(http.StatusOK, types.MessageResponse("updating description already in progress"))
}

func (h skillHandler) UpdateLogo(c *gin.Context) {
	key := c.Param("key")

	var req UpdateSkillLogoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse("not be able to update logo"))
		return
	}

	skill, err := h.skillStorage.GetSkill(key)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse("skill not found"))
		return
	}

	if err := h.skillQueue.PublishSkill(UpdateLogoAction, &key, req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to update logo"))
		return
	}

	c.JSON(http.StatusOK, types.MessageResponse("updating logo already in progress"))
}

func (h skillHandler) UpdateTags(c *gin.Context) {
	key := c.Param("key")

	var req UpdateSkillTagsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, types.ErrorResponse("not be able to update tags"))
		return
	}

	skill, err := h.skillStorage.GetSkill(key)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if skill == nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse("skill not found"))
		return
	}

	if err := h.skillQueue.PublishSkill(UpdateTagsAction, &key, req); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to update tags"))
		return
	}

	c.JSON(http.StatusOK, types.MessageResponse("updating tags already in progress"))
}

func (h skillHandler) DeleteSkill(c *gin.Context) {
	key := c.Param("key")

	_, err := h.skillStorage.GetSkill(key)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, types.ErrorResponse("skill not found"))
		return
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to get skill"))
		return
	}

	if err := h.skillQueue.PublishSkill(DeleteSkillAction, nil, key); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, types.ErrorResponse("not be able to delete skill"))
		return
	}

	c.JSON(http.StatusOK, types.MessageResponse("deleting skill already in progress"))
}
