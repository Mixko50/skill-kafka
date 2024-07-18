package config

import (
	"log"
	"os"
)

type Config struct {
	PostgresURI string
	Port        string
	Kafka       KafkaConfig
}

type KafkaConfig struct {
	KafkaConsumer string
	SkillTopic    string
}

func Configuration() Config {
	if os.Getenv("PORT") == "" {
		log.Fatal("PORT is not set")
	}

	if os.Getenv("POSTGRES_URI") == "" {
		log.Fatal("POSTGRES_URI is not set")
	}

	if os.Getenv("KAFKA_CONSUMER") == "" {
		log.Fatal("KAFKA_CONSUMER is not set")
	}

	if os.Getenv("KAFKA_SKILL_TOPIC") == "" {
		log.Fatal("KAFKA_SKILL_TOPIC is not set")
	}

	return Config{
		PostgresURI: os.Getenv("POSTGRES_URI"),
		Port:        os.Getenv("PORT"),
		Kafka: KafkaConfig{
			KafkaConsumer: os.Getenv("KAFKA_CONSUMER"),
			SkillTopic:    os.Getenv("KAFKA_SKILL_TOPIC"),
		},
	}
}
