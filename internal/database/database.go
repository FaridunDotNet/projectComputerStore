package database

import (
	"apiAcademy/internal/database/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "qwerty"
	DBName     = "academy_db"
)

func GormConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GormAutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Admin{},
		&models.Category{},
		&models.Customer{},
		&models.Order{},
		&models.Product{},
		&models.OrderDetail{},
		&models.Review{},
	)
}
