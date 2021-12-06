package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	AddressLineFirst  string `json:"address_line_1,omitempty"`
	AddressLineSecond string `json:"address_line_2,omitempty"`
	City              string `json:"city,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
	Telephone         string `json:"telephone,omitempty"`
}

type Payment struct {
	gorm.Model
	PaymentType string `json:"payment_type,omitempty"`
	CardNumber  int32  `json:"card_number,omitempty"`
	Expiry      string `json:"expiry,omitempty"`
}

type User struct {
	gorm.Model
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	AddressID uint
	Address   Address `gorm:"foreignKey:AddressID" json:"address_id,omitempty"`
	PaymentID uint
	Payment   Payment `gorm:"foreignKey:PaymentID" json:"payment_id,omitempty"`
}

func CreateAddress(db *gorm.DB, Address *Address) (err error) {
	err = db.Create(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAddress(db *gorm.DB, Address *Address) (err error) {
	err = db.Save(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAddress(db *gorm.DB, Address *Address) (err error) {
	err = db.Delete(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAddress(db *gorm.DB, Address *Address, id uint32) (err error) {
	err = db.Where("id = ?", id).First(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAddresses(db *gorm.DB, Address *[]Address) (err error) {
	err = db.Find(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func CreatePayment(db *gorm.DB, Payment *Payment) (err error) {
	err = db.Create(Payment).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePayment(db *gorm.DB, Payment *Payment) (err error) {
	err = db.Save(Payment).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePayment(db *gorm.DB, Payment *Payment) (err error) {
	err = db.Delete(Payment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPayment(db *gorm.DB, Payment *Payment, id uint32) (err error) {
	err = db.Where("id = ?", id).First(Payment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPayments(db *gorm.DB, Payment *[]Payment) (err error) {
	err = db.Find(Payment).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	err = db.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, User *User) (err error) {
	err = db.Delete(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, User *User, id uint32) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}
