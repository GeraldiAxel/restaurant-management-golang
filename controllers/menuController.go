package controllers

import "github.com/gin-gonic/gin"

func GetAllMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all menus
	}
}

func GetMenuById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get menu by ID
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a menu
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update a menu
	}
}
