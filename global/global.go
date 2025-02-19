package global

import (
	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"hmshop/config"
	"hmshop/logger"
)

var (
	AppConfig   *config.Config
	Log         logger.ILog
	DB          *gorm.DB
	Redis       *redis.Client
	KafkaConfig config.Kafka
	Producer    sarama.SyncProducer
	Consumer    sarama.ConsumerGroup
)
