package common

import (
	"github.com/mattn/go-colorable"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DBconfig struct{}

func NewDBconfig() *DBconfig {
	return &DBconfig{}
}

func (this *DBconfig) GormDB() *gorm.DB {
	newLogger := logger.New(
		log.New(colorable.NewColorableStdout(), "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)
	db, err := gorm.Open(mysql.Open("devuser:123~!@@tcp(39.105.28.235:3320)/tech?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, _ := db.DB()
	mysqlDB.SetConnMaxLifetime(30 * time.Second)
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)

	return db
}
