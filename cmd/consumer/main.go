package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "consumer-xx",
		"group.id":          "group-xx",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	topics := []string{"tp.name"}
	consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
