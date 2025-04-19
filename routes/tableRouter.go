package routes

import (
	"restaurant-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetAllTable())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTableById())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}

//
