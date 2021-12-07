package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	AddressLineFirst  string `json:"address_line_1,omitempty"`
	AddressLineSecond string `json:"address_line_2,omitempty"`
	City              string `json:"city,omitempty"`
	PostalCode        string `json:"postal_code,omitempty"`
	Country           string `json:"country,omitempty"`
	Telephone         string `json:"telephone,omitempty"`
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

func DeleteAddress(db *gorm.DB, Address *Address, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAddress(db *gorm.DB, Address *Address, id string) (err error) {
	err = db.Where("id = ?", id).First(Address).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAddresses(db *gorm.DB, Addresses *[]Address) (err error) {
	err = db.Find(Addresses).Error
	if err != nil {
		return err
	}
	return nil
}
