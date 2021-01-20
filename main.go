package main

import (
	"fmt"
	"jtthink.base/src/models"
	"time"
)

func main() {
	g := models.NewGoods(
		models.WithGoodsSalePrices(180),
		models.WithGoodsMarketPrices(200),
		models.WithGoodsCreateTime(time.Now()),
		models.WithGoodsGoodsComment("dfefef"),
		models.WithGoodsGoodsName(1),
	)
	fmt.Println(g)
}
