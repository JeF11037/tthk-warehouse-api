package models

type Discount struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	Percent          int    `json:"percent,omitempty"`
	Active           string `json:"active,omitempty"`
	CreationTime     string `json:"created_at,omitempty"`
	ModificationTime string `json:"modified_at,omitempty"`
}
