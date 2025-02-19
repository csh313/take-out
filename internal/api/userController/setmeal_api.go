package userController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/service/userService"
	"strconv"
)

type SetmealApi struct {
	service userService.SetmealService
}

func (sm SetmealApi) SetmealByCategoryId(c *gin.Context) {

	value := c.Query("categoryId")
	categoryId, err := strconv.Atoi(value)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	setMealPageQueryVo := sm.service.SetmealByCategoryId(categoryId, c)
	res.OkWithData(setMealPageQueryVo, c)
}

func (sm SetmealApi) GetDishById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	dishItemVo, err := sm.service.GetDishById(id, c)
	if err != nil {
		return
	}
	if len(dishItemVo) == 0 {
		res.OkWithMessage(code.DataNotFound, c)
		return
	}
	res.OkWithData(dishItemVo, c)
}
