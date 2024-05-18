package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	DateOfBirth time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:now()"`
	UpdatedAt   time.Time `gorm:"default:now()"`

	Order []Order `json:"orders" gorm:"foreignKey:CustomerID"`
}
