package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
	"hmshop/middle"
)

func DishRouter(router *gin.RouterGroup) {
	dishRouter := router.Group("dish")
	{

		dishRouter.GET("/page", adminController.DishApi{}.DishPage)
		dishRouter.GET("/:id", adminController.DishApi{}.GetById)
		dishRouter.GET("/list", adminController.DishApi{}.List) //根据分类id获取菜品

	}
	dishRouter.Use(middle.AuthMiddleWare())
	{
		dishRouter.POST("", adminController.DishApi{}.AddDish)
		dishRouter.DELETE("", adminController.DishApi{}.DeleteDish)
		dishRouter.PUT("", adminController.DishApi{}.UpdateDish)
	}
}
