package models

import (
	"fmt"
	"ginchat/config"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserBasic(t *testing.T) {
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
		LoginTime:     time.Now(),
		LogoutTime:    time.Now(),
		HeartbeatTime: time.Now(),
	}
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // 根据整型主键查找

	db.Model(user).Update("Name", "张三")
}
