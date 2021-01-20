package models

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

const GoodsTableName = "goods"

type Goods struct {
	Id           int             `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	GoodsName    int             `gorm:"column:goods_name" json:"goods_name"`
	GoodsComment string          `gorm:"column:goods_comment" json:"goods_comment"`
	CreateTime   time.Time       `gorm:"column:create_time" json:"create_time"`
	MarketPrices decimal.Decimal `gorm:"column:market_prices" json:"market_prices"`
	SalePrices   decimal.Decimal `gorm:"column:sale_prices" json:"sale_prices"`
	DisCount     decimal.Decimal `gorm:"column:dis_count" json:"dis_count"`
}

func (this *Goods) String() string {
	return fmt.Sprintf("Id:%d\nGoodsName:%d\nGoodsComment:%s\nCreateTime:%s\nMarketPrices:%s\nSalePrices:%s\nDisCount:%s\n",
		this.Id, this.GoodsName, this.GoodsComment, this.CreateTime, this.MarketPrices, this.MarketPrices, this.DisCount)
}

func (this *Goods) setDiscount() {
	zero := decimal.NewFromInt(0)
	if this.MarketPrices.GreaterThan(zero) && this.SalePrices.GreaterThan(zero) {
		this.DisCount = this.MarketPrices.Sub(this.SalePrices).
			DivRound(this.MarketPrices, 2).Sub(decimal.NewFromInt(1)).Abs()
	}
}

func NewGoods(f ...ModelAttrFunc) *Goods {
	g := &Goods{}
	UserAtts(f).Apply(g)
	return g
}

func WithGoodsSalePrices(price float64) ModelAttrFunc {
	return func(m Mode) {
		m.(*Goods).SalePrices = decimal.NewFromFloat(price)
		m.(*Goods).setDiscount()
	}
}

func WithGoodsMarketPrices(price float64) ModelAttrFunc {
	return func(m Mode) {
		m.(*Goods).MarketPrices = decimal.NewFromFloat(price)
		m.(*Goods).setDiscount()
	}
}

func WithGoodsCreateTime(createTime time.Time) ModelAttrFunc {
	return func(m Mode) {
		m.(*Goods).CreateTime = createTime
	}
}

func WithGoodsGoodsComment(goodsComment string) ModelAttrFunc {
	return func(m Mode) {
		m.(*Goods).GoodsComment = goodsComment
	}
}

func WithGoodsGoodsName(goodsName int) ModelAttrFunc {
	return func(m Mode) {
		m.(*Goods).GoodsName = goodsName
	}
}
