package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Person struct {
	gorm.Model
	Name        string
	CerditCards []CerditCard
}

type CerditCard struct {
	gorm.Model
	Code     string
	PersonID uint
}

func main() {
	//日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	//建立数据库连接
	dsn := "root:123456!@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//
	////生成表结构
	//db.AutoMigrate(&Person{})
	//db.AutoMigrate(&CerditCard{})
	//
	//person := Person{
	//	Name: "lml",
	//}
	//db.Create(&person)
	//db.Create(&CerditCard{
	//	Code:     "12",
	//	PersonID: person.ID,
	//})
	//db.Create(&CerditCard{
	//	Code:     "34",
	//	PersonID: person.ID,
	//})
	var person Person
	db.Preload("CerditCards").First(&person)
	fmt.Println(person)
}
