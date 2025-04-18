package models

import "gorm.io/gorm"

type MoneyTransaction struct {
	gorm.Model
	FromUserID uint `gorm:"not null;index"`
	ToUserID   uint `gorm:"not null;index"`
	Value      uint `gorm:"not null"`

	FromUser User `gorm:"foreignKey:FromUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // связи с пользователями
	ToUser   User `gorm:"foreignKey:ToUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
