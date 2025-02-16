package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserBasicList 获取用户列表
// @Tags example
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserBasicList [get]
func GetUserBasicList(context *gin.Context) {
	userBasicList := models.GetUserBasicList()
	context.JSON(200, gin.H{
		"msg":  "success",
		"data": userBasicList,
	})
}
