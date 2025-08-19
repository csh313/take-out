package global

import (
	"hmshop/config"
	"hmshop/logger"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	AppConfig   *config.Config
	Log         logger.ILog
	DB          *gorm.DB
	Redis       *redis.Client
	KafkaConfig config.Kafka
	Producer    sarama.SyncProducer
	Consumer    sarama.ConsumerGroup
	DBs         *gorm.DB
)
