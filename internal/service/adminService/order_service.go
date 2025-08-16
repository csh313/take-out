package adminService

import (
	"errors"
	"hmshop/common/code"
	"hmshop/common/res"
	"hmshop/global"
	"hmshop/internal/api/req"
	"hmshop/internal/api/resp"
	"hmshop/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderService struct {
}

func (s OrderService) Confirm(id int, c *gin.Context) error {

	var order = model.Order{
		Id:     id,
		Status: 3,
	}
	err := s.UpdateOrder(order)
	return err
}

func (s OrderService) UpdateOrder(order model.Order) error {
	if err := global.DBs.Model(&order).Updates(order).Error; err != nil {
		global.Log.Error(err.Error())
		return errors.New(code.EditError)
	}
	return nil

}

func (s OrderService) Details(id int) (*resp.OrderVO, error) {
	tx := global.DBs.Begin()

	//抛出异常
	defer func() {
		if r := recover(); r != nil {
			//遇到错误时回滚
			tx.Rollback()
		}
	}()

	var order = model.Order{}
	if err := tx.Where("id = ?", id).First(&order).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return nil, errors.New(code.QueryError)
	}
	var orderDetails []model.OrderDetail
	if err := tx.Where("order_id= ?", id).Find(&orderDetails).Error; err != nil {
		global.Log.Error(err.Error())
		tx.Rollback()
		return nil, errors.New(code.QueryError)
	}
	tx.Commit()
	orderDishes := Str(orderDetails)
	return &resp.OrderVO{
		Order:           order,
		OrderDishes:     orderDishes,
		OrderDetailList: orderDetails,
	}, nil

}

func (s OrderService) Rejection(rejection req.OrderRejectionDTO) error {
	orderModel, err := s.GetById(rejection.OrderId)
	if err != nil {
		return err
	}
	if orderModel == nil || orderModel.Status != code.ToBeConfirmed {
		global.Log.Error(errors.New("订单状态错误"))
		return errors.New("订单状态错误")
	}
	if orderModel.PayStatus == code.Paid {
		//用户已支付，拒单后退款
		global.Log.Info("商家拒绝该订单，退款金额：%v元", orderModel.Amount)
	}
	orderModel.Status = code.Cancelled
	orderModel.RejectionReason = rejection.RejectionReason
	orderModel.CancelTime = time.Now()
	err = s.UpdateOrder(*orderModel)
	return err
}

func (s OrderService) CancelOrder(cancel req.OrderCancelDTO) error {
	orderModel, err := s.GetById(cancel.OrderId)
	if err != nil {
		return err
	}
	if orderModel == nil {
		return errors.New("订单不存在")
	}
	if orderModel.PayStatus == code.Paid {
		//用户已支付，需要退款
		global.Log.Info("已支付订单已取消，退款金额：%v元", orderModel.Amount)
	}
	orderModel.Status = code.Cancelled
	orderModel.CancelTime = time.Now()
	orderModel.CancelReason = cancel.CancelReason
	err = s.UpdateOrder(*orderModel)
	return err
}

func (s OrderService) Delivery(orderId int, c *gin.Context) error {
	orderModel, err := s.GetById(orderId)
	if err != nil {
		return err
	}
	if orderModel == nil || orderModel.Status != code.Confirmed {
		return errors.New("订单状态错误")
	}
	orderModel.Status = code.DeliveryInProgress
	err = s.UpdateOrder(*orderModel)
	return err
}

func (s OrderService) Complete(orderId int) error {
	orderModel, err := s.GetById(orderId)
	if err != nil {
		return err
	}
	if orderModel == nil || orderModel.Status != code.DeliveryInProgress {
		return errors.New("订单状态错误")
	}
	orderModel.Status = code.Completed
	orderModel.DeliveryTime = time.Now()
	err = s.UpdateOrder(*orderModel)
	return err

}

func (s OrderService) Search(condition req.OrderPageQueryDTO) (*res.PageResult, error) {
	req := req.PageInfo{
		Page:     condition.Page,
		PageSize: condition.PageSize,
		Status:   condition.Status,
	}
	var orderModel model.Order
	query, pageSize, offset := res.PageListRow(req, orderModel)
	if condition.UserId != 0 {
		query = query.Where("user_id = ?", condition.UserId)
	}
	if condition.Number != "" {
		query = query.Where("number = ?", condition.Number)
	}

	layout := "2006-01-02 15:04:05" // 时间格式
	if condition.BeginTime != "" {
		t1, err := time.Parse(layout, condition.BeginTime)
		if err != nil {
			global.Log.Error(err.Error())
			return nil, err
		}
		query = query.Where("order_time >= ?", t1)
	}
	if condition.EndTime != "" {
		t2, err := time.Parse(layout, condition.EndTime)
		if err != nil {
			global.Log.Error(err.Error())
			return nil, err
		}
		query = query.Where("order_time<=", t2)
	}
	var page res.PageResult
	var orders []model.Order
	var orderVos []resp.OrderVO
	query = query.Count(&page.Total)
	if err := query.Limit(pageSize).Offset(offset).Order("order_time desc").Find(&orders).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}
	if orders != nil && len(orders) > 0 {
		for _, order := range orders {
			details, err := s.Details(order.Id)
			if err != nil {
				return nil, err
			}
			orderVos = append(orderVos, *details)
		}
	}
	page.Rows = orderVos
	return &page, nil
}

func (s OrderService) GetById(orderId int) (*model.Order, error) {
	var orderModel = model.Order{}
	if err := global.DBs.Where("id = ?", orderId).Find(&orderModel).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, errors.New(code.QueryError)
	}
	return &orderModel, nil
}

func (s OrderService) Statistics() (*resp.OrderStatisticsVO, error) {
	var orderStatistics = resp.OrderStatisticsVO{}
	tx := global.DBs.Model(&model.Order{})
	if err := tx.Where("status=?", code.ToBeConfirmed).
		Count(&orderStatistics.ToBeConfirmed).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}
	if err := tx.Where("status=?", code.Confirmed).
		Count(&orderStatistics.Confirmed).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}
	if err := tx.Where("status=?", code.DeliveryInProgress).
		Count(&orderStatistics.DeliveryInProgress).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}
	return &orderStatistics, nil
}

func Str(dishDetails []model.OrderDetail) string {
	var orderDishes string
	for _, dishDetail := range dishDetails {
		orderDishes += dishDetail.Name + ":" + strconv.Itoa(dishDetail.Number) + ";"
	}
	return orderDishes
}
