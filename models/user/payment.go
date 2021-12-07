package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentType string `json:"payment_type,omitempty"`
	CardNumber  int32  `json:"card_number,omitempty"`
	Expiry      string `json:"expiry,omitempty"`
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

func DeletePayment(db *gorm.DB, Payment *Payment, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Payment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPayment(db *gorm.DB, Payment *Payment, id string) (err error) {
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
