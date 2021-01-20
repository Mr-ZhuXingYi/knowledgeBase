package models

import (
	"fmt"
	"github.com/shopspring/decimal"
)

const SubOrdersTableName = "suborders"

type SubOrders struct {
	Id          int             `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	OrderId     int             `gorm:"column:order_id" json:"order_id"`
	GoodsName   int             `gorm:"column:goods_name" json:"goods_name"`
	GoodsPrices decimal.Decimal `gorm:"column:goods_prices" json:"goods_prices"`
	GoodsCount  int             `gorm:"column:goods_count" json:"goods_count"`
}

func (this *SubOrders) String() string {
	return fmt.Sprintf("Id:%d\nOrderId:%d\nGoodsName:%d\nGoodsPrices:%s\nGoodsCount:%d\n",
		this.Id, this.OrderId, this.GoodsName, this.GoodsPrices, this.GoodsCount)
}

func NewSubOrders(f ...ModelAttrFunc) *SubOrders {
	so := &SubOrders{}
	UserAttsFuncs(f).Apply(so)
	return so
}

func WithSubOrdersOrderId(orderId int) ModelAttrFunc {
	return func(m Mode) {
		m.(*SubOrders).OrderId = orderId
	}
}

func WithSubOrdersGoodsName(goodsName int) ModelAttrFunc {
	return func(m Mode) {
		m.(*SubOrders).GoodsName = goodsName
	}
}

func WithSubOrdersGoodsPrices(price float64) ModelAttrFunc {
	return func(m Mode) {
		m.(*SubOrders).GoodsPrices = decimal.NewFromFloat(price)
	}
}

func WithSubOrdersGoodsCount(goodsCount int) ModelAttrFunc {
	return func(m Mode) {
		m.(*SubOrders).GoodsCount = goodsCount
	}
}
