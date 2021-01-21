package main

import (
	"fmt"
	"github.com/shenyisyn/goft-gin/goft"
	"jtthink.base/src/common"
	"jtthink.base/src/controller"
	"jtthink.base/src/models/GoodsModel"
)

func main() {
	common.InitDB()
	goft.Ignite().
		Mount("v1", controller.NewOrdercontroller()).
		Launch()
}

func test() {
	g := GoodsModel.NewGoods().Mutate(
		GoodsModel.WithGoodsName(1),
		GoodsModel.WithMarketPrices(1800),
	)
	fmt.Println(g)
}
