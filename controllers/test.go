package controllers

import (
	"main/models"
	"main/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []models.User
	modules.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
