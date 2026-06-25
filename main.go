package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"golang-restaurant-management/database"
	"golang-restaurant-management/middleware"
	"golang-restaurant-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodColection *mongo.Collection = database.OpenCollection(database.Client, "food")


func main() {
	port := os.Getenv("PORT")

	if port =="" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.UserRoutes(router)
	// database.ConnectDB()

	router.Run(":" + port)
	
}