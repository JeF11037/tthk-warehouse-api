package main

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()
	router.Run()
}
