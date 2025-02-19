package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func ShoppingCartRouter(router *gin.RouterGroup) {

	shoppingCartRouter := router.Group("shoppingCart")
	shoppingCartRouter.Use(middle.AuthUserMiddleWare())
	{
		shoppingCartRouter.GET("", userController.ShoppingCartApi{}.List)
		shoppingCartRouter.POST("add", userController.ShoppingCartApi{}.AddCart)
		shoppingCartRouter.POST("/sub", userController.ShoppingCartApi{}.Delete)
		shoppingCartRouter.DELETE("/clean", userController.ShoppingCartApi{}.Clean)
	}
}
