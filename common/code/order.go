package code

const (
	Disable = 0
	Enable  = 1
)

// 订单状态 1待付款 2待接单 3已接单 4派送中 5已完成 6已取消 7退款
const (
	PendingPayment = 1 + iota
	ToBeConfirmed
	Confirmed
	DeliveryInProgress
	Completed
	Cancelled
	Refunds
)

// 支付方式
const (
	Wechat = 1
	AliPay = 2
)

// 支付状态
const (
	Unpaid = iota
	Paid
	Refund
)
