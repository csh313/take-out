package kafkaService

import (
	"github.com/IBM/sarama"
	"hmshop/global"
	"log"
)

// 发送消息到 Kafka
func SendMessage(producer sarama.SyncProducer, message string) error {

	msg := &sarama.ProducerMessage{
		Topic: global.KafkaConfig.Topic,
		Value: sarama.StringEncoder(message),
	}

	// 发送消息
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent to Kafka topic %s: %s", global.KafkaConfig.Topic, message)
	return nil
}
