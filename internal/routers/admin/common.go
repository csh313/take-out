package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
	"hmshop/middle"
)

func CommonRouter(router *gin.RouterGroup) {
	commonRouter := router.Group("common")

	commonRouter.Use(middle.AuthMiddleWare())
	{
		commonRouter.POST("/upload", adminController.CommonApi{}.Upload)
	}

}
