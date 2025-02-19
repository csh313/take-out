package model

import "time"

// Order 订单数据模型
type Order struct {
	Id            int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Number        string `json:"number"`
	Status        int    `json:"status"`
	UserId        int    `json:"userId"`
	AddressBookId int    `json:"addressBookId"`
	// 下单时间，以当前时间自动插入
	OrderTime             time.Time `json:"orderTime" gorm:"autoCreateTime"`
	CheckoutTime          time.Time `json:"checkoutTime" gorm:"default:NULL"`
	CancelTime            time.Time `json:"cancelTime" gorm:"default:NULL"`
	EstimatedDeliveryTime time.Time `json:"estimatedDeliveryTime" gorm:"default:NULL"`
	DeliveryTime          time.Time `json:"deliveryTime" gorm:"default:NULL"`
	PayMethod             int       `json:"payMethod"`
	PayStatus             int       `json:"payStatus"`
	Amount                float64   `json:"amount"`
	Remark                string    `json:"remark"`
	Username              string    `json:"username" gorm:"column:user_name"`
	Phone                 string    `json:"phone"`
	Address               string    `json:"address"`
	Consignee             string    `json:"consignee"`
	CancelReason          string    `json:"cancelReason"`
	RejectionReason       string    `json:"rejectionReason"`
	DeliveryStatus        int       `json:"deliveryStatus"`
	PackAmount            float64   `json:"packAmount"`
	TablewareNumber       int       `json:"tablewareNumber"`
	TablewareStatus       int       `json:"tablewareStatus"`
}

// TableName 指定表名
func (Order) TableName() string {
	return "orders"
}
