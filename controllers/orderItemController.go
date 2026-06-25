package controller

import (
	
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get order items
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get a specific order item by ID
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new order item
	}	
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update a specific order item by ID
	}	
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get order items by order ID
	}
}


func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
		
}