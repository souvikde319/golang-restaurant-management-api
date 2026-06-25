package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controller"
)

func TableRoutes(router *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PUT("/tables/:table_id", controller.UpdateTable())
}