package models

import (
	"time"
)

type Review struct {
	ID         uint      `json:"id" gorm:"primaryKey" fake:"-"`
	ProductID  uint      `json:"product_id" binding:"required" fake:"-"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" fake:"-"`
	CustomerID uint      `json:"customer_id" binding:"required"`
	Customer   Customer  `json:"customer" gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" fake:"-"`
	Rating     int       `json:"rating" binding:"required" fake:"{number:1-10}"`
	Comment    string    `json:"comment" fake:"{comment}"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime" fake:"-"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime" fake:"-"`
}
