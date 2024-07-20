package main

import (
	"database/sql"
	"fmt"
	"log"
	"skill-api-kafka-consumer/config"
	"skill-api-kafka-consumer/database"
	"skill-api-kafka-consumer/kafka"
	"skill-api-kafka-consumer/skill"
)

func main() {
	c := config.Configuration()

	db := database.Postgres(c.PostgresURI)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Fail to close database connection: %v", err)
		}
		log.Println("Database connection closed")
	}(db)

	skillStorage := skill.NewSkillStorage(db)
	skillService := skill.NewSkillService(skillStorage)
	skillHandler := skill.NewSkillHandler(skillService)

	consumer := kafka.NewConsumer(c.Kafka.KafkaConsumer, c.Kafka.SkillTopic)
	defer func() {
		consumer.ClosePartition()
		fmt.Println("Kafka Partition closed")
		consumer.Close()
		fmt.Println("Kafka Consumer closed")
	}()

	consumer.Run(skillHandler)
}
