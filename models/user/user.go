package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	AddressID string
	Address   Address `gorm:"foreignKey:AddressID" json:"address_id,omitempty"`
	PaymentID string
	Payment   Payment `gorm:"foreignKey:PaymentID" json:"payment_id,omitempty"`
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

func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).Delete(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, User *User, id string) (err error) {
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

func VerifyUser(db *gorm.DB, User *User) (err error) {
	err = db.Where("email = ? AND password = ?", User.Email, User.Password).First(User).Error
	if err != nil {
		return err
	}
	return nil
}
