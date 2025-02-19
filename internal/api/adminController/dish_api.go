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

type DishApi struct {
	service adminService.DishService
}

func (ds DishApi) AddDish(c *gin.Context) {
	var dishReq req.DishDTO

	if err := c.ShouldBind(&dishReq); err != nil {
		fmt.Println(dishReq)
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err := ds.service.AddDish(c, dishReq)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.AddError, c)
		return
	}
	res.OkWithMessage(code.AddSuccess, c)
}

func (ds DishApi) DishPage(c *gin.Context) {
	var dishReq req.PageInfo
	if err := c.ShouldBindQuery(&dishReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	page, err := ds.service.DishPage(c, dishReq)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
		return
	}

	res.OkWithData(page, c)
}

func (ds DishApi) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	dish, err := ds.service.GetById(c, id)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	res.OkWithData(dish, c)
}

func (ds DishApi) List(c *gin.Context) {
	//query：是查询参数,需要指明参数名称 如/id=2
	//param：是路由参数，不需要指明参数名称 如/2
	categoryId, err := strconv.ParseUint(c.Query("categoryId"), 10, 64)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	list, err := ds.service.List(categoryId, c)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
		return
	}
	res.OkWithData(list, c)

}

func (ds DishApi) DeleteDish(c *gin.Context) {
	var ids string

	ids = c.Query("ids")

	fmt.Println(ids)
	//ids为id用逗号隔开的字符串,分割字符串
	idList := strings.Split(ids, ",")
	err := ds.service.DeleteDish(idList, c)
	if err != nil {
		res.FailWithMessage(code.DeleteError, c)
		global.Log.Error(err)
		return
	}
	res.OkWithMessage(code.DeleteSuccess, c)
}
func (ds DishApi) UpdateDish(c *gin.Context) {
	var dishReq req.DishUpdateDTO
	if err := c.ShouldBind(&dishReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err := ds.service.UpdateDish(c, dishReq)
	if err != nil {
		res.FailWithMessage(code.EditError, c)
		global.Log.Error(err)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)
}
