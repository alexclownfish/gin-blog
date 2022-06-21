package routes

import (
	v1 "gin-blog/api/v1"
	"gin-blog/middleware"
	"gin-blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	pubRoute := r.Group("/api/v1")
	{
		pubRoute.GET("users", v1.UserMethod.GetUserList)
		pubRoute.GET("categorys", v1.CategoryMethod.GetCategoryList)
		pubRoute.GET("articles", v1.ArticleMethod.GetArticleList)
		pubRoute.GET("article/list/:id", v1.ArticleMethod.GetCategoryArticle)
		pubRoute.GET("article/info/:id", v1.GetArticleInfo)
		pubRoute.POST("user/add", v1.UserMethod.AddUser)
		pubRoute.POST("login", v1.Login)
	}

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("user/:id", v1.UserMethod.EditUser)
		auth.DELETE("user/:id", v1.UserMethod.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

	}

	r.Run(utils.HttpPort)
}
