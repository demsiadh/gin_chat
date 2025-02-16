package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex 索引页面
func GetIndex(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "hello, world",
	})
}
