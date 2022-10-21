package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(MsgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
        MsgChan: MsgChan,
    }
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv(key: "KafkaBootstrapServers"),
		"group.id":          os.Getenv(key: "KafkaConsumerGroupID"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err!= nil {
        log.Fatal(format: "error consuming kafka message" + err.Error())
    }
	topics := []string{os.Getenv(key: "KafkaReadTopic")}
	c.SubscribeTopics(topics, rebalanceCb: nil)
	fmt.Println(a...: "Kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(timeout: -1)
        if err == nil {
            k.MsgChan <- msg
        }
	}
}