package skill

import "testing"

func TestConvertSkillType(t *testing.T) {
	t.Run("should return createSkillRequest", func(t *testing.T) {
		// Arrange
		key := "python"
		skillPayload := SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"key":         "figma",
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        []string{"tag"},
			},
			Action: CreateSkillAction,
		}

		// Act
		_, err := ConvertSkillType[CreateSkillRequest](skillPayload.Payload)
		if err != nil {
			t.Errorf("should return createSkillRequest")
		}
	})

	t.Run("should return error when failed to unmarshal skill data", func(t *testing.T) {
		// Arrange
		key := "python"
		skillPayload := SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"key":         "figma",
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        "tag",
			},
			Action: CreateSkillAction,
		}

		// Act
		_, err := ConvertSkillType[CreateSkillRequest](skillPayload.Payload)
		if err == nil {
			t.Errorf("should return error when failed to marshal skill data")
		}
	})

	t.Run("should return error when failed to marshal skill data", func(t *testing.T) {
		// Arrange
		key := "python"
		skillPayload := SkillQueuePayload{
			Key:     &key,
			Payload: make(chan int),
			Action:  CreateSkillAction,
		}

		// Act
		_, err := ConvertSkillType[CreateSkillRequest](skillPayload.Payload)
		if err.Error() != "json: unsupported type: chan int" {
			t.Errorf("should return error when failed to marshal skill data")
		}
	})
}
