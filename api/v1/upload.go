package v1

import (
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")

	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
