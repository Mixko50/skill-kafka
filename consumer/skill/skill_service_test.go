package skill

import (
	"database/sql"
	"testing"
)

func TestSkillService_CreateSkill(t *testing.T) {
	t.Run("should be able to create new skill", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.CreateSkill(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"key":         "figma",
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        []string{"tag"},
			},
			Action: CreateSkillAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when json unmarshall error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.CreateSkill(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"key":         "figma",
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        "tag",
			},
			Action: "invalid",
		})

		// Assert
		if err.Error() != "failed to convert payload to CreateSkillRequest" {
			t.Errorf("expected error to be failed to convert payload to CreateSkillRequest, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.CreateSkill(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"key":         "figma",
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        []string{"tag"},
			},
			Action: CreateSkillAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}

func TestSkillService_UpdateSkill(t *testing.T) {
	t.Run("should be able to update skill", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateSkill(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        []string{"tag"},
			},
			Action: UpdateSkillAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when json unmarshall error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateSkill(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        "tag",
			},
			Action: "invalid",
		})

		// Assert
		if err.Error() != "failed to convert payload to UpdateSkillRequest" {
			t.Errorf("expected error to be failed to convert payload to UpdateSkillRequest, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateSkill(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"name":        "Figma",
				"description": "Figma is a vector bla bla",
				"logo":        "logo",
				"tags":        []string{"tag"},
			},
			Action: UpdateSkillAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}

func TestSkillService_UpdateName(t *testing.T) {
	t.Run("should be able to update name", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateName(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"name": "Figma",
			},
			Action: UpdateNameAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when json unmarshall error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateName(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"name": 1,
			},
			Action: UpdateNameAction,
		})

		// Assert
		if err.Error() != "failed to convert payload to UpdateSkillNameRequest" {
			t.Errorf("expected error to be failed to convert payload to UpdateSkillNameRequest, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateName(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"name": "Figma",
			},
			Action: UpdateNameAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}

func TestSkillService_UpdateDescription(t *testing.T) {
	t.Run("should be able to update description", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateDescription(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"description": "Figma is a vector bla bla",
			},
			Action: UpdateDescAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when json unmarshall error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateDescription(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"description": 1,
			},
			Action: UpdateDescAction,
		})

		// Assert
		if err.Error() != "failed to convert payload to UpdateSkillDescriptionRequest" {
			t.Errorf("expected error to be failed to convert payload to UpdateSkillDescriptionRequest, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateDescription(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"description": "Figma is a vector bla bla",
			},
			Action: UpdateDescAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}

func TestSkillService_UpdateLogo(t *testing.T) {
	t.Run("should be able to update logo", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateLogo(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"logo": "logo",
			},
			Action: UpdateLogoAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when json unmarshall error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateLogo(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"logo": 1,
			},
			Action: UpdateLogoAction,
		})

		// Assert
		if err.Error() != "failed to convert payload to UpdateSkillLogoRequest" {
			t.Errorf("expected error to be failed to convert payload to UpdateSkillLogoRequest, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateLogo(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"logo": "logo",
			},
			Action: UpdateLogoAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}

func TestSkillService_UpdateTags(t *testing.T) {
	t.Run("should be able to update tags", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateTags(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"tags": []string{"tag"},
			},
			Action: UpdateTagsAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when json unmarshall error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateTags(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"tags": "tag",
			},
			Action: UpdateTagsAction,
		})

		// Assert
		if err.Error() != "failed to convert payload to UpdateSkillTagsRequest" {
			t.Errorf("expected error to be failed to convert payload to UpdateSkillTagsRequest, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.UpdateTags(SkillQueuePayload{
			Key: &key,
			Payload: map[string]interface{}{
				"tags": []string{"tag"},
			},
			Action: UpdateTagsAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}

func TestSkillService_DeleteSkill(t *testing.T) {
	t.Run("should be able to delete skill", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.DeleteSkill(SkillQueuePayload{
			Key:    &key,
			Action: DeleteSkillAction,
		})

		// Assert
		if err != nil {
			t.Errorf("expected error to be nil, got %s", err)
		}
	})

	t.Run("should return error when database error", func(t *testing.T) {
		// Arrange
		s := mockSkillStorage{
			err: sql.ErrConnDone,
		}
		service := NewSkillService(&s)
		key := "figma"

		// Act
		err := service.DeleteSkill(SkillQueuePayload{
			Key:    &key,
			Action: DeleteSkillAction,
		})

		// Assert
		if err == nil {
			t.Error("expected error to be not nil")
		}
	})
}
