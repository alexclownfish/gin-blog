package v1

import (
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"gin-blog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var UserMethod usermethod

type usermethod struct {
}

var code int

//查询用户是否存在
func UserExist(ctx *gin.Context) {
	//
}

//添加用户
func (u *usermethod) AddUser(ctx *gin.Context) {
	var data model.User
	var msg string
	if err := ctx.ShouldBindJSON(&data); err != nil {
		logger.Error("ShouldBind参数绑定失败，", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.UserMethod.CheckUser(data.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

//查询用户列表
func (u *usermethod) GetUserList(ctx *gin.Context) {

	params := new(struct {
		PageSize int `form:"page_size"`
		PageNum  int `form:"page_num"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind参数绑定失败，", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bind参数绑定失败",
		})
		return
	}
	if params.PageNum == 0 {
		params.PageNum = -1
	}

	data, total := model.UserMethod.GetUserList(params.PageSize, params.PageNum)

	code = errmsg.SUCCESS
	ctx.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//编辑用户
func (u *usermethod) EditUser(ctx *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code = model.UserMethod.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"messgae": errmsg.GetErrMsg(code),
	})

}

//删除用户
func (u *usermethod) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.UserMethod.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
