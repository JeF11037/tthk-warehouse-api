package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Discount struct {
	gorm.Model
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Percent     float32 `json:"percent,omitempty"`
	Active      string  `json:"active,omitempty"`
}

type Inventory struct {
	gorm.Model
	Price    float32  `json:"price,omitempty"`
	Quantity int32    `json:"quantity,omitempty"`
	Discount Discount `sql:"index" json:"discount_id,omitempty"`
}

type Product struct {
	gorm.Model
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	SKU         string    `json:"sku,omitempty"`
	Category    Category  `sql:"index" json:"category_id,omitempty"`
	Inventory   Inventory `sql:"index" json:"inventory_id,omitempty"`
}
