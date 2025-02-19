package initialize

import (
	"github.com/IBM/sarama"
	"hmshop/config"
	"hmshop/global"
)

func InitKafkaConfig() config.Kafka {
	kafkaConfig := global.AppConfig.Kafka
	return kafkaConfig
}

// 初始化 Kafka 生产者
func InitProducer() sarama.SyncProducer {
	kafkaConfig := global.KafkaConfig // 直接从 global 包中读取配置

	// 创建生产者配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer(kafkaConfig.Brokers, config)
	if err != nil {
		global.Log.Error("创建生产者失败")
		panic(err)
		return nil
	}

	return producer
}

// 初始化 Kafka 消费者
func InitConsumer() sarama.ConsumerGroup {
	kafkaConfig := global.KafkaConfig // 直接从 global 包中读取配置

	// 创建消费者配置
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	// 创建 Kafka 消费者
	consumerGroup, err := sarama.NewConsumerGroup(kafkaConfig.Brokers, kafkaConfig.GroupId, config)
	if err != nil {
		global.Log.Error("创建消费者失败。。。")
		panic(err)
		return nil
	}

	return consumerGroup
}
