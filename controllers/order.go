package controllers

import (
	"net/http"
	model "warehouse/models"
	productModel "warehouse/models/product"
	"warehouse/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrder() *OrderRepository {
	db := modules.Connect()
	db.AutoMigrate(&model.Order{}, &productModel.Product{})
	return &OrderRepository{Db: db}
}

func (repository *OrderRepository) ChangeAmountOfOrder(c *gin.Context) {
	var order model.Order
	change, _ := c.Params.Get("change")
	c.BindJSON(&order)
	if repository.Db.Model(&order).Where("user_id = ? AND product_id = ?", order.UserID, order.ProductID).Updates(&order).RowsAffected == 0 {
		order.Total = 1
		err := model.CreateOrder(repository.Db, &order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "order": order})
		return
	}
	if change == "1" {
		err := model.GetOrder(repository.Db, &order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
			return
		}
		order.Total += 1
		err = model.UpdateOrder(repository.Db, &order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
			return
		}
	}
	if change == "0" {
		err := model.GetOrder(repository.Db, &order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
			return
		}
		order.Total -= 1
		if order.Total > 0 {
			err = model.UpdateOrder(repository.Db, &order)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
				return
			}
		} else {
			err = model.DeleteOrder(repository.Db, &order)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "order": order})
}

func (repository *OrderRepository) GetOrders(c *gin.Context) {
	var orders []model.Order
	id, _ := c.Params.Get("id")
	err := model.GetOrdersByUser(repository.Db, &orders, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}
	for i := 0; i < len(orders); i++ {
		err = productModel.GetProduct(repository.Db, &orders[i].Product, orders[i].ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "orders": orders})
}
