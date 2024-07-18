package skill

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"skill-api-kafka/config"
)

type SkillAction string

const (
	CreateSkillAction SkillAction = "create"
	UpdateSkillAction SkillAction = "update"
	DeleteSkillAction SkillAction = "delete"
	UpdateNameAction  SkillAction = "update_name"
	UpdateDescAction  SkillAction = "update_desc"
	UpdateLogoAction  SkillAction = "update_logo"
	UpdateTagsAction  SkillAction = "update_tags"
)

type SkillQueuePayload struct {
	Action  SkillAction `json:"action"`
	Key     *string     `json:"key"`
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

func (q skillQueue) PublishSkill(action SkillAction, key *string, skillPayload interface{}) error {
	payload := SkillQueuePayload{
		Action:  action,
		Key:     key,
		Payload: skillPayload,
	}

	message, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, _, err = q.producer.SendMessage(&sarama.ProducerMessage{
		Topic: q.config.SkillTopic,
		Value: sarama.StringEncoder(message),
	})

	return err
}
