package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func OrderRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("order")

	orderRouter.Use(middle.AuthUserMiddleWare())
	{
		orderRouter.POST("submit", userController.OrderApi{}.Submit)
		orderRouter.PUT("payment", userController.OrderApi{}.Pay)
		orderRouter.GET("orderDetail/:id", userController.OrderApi{}.OrderDetail)
		orderRouter.GET("historyOrders", userController.OrderApi{}.HistoryOrders)
		orderRouter.PUT("cancel/:id", userController.OrderApi{}.CancelOrder)
		orderRouter.POST("repetition/:id", userController.OrderApi{}.RepeatOrder)
		orderRouter.GET("reminder/:id", userController.OrderApi{}.RemindOrder)
	}

}
