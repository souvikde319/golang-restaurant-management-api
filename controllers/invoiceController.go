package controller

import (

)

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get invoices
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get a specific invoice by ID
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new invoice
	}	
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update a specific invoice by ID
	}
}