package userService

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/res"
	"hmshop/global"
)

type ShopService struct {
}

func (s ShopService) ShopStatus(c *gin.Context) string {
	result, err := global.Redis.Get(c, "shop_status").Result()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("获取店铺状态失败", c)
	}
	return result
}
