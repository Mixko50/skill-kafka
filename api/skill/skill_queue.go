package skill

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"skill-api-kafka/config"
)

type SkillQueuePayload struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type skillQueue struct {
	producer sarama.SyncProducer
	config   config.KafkaConfig
}

func NewSkillQueue(producer sarama.SyncProducer, config config.KafkaConfig) skillQueue {
	return skillQueue{
		producer: producer,
		config:   config,
	}
}

func (q skillQueue) PublishSkill(action string, skillPayload interface{}) error {
	payload := SkillQueuePayload{
		Action:  action,
		Payload: skillPayload,
	}

	message, err := json.Marshal(payload)

	_, _, err = q.producer.SendMessage(&sarama.ProducerMessage{
		Topic: q.config.SkillTopic,
		Value: sarama.StringEncoder(message),
	})

	return err
}
