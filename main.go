package main

import (
	"flag"
	"gowebapp/handler"
	"gowebapp/middle"
	"gowebapp/model"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	endpoint = flag.String("endpoint", "0.0.0.0:1888", "web监听地址")
)

func main() {
	flag.Parse()

	// 初始化web基础配置
	e := gin.Default()
	e.Use(middle.ServeCORS("Authorization"))
	e.Use(middle.ServeAssets("/assets/"))
	e.StaticFS("/assets/", Assets)

	// 初始化web接口
	handler.Init(e)

	// 初始化数据模型
	err := model.Init()
	if err != nil {
		log.Fatal(err)
	}

	// 启动web服务
	e.Run(*endpoint)
}
