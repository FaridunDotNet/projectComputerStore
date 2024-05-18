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
		&models.Course{},
		&models.Student{},
		&models.Teacher{},
		&models.Group{},
		&models.Lesson{},
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

	teachers := router.Group("/teachers")
	{
		teachers.GET("/", h.GetAllCategories)
		teachers.POST("/", h.CreateCategory)
		teachers.GET("/:id", h.GetOneCategory)
		teachers.PUT("/:id", h.UpdateCategory)
		teachers.DELETE("/:id", h.DeleteCategory)
	}

	courses := router.Group("/courses")
	{
		courses.GET("/", h.GetAllProducts)
		courses.POST("/", h.CreateProduct)
		courses.GET("/:id", h.GetOneProduct)
		courses.PUT("/:id", h.UpdateProduct)
		courses.DELETE("/:id", h.DeleteProduct)
	}

	lessons := router.Group("/lessons")
	{

		lessons.GET("/", h.GetAllLessons)
		lessons.POST("/", h.CreateLesson)
		lessons.GET("/:id", h.GetOneLesson)
		lessons.PUT("/:id", h.UpdateLesson)
		lessons.DELETE("/:id", h.DeleteLesson)
	}

	students := router.Group("/students")
	{
		students.GET("/", h.GetAllStudents)
		students.POST("/", h.CreateStudent)
		students.GET("/:id", h.GetOneStudent)
		students.PUT("/:id", h.UpdateStudent)
		students.DELETE("/:id", h.DeleteStudent)
	}
	router.Run(":4000")
}
