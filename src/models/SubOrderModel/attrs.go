package SubOrderModel

import (
	"github.com/shopspring/decimal"
	"jtthink.base/src/models"
)

func WithOrderId(orderId int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*SubOrders).OrderId = orderId
	}
}

func WithGoodsName(goodsName int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*SubOrders).GoodsName = goodsName
	}
}

func WithGoodsPrices(price float64) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*SubOrders).GoodsPrices = decimal.NewFromFloat(price)
	}
}

func WithGoodsCount(goodsCount int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*SubOrders).GoodsCount = goodsCount
	}
}

func WithGoodsId(goodsId int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*SubOrders).GoodsId = goodsId
	}
}
