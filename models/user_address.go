package models

type UserAddress struct {
	ID                int    `json:"id,omitempty"`
	AddressLineFirst  string `json:"address_line_1,omitempty"`
	AddressLineSecond string `json:"address_line_2,omitempty"`
	City              string `json:"city,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
	Telephone         string `json:"telephone,omitempty"`
	CreationTime      string `json:"created_at,omitempty"`
	ModificationTime  string `json:"modified_at,omitempty"`
}
