package models

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMessage(test *testing.T) {
	db, err := gorm.Open(mysql.Open("gin_chat:gin_chat@tcp(192.168.88.130:3306)/gin_chat?charset=utf8mb4&parseTime=1&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.Message{})
}
