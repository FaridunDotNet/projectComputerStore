package models

import (
	"time"
)

type Customer struct {
	ID          uint      `json:"id" gorm:"primaryKey" fake:"-"`
	FullName    string    `json:"full_name" binding:"required" fake:"{name}"`
	Email       string    `json:"email" binding:"required" fake:"{email}"`
	Phone       string    `json:"phone" binding:"required" fake:"{phone}"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required" gorm:"not null" fake:"{date}"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" fake:"-"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime" fake:"-"`
	Orders      []Order   `json:"orders" gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" fake:"-"`
}
