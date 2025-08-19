package routers

import (
	"hmshop/internal/routers/admin"
	"hmshop/internal/routers/user"
	"hmshop/internal/service"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	adminGroup := r.Group("/admin")
	admin.EmployeeRouter(adminGroup)
	admin.CategoryRouter(adminGroup)
	admin.DishRouter(adminGroup)
	admin.MealRouter(adminGroup)
	admin.OrderRouter(adminGroup)
	admin.CommonRouter(adminGroup)
	admin.ShopRouter(adminGroup)

	userGroup := r.Group("/user")
	user.UserRouter(userGroup)
	user.ShopRouter(userGroup)
	user.CategoryRouter(userGroup)
	user.DishRouter(userGroup)
	user.SetmealRouter(userGroup)
	user.AddressRouter(userGroup)
	user.ShoppingCartRouter(userGroup)
	user.OrderRouter(userGroup)

	r.GET("/ws/:id", service.WebsocketHandler)
	return r
}
