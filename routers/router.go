package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/handler"
)

func InitRouter(r *gin.Engine) {
	orderGroup := r.Group("/api/v1/order")
	{
		orderGroup.GET("/", handler.GetOrder)
		orderGroup.GET("/:id", handler.GetOrderById)
		orderGroup.POST("/create/:id", handler.CreateOrder)
		orderGroup.POST("/update", handler.UpdateOrderById)

	}
	userinfoGroup := r.Group("/api/v1/user")
	{
		userinfoGroup.GET("/:id", handler.GetUser)
	}
	foodGroup := r.Group("/api/v1/menu")
	{
		foodGroup.GET("/:id", handler.GetFoodMenu)

	}

}
