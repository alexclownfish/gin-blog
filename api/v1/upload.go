package v1

import (
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

func Upload(ctx *gin.Context) {
	//kind := ctx.Param("kind")
	//_, file, _ := ctx.Request.FormFile("file")
	//fileSize := file.Size
	f, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 10010,
			"msg":  err.Error(),
		})
		return
	}
	url, code := model.Upload(f)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}

func GetImageUrls(ctx *gin.Context) {
	params := new(struct {
		Prefix    string `form:"prefix"`
		Delimiter string `form:"delimiter"`
		Marker    string `form:"marker"`
		Limit     int    `form:"limit"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("Bind参数绑定失败，", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bind参数绑定失败",
		})
		return
	}
	data, code, err := model.GetImages(params.Prefix, params.Delimiter, params.Marker, params.Limit)
	if err != nil {
		logger.Error("查询列表失败，", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": errmsg.GetErrMsg(code),
		"status":  code,
		"data":    data,
	})
}
