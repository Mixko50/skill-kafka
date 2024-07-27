package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os/signal"
	"skill-api-kafka-consumer/config"
	"skill-api-kafka-consumer/database"
	"skill-api-kafka-consumer/kafka"
	"skill-api-kafka-consumer/skill"
	"syscall"
	"time"
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

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func() {
		<-ctx.Done()

		timeOut, cancelTimeout := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelTimeout()
		consumer.ClosePartition()
		fmt.Println("Kafka Partition closed")
		consumer.Close()
		fmt.Println("Kafka Consumer closed")

		<-timeOut.Done()

		err := db.Close()
		if err != nil {
			log.Fatalf("Fail to close database connection: %v", err)
		}
	}()

	consumer.Run(ctx, skillHandler)
}
