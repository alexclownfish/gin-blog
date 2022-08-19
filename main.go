package main

import (
	"gin-blog/model"
	"gin-blog/routes"
)

func main() {
	//引入数据库
	model.InitDb()

	routes.InitRouter()
}
