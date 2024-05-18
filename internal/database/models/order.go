package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	CustomerID  uint      `json:"customer_id" binding:"required"`
	OrderDate   string    `json:"order_date" binding:"required"`
	TotalAmount float64   `json:"total_amount" binding:"required"`
	CreatedAt   time.Time `gorm:"default:now()"`
	UpdatedAt   time.Time `gorm:"default:now()"`

	OrderDetail []OrderDetail `json:"orderDetails" gorm:"foreignKey:OrderID"`
}
