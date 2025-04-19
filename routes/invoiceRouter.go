package routes

import (
	"restaurant-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetAllInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoiceById())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
