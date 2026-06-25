package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controller"
)

func OrderRoutes(router *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PUT("/orders/:order_id", controller.UpdateOrder())
}