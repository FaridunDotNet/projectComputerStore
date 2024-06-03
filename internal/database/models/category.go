package models

import (
	"time"
)

type Category struct {
	ID        int       `json:"id" gorm:"primaryKey" fake:"-"`
	Name      string    `json:"name" binding:"required" fake:"{categoryName}"`
	CreatedAt time.Time `gorm:"default:now()" fake:"-"`
	UpdatedAt time.Time `gorm:"default:now()" fake:"-"`

	Product []Product `json:"products" gorm:"foreignKey:CategoryID" fake:"-"`
}
