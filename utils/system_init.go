package utils

import (
	"fmt"
	"ginchat/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// InitConfig 初始化项目配置
func InitConfig() {
	// 设置配置文件名和路径
	viper.SetConfigName("app")
	viper.AddConfigPath(".")

	// 设置配置文件类型
	viper.SetConfigType("yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("yaml文件读取失败!")
	}
}

// InitDB 初始化数据库
func InitDB() {
	dnC := config.GetDBConfig()
	DB, _ = gorm.Open(mysql.Open(dnC.String()), &gorm.Config{})
	fmt.Println("MySQl init ...")
}
