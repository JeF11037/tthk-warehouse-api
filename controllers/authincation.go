package controllers

import (
	"main/models"
	"main/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	var verified bool
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	verified, err = modules.VerifyUser(user)
	message := gin.H{"success": verified}
	c.JSON(http.StatusOK, message)
}

func Registration(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	err = modules.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	message := gin.H{"success": true}
	c.JSON(http.StatusOK, message)
}
