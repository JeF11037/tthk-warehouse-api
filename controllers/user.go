package controllers

import (
	"net/http"
	"warehouse/models"
	"warehouse/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUser() *UserRepository {
	db := modules.Connect()
	db.AutoMigrate(&models.Address{}, &models.Payment{}, &models.User{})
	return &UserRepository{Db: db}
}

func (repository *UserRepository) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateAddress(repository.Db, &user.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.CreatePayment(repository.Db, &user.Payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.CreateUser(repository.Db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepository) UpdateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.UpdateAddress(repository.Db, &user.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.UpdatePayment(repository.Db, &user.Payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepository) DeleteUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.UpdateAddress(repository.Db, &user.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.UpdatePayment(repository.Db, &user.Payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
