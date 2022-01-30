package main

import (
	"flag"
	"log"
	"net/http"

	"gowebapp/assets"
	"gowebapp/handler"
	"gowebapp/middle"
	"gowebapp/model"

	"github.com/gin-gonic/gin"
)

var (
	endpoint        = flag.String("endpoint", "0.0.0.0:1888", "web监听地址")
	maxReqPerSecond = flag.Int("maxreq", 1000, "每秒最多请求数")
	dsn             = flag.String("dsn", "unknown", "数据库链接信息")
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	flag.Parse()

	// 初始化web基础配置
	e := gin.Default()
	e.Use(middle.ServeCORS("Authorization"))
	e.Use(middle.ServeAssets("/assets/"))
	e.Use(middle.RateLimiter(*maxReqPerSecond))
	e.StaticFS("/assets/", http.FS(assets.FS))

	// 初始化web接口
	handler.Init(e)

	// 初始化数据模型
	if *dsn == "unknown" {
		log.Fatal("不合法的数据库链接信息")
	}
	err := model.Init(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	// 启动web服务
	e.Run(*endpoint)
}
