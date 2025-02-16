package utils

import (
	"fmt"
	"ginchat/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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
	// 创建日志记录器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	// 获取数据库配置
	dnC := config.GetDBConfig()
	// 使用GORM打开MySQL数据库连接，并应用自定义的日志配置
	DB, _ = gorm.Open(mysql.Open(dnC.String()), &gorm.Config{Logger: newLogger})
	// 打印数据库初始化信息
	fmt.Println("MySQl init ...")
}
