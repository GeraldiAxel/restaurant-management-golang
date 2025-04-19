package routes

import (
	"restaurant-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetAllOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrderById())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
}
