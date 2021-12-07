package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID string
	User   User    `gorm:"foreignKey:UserID" json:"user_id,omitempty"`
	Total  float32 `json:"total,omitempty"`
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

func DeleteOrder(db *gorm.DB, Order *Order, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(db *gorm.DB, Order *Order, id string) (err error) {
	err = db.Where("id = ?", id).First(Order).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrders(db *gorm.DB, Order *[]Order) (err error) {
	err = db.Find(Order).Error
	if err != nil {
		return err
	}
	return nil
}
