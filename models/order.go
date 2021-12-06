package models

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	Order   Order   `sql:"index" json:"order_id,omitempty"`
	Product Product `sql:"index" json:"product_id,omitempty"`
}

type Order struct {
	gorm.Model
	User  int32   `json:"user_id,omitempty"`
	Total float32 `json:"total,omitempty"`
}
