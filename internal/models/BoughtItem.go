package models

import (
	"time"
)

type BoughtItem struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`

	BuyerUsername string `gorm:"not null;index"`
	ItemName      string `gorm:"not null"`

	Buyer *User `gorm:"foreignKey:BuyerUsername;references:Username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
