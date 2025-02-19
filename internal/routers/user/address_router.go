package user

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/userController"
	"hmshop/middle"
)

func AddressRouter(router *gin.RouterGroup) {
	addressRouter := router.Group("addressBook")

	addressRouter.Use(middle.AuthUserMiddleWare())
	{
		addressRouter.GET("list", userController.AddressApi{}.AddressList)
		addressRouter.GET("/:id", userController.AddressApi{}.GetById)
		addressRouter.GET("default", userController.AddressApi{}.GetDefaultAddress)
		addressRouter.PUT("default", userController.AddressApi{}.SetDefaultAddress)
		addressRouter.PUT("", userController.AddressApi{}.UpdateAddress)
		addressRouter.POST("", userController.AddressApi{}.AddAddress)
		addressRouter.DELETE("", userController.AddressApi{}.DeleteAddress)
	}

}
