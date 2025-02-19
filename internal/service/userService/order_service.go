package userService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iWyh2/go-myUtils/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hmshop/common/code"
	"hmshop/common/enum"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/api/resp"
	"hmshop/internal/model"
	"hmshop/internal/service"
	"hmshop/internal/service/adminService"
	"hmshop/internal/service/kafkaService"
	"strconv"
	"time"
)

var admin adminService.OrderService

type OrderService struct{}

func (s OrderService) Submit(req req.OrderSubmitDTO, c *gin.Context) (*resp.OrderSubmitVO, error) {
	//查询收货地址是否为空
	address, err := AddressService{}.GetById(req.AddressBookId, c)
	if address == nil {
		res.FailWithMessage(code.DataNotFound, c)
		return nil, errors.New(code.DataNotFound)
	}
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		res.FailWithMessage(code.UserError, c)
		return nil, errors.New(code.UserError)
	}
	userId := value.(uint64)
	cart, err := ShoppingCartService{}.QueryShoppingCart(model.ShoppingCart{UserId: userId}, c)
	if err != nil {
		return nil, err
	}
	//查看当前用户的购物车
	if cart == nil || len(cart) == 0 {
		res.FailWithMessage(code.DataNotFound, c)
		return nil, errors.New(code.DataNotFound)
	}
	//构造订单数据
	var order model.Order
	err = copier.Copy(&order, &req)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return nil, err
	}
	order.UserId = int(userId)
	order.Number = strconv.FormatInt(time.Now().UnixMilli(), 10)
	order.Phone = address.Phone
	order.Address = address.Detail
	order.Consignee = address.Consignee
	order.Status = code.PendingPayment
	order.PayStatus = code.Unpaid
	order.OrderTime = time.Now()
	//var defaultTime time.Time
	//order.CheckoutTime = time.Now()
	//order.CancelTime = time.Now()
	//order.DeliveryTime = time.Now()

	//插入订单数据
	if err = global.DB.Create(&order).Error; err != nil {
		res.FailWithMessage(code.AddError, c)
		return nil, err
	}
	//插入订单明细表
	var orderDetailList []model.OrderDetail
	for _, shopCart := range cart {
		var orderDetail model.OrderDetail
		if err := copier.Copy(&orderDetail, &shopCart); err != nil {
			res.FailWithMessage(code.ServerInternalError, c)
			return nil, err
		}
		orderDetail.OrderId = order.Id
		orderDetailList = append(orderDetailList, orderDetail)
	}
	if err = global.DB.Create(&orderDetailList).Error; err != nil {
		res.FailWithMessage(code.AddError, c)
		return nil, err
	}
	//清空购物车数据
	err = ShoppingCartService{}.Clean(c)
	if err != nil {
		return nil, err
	}
	return &resp.OrderSubmitVO{OrderId: order.Id,
		OrderNumber: order.Number,
		OrderAmount: order.Amount,
		OrderTime:   order.OrderTime}, nil

}

func (s OrderService) Pay(orderReq req.OrderPaymentDTO, c *gin.Context) (*resp.OrderPaymentVO, error) {
	var order model.Order
	//先查看订单状态，若为待付款，则可进行下一步
	if err := global.DB.Where("number=?", orderReq.OrderNumber).Take(&order).Error; err != nil {
		global.Log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.FailWithMessage(code.DataNotFound, c)
		} else {
			res.FailWithMessage(err.Error(), c)
		}
		return nil, err
	}
	if order.Status != 1 {
		res.FailWithMessage("该订单不是待付款状态", c)
		return nil, errors.New("该订单不是待付款状态")
	}
	//todo 调用微信支付接口，此处手动生成
	go s.PaySuccess(orderReq.OrderNumber, c)
	// 构建消息内容
	message := fmt.Sprintf("OrderNumber %d has been paid successfully", orderReq.OrderNumber)
	// 初始化 Kafka 生产者
	err := kafkaService.SendMessage(global.Producer, message)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return nil, err
	}

	return &resp.OrderPaymentVO{
		NonceStr:   iUtils.UUID(),
		PaySign:    "iWyh2",
		SignType:   "iwyh2",
		PackageStr: iUtils.UUID(),
		TimeStamp:  strconv.FormatInt(time.Now().UnixMilli(), 10),
	}, nil
}
func (s OrderService) PaySuccess(orderId string, c *gin.Context) {
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		res.FailWithMessage(code.UserError, c)
		return
	}
	userId := value.(uint64)
	orderid, err := strconv.Atoi(orderId)
	if err != nil {
		res.FailWithMessage(code.ServerInternalError, c)
		return
	}
	if err = global.DB.Model(&model.Order{}).Where("id=?", orderid).Updates(&model.Order{UserId: int(userId),
		PayStatus: code.Paid, Status: code.ToBeConfirmed, CheckoutTime: time.Now()}).Error; err != nil {
		res.FailWithMessage(code.SqlError, c)
		return
	}

}

