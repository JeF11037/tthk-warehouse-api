package controllers

import (
	"net/http"
	model "warehouse/models/user"
	"warehouse/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUser() *UserRepository {
	db := modules.Connect()
	db.AutoMigrate(&model.Address{}, &model.Payment{}, &model.User{})
	return &UserRepository{Db: db}
}

func (repository *UserRepository) CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateAddress(repository.Db, &user.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = model.CreatePayment(repository.Db, &user.Payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = model.CreateUser(repository.Db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepository) UpdateUser(c *gin.Context) {
	var user model.User
	id, _ := c.Params.Get("id")
	err := model.GetUser(repository.Db, &user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = model.UpdateUser(repository.Db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepository) DeleteUser(c *gin.Context) {
	var user model.User
	id, _ := c.Params.Get("id")
	err := model.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = model.DeleteAddress(repository.Db, &user.Address, user.AddressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = model.DeletePayment(repository.Db, &user.Payment, user.PaymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepository) GetUser(c *gin.Context) {
	var user model.User
	id, _ := c.Params.Get("id")
	err := model.GetUser(repository.Db, &user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepository) GetUsers(c *gin.Context) {
	var user []model.User
	err := model.GetUsers(repository.Db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
