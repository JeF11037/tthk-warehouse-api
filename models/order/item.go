package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemID    string
	Item      Item `gorm:"foreignKey:ItemID" json:"Item_id,omitempty"`
	ProductID string
	Product   Product `gorm:"foreignKey:ProductID" json:"product_id,omitempty"`
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

func GetItems(db *gorm.DB, Item *[]Item) (err error) {
	err = db.Find(Item).Error
	if err != nil {
		return err
	}
	return nil
}
