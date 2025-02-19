package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func DishRouter(router *gin.RouterGroup) {
	dishRouter := router.Group("dish")

	dishRouter.Use(middle.AuthUserMiddleWare())
	{
		dishRouter.GET("list", userController.DishApi{}.GetDishList)
	}

}
