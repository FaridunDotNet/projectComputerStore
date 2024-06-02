package models

import (
	"time"
)

type Order struct {
	ID          uint          `json:"id" gorm:"primaryKey" fake:"-"`
	CustomerID  uint          `json:"customer_id" binding:"required" fake:"-"`
	OrderData   string        `json:"order_data" binding:"required" fake:"{sentence}"`
	TotalAmount float64       `json:"total_amount" binding:"required" fake:"{100-100000}"`
	CreatedAt   time.Time     `json:"created_at" gorm:"autoCreateTime" fake:"-"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"autoUpdateTime" fake:"-"`
	OrderDetail []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID" fake:"-"`
}
