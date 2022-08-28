package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/Response"
	"github.com/hsuehly/presideService/service"
)

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	fmt.Println("userId", userId)
	userInfo := service.GetUserByIdService(userId)
	fmt.Println("userInfo", userInfo)
	//data, err := json.Marshal(userInfo)
	//if err != nil {
	//	fmt.Println(err, "err")
	//}
	//fmt.Println(data, "data")
	Response.Success(c, userInfo)
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "success",
	//	"data": userInfo,
	//})
}
