package kafkaService

import (
	"context"
	"github.com/IBM/sarama"
	"hmshop/global"
	"hmshop/internal/service"
	"log"
)

// 消费 Kafka 消息
func ConsumeMessages(consumerGroup sarama.ConsumerGroup) {
	ctx := context.Background() // 使用 Go 的标准 context
	for {
		// 使用你设定的 Kafka topic 和消费组
		if err := consumerGroup.Consume(ctx, []string{global.KafkaConfig.Topic}, &MessageHandler{}); err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

// 消息处理器
type MessageHandler struct{}

func (handler *MessageHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (handler *MessageHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (handler *MessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// 这里处理每一条消息
		log.Printf("Consumed message: %s", string(message.Value))

		// 假设消息中包含的是订单号或者其他信息
		orderNumber := string(message.Value) // 假设消息是订单号
		// 发送通知给商家（例如，WebSocket 通知）
		service.Server{}.SendToAllClients(orderNumber)

		log.Printf("Merchant notified for order %s", orderNumber)

		session.MarkMessage(message, "")
	}
	return nil
}
