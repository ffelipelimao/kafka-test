package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	p := NewKafkaProducer()
	Publish("mensagem test", "tp.name", p, nil, deliveryChan)
	go DeliveryReport(deliveryChan)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return producer
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, channel chan kafka.Event) error {
	kafkaMsg := &kafka.Message{
		Key:   key,
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
	}

	// produce at goroutines, so i put this channel to ack my publish
	err := producer.Produce(kafkaMsg, channel)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func DeliveryReport(channel chan kafka.Event) {
	for e := range channel {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("error to send message")
			} else {
				// use case to publish in another message like DLQ
				fmt.Println("message delivery", ev.TopicPartition)
			}

		}
	}
}
