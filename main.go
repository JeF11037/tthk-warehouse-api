package main

import (
	userEndpoints "warehouse/routes/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userEndpoints.Activate(router)
	router.Run()
}
