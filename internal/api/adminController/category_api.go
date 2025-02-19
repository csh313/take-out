package adminController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/model"
	"hmshop/internal/service/adminService"
	"strconv"
)

type CategoryApi struct {
	service adminService.CategoryService
}

func (cg CategoryApi) PageCategory(c *gin.Context) {
	var page req.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		global.Log.Error(code.ReqError)
		res.FailWithMessage(code.ReqError, c)
	}
	var model model.Category
	list, err := res.PageList(page, model)
	if err != nil {
		res.FailWithMessage(code.ReqError, c)
	}
	res.OkWithData(list, c)
}

func (cg CategoryApi) AddCategory(c *gin.Context) {
	var request req.CategoryDTO
	if err := c.ShouldBind(&request); err != nil {
		global.Log.Error(code.ReqError)
		res.FailWithMessage(code.ReqError, c)
	}
	err := cg.service.AddCategory(c, request)
	if err != nil {
		res.FailWithMessage(code.ReqError, c)
	}
	res.OkWithMessage(code.AddSuccess, c)
}

func (cg CategoryApi) List(c *gin.Context) {
	typeId, err := strconv.Atoi(c.Query("type"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	fmt.Println("typeId", typeId)

	categoryModel, err := cg.service.List(c, typeId)
	if err != nil {
		res.FailWithMessage(code.ReqError, c)
	}
	res.OkWithData(categoryModel, c)

}

func (cg CategoryApi) DeleteById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err = cg.service.DeleteById(c, id)
	if err != nil {
		res.FailWithMessage(code.ReqError, c)
		return
	}
	res.OkWithMessage(code.DeleteSuccess, c)
}

func (cg CategoryApi) UpdateCategory(c *gin.Context) {
	var request req.CategoryDTO
	if err := c.ShouldBind(&request); err != nil {
		global.Log.Error(code.ReqError)
		fmt.Println("----cscscs")
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err := cg.service.EditCategory(c, request)
	if err != nil {
		res.FailWithMessage(code.ReqError, c)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)

}

func (cg CategoryApi) SetStatus(c *gin.Context) {
	//query：是查询参数,需要指明参数名称 如/id=2
	//param：是路由参数，不需要指明参数名称 如/2
	id, _ := strconv.ParseUint(c.Query("id"), 10, 64)
	status, _ := strconv.Atoi(c.Param("status"))
	fmt.Println(id, status, "------")
	err := cg.service.SetStatus(c, id, status)
	if err != nil {
		res.FailWithMessage(code.EditError, c)
		return
	}
	res.OkWithMessage(code.EditSuccess, c)
}
