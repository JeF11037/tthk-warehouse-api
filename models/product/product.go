package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SKU         string `json:"sku,omitempty"`
	CategoryID  string
	Category    Category `gorm:"foreignKey:CategoryID" json:"category_id,omitempty"`
	InventoryID string
	Inventory   Inventory `gorm:"foreignKey:InventoryID" json:"inventory_id,omitempty"`
}

func CreateProduct(db *gorm.DB, Product *Product) (err error) {
	err = db.Create(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, Product *Product) (err error) {
	err = db.Save(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(db *gorm.DB, Product *Product, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProduct(db *gorm.DB, Product *Product, id string) (err error) {
	err = db.Where("id = ?", id).First(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProducts(db *gorm.DB, Product *[]Product) (err error) {
	err = db.Find(Product).Error
	if err != nil {
		return err
	}
	return nil
}
