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
	dbConfig := config.GetDBConfig()
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
	db.First(user, 1)
	fmt.Println(user) // 根据整型主键查找

	db.Model(user).Update("Name", "张三")
}
