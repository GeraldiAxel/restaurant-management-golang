package routes

import (
	"restaurant-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetAllOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItemById())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemByOrderId())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
