package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/Response"
	"github.com/hsuehly/presideService/service"
)

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	userInfo, err := service.GetUserByIdService(userId)
	if err != nil {
		Response.Error(c, 500, "查询失败")
		return
	}

	Response.Success(c, userInfo)

}
