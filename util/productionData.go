package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	. "jtthink.base/src/models"
	"time"
)

func main() {
	productionOrderData(199)
}

func productionOrderData(acount int) {
	db, err := gorm.Open("mysql", "devuser:123~!@@tcp(39.105.28.235:3320)/tech?charset=utf8&parseTime=true&loc=Local&multiStatements=true")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	for i := 0; i < acount; i++ {
		time.Sleep(time.Second * 3)
		p, goodsname := getOrderAmount()
		goodsCount := getGoodCount()
		o := NewOrders(
			WithOrdersUpdateTime(time.Now()),
			WithOrdersCreateTime(time.Now()),
			WithOrdersOrderAmount(p*goodsCount),
			WithOrdersOrderNum(getOrderNum()),
			WithOrdersUserId(getUserId()),
			WithOrdersStatus(getGoodStatus()),
		)
		if err := db.Table("orders").Create(o).Error; err != nil {
			panic(err)
		}

		so := NewSubOrders(
			WithSubOrdersOrderId(o.Id),
			WithSubOrdersGoodsName(goodsname),
			WithSubOrdersGoodsPrices(p),
			WithSubOrdersGoodsCount(int(goodsCount)),
		)
		if err := db.Table("suborders").Create(so).Error; err != nil {
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
