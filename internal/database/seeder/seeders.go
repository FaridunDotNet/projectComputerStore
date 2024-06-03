package seeder

import (
	"apiAcademy/internal/database/models"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func SeedAdmins(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var admin models.Admin

		if err := gofakeit.Struct(&admin); err != nil {
			return err
		}

		if err := db.Create(&admin).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedCustomers(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Customer

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedCategories(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Category

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedProducts(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Product

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedOrders(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.Order

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedOrdersDetails(db *gorm.DB, number int) error {
	for i := 0; i < number; i++ {
		var model models.OrderDetail

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedReview(db *gorm.DB, number int) error {
	for i := 0; i < 10; i++ {
		var model models.Review

		if err := gofakeit.Struct(&model); err != nil {
			return err
		}

		if err := db.Create(&model).Error; err != nil {
			return err
		}
	}

	return nil
}
