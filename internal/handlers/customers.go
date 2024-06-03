package handlers

import (
	"apiAcademy/internal/database/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (h *Handlers) GetAllCustomers(c *gin.Context) {
	var customers []models.Customer
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := (page - 1) * pageSize

	if err := h.DB.Limit(pageSize).Offset(offset).Find(&customers).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
	})
}

func (h *Handlers) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Create(&customer).Error; err != nil {
		log.Println("Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handlers) GetOneCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if err := h.DB.Preload("Orders").First(&customer, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handlers) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	// Поиск клиента по ID
	if err := h.DB.First(&customer, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	// Привязка JSON-данных к клиенту
	if err := c.ShouldBindJSON(&customer); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Сохранение обновленного клиента
	if err := h.DB.Save(&customer).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handlers) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	// Поиск клиента по ID
	if err := h.DB.First(&customer, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Customer not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	// Удаление клиента
	if err := h.DB.Delete(&customer).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}
