package main

import (
	"skill-api-kafka-consumer/config"
	"skill-api-kafka-consumer/database"
	"skill-api-kafka-consumer/kafka"
	"skill-api-kafka-consumer/skill"
)

func main() {
	c := config.Configuration()

	db := database.Postgres(c.PostgresURI)
	defer db.Close()

	skillStorage := skill.NewSkillStorage(db)
	skillService := skill.NewSkillService(skillStorage)
	skillHandler := skill.NewSkillHandler(skillService)

	consumer := kafka.NewConsumer(c.Kafka.KafkaConsumer, c.Kafka.SkillTopic)
	defer func() {
		consumer.ClosePartition()
		consumer.Close()
	}()

	consumer.Run(skillHandler)
}
