package main

import (
	"main/controllers"
	"main/modules"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	modules.Connect()
	// test example
	router.GET("/users", controllers.FindUsers)
	modules.Disconnect()
	router.Run()
}
