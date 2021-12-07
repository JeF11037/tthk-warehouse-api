package models

import (
	models "warehouse/models/product"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	OrderID   string
	Order     Order `gorm:"foreignKey:OrderID" json:"order_id,omitempty"`
	ProductID string
	Product   models.Product `gorm:"foreignKey:ProductID" json:"product_id,omitempty"`
}

func CreateItem(db *gorm.DB, Item *Item) (err error) {
	err = db.Create(Item).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateItem(db *gorm.DB, Item *Item) (err error) {
	err = db.Save(Item).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteItem(db *gorm.DB, Item *Item, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Item).Error
	if err != nil {
		return err
	}
	return nil
}

func GetItem(db *gorm.DB, Item *Item, id string) (err error) {
	err = db.Where("id = ?", id).First(Item).Error
	if err != nil {
		return err
	}
	return nil
}

func GetItems(db *gorm.DB, Items *[]Item) (err error) {
	err = db.Find(Items).Error
	if err != nil {
		return err
	}
	return nil
}
