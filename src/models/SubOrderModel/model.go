package SubOrderModel

import (
	"fmt"
	"github.com/shopspring/decimal"
	"jtthink.base/src/models"
)

const SubOrdersTableName = "suborders"

type SubOrders struct {
	Id          int             `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	OrderId     int             `gorm:"column:order_id" json:"order_id" binding:"gt=0"`
	GoodsName   int             `gorm:"column:goods_name" json:"goods_name"  binding:"gte=1,lte=3"`
	GoodsPrices decimal.Decimal `gorm:"column:goods_prices" json:"goods_prices"`
	GoodsCount  int             `gorm:"column:goods_count" json:"goods_count" binding:"gt=0"`
}

func (this *SubOrders) String() string {
	return fmt.Sprintf("Id:%d\nOrderId:%d\nGoodsName:%d\nGoodsPrices:%s\nGoodsCount:%d\n",
		this.Id, this.OrderId, this.GoodsName, this.GoodsPrices, this.GoodsCount)
}

func NewSubOrders(f ...models.ModelAttrFunc) *SubOrders {
	so := &SubOrders{}
	models.UserAttsFuncs(f).Apply(so)
	return so
}

func (this *SubOrders) Mutate(f ...models.ModelAttrFunc) *SubOrders {
	models.UserAttsFuncs(f).Apply(this)
	return this
}