func (s OrderService) OrderDetail(id int, c *gin.Context) (*resp.OrderVO, error) {
	details, err := admin.Details(id)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return nil, err
	}
	return details, nil
}

func (s OrderService) History(pageReq req.OrderPageQueryDTO, c *gin.Context) (*res.PageResult, error) {
	list, err := admin.Search(pageReq)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return nil, err
	}
	return list, nil
}

func (s OrderService) CancelOrder(id int, c *gin.Context) error {
	err := admin.CancelOrder(req.OrderCancelDTO{OrderId: id,
		CancelReason: "用户取消订单"})
	return err
	//order, err := adminService.OrderService{}.GetById(id)
	//if err != nil {
	//	res.FailWithMessage(code.SqlError, c)
	//	return
	//}
	//if order == nil {
	//	res.FailWithMessage(code.QueryError, c)
	//	return
	//}
	//if order.Status>code.ToBeConfirmed{
	//	res.FailWithMessage("商家已接单，退款失败",c)
	//	return
	//}
	//if order.Status==code.ToBeConfirmed {
	//	order.PayStatus=code.Refund
	//}
	//order.Status=code.Cancelled
	//order.CancelReason="用户取消订单"
	//order.CancelTime=time.Now()
	//

}

func (s OrderService) RepeatOrder(id int, c *gin.Context) error {
	//查询当前用户id
	value, exists := c.Get(enum.CurrentUserId)
	if !exists {
		res.FailWithMessage(code.UserError, c)
		return errors.New(code.UserError)
	}
	userId := value.(uint64)
	//根据订单id查询订单详情信息
	var orderDetailList []model.OrderDetail
	if err := global.DB.Where(&model.OrderDetail{OrderId: id}).Find(&orderDetailList).Error; err != nil {
		res.FailWithMessage(code.QueryError, c)
		return err
	}
	if len(orderDetailList) == 0 {
		res.FailWithMessage(code.DataNotFound, c)
		return errors.New(code.DataNotFound)
	}

	//将订单详细的数据再次放到购物车中
	var cartList []model.ShoppingCart
	for _, order := range orderDetailList {
		var shoppingCart model.ShoppingCart
		if err := copier.Copy(&shoppingCart, &order); err != nil {
			res.FailWithMessage(err.Error(), c)
			return err
		}
		shoppingCart.UserId = userId
		cartList = append(cartList, shoppingCart)
	}
	if err := global.DB.Create(&cartList).Error; err != nil {
		res.FailWithMessage(code.SqlError, c)
		return err
	}
	return nil
}

func (s OrderService) Remind(id int, c *gin.Context) error {
	order, err := admin.GetById(id)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return err
	}
	if order == nil {
		res.FailWithMessage(code.DataNotFound, c)
		return errors.New(code.DataNotFound)
	}
	jsonMap := map[string]any{
		"type":    2,
		"orderId": id,
		"content": "订单号" + order.Number,
	}
	service.WSServer.SendToAllClients(jsonMap)
	return nil
}
