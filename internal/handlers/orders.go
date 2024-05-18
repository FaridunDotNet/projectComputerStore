package handlers

import (
	"apiAcademy/internal/database/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (h *Handlers) GetAllOrders(c *gin.Context) {
	orders := []models.Order{}
	if err := h.DB.Find(&orders).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func (h *Handlers) CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Create(&order).Error; err != nil {
		log.Println("Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handlers) GetOneOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := h.DB.First(&order, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Order not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, order)

}

func (h *Handlers) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := h.DB.First(&order, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Order not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Save(&order).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, order)

}

func (h *Handlers) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := h.DB.First(&order, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Order not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	if err := h.DB.Delete(&order).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}
