package models

import (
	"time"
)

type MoneyTransaction struct {
	CreatedAt time.Time `gorm:"autoCreateTime"` // Только время создания

	FromUsername string `gorm:"not null;index"`
	ToUsername   string `gorm:"not null;index"`
	Value        uint   `gorm:"not null"`

	FromUser User `gorm:"foreignKey:FromUsername;references:Username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ToUser   User `gorm:"foreignKey:ToUsername;references:Username;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
