package main

import (
	"github.com/gin-gonic/gin"
	"shen.com/ginvue/common"
	"shen.com/ginvue/routers"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 初始化数据库
	db := common.InitDB()
	defer db.Close()
	// 注册总路由
	routers.TotalRouter(r)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	panic(r.Run())
}
