package models

import (
	"time"
)

type Product struct {
	ID                uint          `json:"id" gorm:"primaryKey" fake:"-"`
	Name              string        `json:"name" binding:"required" fake:"{name}"`
	Description       string        `json:"description" binding:"required" fake:"{sentence}"`
	Price             uint          `json:"price" binding:"required" fake:"{number:1,100}"`
	QuantityAvailable int           `json:"quantity_available" binding:"required" fake:"{number:1,100}"`
	CategoryID        uint          `json:"category_id" binding:"required" fake:"{number: 1,10}"`
	Category          Category      `json:"category" fake:"-"`
	CreatedAt         time.Time     `json:"created_at" gorm:"autoCreateTime" fake:"-"`
	UpdatedAt         time.Time     `json:"updated_at" gorm:"autoUpdateTime" fake:"-"`
	OrderDetails      []OrderDetail `json:"order_details" gorm:"foreignKey:ProductID" fake:"-"`
	Reviews           []Review      `json:"reviews" gorm:"foreignKey:ProductID" fake:"-"`
}
