package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Name              string    `json:"name" binding:"required"`
	Description       string    `json:"description"`
	Price             float64   `json:"price" binding:"required"`
	QuantityAvailable int       `json:"quantity_available" binding:"required"`
	CategoryID        uint      `json:"category_id" binding:"required"`
	CreatedAt         time.Time `gorm:"default:now()"`
	UpdatedAt         time.Time `gorm:"default:now()"`

	Category    Category      `json:"category"`
	OrderDetail []OrderDetail `json:"orderDetail" gorm:"foreignKey:ProductID"`
	Review      []Review      `json:"reviews" gorm:"foreignKey:ProductID"`
	// course has many groups
}
