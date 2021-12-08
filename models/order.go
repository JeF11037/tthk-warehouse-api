package models

import (
	product "warehouse/models/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    string          `json:"user_id,omitempty"`
	ProductID string          `json:"product_id,omitempty"`
	Product   product.Product `gorm:"foreignKey:ProductID;references:id"  `
	Total     float32         `json:"total,omitempty"`
}

func CreateOrder(db *gorm.DB, Order *Order) (err error) {
	err = db.Create(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrder(db *gorm.DB, Order *Order) (err error) {
	err = db.Save(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(db *gorm.DB, Order *Order) (err error) {
	err = db.Where("user_id = ? AND product_id = ?", Order.UserID, Order.ProductID).Delete(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(db *gorm.DB, Order *Order) (err error) {
	err = db.Where("user_id = ? AND product_id = ?", Order.UserID, Order.ProductID).First(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrdersByUser(db *gorm.DB, Orders *[]Order, id string) (err error) {
	err = db.Where("user_id = ?", id).Find(Orders).Error
	if err != nil {
		return err
	}
	return nil
}
