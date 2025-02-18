package config

import "github.com/spf13/viper"

type RedisConfig struct {
	Addr        string `mapstructure:"addr"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"db"`
	PoolSize    int    `mapstructure:"poolSize"`
	MinIdleConn int    `mapstructure:"minIdleConn"`
}

// GetRedisConfig 获取Redis连接信息
func GetRedisConfig() (redisConfig *RedisConfig) {
	if err := viper.UnmarshalKey("redis", &redisConfig); err != nil {
	}
	return
}
