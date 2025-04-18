package services

import (
	"AvitoShop/internal/models"
	"errors"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func SendMoney(db *gorm.DB, fromUsername, toUsername string, value uint) error {
	if value == 0 || fromUsername == toUsername {
		log.Warnf("Useless send money operation: from %s to %s with value %d", fromUsername, toUsername, value)
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var fromUser, toUser models.User

		if err := tx.First(&fromUser, "username = ?", fromUsername).Error; err != nil {
			return errors.New("sender not found")
		}

		if err := tx.First(&toUser, "username = ?", toUsername).Error; err != nil {
			return errors.New("receiver not found")
		}

		if fromUser.Money < value {
			return errors.New("sender doesn't have enough money")
		}

		if err := tx.Create(&models.MoneyTransaction{
			FromUsername: fromUsername,
			ToUsername:   toUsername,
			Value:        value,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&fromUser).Update("money", fromUser.Money-value).Error; err != nil {
			return err
		}

		if err := tx.Model(&toUser).Update("money", toUser.Money+value).Error; err != nil {
			return err
		}

		return nil
	})
}

func SendedMoney(db *gorm.DB, senderUsername string) ([]models.MoneyTransaction, error) {
	var moneyTransactions []models.MoneyTransaction

	res := db.Where("from_username = ?", senderUsername).
		Preload("FromUser").
		Preload("ToUser").
		Find(&moneyTransactions)

	return moneyTransactions, res.Error
}

func RecievedMoney(db *gorm.DB, receiverUsername string) ([]models.MoneyTransaction, error) {
	var moneyTransactions []models.MoneyTransaction

	res := db.Where("to_username = ?", receiverUsername).
		Preload("FromUser").
		Preload("ToUser").
		Find(&moneyTransactions)

	return moneyTransactions, res.Error
}

func BuyItem(db *gorm.DB, buyerUsername string, productName string) error {
	price, err := ProductPrice(productName)
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var buyer models.User

		if err := tx.First(&buyer, "username = ?", buyerUsername).Error; err != nil {
			return errors.New("buyer not found")
		}

		if buyer.Money < price {
			return errors.New("buyer doesn't have enough money")
		}

		if err := tx.Create(&models.BoughtItem{
			BuyerUsername: buyerUsername,
			ItemName:      productName,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&buyer).Update("money", buyer.Money-price).Error; err != nil {
			return err
		}

		return nil
	})
}
