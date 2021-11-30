package models

type ProductInventory struct {
	ID               int      `json:"id,omitempty"`
	Price            float32  `json:"price,omitempty"`
	Quantity         int      `json:"quantity,omitempty"`
	Discount         Discount `json:"discount_id,omitempty"`
	CreationTime     string   `json:"created_at,omitempty"`
	ModificationTime string   `json:"modified_at,omitempty"`
}
