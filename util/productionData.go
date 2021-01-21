package main

import (
	"fmt"
	"jtthink.base/src/common"
	"jtthink.base/src/models/OrderModel"
	"jtthink.base/src/models/SubOrderModel"
	"time"
)

func main() {
	productionOrderData(300)
}

func productionOrderData(acount int) {
	for i := 0; i < acount; i++ {
		time.Sleep(time.Second * 3)
		p, goodsname := getOrderAmount()
		goodsCount := getGoodCount()
		o := OrderModel.NewOrders(
			OrderModel.WithUpdateTime(time.Now()),
			OrderModel.WithCreateTime(time.Now()),
			OrderModel.WithOrderAmount(p*goodsCount),
			OrderModel.WithOrderNum(getOrderNum()),
			OrderModel.WithUserId(getUserId()),
			OrderModel.WithStatus(getGoodStatus()),
		)
		if err := common.NewDBconfig().GormDB().Table("orders").Create(o).Error; err != nil {
			panic(err)
		}

		so := SubOrderModel.NewSubOrders(
			SubOrderModel.WithOrderId(o.Id),
			SubOrderModel.WithGoodsName(goodsname),
			SubOrderModel.WithGoodsPrices(p),
			SubOrderModel.WithGoodsCount(int(goodsCount)),
			SubOrderModel.WithGoodsId(goodsname),
		)
		if err := common.NewDBconfig().GormDB().Table("suborders").Create(so).Error; err != nil {
			panic(err)
		}
	}
}

func getUserId() int {
	t := time.Now().UnixNano()
	return int(t % 26)
}

func getGoodStatus() int {
	t := time.Now().UnixNano()
	return int(t%3 + 1)
}

func getGoodCount() float64 {
	t := time.Now().UnixNano()
	switch t % 3 {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return 1
	}
}

func getOrderNum() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func getOrderAmount() (float64, int) {
	t := time.Now().UnixNano()
	switch t % 3 {
	case 0:
		return 240, 1
	case 1:
		return 560, 2
	case 2:
		return 1800, 3

	}
	return 0, 0
}
