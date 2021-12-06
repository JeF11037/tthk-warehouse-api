package main

import (
	"warehouse/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userRepository := controllers.NewUser()
	router.POST("/users", userRepository.CreateUser)
	router.PUT("/users", userRepository.CreateUser)
	router.DELETE("/users", userRepository.CreateUser)
	router.GET("/users", userRepository.CreateUser)
	router.GET("/users/:id", userRepository.CreateUser)
	router.Run()
}
