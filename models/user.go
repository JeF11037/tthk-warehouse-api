package models

type User struct {
	ID               int         `json:"id,omitempty"`
	Username         string      `json:"username,omitempty"`
	Password         string      `json:"password,omitempty"`
	FirstName        string      `json:"first_name,omitempty"`
	LastName         string      `json:"last_name,omitempty"`
	Email            string      `json:"email,omitempty"`
	Address          UserAddress `json:"address_id,omitempty"`
	Payment          UserPayment `json:"payment_id,omitempty"`
	RegistrationTime string      `json:"created_at,omitempty"`
	ModificationTime string      `json:"modified_at,omitempty"`
}
