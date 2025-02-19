package adminController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/internal/service/adminService"
	"strconv"
)

type ShopApi struct {
	service adminService.ShopService
}

func (a ShopApi) GetStatus(c *gin.Context) {
	result, err := a.service.GetStatus(c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	status, _ := strconv.Atoi(result)

	res.OkWithData(status, c)

}

func (a ShopApi) SetStatus(c *gin.Context) {
	status := c.Param("status")
	a.service.SetStatus(status, c)
	res.OkWithMessage(code.EditSuccess, c)
}
