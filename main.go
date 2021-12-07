package main

import (
	userEndpoints "warehouse/routes/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	userEndpoints.Activate(router)
	router.Run()
}
