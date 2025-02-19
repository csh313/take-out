package adminService

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"hmshop/common/res"
	"hmshop/global"
)

type ShopService struct {
}

func (s ShopService) GetStatus(c *gin.Context) (string, error) {
	result, err := global.Redis.Get(c, "shop_status").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", errors.New("redis is null")
		} else {
			return "", err
		}
	}
	return result, nil
}

func (s ShopService) SetStatus(status string, c *gin.Context) {
	if err := global.Redis.Set(c, "shop_status", status, 0).Err(); err != nil {
		res.FailWithMessage(err.Error(), c)
	}
}
