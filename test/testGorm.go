package main

import (
	"fmt"
	"ginchat/config"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dbConfig := config.NewDBConnection("gin_chat", "gin_chat", "192.168.88.130", 3306, "gin_chat", "utf8mb4")
	db, err := gorm.Open(mysql.Open(dbConfig.String()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{
		Name:          "admin",
		PassWord:      "admin",
		Phone:         "123456789",
		Email:         "admin@admin.com",
		Identity:      "admin",
		ClientIp:      "127.0.0.1",
		ClientPort:    "8080",
		LoginTime:     0,
		HeartbeatTime: 0,
		LogOutTime:    0,
		IsLogOut:      false,
		DeviceInfo:    "pc",
	}
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // 根据整型主键查找
	//db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("Name", "张三")
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
