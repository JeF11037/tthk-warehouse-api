package models

type OrderItems struct {
	ID               int    `json:"id,omitempty"`
	Order            int    `json:"order_id,omitempty"`
	Product          int    `json:"product_id,omitempty"`
	CreationTime     string `json:"created_at,omitempty"`
	ModificationTime string `json:"modified_at,omitempty"`
}
