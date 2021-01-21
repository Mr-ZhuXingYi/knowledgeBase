package main

import (
	"fmt"
	"github.com/shenyisyn/goft-gin/goft"
	"jtthink.base/src/common"
	"jtthink.base/src/configuration"
	"jtthink.base/src/controller"
	"jtthink.base/src/models/GoodsModel"
)

func main() {
	goft.Ignite().Config(common.NewDBconfig(), configuration.NewServiceConfig()).
		Mount("v1",
			controller.NewOrderController(),
			controller.NewGoodsController(),
		).
		Launch()
}

func test() {
	g := GoodsModel.NewGoods().Mutate(
		GoodsModel.WithGoodsName(1),
		GoodsModel.WithMarketPrices(1800),
	)
	fmt.Println(g)
}
