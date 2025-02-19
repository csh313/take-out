package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func UserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userRouter.POST("/login", userController.UserApi{}.Login)
	userRouter.Use(middle.AuthUserMiddleWare())

}
