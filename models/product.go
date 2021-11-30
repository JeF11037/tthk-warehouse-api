package models

type Product struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	SKU              string `json:"sku,omitempty"`
	Category         int    `json:"category_id,omitempty"`
	Inventory        int    `json:"inventory_id,omitempty"`
	CreationTime     string `json:"created_at,omitempty"`
	ModificationTime string `json:"modified_at,omitempty"`
}
