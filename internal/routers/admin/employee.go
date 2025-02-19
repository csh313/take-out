package admin

import (
	"github.com/gin-gonic/gin"
	"hmshop/internal/api/adminController"
	"hmshop/middle"
)

func EmployeeRouter(router *gin.RouterGroup) {
	employeeRouter := router.Group("employee")
	{
		employeeRouter.POST("/login", adminController.EmployeeApi{}.Login)
		employeeRouter.POST("/register", adminController.EmployeeApi{}.Register)
	}
	employeeRouter.Use(middle.AuthMiddleWare())
	{
		employeeRouter.GET("/logout", adminController.EmployeeApi{}.Logout)
		employeeRouter.GET("/:id", adminController.EmployeeApi{}.GetEmployee)
		employeeRouter.GET("/page", adminController.EmployeeApi{}.PageEmployee)
		employeeRouter.PUT("", adminController.EmployeeApi{}.UpdateEmployee)
		employeeRouter.PUT("/editPassword", adminController.EmployeeApi{}.UpdatePassword)
		employeeRouter.POST("/status/:status", adminController.EmployeeApi{}.UpdateStatus)
	}

}
