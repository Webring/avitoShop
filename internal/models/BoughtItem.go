package models

import "gorm.io/gorm"

type BoughtItem struct {
	gorm.Model
	BuyerID  uint   `gorm:"not null;index"`
	ItemName string `gorm:"not null"`

	Buyer *User `gorm:"foreignKey:BuyerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
