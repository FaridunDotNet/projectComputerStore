package main

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "qwerty"
	DBName     = "academy_db"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")

	//h := handler{DB: db}

	if err = db.AutoMigrate(
		&models.Admin{},
		&models.Product{},
		&models.Order{},
		&models.Category{},
		&models.Customer{},
		&models.Review{},
		&models.OrderDetail{},
	); err != nil {
		log.Fatal("cannot migrate: ", err.Error())
	}

	h := handlers.Handlers{
		DB: db,
	}

	router := gin.Default()

	admins := router.Group("/admins")
	{
		admins.GET("/", h.GetAllAdmins) // #inprogress
		admins.POST("/", h.CreateAdmin) // #inprogress
		admins.GET("/:id", h.GetOneAdmin)
		admins.PUT("/:id", h.UpdateAdmin)
		admins.DELETE("/:id", h.DeleteAdmin)
	}

	categories := router.Group("/categories")
	{
		categories.GET("/", h.GetAllCategories)
		categories.POST("/", h.CreateCategory)
		categories.GET("/:id", h.GetOneCategory)
		categories.PUT("/:id", h.UpdateCategory)
		categories.DELETE("/:id", h.DeleteCategory)
	}

	products := router.Group("/products")
	{
		products.GET("/", h.GetAllProducts)
		products.POST("/", h.CreateProduct)
		products.GET("/:id", h.GetOneProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
	}

	customers := router.Group("/customers")
	{

		customers.GET("/", h.GetAllCustomers)
		customers.POST("/", h.CreateCustomer)
		customers.GET("/:id", h.GetOneCustomer)
		customers.PUT("/:id", h.UpdateCustomer)
		customers.DELETE("/:id", h.DeleteCustomer)
	}

	orders := router.Group("/orders")
	{
		orders.GET("/", h.GetAllOrders)
		orders.POST("/", h.CreateOrder)
		orders.GET("/:id", h.GetOneOrder)
		orders.PUT("/:id", h.UpdateOrder)
		orders.DELETE("/:id", h.DeleteOrder)
	}

	ordersDetails := router.Group("/order_details")
	{
		ordersDetails.GET("/", h.GetAllOrderDetails)
		ordersDetails.POST("/", h.CreateOrderDetail)
		ordersDetails.GET("/:id", h.GetOneOrderDetail)
		ordersDetails.PUT("/:id", h.UpdateOrderDetail)
		ordersDetails.DELETE("/:id", h.DeleteOrderDetail)
	}

	reviews := router.Group("/order_details")
	{
		reviews.GET("/", h.GetAllReviews)
		reviews.POST("/", h.CreateReview)
		reviews.GET("/:id", h.GetOneReview)
		reviews.PUT("/:id", h.UpdateReview)
		reviews.DELETE("/:id", h.DeleteReview)
		reviews.GET("/id", h.GetProductWithReviews)
	}

	router.Run(":4000")
}
