package service

import (
	"fmt"
	"ginchat/common"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
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
	context.JSON(200, common.NewSuccessResponseWithData(userBasicList))
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
	// 判断两次密码是否一致
	if password != rePassword {
		context.JSON(http.StatusBadRequest, common.NewErrorResponse("两次密码不一致!"))
		return
	}

	// 判断用户名是否被注册
	if data := models.FindUserByName(user.Name); data.Name != "" {
		context.JSON(http.StatusBadRequest, common.NewErrorResponse("用户名已存在!"))
		return
	}

	// 加密密码
	user.Salt = fmt.Sprintf("%06d", rand.Int31()%1000000)
	user.PassWord = utils.MakePassword(password, user.Salt)
	models.CreateUser(user)
	context.JSON(http.StatusOK, common.NewSuccessResponseWithData(user))
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
	context.JSON(http.StatusOK, common.NewSuccessResponseWithData(user))
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
		context.JSON(http.StatusBadRequest, common.NewErrorResponse("修改用户失败!(传入参数不规范)"))
		return
	}

	models.UpdateUser(user)
	context.JSON(http.StatusOK, common.NewSuccessResponseWithData(user))
}

// LoginUser 登录用户
// @Summary 登录用户
// @Tags 用户模块
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code", "message"}
// @Router /user/loginUser [post]
func LoginUser(context *gin.Context) {
	name := context.PostForm("name")
	password := context.PostForm("password")

	// 查询用户记录
	user := models.FindUserByName(name)
	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, common.NewErrorResponse("用户不存在"))
		return
	}

	// 校验用户密码
	if !utils.ValidPassword(password, user.Salt, user.PassWord) {
		context.JSON(http.StatusBadRequest, common.NewErrorResponse("用户名或密码错误!"))
		return
	}

	// 发放token
	user.Identity = utils.MakeToken()
	// 修改数据库唯一标识
	models.UpdateUser(user)

	context.JSON(http.StatusOK, common.NewSuccessResponseWithData(user))
}
