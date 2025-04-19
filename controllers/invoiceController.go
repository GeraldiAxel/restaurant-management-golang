package controllers

import "github.com/gin-gonic/gin"

func GetAllInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all invoices
	}
}

func GetInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get invoice by ID
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create invoice
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update invoice
	}
}
