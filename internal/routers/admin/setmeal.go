package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
	"hmshop/middle"
)

func MealRouter(router *gin.RouterGroup) {
	setmealRouter := router.Group("setmeal")
	{

	}
	setmealRouter.Use(middle.AuthMiddleWare())
	{
		setmealRouter.POST("", adminController.SetmealApi{}.AddSetmeal)
		setmealRouter.GET("/:id", adminController.SetmealApi{}.GetById)
		setmealRouter.GET("/page", adminController.SetmealApi{}.PageQuery)
		setmealRouter.DELETE("", adminController.SetmealApi{}.DeleteByIds)
		setmealRouter.PUT("", adminController.SetmealApi{}.UpdateMeal)
		setmealRouter.POST("/status/:status", adminController.SetmealApi{}.SetStatus)
	}
}
