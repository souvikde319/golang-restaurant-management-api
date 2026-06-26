package controllers

import (
	"github.com/gin-gonic/gin"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

var MenuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// Logic to get menus
		result, err := MenuCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching menus"})
			return
		}
		var allMenes []models.Menu
		if err  = result.All(context.Background(), &allMenes); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenes)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get a specific menu by ID
	}	
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new menu
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update a specific menu by ID
	}
}

