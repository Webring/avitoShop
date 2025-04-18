package services

import (
	"AvitoShop/internal/models"
	"errors"
	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, userID uint) (*models.User, error) {
	var user models.User

	if err := db.First(&user, userID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
