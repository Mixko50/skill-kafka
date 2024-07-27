package skill

import (
	"errors"
	"testing"
)

func TestValidateSkillMessageHandler(t *testing.T) {
	t.Run("should be able to validate skill message", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)

		// Act
		_, err := h.ValidateSkillMessage([]byte(`{"action":"create","key":"python"}`))

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should not be able to perform when message is empty", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)

		// Act
		_, err := h.ValidateSkillMessage([]byte(``))

		// Assert
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should not be able to perform when action is empty", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)

		// Act
		_, err := h.ValidateSkillMessage([]byte(`{"data" : "test"}`))

		// Assert
		if err.Error() != "action is empty" {
			t.Errorf("expected error, got %s", err)
		}
	})

	t.Run("should not be able to perform without key", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)

		// Act
		_, err := h.ValidateSkillMessage([]byte(`{"action":"create"}`))

		// Assert
		if err.Error() != "key is nil" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleSkill(t *testing.T) {
	t.Run("should be able to handle skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action:  CreateSkillAction,
			Key:     nil,
			Payload: nil,
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should not be able to handle skill when action is invalid", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action:  "invalid",
			Key:     nil,
			Payload: nil,
		})

		// Assert
		if err.Error() != ErrInvalidSkillAction.Error() {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleCreateSkill(t *testing.T) {
	t.Run("should be able to create new skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: CreateSkillAction,
			Key:    &key,
			Payload: &CreateSkillRequest{
				Key:         "python",
				Name:        "Python",
				Description: "description",
				Logo:        "logo",
				Tags:        []string{"tag1"},
			},
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to create skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: CreateSkillAction,
			Key:    &key,
			Payload: &CreateSkillRequest{
				Key:         "python",
				Name:        "Python",
				Description: "description",
				Logo:        "logo",
				Tags:        []string{"tag1"},
			},
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleUpdateSkill(t *testing.T) {
	t.Run("should be able to update skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateSkillAction,
			Key:    &key,
			Payload: &UpdateSkillRequest{
				Name:        "Python",
				Description: "description",
				Logo:        "logo",
				Tags:        []string{"tag1"},
			},
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to update skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateSkillAction,
			Key:    &key,
			Payload: &UpdateSkillRequest{
				Name:        "Python",
				Description: "description",
				Logo:        "logo",
				Tags:        []string{"tag1"},
			},
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleUpdateName(t *testing.T) {
	t.Run("should be able to update skill name", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateNameAction,
			Key:    &key,
			Payload: &UpdateSkillNameRequest{
				Name: "Python",
			},
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to update skill name", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateNameAction,
			Key:    &key,
			Payload: &UpdateSkillNameRequest{
				Name: "Python",
			},
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleUpdateDescription(t *testing.T) {
	t.Run("should be able to update skill description", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateDescAction,
			Key:    &key,
			Payload: &UpdateSkillDescriptionRequest{
				Description: "description",
			},
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to update skill description", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateDescAction,
			Key:    &key,
			Payload: &UpdateSkillDescriptionRequest{
				Description: "description",
			},
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleUpdateLogo(t *testing.T) {
	t.Run("should be able to update skill logo", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateLogoAction,
			Key:    &key,
			Payload: &UpdateSkillLogoRequest{
				Logo: "logo",
			},
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to update skill logo", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateLogoAction,
			Key:    &key,
			Payload: &UpdateSkillLogoRequest{
				Logo: "logo",
			},
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleUpdateTags(t *testing.T) {
	t.Run("should be able to update skill tags", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateTagsAction,
			Key:    &key,
			Payload: &UpdateSkillTagsRequest{
				Tags: []string{"tag1"},
			},
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to update skill tags", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: UpdateTagsAction,
			Key:    &key,
			Payload: &UpdateSkillTagsRequest{
				Tags: []string{"tag1"},
			},
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}

func TestHandleDeleteSkill(t *testing.T) {
	t.Run("should be able to delete skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: DeleteSkillAction,
			Key:    &key,
		})

		// Assert
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("should error to delete skill", func(t *testing.T) {
		// Arrange
		s := mockSkillService{
			err: errors.New("error"),
		}
		h := NewSkillHandler(s)
		key := "python"

		// Act
		err := h.HandleSkill(&SkillQueuePayload{
			Action: DeleteSkillAction,
			Key:    &key,
		})

		// Assert
		if err.Error() != "error" {
			t.Errorf("expected error, got %s", err)
		}
	})
}
