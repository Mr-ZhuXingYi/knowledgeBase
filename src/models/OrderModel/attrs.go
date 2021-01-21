package OrderModel

import (
	"github.com/shopspring/decimal"
	"jtthink.base/src/models"
	"time"
)

func WithUpdateTime(updateTime time.Time) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).UpdateTime = updateTime
	}
}

func WithCouponCode(couponCode string) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).CouponCode = couponCode
	}
}

func WithOrderAmount(price float64) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).OrderAmount = decimal.NewFromFloat(price)
	}
}

func WithStatus(status int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).Status = status
	}
}

func WithCreateTime(createTime time.Time) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).CreateTime = createTime
	}
}

func WithOrderNum(orderNum string) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).OrderNum = orderNum
	}
}

func WithUserId(UserId int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Orders).UserId = UserId
	}
}
