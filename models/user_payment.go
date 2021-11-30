package models

type UserPayment struct {
	ID               int    `json:"id,omitempty"`
	PaymentType      string `json:"payment_type,omitempty"`
	CardNumber       int    `json:"card_number,omitempty"`
	Expiry           string `json:"expiry,omitempty"`
	CreationTime     string `json:"created_at,omitempty"`
	ModificationTime string `json:"modified_at,omitempty"`
}
