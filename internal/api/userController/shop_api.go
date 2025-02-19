package userController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/res"
	"hmshop/internal/service/userService"
	"strconv"
)

type ShopApi struct {
	service userService.ShopService
}

func (a ShopApi) ShopStatus(c *gin.Context) {

	result := a.service.ShopStatus(c)
	status, _ := strconv.Atoi(result)
	res.OkWithData(status, c)
}
