package v1

import (
	"gin-blog/model"
	"gin-blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var CategoryMethod categorymethod

type categorymethod struct {
}

//查询分类是否存在
//func UserExist(ctx *gin.Context) {
//	//
//}

//添加分类
func AddCategory(ctx *gin.Context) {
	var data model.Category
	_ = ctx.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)

	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询分类列表
func (c *categorymethod) GetCategoryList(ctx *gin.Context) {

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

	data, total := model.CategoryMethod.GetCategoryList(params.PageSize, params.PageNum)

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

//编辑分类
func EditCategory(ctx *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"messgae": errmsg.GetErrMsg(code),
	})

}

//删除分类
func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code = model.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类信息
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}
