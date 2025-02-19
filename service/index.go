package service

import (
	"ginchat/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex 索引页面
// @Tags example
// @Success 200 {string} hello, world
// @Router /index [get]
func GetIndex(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", common.NewSuccessResponse())
}

func Register(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", common.NewSuccessResponse())
}
