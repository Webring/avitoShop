package models

type User struct {
	Username string `gorm:"primaryKey;size:100;not null"`
	Password string `gorm:"not null"`
	Money    uint   `gorm:"default:1000;not null"`

	MoneyTransactionsSent     []MoneyTransaction `gorm:"foreignKey:FromUsername;references:Username"`
	MoneyTransactionsReceived []MoneyTransaction `gorm:"foreignKey:ToUsername;references:Username"`
	BoughtItems               []BoughtItem       `gorm:"foreignKey:BuyerUsername;references:Username"`
}
