package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/soft_delete"
)

type TestUser struct {
	ID        uint `gorm:"primarykey"`
	IsDelete  bool
	UserName  string
	AgeInt8   int8
	AgeInt16  int16
	AgeInt32  int32
	AgeInt64  int64
	AgeInt    int
	AgeUint8  uint8
	AgeUint16 uint16
	AgeUint32 uint32
	AgeUint64 uint64
	AgeUint   uint
}

type User struct {
	ID        uint                  `gorm:"primarykey"`
	UserName  string                `gorm:"column:user_name;type:varchar(20);not null;unique;index:index_user_name;comment:用户名"`
	Mobile    string                `gorm:"column:mobile;type:varchar(11);not null;unique;index:index_mobile;comment:电话"`
	Gender    uint8                 `gorm:"comment:性别 0:未知 1:男 2:女"`
	Birthday  *time.Time            `gorm:"column:birthday;type:datetime; comment:生日"`
	CreatedAt time.Time             `gorm:"column:create_time;type:datetime"`
	UpdatedAt int                   `gorm:"column:update_time"`
	IsDelete  bool                  `gorm:"not null;default:0;comment:0:enabled 1:disabled"`
	Delete    soft_delete.DeletedAt `gorm:"softDelete:flag"`
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
	dsn := fmt.Sprintf("root:123456!@tcp(%s:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local", DbIP)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	var result *gorm.DB

	//生成表结构
	db.AutoMigrate(&User{})

	//单独创建用户
	createUser := User{
		UserName: "lml0",
		Mobile:   "18888888880",
	}
	result = db.Create(&createUser)
	fmt.Println(createUser.ID)

	//批量添加用户
	createUserList := []User{{UserName: "lml1", Mobile: "18888888881"}, {UserName: "lml2", Mobile: "18888888882"}, {UserName: "lml3", Mobile: "18888888883"}}
	result = db.Create(&createUserList)
	for _, user := range createUserList {
		fmt.Println(user.ID)
	}

	//普通查询用户
	var user User
	result = db.First(&user)
	fmt.Println(user)

	user = User{}
	result = db.Take(&user)
	fmt.Println(user)

	user = User{}
	result = db.Last(&user)
	fmt.Println(user)

	user = User{}
	result = db.Limit(1).Find(&user)
	fmt.Println(user)

	user = User{}
	result = db.First(&user, 2)
	fmt.Println(user)

	var userList []User
	result = db.Find(&userList, []int{1, 2, 3, 5})
	for _, u := range userList {
		fmt.Println(u)
	}

	//where查询
	user = User{}
	result = db.Where(&User{UserName: "lml2", IsDelete: false}).First(&user)
	fmt.Println(user)

	user = User{}
	result = db.Where(map[string]interface{}{"user_name": "lml2", "is_delete": false}).First(&user)
	fmt.Println(user)

	//更新
	var updateUser User
	db.First(&updateUser)
	//save更新所有字段
	updateUser.UserName = "update_lml"
	db.Save(&updateUser)
	//update更新单个字段
	db.Model(&updateUser).Update("IsDelete", true)
	db.Model(&updateUser).Update("IsDelete", false)
	//更新多个字段 结构体更新，只会更新非零值字段
	db.Model(&updateUser).Updates(User{UserName: "update_lml_multiply", IsDelete: false})
	db.Model(&updateUser).Updates(map[string]interface{}{"UserName": "update_lml_map", "IsDelete": false})
	//批量更新
	db.Model(&User{}).Where("id IN ?", []int{1}).Updates(&User{UserName: "lml_batch", IsDelete: false})
	db.Model(&User{}).Where("id IN ?", []int{1}).
		Updates(map[string]interface{}{"UserName": "lml_batch", "IsDelete": false})

	//删除
	var deleteUser User
	db.Last(&deleteUser)

	//主键删除
	db.Delete(&deleteUser)
	//主键批量删除
	db.Delete(&User{}, []uint{1, 2, 3})
	//批量删除
	db.Where("user_name = ?", "jinzhu").Delete(&User{})

	fmt.Println(result)

}
