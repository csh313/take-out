package userController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/service/userService"
	"strconv"
)

type OrderApi struct {
	service userService.OrderService
}

func (oa OrderApi) Pay(c *gin.Context) {
	var orderReq req.OrderPaymentDTO
	if err := c.ShouldBind(&orderReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	paymentVo, err := oa.service.Pay(orderReq, c)
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithData(paymentVo, c)

}

func (oa OrderApi) Submit(c *gin.Context) {
	var orderReq req.OrderSubmitDTO
	if err := c.ShouldBind(&orderReq); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	fmt.Println(orderReq.EstimatedDeliveryTime, "----------------")
	submitVO, err := oa.service.Submit(orderReq, c)
	if err != nil {
		return
	}
	res.OkWithData(*submitVO, c)
}

func (oa OrderApi) OrderDetail(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	detail, err := oa.service.OrderDetail(orderId, c)
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithData(*detail, c)
}

func (oa OrderApi) HistoryOrders(c *gin.Context) {
	var page req.OrderPageQueryDTO
	if err := c.ShouldBind(&page); err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	history, err := oa.service.History(page, c)
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithData(*history, c)

}

func (oa OrderApi) CancelOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err = oa.service.CancelOrder(orderId, c)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("取消订单成功", c)
}

func (oa OrderApi) RepeatOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err = oa.service.RepeatOrder(orderId, c)
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithMessage("再来一单成功", c)
}

func (oa OrderApi) RemindOrder(c *gin.Context) {

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(code.ReqError, c)
		return
	}
	err = oa.service.Remind(orderId, c)
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithMessage("提醒商家成功", c)
}
