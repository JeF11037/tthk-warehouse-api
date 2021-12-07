package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func CreateCategory(db *gorm.DB, Category *Category) (err error) {
	err = db.Create(Category).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(db *gorm.DB, Category *Category) (err error) {
	err = db.Save(Category).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(db *gorm.DB, Category *Category, id string) (err error) {
	err = db.Where("id = ?", id).Delete(Category).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(db *gorm.DB, Category *Category, id string) (err error) {
	err = db.Where("id = ?", id).First(Category).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategories(db *gorm.DB, Categories *[]Category) (err error) {
	err = db.Find(Categories).Error
	if err != nil {
		return err
	}
	return nil
}
