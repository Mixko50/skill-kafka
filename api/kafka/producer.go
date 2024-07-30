package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"skill-api-kafka/config"
	"strings"
)

func Producer(c config.KafkaConfig) (sarama.SyncProducer, func()) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Return.Errors = true
	kafkaConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(strings.Split(c.KafkaBroker, ","), kafkaConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// Test
	msg := &sarama.ProducerMessage{Topic: "my-topic-test", Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

	return producer, func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}
}
