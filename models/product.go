package models

type Product struct {
	ID               int              `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Description      string           `json:"description,omitempty"`
	SKU              string           `json:"sku,omitempty"`
	Category         ProductCategory  `json:"category_id,omitempty"`
	Inventory        ProductInventory `json:"inventory_id,omitempty"`
	CreationTime     string           `json:"created_at,omitempty"`
	ModificationTime string           `json:"modified_at,omitempty"`
}
