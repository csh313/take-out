package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
	"hmshop/middle"
)

func OrderRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("order")

	orderRouter.Use(middle.AuthMiddleWare())
	{
		orderRouter.GET("details/:id", adminController.OrderApi{}.Details)
		orderRouter.PUT("confirm", adminController.OrderApi{}.OrderConfirm)
		orderRouter.PUT("rejection", adminController.OrderApi{}.Rejection)
		orderRouter.PUT("cancel", adminController.OrderApi{}.CancelOrder)
		orderRouter.PUT("delivery/:id", adminController.OrderApi{}.Delivery)
		orderRouter.PUT("complete/:id", adminController.OrderApi{}.Complete)
		orderRouter.GET("conditionSearch", adminController.OrderApi{}.Search)
		orderRouter.GET("statistics", adminController.OrderApi{}.Statistics)
	}

}
