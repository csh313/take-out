package adminController

import (
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/service/adminService"
	"strconv"
)

type OrderApi struct {
	service adminService.OrderService
}

func (od OrderApi) Details(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	details, err := od.service.Details(id)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
		return
	}
	res.OkWithData(*details, c)
}

func (od OrderApi) OrderConfirm(c *gin.Context) {
	var confirm int
	if err := c.ShouldBind(&confirm); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}

	if err := od.service.Confirm(confirm, c); err != nil {
		global.Log.Error(err)
		res.FailWithMessage("接单失败", c)
		return
	}
	res.OkWithMessage("接单成功", c)
}

func (od OrderApi) Rejection(c *gin.Context) {
	var rejection req.OrderRejectionDTO
	if err := c.ShouldBind(&rejection); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err := od.service.Rejection(rejection)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("拒单失败", c)
		return
	}
	res.OkWithMessage("拒单成功", c)

}

func (od OrderApi) CancelOrder(c *gin.Context) {
	var cancel req.OrderCancelDTO
	if err := c.ShouldBind(&cancel); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err := od.service.CancelOrder(cancel)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("取消订单成功", c)
}

func (od OrderApi) Delivery(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("id"))
	err := od.service.Delivery(orderId, c)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.EditError+"-派送订单", c)
		return
	}
	res.OkWithMessage(code.EditSuccess+"派送订单", c)
}

func (od OrderApi) Complete(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("id"))
	err := od.service.Complete(orderId)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.EditError+"完成订单", c)
		return
	}
	res.OkWithMessage(code.EditSuccess+"完成订单", c)

}

func (od OrderApi) Search(c *gin.Context) {
	var condition req.OrderPageQueryDTO
	if err := c.ShouldBindQuery(&condition); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	search, err := od.service.Search(condition)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	res.OkWithData(*search, c)
}

func (od OrderApi) Statistics(c *gin.Context) {
	statistics, err := od.service.Statistics()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.QueryError, c)
		return
	}
	res.OkWithData(*statistics, c)
}
