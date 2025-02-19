package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
	"hmshop/middle"
)

func CategoryRouter(router *gin.RouterGroup) {
	categoryRouter := router.Group("category")
	{
		categoryRouter.GET("/page", adminController.CategoryApi{}.PageCategory)
	}
	categoryRouter.Use(middle.AuthMiddleWare())
	{
		categoryRouter.POST("/add", adminController.CategoryApi{}.AddCategory)
		categoryRouter.GET("/list", adminController.CategoryApi{}.List) //根据类型查询分类
		categoryRouter.DELETE("", adminController.CategoryApi{}.DeleteById)
		categoryRouter.PUT("", adminController.CategoryApi{}.UpdateCategory)
		categoryRouter.POST("status/:status", adminController.CategoryApi{}.SetStatus)

	}

}
