package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
)

func ShopRouter(router *gin.RouterGroup) {
	shopRouter := router.Group("shop")
	{
		shopRouter.GET("/status", adminController.ShopApi{}.GetStatus)
		shopRouter.PUT(":status", adminController.ShopApi{}.SetStatus)
	}
}
