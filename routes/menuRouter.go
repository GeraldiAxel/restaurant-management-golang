package routes

import (
	"restaurant-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetAllMenu())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenuById())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
