package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/routers"
)

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	r.Run() // 监听并在 0.0.0.0:8081 上启动服务
}
