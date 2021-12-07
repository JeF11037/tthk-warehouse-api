package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Percent     float32 `json:"percent,omitempty"`
	Active      string  `json:"active,omitempty"`
}

func CreateDiscount(db *gorm.DB, Discount *Discount) (err error) {
	err = db.Create(Discount).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateDiscount(db *gorm.DB, Discount *Discount) (err error) {
	err = db.Save(Discount).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDiscount(db *gorm.DB, Discount *Discount, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Discount).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDiscount(db *gorm.DB, Discount *Discount, id string) (err error) {
	err = db.Where("id = ?", id).First(Discount).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDiscounts(db *gorm.DB, Discounts *[]Discount) (err error) {
	err = db.Find(Discounts).Error
	if err != nil {
		return err
	}
	return nil
}
