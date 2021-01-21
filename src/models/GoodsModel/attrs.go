package GoodsModel

import (
	"github.com/shopspring/decimal"
	"jtthink.base/src/models"
	"time"
)

func WithSalePrices(price float64) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Goods).SalePrices = decimal.NewFromFloat(price)
		m.(*Goods).setDiscount()
	}
}

func WithMarketPrices(price float64) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Goods).MarketPrices = decimal.NewFromFloat(price)
		m.(*Goods).setDiscount()
	}
}

func WithCreateTime(createTime time.Time) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Goods).CreateTime = createTime
	}
}

func WithGoodsComment(goodsComment string) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Goods).GoodsComment = goodsComment
	}
}

func WithGoodsName(goodsName int) models.ModelAttrFunc {
	return func(m models.Mode) {
		m.(*Goods).GoodsName = goodsName
	}
}
