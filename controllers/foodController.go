package controller

import (

)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get foods
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get a specific food by ID
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new food item
	}	
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update a specific food item by ID
	}	
}


func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {
	
}