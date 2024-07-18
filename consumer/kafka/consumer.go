package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os/signal"
	"skill-api-kafka-consumer/skill"
	"syscall"
)

type Consumer struct {
	broker    string
	topic     string
	consumer  sarama.Consumer
	partition sarama.PartitionConsumer
}

func NewConsumer(broker string, topic string) *Consumer {
	consumer, err := sarama.NewConsumer([]string{broker}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	partition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	return &Consumer{
		consumer:  consumer,
		partition: partition,
		topic:     topic,
		broker:    broker,
	}
}

func (c *Consumer) Close() {
	if err := c.consumer.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Consumer) ClosePartition() {
	if err := c.partition.Close(); err != nil {
		log.Fatal(err)
	}
}

func (c *Consumer) Run(h skill.SkillHandler) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	fmt.Printf("Consuming topic %s at %s.\n", c.topic, c.broker)

ConsumerLoop:
	for {
		select {
		case msg := <-c.partition.Messages():
			log.Printf("Consumed message offset %d", msg.Offset)

			var payload *skill.SkillQueuePayload

			err := json.Unmarshal(msg.Value, &payload)
			if err != nil {
				log.Printf("Error unmarshalling message: %s, offset: %d", err, msg.Offset)
				continue
			}

			if payload == nil {
				log.Printf("Payload is nil, offset: %d", msg.Offset)
				continue
			}

			h.HandleSkill(payload)
		case <-ctx.Done():
			log.Print("Shutting down consumer...")
			break ConsumerLoop
		}
	}
}
