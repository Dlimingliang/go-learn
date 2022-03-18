package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Name  *string
	Price uint
}

type User struct {
	UserID uint   `gorm:"primarykey"`
	Name   string `gorm:"column:user_name;size:200"`
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

	//生成表结构
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})

	p := Product{
		Code: sql.NullString{
			String: "D42",
			Valid:  true,
		},
		Price: 200,
	}
	result := db.Create(&p)
	fmt.Println(p.ID)
	fmt.Println(result.RowsAffected)
	db.Create(&Product{
		Code: sql.NullString{
			String: "D41",
			Valid:  true,
		},
		Price: 100,
	})

	var product Product
	result = db.First(&product, 2) //主键查找
	result = db.First(&product, []int{1, 2, 3})
	//判断是否查询到
	//errors.Is(result.Error, gorm.ErrRecordNotFound)
	db.First(&Product{}, "code = ?", "D42") //条件查找
	//检索全部对象
	db.Find(&Product{})
	//条件查询
	//获取一条
	db.Where("code = ?", "D42").First(&product)
	db.Where(&Product{Price: 100}).First(&product)
	//获取全部匹配的记录
	var products []Product
	db.Where("code = ?", "D42").Find(&products)

	db.Model(&product).Update("Price", 1000)
	// updates仅更新非零值字段, 但是update会更新零值
	// 解决非零值的方法有俩种
	// 1. 使用指针
	// 2. 使用sql.Null***
	db.Model(&product).Updates(Product{Price: 100, Code: sql.NullString{
		String: "",
		Valid:  true,
	}})
	//db.Model(&product).Updates(map[string]interface{}{"Price": 100, "Code": "F41"})

	db.Delete(&product)
}
