package controllers

import (
	"net/http"
	model "warehouse/models/product"
	"warehouse/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func NewProduct() *ProductRepository {
	db := modules.Connect()
	db.AutoMigrate(&model.Product{})
	return &ProductRepository{Db: db}
}

func (repository *ProductRepository) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	err := model.CreateProduct(repository.Db, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *ProductRepository) UpdateProduct(c *gin.Context) {
	var product model.Product
	id, _ := c.Params.Get("id")
	err := model.GetProduct(repository.Db, &product, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.BindJSON(&product)
	err = model.UpdateProduct(repository.Db, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *ProductRepository) DeleteProduct(c *gin.Context) {
	var product model.Product
	id, _ := c.Params.Get("id")
	err := model.DeleteProduct(repository.Db, &product, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *ProductRepository) GetProduct(c *gin.Context) {
	var product model.Product
	id, _ := c.Params.Get("id")
	err := model.GetProduct(repository.Db, &product, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "product": product})
}

func (repository *ProductRepository) GetProducts(c *gin.Context) {
	var products []model.Product
	err := model.GetProducts(repository.Db, &products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "products": products})
}
