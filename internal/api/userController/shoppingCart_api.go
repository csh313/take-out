package userController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/service/userService"
)

type ShoppingCartApi struct {
	service userService.ShoppingCartService
}

func (sc ShoppingCartApi) List(c *gin.Context) {
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		global.Log.Error(errors.New("查询用户信息错误"))
		res.FailWithMessage(code.ReqError, c)
		return
	}
	list := sc.service.List(value, c)
	res.OkWithData(list, c)
}

func (sc ShoppingCartApi) AddCart(c *gin.Context) {
	var cartReq req.ShoppingCartDTO
	if err := c.ShouldBind(&cartReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err := sc.service.AddCart(cartReq, c)
	if err != nil {
		return
	}
	res.OkWithMessage(code.AddSuccess, c)
}

func (sc ShoppingCartApi) Delete(c *gin.Context) {
	var cartReq req.ShoppingCartDTO
	if err := c.ShouldBind(&cartReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	err := sc.service.Delete(cartReq, c)
	if err != nil {
		return
	}
	res.OkWithMessage(code.DeleteSuccess, c)
}

func (sc ShoppingCartApi) Clean(c *gin.Context) {
	err := sc.service.Clean(c)
	if err != nil {
		return
	}
	res.OkWithMessage(code.DeleteSuccess, c)
}
