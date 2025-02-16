package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex 索引页面
// @Tags example
// @Success 200 {string} hello, world
// @Router /index [get]
func GetIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "hello, world",
	})
}
