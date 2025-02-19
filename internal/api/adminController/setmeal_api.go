package adminController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/service/adminService"
	"strconv"
	"strings"
)

type SetmealApi struct {
	service adminService.SetmealService
}

func (sm SetmealApi) AddSetmeal(c *gin.Context) {
	var mealReq req.SetMealDTO
	if err := c.ShouldBind(&mealReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	err := sm.service.AddSetmeal(mealReq, c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(code.AddSuccess, c)
}

func (sm SetmealApi) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.FailWithMessage(code.ReqError, c)
		return
	}

	resp, err := sm.service.GetById(id, c)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(*resp, c)
}

func (sm SetmealApi) PageQuery(c *gin.Context) {
	var page req.PageInfo
	if err := c.ShouldBind(&page); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	query, err := sm.service.PageQuery(page, c)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(query, c)
}

func (sm SetmealApi) DeleteByIds(c *gin.Context) {
	var ids string

	ids = c.Query("ids")

	fmt.Println(ids)
	//ids为id用逗号隔开的字符串,分割字符串
	idList := strings.Split(ids, ",")
	err := sm.service.DeleteByIds(idList, c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(code.DeleteSuccess, c)
}

func (sm SetmealApi) UpdateMeal(c *gin.Context) {
	var mealReq req.SetMealDTO
	if err := c.ShouldBind(&mealReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	err := sm.service.UpdateMeal(mealReq, c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)
}

func (sm SetmealApi) SetStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	err = sm.service.SetStatus(id, status, c)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)
}
