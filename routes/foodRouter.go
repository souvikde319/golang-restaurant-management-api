package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controller"
)

func FoodRoutes(router *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	inccomingRoutes.PUT("/foods/:food_id", controller.UpdateFood())
}