package models

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

const OrdersTableName = "orders"

type Orders struct {
	Id          int             `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id" `
	OrderNum    string          `gorm:"column:order_num" json:"order_num"  binding:"required,max=500"`
	UserId      int             `gorm:"column:user_id" json:"user_id"  binding:"gt=0"`
	CreateTime  time.Time       `gorm:"column:create_time" json:"create_time"`
	Status      int             `gorm:"column:status" json:"status"  binding:"gte=1,lte=3"`
	OrderAmount decimal.Decimal `json:"order_amount" gorm:"column:order_amount"`
	CouponCode  string          `gorm:"column:coupon_code" json:"coupon_code"  binding:"min=0,max=500"`
	UpdateTime  time.Time       `gorm:"column:update_time" json:"update_time"`
}

func (this *Orders) String() string {
	return fmt.Sprintf("Id:%d\nOrderNum:%s\nUserId:%d\nCreateTime:%s\nStatus:%d\nOrderAmount:%s\nCouponCode:%s\nUpdateTime:%s",
		this.Id, this.OrderNum, this.UserId, this.CreateTime, this.Status, this.OrderAmount, this.CouponCode, this.UpdateTime)
}

func NewOrders(f ...ModelAttrFunc) *Orders {
	o := &Orders{}
	UserAttsFuncs(f).Apply(o)
	return o
}

func WithOrdersUpdateTime(updateTime time.Time) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).UpdateTime = updateTime
	}
}

func WithOrdersCouponCode(couponCode string) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).CouponCode = couponCode
	}
}

func WithOrdersOrderAmount(price float64) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).OrderAmount = decimal.NewFromFloat(price)
	}
}

func WithOrdersStatus(status int) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).Status = status
	}
}

func WithOrdersCreateTime(createTime time.Time) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).CreateTime = createTime
	}
}

func WithOrdersOrderNum(orderNum string) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).OrderNum = orderNum
	}
}

func WithOrdersUserId(UserId int) ModelAttrFunc {
	return func(m Mode) {
		m.(*Orders).UserId = UserId
	}
}
