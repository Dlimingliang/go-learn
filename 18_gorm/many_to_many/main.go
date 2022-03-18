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

type User3 struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
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

	//db.AutoMigrate(&User3{})
	//
	//var languages []Language
	//languages = append(languages, Language{Name: "go"})
	//languages = append(languages, Language{Name: "java"})
	//user := User3{
	//	Name:      "lml",
	//	Languages: languages,
	//}
	//db.Create(&user)

	//var user User3
	//db.Preload("Languages").First(&user)
	//fmt.Println(user)

	var user User3
	db.First(&user)
	var languages []Language
	db.Model(&user).Association("Languages").Find(&languages)
	fmt.Println(languages)

}
