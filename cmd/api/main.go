package main

import (
	"apiAcademy/internal/database"
	"apiAcademy/internal/handlers"
	"apiAcademy/internal/helpers"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := database.GormConnect()
	if err != nil {
		logger.Error("cannot connect to DB via gorm", "err", err.Error())
		return
	}

	if err := database.GormAutoMigrate(db); err != nil {
		logger.Error("cannot automigrate models", "err", err.Error())
		return
	}

	h := handlers.Handlers{
		DB: db,
	}

	helpers.SetValidatorEngineToUseJSONTags()

	router := gin.Default() //Для восстанавления из паники

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

	reviews := router.Group("/reviews")
	{
		reviews.GET("/", h.GetAllReviews)
		reviews.POST("/", h.CreateReview)
		reviews.GET("/:id", h.GetOneReview)
		reviews.PUT("/:id", h.UpdateReview)
		reviews.DELETE("/:id", h.DeleteReview)
	}

	router.Run(":4000")
}
