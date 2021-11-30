package models

type ProductCategory struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	CreationTime     string `json:"created_at,omitempty"`
	ModificationTime string `json:"modified_at,omitempty"`
}
