package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`

	Product []Product `json:"products" gorm:"foreignKey:CategoryID"`
}
