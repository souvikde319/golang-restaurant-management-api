package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controller"
)

func MenuRoutes(router *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PUT("/menus/:menu_id", controller.UpdateMenu())
}