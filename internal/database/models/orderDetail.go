package models

import (
	"gorm.io/gorm"
	"time"
)

type OrderDetail struct {
	gorm.Model
	OrderID         uint      `json:"order_id" binding:"required"`
	Order           Order     `json:"order"`
	ProductID       uint      `json:"product_id" binding:"required"`
	Product         Product   `json:"product"`
	QuantityOrdered int       `json:"quantity_ordered" binding:"required"`
	UnitPrice       float64   `json:"unit_price" binding:"required"`
	CreatedAt       time.Time `gorm:"default:now()"`
	UpdatedAt       time.Time `gorm:"default:now()"`
}
