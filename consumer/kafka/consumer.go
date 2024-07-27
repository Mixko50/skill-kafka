package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"skill-api-kafka-consumer/skill"
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

func (c *Consumer) Run(ctx context.Context, h skill.SkillHandler) {
	fmt.Printf("Consuming topic %s at %s.\n", c.topic, c.broker)

ConsumerLoop:
	for {
		select {
		case msg := <-c.partition.Messages():

			payload, err := h.ValidateSkillMessage(msg.Value)
			if err != nil {
				log.Printf("Error validating message at topic: %s, partition: %d, offset: %d, error: %s", msg.Topic, msg.Partition, msg.Offset, err)
				continue
			}

			if err := h.HandleSkill(payload); err != nil {
				log.Printf("Error handling message at topic: %s, partition: %d, offset: %d, error: %s", msg.Topic, msg.Partition, msg.Offset, err)
				continue
			}

			log.Printf("Successfully handled message at topic: %s, partition: %d, offset %d", msg.Topic, msg.Partition, msg.Offset)

		case <-ctx.Done():
			log.Print("Shutting down consumer...")
			break ConsumerLoop
		}
	}
}
