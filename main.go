package main

import (
	"fmt"
	"jtthink.base/src/common"
)

func main() {
	common.InitDB()
	var c int64
	if err := common.DB.Table("goods").Count(&c).Error; err != nil {
		panic(err)
	}
	fmt.Println(c)

}
