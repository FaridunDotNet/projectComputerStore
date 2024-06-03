package models

import (
	"time"
)

type OrderDetail struct {
	ID              uint      `json:"id" gorm:"primaryKey" fake:"-"`
	OrderID         uint      `json:"order_id" binding:"required" fake:"-"`
	Order           Order     `json:"order" gorm:"foreignKey:OrderID;references:ID" fake:"-"`
	ProductID       uint      `json:"product_id" binding:"required" fake:"-"`
	Product         Product   `json:"product" gorm:"foreignKey:ProductID;references:ID" fake:"-"`
	QuantityOrdered int       `json:"quantity_ordered" binding:"required" fake:"-"`
	UnitPrice       float64   `json:"unit_price" binding:"required" fake:"number:1,10"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime" fake:"-"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime" fake:"-"`
}
