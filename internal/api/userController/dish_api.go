package userController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/service/userService"
	"strconv"
)

type DishApi struct {
	service userService.DishService
}

func (d DishApi) GetDishList(c *gin.Context) {
	value := c.Query("categoryId")
	categoryId, err := strconv.Atoi(value)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	dishes := d.service.ListDish(categoryId, c)
	if len(dishes) == 0 || dishes == nil {
		res.FailWithMessage(code.DataNotFound, c)
		return
	}
	res.OkWithData(dishes, c)
}
