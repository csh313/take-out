package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hmshop/global"
)

func InitRedis() *redis.Client {

	redisOpt := global.AppConfig.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisOpt.Host + ":" + redisOpt.Port,
		Password: redisOpt.Password,
		DB:       redisOpt.Database,
	})
	ping := redisClient.Ping(context.Background())
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	return redisClient
}
