package req

import (
	times "hmshop/common/time"
)

// OrderCancelDTO 商家取消订单接收数据模型
type OrderCancelDTO struct {
	OrderId      int    `json:"id"`
	CancelReason string `json:"cancelReason"`
}

// OrderConfirmDTO 接单接收数据模型
type OrderConfirmDTO struct {
	OrderId any `json:"id"`
	Status  int `json:"status"`
}

// OrderPageQueryDTO 订单分页查询数据模型
type OrderPageQueryDTO struct {
	Page      int    `form:"page" binding:"required"`
	PageSize  int    `form:"pageSize" binding:"required"`
	UserId    int    `form:"userId"`
	Number    string `form:"number"`
	Phone     string `form:"phone"`
	Status    int    `form:"status"`
	BeginTime string `form:"beginTime"`
	EndTime   string `form:"endTime"`
}

// OrderPaymentDTO 订单支付传递数据模型
type OrderPaymentDTO struct {
	OrderNumber string `json:"orderNumber"`
	PayMethod   int    `json:"payMethod"`
}

// OrderRejectionDTO 拒单接收数据模型
type OrderRejectionDTO struct {
	OrderId         int    `json:"id"`
	RejectionReason string `json:"rejectionReason"`
}

// OrderSubmitDTO 用户下单接口参数
type OrderSubmitDTO struct {
	AddressBookId         int              `json:"addressBookId"`
	Amount                float64          `json:"amount"`
	DeliveryStatus        int              `json:"deliveryStatus"`
	EstimatedDeliveryTime times.CustomTime `json:"estimatedDeliveryTime"`
	PackAmount            float64          `json:"packAmount"`
	PayMethod             int              `json:"payMethod"`
	Remark                string           `json:"remark"`
	TablewareNumber       int              `json:"tablewareNumber"`
	TablewareStatus       int              `json:"tablewareStatus"`
}
