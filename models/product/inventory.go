package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Price      float32 `json:"price,omitempty"`
	Quantity   int32   `json:"quantity,omitempty"`
	DiscountID string
	Discount   Discount `gorm:"foreignKey:DiscountID" json:"discount_id,omitempty"`
}

func CreateInventory(db *gorm.DB, Inventory *Inventory) (err error) {
	err = db.Create(Inventory).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateInventory(db *gorm.DB, Inventory *Inventory) (err error) {
	err = db.Save(Inventory).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteInventory(db *gorm.DB, Inventory *Inventory, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Inventory).Error
	if err != nil {
		return err
	}
	return nil
}

func GetInventory(db *gorm.DB, Inventory *Inventory, id string) (err error) {
	err = db.Where("id = ?", id).First(Inventory).Error
	if err != nil {
		return err
	}
	return nil
}

func GetInventories(db *gorm.DB, Inventories *[]Inventory) (err error) {
	err = db.Find(Inventories).Error
	if err != nil {
		return err
	}
	return nil
}
