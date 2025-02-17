package service

import (
	"ginchat/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserBasicList 获取用户列表
// @Summary 查询用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserBasicList [get]
func GetUserBasicList(context *gin.Context) {
	userBasicList := models.GetUserBasicList()
	context.JSON(200, gin.H{
		"msg":  "success",
		"data": userBasicList,
	})
}

// CreateUser 创建用户
// @Summary 创建用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param rePassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(context *gin.Context) {
	user := models.UserBasic{}
	user.Name = context.Query("name")
	password := context.Query("password")
	rePassword := context.Query("rePassword")
	if password != rePassword {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致!",
		})
		return
	}
	user.PassWord = password
	models.CreateUser(user)
	context.JSON(http.StatusOK, gin.H{
		"message": "创建成功!",
	})
}

// DeleteUser 删除用户(逻辑删除)
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(context *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(context.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	context.JSON(http.StatusOK, gin.H{
		"message": "删除成功!",
	})
}

// UpdateUser 修改用户
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(context *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(context.PostForm("id"))
	user.ID = uint(id)
	user.Name = context.PostForm("name")
	user.PassWord = context.PostForm("password")
	user.Phone = context.PostForm("phone")
	user.Email = context.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "修改用户失败!(传入参数不规范)",
		})
		return
	}

	models.UpdateUser(user)
	context.JSON(http.StatusOK, gin.H{
		"message": "修改用户成功!",
	})
}
