package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Money    uint   `gorm:"default:1000;not null"`

	MoneyTransactionsSent     []MoneyTransaction `gorm:"foreignKey:FromUserID"`
	MoneyTransactionsReceived []MoneyTransaction `gorm:"foreignKey:ToUserID"`
}
