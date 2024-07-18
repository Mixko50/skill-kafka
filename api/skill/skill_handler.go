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

type skillHandler struct {
	skillStorage SkillStorage
}

func NewSkillHandler(skillStorage SkillStorage) skillHandler {
	return skillHandler{
		skillStorage: skillStorage,
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
