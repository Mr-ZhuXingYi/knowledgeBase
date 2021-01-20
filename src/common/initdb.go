package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	DB *gorm.DB
)

func InitDB() {

	orm, err := gorm.Open(mysql.Open("devuser:123~!@@tcp(39.105.28.235:3320)/tech?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, _ := orm.DB()
	mysqlDB.SetConnMaxLifetime(30 * time.Second)
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	DB = orm
}
