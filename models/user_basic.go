package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogoutTime    *time.Time `gorm:"column:logout_time" json:"logout_time"`
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// GetUserBasicList 获取用户列表
func GetUserBasicList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

// CreateUser 创建用户
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

// DeleteUser 删除用户
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

// UpdateUser 修改用户
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
