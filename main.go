package main

import (
	"main/controllers"
	"main/modules"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	modules.Connect()
	router.POST("/login", controllers.Login)
	router.POST("/registration", controllers.Registration)
	router.Run()
}
