package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func NewKafkaProducer() *kafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv(key: "KafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
        log.Println(err.error())
    }
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		value: []byte(msg),
	}
	err := producer.Producer(message, deliveryChan: nil)
    if err != nil {
		return err
	}
	return nil
}