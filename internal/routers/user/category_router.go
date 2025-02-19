package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func CategoryRouter(router *gin.RouterGroup) {
	categoryRouter := router.Group("category")
	categoryRouter.Use(middle.AuthUserMiddleWare())
	{
		categoryRouter.GET("/list", userController.CategoryApi{}.GetCategoryByType)
	}

}
