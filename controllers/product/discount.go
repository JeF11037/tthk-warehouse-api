package controllers

import (
	"net/http"
	model "warehouse/models/product"
	"warehouse/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DiscountRepository struct {
	Db *gorm.DB
}

func NewDiscount() *DiscountRepository {
	db := modules.Connect()
	db.AutoMigrate(&model.Discount{})
	return &DiscountRepository{Db: db}
}

func (repository *DiscountRepository) CreateDiscount(c *gin.Context) {
	var discount model.Discount
	c.BindJSON(&discount)
	err := model.CreateDiscount(repository.Db, &discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *DiscountRepository) UpdateDiscount(c *gin.Context) {
	var discount model.Discount
	id, _ := c.Params.Get("id")
	err := model.GetDiscount(repository.Db, &discount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.BindJSON(&discount)
	err = model.UpdateDiscount(repository.Db, &discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *DiscountRepository) DeleteDiscount(c *gin.Context) {
	var discount model.Discount
	id, _ := c.Params.Get("id")
	err := model.DeleteDiscount(repository.Db, &discount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *DiscountRepository) GetDiscount(c *gin.Context) {
	var discount model.Discount
	id, _ := c.Params.Get("id")
	err := model.GetDiscount(repository.Db, &discount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "discount": discount})
}

func (repository *DiscountRepository) GetDiscounts(c *gin.Context) {
	var discounts []model.Discount
	err := model.GetDiscounts(repository.Db, &discounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "discounts": discounts})
}
