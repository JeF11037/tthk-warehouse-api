package routes

import (
	"warehouse/controllers"

	"github.com/gin-gonic/gin"
)

func Activate(router *gin.Engine) {
	userRepository := controllers.NewUser()
	router.POST("/users", userRepository.CreateUser)
	router.PUT("/users/:id", userRepository.UpdateUser)
	router.DELETE("/users/:id", userRepository.DeleteUser)
	router.GET("/users", userRepository.GetUsers)
	router.GET("/users/:id", userRepository.GetUser)
}
