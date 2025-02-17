package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// DBConnection 数据库连接信息
type DBConnection struct {
	UserName     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DatabaseName string `mapstructure:"dbname"`
	Charset      string `mapstructure:"charset"`
	ParseTime    string `mapstructure:"parseTime"`
	Loc          string `mapstructure:"loc"`
}

func (dbC *DBConnection) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		dbC.UserName,
		dbC.Password,
		dbC.Host,
		dbC.Port,
		dbC.DatabaseName,
		dbC.Charset,
		dbC.ParseTime,
		dbC.Loc)
}

// GetDBConfig 获取数据库连接信息
func GetDBConfig() (dbConfig *DBConnection) {
	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
	}
	return
}
