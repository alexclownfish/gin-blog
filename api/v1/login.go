package v1

import (
	"gin-blog/middleware"
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录接口
func Login(ctx *gin.Context) {
	var data model.User
	var token string
	var code int
	ctx.ShouldBindJSON(&data)
	code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, _ = middleware.SetToken(data.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"loginName": data.Username,
		"status":    code,
		"message":   errmsg.GetErrMsg(code),
		"token":     token,
	})
}
