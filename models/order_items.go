package models

type OrderItems struct {
	ID               int          `json:"id,omitempty"`
	Order            OrderDetails `json:"order_id,omitempty"`
	Product          Product      `json:"product_id,omitempty"`
	CreationTime     string       `json:"created_at,omitempty"`
	ModificationTime string       `json:"modified_at,omitempty"`
}
