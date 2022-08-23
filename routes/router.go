package routes

import (
	v1 "gin-blog/api/v1"
	"gin-blog/middleware"
	"gin-blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	//自定义日志：日志切割，日志软链接
	//实现日志每天一个log文件，日志软链接最新日志，
	//example：
	//time="2022-06-22 20:47:45" level=info Agent="ApiPOST Runtime +https://www.apipost.cn" DataSize=3851 HostName=DESKTOP-SCTNE5E Ip=172.21.80.1 Method=GET Path="/api/v1/article/list/5?page_size=10&page_num=1" SpendTime="5 ms" status=200
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	//跨域配置
	r.Use(middleware.Cors())
	//公共路由
	pubRoute := r.Group("/api/v1")
	{
		pubRoute.GET("categorys", v1.CategoryMethod.GetCategoryList)
		pubRoute.GET("articles", v1.ArticleMethod.GetArticleList)
		pubRoute.GET("article/list/:id", v1.ArticleMethod.GetCategoryArticle)
		pubRoute.GET("article/info/:id", v1.GetArticleInfo)
		pubRoute.POST("user/add", v1.UserMethod.AddUser)
		pubRoute.POST("login", v1.Login)
		pubRoute.GET("getimgurls", v1.GetImageUrls)
	}
	//携带token路由
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.GET("users", v1.UserMethod.GetUserList)
		auth.PUT("user/:id", v1.UserMethod.EditUser)
		auth.DELETE("user/:id", v1.UserMethod.DeleteUser)
		auth.GET("user/:id", v1.GetUserInfo)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		auth.GET("category/:id", v1.GetCateInfo)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("upload", v1.Upload)
		auth.DELETE("deleteossfiles", v1.DeleteOssFiles)
	}

	r.Run(utils.HttpPort)
}
