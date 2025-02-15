package config

import "fmt"

// DBConnection 数据库连接信息
type DBConnection struct {
	userName     string
	password     string
	host         string
	port         int
	databaseName string
	charset      string
}

func (dbC *DBConnection) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbC.userName, dbC.password, dbC.host, dbC.port, dbC.databaseName, dbC.charset)
}

// NewDBConnection 创建数据库配置信息
func NewDBConnection(userName, password, host string, port int, databaseName, charset string) *DBConnection {
	return &DBConnection{
		userName:     userName,
		password:     password,
		host:         host,
		port:         port,
		databaseName: databaseName,
		charset:      charset,
	}
}
