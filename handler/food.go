package handler

import (
	"github.com/hsuehly/presideService/Response"
	"github.com/hsuehly/presideService/service"

	"github.com/gin-gonic/gin"
)

func GetFoodMenu(c *gin.Context) {
	id := c.Param("id")
	result := service.GetFoodByIdService(id)
	Response.Success(c, result)

}
