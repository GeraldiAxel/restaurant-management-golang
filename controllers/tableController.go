package controllers

import "github.com/gin-gonic/gin"

func GetAllTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all tables
	}
}

func GetTableById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get table by ID
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a table
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update a table
	}
}
