package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
)

func ShopRouter(router *gin.RouterGroup) {
	shopRouter := router.Group("shop")
	shopRouter.GET("/status", userController.ShopApi{}.ShopStatus)
}
