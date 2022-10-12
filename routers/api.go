package routers

import (
	"Assignment_02/controllers"
	"Assignment_02/database"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	db := database.StartDB()

	controller := controllers.New(db)

	router := gin.Default()

	router.GET("/orders", controller.GetOrders)
	router.POST("/order", controller.CreateOrder)
	router.PUT("/order/:id", controller.UpdateOrder)
	router.DELETE("/order/:id", controller.DeleteOrder)

	router.Run()

}
