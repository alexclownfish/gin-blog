package v1

import (
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var ArticleMethod articlemethod

type articlemethod struct {
}

//添加文章
func AddArticle(ctx *gin.Context) {
	var data model.Article
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		logger.Error("参数绑定失败", err)
	}

	code = model.CreateArticle(&data)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func (a *articlemethod) GetArticleList(ctx *gin.Context) {
	title := ctx.Query("title")
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
	//switch {
	//case pageSize >= 100:
	//	pageSize = 100
	//case pageSize <= 0:
	//	pageSize = 10
	//}
	if params.PageNum == 0 {
		params.PageNum = -1
	}
	if len(title) == 0 {
		data, code, total := model.ArticleMethod.GetArticleList(params.PageSize, params.PageNum)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	data, code, total := model.SearchArticle(title, params.PageSize, params.PageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(ctx *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code = model.EditArticle(id, &data)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"messgae": errmsg.GetErrMsg(code),
	})

}

//删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章信息
func GetArticleInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArticleInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

//查询分类下的所有文章
func (a *articlemethod) GetCategoryArticle(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Param("id"))
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
	data, code, total := model.ArticleMethod.GetCategoryArticle(cid, params.PageSize, params.PageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
