package v1

import (
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"sync"
)

func Upload(ctx *gin.Context) {
	from, err := ctx.MultipartForm()
	if err != nil {
		panic(err.Error())
	}
	var (
		url  string
		urls []string
		code int
		//rcode int
		//code  []int
		wg sync.WaitGroup
	)
	files := from.File["file"]

	for _, file := range files {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 10010,
				"msg":  err.Error(),
			})
			return
		}
		wg.Add(1)
		go func() {
			url, code = model.Upload(file, &wg)
		}()
		wg.Wait()
		urls = append(urls, url)
		//code = append(code, cod)
		//for _, v := range code {
		//	rcode = v
		//	if v != 200 {
		//		logger.Error("上传失败，code：" + string(v))
		//	}
		//}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     urls,
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
	total := len(data)
	ctx.JSON(http.StatusOK, gin.H{
		"message": errmsg.GetErrMsg(code),
		"status":  code,
		"data":    data,
		"total":   total,
	})
}

func DeleteQNOssFiles(ctx *gin.Context) {
	params := new(struct {
		Keys []string `json:"keys"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("Bind参数绑定失败，", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bind参数绑定失败",
		})
		return
	}
	code, err := model.DeleteQNFiles(params.Keys)
	if err != nil {
		logger.Error("删除失败：%s", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"code":    code,
	})

}
