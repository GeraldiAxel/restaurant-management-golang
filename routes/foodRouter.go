package routes

import (
	"restaurant-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetAllFood())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFoodById())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
}
