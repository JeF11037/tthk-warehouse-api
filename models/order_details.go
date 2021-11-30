package models

type OrderDetails struct {
	ID               int     `json:"id,omitempty"`
	User             User    `json:"user_id,omitempty"`
	Total            float32 `json:"total,omitempty"`
	CreationTime     string  `json:"created_at,omitempty"`
	ModificationTime string  `json:"modified_at,omitempty"`
}
