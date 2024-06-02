package main

import (
	"apiAcademy/internal/database"
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/database/seeder"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	fmt.Println("Academy API Seeder\n")
	fmt.Println("Connecting to DB...")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := database.GormConnect()
	if err != nil {
		logger.Error("cannot connect to DB via gorm", "err", err.Error())
		return
	}

	db.Migrator().DropTable(&models.Customer{})
	db.Migrator().DropTable(&models.Product{})
	db.Migrator().DropTable(&models.Order{})
	db.Migrator().DropTable(&models.OrderDetail{})
	db.Migrator().DropTable(&models.Category{})
	db.Migrator().DropTable(&models.Review{})
	db.Migrator().DropTable(&models.Admin{})

	if err := database.GormAutoMigrate(db); err != nil {
		logger.Error("cannot automigrate models", "err", err.Error())
		return
	}

	// Admins
	if err := seeder.SeedAdmins(db, 10); err != nil {
		logger.Error("cannot seed admins", "err", err.Error())
	}

	// Teachers
	if err := seeder.SeedCustomers(db, 10); err != nil {
		logger.Error("cannot seed customers", "err", err.Error())
	}

	// Courses
	if err := seeder.SeedCategories(db, 10); err != nil {
		logger.Error("cannot seed categories", "err", err.Error())
	}

	// Groups
	if err := seeder.SeedProducts(db, 10); err != nil {
		logger.Error("cannot seed products", "err", err.Error())
	}

	// Students
	if err := seeder.SeedOrders(db, 100); err != nil {
		logger.Error("cannot seed orders", "err", err.Error())
	}

	// Lessons
	if err := seeder.SeedOrdersDetails(db, 10); err != nil {
		logger.Error("cannot seed details", "err", err.Error())
	}

	if err := seeder.SeedReview(db, 10); err != nil {
		logger.Error("cannot seed reviews", "err", err.Error())
	}
}
