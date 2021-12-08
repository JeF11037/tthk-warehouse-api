package controllers

import (
	"net/http"
	model "warehouse/models/product"
	"warehouse/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Db *gorm.DB
}

func NewCategory() *CategoryRepository {
	db := modules.Connect()
	db.AutoMigrate(&model.Category{})
	return &CategoryRepository{Db: db}
}

func (repository *CategoryRepository) CreateCategory(c *gin.Context) {
	var category model.Category
	c.BindJSON(&category)
	err := model.CreateCategory(repository.Db, &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *CategoryRepository) UpdateCategory(c *gin.Context) {
	var category model.Category
	id, _ := c.Params.Get("id")
	err := model.GetCategory(repository.Db, &category, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.BindJSON(&category)
	err = model.UpdateCategory(repository.Db, &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *CategoryRepository) DeleteCategory(c *gin.Context) {
	var category model.Category
	id, _ := c.Params.Get("id")
	err := model.DeleteCategory(repository.Db, &category, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (repository *CategoryRepository) GetCategory(c *gin.Context) {
	var category model.Category
	id, _ := c.Params.Get("id")
	err := model.GetCategory(repository.Db, &category, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "category": category})
}

func (repository *CategoryRepository) GetCategories(c *gin.Context) {
	var categories []model.Category
	err := model.GetCategories(repository.Db, &categories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "categories": categories})
}
