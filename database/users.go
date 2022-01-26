package database

import (
	"TodoGolang/models"

	"github.com/jinzhu/gorm"
)

func GetUserByID(id string, db *gorm.DB) (models.User, bool, error) {
	b := models.User{}

	query := db.Select("users.*")
	query = query.Group("users.id")
	err := query.Where("users.id = ?", id).First(&b).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return b, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return b, false, nil
	}
	return b, true, nil
}

func GetUsers(db *gorm.DB) ([]models.User, error) {
	users := []models.User{}
	query := db.Select("users.*").
		Group("users.id")
	if err := query.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func DeleteUser(id string, db *gorm.DB) error {
	var b models.User
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, b *models.User) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
