package cmd

import (
	"ecommerce-ums/helpers"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
)

func ServeKafka() {
	brokers := strings.Split(helpers.GetEnv("KAFKA_HOST", "localhost:9092"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC", "example-topid")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("failed to connect with kafka as consumer", err)
		return
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		helpers.Logger.Error("failed to create consumer partition 01", err)
		return
	}

	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s from partition %d\n", string(msg.Value), msg.Partition)
	}

}
