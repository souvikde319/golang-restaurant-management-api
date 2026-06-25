package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controller"
)


func InvoiceRoutes(router *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PUT("/invoices/:invoice_id", controller.UpdateInvoice())
}