package controllers

import "github.com/gin-gonic/gin"

func GetAllOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all order items
	}
}

func GetOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get order item by ID
	}
}

func GetOrderItemByOrderId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get order items by order ID
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update order item
	}
}
