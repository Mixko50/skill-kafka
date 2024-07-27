package kafka

import (
	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
	"golang.org/x/net/context"
	"skill-api-kafka-consumer/skill"
	"testing"
	"time"
)

type handlerMock struct {
	skill.SkillHandler
	msg string
}

func (h *handlerMock) ValidateSkillMessage(msg []byte) (*skill.SkillQueuePayload, error) {
	h.msg = string(msg)
	return &skill.SkillQueuePayload{}, nil
}

func (h *handlerMock) HandleSkill(payload *skill.SkillQueuePayload) error {
	return nil
}

func TestConsumer(t *testing.T) {
	// Arrange
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	handler := &handlerMock{}
	consumerMock := mocks.NewConsumer(t, nil)

	// Set expectation for the partition
	consumerMock.ExpectConsumePartition("skill", 0, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte(`create`)})

	// Start consuming the partition
	partition, err := consumerMock.ConsumePartition("skill", 0, sarama.OffsetOldest)
	if err != nil {
		t.Fatalf("Failed to start consuming partition: %v", err)
	}

	c := Consumer{
		broker:    "localhost:9092",
		topic:     "skill",
		consumer:  consumerMock,
		partition: partition,
	}
	defer c.Close()

	// Act
	c.Run(ctx, handler)

	// Assert
	if handler.msg != "create" {
		t.Errorf("expected %q but got %q", "create", handler.msg)
	}
}
