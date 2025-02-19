package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func SetmealRouter(router *gin.RouterGroup) {
	setmealRouter := router.Group("setmeal")
	setmealRouter.Use(middle.AuthUserMiddleWare())
	{
		setmealRouter.GET("list", userController.SetmealApi{}.SetmealByCategoryId)
		setmealRouter.GET("dish/:id", userController.SetmealApi{}.GetDishById)
	}

}
