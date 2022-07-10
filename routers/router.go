package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/handler"
)

func InitRouter(r *gin.Engine) {
	r.POST("/test", func(c *gin.Context) {
		test, ok := c.GetPostForm("test")
		if !ok {
			fmt.Println("err", ok, test)
		}
		c.JSON(200, gin.H{
			"test": test,
		})
	})
	orderGroup := r.Group("/api/v1/order")
	{
		orderGroup.GET("/", handler.GetOrder)
		orderGroup.GET("/:id", handler.GetOrderById)
		orderGroup.POST("/create", handler.CreateOrder)
		orderGroup.POST("/update", handler.UpdateOrderById)

	}

}
