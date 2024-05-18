package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ProductID  uint      `json:"product_id" binding:"required"`
	Product    Product   `json:"product"`
	CustomerID uint      `json:"customer_id" binding:"required"`
	Customer   Customer  `json:"customer"`
	Rating     int       `json:"rating" binding:"required,gte=1,lte=10"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time `gorm:"default:now()"`
}
