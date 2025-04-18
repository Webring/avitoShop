package services

import (
	"AvitoShop/internal/models"
	"gorm.io/gorm"
)

type InventoryDTO struct {
	ItemName string `json:"name"`
	Amount   uint   `json:"amount"`
}

func UserInventory(db *gorm.DB, Username string) ([]InventoryDTO, error) {
	var items []InventoryDTO
	db.Model(&models.BoughtItem{}).
		Where("buyer_username = ?", Username).
		Select("item_name, COUNT(item_name) as amount").
		Group("item_name").
		Find(&items)

	return items, nil
}
