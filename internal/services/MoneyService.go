package services

import (
	"AvitoShop/internal/models"
	"errors"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func SendMoney(db *gorm.DB, fromUserID, toUserID, value uint) error {
	if value == 0 || fromUserID == toUserID {
		log.Warnf("Useless send money operation: from %u to %u with value %u", fromUserID, toUserID, value)
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var fromUser, toUser models.User

		if err := tx.First(&fromUser, fromUserID).Error; err != nil {
			return errors.New("sender not found")
		}

		if err := tx.First(&toUser, toUserID).Error; err != nil {
			return errors.New("receiver not found")
		}

		if fromUser.Money < value {
			return errors.New("sender haven't enough money")
		}

		if err := tx.Create(&models.MoneyTransaction{
			FromUserID: fromUserID,
			ToUserID:   toUserID,
			Value:      value,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&fromUser).Update("Money", fromUser.Money-value).Error; err != nil {
			return err
		}

		if err := tx.Model(&toUser).Update("Money", toUser.Money+value).Error; err != nil {
			return err
		}

		return nil
	})
}

func SendedMoney(db *gorm.DB, SenderID uint) ([]models.MoneyTransaction, error) {
	var moneyTransactions []models.MoneyTransaction

	res := db.Where("from_user_id = ?", SenderID).
		Preload("FromUser").
		Preload("ToUser").
		Find(&moneyTransactions)

	return moneyTransactions, res.Error
}

func RecievedMoney(db *gorm.DB, RecieverID uint) ([]models.MoneyTransaction, error) {
	var moneyTransactions []models.MoneyTransaction

	res := db.Where("to_user_id = ?", RecieverID).
		Preload("FromUser").
		Preload("ToUser").
		Find(&moneyTransactions)

	return moneyTransactions, res.Error
}

func BuyItem(db *gorm.DB, BuyerID uint, ProductName string) error {
	price, err := ProductPrice(ProductName)

	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var buyer models.User

		if err := tx.First(&buyer, BuyerID).Error; err != nil {
			return errors.New("buyer not found")
		}

		if buyer.Money < price {
			return errors.New("buyer haven't enough money")
		}

		if err := tx.Create(&models.BoughtItem{
			BuyerID:  BuyerID,
			ItemName: ProductName,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&buyer).Update("Money", buyer.Money-price).Error; err != nil {
			return err
		}

		return nil
	})
}
