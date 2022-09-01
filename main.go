package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/Response"
	"github.com/hsuehly/presideService/config"
	"github.com/hsuehly/presideService/middleware"
	"github.com/hsuehly/presideService/routers"
	"github.com/hsuehly/presideService/util"
)

func main() {
	config.InitConfigData()
	util.WxinitData()
	r := gin.Default()
	r.Use(middleware.CORS())
	err := Response.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	routers.InitRouter(r)
	//userInfo := service.GetUserByIdService("100")
	//fmt.Println(userInfo)
	r.Run() // 监听并在 0.0.0.0:8081 上启动服务
}
